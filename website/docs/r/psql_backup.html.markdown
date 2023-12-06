---
subcategory: "Psql"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_psql_backup"
sidebar_current: "docs-oci-resource-psql-backup"
description: |-
  Provides the Backup resource in Oracle Cloud Infrastructure Psql service
---

# oci_psql_backup
This resource provides the Backup resource in Oracle Cloud Infrastructure Psql service.

Creates a new Backup.


## Example Usage

```hcl
resource "oci_psql_backup" "test_backup" {
	#Required
	compartment_id = var.compartment_id
	db_system_id = oci_psql_db_system.test_db_system.id
	display_name = var.backup_display_name

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.backup_description
	freeform_tags = {"bar-key"= "value"}
	retention_period = var.backup_retention_period
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) Compartment identifier
* `db_system_id` - (Required) Posgresql DbSystem identifier
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) Backup description
* `display_name` - (Required) (Updatable) Backup display name.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `retention_period` - (Optional) (Updatable) Backup retention period in days.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `backup_size` - Backup size in GB.
* `compartment_id` - Backup compartment identifier
* `db_system_details` - Information about the DbSystem associated to a backup.
	* `db_version` - The major and minor versions of the DbSystem software.
	* `system_type` - Type of the DbSystem.
* `db_system_id` - The source DbSystem OCID.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - Backup description
* `display_name` - Backup display name
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - Unique identifier that is immutable on creation
* `last_accepted_request_token` - lastAcceptedRequestToken from MP.
* `last_completed_request_token` - lastCompletedRequestToken from MP.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `retention_period` - Backup retention period in days.
* `source_type` - Specifies whether the backup was created manually, or via scheduled backup policy
* `state` - The current state of the Backup.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the the Backup was created. An RFC3339 formatted datetime string
* `time_updated` - The time the Backup was updated. An RFC3339 formatted datetime string

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Backup
	* `update` - (Defaults to 20 minutes), when updating the Backup
	* `delete` - (Defaults to 20 minutes), when destroying the Backup


## Import

Backups can be imported using the `id`, e.g.

```
$ terraform import oci_psql_backup.test_backup "id"
```

