PROJECT=io-benchmarks
ORGANIZATION=giantswarm

SOURCE := $(shell find . -name '*.go')
VERSION := $(shell cat VERSION)
GOPATH := $(shell pwd)/.gobuild
PROJECT_PATH := $(GOPATH)/src/github.com/$(ORGANIZATION)

.PHONY=all clean test deps $(PROJECT) install fmt

all: deps $(PROJECT)

clean:
	rm -rf $(GOPATH) $(PROJECT) 

test:
	GOPATH=$(GOPATH) go test ./...

# deps
deps: .gobuild
.gobuild:
	mkdir -p $(PROJECT_PATH)
	cd $(PROJECT_PATH) && ln -s ../../../.. $(PROJECT)

	# Fetch private packages first (so `go get` skips them later)

	#
	# Fetch public packages
	GOPATH=$(GOPATH) go get -d github.com/$(ORGANIZATION)/$(PROJECT)

# build
$(PROJECT): $(SOURCE)
	GOPATH=$(GOPATH) go build -ldflags "-X main.projectVersion $(VERSION)" -o $(PROJECT)

install: $(PROJECT)
	cp $(PROJECT) /usr/local/bin/

fmt:
	gofmt -l -w .
