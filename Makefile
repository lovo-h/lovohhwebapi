# The name of the executable (default is current directory name)
TARGET := lovohhwebapi
.DEFAULT_GOAL := $(TARGET)

VET_REPORT := vet.report
GO_TEST_OUT := go-test.out
GO_TEST_XML := go-test.report.xml

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
		if ! hash godep 2>/dev/null; then go get github.com/tools/godep; fi
		godep go test -v $(GOPACKAGES) | tee $(GO_TEST_OUT)
		cat $(GO_TEST_OUT) | go-junit-report > $(GO_TEST_XML)

#TODO: waiting for February 2018, go.1.10 release
#coverage:

clean:
		-rm -f $(TARGET)
		-rm -f $(VET_REPORT)
		-rm -f $(GO_TEST_OUT)
		-rm -f $(GO_TEST_XML)

vet:
		if ! hash godep 2>/dev/null; then go get github.com/tools/godep; fi
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
