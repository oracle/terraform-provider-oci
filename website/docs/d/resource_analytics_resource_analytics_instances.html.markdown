---
subcategory: "Resource Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_resource_analytics_resource_analytics_instances"
sidebar_current: "docs-oci-datasource-resource_analytics-resource_analytics_instances"
description: |-
  Provides the list of Resource Analytics Instances in Oracle Cloud Infrastructure Resource Analytics service
---

# Data Source: oci_resource_analytics_resource_analytics_instances
This data source provides the list of Resource Analytics Instances in Oracle Cloud Infrastructure Resource Analytics service.

Gets a list of ResourceAnalyticsInstances.


## Example Usage

```hcl
data "oci_resource_analytics_resource_analytics_instances" "test_resource_analytics_instances" {

	#Optional
	compartment_id = var.compartment_id
	display_name = var.resource_analytics_instance_display_name
	id = var.resource_analytics_instance_id
	state = var.resource_analytics_instance_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly.
* `id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ResourceAnalyticsInstance.
* `state` - (Optional) A filter to return only resources that match the given lifecycle state. The state value is case-insensitive. 


## Attributes Reference

The following attributes are exported:

* `resource_analytics_instance_collection` - The list of resource_analytics_instance_collection.

### ResourceAnalyticsInstance Reference

The following attributes are exported:

* `adw_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the created ADW instance.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - A description of the ResourceAnalyticsInstance instance.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ResourceAnalyticsInstance.
* `lifecycle_details` - A message that describes the current state of the ResourceAnalyticsInstance in more detail. For example, can be used to provide actionable information for a resource in the Failed state. 
* `oac_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the OAC enabled for the ResourceAnalyticsInstance.
* `state` - The current state of the ResourceAnalyticsInstance.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the ResourceAnalyticsInstance was created, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339). Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time the ResourceAnalyticsInstance was updated, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 

