---
subcategory: "Apm"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_apm_data_keys"
sidebar_current: "docs-oci-datasource-apm-data_keys"
description: |-
  Provides the list of Data Keys in Oracle Cloud Infrastructure Apm service
---

# Data Source: oci_apm_data_keys
This data source provides the list of Data Keys in Oracle Cloud Infrastructure Apm service.

Lists all Data Keys for the specified APM domain. The caller may filter the list by specifying the 'dataKeyType'
query parameter.


## Example Usage

```hcl
data "oci_apm_data_keys" "test_data_keys" {
	#Required
	apm_domain_id = oci_apm_apm_domain.test_apm_domain.id

	#Optional
	data_key_type = var.data_key_data_key_type
}
```

## Argument Reference

The following arguments are supported:

* `apm_domain_id` - (Required) The OCID of the APM domain
* `data_key_type` - (Optional) Data key type.


## Attributes Reference

The following attributes are exported:

* `data_keys` - The list of data_keys.

### DataKey Reference

The following attributes are exported:

* `name` - Name of the Data Key. The name uniquely identifies a Data Key within an APM domain. 
* `type` - Type of the Data Key.
* `value` - Value of the Data Key.

