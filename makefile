BINARY_NAME=bilivdtool
VERSION=0.0.1
GIT_VERSION=$(shell git rev-parse HEAD | head -c 6)
OS=$(shell uname)

ifeq ($(OS),Linux)
	BINARY_SUFFIX=
else ifeq ($(OS),Darwin)
	BINARY_SUFFIX=
else ifeq ($(OS),CYGWIN_NT-10.0)
    BINARY_SUFFIX=.exe
else ifeq ($(OS),MINGW32_NT-6.2)
	BINARY_SUFFIX=.exe
else ifeq ($(OS),Windows_NT)
	BINARY_SUFFIX=.exe
else
	BINARY_SUFFIX=
endif

all: build

build:
	@echo "Compiling for $(OS)..."
	go build -ldflags "-X main.version=$(VERSION)-$(GIT_VERSION)" -o $(BINARY_NAME)$(BINARY_SUFFIX) main.go

clean:
	rm -f $(BINARY_NAME)$(BINARY_SUFFIX)

.PHONY: all build clean