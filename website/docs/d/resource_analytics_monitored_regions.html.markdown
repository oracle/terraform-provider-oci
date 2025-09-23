---
subcategory: "Resource Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_resource_analytics_monitored_regions"
sidebar_current: "docs-oci-datasource-resource_analytics-monitored_regions"
description: |-
  Provides the list of Monitored Regions in Oracle Cloud Infrastructure Resource Analytics service
---

# Data Source: oci_resource_analytics_monitored_regions
This data source provides the list of Monitored Regions in Oracle Cloud Infrastructure Resource Analytics service.

Gets a list of MonitoredRegions.


## Example Usage

```hcl
data "oci_resource_analytics_monitored_regions" "test_monitored_regions" {

	#Optional
	id = var.monitored_region_id
	resource_analytics_instance_id = oci_resource_analytics_resource_analytics_instance.test_resource_analytics_instance.id
	state = var.monitored_region_state
}
```

## Argument Reference

The following arguments are supported:

* `id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the MonitoredRegion.
* `resource_analytics_instance_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a ResourceAnalyticsInstance.
* `state` - (Optional) A filter to return only resources that match the given lifecycle state. The state value is case-insensitive. 


## Attributes Reference

The following attributes are exported:

* `monitored_region_collection` - The list of monitored_region_collection.

### MonitoredRegion Reference

The following attributes are exported:

* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the MonitoredRegion.
* `lifecycle_details` - A message that describes the current state of the MonitoredRegion in more detail. For example, can be used to provide actionable information for a resource in the Failed state. 
* `region_id` - The [Region Identifier](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/regions.htm) of this MonitoredRegion.
* `resource_analytics_instance_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ResourceAnalyticsInstance associated with this MonitoredRegion.
* `state` - The current state of the MonitoredRegion.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the MonitoredRegion was created, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time the MonitoredRegion was updated, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 

