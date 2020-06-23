---
subcategory: "Mysql"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_mysql_mysql_backup"
sidebar_current: "docs-oci-datasource-mysql-mysql_backup"
description: |-
  Provides details about a specific Mysql Backup in Oracle Cloud Infrastructure Mysql service
---

# Data Source: oci_mysql_mysql_backup
This data source provides details about a specific Mysql Backup resource in Oracle Cloud Infrastructure Mysql service.

Get information about the specified Backup

## Example Usage

```hcl
data "oci_mysql_mysql_backup" "test_mysql_backup" {
	#Required
	backup_id = "${oci_mysql_mysql_backup.test_backup.id}"
}
```

## Argument Reference

The following arguments are supported:

* `backup_id` - (Required) The OCID of the Backup


## Attributes Reference

The following attributes are exported:

* `backup_size_in_gbs` - The size of the backup in base-2 (IEC) gibibytes. (GiB).
* `backup_type` - The type of backup.
* `compartment_id` - The OCID of the compartment.
* `creation_type` - If the backup was created automatically, or by a manual request.
* `data_storage_size_in_gb` - Initial size of the data volume in GiBs. 
* `db_system_id` - The OCID of the DB System the backup is associated with.
* `defined_tags` - Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - A user-supplied description for the backup.
* `display_name` - A user-supplied display name for the backup.
* `freeform_tags` - Simple key-value pair applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - OCID of the backup itself
* `lifecycle_details` - Additional information about the current lifecycleState.
* `mysql_version` - The MySQL server version of the DB System used for backup.
* `retention_in_days` - Number of days to retain this backup.
* `shape_name` - The shape of the DB System instance used for backup.
* `state` - The state of the backup.
* `time_created` - The time the backup record was created.
* `time_updated` - The time at which the backup was updated.

