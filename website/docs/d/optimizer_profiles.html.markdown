---
subcategory: "Optimizer"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_optimizer_profiles"
sidebar_current: "docs-oci-datasource-optimizer-profiles"
description: |-
  Provides the list of Profiles in Oracle Cloud Infrastructure Optimizer service
---

# Data Source: oci_optimizer_profiles
This data source provides the list of Profiles in Oracle Cloud Infrastructure Optimizer service.

Lists the existing profiles.


## Example Usage

```hcl
data "oci_optimizer_profiles" "test_profiles" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	name = var.profile_name
	state = var.profile_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `name` - (Optional) Optional. A filter that returns results that match the name specified.
* `state` - (Optional) A filter that returns results that match the lifecycle state specified. 


## Attributes Reference

The following attributes are exported:

* `profile_collection` - The list of profile_collection.

### Profile Reference

The following attributes are exported:

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
* `time_created` - The date and time the profile was created, in the format defined by RFC3339.
* `time_updated` - The date and time the profile was last updated, in the format defined by RFC3339.

