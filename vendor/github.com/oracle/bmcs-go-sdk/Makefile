GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)
PACKAGE=github.com/oracle/bmcs-go-sdk

default: fmt build

deps:
	go get github.com/kardianos/govendor

fmt:
	goimports -w -local ${PACKAGE} $(GOFMT_FILES)

build: test
	@go build ${PACKAGE}
	@go build ${PACKAGE}/cmd/obmc-container-clean

test:
	@DEBUG=true go test -v ${PACKAGE}

acceptance_test:
	@TEST=true go test -v ${PACKAGE}/acceptance-test

acceptance_cover:
	@TEST=true go test -v -coverprofile=coverage.out -coverpkg ${PACKAGE} ${PACKAGE}/acceptance-test
	@go tool cover -html coverage.out

record_acceptance_test:
	@go test -v -timeout 120m -tags "recording all" ${PACKAGE}/acceptance-test

.PHONY: build test regression_test acceptance_test record_acceptance_test
