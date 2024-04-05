---
subcategory: "Stack Monitoring"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_stack_monitoring_monitored_resources_search"
sidebar_current: "docs-oci-resource-stack_monitoring-monitored_resources_search"
description: |-
  Provides the Monitored Resources Search resource in Oracle Cloud Infrastructure Stack Monitoring service
---

# oci_stack_monitoring_monitored_resources_search
This resource provides the Monitored Resources Search resource in Oracle Cloud Infrastructure Stack Monitoring service.

Gets a list of all monitored resources in a compartment for the given search criteria.


## Example Usage

```hcl
resource "oci_stack_monitoring_monitored_resources_search" "test_monitored_resources_search" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	compartment_ids = var.monitored_resources_search_compartment_ids
	exclude_fields = var.monitored_resources_search_exclude_fields
	external_id = oci_stack_monitoring_external.test_external.id
	fields = var.monitored_resources_search_fields
	host_name = var.monitored_resources_search_host_name
	host_name_contains = var.monitored_resources_search_host_name_contains
	license = var.monitored_resources_search_license
	lifecycle_states = var.monitored_resources_search_lifecycle_states
	management_agent_id = oci_management_agent_management_agent.test_management_agent.id
	name = var.monitored_resources_search_name
	name_contains = var.monitored_resources_search_name_contains
	property_equals = var.monitored_resources_search_property_equals
	resource_category = var.monitored_resources_search_resource_category
	resource_time_zone = var.monitored_resources_search_resource_time_zone
	source_type = var.monitored_resources_search_source_type
	state = var.monitored_resources_search_state
	time_created_greater_than_or_equal_to = var.monitored_resources_search_time_created_greater_than_or_equal_to
	time_created_less_than = var.monitored_resources_search_time_created_less_than
	time_updated_greater_than_or_equal_to = var.monitored_resources_search_time_updated_greater_than_or_equal_to
	time_updated_less_than = var.monitored_resources_search_time_updated_less_than
	type = var.monitored_resources_search_type
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) Compartment Identifier [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). 
* `compartment_ids` - (Optional) Multiple compartment identifiers [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). 
* `exclude_fields` - (Optional) Partial response refers to an optimization technique offered by the RESTful web APIs, to return all the information except the fields requested to be excluded (excludeFields) by the client. In this mechanism, the client sends the exclude field names as the query parameters for an API to the server, and the server trims down the default response content by removing the fields that are not required by the client. The parameter controls which fields to exlude and to return and should be a query string parameter called "excludeFields" of an array type, provide the values as enums, and use collectionFormat. 
* `external_id` - (Optional) External resource is any Oracle Cloud Infrastructure resource identifier [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) which is not a Stack Monitoring service resource. Currently supports only following resource types - Container database, non-container database,  pluggable database and Oracle Cloud Infrastructure compute instance. 
* `fields` - (Optional) Partial response refers to an optimization technique offered by the RESTful web APIs, to return only the information (fields) required by the client. In this mechanism, the client sends the required field names as the query parameters for an API to the server, and the server trims down the default response content by removing the fields that are not required by the client. The parameter controls which fields to return and should be a query string parameter called "fields" of an array type, provide the values as enums, and use collectionFormat. 
* `host_name` - (Optional) A filter to return resources with host name match. 
* `host_name_contains` - (Optional) A filter to return resources with host name pattern. 
* `license` - (Optional) License edition of the monitored resource.
* `lifecycle_states` - (Optional) Multiple lifecycle states filter. 
* `management_agent_id` - (Optional) A filter to return resources with matching management agent id.
* `name` - (Optional) A filter to return resources that match exact resource name. 
* `name_contains` - (Optional) A filter to return resources that match resource name pattern given. The match is not case sensitive.
* `property_equals` - (Optional) Criteria based on resource property.
* `resource_category` - (Optional) Resource category filter.
* `resource_time_zone` - (Optional) Time zone in the form of tz database canonical zone ID. Specifies the preference with a value that uses the IANA Time Zone Database format (x-obmcs-time-zone). For example - America/Los_Angeles 
* `source_type` - (Optional) Source type filter.
* `state` - (Optional) A filter to return resources with matching lifecycle state.
* `time_created_greater_than_or_equal_to` - (Optional) Search for resources that were created within a specific date range, using this parameter to specify the earliest creation date for the returned list (inclusive). Specifying this parameter without the corresponding `timeCreatedLessThan` parameter will retrieve resources created from the given `timeCreatedGreaterThanOrEqualTo` to the current time, in "YYYY-MM-ddThh:mmZ" format with a Z offset, as defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).

	**Example:** 2016-12-19T16:39:57.600Z 
* `time_created_less_than` - (Optional) Search for resources that were created within a specific date range, using this parameter to specify the latest creation date for the returned list (exclusive). Specifying this parameter without the corresponding `timeCreatedGreaterThanOrEqualTo` parameter will retrieve all resources created before the specified end date, in "YYYY-MM-ddThh:mmZ" format with a Z offset, as defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).

	**Example:** 2016-12-19T16:39:57.600Z 
* `time_updated_greater_than_or_equal_to` - (Optional) Search for resources that were updated within a specific date range, using this parameter to specify the earliest update date for the returned list (inclusive). Specifying this parameter without the corresponding `timeUpdatedLessThan` parameter will retrieve resources updated from the given `timeUpdatedGreaterThanOrEqualTo` to the current time, in "YYYY-MM-ddThh:mmZ" format with a Z offset, as defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).

	**Example:** 2016-12-19T16:39:57.600Z 
* `time_updated_less_than` - (Optional) Search for resources that were updated within a specific date range, using this parameter to specify the latest creation date for the returned list (exclusive). Specifying this parameter without the corresponding `timeUpdatedGreaterThanOrEqualTo` parameter will retrieve all resources updated before the specified end date, in "YYYY-MM-ddThh:mmZ" format with a Z offset, as defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).

	**Example:** 2016-12-19T16:39:57.600Z 
* `type` - (Optional) A filter to return resources that match resource type. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `items` - List of monitored resources.
	* `compartment_id` - Compartment Identifier [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
	* `display_name` - Monitored resource display name.
	* `external_id` - External resource is any Oracle Cloud Infrastructure resource identifier [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) which is not a Stack Monitoring service resource. 
	* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
	* `host_name` - Monitored Resource Host Name. 
	* `id` - Monitored resource identifier [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). 
	* `license` - License edition of the monitored resource.
	* `management_agent_id` - Management Agent Identifier [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	* `name` - Monitored Resource Name. 
	* `properties` - List of monitored resource properties. 
		* `name` - Property Name. 
		* `value` - Property Value. 
	* `resource_category` - Resource Category to indicate the kind of resource type. 
	* `source_type` - Source type to indicate if the resource is stack monitoring discovered, Oracle Cloud Infrastructure native resource, etc. 
	* `state` - The current state of the monitored resource.
	* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
	* `time_created` - Monitored resource creation time. An RFC3339 formatted datetime string. 
	* `time_updated` - Monitored resource update time. An RFC3339 formatted datetime string. 
	* `type` - Monitored Resource Type. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Monitored Resources Search
	* `update` - (Defaults to 20 minutes), when updating the Monitored Resources Search
	* `delete` - (Defaults to 20 minutes), when destroying the Monitored Resources Search


## Import

MonitoredResourcesSearch can be imported using the `id`, e.g.

```
$ terraform import oci_stack_monitoring_monitored_resources_search.test_monitored_resources_search "id"
```

