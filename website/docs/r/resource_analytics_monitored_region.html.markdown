---
subcategory: "Resource Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_resource_analytics_monitored_region"
sidebar_current: "docs-oci-resource-resource_analytics-monitored_region"
description: |-
  Provides the Monitored Region resource in Oracle Cloud Infrastructure Resource Analytics service
---

# oci_resource_analytics_monitored_region
This resource provides the Monitored Region resource in Oracle Cloud Infrastructure Resource Analytics service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/resource-analytics/latest/MonitoredRegion

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/

Creates a MonitoredRegion.


## Example Usage

```hcl
resource "oci_resource_analytics_monitored_region" "test_monitored_region" {
	#Required
	region_id = oci_identity_region.test_region.id
	resource_analytics_instance_id = oci_resource_analytics_resource_analytics_instance.test_resource_analytics_instance.id
}
```

## Argument Reference

The following arguments are supported:

* `region_id` - (Required) The [Region Identifier](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/regions.htm) of this MonitoredRegion.
* `resource_analytics_instance_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ResourceAnalyticsInstance associated with this MonitoredRegion.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the MonitoredRegion.
* `lifecycle_details` - A message that describes the current state of the MonitoredRegion in more detail. For example, can be used to provide actionable information for a resource in the Failed state. 
* `region_id` - The [Region Identifier](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/regions.htm) of this MonitoredRegion.
* `resource_analytics_instance_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ResourceAnalyticsInstance associated with this MonitoredRegion.
* `state` - The current state of the MonitoredRegion.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the MonitoredRegion was created, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time the MonitoredRegion was updated, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Monitored Region
	* `update` - (Defaults to 20 minutes), when updating the Monitored Region
	* `delete` - (Defaults to 20 minutes), when destroying the Monitored Region


## Import

MonitoredRegions can be imported using the `id`, e.g.

```
$ terraform import oci_resource_analytics_monitored_region.test_monitored_region "id"
```

