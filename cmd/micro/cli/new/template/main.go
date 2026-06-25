package template

var (
	MainSRV = `package main

import (
	"{{.Dir}}/handler"
	pb "{{.Dir}}/proto"

	"micro.labqa.pp.ua/v6"
	"micro.labqa.pp.ua/v6/gateway/mcp"
)

func main() {
	// Create service
	service := micro.NewService("{{lower .Alias}}",
		mcp.WithMCP(":3001"),
	)

	// Initialize service
	service.Init()

	// Register handler
	pb.Register{{title .Alias}}Handler(service.Server(), handler.New())

	// Run service
	service.Run()
}
`

	MainSRVNoMCP = `package main

import (
	"{{.Dir}}/handler"
	pb "{{.Dir}}/proto"

	"micro.labqa.pp.ua/v6"
)

func main() {
	// Create service
	service := micro.NewService("{{lower .Alias}}")

	// Initialize service
	service.Init()

	// Register handler
	pb.Register{{title .Alias}}Handler(service.Server(), handler.New())

	// Run service
	service.Run()
}
`

	// MainNoProto is the default template: handlers are registered by
	// reflection, so the service builds and runs with no protoc toolchain.
	MainNoProto = `package main

import (
	"{{.Dir}}/handler"

	"micro.labqa.pp.ua/v6"
	"micro.labqa.pp.ua/v6/gateway/mcp"
	log "micro.labqa.pp.ua/v6/logger"
)

func main() {
	// Create service
	service := micro.NewService("{{lower .Alias}}",
		mcp.WithMCP(":3001"),
	)

	// Initialize service
	service.Init()

	// Register handler (reflection-based — no protoc required)
	if err := service.Handle(handler.New()); err != nil {
		log.Fatal(err)
	}

	// Run service
	service.Run()
}
`

	MainNoProtoNoMCP = `package main

import (
	"{{.Dir}}/handler"

	"micro.labqa.pp.ua/v6"
	log "micro.labqa.pp.ua/v6/logger"
)

func main() {
	// Create service
	service := micro.NewService("{{lower .Alias}}")

	// Initialize service
	service.Init()

	// Register handler (reflection-based — no protoc required)
	if err := service.Handle(handler.New()); err != nil {
		log.Fatal(err)
	}

	// Run service
	service.Run()
}
`
)
