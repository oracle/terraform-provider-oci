GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)

default: build
build: ;go build -o terraform-provider-oci
clean: ;@rm -rf terraform-provider-oci  rm -rf bin/*  rm bin
fmt: ;goimports -w -local github.com/oracle/terraform-provider-oci $(GOFMT_FILES)

### `make release version=2.0.1`
release: clean
ifdef version
	sed -i '' -e 's/Version = ".*"/Version = "$(version)"/g' provider/version.go
	gox -output ./bin/{{.OS}}_{{.Arch}}/terraform-provider-oci_v$(version)
else
	@echo Err! `make release` requires a version argument 
endif

zip: 
	@cd bin; \
	zip -r windows.zip windows_386 windows_amd64; \
	tar -czvf darwin.tar.gz darwin_386 darwin_amd64; \
	tar -czvf freebsd.tar.gz freebsd_386 freebsd_amd64 freebsd_arm; \
	tar -czvf linux.tar.gz linux_386 linux_amd64 linux_arm; \
	tar -czvf openbsd.tar.gz openbsd_386 openbsd_amd64

### `make test run=TestResourceCore debug=1`
cmd := TF_ACC=1 TF_ORACLE_ENV=test go test ./provider -v -timeout 120m
ifdef run
  cmd := $(cmd) -run $(run)
endif
ifdef debug
  cmd := DEBUG=true TF_LOG=DEBUG OCI_GO_SDK_DEBUG=1 $(cmd)
endif
test: ;$(cmd)

test_print:
	@grep -ohi "Test.*$(test).*TestSuite" provider/*.go
	@grep -oh "TestAcc.*\*testing.T" provider/*.go | cut -d \( -f 1

.PHONY: build clean fmt release zip test test_print
