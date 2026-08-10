# micro.mu — service definitions for `micro run`
# Format reference: internal/website/content/en/docs/guides/micro-run.md
# Run from this directory: `micro run`

service greeter
    path ./examples/hello-world

service web
    path ./examples/web-service
    port 9090
    depends greeter
