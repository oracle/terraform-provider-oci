---
subcategory: "Optimizer"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_optimizer_profile"
sidebar_current: "docs-oci-datasource-optimizer-profile"
description: |-
  Provides details about a specific Profile in Oracle Cloud Infrastructure Optimizer service
---

# Data Source: oci_optimizer_profile
This data source provides details about a specific Profile resource in Oracle Cloud Infrastructure Optimizer service.

Gets the specified profile's information. Uses the profile's OCID to determine which profile to retrieve.


## Example Usage

```hcl
data "oci_optimizer_profile" "test_profile" {
	#Required
	profile_id = oci_optimizer_profile.test_profile.id
}
```

## Argument Reference

The following arguments are supported:

* `profile_id` - (Required) The unique OCID of the profile.


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

