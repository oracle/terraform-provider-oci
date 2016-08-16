build: test
	@go build github.com/mustwin/baremetal-sdk-go

test:
	@go test -v github.com/mustwin/baremetal-sdk-go

identity_test:
	@go install github.com/mustwin/baremetal-sdk-go/test/identity
	@identity

core_test:
	@go install github.com/mustwin/baremetal-sdk-go/test/core
	@core

regression_test: test identity_test core_test


.PHONY: build test regression_test identity_test core_test
