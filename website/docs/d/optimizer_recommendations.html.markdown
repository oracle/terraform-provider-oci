---
subcategory: "Optimizer"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_optimizer_recommendations"
sidebar_current: "docs-oci-datasource-optimizer-recommendations"
description: |-
  Provides the list of Recommendations in Oracle Cloud Infrastructure Optimizer service
---

# Data Source: oci_optimizer_recommendations
This data source provides the list of Recommendations in Oracle Cloud Infrastructure Optimizer service.

Lists the Cloud Advisor recommendations that are currently supported in the specified category.


## Example Usage

```hcl
data "oci_optimizer_recommendations" "test_recommendations" {
	#Required
	category_id = oci_optimizer_category.test_category.id
	compartment_id = var.compartment_id
	compartment_id_in_subtree = var.recommendation_compartment_id_in_subtree

	#Optional
	name = var.recommendation_name
	state = var.recommendation_state
	status = var.recommendation_status
}
```

## Argument Reference

The following arguments are supported:

* `category_id` - (Required) The unique OCID associated with the category.
* `compartment_id` - (Required) The OCID of the compartment.
* `compartment_id_in_subtree` - (Required) When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned depending on the the setting of `accessLevel`.

	Can only be set to true when performing ListCompartments on the tenancy (root compartment). 
* `name` - (Optional) Optional. A filter that returns results that match the name specified.
* `state` - (Optional) A filter that returns results that match the lifecycle state specified. 
* `status` - (Optional) A filter that returns recommendations that match the status specified. 


## Attributes Reference

The following attributes are exported:

* `recommendation_collection` - The list of recommendation_collection.

### Recommendation Reference

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

