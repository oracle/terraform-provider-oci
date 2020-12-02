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
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}` 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `is_enabled` - Indicates if Data Safe is enabled.
* `state` - The current state of Data Safe.
* `time_enabled` - The date and time Data Safe was enabled, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `url` - The URL of the Data Safe service.

