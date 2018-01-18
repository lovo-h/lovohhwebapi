# The name of the executable (default is current directory name)
TARGET := lovohhwebapi
.DEFAULT_GOAL := $(TARGET)

VET_REPORT := vet.report

# These will be provided to the target
VERSION := 1.0.0
BUILD := $(shell git rev-parse HEAD)

# Use linker flags to provide version/build settings to the target
LDFLAGS=-ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD)"

# go source files, ignore vendor directory
SRC = $(shell find . -name '*.go' -not -path "./vendor/*")
# go packages, ignore vendor directory
GOPACKAGES = $(shell go list ./...  | grep -v /vendor/)

all: prepare $(TARGET)

prepare: clean test vet simplify

$(TARGET): $(SRC)
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a $(LDFLAGS) -o $(TARGET)

build: $(TARGET)
		@true

test: test-all

test-all:
		godep go test -v $(GOPACKAGES)

clean:
		-rm -f $(TARGET)
		-rm -f $(VET_REPORT)

vet:
		godep go vet ./... > $(VET_REPORT)

install:
		@go install $(LDFLAGS)

uninstall: clean
		-rm -f $$(which ${TARGET})

fmt:
		@gofmt -l -w $(SRC)

simplify:
		@gofmt -s -l -w $(SRC)

.PHONY: all prepare build test test-all clean vet install uninstall fmt simplify
