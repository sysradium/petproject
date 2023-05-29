GO ?= CGO_ENABLED=0 go
CMD_GO := $(foreach pb, $(wildcard cmd/*), $(pb)_cmd)
GO_LDFLAGS := -w -s


proto:: 
	buf generate

test::
	go test -v ./...

build: $(CMD_GO)

$(CMD_GO): %_cmd:
	$(GO) build -v $(GO_FLAGS) -ldflags "$(GO_LDFLAGS)" ./$*
