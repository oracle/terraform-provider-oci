GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)

default: build

build_mocks:
	cd client && mockery -case underscore -name BareMetalClient

clean:
	rm -rf terraform-provider-baremetal
	rm -rf bin/*

fmt:
	gofmt -w $(GOFMT_FILES)

test_acceptance_debug:
	TF_LOG=DEBUG DEBUG=true TF_ORACLE_ENV=test TF_ACC=1 go test -v -timeout 120m

test_acceptance:
	# You MUST export these variables
	# export TF_VAR_private_key_path=/Users/Mike/.ssh/oracle2
	# export TF_VAR_fingerprint=46:08:e3:7b:95:0a:d6:5f:78:24:32:87:23:3f:56:31
	# export TF_VAR_user_ocid=ocid1.user.oc1..aaaaaaaa5hpflij6krfusympeuugy2bawg25pralmnw7v4xdveysdpoxdjsk
	# export TF_VAR_tenancy_ocid=ocid1.tenancy.oc1..aaaaaaaayfzsknaowsjdlheebqsaicjddtlubq7dnwz5izbvs3vfs4xmkargta
	# export TF_VAR_compartment_id=ocid1.compartment.oc1..aaaaaaaajszpk2siudrmdhaknxvny7vktxk2dm43xpk7sa5d4vmrol2n2qsa
	# export TF_VAR_namespace=mustwin
	TF_ORACLE_ENV=test TF_ACC=1 go test -v -timeout 120m

build:
	go build -o terraform-provider-baremetal

version:
	sed -i '' -e 's/version = ".*"/version = "\
	$(shell curl -s https://api.github.com/repos/oracle/terraform-provider-baremetal/releases/latest | \
	jq -r '.tag_name')\
	"/g' version.go

release: version test_acceptance
	gox -output "./bin/{{.OS}}_{{.Arch}}/terraform-provider-baremetal"

zip:
	cd bin \
	&& zip -r windows.zip windows_386 windows_amd64 \
	&& tar -czvf darwin.tar.gz darwin_386 darwin_amd64 \
	&& tar -czvf freebsd.tar.gz freebsd_386 freebsd_amd64 freebsd_arm \
	&& tar -czvf linux.tar.gz linux_386 linux_amd64 linux_arm \
	&& tar -czvf openbsd.tar.gz openbsd_386 openbsd_amd64

.PHONY: clean fmt build release test test_unit zip version
