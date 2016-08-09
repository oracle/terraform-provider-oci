build: test
	go build github.com/mustwin/baremtlclient

test:
	go test -v github.com/mustwin/baremtlclient

regression_test: test
	go install github.com/mustwin/baremtlclient/test/identity
	identity

.PHONY: build test regression_test
