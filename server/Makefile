GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
BINARY_NAME=spApi
VERSION=$(shell cat ./VERSION)
BUILD_TIME=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
COMMIT=$(shell git rev-parse --short HEAD)
LDFLAGS=-ldflags "-X=main.version=$(VERSION) -X=main.buildTime=$(BUILD_TIME) -X main.gitCommit=$(COMMIT)"

all: test build

.PHONY: test build
build:
	go mod vendor;
	rm -rf ./build/;
	mkdir ./build;
	cp -r ./config.toml ./build/
	$(MAKE) -s go-build

go-build:
	@GOPATH=$(GOPATH) $(GOBUILD) $(LDFLAGS) -o ./build/$(BINARY_NAME)

test:
	$(GOTEST) -v ./...
