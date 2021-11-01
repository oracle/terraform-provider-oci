---
subcategory: "Optimizer"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_optimizer_category"
sidebar_current: "docs-oci-datasource-optimizer-category"
description: |-
  Provides details about a specific Category in Oracle Cloud Infrastructure Optimizer service
---

# Data Source: oci_optimizer_category
This data source provides details about a specific Category resource in Oracle Cloud Infrastructure Optimizer service.

Gets the category that corresponds to the specified OCID.


## Example Usage

```hcl
data "oci_optimizer_category" "test_category" {
	#Required
	category_id = oci_optimizer_category.test_category.id
}
```

## Argument Reference

The following arguments are supported:

* `category_id` - (Required) The unique OCID associated with the category.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the tenancy. The tenancy is the root compartment.
* `description` - Text describing the category.
* `estimated_cost_saving` - The estimated cost savings, in dollars, for the category.
* `extended_metadata` - Additional metadata key/value pairs for the category.

	For example:

	`{"EstimatedSaving": "200"}` 
* `id` - The unique OCID of the category.
* `name` - The name assigned to the category.
* `recommendation_counts` - An array of `RecommendationCount` objects grouped by the level of importance assigned to the recommendation.
	* `count` - The count of recommendations.
	* `importance` - The level of importance assigned to the recommendation.
* `resource_counts` - An array of `ResourceCount` objects grouped by the status of the recommendation.
	* `count` - The count of resources.
	* `status` - The recommendation status of the resource.
* `state` - The category's current state.
* `time_created` - The date and time the category details were created, in the format defined by RFC3339.
* `time_updated` - The date and time the category details were last updated, in the format defined by RFC3339.

