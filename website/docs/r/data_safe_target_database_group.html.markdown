---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_target_database_group"
sidebar_current: "docs-oci-resource-data_safe-target_database_group"
description: |-
  Provides the Target Database Group resource in Oracle Cloud Infrastructure Data Safe service
---

# oci_data_safe_target_database_group
This resource provides the Target Database Group resource in Oracle Cloud Infrastructure Data Safe service.

Creates a new target database group.


## Example Usage

```hcl
resource "oci_data_safe_target_database_group" "test_target_database_group" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.target_database_group_display_name
	matching_criteria {
		#Required
		include {

			#Optional
			compartments {
				#Required
				id = var.target_database_group_matching_criteria_include_compartments_id

				#Optional
				is_include_subtree = var.target_database_group_matching_criteria_include_compartments_is_include_subtree
			}
			defined_tags = {"Operations.CostCenter"= "42"}
			freeform_tags = {"Department"= "Finance"}
			target_database_ids = var.target_database_group_matching_criteria_include_target_database_ids
		}

		#Optional
		exclude {
			#Required
			target_database_ids = var.target_database_group_matching_criteria_exclude_target_database_ids
		}
	}

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.target_database_group_description
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The OCID of the compartment to create the target database group.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm) Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) Description of the target database group (optional).
* `display_name` - (Required) (Updatable) The name of the target database group.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `matching_criteria` - (Required) (Updatable) Criteria to either include or exclude target databases from the target database group. These criteria can be based on compartments or tags or a list of target databases. See examples below for more details. Include: Target databases will be added to the target database group if they match at least one of the include criteria. Exclude: Target databases that will be excluded from the target database group (even if they match any of the include criteria). 
	* `exclude` - (Optional) (Updatable) Criteria to exclude certain target databases from the target database group.
		* `target_database_ids` - (Required) (Updatable) The list of target database OCIDS, that should be excluded from the target database group (even if they match some of the other criteria).
	* `include` - (Required) (Updatable) Criteria to determine whether a target database should be included in the target database group. If the database satisfies any of compartments, targetDatabaseIds, freeformTags, or definedTags criteria, it qualifies for inclusion in the target database group 
		* `compartments` - (Optional) (Updatable) List of compartment objects, each containing the OCID of the compartment and a boolean value that indicates whether the target databases in the compartments and sub-compartments should also be included in the target database group.
			* `id` - (Required) (Updatable) The OCID of the compartment for including target databases to the target database group. All target databases in the compartment will be members of the target database group.
			* `is_include_subtree` - (Optional) (Updatable) This indicates whether the target databases of sub-compartments should also be included in the target database group. By default, this parameter is set to false.
		* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm) Example: `{"Operations.CostCenter": "42"}` 
		* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
		* `target_database_ids` - (Optional) (Updatable) The list of target database OCIDs to be included in the target database group.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID for the compartment containing the target database group.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm) Example: `{"Operations.CostCenter": "42"}` 
* `description` - Description of the target database group.
* `display_name` - The name of the target database group.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the specified target database group.
* `lifecycle_details` - Details for the lifecycle status of the target database group.
* `matching_criteria` - Criteria to either include or exclude target databases from the target database group. These criteria can be based on compartments or tags or a list of target databases. See examples below for more details. Include: Target databases will be added to the target database group if they match at least one of the include criteria. Exclude: Target databases that will be excluded from the target database group (even if they match any of the include criteria). 
	* `exclude` - Criteria to exclude certain target databases from the target database group.
		* `target_database_ids` - The list of target database OCIDS, that should be excluded from the target database group (even if they match some of the other criteria).
	* `include` - Criteria to determine whether a target database should be included in the target database group. If the database satisfies any of compartments, targetDatabaseIds, freeformTags, or definedTags criteria, it qualifies for inclusion in the target database group 
		* `compartments` - List of compartment objects, each containing the OCID of the compartment and a boolean value that indicates whether the target databases in the compartments and sub-compartments should also be included in the target database group.
			* `id` - The OCID of the compartment for including target databases to the target database group. All target databases in the compartment will be members of the target database group.
			* `is_include_subtree` - This indicates whether the target databases of sub-compartments should also be included in the target database group. By default, this parameter is set to false.
		* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm) Example: `{"Operations.CostCenter": "42"}` 
		* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
		* `target_database_ids` - The list of target database OCIDs to be included in the target database group.
* `membership_count` - The number of target databases in the specified target database group.
* `membership_update_time` - Time when the members of the target database group were last changed, i.e. the list was refreshed, a target database was added or removed.
* `state` - The lifecycle status of the target database group.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - Time when the target database group was created.
* `time_updated` - Time when the target database group was last updated.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Target Database Group
	* `update` - (Defaults to 20 minutes), when updating the Target Database Group
	* `delete` - (Defaults to 20 minutes), when destroying the Target Database Group


## Import

TargetDatabaseGroups can be imported using the `id`, e.g.

```
$ terraform import oci_data_safe_target_database_group.test_target_database_group "id"
```

