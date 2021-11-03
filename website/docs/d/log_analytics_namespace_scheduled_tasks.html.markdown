---
subcategory: "Log Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_log_analytics_namespace_scheduled_tasks"
sidebar_current: "docs-oci-datasource-log_analytics-namespace_scheduled_tasks"
description: |-
  Provides the list of Namespace Scheduled Tasks in Oracle Cloud Infrastructure Log Analytics service
---

# Data Source: oci_log_analytics_namespace_scheduled_tasks
This data source provides the list of Namespace Scheduled Tasks in Oracle Cloud Infrastructure Log Analytics service.

Lists scheduled tasks.

## Example Usage

```hcl
data "oci_log_analytics_namespace_scheduled_tasks" "test_namespace_scheduled_tasks" {
	#Required
	compartment_id = var.compartment_id
	namespace = var.namespace_scheduled_task_namespace

	#Optional
	display_name = var.namespace_scheduled_task_display_name
	task_type = var.namespace_scheduled_task_task_type
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. 
* `namespace` - (Required) The Logging Analytics namespace used for the request. 
* `task_type` - (Required when kind=STANDARD) Required parameter to specify schedule task type.


## Attributes Reference

The following attributes are exported:

* `scheduled_task_collection` - The list of scheduled_task_collection.

### NamespaceScheduledTask Reference

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

