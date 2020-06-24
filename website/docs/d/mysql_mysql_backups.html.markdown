---
subcategory: "Mysql"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_mysql_mysql_backups"
sidebar_current: "docs-oci-datasource-mysql-mysql_backups"
description: |-
  Provides the list of Mysql Backups in Oracle Cloud Infrastructure Mysql service
---

# Data Source: oci_mysql_mysql_backups
This data source provides the list of Mysql Backups in Oracle Cloud Infrastructure Mysql service.

Get a list of DB System backups.


## Example Usage

```hcl
data "oci_mysql_mysql_backups" "test_mysql_backups" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	backup_id = "${oci_mysql_mysql_backup.test_backup.id}"
	db_system_id = "${oci_mysql_mysql_db_system.test_db_system.id}"
	display_name = "${var.mysql_backup_display_name}"
	state = "${var.mysql_backup_state}"
}
```

## Argument Reference

The following arguments are supported:

* `backup_id` - (Optional) Backup OCID
* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `db_system_id` - (Optional) The DB System [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `display_name` - (Optional) A filter to return only the resource matching the given display name exactly.
* `state` - (Optional) Backup Lifecycle State


## Attributes Reference

The following attributes are exported:

* `backups` - The list of backups.

### MysqlBackup Reference

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

