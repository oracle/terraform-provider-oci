---
subcategory: "Cloud Migrations"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_migrations_migrations"
sidebar_current: "docs-oci-datasource-cloud_migrations-migrations"
description: |-
  Provides the list of Migrations in Oracle Cloud Infrastructure Cloud Migrations service
---

# Data Source: oci_cloud_migrations_migrations
This data source provides the list of Migrations in Oracle Cloud Infrastructure Cloud Migrations service.

Returns a list of migrations.


## Example Usage

```hcl
data "oci_cloud_migrations_migrations" "test_migrations" {

	#Optional
	compartment_id = var.compartment_id
	display_name = var.migration_display_name
	migration_id = oci_cloud_migrations_migration.test_migration.id
	state = var.migration_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The ID of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the entire given display name.
* `migration_id` - (Optional) Unique migration identifier
* `state` - (Optional) A filter to return only resources where the resource's lifecycle state matches the given lifecycle state.


## Attributes Reference

The following attributes are exported:

* `migration_collection` - The list of migration_collection.

### Migration Reference

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

