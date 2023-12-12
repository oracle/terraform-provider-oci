---
subcategory: "Optimizer"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_optimizer_histories"
sidebar_current: "docs-oci-datasource-optimizer-histories"
description: |-
  Provides the list of Histories in Oracle Cloud Infrastructure Optimizer service
---

# Data Source: oci_optimizer_histories
This data source provides the list of Histories in Oracle Cloud Infrastructure Optimizer service.

Lists changes to the recommendations based on user activity. 
For example, lists when recommendations have been implemented, dismissed, postponed, or reactivated.


## Example Usage

```hcl
data "oci_optimizer_histories" "test_histories" {
	#Required
	compartment_id = var.compartment_id
	compartment_id_in_subtree = var.history_compartment_id_in_subtree

	#Optional
	include_resource_metadata = var.history_include_resource_metadata
	name = var.history_name
	recommendation_id = oci_optimizer_recommendation.test_recommendation.id
	recommendation_name = oci_optimizer_recommendation.test_recommendation.name
	resource_type = var.history_resource_type
	state = var.history_state
	status = var.history_status
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `compartment_id_in_subtree` - (Required) When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned depending on the the setting of `accessLevel`.

	Can only be set to true when performing ListCompartments on the tenancy (root compartment). 
* `include_resource_metadata` - (Optional) Supplement additional resource information in extended metadata response.
* `name` - (Optional) Optional. A filter that returns results that match the name specified.
* `recommendation_id` - (Optional) The unique OCID associated with the recommendation.
* `recommendation_name` - (Optional) Optional. A filter that returns results that match the recommendation name specified.
* `resource_type` - (Optional) Optional. A filter that returns results that match the resource type specified.
* `state` - (Optional) A filter that returns results that match the lifecycle state specified. 
* `status` - (Optional) A filter that returns recommendations that match the status specified. 


## Attributes Reference

The following attributes are exported:

* `history_collection` - The list of history_collection.

### History Reference

The following attributes are exported:

* `items` - A collection of history summaries.
	* `action` - Details about the recommended action. 
		* `description` - Text describing the recommended action.
		* `type` - The status of the resource action.
		* `url` - The URL path to documentation that explains how to perform the action.
	* `category_id` - The unique OCID associated with the category.
	* `compartment_id` - The OCID of the compartment.
	* `compartment_name` - The name assigned to the compartment.
	* `estimated_cost_saving` - The estimated cost savings, in dollars, for the resource action.
	* `extended_metadata` - Additional metadata key/value pairs that you provide. They serve the same purpose and functionality as fields in the `metadata` object.

		They are distinguished from `metadata` fields in that these can be nested JSON objects (whereas `metadata` fields are string/string maps only).

		For example:

		`{"CurrentShape": {"name":"VM.Standard2.16"}, "RecommendedShape": {"name":"VM.Standard2.8"}}` 
	* `id` - The unique OCID associated with the recommendation history.
	* `metadata` - Custom metadata key/value pairs for the resource action.

		**Metadata Example**

		"metadata" : { "cpuRecommendedShape": "VM.Standard1.1", "computeMemoryUtilization": "26.05734124418388", "currentShape": "VM.Standard1.2", "instanceRecommendedShape": "VM.Standard1.1", "computeCpuUtilization": "7.930035319720132", "memoryRecommendedShape": "None" } 
	* `name` - The name assigned to the resource.
	* `recommendation_id` - The unique OCID associated with the recommendation.
	* `recommendation_name` - The name assigned to the recommendation.
	* `resource_action_id` - The unique OCID associated with the resource action.
	* `resource_id` - The unique OCID associated with the resource.
	* `resource_type` - The kind of resource.
	* `state` - The recommendation history's current state.
	* `status` - The current status of the resource action.
	* `time_created` - The date and time the recommendation history was created, in the format defined by RFC3339.

