GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)

default: build

clean:
	rm -rf terraform-provider-baremetal
	rm -rf bin/*

fmt:
	gofmt -w $(GOFMT_FILES)

build:
	go build -o terraform-provider-baremetal

cross:
	gox -output "./bin/{{.OS}}_{{.Arch}}/terraform-provider-baremetal"

.PHONY: clean fmt build cross
