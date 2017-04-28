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
	TF_ORACLE_ENV=test TF_ACC=1 go test -v

build: test
	go build -o terraform-provider-baremetal

cross: test_acceptance
	gox -output "./bin/{{.OS}}_{{.Arch}}/terraform-provider-baremetal"

zip:
	cd bin \
	&& zip -r windows.zip windows_386 windows_amd64 \
	&& tar -czvf darwin.tar.gz darwin_386 darwin_amd64 \
	&& tar -czvf freebsd.tar.gz freebsd_386 freebsd_amd64 freebsd_arm \
	&& tar -czvf linux.tar.gz linux_386 linux_amd64 linux_arm \
	&& tar -czvf openbsd.tar.gz openbsd_386 openbsd_amd64

.PHONY: clean fmt build cross test test_unit zip
