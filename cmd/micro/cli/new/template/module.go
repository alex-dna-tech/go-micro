package template

var (
	Module = `module {{.Dir}}

go 1.22

require (
	micro.labqa.pp.ua latest
	github.com/golang/protobuf latest
	google.golang.org/protobuf latest
)
`
)
