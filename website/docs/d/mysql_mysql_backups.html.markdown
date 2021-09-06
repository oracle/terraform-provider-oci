---
subcategory: "MySQL Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_mysql_mysql_backups"
sidebar_current: "docs-oci-datasource-mysql-mysql_backups"
description: |-
  Provides the list of Mysql Backups in Oracle Cloud Infrastructure MySQL Database service
---

# Data Source: oci_mysql_mysql_backups
This data source provides the list of Mysql Backups in Oracle Cloud Infrastructure MySQL Database service.

Get a list of DB System backups.


## Example Usage

```hcl
data "oci_mysql_mysql_backups" "test_mysql_backups" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	backup_id = oci_mysql_mysql_backup.test_backup.id
	creation_type = var.mysql_backup_creation_type
	db_system_id = oci_mysql_mysql_db_system.test_db_system.id
	display_name = var.mysql_backup_display_name
	state = var.mysql_backup_state
}
```

## Argument Reference

The following arguments are supported:

* `backup_id` - (Optional) Backup OCID
* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `creation_type` - (Optional) Backup creationType
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
	* `data_storage_size_in_gb` - Initial size of the data volume in GiBs that will be created and attached. 
	* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
	* `description` - User-provided data about the DB System.
	* `display_name` - The user-friendly name for the DB System. It does not have to be unique.
	* `endpoints` - The network endpoints available for this DB System. 
		* `hostname` - The network address of the DB System.
		* `ip_address` - The IP address the DB System is configured to listen on.
		* `modes` - The access modes from the client that this endpoint supports.
		* `port` - The port the MySQL instance listens on.
		* `port_x` - The network port where to connect to use this endpoint using the X protocol.
		* `status` - The state of the endpoints, as far as it can seen from the DB System. There may be some inconsistency with the actual state of the MySQL service. 
		* `status_details` - Additional information about the current endpoint status.
	* `fault_domain` - The name of the Fault Domain the DB System is located in. 
	* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
	* `hostname_label` - The hostname for the primary endpoint of the DB System. Used for DNS. The value is the hostname portion of the primary private IP's fully qualified domain name (FQDN) (for example, "dbsystem-1" in FQDN "dbsystem-1.subnet123.vcn1.oraclevcn.com"). Must be unique across all VNICs in the subnet and comply with RFC 952 and RFC 1123. 
	* `id` - The OCID of the DB System.
	* `ip_address` - The IP address the DB System is configured to listen on. A private IP address of the primary endpoint of the DB System. Must be an available IP address within the subnet's CIDR. This will be a "dotted-quad" style IPv4 address. 
	* `is_highly_available` - If the policy is to enable high availability of the instance, by maintaining secondary/failover capacity as necessary. 
	* `maintenance` - The Maintenance Policy for the DB System. 
		* `window_start_time` - The start time of the maintenance window.

			This string is of the format: "{day-of-week} {time-of-day}".

			"{day-of-week}" is a case-insensitive string like "mon", "tue", &c.

			"{time-of-day}" is the "Time" portion of an RFC3339-formatted timestamp. Any second or sub-second time data will be truncated to zero. 
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

