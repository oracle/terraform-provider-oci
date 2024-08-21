# Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.

TEST?=./...
GOFMT_FILES?=$(if $(SERVICE), $$(find . -name '$(SERVICE)*.go' |grep -v vendor), $$(find . -name '*.go' |grep -v vendor))
PKG_NAME=oci
TEST_PKG_NAME=internal/integrationtest
WEBSITE_REPO=github.com/hashicorp/terraform-website
release_date=$(shell date -v +5d +%F)

prefix := $(if $(debug),TF_LOG=DEBUG OCI_GO_SDK_DEBUG=v, )
timeout := $(if $(timeout), $(timeout), 120m)
run_regex := $(if $(run), -run $(run), )
test_tags := $(if $(tags), -tags $(tags), )
skip_goimports_check_flag := $(if $(skip_goimports_check), -s, )

## This rule will set GO mod environment variables so that builds/tests are using the vendor folder
## May need to remove this in future so that it doesn't interfere with environment settings of .travis.yml file
gomodenv:
	export GO111MODULE=on
	export GOFLAGS=-mod=vendor

default: build

## IMPORTANT: Do not modify the following `build` target. The following steps are a requirement of the provider release process.
build: fmtcheck gomodenv
	go install

buildnovcs: fmtcheck errcheck gomodenv
	go install -buildvcs=false

### TODO: Fix this so that only unit tests are running
test: fmtcheck

sweep: fmtcheck gomodenv
	@echo "WARNING: This will destroy infrastructure. Use only in development accounts."
	TF_ACC=1 $(prefix) go test $(TEST) -v -run TestMain -sweep=$(sweep) -sweep-run=$(sweep-run) -timeout $(timeout)

testacc: build
	TF_ACC=1 $(prefix) go test $(TEST) -v $(TESTARGS) $(run_regex) $(test_tags) -timeout $(timeout)

localinstall:
	mkdir -p $(GOPATH)/bin/registry.terraform.io/hashicorp/oci/1.0.0/darwin_amd64
	cp $(GOPATH)/bin/terraform-provider-oci $(HOME)/.terraform.d/plugins/registry.terraform.io/hashicorp/oci/1.0.0/darwin_amd64
	cp $(GOPATH)/bin/terraform-provider-oci $(GOPATH)/bin/registry.terraform.io/hashicorp/oci/1.0.0/darwin_amd64

vet:
	@echo "go vet ."
	@go vet $$(go list ./... | grep -v vendor/) ; if [ $$? -ne 0 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

fmt:
	gofmt -w $(GOFMT_FILES)
	goimports -w -local terraform-provider-oci $(GOFMT_FILES)
	@if [ -x "$$(command -v terraform)" ]; then \
		terraform fmt; \
	else \
		echo "No terraform command found. Not running 'terraform fmt'"; \
	fi

## IMPORTANT: Do not modify the following `fmtcheck` target. The following steps are a requirement of the provider release process.
## To add custom checks, use the `ocicheck` target instead.
fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh' $(skip_goimports_check_flag)"

ocicheck:
	@if [ -x "$$(command -v terraform)" ]; then \
		echo "==> Checking terraform formatting of files"; \
		terraform fmt -check=true || (echo "Terraform files are not appropriately formatted. Please run make fmt to format them." && exit 1); \
	else \
		echo "No terraform command found. Make sure it is installed in your PATH"; \
		exit 1; \
	fi

errcheck:
	@sh -c "'$(CURDIR)/scripts/errcheck.sh'"

test-compile:
	@if [ "$(TEST)" = "./..." ]; then \
		echo "ERROR: Set TEST to a specific package. For example,"; \
		echo "  make test-compile TEST=./$(TEST_PKG_NAME)"; \
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
endif
	@$(MAKE) -C $(GOPATH)/src/$(WEBSITE_REPO) website-provider-test PROVIDER_PATH=$(shell pwd) PROVIDER_NAME=$(PKG_NAME)

test-docscheck:
	@sh -c "'$(CURDIR)/scripts/docscheck.sh'"

## Additional OCI stuff that will need to be moved eventually
get: ;go get golang.org/x/tools/cmd/goimports; go get github.com/mitchellh/gox

### `make update-version version=2.0.1`
update-version:
ifdef version
	sed -i -e 's/ReleaseDate = ".*"/ReleaseDate = "$(release_date)"/g' internal/globalvar/version.go
	sed -i -e 's/Version = ".*"/Version = "$(version)"/g' internal/globalvar/version.go && rm -f internal/globalvar/version.go-e
	sed -i -e 's/version: ".*"/version: "$(version)"/g' ocibuild.conf && rm -f ocibuild.conf-e
else
	@echo Err! `make update-version` requires a version argument
endif

### `make release version=2.0.1`
release: clean
ifdef version
	sed -i -e 's/Version = ".*"/Version = "$(version)"/g' internal/globalvar/version.go && rm -f internal/globalvar/version.go-e
ifdef platform
	gox -output ./bin/{{.OS}}_{{.Arch}}/terraform-provider-oci_v$(version) -osarch=$(platform)
else
	gox -output ./bin/{{.OS}}_{{.Arch}}/terraform-provider-oci_v$(version)
	gox -output ./bin/solaris_amd64/terraform-provider-oci_v$(version) -osarch="solaris/amd64"
endif
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

.PHONY: build test testacc vet fmt fmtcheck errcheck test-compile website website-test test-docscheck

replace_sdk_version:
ifdef version
	sed -i -e 's/\("github.com\/oracle\/oci-go-sdk\/v\)[0-9]*/\1$(version)/g' oci/*.go && rm -f oci/*.go-e
else
	@echo Error! replace_sdk_version requires a version argument
endif

check-untagged-tests:
	@sh -c "'$(CURDIR)/scripts/check-untagged-tests.sh' -s ''$(SERVICE)"

check-module-name:
	@sh -c "'$(CURDIR)/scripts/gomodnamecheck.sh'"
