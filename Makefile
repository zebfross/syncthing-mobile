BUILD_MODE?=c-shared
OUTPUT_DIR?=output
GO_BINARY?=go
BINDING_NAME?=libsyncthing_bridge
BINDING_FILE?=$(BINDING_NAME).so
BINDING_ARGS?=
BINDING_OUTPUT?=$(OUTPUT_DIR)/binding
EXTRA_LD_FLAGS?=

default: fmt test

deps:
	go mod download

test:
	go test ./lib/... -coverprofile=profile.cov -cover -short -count 1

fmt:
	go fmt ./...

clean:
	rm -rf output

binding: deps
	mkdir -p $(BINDING_OUTPUT)
	$(GO_BINARY) build -ldflags="-w -s $(EXTRA_LD_FLAGS)" -o $(BINDING_OUTPUT)/$(BINDING_FILE) -buildmode=$(BUILD_MODE) $(BINDING_ARGS) binding/main.go

proto: $(wildcard proto/lib/**/*.proto)
	go run build.go proto

include Makefile.android
include Makefile.ios
include Makefile.windows
include Makefile.gomobile
include Makefile.flatbuffers