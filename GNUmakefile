TEST?=./...
GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)
PKG_NAME=oci
WEBSITE_REPO=github.com/hashicorp/terraform-website

prefix := $(if $(debug),TF_LOG=DEBUG OCI_GO_SDK_DEBUG=1, )
timeout := $(if $(timeout), $(timeout), 120m)
run_regex := $(if $(run), -run $(run), )
skip_goimports_check_flag := $(if $(skip_goimports_check), -s, )

default: build

build: fmtcheck
	go install

### TODO: Fix this so that only unit tests are running
test: fmtcheck

testacc: fmtcheck
	TF_ACC=1 $(prefix) go test $(TEST) -v $(TESTARGS) $(run_regex) -timeout $(timeout)

vet:
	@echo "go vet ."
	@go vet $$(go list ./... | grep -v vendor/) ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

fmt:
	gofmt -w $(GOFMT_FILES)
	goimports -w -local github.com/oracle/terraform-provider-oci $(GOFMT_FILES)

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh' $(skip_goimports_check_flag)"

errcheck:
	@sh -c "'$(CURDIR)/scripts/errcheck.sh'"

vendor-status:
	@govendor status

test-compile:
	@if [ "$(TEST)" = "./..." ]; then \
		echo "ERROR: Set TEST to a specific package. For example,"; \
		echo "  make test-compile TEST=./$(PKG_NAME)"; \
		exit 1; \
	fi
	go test -c $(TEST) $(TESTARGS)

website:
ifeq (,$(wildcard $(GOPATH)/src/$(WEBSITE_REPO)))
	echo "$(WEBSITE_REPO) not found in your GOPATH (necessary for layouts and assets), get-ting..."
	git clone https://$(WEBSITE_REPO) $(GOPATH)/src/$(WEBSITE_REPO)
endif
	@$(MAKE) -C $(GOPATH)/src/$(WEBSITE_REPO) website-provider PROVIDER_PATH=$(shell pwd) PROVIDER_NAME=$(PKG_NAME)

website-test:
ifeq (,$(wildcard $(GOPATH)/src/$(WEBSITE_REPO)))
	echo "$(WEBSITE_REPO) not found in your GOPATH (necessary for layouts and assets), get-ting..."
	git clone https://$(WEBSITE_REPO) $(GOPATH)/src/$(WEBSITE_REPO)
	# Additional steps before registration is complete
	ln -s ../../../../ext/providers/oci/website/docs $(GOPATH)/src/$(WEBSITE_REPO)/content/source/docs/providers/oci
	ln -s ../../../ext/providers/oci/website/oci.erb $(GOPATH)/src/$(WEBSITE_REPO)/content/source/layouts/oci.erb
endif
	@$(MAKE) -C $(GOPATH)/src/$(WEBSITE_REPO) website-provider-test PROVIDER_PATH=$(shell pwd) PROVIDER_NAME=$(PKG_NAME)

## Additional OCI stuff that will need to be moved eventually
get: ;go get -u github.com/kardianos/govendor; go get golang.org/x/tools/cmd/goimports; go get github.com/mitchellh/gox

### `make update-version version=2.0.1`
update-version:
ifdef version
	sed -i -e 's/Version = ".*"/Version = "$(version)"/g' oci/version.go && rm -f oci/version.go-e
else
	@echo Err! `make update-version` requires a version argument
endif

### `make release version=2.0.1`
release: clean get
ifdef version
	sed -i -e 's/Version = ".*"/Version = "$(version)"/g' oci/version.go && rm -f oci/version.go-e
	gox -output ./bin/{{.OS}}_{{.Arch}}/terraform-provider-oci_v$(version)
	gox -output ./bin/solaris_amd64/terraform-provider-oci_v$(version) -osarch="solaris/amd64"
else
	@echo Err! `make release` requires a version argument
endif

clean: ;@rm -rf terraform-provider-oci  rm -rf bin/*  rm bin

zip:
	@cd bin; \
	zip -r windows_386.zip windows_386; \
	zip -r windows_amd64.zip windows_amd64; \
	tar -czvf darwin_386.tar.gz darwin_386; \
	tar -czvf darwin_amd64.tar.gz darwin_amd64; \
	tar -czvf freebsd_386.tar.gz freebsd_386; \
	tar -czvf freebsd_amd64.tar.gz freebsd_amd64; \
	tar -czvf freebsd_arm.tar.gz freebsd_arm; \
	tar -czvf linux_386.tar.gz linux_386; \
	tar -czvf linux_amd64.tar.gz linux_amd64; \
	tar -czvf linux_arm.tar.gz linux_arm; \
	tar -czvf openbsd_386.tar.gz openbsd_386; \
	tar -czvf openbsd_amd64.tar.gz openbsd_amd64; \
	tar -czvf solaris_amd64.tar.gz solaris_amd64

.PHONY: build test testacc vet fmt fmtcheck errcheck vendor-status test-compile website website-test

