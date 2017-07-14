GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)

default: fmt build

deps:
	go get github.com/kardianos/govendor

fmt:
	gofmt -w $(GOFMT_FILES)

build: test
	@go build github.com/MustWin/baremetal-sdk-go
	@go build github.com/MustWin/baremetal-sdk-go/cmd/obmc-container-clean

test:
	@DEBUG=true go test -v github.com/MustWin/baremetal-sdk-go

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

acceptance_test:
	@go test -v github.com/MustWin/baremetal-sdk-go/acceptance-test

record_acceptance_test:
	@go test -v -tags "recording all" github.com/MustWin/baremetal-sdk-go/acceptance-test

record_identity_user_test:
	@go test -v -tags "recording identity_user" github.com/MustWin/baremetal-sdk-go/acceptance-test

record_identity_group_test:
	@go test -v -tags "recording identity_group" github.com/MustWin/baremetal-sdk-go/acceptance-test

record_identity_user_group_membership_test:
	@go test -v -tags "recording identity_user_group_membership" github.com/MustWin/baremetal-sdk-go/acceptance-test

record_identity_compartment_test:
	@go test -v -tags "recording identity_compartment" github.com/MustWin/baremetal-sdk-go/acceptance-test

record_identity_policy_test:
	@go test -v -tags "recording identity_policy" github.com/MustWin/baremetal-sdk-go/acceptance-test

record_core_test:
	@go test -v -timeout 120m -tags "recording core" github.com/MustWin/baremetal-sdk-go/acceptance-test

record_core_volume_test:
	@go test -v -timeout 120m -tags "recording core_volume" github.com/MustWin/baremetal-sdk-go/acceptance-test

record_core_volume_attachment_test:
	@go test -v -timeout 120m -tags "recording core_volume_attachment" github.com/MustWin/baremetal-sdk-go/acceptance-test

record_core_instance_test:
	@go test -v -timeout 120m -tags "recording core_instance" github.com/MustWin/baremetal-sdk-go/acceptance-test

record_objectstorage_test:
	@go test -v -tags "recording objectstorage" github.com/MustWin/baremetal-sdk-go/acceptance-test

record_load_balancer_test:
	@go test -v -timeout 120m -tags "recording load_balancer" github.com/MustWin/baremetal-sdk-go/acceptance-test

record_load_balancer_listener_test:
	@go test -v -timeout 120m -tags "recording load_balancer_listener" github.com/MustWin/baremetal-sdk-go/acceptance-test

record_load_balancer_backendset_test:
	@go test -v -timeout 120m -tags "recording load_balancer_backendset" github.com/MustWin/baremetal-sdk-go/acceptance-test

record_load_balancer_certificate_test:
	@go test -v -timeout 120m -tags "recording load_balancer_certificate" github.com/MustWin/baremetal-sdk-go/acceptance-test


.PHONY: build test regression_test identity_test core_test database_test object_test acceptance_test \
	record_identity_user_test record_identity_group_test record_identity_user_group_membership_test \
	record_core_test record_identity_policy_test
