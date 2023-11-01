---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_sql_collection_analytics"
sidebar_current: "docs-oci-datasource-data_safe-sql_collection_analytics"
description: |-
  Provides the list of Sql Collection Analytics in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_sql_collection_analytics
This data source provides the list of Sql Collection Analytics in Oracle Cloud Infrastructure Data Safe service.

Retrieves a list of all SQL collection analytics in Data Safe.

The ListSqlCollectionAnalytics operation returns only the analytics for the SQL collections in the specified `compartmentId`.

The parameter `accessLevel` specifies whether to return only those compartments for which the
requestor has INSPECT permissions on at least one resource directly
or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
Principal doesn't have access to even one of the child compartments. This is valid only when
`compartmentIdInSubtree` is set to `true`.

The parameter `compartmentIdInSubtree` applies when you perform ListSqlCollections on the
`compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
To get a full list of all compartments and subcompartments in the tenancy (root compartment),
set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.


## Example Usage

```hcl
data "oci_data_safe_sql_collection_analytics" "test_sql_collection_analytics" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	access_level = var.sql_collection_analytic_access_level
	compartment_id_in_subtree = var.sql_collection_analytic_compartment_id_in_subtree
	group_by = var.sql_collection_analytic_group_by
	state = var.sql_collection_analytic_state
	target_id = oci_cloud_guard_target.test_target.id
	time_ended = var.sql_collection_analytic_time_ended
	time_started = var.sql_collection_analytic_time_started
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED. Setting this to ACCESSIBLE returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed. 
* `compartment_id` - (Required) A filter to return only resources that match the specified compartment OCID.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting. 
* `group_by` - (Optional) The group by parameter to summarize SQL collection aggregation.
* `state` - (Optional) The current state of the SQL collection.
* `target_id` - (Optional) A filter to return only items related to a specific target OCID.
* `time_ended` - (Optional) An optional filter to return the stats of the SQL collection logs collected before the date-time specified, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). 
* `time_started` - (Optional) An optional filter to return the stats of the SQL collection logs collected after the date-time specified, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). 


## Attributes Reference

The following attributes are exported:

* `sql_collection_analytics_collection` - The list of sql_collection_analytics_collection.

### SqlCollectionAnalytic Reference

The following attributes are exported:

* `items` - The aggregated data point items.
	* `dimensions` - The dimensions available for SQL collection analytics.
		* `state` - The current state of the SQL collection.
		* `target_id` - The OCID of the target corresponding to the security policy deployment.
	* `sql_collection_analytic_count` - The total count of the aggregated metric.

