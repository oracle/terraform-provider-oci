---
subcategory: "Mysql"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_mysql_mysql_db_system"
sidebar_current: "docs-oci-resource-mysql-mysql_db_system"
description: |-
  Provides the Mysql Db System resource in Oracle Cloud Infrastructure Mysql service
---

# oci_mysql_mysql_db_system
This resource provides the Mysql Db System resource in Oracle Cloud Infrastructure Mysql service.

Creates and launches a DB System.


## Example Usage

```hcl
resource "oci_mysql_mysql_db_system" "test_mysql_db_system" {
	#Required
	admin_password = var.mysql_db_system_admin_password
	admin_username = var.mysql_db_system_admin_username
	availability_domain = var.mysql_db_system_availability_domain
	compartment_id = var.compartment_id
	shape_name = var.mysql_shape_name
	subnet_id = oci_core_subnet.test_subnet.id

	#Optional
	backup_policy {

		#Optional
		defined_tags = {"foo-namespace.bar-key"= "value"}
		freeform_tags = {"bar-key"= "value"}
		is_enabled = var.mysql_db_system_backup_policy_is_enabled
		retention_in_days = var.mysql_db_system_backup_policy_retention_in_days
		window_start_time = var.mysql_db_system_backup_policy_window_start_time
	}
	configuration_id = oci_audit_configuration.test_configuration.id
	data_storage_size_in_gb = var.mysql_db_system_data_storage_size_in_gb
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.mysql_db_system_description
	display_name = var.mysql_db_system_display_name
	fault_domain = var.mysql_db_system_fault_domain
	freeform_tags = {"bar-key"= "value"}
	hostname_label = var.mysql_db_system_hostname_label
	ip_address = var.mysql_db_system_ip_address
	maintenance {
		#Required
		window_start_time = var.mysql_db_system_maintenance_window_start_time
	}
	port = var.mysql_db_system_port
	port_x = var.mysql_db_system_port_x
	source {
		#Required
		source_type = var.mysql_db_system_source_source_type

		#Optional
		backup_id = oci_mysql_mysql_backup.test_backup.id
	}
}
```

## Argument Reference

The following arguments are supported:

* `admin_password` - (Required) The password for the administrative user. The password must be between 8 and 32 characters long, and must contain at least 1 numeric character, 1 lowercase character, 1 uppercase character, and 1 special (nonalphanumeric) character. 
* `admin_username` - (Required) The username for the administrative user.
* `availability_domain` - (Required) The Availability Domain where the primary instance should be located. 
* `backup_policy` - (Optional) (Updatable) Backup policy as optionally used for DB System Creation. 
	* `defined_tags` - (Optional) (Updatable) Usage of predefined tag keys. These predefined keys are scoped to namespaces.

		Tags defined here will be copied verbatim as tags on the Backup resource created by this BackupPolicy.

		Example: `{"foo-namespace.bar-key": "value"}` 
	* `freeform_tags` - (Optional) (Updatable) Simple key-value pair applied without any predefined name, type or scope. Exists for cross-compatibility only.

		Tags defined here will be copied verbatim as tags on the Backup resource created by this BackupPolicy.

		Example: `{"bar-key": "value"}` 
	* `is_enabled` - (Optional) (Updatable) Specifies if automatic backups are enabled. 
	* `retention_in_days` - (Optional) (Updatable) Number of days to retain an automatic backup.
	* `window_start_time` - (Optional) (Updatable) The start of a 30-minute window of time in which daily, automated backups occur.

		This should be in the format of the "Time" portion of an RFC3339-formatted timestamp. Any second or sub-second time data will be truncated to zero.

		At some point in the window, the system may incur a brief service disruption as the backup is performed. 
* `compartment_id` - (Required) The OCID of the compartment.
* `configuration_id` - (Optional) The OCID of the Configuration to be used for this DB System.
* `data_storage_size_in_gb` - (Optional) Initial size of the data volume in GBs that will be created and attached. Keep in mind that this only specifies the size of the database data volume, the log volume for the database will be scaled appropriately with its shape. It is required if you are creating a new database. It cannot be set if you are creating a database from a backup.
* `defined_tags` - (Optional) (Updatable) Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) User-provided data about the DB System.
* `display_name` - (Optional) (Updatable) The user-friendly name for the DB System. It does not have to be unique.
* `fault_domain` - (Optional) The name of the Fault Domain the DB System is located in. 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `hostname_label` - (Optional) The hostname for the primary endpoint of the DB System. Used for DNS.

	The value is the hostname portion of the primary private IP's fully qualified domain name (FQDN) (for example, "dbsystem-1" in FQDN "dbsystem-1.subnet123.vcn1.oraclevcn.com").

	Must be unique across all VNICs in the subnet and comply with RFC 952 and RFC 1123. 
