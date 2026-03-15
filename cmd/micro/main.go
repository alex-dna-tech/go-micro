package main

import (
	"embed"
	"micro.labqa.pp.ua/v5/cmd"

	_ "micro.labqa.pp.ua/v5/cmd/micro/cli"
	_ "micro.labqa.pp.ua/v5/cmd/micro/cli/build"
	_ "micro.labqa.pp.ua/v5/cmd/micro/cli/deploy"
	_ "micro.labqa.pp.ua/v5/cmd/micro/mcp"
	_ "micro.labqa.pp.ua/v5/cmd/micro/run"
	"micro.labqa.pp.ua/v5/cmd/micro/server"
)

//go:embed web/styles.css web/main.js web/templates/*
var webFS embed.FS

var version = "5.0.0-dev"

func init() {
	server.HTML = webFS
}

func main() {
	cmd.Init(
		cmd.Name("micro"),
		cmd.Version(version),
	)
}
