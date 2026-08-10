# micro.mu — service definitions for `micro run`
# Format reference: internal/website/content/en/docs/guides/micro-run.md
# Run from this directory: `micro run`
#
# Conflict-free subset of examples/ for the compose stack
# (internal/docker/docker-compose.yml, micro-run container). The micro-run
# container runs this with --no-gateway; the gw container owns the gateway.
#
# Excluded examples (they collide on a shared port or service name, or need
# their own module/infra): agent-demo, mcp/{hello,documented,platform,workflow}
# (:9090/:3000/:3001 ports, "greeter"/"users" names), auth, deployment,
# grpc-interop (own modules), agent-human-input (no main.go), agent-ollama
# (needs OLLAMA_API_KEY).

# Long-running services — register in NATS, reachable through the gw gateway.
service greeter
    path ./examples/hello-world

service web
    path ./examples/web-service
    port 9090
    depends greeter

service multi-service
    path ./examples/multi-service

service contacts
    path ./examples/mcp/crud

# One-shot examples — build and run once at startup, output captured in logs.
service first-agent
    path ./examples/first-agent

service support
    path ./examples/support

service flow-durable
    path ./examples/flow-durable

service flow-loop
    path ./examples/flow-loop

service agent-durable
    path ./examples/agent-durable

service agent-x402-buyer
    path ./examples/agent-x402-buyer

service graceful-stop
    path ./examples/graceful-stop

service agent-plan-delegate
    path ./examples/agent-plan-delegate

service agent-wrap-tool
    path ./examples/agent-wrap-tool