* `ip_address` - (Optional) The IP address the DB System is configured to listen on. A private IP address of your choice to assign to the primary endpoint of the DB System. Must be an available IP address within the subnet's CIDR. If you don't specify a value, Oracle automatically assigns a private IP address from the subnet. This should be a "dotted-quad" style IPv4 address. 
* `maintenance` - (Optional) (Updatable) The Maintenance Policy for the DB System. `maintenance` and `backup_policy` cannot be updated in the same request.
	* `window_start_time` - (Required) (Updatable) The start of the 2 hour maintenance window.

		This string is of the format: "{day-of-week} {time-of-day}".

		"{day-of-week}" is a case-insensitive string like "mon", "tue", &c.

		"{time-of-day}" is the "Time" portion of an RFC3339-formatted timestamp. Any second or sub-second time data will be truncated to zero. 
* `port` - (Optional) The port for primary endpoint of the DB System to listen on.
* `port_x` - (Optional) The TCP network port on which X Plugin listens for connections. This is the X Plugin equivalent of port. 
* `shape_name` - (Required) The name of the shape. The shape determines the resources allocated
	* CPU cores and memory for VM shapes; CPU cores, memory and storage for non-VM (or bare metal) shapes. To get a list of shapes, use the [ListShapes](https://docs.cloud.oracle.com/iaas/api/#/en/mysql/20190415/ShapeSummary/ListShapes) operation. 
* `source` - (Optional) Parameters detailing how to provision the initial data of the system. 
	* `backup_id` - (Required when source_type=BACKUP) The OCID of the backup to be used as the source for the new DB System. 
	* `source_type` - (Required) The specific source identifier. Use `BACKUP` for creating a new database by restoring from a backup.
* `subnet_id` - (Required) The OCID of the subnet the DB System is associated with. 
* `state` - (Optional) (Updatable) The target state for the DB System. Could be set to `ACTIVE` or `INACTIVE`. 
* `shutdown_type` - (Optional) It is applicable only for stopping a DB System. Could be set to `FAST`, `SLOW` or `IMMEDIATE`. Default value is `FAST`.

** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `analytics_cluster` - 
	* `cluster_size` - The number of analytics-processing compute instances, of the specified shape, in the Analytics Cluster. 
	* `shape_name` - The shape determines resources to allocate to the Analytics Cluster nodes - CPU cores, memory. 
	* `state` - The current state of the MySQL Analytics Cluster.
	* `time_created` - The date and time the Analytics Cluster was created, as described by [RFC 3339](https://tools.ietf.org/rfc/rfc3339).
	* `time_updated` - The time the Analytics Cluster was last updated, as described by [RFC 3339](https://tools.ietf.org/rfc/rfc3339).
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
* `is_analytics_cluster_attached` - If the DB System has an Analytics Cluster attached. 
* `lifecycle_details` - Additional information about the current lifecycleState.
* `maintenance` - The Maintenance Policy for the DB System. 
	* `window_start_time` - The start time of the maintenance window.

		This string is of the format: "{day-of-week} {time-of-day}".

		"{day-of-week}" is a case-insensitive string like "mon", "tue", &c.

		"{time-of-day}" is the "Time" portion of an RFC3339-formatted timestamp. Any second or sub-second time data will be truncated to zero. 
* `mysql_version` - Name of the MySQL Version in use for the DB System.
* `port` - The port for primary endpoint of the DB System to listen on.
* `port_x` - The network port on which X Plugin listens for TCP/IP connections. This is the X Plugin equivalent of port. 
* `shape_name` - The shape of the primary instances of the DB System. The shape determines resources allocated to a DB System - CPU cores and memory for VM shapes; CPU cores, memory and storage for non-VM (or bare metal) shapes. To get a list of shapes, use (the [ListShapes](https://docs.cloud.oracle.com/iaas/api/#/en/mysql/20181021/ShapeSummary/ListShapes) operation. 
* `source` - Parameters detailing how to provision the initial data of the DB System. 
	* `backup_id` - The OCID of the backup to be used as the source for the new DB System. 
	* `source_type` - The specific source identifier. 
* `state` - The current state of the DB System.
* `subnet_id` - The OCID of the subnet the DB System is associated with. 
* `time_created` - The date and time the DB System was created.
* `time_updated` - The time the DB System was last updated.

## Import

MysqlDbSystems can be imported using the `id`, e.g.

```
$ terraform import oci_mysql_mysql_db_system.test_mysql_db_system "id"
```

