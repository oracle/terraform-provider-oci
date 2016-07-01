GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)

default: build

clean:
	rm -rf bin/*

fmt:
	gofmt -w $(GOFMT_FILES)

build:
	go build -o bin/terraform-baremetal

.PHONY: clean fmt build
