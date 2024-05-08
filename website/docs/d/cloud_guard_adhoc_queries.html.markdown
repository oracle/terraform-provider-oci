---
subcategory: "Cloud Guard"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_guard_adhoc_queries"
sidebar_current: "docs-oci-datasource-cloud_guard-adhoc_queries"
description: |-
  Provides the list of Adhoc Queries in Oracle Cloud Infrastructure Cloud Guard service
---

# Data Source: oci_cloud_guard_adhoc_queries
This data source provides the list of Adhoc Queries in Oracle Cloud Infrastructure Cloud Guard service.

Returns a list of all adhoc queries (AdhocQuery resources) for a compartment
identified by compartmentId. List is returned in a AdhocQueryCollection resource
with page of AdhocQuerySummary resources.

The ListAdhocQueries operation returns only the adhoc queries in 'compartmentId' passed.
The list does not include any subcompartments of the compartmentId passed.

The parameter `accessLevel` specifies whether to return only those compartments for which the
requestor has INSPECT permissions on at least one resource directly
or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
Principal doesn't have access to even one of the child compartments. This is valid only when
`compartmentIdInSubtree` is set to `true`.

The parameter `compartmentIdInSubtree` applies when you perform ListAdhocQueries on the
`compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
To get a full list of all compartments and subcompartments in the tenancy (root compartment),
set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.


## Example Usage

```hcl
data "oci_cloud_guard_adhoc_queries" "test_adhoc_queries" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	access_level = var.adhoc_query_access_level
	adhoc_query_status = var.adhoc_query_adhoc_query_status
	compartment_id_in_subtree = var.adhoc_query_compartment_id_in_subtree
	time_ended_filter_query_param = var.adhoc_query_time_ended_filter_query_param
	time_started_filter_query_param = var.adhoc_query_time_started_filter_query_param
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are `RESTRICTED` and `ACCESSIBLE`. Default is `RESTRICTED`. Setting this to `ACCESSIBLE` returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to `RESTRICTED` permissions are checked and no partial results are displayed. 
* `adhoc_query_status` - (Optional) The status of the adhoc query created. Default value for state is provisioning. If no value is specified state is provisioning.
* `compartment_id` - (Required) The OCID of the compartment in which to list resources.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned depending on the setting of `accessLevel`. 
* `time_ended_filter_query_param` - (Optional) End time for a filter. If end time is not specified, end time will be set to current time.
* `time_started_filter_query_param` - (Optional) Start time for a filter. If start time is not specified, start time will be set to current time - 30 days.


## Attributes Reference

The following attributes are exported:

* `adhoc_query_collection` - The list of adhoc_query_collection.

### AdhocQuery Reference

The following attributes are exported:

* `adhoc_query_details` - Detailed information about the adhoc query.
	* `adhoc_query_resources` - Target information in which adhoc query will be run
		* `region` - Region in which adhoc query needs to be run
		* `resource_ids` - List of OCIDs on which query needs to be run
		* `resource_type` - Type of resource
	* `query` - The adhoc query expression that is run
* `adhoc_query_regional_details` - Instance level status for each region
	* `expected_count` - Expected number of instances on which query should run
	* `expired_count` - Number of instances on which query expired
	* `failed_count` - Number of instances on which query failed
	* `region` - Region name
	* `regional_error` - error message to show if adhoc query fails in a region
	* `regional_status` - adhoc query status of the region
	* `succeeded_count` - Number of instances on which query succeeded
* `compartment_id` - Compartment OCID of the adhoc query
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `error_message` - Error message to show on UI in case of failure
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`

	Avoid entering confidential information. 
* `id` - OCID for the adhoc query
* `state` - The current lifecycle state of the resource.
* `status` - Status of the adhoc query
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). System tags can be viewed by users, but can only be created by the system.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the adhoc query was created. Format defined by RFC3339.
* `time_updated` - The date and time the adhoc query was updated. Format defined by RFC3339.

