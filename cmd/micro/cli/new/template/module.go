package template

var (
	Module = `module {{.Dir}}

go 1.22

require (
	micro.labqa.pp.ua/v6 {{.MicroVersion}}
	github.com/golang/protobuf latest
	google.golang.org/protobuf latest
)
`

	// ModuleNoProto is the default go.mod: no protobuf dependencies.
	// MicroVersion is the version this CLI was built from (or "latest"), so a
	// generated service tracks the framework the user is actually running.
	ModuleNoProto = `module {{.Dir}}

go 1.23

require micro.labqa.pp.ua/v6 {{.MicroVersion}}
`
)
