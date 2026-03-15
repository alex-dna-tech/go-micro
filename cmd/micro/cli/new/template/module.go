package template

var (
	Module = `module {{.Dir}}

go 1.22

require (
	micro.labqa.pp.ua/v5 latest
	github.com/golang/protobuf latest
	google.golang.org/protobuf latest
)
`
)
