GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)

default: build

clean:
	rm -rf terraform-provider-baremetal
	rm -rf bin/*

fmt:
	gofmt -w $(GOFMT_FILES)

test:
	go vet; go test

build: test
	go build -o terraform-provider-baremetal

cross: test
	gox -output "./bin/{{.OS}}_{{.Arch}}/terraform-provider-baremetal"

.PHONY: clean fmt build cross test
