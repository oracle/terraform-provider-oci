---
subcategory: "Psql"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_psql_backup"
sidebar_current: "docs-oci-datasource-psql-backup"
description: |-
  Provides details about a specific Backup in Oracle Cloud Infrastructure Psql service
---

# Data Source: oci_psql_backup
This data source provides details about a specific Backup resource in Oracle Cloud Infrastructure Psql service.

Gets a Backup by identifier

## Example Usage

```hcl
data "oci_psql_backup" "test_backup" {
	#Required
	backup_id = oci_psql_backup.test_backup.id
}
```

## Argument Reference

The following arguments are supported:

* `backup_id` - (Required) unique Backup identifier


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

