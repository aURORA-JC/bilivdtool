BINARY_NAME=bilivdtool
VERSION=0.0.1
GIT_VERSION=$(shell git describe --tags --abbrev=0)

all: build

build:
    go build -ldflags "-X main.version=$(VERSION)$(GIT_VERSION)" -o $(BINARY_NAME) main.go

clean:
    if exist $(BINARY_NAME) del $(BINARY_NAME)

.PHONY: all build clean