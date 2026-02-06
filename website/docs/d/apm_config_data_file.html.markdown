---
subcategory: "Apm Config"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_apm_config_data_file"
sidebar_current: "docs-oci-datasource-apm_config-data_file"
description: |-
  Provides details about a specific Data File in Oracle Cloud Infrastructure Apm Config service
---

# Data Source: oci_apm_config_data_file
This data source provides details about a specific Data File resource in Oracle Cloud Infrastructure Apm Config service.

Retrieves the Data file with the specified name and type.


## Example Usage

```hcl
data "oci_apm_config_data_file" "test_data_file" {
	#Required
	apm_domain_id = oci_apm_apm_domain.test_apm_domain.id
	apm_type = var.data_file_apm_type
	data_file_name = oci_apm_config_data_file.test_data_file.name
}
```

## Argument Reference

The following arguments are supported:

* `apm_domain_id` - (Required) The APM Domain ID the request is intended for. 
* `apm_type` - (Required) The type of the data file.
* `data_file_name` - (Required) The name of the data file.
* `base64_encode_content` - (Optional) Encodes the downloaded content in base64.


## Attributes Reference

The following attributes are exported:


