K6_IMAGE  := grafana/k6:latest
BASE_URL  ?= http://host.docker.internal:8080
ADMIN_USER ?= admin
ADMIN_PASS ?= micro
VUS       ?= 10
ITERATIONS ?= 50

.PHONY: run test stress dev

run:
	docker-compose up -d --force-recreate --remove-orphans --build

test:
	docker run --rm \
		-e BASE_URL=$(BASE_URL) \
		-e MICRO_ADMIN_USER=$(ADMIN_USER) \
		-e MICRO_ADMIN_PASSWORD=$(ADMIN_PASS) \
		-e VUS=$(VUS) \
		-e ITERATIONS=$(ITERATIONS) \
		-v "$(CURDIR)/k6:/scripts:ro" \
		$(K6_IMAGE) run /scripts/loadtest.js

stress:
	docker run --rm \
		-e BASE_URL=$(BASE_URL) \
		-e MICRO_ADMIN_USER=$(ADMIN_USER) \
		-e MICRO_ADMIN_PASSWORD=$(ADMIN_PASS) \
		-v "$(CURDIR)/k6:/scripts:ro" \
		$(K6_IMAGE) run /scripts/stress.js

dev:
	MICRO_REGISTRY="nats" \
	MICRO_BROKER="nats" \
	MICRO_TRANSPORT="nats" \
	MICRO_REGISTRY_ADDRESS="nats://127.0.0.1:4222" \
	MICRO_BROKER_ADDRESS="nats://127.0.0.1:4222" \
	MICRO_TRANSPORT_ADDRESS="nats://127.0.0.1:4222" \
	MICRO_AI_PROVIDER="openai" \
	MICRO_AI_BASE_URL="http://127.0.0.1:8000" \
	MICRO_AI_MODEL="any" \
	MICRO_AI_API_KEY="not-required" \
	MICRO_ADMIN_USER="admin" \
	MICRO_ADMIN_PASSWORD="micro" \
	MICRO_LOG_LEVEL="debug" \
	micro run --address :8080 --mcp-address :3000
