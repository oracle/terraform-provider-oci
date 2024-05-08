---
subcategory: "Cloud Guard"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_guard_adhoc_query"
sidebar_current: "docs-oci-resource-cloud_guard-adhoc_query"
description: |-
  Provides the Adhoc Query resource in Oracle Cloud Infrastructure Cloud Guard service
---

# oci_cloud_guard_adhoc_query
This resource provides the Adhoc Query resource in Oracle Cloud Infrastructure Cloud Guard service.

Creates a AdhocQuery resource.


## Example Usage

```hcl
resource "oci_cloud_guard_adhoc_query" "test_adhoc_query" {
	#Required
	adhoc_query_details {
		#Required
		adhoc_query_resources {

			#Optional
			region = var.adhoc_query_adhoc_query_details_adhoc_query_resources_region
			resource_ids = var.adhoc_query_adhoc_query_details_adhoc_query_resources_resource_ids
			resource_type = var.adhoc_query_adhoc_query_details_adhoc_query_resources_resource_type
		}
		query = var.adhoc_query_adhoc_query_details_query
	}
	compartment_id = var.compartment_id

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	freeform_tags = {"bar-key"= "value"}
}
```

## Argument Reference

The following arguments are supported:

* `adhoc_query_details` - (Required) Detailed information about the adhoc query.
	* `adhoc_query_resources` - (Required) Target information in which adhoc query will be run
		* `region` - (Optional) Region in which adhoc query needs to be run
		* `resource_ids` - (Optional) List of OCIDs on which query needs to be run
		* `resource_type` - (Optional) Type of resource
	* `query` - (Required) The adhoc query expression that is run
* `compartment_id` - (Required) Compartment OCID of adhoc query
* `defined_tags` - (Optional) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `freeform_tags` - (Optional) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`

	Avoid entering confidential information. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Adhoc Query
	* `update` - (Defaults to 20 minutes), when updating the Adhoc Query
	* `delete` - (Defaults to 20 minutes), when destroying the Adhoc Query


## Import

AdhocQueries can be imported using the `id`, e.g.

```
$ terraform import oci_cloud_guard_adhoc_query.test_adhoc_query "id"
```

