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

Lists the Cloud Advisor recommendations that are currently supported.


## Example Usage

```hcl
data "oci_optimizer_recommendations" "test_recommendations" {
	#Required
	compartment_id = var.compartment_id
	compartment_id_in_subtree = var.recommendation_compartment_id_in_subtree

	#Optional
	category_id = oci_optimizer_category.test_category.id
	category_name = oci_optimizer_category.test_category.name
	child_tenancy_ids = var.recommendation_child_tenancy_ids
	include_organization = var.recommendation_include_organization
	name = var.recommendation_name
	state = var.recommendation_state
	status = var.recommendation_status
}
```

## Argument Reference

The following arguments are supported:

* `category_id` - (Optional) The unique OCID associated with the category.
* `category_name` - (Optional) Optional. A filter that returns results that match the category name specified.
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

