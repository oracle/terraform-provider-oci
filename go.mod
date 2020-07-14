module github.com/terraform-providers/terraform-provider-oci

require (
	github.com/fatih/color v1.7.0
	github.com/hashicorp/hcl v0.0.0-20180404174102-ef8a98b0bbce // indirect
	github.com/hashicorp/hcl2 v0.0.0-20190618163856-0b64543c968c
	github.com/hashicorp/terraform-exec v0.2.3-0.20200714005955-7764650a3a04
	github.com/hashicorp/terraform-plugin-sdk v1.15.0
    github.com/oracle/oci-go-sdk v24.3.0+incompatible
    github.com/oracle/oci-go-sdk/v25 v25.1.0
	github.com/stretchr/objx v0.1.1 // indirect
	github.com/stretchr/testify v1.6.1
	gopkg.in/yaml.v2 v2.2.2
)

// Uncomment this line to get OCI Go SDK from local source instead of github
//replace github.com/oracle/oci-go-sdk => ../../oracle/oci-go-sdk

replace github.com/oracle/oci-go-sdk/v25 => ../../oracle/oci-go-sdk
