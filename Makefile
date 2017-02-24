GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)

default: build

build_mocks:
	cd client && mockery -case underscore -name BareMetalClient

clean:
	rm -rf terraform-provider-baremetal
	rm -rf bin/*

fmt:
	gofmt -w $(GOFMT_FILES)

test:
	TF_ORACLE_ENV=test go test -v

test_acceptance:
	TF_ACC=1 go test -v

build: test
	go build -o terraform-provider-baremetal

cross: test_acceptance
	gox -output "./bin/{{.OS}}_{{.Arch}}/terraform-provider-baremetal"

.PHONY: clean fmt build cross test test_unit
