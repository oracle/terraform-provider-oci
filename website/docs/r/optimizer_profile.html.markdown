
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
	defined_tags = {"foo-namespace.bar-key"= "value"}
	freeform_tags = {"bar-key"= "value"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the tenancy. The tenancy is the root compartment.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Required) (Updatable) Text describing the profile.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair applied without any predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Exists for cross-compatibility only.  Example: `{"bar-key": "value"}` 
* `levels_configuration` - (Required) (Updatable) A list of configuration levels for each recommendation.
	* `items` - (Optional) (Updatable) The array of configuration levels.
		* `level` - (Optional) (Updatable) The pre-defined profile level.
		* `recommendation_id` - (Optional) (Updatable) The unique OCID of the recommendation.
* `name` - (Required) The name assigned to the profile.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the tenancy. The tenancy is the root compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - Text describing the profile.
* `freeform_tags` - Simple key-value pair applied without any predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Exists for cross-compatibility only.  Example: `{"bar-key": "value"}` 
* `id` - The unique OCID of the profile.
* `levels_configuration` - A list of configuration levels for each recommendation.
	* `items` - The array of configuration levels.
		* `level` - The pre-defined profile level.
		* `recommendation_id` - The unique OCID of the recommendation.
* `name` - The name assigned to the profile.
* `state` - The profile's current state.
* `time_created` - The date and time the profile was created, in the format defined by RFC3339.
* `time_updated` - The date and time the profile was last updated, in the format defined by RFC3339.

## Import

Profiles can be imported using the `id`, e.g.

```
$ terraform import oci_optimizer_profile.test_profile "id"
```

