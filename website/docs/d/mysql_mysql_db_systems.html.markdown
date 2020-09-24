---
subcategory: "Mysql"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_mysql_mysql_db_systems"
sidebar_current: "docs-oci-datasource-mysql-mysql_db_systems"
description: |-
  Provides the list of Mysql Db Systems in Oracle Cloud Infrastructure Mysql service
---

# Data Source: oci_mysql_mysql_db_systems
This data source provides the list of Mysql Db Systems in Oracle Cloud Infrastructure Mysql service.

Get a list of DB Systems in the specified compartment.
The default sort order is by timeUpdated, descending.


## Example Usage

```hcl
data "oci_mysql_mysql_db_systems" "test_mysql_db_systems" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	configuration_id = var.mysql_configuration_id
	db_system_id = oci_mysql_mysql_db_system.test_db_system.id
	display_name = var.mysql_db_system_display_name
	is_analytics_cluster_attached = var.mysql_db_system_is_analytics_cluster_attached
	is_up_to_date = var.mysql_db_system_is_up_to_date
	state = var.mysql_db_system_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `configuration_id` - (Optional) The requested Configuration instance.
* `db_system_id` - (Optional) The DB System [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `display_name` - (Optional) A filter to return only the resource matching the given display name exactly.
* `is_analytics_cluster_attached` - (Optional) If true, return only DB Systems with an Analytics Cluster attached, if false return only DB Systems with no Analytics Cluster attached. If not present, return all DB Systems. 
* `is_up_to_date` - (Optional) Filter instances if they are using the latest revision of the Configuration they are associated with. 
* `state` - (Optional) DbSystem Lifecycle State


## Attributes Reference

The following attributes are exported:

* `db_systems` - The list of db_systems.

### MysqlDbSystem Reference

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
* `defined_tags` - Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"foo-namespace.bar-key": "value"}` 
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
* `freeform_tags` - Simple key-value pair applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
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

