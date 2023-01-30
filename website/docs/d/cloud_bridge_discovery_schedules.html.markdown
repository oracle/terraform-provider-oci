---
subcategory: "Cloud Bridge"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_bridge_discovery_schedules"
sidebar_current: "docs-oci-datasource-cloud_bridge-discovery_schedules"
description: |-
  Provides the list of Discovery Schedules in Oracle Cloud Infrastructure Cloud Bridge service
---

# Data Source: oci_cloud_bridge_discovery_schedules
This data source provides the list of Discovery Schedules in Oracle Cloud Infrastructure Cloud Bridge service.

Lists discovery schedules.

## Example Usage

```hcl
data "oci_cloud_bridge_discovery_schedules" "test_discovery_schedules" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	discovery_schedule_id = oci_cloud_bridge_discovery_schedule.test_discovery_schedule.id
	display_name = var.discovery_schedule_display_name
	state = var.discovery_schedule_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `discovery_schedule_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the discovery schedule.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `state` - (Optional) The current state of the discovery schedule.


## Attributes Reference

The following attributes are exported:

* `discovery_schedule_collection` - The list of discovery_schedule_collection.

### DiscoverySchedule Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which the discovery schedule exists.
* `defined_tags` - The defined tags associated with this resource, if any. Each key is predefined and scoped to namespaces. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name for the discovery schedule. Does not have to be unique, and it's mutable. Avoid entering confidential information. 
* `execution_recurrences` - Recurrence specification for the discovery schedule execution.
* `freeform_tags` - The freeform tags associated with this resource, if any. Each tag is a simple key-value pair with no predefined name, type, or namespace/scope. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the discovery schedule.
* `lifecycle_details` - The detailed state of the discovery schedule.
* `state` - Current state of the discovery schedule.
* `system_tags` - The system tags associated with this resource, if any. The system tags are set by Oracle cloud infrastructure services. Each key is predefined and scoped to namespaces. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{orcl-cloud: {free-tier-retain: true}}` 
* `time_created` - The time when the discovery schedule was created in RFC3339 format.
* `time_updated` - The time when the discovery schedule was last updated in RFC3339 format.

