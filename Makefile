GOCC=go
GOBUILD=$(GOCC) build
BUILD_PATH=./build
BINARY=$(BUILD_PATH)/orbitdb

build:
	$(GOBUILD) -o $(BINARY) -v ./cmd/orbitdb/
