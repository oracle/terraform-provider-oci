---
subcategory: "Cloud Migrations"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_migrations_replication_schedule"
sidebar_current: "docs-oci-datasource-cloud_migrations-replication_schedule"
description: |-
  Provides details about a specific Replication Schedule in Oracle Cloud Infrastructure Cloud Migrations service
---

# Data Source: oci_cloud_migrations_replication_schedule
This data source provides details about a specific Replication Schedule resource in Oracle Cloud Infrastructure Cloud Migrations service.

Gets a replication schedule by identifier.

## Example Usage

```hcl
data "oci_cloud_migrations_replication_schedule" "test_replication_schedule" {
	#Required
	replication_schedule_id = oci_cloud_migrations_replication_schedule.test_replication_schedule.id
}
```

## Argument Reference

The following arguments are supported:

* `replication_schedule_id` - (Required) Unique replication schedule identifier in path


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

