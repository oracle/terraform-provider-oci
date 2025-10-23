---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_sensitive_type_group_grouped_sensitive_type"
sidebar_current: "docs-oci-resource-data_safe-sensitive_type_group_grouped_sensitive_type"
description: |-
  Provides the Sensitive Type Group Grouped Sensitive Type resource in Oracle Cloud Infrastructure Data Safe service
---

# oci_data_safe_sensitive_type_group_grouped_sensitive_type
This resource provides the Sensitive Type Group Grouped Sensitive Type resource in Oracle Cloud Infrastructure Data Safe service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/data-safe/latest/SensitiveTypeGroupGroupedSensitiveType

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/datasafe


  Patches one or more sensitive types in a sensitive type group. You can use this operation to add or remove
sensitive type ids in a sensitive type group.


## Example Usage

```hcl
resource "oci_data_safe_sensitive_type_group_grouped_sensitive_type" "test_sensitive_type_group_grouped_sensitive_type" {
	#Required
	sensitive_type_group_id = oci_data_safe_sensitive_type_group.test_sensitive_type_group.id

	#Optional
	patch_operations {
		#Required
		operation = var.sensitive_type_group_grouped_sensitive_type_patch_operations_operation
		selection = var.sensitive_type_group_grouped_sensitive_type_patch_operations_selection

		#Optional
		value = var.sensitive_type_group_grouped_sensitive_type_patch_operations_value
	}
}
```

## Argument Reference

The following arguments are supported:

* `patch_operations` - (Optional) (Updatable) 
	* `operation` - (Required) (Updatable) The operation can be one of these values: `INSERT`, `MERGE`, `REMOVE`
	* `selection` - (Required) (Updatable) 
	* `value` - (Required when operation=INSERT | MERGE) (Updatable) 
* `sensitive_type_group_id` - (Required) 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `items` - List of sensitive type id summary objects present in the sensitive type group.
	* `sensitive_type_id` - The OCID of the sensitive type.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Sensitive Type Group Grouped Sensitive Type
	* `update` - (Defaults to 20 minutes), when updating the Sensitive Type Group Grouped Sensitive Type
	* `delete` - (Defaults to 20 minutes), when destroying the Sensitive Type Group Grouped Sensitive Type


## Import

SensitiveTypeGroupGroupedSensitiveTypes can be imported using the `id`, e.g.

```
$ terraform import oci_data_safe_sensitive_type_group_grouped_sensitive_type.test_sensitive_type_group_grouped_sensitive_type "sensitiveTypeGroups/{sensitiveTypeGroupId}/groupedSensitiveTypes" 
```

