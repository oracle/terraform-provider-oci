---
subcategory: "Optimizer"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_optimizer_profile"
sidebar_current: "docs-oci-resource-optimizer-profile"
description: |-
  Provides the Profile resource in Oracle Cloud Infrastructure Optimizer service
---

# oci_optimizer_profile
This resource provides the Profile resource in Oracle Cloud Infrastructure Optimizer service.

Creates a new profile.


## Example Usage

```hcl
resource "oci_optimizer_profile" "test_profile" {
	#Required
	compartment_id = var.compartment_id
	description = var.profile_description
	levels_configuration {

		#Optional
		items {

			#Optional
			level = var.profile_levels_configuration_items_level
			recommendation_id = oci_optimizer_recommendation.test_recommendation.id
		}
	}
	name = var.profile_name

	#Optional
	aggregation_interval_in_days = var.profile_aggregation_interval_in_days
	defined_tags = {"foo-namespace.bar-key"= "value"}
	freeform_tags = {"bar-key"= "value"}
	target_compartments {
		#Required
		items = var.profile_target_compartments_items
	}
	target_tags {
		#Required
		items {
			#Required
			tag_definition_name = var.profile_target_tags_items_tag_definition_name
			tag_namespace_name = oci_identity_tag_namespace.test_tag_namespace.name
			tag_value_type = var.profile_target_tags_items_tag_value_type

			#Optional
			tag_values = var.profile_target_tags_items_tag_values
		}
	}
}
```

## Argument Reference

The following arguments are supported:

* `aggregation_interval_in_days` - (Optional) (Updatable) The time period over which to collect data for the recommendations, measured in number of days.
* `compartment_id` - (Required) The OCID of the tenancy. The tenancy is the root compartment.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Required) (Updatable) Text describing the profile. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair applied without any predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Exists for cross-compatibility only.  Example: `{"bar-key": "value"}` 
* `levels_configuration` - (Required) (Updatable) A list of configuration levels for each recommendation.
	* `items` - (Optional) (Updatable) The array of configuration levels.
		* `level` - (Optional) (Updatable) The pre-defined profile level.
		* `recommendation_id` - (Optional) (Updatable) The unique OCID of the recommendation.
* `name` - (Required) (Updatable) The name assigned to the profile. Avoid entering confidential information.
* `target_compartments` - (Optional) (Updatable) Optional. The compartments specified in the profile override for a recommendation. 
	* `items` - (Required) (Updatable) The list of OCIDs attached to the compartments specified in the current profile override.
* `target_tags` - (Optional) (Updatable) Optional. The tags specified in the profile override for a recommendation. 
	* `items` - (Required) (Updatable) The list of tags specified in the current profile override.
		* `tag_definition_name` - (Required) (Updatable) The name you use to refer to the tag, also known as the tag key.
		* `tag_namespace_name` - (Required) (Updatable) The name of the tag namespace.
		* `tag_value_type` - (Required) (Updatable) Specifies which tag value types in the `tagValues` field result in overrides of the recommendation criteria.

			When the value for this field is `ANY`, the `tagValues` field should be empty, which enforces overrides to the recommendation for resources with any tag values attached to them.

			When the value for this field value is `VALUE`, the `tagValues` field must include a specific value or list of values. Overrides to the recommendation criteria only occur for resources that match the values in the `tagValues` fields. 
		* `tag_values` - (Optional) (Updatable) The list of tag values. The tag value is the value that the user applying the tag adds to the tag key.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `aggregation_interval_in_days` - The time period over which to collect data for the recommendations, measured in number of days.
* `compartment_id` - The OCID of the tenancy. The tenancy is the root compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - Text describing the profile. Avoid entering confidential information.
* `freeform_tags` - Simple key-value pair applied without any predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Exists for cross-compatibility only.  Example: `{"bar-key": "value"}` 
* `id` - The unique OCID of the profile.
* `levels_configuration` - A list of configuration levels for each recommendation.
	* `items` - The array of configuration levels.
		* `level` - The pre-defined profile level.
		* `recommendation_id` - The unique OCID of the recommendation.
* `name` - The name assigned to the profile. Avoid entering confidential information.
* `state` - The profile's current state.
* `target_compartments` - Optional. The compartments specified in the profile override for a recommendation. 
	* `items` - The list of OCIDs attached to the compartments specified in the current profile override.
* `target_tags` - Optional. The tags specified in the profile override for a recommendation. 
	* `items` - The list of tags specified in the current profile override.
		* `tag_definition_name` - The name you use to refer to the tag, also known as the tag key.
		* `tag_namespace_name` - The name of the tag namespace.
		* `tag_value_type` - Specifies which tag value types in the `tagValues` field result in overrides of the recommendation criteria.

			When the value for this field is `ANY`, the `tagValues` field should be empty, which enforces overrides to the recommendation for resources with any tag values attached to them.

			When the value for this field value is `VALUE`, the `tagValues` field must include a specific value or list of values. Overrides to the recommendation criteria only occur for resources that match the values in the `tagValues` fields. 
		* `tag_values` - The list of tag values. The tag value is the value that the user applying the tag adds to the tag key.
* `time_created` - The date and time the profile was created, in the format defined by RFC3339.
* `time_updated` - The date and time the profile was last updated, in the format defined by RFC3339.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Profile
	* `update` - (Defaults to 20 minutes), when updating the Profile
	* `delete` - (Defaults to 20 minutes), when destroying the Profile


## Import

Profiles can be imported using the `id`, e.g.

```
$ terraform import oci_optimizer_profile.test_profile "id"
```

