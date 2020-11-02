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

