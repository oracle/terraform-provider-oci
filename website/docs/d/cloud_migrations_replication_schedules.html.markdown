---
subcategory: "Cloud Migrations"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_migrations_replication_schedules"
sidebar_current: "docs-oci-datasource-cloud_migrations-replication_schedules"
description: |-
  Provides the list of Replication Schedules in Oracle Cloud Infrastructure Cloud Migrations service
---

# Data Source: oci_cloud_migrations_replication_schedules
This data source provides the list of Replication Schedules in Oracle Cloud Infrastructure Cloud Migrations service.

Returns a list of replication schedules.


## Example Usage

```hcl
data "oci_cloud_migrations_replication_schedules" "test_replication_schedules" {

	#Optional
	compartment_id = var.compartment_id
	display_name = var.replication_schedule_display_name
	replication_schedule_id = oci_cloud_migrations_replication_schedule.test_replication_schedule.id
	state = var.replication_schedule_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The ID of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the entire given display name.
* `replication_schedule_id` - (Optional) Unique replication schedule identifier in query
* `state` - (Optional) The current state of the replication schedule.


## Attributes Reference

The following attributes are exported:

* `replication_schedule_collection` - The list of replication_schedule_collection.

### ReplicationSchedule Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which the replication schedule exists.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - A name of the replication schedule.
* `execution_recurrences` - Recurrence specification for the replication schedule execution.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. It exists only for cross-compatibility. Example: `{"bar-key": "value"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the replication schedule.
* `lifecycle_details` - The detailed state of the replication schedule.
* `state` - Current state of the replication schedule.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time when the replication schedule was created in RFC3339 format.
* `time_updated` - The time when the replication schedule was last updated in RFC3339 format.

