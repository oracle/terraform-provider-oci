---
subcategory: "Optimizer"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_optimizer_recommendation"
sidebar_current: "docs-oci-datasource-optimizer-recommendation"
description: |-
  Provides details about a specific Recommendation in Oracle Cloud Infrastructure Optimizer service
---

# Data Source: oci_optimizer_recommendation
This data source provides details about a specific Recommendation resource in Oracle Cloud Infrastructure Optimizer service.

Gets the recommendation for the specified OCID.


## Example Usage

```hcl
data "oci_optimizer_recommendation" "test_recommendation" {
	#Required
	recommendation_id = oci_optimizer_recommendation.test_recommendation.id
}
```

## Argument Reference

The following arguments are supported:

* `recommendation_id` - (Required) The unique OCID associated with the recommendation.


## Attributes Reference

The following attributes are exported:

* `category_id` - The unique OCID associated with the category.
* `compartment_id` - The OCID of the tenancy. The tenancy is the root compartment.
* `description` - Text describing the recommendation.
* `estimated_cost_saving` - The estimated cost savings, in dollars, for the recommendation.
* `extended_metadata` - Additional metadata key/value pairs for the recommendation.

	For example:

	`{"EstimatedSaving": "200"}` 
* `id` - The unique OCID associated with the recommendation.
* `importance` - The level of importance assigned to the recommendation.
* `name` - The name assigned to the recommendation.
* `resource_counts` - An array of `ResourceCount` objects grouped by the status of the resource actions.
	* `count` - The count of resources.
	* `status` - The recommendation status of the resource.
* `state` - The recommendation's current state.
* `status` - The current status of the recommendation.
* `supported_levels` - Optional. The profile levels supported by a recommendation. For example, profile level values could be `Low`, `Medium`, and `High`. Not all recommendations support this field. 
	* `items` - The list of supported levels.
		* `name` - The name of the profile level.
* `time_created` - The date and time the recommendation details were created, in the format defined by RFC3339.
* `time_status_begin` - The date and time that the recommendation entered its current status. The format is defined by RFC3339.

	For example, "The status of the recommendation changed from `pending` to `current(ignored)` on this date and time." 
* `time_status_end` - The date and time the current status will change. The format is defined by RFC3339.

	For example, "The current `postponed` status of the recommendation will end and change to `pending` on this date and time." 
* `time_updated` - The date and time the recommendation details were last updated, in the format defined by RFC3339.

