.PHONY: test clean all

gen::
	oapi-codegen -config api/server.cfg.yaml api/openapi.yaml
	oapi-codegen -config api/models.yaml api/openapi.yaml
	wire cmd/http/wire.go

test::
	go test -v ./...

run::
	go run cmd/http/http.go cmd/http/wire_gen.go
