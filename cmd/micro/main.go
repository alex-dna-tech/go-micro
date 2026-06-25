package main

import (
	"embed"
	"micro.labqa.pp.ua/v6/cmd"

	_ "micro.labqa.pp.ua/v6/cmd/micro/a2a"
	_ "micro.labqa.pp.ua/v6/cmd/micro/api"
	_ "micro.labqa.pp.ua/v6/cmd/micro/chat"
	_ "micro.labqa.pp.ua/v6/cmd/micro/cli"
	_ "micro.labqa.pp.ua/v6/cmd/micro/cli/build"
	_ "micro.labqa.pp.ua/v6/cmd/micro/cli/deploy"
	_ "micro.labqa.pp.ua/v6/cmd/micro/flow"
	_ "micro.labqa.pp.ua/v6/cmd/micro/mcp"
	_ "micro.labqa.pp.ua/v6/cmd/micro/resource"
	_ "micro.labqa.pp.ua/v6/cmd/micro/run"
	"micro.labqa.pp.ua/v6/cmd/micro/server"
)

//go:embed web/styles.css web/main.js web/templates/*
var webFS embed.FS

var version = "5.0.0-dev"

func init() {
	server.HTML = webFS
}

func main() {
	_ = cmd.Init(
		cmd.Name("micro"),
		cmd.Version(version),
	)
}
