---
subcategory: "Cloud Guard"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_guard_adhoc_query"
sidebar_current: "docs-oci-datasource-cloud_guard-adhoc_query"
description: |-
  Provides details about a specific Adhoc Query in Oracle Cloud Infrastructure Cloud Guard service
---

# Data Source: oci_cloud_guard_adhoc_query
This data source provides details about a specific Adhoc Query resource in Oracle Cloud Infrastructure Cloud Guard service.

Returns an adhoc query identified by adhocQueryId.

## Example Usage

```hcl
data "oci_cloud_guard_adhoc_query" "test_adhoc_query" {
	#Required
	adhoc_query_id = oci_cloud_guard_adhoc_query.test_adhoc_query.id
}
```

## Argument Reference

The following arguments are supported:

* `adhoc_query_id` - (Required) Adhoc query OCID.


## Attributes Reference

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

