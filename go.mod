module github.com/terraform-providers/terraform-provider-oci

require (
	github.com/fatih/color v1.7.0
	github.com/hashicorp/go-multierror v1.0.0
	github.com/hashicorp/hcl v0.0.0-20180404174102-ef8a98b0bbce // indirect
	github.com/hashicorp/hcl2 v0.0.0-20190618163856-0b64543c968c
	github.com/hashicorp/terraform-exec v0.13.3
	github.com/hashicorp/terraform-plugin-sdk v1.17.2
	github.com/oracle/oci-go-sdk/v56 v56.0.0
	github.com/stretchr/testify v1.7.0
	golang.org/x/mod v0.4.2
	gopkg.in/yaml.v2 v2.3.0
)

// Uncomment this line to get OCI Go SDK from local source instead of github
//replace github.com/oracle/oci-go-sdk => ../../oracle/oci-go-sdk

go 1.13
