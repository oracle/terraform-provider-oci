---
subcategory: "Optimizer"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_optimizer_resource_actions"
sidebar_current: "docs-oci-datasource-optimizer-resource_actions"
description: |-
  Provides the list of Resource Actions in Oracle Cloud Infrastructure Optimizer service
---

# Data Source: oci_optimizer_resource_actions
This data source provides the list of Resource Actions in Oracle Cloud Infrastructure Optimizer service.

Lists the Cloud Advisor resource actions that are supported.


## Example Usage

```hcl
data "oci_optimizer_resource_actions" "test_resource_actions" {
	#Required
	compartment_id = var.compartment_id
	compartment_id_in_subtree = var.resource_action_compartment_id_in_subtree

	#Optional
	child_tenancy_ids = var.resource_action_child_tenancy_ids
	include_organization = var.resource_action_include_organization
	include_resource_metadata = var.resource_action_include_resource_metadata
	name = var.resource_action_name
	recommendation_id = oci_optimizer_recommendation.test_recommendation.id
	recommendation_name = oci_optimizer_recommendation.test_recommendation.name
	resource_type = var.resource_action_resource_type
	state = var.resource_action_state
	status = var.resource_action_status
}
```

## Argument Reference

The following arguments are supported:

* `child_tenancy_ids` - (Optional) A list of child tenancies for which the respective data will be returned. Please note that  the parent tenancy id can also be included in this list. For example, if there is a parent P with two children A and B, to return results of only parent P and child A, this list should be populated with  tenancy id of parent P and child A. 

	If this list contains a tenancy id that isn't part of the organization of parent P, the request will  fail. That is, let's say there is an organization with parent P with children A and B, and also one  other tenant T that isn't part of the organization. If T is included in the list of  childTenancyIds, the request will fail.

	It is important to note that if you are setting the includeOrganization parameter value as true and  also populating the childTenancyIds parameter with a list of child tenancies, the request will fail. The childTenancyIds and includeOrganization should be used exclusively.

	When using this parameter, please make sure to set the compartmentId with the parent tenancy ID. 
* `compartment_id` - (Required) The OCID of the compartment.
* `compartment_id_in_subtree` - (Required) When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned depending on the the setting of `accessLevel`.

	Can only be set to true when performing ListCompartments on the tenancy (root compartment). 
* `include_organization` - (Optional) When set to true, the data for all child tenancies including the parent is returned. That is, if  there is an organization with parent P and children A and B, to return the data for the parent P, child  A and child B, this parameter value should be set to true.

	Please note that this parameter shouldn't be used along with childTenancyIds parameter. If you would like  to get results specifically for parent P and only child A, use the childTenancyIds parameter and populate the list with tenancy id of P and A.

	When using this parameter, please make sure to set the compartmentId with the parent tenancy ID. 
* `include_resource_metadata` - (Optional) Supplement additional resource information in extended metadata response.
* `name` - (Optional) Optional. A filter that returns results that match the name specified.
* `recommendation_id` - (Optional) The unique OCID associated with the recommendation.
* `recommendation_name` - (Optional) Optional. A filter that returns results that match the recommendation name specified.
* `resource_type` - (Optional) Optional. A filter that returns results that match the resource type specified.
* `state` - (Optional) A filter that returns results that match the lifecycle state specified. 
* `status` - (Optional) A filter that returns recommendations that match the status specified. 


## Attributes Reference

The following attributes are exported:

* `resource_action_collection` - The list of resource_action_collection.

### ResourceAction Reference

The following attributes are exported:

* `action` - Details about the recommended action. 
	* `description` - Text describing the recommended action.
	* `type` - The status of the resource action.
	* `url` - The URL path to documentation that explains how to perform the action.
* `category_id` - The unique OCID associated with the category.
* `compartment_id` - The OCID of the compartment.
* `compartment_name` - The name associated with the compartment.
* `estimated_cost_saving` - The estimated cost savings, in dollars, for the resource action.
* `extended_metadata` - Additional metadata key/value pairs that you provide. They serve the same purpose and functionality as fields in the `metadata` object.

	They are distinguished from `metadata` fields in that these can be nested JSON objects (whereas `metadata` fields are string/string maps only).

	For example:

	`{"CurrentShape": {"name":"VM.Standard2.16"}, "RecommendedShape": {"name":"VM.Standard2.8"}}` 
* `id` - The unique OCID associated with the resource action.
* `metadata` - Custom metadata key/value pairs for the resource action.

	**Metadata Example**

	"metadata" : { "cpuRecommendedShape": "VM.Standard1.1", "computeMemoryUtilization": "26.05734124418388", "currentShape": "VM.Standard1.2", "instanceRecommendedShape": "VM.Standard1.1", "computeCpuUtilization": "7.930035319720132", "memoryRecommendedShape": "None" } 
* `name` - The name assigned to the resource.
* `recommendation_id` - The unique OCID associated with the recommendation.
* `resource_id` - The unique OCID associated with the resource.
* `resource_type` - The kind of resource.
* `state` - The resource action's current state.
* `status` - The current status of the resource action.
* `time_created` - The date and time the resource action details were created, in the format defined by RFC3339.
* `time_status_begin` - The date and time that the resource action entered its current status. The format is defined by RFC3339.

	For example, "The status of the resource action changed from `pending` to `current(ignored)` on this date and time." 
* `time_status_end` - The date and time the current status will change. The format is defined by RFC3339.

	For example, "The current `postponed` status of the resource action will end and change to `pending` on this date and time." 
* `time_updated` - The date and time the resource action details were last updated, in the format defined by RFC3339.

