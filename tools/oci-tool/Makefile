
default: build

build:
	go build -o oci-tool

clean:
	rm -rf oci-tool
	rm -rf bin/*

release: clean
	gox -osarch="darwin/amd64 linux/amd64 windows/amd64" -output "./bin/{{.OS}}_{{.Arch}}/oci-tool"

zip:
	cd bin && zip -r oci-tool.zip windows_amd64 darwin_amd64 linux_amd64

.PHONY: build clean release zip
