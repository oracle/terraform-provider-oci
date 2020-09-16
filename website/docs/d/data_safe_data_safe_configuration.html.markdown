---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_data_safe_configuration"
sidebar_current: "docs-oci-datasource-data_safe-data_safe_configuration"
description: |-
  Provides details about a specific Data Safe Configuration in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_data_safe_configuration
This data source provides details about a specific Data Safe Configuration resource in Oracle Cloud Infrastructure Data Safe service.

Gets the details of the Data Safe configuration.

## Example Usage

```hcl
data "oci_data_safe_data_safe_configuration" "test_data_safe_configuration" {

	#Optional
	compartment_id = var.compartment_id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) A filter to return only resources that match the specified compartment OCID.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the tenancy used to enable Data Safe.
* `is_enabled` - Indicates if Data Safe is enabled.
* `state` - The current state of Data Safe configuration.
* `time_enabled` - The specific time when Data Safe configuration was enabled.
* `url` - The URL of the Data Safe service.

