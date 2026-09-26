package server

import (
	"go/ast"
	"go/doc"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"strings"
	"sync"
)

var (
	examplePattern = regexp.MustCompile(`@example\s+([\s\S]+?)(?:\n\s*\n|$)`)
)

// docSource caches handler source files registered via RegisterDocFS, keyed
// by base filename. The doc-comment extractor consults it before touching the
// filesystem so endpoint descriptions work in -trimpath / containerized
// builds where handler source is not on disk.
var (
	docSourcesMu sync.RWMutex
	docSources   = map[string][]byte{}
)

// RegisterDocFS makes handler source files available to the doc-comment
// extractor without reading the filesystem at runtime. Services embed their
// handler package (//go:embed *.go) and register it; the extractor matches
// files by base name against what runtime.FuncForPC reports.
func RegisterDocFS(src fs.FS) {
	if src == nil {
		return
	}
	// The walk func always returns nil, so the only error WalkDir can report
	// is a failure reading the root; there is nothing to recover from.
	_ = fs.WalkDir(src, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return nil
		}
		b, rerr := fs.ReadFile(src, path)
		if rerr != nil {
			return nil
		}
		docSourcesMu.Lock()
		docSources[filepath.Base(path)] = b
		docSourcesMu.Unlock()
		return nil
	})
}

// docSource returns the registered source for a compiled file path, if any.
// A file is matched by base name, so both -trimpath module-relative paths
// (e.g. "examples/crud/main.go") and absolute build paths resolve.
func docSource(file string) []byte {
	docSourcesMu.RLock()
	defer docSourcesMu.RUnlock()
	return docSources[filepath.Base(file)]
}

// extractMethodDoc extracts documentation from a method's Go doc comment
func extractMethodDoc(method reflect.Method, rcvrType reflect.Type) (description, example string) {
	// Get the function's source location
	fn := method.Func
	if !fn.IsValid() {
		return "", ""
	}

	pc := fn.Pointer()
	if pc == 0 {
		return "", ""
	}

	// Get the source file location
	funcForPC := runtime.FuncForPC(pc)
	if funcForPC == nil {
		return "", ""
	}

	file, _ := funcForPC.FileLine(pc)
	if file == "" {
		return "", ""
	}

	// Resolve the source: registered/embedded copy first, then the file on
	// disk (dev machines running next to their source).
	src := docSource(file)
	if src == nil {
		var err error
		src, err = os.ReadFile(file)
		if err != nil {
			return "", ""
		}
	}

	// Parse the source file
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, file, src, parser.ParseComments)
	if err != nil {
		return "", ""
	}

	// Find the receiver type name (e.g., "Users" from *Users)
	rcvrTypeName := rcvrType.Name()
	if rcvrTypeName == "" && rcvrType.Kind() == reflect.Pointer {
		rcvrTypeName = rcvrType.Elem().Name()
	}

	// Search for the method in the AST
	for _, decl := range f.Decls {
		funcDecl, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}

		// Check if this is a method (has receiver)
		if funcDecl.Recv == nil {
			continue
		}

		// Check if method name matches
		if funcDecl.Name.Name != method.Name {
			continue
		}

		// Check if receiver type matches
		if len(funcDecl.Recv.List) > 0 {
			recvTypeName := getTypeName(funcDecl.Recv.List[0].Type)
			if recvTypeName != rcvrTypeName {
				continue
			}
		}

		// Found the method! Extract its doc comment
		if funcDecl.Doc != nil {
			comment := funcDecl.Doc.Text()
			return parseComment(comment)
		}
	}

	return "", ""
}

// getTypeName extracts the type name from an AST expression
func getTypeName(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.StarExpr:
		return getTypeName(t.X)
	default:
		return ""
	}
}

// parseComment extracts description and example from a doc comment
func parseComment(comment string) (description, example string) {
	// Extract @example if present
	if match := examplePattern.FindStringSubmatch(comment); len(match) > 1 {
		example = strings.TrimSpace(match[1])
		// Remove @example section from description
		comment = examplePattern.ReplaceAllString(comment, "")
	}

	// Clean up the description
	description = strings.TrimSpace(comment)

	// Use doc.Synopsis for the first sentence if description is long
	if len(description) > 200 {
		synopsis := doc.Synopsis(description)
		if synopsis != "" {
			description = synopsis
		}
	}

	return description, example
}

// extractHandlerDocs extracts documentation for all methods of a handler
func extractHandlerDocs(handler interface{}) map[string]map[string]string {
	metadata := make(map[string]map[string]string)

	typ := reflect.TypeOf(handler)
	if typ == nil {
		return metadata
	}

	// Get the receiver type for methods
	rcvrType := typ
	if rcvrType.Kind() == reflect.Pointer {
		rcvrType = rcvrType.Elem()
	}

	// Iterate through methods
	for i := 0; i < typ.NumMethod(); i++ {
		method := typ.Method(i)

		// Skip non-exported methods
		if method.PkgPath != "" {
			continue
		}

		// Extract documentation from source
		description, example := extractMethodDoc(method, rcvrType)

		if description != "" || example != "" {
			metadata[method.Name] = make(map[string]string)
			if description != "" {
				metadata[method.Name]["description"] = description
			}
			if example != "" {
				metadata[method.Name]["example"] = example
			}
		}
	}

	return metadata
}
