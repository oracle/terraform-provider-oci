GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)
PACKAGE=github.com/oracle/bmcs-go-sdk

default: fmt build
build: test
	@go build ${PACKAGE}
	@go build ${PACKAGE}/cmd/obmc-container-clean
fmt: ;goimports -w -local ${PACKAGE} $(GOFMT_FILES)
test: unit_test acc_test

deps:
	go get github.com/kardianos/govendor

acceptance_cover:
	@TEST=true go test -v -coverprofile=coverage.out -coverpkg ${PACKAGE} ${PACKAGE}/acceptance-test
	@go tool cover -html coverage.out

### `make unit_test run=TestResourceCore debug=1`
cmd := go test
ifdef run
  cmd := $(cmd) -run $(run)
endif
ifdef debug
  cmd := $(cmd) -v
endif
ifdef verbose
  cmd := DEBUG=true $(cmd)
endif
unit_test: ;$(cmd)


### `make acc_test rec=identity_user debug=1 verbose=1`
cmd2 := TEST=true go test
ifdef debug
  cmd2 := $(cmd2) -v
endif
ifdef verbose
  cmd2 := DEBUG=true $(cmd2)
endif
ifdef rec
  cmd2 := $(cmd2) -timeout 120m -tags "recording $(rec)"
endif
acc_test: ;$(cmd2) ${PACKAGE}/acceptance-test

test_print:
	@grep -ohi "TestRun.*$(test).*Tests" *.go
	@grep -oh "Test.*\*testing.T" *.go | cut -d \( -f 1

.PHONY: build fmt test unit_test acc_test
