---
subcategory: "Optimizer"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_optimizer_profile_level"
sidebar_current: "docs-oci-datasource-optimizer-profile_level"
description: |-
  Provides details about a specific Profile Level in Oracle Cloud Infrastructure Optimizer service
---

# Data Source: oci_optimizer_profile_level
This data source provides details about a specific Profile Level resource in Oracle Cloud Infrastructure Optimizer service.

Lists the existing profile levels.


## Example Usage

```hcl
data "oci_optimizer_profile_level" "test_profile_level" {
	#Required
	compartment_id = var.compartment_id
	compartment_id_in_subtree = var.profile_level_compartment_id_in_subtree

	#Optional
	name = var.profile_level_name
	recommendation_name = oci_optimizer_recommendation.test_recommendation.name
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `compartment_id_in_subtree` - (Required) When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned depending on the the setting of `accessLevel`.

	Can only be set to true when performing ListCompartments on the tenancy (root compartment). 
* `name` - (Optional) Optional. A filter that returns results that match the name specified.
* `recommendation_name` - (Optional) Optional. A filter that returns results that match the recommendation name specified.


## Attributes Reference

The following attributes are exported:

* `items` - A collection of profile levels.
	* `default_interval` - The default aggregation interval (in days) for profiles using this profile level. 
	* `metrics` - The metrics that will be evaluated by profiles using this profile level.
		* `name` - The name of the metric (e.g., `CpuUtilization`).
		* `statistic` - The name of the statistic (e.g., `p95`).
		* `target` - Optional. The metric value that the recommendation will target.
		* `threshold` - The threshold that must be crossed for the recommendation to appear.
	* `name` - A unique name for the profile level.
	* `recommendation_name` - The name of the recommendation this profile level applies to.
	* `time_created` - The date and time the category details were created, in the format defined by RFC3339.
	* `time_updated` - The date and time the category details were last updated, in the format defined by RFC3339.
	* `valid_intervals` - An array of aggregation intervals (in days) allowed for profiles using this profile level. 

