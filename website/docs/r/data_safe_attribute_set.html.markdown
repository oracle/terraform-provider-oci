---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_attribute_set"
sidebar_current: "docs-oci-resource-data_safe-attribute_set"
description: |-
  Provides the Attribute Set resource in Oracle Cloud Infrastructure Data Safe service
---

# oci_data_safe_attribute_set
This resource provides the Attribute Set resource in Oracle Cloud Infrastructure Data Safe service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/data-safe/latest/AttributeSet

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/datasafe

Creates an attribute set.


## Example Usage

```hcl
resource "oci_data_safe_attribute_set" "test_attribute_set" {
	#Required
	attribute_set_type = var.attribute_set_attribute_set_type
	attribute_set_values = var.attribute_set_attribute_set_values
	compartment_id = var.compartment_id
	display_name = var.attribute_set_display_name

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.attribute_set_description
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `attribute_set_type` - (Required) The type of attribute set.
* `attribute_set_values` - (Required) (Updatable) The list of values in an attribute set
* `compartment_id` - (Required) (Updatable) The OCID of the compartment that contains the attribute set.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm) Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) Description of the attribute set.
* `display_name` - (Required) (Updatable) The display name of the attribute set. The name is unique and changeable.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `attribute_set_type` - The type of attribute set.
* `attribute_set_values` - The list of values in an attribute set
* `compartment_id` - The OCID of the compartment where the attribute set is stored.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm) Example: `{"Operations.CostCenter": "42"}` 
* `description` - Description of an attribute set.
* `display_name` - The display name of an attribute set. The name does not have to be unique, and is changeable.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `id` - The OCID of an attribute set.
* `in_use` - Indicates whether the attribute set is in use by other resource.
* `is_user_defined` - A boolean flag indicating to list user defined or seeded attribute sets.
* `state` - The current state of an attribute set.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time an attribute set was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `time_updated` - The date and time an attribute set was updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Attribute Set
	* `update` - (Defaults to 20 minutes), when updating the Attribute Set
	* `delete` - (Defaults to 20 minutes), when destroying the Attribute Set


## Import

AttributeSets can be imported using the `id`, e.g.

```
$ terraform import oci_data_safe_attribute_set.test_attribute_set "id"
```

