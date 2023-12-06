---
subcategory: "MySQL Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_mysql_mysql_backup"
sidebar_current: "docs-oci-resource-mysql-mysql_backup"
description: |-
  Provides the Mysql Backup resource in Oracle Cloud Infrastructure MySQL Database service
---

# oci_mysql_mysql_backup
This resource provides the Mysql Backup resource in Oracle Cloud Infrastructure MySQL Database service.

Create a backup of a DB System.


## Example Usage

```hcl
resource "oci_mysql_mysql_backup" "test_mysql_backup" {
	#Required
	db_system_id = oci_mysql_mysql_db_system.test_db_system.id

	#Optional
	backup_type = var.mysql_backup_backup_type
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.mysql_backup_description
	display_name = var.mysql_backup_display_name
	freeform_tags = {"bar-key"= "value"}
	retention_in_days = var.mysql_backup_retention_in_days
}
```

## Argument Reference

The following arguments are supported:

* `backup_type` - (Optional) The type of backup.
* `compartment_id` - (Optional) (Updatable) The OCID of the compartment the backup exists in.
* `db_system_id` - (Required) The OCID of the DB System the Backup is associated with.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) A user-supplied description for the backup.
* `display_name` - (Optional) (Updatable) A user-supplied display name for the backup.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `retention_in_days` - (Optional) (Updatable) Number of days to retain this backup.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `backup_size_in_gbs` - The size of the backup in base-2 (IEC) gibibytes. (GiB).
* `backup_type` - The type of backup.
* `compartment_id` - The OCID of the compartment the backup exists in.
* `creation_type` - Indicates how the backup was created: manually, automatic, or by an Operator. 
* `data_storage_size_in_gb` - Initial size of the data volume in GiBs. 
* `db_system_id` - The OCID of the DB System the backup is associated with.
* `db_system_snapshot` - Snapshot of the DbSystem details at the time of the backup 
	* `admin_username` - The username for the administrative user.
	* `availability_domain` - The Availability Domain where the primary DB System should be located. 
	* `backup_policy` - The Backup policy for the DB System.
		* `defined_tags` - Usage of predefined tag keys. These predefined keys are scoped to namespaces.

			Tags defined here will be copied verbatim as tags on the Backup resource created by this BackupPolicy.

			Example: `{"foo-namespace.bar-key": "value"}` 
		* `freeform_tags` - Simple key-value pair applied without any predefined name, type or scope. Exists for cross-compatibility only.

			Tags defined here will be copied verbatim as tags on the Backup resource created by this BackupPolicy.

			Example: `{"bar-key": "value"}` 
		* `is_enabled` - If automated backups are enabled or disabled.
		* `pitr_policy` - The PITR policy for the DB System.
			* `is_enabled` - Specifies if PITR is enabled or disabled.
		* `retention_in_days` - The number of days automated backups are retained. 
		* `window_start_time` - The start of a 30-minute window of time in which daily, automated backups occur.

			This should be in the format of the "Time" portion of an RFC3339-formatted timestamp. Any second or sub-second time data will be truncated to zero.

			At some point in the window, the system may incur a brief service disruption as the backup is performed.

			If not defined, a window is selected from the following Region-based time-spans:
			* eu-frankfurt-1: 20:00 - 04:00 UTC
			* us-ashburn-1: 03:00 - 11:00 UTC
			* uk-london-1: 06:00 - 14:00 UTC
			* ap-tokyo-1: 13:00 - 21:00
			* us-phoenix-1: 06:00 - 14:00 
	* `compartment_id` - The OCID of the compartment the DB System belongs in.
	* `configuration_id` - The OCID of the Configuration to be used for Instances in this DB System.
	* `crash_recovery` - Whether to run the DB System with InnoDB Redo Logs and the Double Write Buffer enabled or disabled, and whether to enable or disable syncing of the Binary Logs. 
	* `data_storage_size_in_gb` - Initial size of the data volume in GiBs that will be created and attached. 
	* `database_management` - Whether to enable monitoring via the Database Management service. 
	* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
	* `deletion_policy` - The Deletion policy for the DB System.
		* `automatic_backup_retention` - Specifies if any automatic backups created for a DB System should be retained or deleted when the DB System is deleted. 
		* `final_backup` - Specifies whether or not a backup is taken when the DB System is deleted. REQUIRE_FINAL_BACKUP: a backup is taken if the DB System is deleted. SKIP_FINAL_BACKUP: a backup is not taken if the DB System is deleted. 
		* `is_delete_protected` - Specifies whether the DB System can be deleted. Set to true to prevent deletion, false (default) to allow. 
	* `description` - User-provided data about the DB System.
	* `display_name` - The user-friendly name for the DB System. It does not have to be unique.
	* `endpoints` - The network endpoints available for this DB System. 
		* `hostname` - The network address of the DB System.
		* `ip_address` - The IP address the DB System is configured to listen on.
		* `modes` - The access modes from the client that this endpoint supports.
		* `port` - The port the MySQL instance listens on.
		* `port_x` - The network port where to connect to use this endpoint using the X protocol.
		* `resource_id` - The OCID of the resource that this endpoint is attached to.
		* `resource_type` - The type of endpoint that clients and connectors can connect to.
		* `status` - The state of the endpoints, as far as it can seen from the DB System. There may be some inconsistency with the actual state of the MySQL service. 
		* `status_details` - Additional information about the current endpoint status.
	* `fault_domain` - The name of the Fault Domain the DB System is located in. 
	* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
	* `hostname_label` - The hostname for the primary endpoint of the DB System. Used for DNS. The value is the hostname portion of the primary private IP's fully qualified domain name (FQDN) (for example, "dbsystem-1" in FQDN "dbsystem-1.subnet123.vcn1.oraclevcn.com"). Must be unique across all VNICs in the subnet and comply with RFC 952 and RFC 1123. 
	* `id` - The OCID of the DB System.
	* `ip_address` - The IP address the DB System is configured to listen on. A private IP address of the primary endpoint of the DB System. Must be an available IP address within the subnet's CIDR. This will be a "dotted-quad" style IPv4 address. 
	* `is_highly_available` - Specifies if the DB System is highly available. 
	* `maintenance` - The Maintenance Policy for the DB System or Read Replica that this model is included in. 
		* `window_start_time` - The start time of the maintenance window.

			This string is of the format: "{day-of-week} {time-of-day}".

			"{day-of-week}" is a case-insensitive string like "mon", "tue", &c.

			"{time-of-day}" is the "Time" portion of an RFC3339-formatted timestamp. Any second or sub-second time data will be truncated to zero.

			If you set the read replica maintenance window to "" or if not specified, the read replica is set same as the DB system maintenance window. 
	* `mysql_version` - Name of the MySQL Version in use for the DB System.
	* `port` - The port for primary endpoint of the DB System to listen on.
	* `port_x` - The network port on which X Plugin listens for TCP/IP connections. This is the X Plugin equivalent of port. 
	* `shape_name` - The shape of the primary instances of the DB System. The shape determines resources allocated to a DB System - CPU cores and memory for VM shapes; CPU cores, memory and storage for non-VM (or bare metal) shapes. To get a list of shapes, use (the [ListShapes](https://docs.cloud.oracle.com/iaas/api/#/en/mysql/20181021/ShapeSummary/ListShapes) operation. 
	* `subnet_id` - The OCID of the subnet the DB System is associated with. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - A user-supplied description for the backup.
* `display_name` - A user-supplied display name for the backup.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - OCID of the backup itself
* `lifecycle_details` - Additional information about the current lifecycleState.
* `mysql_version` - The MySQL server version of the DB System used for backup.
* `retention_in_days` - Number of days to retain this backup.
* `shape_name` - The shape of the DB System instance used for backup.
* `state` - The state of the backup.
* `time_created` - The time the backup record was created.
* `time_updated` - The time at which the backup was updated.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Mysql Backup
	* `update` - (Defaults to 20 minutes), when updating the Mysql Backup
	* `delete` - (Defaults to 20 minutes), when destroying the Mysql Backup


## Import

MysqlBackups can be imported using the `id`, e.g.

```
$ terraform import oci_mysql_mysql_backup.test_mysql_backup "id"
```

