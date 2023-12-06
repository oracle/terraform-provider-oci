---
subcategory: "Cloud Migrations"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_migrations_replication_schedule"
sidebar_current: "docs-oci-resource-cloud_migrations-replication_schedule"
description: |-
  Provides the Replication Schedule resource in Oracle Cloud Infrastructure Cloud Migrations service
---

# oci_cloud_migrations_replication_schedule
This resource provides the Replication Schedule resource in Oracle Cloud Infrastructure Cloud Migrations service.

Creates a replication schedule.


## Example Usage

```hcl
resource "oci_cloud_migrations_replication_schedule" "test_replication_schedule" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.replication_schedule_display_name
	execution_recurrences = var.replication_schedule_execution_recurrences

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	freeform_tags = {"bar-key"= "value"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which the replication schedule should be created.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Required) (Updatable) A user-friendly name for a replication schedule. Does not have to be unique, and is mutable. Avoid entering confidential information.
* `execution_recurrences` - (Required) (Updatable) Recurrence specification for replication schedule execution.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. It exists only for cross-compatibility. Example: `{"bar-key": "value"}` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Replication Schedule
	* `update` - (Defaults to 20 minutes), when updating the Replication Schedule
	* `delete` - (Defaults to 20 minutes), when destroying the Replication Schedule


## Import

ReplicationSchedules can be imported using the `id`, e.g.

```
$ terraform import oci_cloud_migrations_replication_schedule.test_replication_schedule "id"
```

