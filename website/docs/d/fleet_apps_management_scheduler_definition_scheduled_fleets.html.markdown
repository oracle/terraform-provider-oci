---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_scheduler_definition_scheduled_fleets"
sidebar_current: "docs-oci-datasource-fleet_apps_management-scheduler_definition_scheduled_fleets"
description: |-
  Provides the list of Scheduler Definition Scheduled Fleets in Oracle Cloud Infrastructure Fleet Apps Management service
---

# Data Source: oci_fleet_apps_management_scheduler_definition_scheduled_fleets
This data source provides the list of Scheduler Definition Scheduled Fleets in Oracle Cloud Infrastructure Fleet Apps Management service.

Returns a list of ScheduledFleets


## Example Usage

```hcl
data "oci_fleet_apps_management_scheduler_definition_scheduled_fleets" "test_scheduler_definition_scheduled_fleets" {
	#Required
	scheduler_definition_id = oci_fleet_apps_management_scheduler_definition.test_scheduler_definition.id

	#Optional
	compartment_id = var.compartment_id
	display_name = var.scheduler_definition_scheduled_fleet_display_name
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The ID of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `scheduler_definition_id` - (Required) unique SchedulerDefinition identifier


## Attributes Reference

The following attributes are exported:

* `scheduled_fleet_collection` - The list of scheduled_fleet_collection.

### SchedulerDefinitionScheduledFleet Reference

The following attributes are exported:

* `items` - List of ScheduledFleets.
	* `action_group_types` - All ActionGroup Types part of the schedule.
	* `application_types` - All application types part of the schedule.
	* `count_of_affected_resources` - Count of Resources affected by the Schedule
	* `count_of_affected_targets` - Count of Targets affected by the Schedule
	* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
	* `id` - The OCID of the resource.
	* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
	* `tenancy_id` - OCID of the tenancy to which the resource belongs to.

