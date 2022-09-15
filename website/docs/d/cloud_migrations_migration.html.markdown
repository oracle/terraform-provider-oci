---
subcategory: "Cloud Migrations"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_migrations_migration"
sidebar_current: "docs-oci-datasource-cloud_migrations-migration"
description: |-
  Provides details about a specific Migration in Oracle Cloud Infrastructure Cloud Migrations service
---

# Data Source: oci_cloud_migrations_migration
This data source provides details about a specific Migration resource in Oracle Cloud Infrastructure Cloud Migrations service.

Gets a migration by identifier.

## Example Usage

```hcl
data "oci_cloud_migrations_migration" "test_migration" {
	#Required
	migration_id = oci_cloud_migrations_migration.test_migration.id
}
```

## Argument Reference

The following arguments are supported:

* `migration_id` - (Required) Unique migration identifier


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

