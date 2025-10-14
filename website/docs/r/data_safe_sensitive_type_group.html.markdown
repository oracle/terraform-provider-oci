---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_sensitive_type_group"
sidebar_current: "docs-oci-resource-data_safe-sensitive_type_group"
description: |-
  Provides the Sensitive Type Group resource in Oracle Cloud Infrastructure Data Safe service
---

# oci_data_safe_sensitive_type_group
This resource provides the Sensitive Type Group resource in Oracle Cloud Infrastructure Data Safe service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/data-safe/latest/SensitiveTypeGroup

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/datasafe

Creates a new sensitive type group.


## Example Usage

```hcl
resource "oci_data_safe_sensitive_type_group" "test_sensitive_type_group" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.sensitive_type_group_description
	display_name = var.sensitive_type_group_display_name
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The OCID of the compartment where the sensitive type group should be created.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm) Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) The description of the sensitive type group.
* `display_name` - (Optional) (Updatable) The display name of the sensitive type group. The name does not have to be unique.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains the sensitive type group.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm) Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description of the sensitive type group.
* `display_name` - The display name of the sensitive type group.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the sensitive type group.
* `sensitive_type_count` - The number of sensitive types in the specified sensitive type group.
* `state` - The current state of the sensitive type group.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the sensitive type group was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `time_updated` - The date and time the sensitive type group was last updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Sensitive Type Group
	* `update` - (Defaults to 20 minutes), when updating the Sensitive Type Group
	* `delete` - (Defaults to 20 minutes), when destroying the Sensitive Type Group


## Import

SensitiveTypeGroups can be imported using the `id`, e.g.

```
$ terraform import oci_data_safe_sensitive_type_group.test_sensitive_type_group "id"
```

