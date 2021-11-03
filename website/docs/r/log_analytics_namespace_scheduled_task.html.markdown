---
subcategory: "Log Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_log_analytics_namespace_scheduled_task"
sidebar_current: "docs-oci-resource-log_analytics-namespace_scheduled_task"
description: |-
  Provides the Namespace Scheduled Task resource in Oracle Cloud Infrastructure Log Analytics service
---

# oci_log_analytics_namespace_scheduled_task
This resource provides the Namespace Scheduled Task resource in Oracle Cloud Infrastructure Log Analytics service.

Schedule a task as specified and return task info.

## Example Usage

```hcl
resource "oci_log_analytics_namespace_scheduled_task" "test_namespace_scheduled_task" {
	#Required
	compartment_id = var.compartment_id
	kind = var.namespace_scheduled_task_kind
	namespace = var.namespace_scheduled_task_namespace

	#Optional
	action {
		#Required
		type = var.namespace_scheduled_task_action_type

		#Optional
		compartment_id_in_subtree = var.namespace_scheduled_task_action_compartment_id_in_subtree
		data_type = var.namespace_scheduled_task_action_data_type
		purge_compartment_id = oci_identity_compartment.test_compartment.id
		purge_duration = var.namespace_scheduled_task_action_purge_duration
		query_string = var.namespace_scheduled_task_action_query_string
		saved_search_id = oci_log_analytics_saved_search.test_saved_search.id
	}
	defined_tags = {"foo-namespace.bar-key"= "value"}
	display_name = var.namespace_scheduled_task_display_name
	freeform_tags = {"bar-key"= "value"}
	saved_search_id = oci_log_analytics_saved_search.test_saved_search.id
	schedules {
		#Required
		type = var.namespace_scheduled_task_schedules_type

		#Optional
		expression = var.namespace_scheduled_task_schedules_expression
		misfire_policy = var.namespace_scheduled_task_schedules_misfire_policy
		recurring_interval = var.namespace_scheduled_task_schedules_recurring_interval
		repeat_count = var.namespace_scheduled_task_schedules_repeat_count
		time_zone = var.namespace_scheduled_task_schedules_time_zone
	}
	task_type = var.namespace_scheduled_task_task_type
}
```

## Argument Reference

The following arguments are supported:

* `action` - (Required when kind=STANDARD) Action for scheduled task.
	* `compartment_id_in_subtree` - (Applicable when type=PURGE) if true, purge child compartments data
	* `data_type` - (Required when type=PURGE) the type of the log data to be purged
	* `purge_compartment_id` - (Required when type=PURGE) the compartment OCID under which the data will be purged
	* `purge_duration` - (Required when type=PURGE) The duration of data to be retained, which is used to calculate the timeDataEnded when the task fires. The value should be negative. Purge duration in ISO 8601 extended format as described in https://en.wikipedia.org/wiki/ISO_8601#Durations. The largest supported unit is D, e.g. -P365D (not -P1Y) or -P14D (not -P2W). 
	* `query_string` - (Required when type=PURGE) Purge query string.
	* `saved_search_id` - (Applicable when type=STREAM) The ManagementSavedSearch id [OCID] utilized in the action.
	* `type` - (Required) Action type discriminator.
* `compartment_id` - (Required) (Updatable) Compartment Identifier [OCID] (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name that is changeable and that does not have to be unique. Format: a leading alphanumeric, followed by zero or more alphanumerics, underscores, spaces, backslashes, or hyphens in any order). No trailing spaces allowed. 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `kind` - (Required) Discriminator.
* `namespace` - (Required) The Logging Analytics namespace used for the request. 
* `saved_search_id` - (Required when kind=ACCELERATION) The ManagementSavedSearch id [OCID] to be accelerated.
* `schedules` - (Required when kind=STANDARD) (Updatable) Schedules, typically a single schedule. Note there may only be a single schedule for SAVED_SEARCH and PURGE scheduled tasks. 
	* `expression` - (Required when type=CRON) (Updatable) Value in cron format.
	* `misfire_policy` - (Applicable when kind=STANDARD) (Updatable) Schedule misfire retry policy.
	* `recurring_interval` - (Required when type=FIXED_FREQUENCY) (Updatable) Recurring interval in ISO 8601 extended format as described in https://en.wikipedia.org/wiki/ISO_8601#Durations. The largest supported unit is D, e.g. P14D (not P2W). The value must be at least 5 minutes (PT5M) and at most 3 weeks (P21D or PT30240M). 
	* `repeat_count` - (Applicable when type=FIXED_FREQUENCY) (Updatable) Number of times (0-based) to execute until auto-stop. Default value -1 will execute indefinitely. Value 0 will execute once. 
	* `time_zone` - (Required when type=CRON) (Updatable) Time zone, by default UTC.
	* `type` - (Required) (Updatable) Schedule type discriminator.
* `task_type` - (Required when kind=STANDARD) Task type.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `action` - Action for scheduled task.
	* `compartment_id_in_subtree` - if true, purge child compartments data
	* `data_type` - the type of the log data to be purged
	* `purge_compartment_id` - the compartment OCID under which the data will be purged
	* `purge_duration` - The duration of data to be retained, which is used to calculate the timeDataEnded when the task fires. The value should be negative. Purge duration in ISO 8601 extended format as described in https://en.wikipedia.org/wiki/ISO_8601#Durations. The largest supported unit is D, e.g. -P365D (not -P1Y) or -P14D (not -P2W). 
	* `query_string` - Purge query string.
	* `saved_search_id` - The ManagementSavedSearch id [OCID] utilized in the action.
	* `type` - Action type discriminator.
* `compartment_id` - Compartment Identifier [OCID] (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - A user-friendly name that is changeable and that does not have to be unique. Format: a leading alphanumeric, followed by zero or more alphanumerics, underscores, spaces, backslashes, or hyphens in any order). No trailing spaces allowed. 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the data plane resource. 
* `num_occurrences` - Number of execution occurrences.
* `schedules` - Schedules.
	* `expression` - Value in cron format.
	* `misfire_policy` - Schedule misfire retry policy.
	* `recurring_interval` - Recurring interval in ISO 8601 extended format as described in https://en.wikipedia.org/wiki/ISO_8601#Durations. The largest supported unit is D, e.g. P14D (not P2W). The value must be at least 5 minutes (PT5M) and at most 3 weeks (P21D or PT30240M). 
	* `repeat_count` - Number of times (0-based) to execute until auto-stop. Default value -1 will execute indefinitely. Value 0 will execute once. 
	* `time_zone` - Time zone, by default UTC.
	* `type` - Schedule type discriminator.
* `state` - The current state of the scheduled task.
* `task_status` - Status of the scheduled task. - PURGE_RESOURCE_NOT_FOUND
* `task_type` - Task type.
* `time_created` - The date and time the scheduled task was created, in the format defined by RFC3339. 
* `time_updated` - The date and time the scheduled task was last updated, in the format defined by RFC3339. 
* `work_request_id` - most recent Work Request Identifier [OCID] (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the asynchronous request.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Namespace Scheduled Task
	* `update` - (Defaults to 20 minutes), when updating the Namespace Scheduled Task
	* `delete` - (Defaults to 20 minutes), when destroying the Namespace Scheduled Task


## Import

NamespaceScheduledTasks can be imported using the `id`, e.g.

```
$ terraform import oci_log_analytics_namespace_scheduled_task.test_namespace_scheduled_task "namespaces/{namespaceName}/scheduledTasks/{scheduledTaskId}" 
```

