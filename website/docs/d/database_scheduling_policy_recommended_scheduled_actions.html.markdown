---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_scheduling_policy_recommended_scheduled_actions"
sidebar_current: "docs-oci-datasource-database-scheduling_policy_recommended_scheduled_actions"
description: |-
  Provides the list of Scheduling Policy Recommended Scheduled Actions in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_scheduling_policy_recommended_scheduled_actions
This data source provides the list of Scheduling Policy Recommended Scheduled Actions in Oracle Cloud Infrastructure Database service.

Returns a recommended Scheduled Actions configuration for a given resource, plan intent and scheduling policy.


## Example Usage

```hcl
data "oci_database_scheduling_policy_recommended_scheduled_actions" "test_scheduling_policy_recommended_scheduled_actions" {
	#Required
	plan_intent = var.scheduling_policy_recommended_scheduled_action_plan_intent
	scheduling_policy_id = oci_database_scheduling_policy.test_scheduling_policy.id
	scheduling_policy_target_resource_id = oci_cloud_guard_resource.test_resource.id
}
```

## Argument Reference

The following arguments are supported:

* `plan_intent` - (Required) The scheduling plan intent the scheduled actions will be for.
* `scheduling_policy_id` - (Required) The Scheduling Policy [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `scheduling_policy_target_resource_id` - (Required) The target resource [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) the scheduled actions will be for.


## Attributes Reference

The following attributes are exported:

* `recommended_scheduled_actions_collection` - The list of recommended_scheduled_actions_collection.

### SchedulingPolicyRecommendedScheduledAction Reference

The following attributes are exported:

* `items` - List of scheduled actions.
	* `action_members` - The list of action members in a scheduled action.
		* `estimated_time_in_mins` - The estimated time for the intended action member.
		* `member_id` - The ocid of the action member.
		* `member_order` - The order of the action member in a scheduled action.
	* `action_order` - The order of the scheduled action.
	* `action_params` - Map<ParamName, ParamValue> where a key value pair describes the specific action parameter. Example: `{"count": "3"}` 
	* `action_type` - The type of the scheduled action being performed
	* `display_name` - Description of the scheduled action being performed, i.e. apply full update to DB Servers 1,2,3,4.
	* `estimated_time_in_mins` - The estimated patching time in minutes for the entire scheduled action.
	* `scheduling_window_id` - The id of the scheduling window this scheduled action belongs to.

