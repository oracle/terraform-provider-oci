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
	child_tenancy_ids = var.category_child_tenancy_ids
	include_organization = var.category_include_organization
	name = var.category_name
	state = var.category_state
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
* `name` - (Optional) Optional. A filter that returns results that match the name specified.
* `state` - (Optional) A filter that returns results that match the lifecycle state specified. 


## Attributes Reference

The following attributes are exported:

* `category_collection` - The list of category_collection.

### Category Reference

The following attributes are exported:

* `compartment_id` - The OCID of the tenancy. The tenancy is the root compartment.
* `compartment_name` - The name associated with the compartment.
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

