---
subcategory: "Optimizer"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_optimizer_categories"
sidebar_current: "docs-oci-datasource-optimizer-categories"
description: |-
  Provides the list of Categories in Oracle Cloud Infrastructure Optimizer service
---

# Data Source: oci_optimizer_categories
This data source provides the list of Categories in Oracle Cloud Infrastructure Optimizer service.

Lists the supported Cloud Advisor categories.


## Example Usage

```hcl
data "oci_optimizer_categories" "test_categories" {
	#Required
	compartment_id = var.compartment_id
	compartment_id_in_subtree = var.category_compartment_id_in_subtree

	#Optional
	name = var.category_name
	state = var.category_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `compartment_id_in_subtree` - (Required) When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned depending on the the setting of `accessLevel`.

	Can only be set to true when performing ListCompartments on the tenancy (root compartment). 
* `name` - (Optional) Optional. A filter that returns results that match the name specified.
* `state` - (Optional) A filter that returns results that match the lifecycle state specified. 


## Attributes Reference

The following attributes are exported:

* `category_collection` - The list of category_collection.

### Category Reference

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

