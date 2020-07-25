BIN           := $(shell basename $(PWD))
VERSION       := $(shell git describe --tags --always --dirty)
LDFLAGS       := -ldflags="-s -w -X main.ver=$(VERSION)"
IMAGE         := docker.io/pbar1/$(BIN)
.DEFAULT_GOAL := build

export CGO_ENABLED     := 0
export DOCKER_BUILDKIT := 1

name:
	@echo $(BIN)

build: clean
	GOOS=linux   GOARCH=arm64 go build -o bin/$(BIN)_linux_arm64       $(LDFLAGS) *.go
	GOOS=linux   GOARCH=amd64 go build -o bin/$(BIN)_linux_amd64       $(LDFLAGS) *.go
	GOOS=darwin  GOARCH=amd64 go build -o bin/$(BIN)_darwin_amd64      $(LDFLAGS) *.go
	GOOS=windows GOARCH=amd64 go build -o bin/$(BIN)_windows_amd64.exe $(LDFLAGS) *.go
	du -sh bin/*

image: build
	docker build . -t $(IMAGE):$(VERSION) -t $(IMAGE):latest

image-push: image
	docker push $(IMAGE):$(VERSION)
	docker push $(IMAGE):latest

clean:
	rm -rf bin
