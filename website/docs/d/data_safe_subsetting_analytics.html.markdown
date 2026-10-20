---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_subsetting_analytics"
sidebar_current: "docs-oci-datasource-data_safe-subsetting_analytics"
description: |-
  Provides the list of Subsetting Analytics in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_subsetting_analytics
This data source provides the list of Subsetting Analytics in Oracle Cloud Infrastructure Data Safe service.

Gets consolidated subsetting analytics data based on the specified query parameters.
If CompartmentIdInSubtreeQueryParam is specified as true, the behaviour
is equivalent to accessLevel "ACCESSIBLE" by default.


## Example Usage

```hcl
data "oci_data_safe_subsetting_analytics" "test_subsetting_analytics" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	compartment_id_in_subtree = var.subsetting_analytic_compartment_id_in_subtree
	group_by = var.subsetting_analytic_group_by
	subsetting_policy_id = oci_data_safe_subsetting_policy.test_subsetting_policy.id
	target_database_group_id = oci_data_safe_target_database_group.test_target_database_group.id
	target_id = oci_cloud_guard_target.test_target.id
	time_created_greater_than_or_equal_to = var.subsetting_analytic_time_created_greater_than_or_equal_to
	time_created_less_than = var.subsetting_analytic_time_created_less_than
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) A filter to return only resources that match the specified compartment OCID.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting. 
* `group_by` - (Optional) Attribute by which the subsetting analytics data should be grouped.
* `subsetting_policy_id` - (Optional) A filter to return only the resources that match the specified subsetting policy OCID.
* `target_database_group_id` - (Optional) A filter to return the target database group that matches the specified OCID.
* `target_id` - (Optional) A filter to return only items related to a specific target OCID.
* `time_created_greater_than_or_equal_to` - (Optional) A filter to return only the resources that were created after the specified date and time, as defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Using TimeCreatedGreaterThanOrEqualToQueryParam parameter retrieves all resources created after that date.

	**Example:** 2016-12-19T16:39:57.600Z 
* `time_created_less_than` - (Optional) Search for resources that were created before a specific date. Specifying this parameter corresponding `timeCreatedLessThan` parameter will retrieve all resources created before the specified created date, in "YYYY-MM-ddThh:mmZ" format with a Z offset, as defined by RFC 3339.

	**Example:** 2016-12-19T16:39:57.600Z 


## Attributes Reference

The following attributes are exported:

* `subsetting_analytics_collection` - The list of subsetting_analytics_collection.

### SubsettingAnalytic Reference

The following attributes are exported:

* `items` - An array of subsetting analytics summary objects
	* `dimensions` - The scope of analytics data
		* `policy_id` - The OCID of the subsetting policy
		* `target_id` - The OCID of the target database
	* `metric_name` - The name of the aggregation metric
	* `subsetting_analytic_count` - The total count for the aggregation metric
	* `time_last_subsetted` - The date and time the target database was last subsetted using a subsetting policy, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339)

