module github.com/terraform-providers/terraform-provider-oci

require (
	github.com/fatih/color v1.7.0
	github.com/hashicorp/hcl v0.0.0-20180404174102-ef8a98b0bbce // indirect
	github.com/hashicorp/hcl2 v0.0.0-20190618163856-0b64543c968c
	github.com/hashicorp/terraform-exec v0.6.0
	github.com/hashicorp/terraform-plugin-sdk v1.15.0
	github.com/oracle/oci-go-sdk/v38 v38.1.0
	github.com/stretchr/objx v0.1.1 // indirect
	github.com/stretchr/testify v1.6.1
	golang.org/x/mod v0.3.0
	gopkg.in/yaml.v2 v2.2.2
)

// Uncomment this line to get OCI Go SDK from local source instead of github
//replace github.com/oracle/oci-go-sdk => ../../oracle/oci-go-sdk
