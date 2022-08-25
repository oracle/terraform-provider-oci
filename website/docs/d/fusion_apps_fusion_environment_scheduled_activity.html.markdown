---
subcategory: "Fusion Apps"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fusion_apps_fusion_environment_scheduled_activity"
sidebar_current: "docs-oci-datasource-fusion_apps-fusion_environment_scheduled_activity"
description: |-
  Provides details about a specific Fusion Environment Scheduled Activity in Oracle Cloud Infrastructure Fusion Apps service
---

# Data Source: oci_fusion_apps_fusion_environment_scheduled_activity
This data source provides details about a specific Fusion Environment Scheduled Activity resource in Oracle Cloud Infrastructure Fusion Apps service.

Gets a ScheduledActivity by identifier

## Example Usage

```hcl
data "oci_fusion_apps_fusion_environment_scheduled_activity" "test_fusion_environment_scheduled_activity" {
	#Required
	fusion_environment_id = oci_fusion_apps_fusion_environment.test_fusion_environment.id
	scheduled_activity_id = oci_fusion_apps_scheduled_activity.test_scheduled_activity.id
}
```

## Argument Reference

The following arguments are supported:

* `fusion_environment_id` - (Required) unique FusionEnvironment identifier
* `scheduled_activity_id` - (Required) Unique ScheduledActivity identifier.


## Attributes Reference

The following attributes are exported:

* `actions` - List of actions
	* `action_type` - Type of action
	* `artifact` - patch that delivered the vertex update prerequisite
	* `category` - patch artifact category
	* `description` - A string that describes the details of the action. It does not have to be unique, and you can change it. Avoid entering confidential information.
	* `mode` - A string that describeds whether the change is applied hot or cold
	* `qualifier` - month qualifier
	* `reference_key` - Unique identifier of the object that represents the action
	* `state` - A string that describes whether the change is applied hot or cold
	* `version` - name of the repo
* `delay_in_hours` - Cumulative delay hours
* `display_name` - scheduled activity display name, can be renamed.
* `fusion_environment_id` - FAaaS Environment Identifier.
* `id` - Unique identifier that is immutable on creation.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `run_cycle` - run cadence.
* `service_availability` - Service availability / impact during scheduled activity execution up down
* `state` - The current state of the scheduledActivity.
* `time_created` - The time the scheduled activity record was created. An RFC3339 formatted datetime string.
* `time_expected_finish` - Current time the scheduled activity is scheduled to end. An RFC3339 formatted datetime string.
* `time_finished` - The time the scheduled activity actually completed / cancelled / failed. An RFC3339 formatted datetime string.
* `time_scheduled_start` - Current time the scheduled activity is scheduled to start. An RFC3339 formatted datetime string.
* `time_updated` - The time the scheduled activity record was updated. An RFC3339 formatted datetime string.

