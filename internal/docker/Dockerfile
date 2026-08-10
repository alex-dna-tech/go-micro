# Personal build for compose containers only (internal/docker/docker-compose.yml).
# Compiles the micro CLI and protoc-gen-micro from source; the repo root is
# supplied as the `repo` additional build context (see docker-compose.yml).
FROM golang:1.26.0-alpine AS build
RUN apk add --no-cache git
COPY --from=repo . /src
WORKDIR /src
RUN CGO_ENABLED=0 go build -trimpath -o /out/micro ./cmd/micro \
 && CGO_ENABLED=0 go build -trimpath -o /out/protoc-gen-micro ./cmd/protoc-gen-micro

FROM alpine:latest
RUN apk add --no-cache git make curl protoc
COPY --from=build /out/micro /usr/local/go/bin/micro
COPY --from=build /out/protoc-gen-micro /usr/local/go/bin/protoc-gen-micro
ENV PATH=/usr/local/go/bin:$PATH
COPY --from=golang:1.26.0-alpine /usr/local/go /usr/local/go

WORKDIR /micro
EXPOSE 8080 3000
ENTRYPOINT ["/usr/local/go/bin/micro"]
CMD ["gateway"]
