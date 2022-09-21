---
subcategory: "Cloud Migrations"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_migrations_migration"
sidebar_current: "docs-oci-resource-cloud_migrations-migration"
description: |-
  Provides the Migration resource in Oracle Cloud Infrastructure Cloud Migrations service
---

# oci_cloud_migrations_migration
This resource provides the Migration resource in Oracle Cloud Infrastructure Cloud Migrations service.

Creates a migration.


## Example Usage

```hcl
resource "oci_cloud_migrations_migration" "test_migration" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.migration_display_name

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	freeform_tags = {"bar-key"= "value"}
	is_completed = var.migration_is_completed
	replication_schedule_id = oci_cloud_migrations_replication_schedule.test_replication_schedule.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) Compartment identifier
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Required) (Updatable) Migration identifier
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. It exists only for cross-compatibility. Example: `{"bar-key": "value"}` 
* `is_completed` - (Optional) (Updatable) Indicates whether migration is marked as complete.
* `replication_schedule_id` - (Optional) (Updatable) Replication schedule identifier


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - Compartment Identifier
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - Migration Identifier that can be renamed
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. It exists only for cross-compatibility. Example: `{"bar-key": "value"}` 
* `id` - Unique identifier that is immutable on creation
* `is_completed` - Indicates whether migration is marked as completed.
* `lifecycle_details` - A message describing the current state in more detail. For example, it can be used to provide actionable information for a resource in Failed state.
* `replication_schedule_id` - Replication schedule identifier
* `state` - The current state of migration.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time when the migration project was created. An RFC3339 formatted datetime string
* `time_updated` - The time when the migration project was updated. An RFC3339 formatted datetime string

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Migration
	* `update` - (Defaults to 20 minutes), when updating the Migration
	* `delete` - (Defaults to 20 minutes), when destroying the Migration


## Import

Migrations can be imported using the `id`, e.g.

```
$ terraform import oci_cloud_migrations_migration.test_migration "id"
```

