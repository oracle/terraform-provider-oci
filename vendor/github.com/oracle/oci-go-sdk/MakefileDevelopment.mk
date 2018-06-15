#### Project generation setup
PROJECT_NAME=github.com/oracle/oci-go-sdk
PROJECT_PATH=$(GOPATH)/src/$(PROJECT_NAME)
REMOVE_AFTER_GENERATE=audit/audit_waiters.go objectstorage/objectstorage_waiters.go

#### Versions
#### If you are doing a release, do not forget to increment this versions
VER_MAJOR=1
VER_MINOR=7
###################

##### Integ test setup
INTEGTEST_DIR = integtest
TEST_HELPERS = test_helpers.go test_service_deps_helpers.go
TEST_FILES = $(notdir $(wildcard $(INTEGTEST_DIR)/*_integ_test.go))
TARGETS = $(patsubst %_client_integ_test.go, %, $(TEST_FILES))
TEST_TARGETS = $(patsubst %_client_integ_test.go, test-%, $(TEST_FILES))

define HELP_MESSAGE
make generate to generate the sdk
make test-[name of the package] to integ test a package
make test-all runs all integ tests
endef


.PHONY: help

export HELP_MESSAGE
help:
	@echo "$$HELP_MESSAGE"

list-test:
	@echo $(TEST_TARGETS)

test-all: build-sdk test-sdk-only $(TEST_TARGETS)

$(TEST_TARGETS): test-%:%
	@echo Testing $(INTEGTEST_DIR)/$<_client_integ_test.go
	@(cd $(INTEGTEST_DIR) && go test -v $(TEST_HELPERS) $<_client_integ_test.go)

$(TARGETS): %:integtest/%_client_integ_test.go

generate:
	@echo "Cleaning and generating sdk"
	@(cd $(PROJECT_PATH) && make clean-generate)
	PROJECT_NAME=$(PROJECT_NAME) mvn clean install
	@(cd $(PROJECT_PATH) && rm -f $(REMOVE_AFTER_GENERATE))

build-sdk:
	@echo "Building sdk"
	@(cd $(PROJECT_PATH) && make build)

test-sdk-only:
	@echo "Testing sdk common"
	@(cd $(PROJECT_PATH) && make test)


release-sdk:
	@echo "Building oci-go-sdk with major:$(VER_MAJOR) minor:$(VER_MINOR) patch:$(VER_PATCH) tag:$(VER_TAG)"
	@(cd $(PROJECT_PATH) && VER_MAJOR=$(VER_MAJOR) VER_MINOR=$(VER_MINOR) VER_PATCH=$(VER_PATCH) VER_TAG=$(VER_TAG) make release)

build: generate build-sdk

release: generate release-sdk
