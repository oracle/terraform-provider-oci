GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)

default: build

clean:
	rm -rf terraform-provider-baremetal

fmt:
	gofmt -w $(GOFMT_FILES)

build:
	go build -o terraform-provider-baremetal

.PHONY: clean fmt build
