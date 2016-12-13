build: test
	@go build github.com/MustWin/baremetal-sdk-go

test:
	@go test -v github.com/MustWin/baremetal-sdk-go

identity_test:
	@go install github.com/MustWin/baremetal-sdk-go/test/identity
	@identity

core_test:
	@go install github.com/MustWin/baremetal-sdk-go/test/core
	@core

database_test:
	@go install github.com/MustWin/baremetal-sdk-go/test/database
	@database

object_test:
	@go install github.com/MustWin/baremetal-sdk-go/test/object
	@object

regression_test: test identity_test database_test core_test object_test


.PHONY: build test regression_test identity_test core_test database_test object_test
