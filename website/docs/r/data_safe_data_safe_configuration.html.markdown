---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_data_safe_configuration"
sidebar_current: "docs-oci-resource-data_safe-data_safe_configuration"
description: |-
  Provides the Data Safe Configuration resource in Oracle Cloud Infrastructure Data Safe service
---

# oci_data_safe_data_safe_configuration
This resource provides the Data Safe Configuration resource in Oracle Cloud Infrastructure Data Safe service.

Enables Data Safe in the tenancy and region.


## Example Usage

```hcl
resource "oci_data_safe_data_safe_configuration" "test_data_safe_configuration" {
	#Required
	is_enabled = var.data_safe_configuration_is_enabled

	#Optional
	compartment_id = var.compartment_id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) (Updatable) A filter to return only resources that match the specified compartment OCID.
* `is_enabled` - (Required) (Updatable) Indicates if Data Safe is enabled.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the tenancy used to enable Data Safe.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}` 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `is_enabled` - Indicates if Data Safe is enabled.
* `state` - The current state of Data Safe.
* `time_enabled` - The date and time Data Safe was enabled, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `url` - The URL of the Data Safe service.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Data Safe Configuration
	* `update` - (Defaults to 20 minutes), when updating the Data Safe Configuration
	* `delete` - (Defaults to 20 minutes), when destroying the Data Safe Configuration


## Import

Import is not supported for this resource.

