---
subcategory: "Optimizer"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_optimizer_recommendation"
sidebar_current: "docs-oci-resource-optimizer-recommendation"
description: |-
  Provides the Recommendation resource in Oracle Cloud Infrastructure Optimizer service
---

# oci_optimizer_recommendation
This resource provides the Recommendation resource in Oracle Cloud Infrastructure Optimizer service.

Updates the recommendation that corresponds to the specified OCID.
Use this operation to implement the following actions:

  * Postpone recommendation
  * Dismiss recommendation
  * Reactivate recommendation


## Example Usage

```hcl
resource "oci_optimizer_recommendation" "test_recommendation" {
	#Required
	recommendation_id = oci_optimizer_recommendation.test_recommendation.id
	status = var.recommendation_status

	#Optional
	time_status_end = var.recommendation_time_status_end
}
```

## Argument Reference

The following arguments are supported:

* `recommendation_id` - (Required) The unique OCID associated with the recommendation.
* `status` - (Required) (Updatable) The status of the recommendation.
* `time_status_end` - (Optional) (Updatable) The date and time the current status will change. The format is defined by RFC3339.

	For example, "The current `postponed` status of the recommendation will end and change to `pending` on this date and time." 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Recommendation
	* `update` - (Defaults to 20 minutes), when updating the Recommendation
	* `delete` - (Defaults to 20 minutes), when destroying the Recommendation


## Import

Recommendations can be imported using the `id`, e.g.

```
$ terraform import oci_optimizer_recommendation.test_recommendation "id"
```

