---
subcategory: "Psql"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_psql_backups"
sidebar_current: "docs-oci-datasource-psql-backups"
description: |-
  Provides the list of Backups in Oracle Cloud Infrastructure Psql service
---

# Data Source: oci_psql_backups
This data source provides the list of Backups in Oracle Cloud Infrastructure Psql service.

Returns a list of Backup.


## Example Usage

```hcl
data "oci_psql_backups" "test_backups" {

	#Optional
	backup_id = oci_psql_backup.test_backup.id
	compartment_id = var.compartment_id
	display_name = var.backup_display_name
	id = var.backup_id
	state = var.backup_state
	time_ended = var.backup_time_ended
	time_started = var.backup_time_started
}
```

## Argument Reference

The following arguments are supported:

* `backup_id` - (Optional) unique Backup identifier
* `compartment_id` - (Optional) The ID of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `id` - (Optional) unique DbSystem identifier
* `state` - (Optional) A filter to return only resources their lifecycleState matches the given lifecycleState.
* `time_ended` - (Optional) The End date for getting  backups. An RFC3339 formatted datetime string.
* `time_started` - (Optional) The start date for getting  backups. An RFC3339 formatted datetime string


## Attributes Reference

The following attributes are exported:

* `backup_collection` - The list of backup_collection.

### Backup Reference

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

