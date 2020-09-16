---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_exadata_infrastructure_download_config_file"
sidebar_current: "docs-oci-datasource-database-exadata_infrastructure_download_config_file"
description: |-
  Provides details about a specific Exadata Infrastructure Download Config File in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_exadata_infrastructure_download_config_file
This data source provides details about a specific Exadata Infrastructure Download Config File resource in Oracle Cloud Infrastructure Database service.

Downloads the configuration file for the specified Exadata infrastructure.


## Example Usage

```hcl
data "oci_database_exadata_infrastructure_download_config_file" "test_exadata_infrastructure_download_config_file" {
	#Required
	exadata_infrastructure_id = oci_database_exadata_infrastructure.test_exadata_infrastructure.id

	#Optional
	base64_encode_content = "false"
}
```

## Argument Reference

The following arguments are supported:

* `base64_encode_content` - (Optional) Encodes the downloaded zipped config in base64. It is recommended to set this to `true` to avoid corrupting the zip file in Terraform state. The default value is `false`.
* `exadata_infrastructure_id` - (Required) The Exadata infrastructure [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `content` - content of the downloaded config file for exadata infrastructure. If `base64_encode_content` is set to `true`, then this content will be base64 encoded.

