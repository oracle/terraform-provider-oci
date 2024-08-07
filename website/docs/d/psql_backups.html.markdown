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

Returns a list of backups.


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

* `backup_id` - (Optional) A unique identifier for the backup.
* `compartment_id` - (Optional) The ID of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `id` - (Optional) A unique identifier for the database system.
* `state` - (Optional) A filter to return only resources if their `lifecycleState` matches the given `lifecycleState`.
* `time_ended` - (Optional) The end date for getting backups. An [RFC 3339](https://tools.ietf.org/rfc/rfc3339) formatted datetime string.
* `time_started` - (Optional) The start date for getting backups. An [RFC 3339](https://tools.ietf.org/rfc/rfc3339) formatted datetime string.


## Attributes Reference

The following attributes are exported:

* `backup_collection` - The list of backup_collection.

### Backup Reference

The following attributes are exported:

* `backup_size` - The size of the backup, in gigabytes.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the backup.
* `db_system_details` - Information about the database system associated with a backup.
	* `db_version` - The major and minor versions of the database system software.
	* `system_type` - Type of the database system.
* `db_system_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup's source database system.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - A description for the backup.
* `display_name` - A user-friendly display name for the backup. Avoid entering confidential information.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup.
* `last_accepted_request_token` - lastAcceptedRequestToken from MP.
* `last_completed_request_token` - lastCompletedRequestToken from MP.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `retention_period` - Backup retention period in days.
* `source_type` - Specifies whether the backup was created manually, or by a management policy.
* `state` - The current state of the backup.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the backup request was received, expressed in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format.  Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time the backup was updated, expressed in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format.  Example: `2016-08-25T21:10:29.600Z` 

