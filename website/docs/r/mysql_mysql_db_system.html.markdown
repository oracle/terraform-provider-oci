---
subcategory: "MySQL Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_mysql_mysql_db_system"
sidebar_current: "docs-oci-resource-mysql-mysql_db_system"
description: |-
  Provides the Mysql Db System resource in Oracle Cloud Infrastructure MySQL Database service
---

# oci_mysql_mysql_db_system
This resource provides the Mysql Db System resource in Oracle Cloud Infrastructure MySQL Database service.

Creates and launches a DB System.


## Example Usage

```hcl
resource "oci_mysql_mysql_db_system" "test_mysql_db_system" {
	#Required
	availability_domain = var.mysql_db_system_availability_domain
	compartment_id = var.compartment_id
	shape_name = var.mysql_shape_name
	subnet_id = oci_core_subnet.test_subnet.id

	#Optional
	admin_password = var.mysql_db_system_admin_password
	admin_username = var.mysql_db_system_admin_username
	backup_policy {

		#Optional
		defined_tags = {"foo-namespace.bar-key"= "value"}
		freeform_tags = {"bar-key"= "value"}
		is_enabled = var.mysql_db_system_backup_policy_is_enabled
		pitr_policy {
			#Required
			is_enabled = var.mysql_db_system_backup_policy_pitr_policy_is_enabled
		}
		retention_in_days = var.mysql_db_system_backup_policy_retention_in_days
		window_start_time = var.mysql_db_system_backup_policy_window_start_time
	}
	configuration_id = oci_audit_configuration.test_configuration.id
	crash_recovery = var.mysql_db_system_crash_recovery
	data_storage {

		#Optional
		is_auto_expand_storage_enabled = var.mysql_db_system_data_storage_is_auto_expand_storage_enabled
		max_storage_size_in_gbs = var.mysql_db_system_data_storage_max_storage_size_in_gbs
	}
	data_storage_size_in_gb = var.mysql_db_system_data_storage_size_in_gb
	database_management = var.mysql_db_system_database_management
	defined_tags = {"foo-namespace.bar-key"= "value"}
	deletion_policy {

		#Optional
		automatic_backup_retention = var.mysql_db_system_deletion_policy_automatic_backup_retention
		final_backup = var.mysql_db_system_deletion_policy_final_backup
		is_delete_protected = var.mysql_db_system_deletion_policy_is_delete_protected
	}
	description = var.mysql_db_system_description
	display_name = var.mysql_db_system_display_name
	fault_domain = var.mysql_db_system_fault_domain
	freeform_tags = {"bar-key"= "value"}
	hostname_label = var.mysql_db_system_hostname_label
	ip_address = var.mysql_db_system_ip_address
	is_highly_available = var.mysql_db_system_is_highly_available
	maintenance {
		#Required
		window_start_time = var.mysql_db_system_maintenance_window_start_time
	}
	port = var.mysql_db_system_port
	port_x = var.mysql_db_system_port_x
	secure_connections {
		#Required
		certificate_generation_type = var.mysql_db_system_secure_connections_certificate_generation_type

		#Optional
		certificate_id = oci_apigateway_certificate.test_certificate.id
	}
	source {
		#Required
		source_type = var.mysql_db_system_source_source_type

		#Optional
		# source_url = var.mysql_db_system_source_source_url
		backup_id = oci_mysql_mysql_backup.test_backup.id
	}
}
```

## Argument Reference

The following arguments are supported:

* `admin_password` - (Optional) The password for the administrative user. The password must be between 8 and 32 characters long, and must contain at least 1 numeric character, 1 lowercase character, 1 uppercase character, and 1 special (nonalphanumeric) character. 
* `admin_username` - (Optional) The username for the administrative user.
* `availability_domain` - (Required) The availability domain on which to deploy the Read/Write endpoint. This defines the preferred primary instance.

	In a failover scenario, the Read/Write endpoint is redirected to one of the other availability domains and the MySQL instance in that domain is promoted to the primary instance. This redirection does not affect the IP address of the DB System in any way.

	For a standalone DB System, this defines the availability domain in which the DB System is placed. 
* `backup_policy` - (Optional) (Updatable) Backup policy as optionally used for DB System Creation. 
	* `defined_tags` - (Optional) (Updatable) Usage of predefined tag keys. These predefined keys are scoped to namespaces.

		Tags defined here will be copied verbatim as tags on the Backup resource created by this BackupPolicy.

		Example: `{"foo-namespace.bar-key": "value"}` 
	* `freeform_tags` - (Optional) (Updatable) Simple key-value pair applied without any predefined name, type or scope. Exists for cross-compatibility only.

		Tags defined here will be copied verbatim as tags on the Backup resource created by this BackupPolicy.

		Example: `{"bar-key": "value"}` 
	* `is_enabled` - (Optional) (Updatable) Specifies if automatic backups are enabled. 
	* `pitr_policy` - (Optional) (Updatable) The PITR policy for the DB System.
		* `is_enabled` - (Required) (Updatable) Specifies if PITR is enabled or disabled.
	* `retention_in_days` - (Optional) (Updatable) Number of days to retain an automatic backup.
	* `window_start_time` - (Optional) (Updatable) The start of a 30-minute window of time in which daily, automated backups occur.

		This should be in the format of the "Time" portion of an RFC3339-formatted timestamp. Any second or sub-second time data will be truncated to zero.

		At some point in the window, the system may incur a brief service disruption as the backup is performed. 
* `compartment_id` - (Required) The OCID of the compartment.
* `configuration_id` - (Optional) (Updatable) The OCID of the Configuration to be used for this DB System.
* `crash_recovery` - (Optional) (Updatable) Whether to run the DB System with InnoDB Redo Logs and the Double Write Buffer enabled or disabled, and whether to enable or disable syncing of the Binary Logs. 
* `data_storage` - (Optional) (Updatable) Data Storage configuration properties. 
	* `is_auto_expand_storage_enabled` - (Optional) (Updatable) Enable/disable automatic storage expansion. When set to true, the DB System will automatically add storage incrementally up to the value specified in maxStorageSizeInGBs. 
	* `max_storage_size_in_gbs` - (Optional) (Updatable) Maximum storage size this DB System can expand to. When isAutoExpandStorageEnabled is set to true, the DB System will add storage incrementally up to this value.

		DB Systems with an initial storage size of 400 GB or less can be expanded up to 32 TB. DB Systems with an initial storage size between 401-800 GB can be expanded up to 64 TB. DB Systems with an initial storage size between 801-1200 GB can be expanded up to 96 TB. DB Systems with an initial storage size of 1201 GB or more can be expanded up to 128 TB.

		It is not possible to decrease data storage size. You cannot set the maximum data storage size to less than either current DB System dataStorageSizeInGBs or allocatedStorageSizeInGBs. 
* `data_storage_size_in_gb` - (Optional) (Updatable) Initial size of the data volume in GBs that will be created and attached. Keep in mind that this only specifies the size of the database data volume, the log volume for the database will be scaled appropriately with its shape. It is required if you are creating a new database. It cannot be set if you are creating a database from a backup.
* `database_management` - (Optional) (Updatable) Whether to enable monitoring via the Database Management service. 
* `defined_tags` - (Optional) (Updatable) Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"foo-namespace.bar-key": "value"}` 
* `deletion_policy` - (Optional) (Updatable) Policy for how the DB System and related resources should be handled at the time of its deletion. 
	* `automatic_backup_retention` - (Optional) (Updatable) Specifies if any automatic backups created for a DB System should be retained or deleted when the DB System is deleted. 
	* `final_backup` - (Optional) (Updatable) Specifies whether or not a backup is taken when the DB System is deleted. REQUIRE_FINAL_BACKUP: a backup is taken if the DB System is deleted. SKIP_FINAL_BACKUP: a backup is not taken if the DB System is deleted. 
	* `is_delete_protected` - (Optional) (Updatable) Specifies whether the DB System can be deleted. Set to true to prevent deletion, false (default) to allow. 
* `description` - (Optional) (Updatable) User-provided data about the DB System.
* `display_name` - (Optional) (Updatable) The user-friendly name for the DB System. It does not have to be unique.
* `fault_domain` - (Optional) The fault domain on which to deploy the Read/Write endpoint. This defines the preferred primary instance.

	In a failover scenario, the Read/Write endpoint is redirected to one of the other fault domains and the MySQL instance in that domain is promoted to the primary instance. This redirection does not affect the IP address of the DB System in any way.

	For a standalone DB System, this defines the fault domain in which the DB System is placed. 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `hostname_label` - (Optional) The hostname for the primary endpoint of the DB System. Used for DNS.

	The value is the hostname portion of the primary private IP's fully qualified domain name (FQDN) (for example, "dbsystem-1" in FQDN "dbsystem-1.subnet123.vcn1.oraclevcn.com").

	Must be unique across all VNICs in the subnet and comply with RFC 952 and RFC 1123. 
* `ip_address` - (Optional) The IP address the DB System is configured to listen on. A private IP address of your choice to assign to the primary endpoint of the DB System. Must be an available IP address within the subnet's CIDR. If you don't specify a value, Oracle automatically assigns a private IP address from the subnet. This should be a "dotted-quad" style IPv4 address. 
* `is_highly_available` - (Optional) (Updatable) Specifies if the DB System is highly available.

	When creating a DB System with High Availability, three instances are created and placed according to your region- and subnet-type. The secondaries are placed automatically in the other two availability or fault domains.  You can choose the preferred location of your primary instance, only. 
* `maintenance` - (Optional) (Updatable) The Maintenance Policy for the DB System or Read Replica that this model is included in. `maintenance` and `backup_policy` cannot be updated in the same request.
	* `window_start_time` - (Required) (Updatable) The start of the 2 hour maintenance window.

		This string is of the format: "{day-of-week} {time-of-day}".

		"{day-of-week}" is a case-insensitive string like "mon", "tue", &c.

		"{time-of-day}" is the "Time" portion of an RFC3339-formatted timestamp. Any second or sub-second time data will be truncated to zero.

		If you set the read replica maintenance window to "" or if not specified, the read replica is set same as the DB system maintenance window. 
* `mysql_version` - (Optional) The specific MySQL version identifier.
* `port` - (Optional) The port for primary endpoint of the DB System to listen on.
* `port_x` - (Optional) The TCP network port on which X Plugin listens for connections. This is the X Plugin equivalent of port. 
* `secure_connections` - (Optional) (Updatable) Secure connection configuration details. 
	* `certificate_generation_type` - (Required) (Updatable) Select whether to use MySQL Database Service-managed certificate (SYSTEM) or your own certificate (BYOC). 
	* `certificate_id` - (Optional) (Updatable) The OCID of the certificate to use.
* `shape_name` - (Required) (Updatable) The name of the shape. The shape determines the resources allocated
	* CPU cores and memory for VM shapes; CPU cores, memory and storage for non-VM (or bare metal) shapes. To get a list of shapes, use the [ListShapes](https://docs.cloud.oracle.com/iaas/api/#/en/mysql/20190415/ShapeSummary/ListShapes) operation. 
* `source` - (Optional) Parameters detailing how to provision the initial data of the system. 
	* `backup_id` - (Required when source_type=BACKUP) The OCID of the backup to be used as the source for the new DB System. 
	* `db_system_id` - (Required when source_type=PITR) The OCID of the DB System from which a backup shall be selected to be restored when creating the new DB System. Use this together with recovery point to perform a point in time recovery operation. 
	* `recovery_point` - (Applicable when source_type=PITR) The date and time, as per RFC 3339, of the change up to which the new DB System shall be restored to, using a backup and logs from the original DB System. In case no point in time is specified, then this new DB System shall be restored up to the latest change recorded for the original DB System. 
	* `source_type` - (Required) The specific source identifier. Use `BACKUP` for creating a new database by restoring from a backup. Use `IMPORTURL` for creating a new database from a URL Object Storage PAR.
	* `source_url` - (Required when source_type=IMPORTURL) The Pre-Authenticated Request (PAR) of a bucket/prefix or PAR of a @.manifest.json object from the Object Storage. Check [Using Pre-Authenticated Requests](https://docs.oracle.com/en-us/iaas/Content/Object/Tasks/usingpreauthenticatedrequests.htm) for information related to PAR creation. Please create PAR with "Permit object reads" access type and "Enable Object Listing" permission when using a bucket/prefix PAR. Please create PAR with "Permit object reads" access type when using a @.manifest.json object PAR. 
* `subnet_id` - (Required) The OCID of the subnet the DB System is associated with. 
* `state` - (Optional) (Updatable) The target state for the DB System. Could be set to `ACTIVE` or `INACTIVE`. 
* `shutdown_type` - (Optional) It is applicable only for stopping a DB System. Could be set to `FAST`, `SLOW` or `IMMEDIATE`. Default value is `FAST`.

** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `availability_domain` - The availability domain on which to deploy the Read/Write endpoint. This defines the preferred primary instance.

	In a failover scenario, the Read/Write endpoint is redirected to one of the other availability domains and the MySQL instance in that domain is promoted to the primary instance. This redirection does not affect the IP address of the DB System in any way.

	For a standalone DB System, this defines the availability domain in which the DB System is placed. 
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
* `channels` - A list with a summary of all the Channels attached to the DB System.
	* `compartment_id` - The OCID of the compartment.
	* `defined_tags` - Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"foo-namespace.bar-key": "value"}` 
	* `display_name` - The user-friendly name for the Channel. It does not have to be unique.
	* `freeform_tags` - Simple key-value pair applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
	* `id` - The OCID of the Channel.
	* `is_enabled` - Whether the Channel has been enabled by the user.
	* `lifecycle_details` - A message describing the state of the Channel.
	* `source` - Parameters detailing how to provision the source for the given Channel.
		* `anonymous_transactions_handling` - Specifies how the replication channel handles replicated transactions without an identifier, enabling replication from a source that does not use transaction-id-based replication to a replica that does. 
			* `last_configured_log_filename` - Specifies one of the coordinates (file) at which the replica should begin reading the source's log. As this value specifies the point where replication starts from, it is only used once, when it starts. It is never used again, unless a new UpdateChannel operation modifies it. 
			* `last_configured_log_offset` - Specifies one of the coordinates (offset) at which the replica should begin reading the source's log. As this value specifies the point where replication starts from, it is only used once, when it starts. It is never used again, unless a new UpdateChannel operation modifies it. 
			* `policy` - Specifies how the replication channel handles anonymous transactions.
			* `uuid` - The UUID that is used as a prefix when generating transaction identifiers for anonymous transactions coming from the source. You can change the UUID later. 
		* `hostname` - The network address of the MySQL instance.
		* `port` - The port the source MySQL instance listens on.
		* `source_type` - The specific source identifier.
		* `ssl_ca_certificate` - The CA certificate of the server used for VERIFY_IDENTITY and VERIFY_CA ssl modes.
			* `certificate_type` - The type of CA certificate.
			* `contents` - The string containing the CA certificate in PEM format.
		* `ssl_mode` - The SSL mode of the Channel.
		* `username` - The name of the replication user on the source MySQL instance. The username has a maximum length of 96 characters. For more information, please see the [MySQL documentation](https://dev.mysql.com/doc/refman/8.0/en/change-master-to.html) 
	* `state` - The state of the Channel.
	* `target` - Details about the Channel target.
		* `applier_username` - The username for the replication applier of the target MySQL DB System.
		* `channel_name` - The case-insensitive name that identifies the replication channel. Channel names must follow the rules defined for [MySQL identifiers](https://dev.mysql.com/doc/refman/8.0/en/identifiers.html). The names of non-Deleted Channels must be unique for each DB System. 
		* `db_system_id` - The OCID of the source DB System.
		* `delay_in_seconds` - Specifies the amount of time, in seconds, that the channel waits before  applying a transaction received from the source. 
		* `filters` - Replication filter rules to be applied at the DB System Channel target. 
			* `type` - The type of the filter rule.

				For details on each type, see [Replication Filtering Rules](https://dev.mysql.com/doc/refman/8.0/en/replication-rules.html) 
			* `value` - The body of the filter rule. This can represent a database, a table, or a database pair (represented as "db1->db2"). For more information, see [Replication Filtering Rules](https://dev.mysql.com/doc/refman/8.0/en/replication-rules.html). 
		* `tables_without_primary_key_handling` - Specifies how a replication channel handles the creation and alteration of tables  that do not have a primary key. 
		* `target_type` - The specific target identifier.
	* `time_created` - The date and time the Channel was created, as described by [RFC 3339](https://tools.ietf.org/rfc/rfc3339). 
	* `time_updated` - The time the Channel was last updated, as described by [RFC 3339](https://tools.ietf.org/rfc/rfc3339). 
* `compartment_id` - The OCID of the compartment the DB System belongs in.
* `configuration_id` - The OCID of the Configuration to be used for Instances in this DB System.
* `crash_recovery` - Whether to run the DB System with InnoDB Redo Logs and the Double Write Buffer enabled or disabled, and whether to enable or disable syncing of the Binary Logs. 
* `current_placement` - The availability domain and fault domain a DB System is placed in.
	* `availability_domain` - The availability domain in which the DB System is placed.
	* `fault_domain` - The fault domain in which the DB System is placed.
* `data_storage` - Data Storage information. 
	* `allocated_storage_size_in_gbs` - The actual allocated storage size for the DB System. This may be higher than dataStorageSizeInGBs if an automatic storage expansion has occurred. 
	* `data_storage_size_in_gb` - User specified size of the data volume. May be less than current allocatedStorageSizeInGBs. 
	* `data_storage_size_limit_in_gbs` - The absolute limit the DB System's storage size may ever expand to, either manually or automatically. This limit is based based on the initial dataStorageSizeInGBs when the DB System was first created. Both dataStorageSizeInGBs and maxDataStorageSizeInGBs can not exceed this value.

		DB Systems with an initial storage size of 400 GB or less can be expanded up to 32 TB. DB Systems with an initial storage size between 401-800 GB can be expanded up to 64 TB. DB Systems with an initial storage size between 801-1200 GB can be expanded up to 96 TB. DB Systems with an initial storage size of 1201 GB or more can be expanded up to 128 TB. 
	* `is_auto_expand_storage_enabled` - Enable/disable automatic storage expansion. When set to true, the DB System will automatically add storage incrementally up to the value specified in maxStorageSizeInGBs. 
	* `max_storage_size_in_gbs` - Maximum storage size this DB System can expand to. When isAutoExpandStorageEnabled is set to true, the DB System will add storage incrementally up to this value.

		DB Systems with an initial storage size of 400 GB or less can be expanded up to 32 TB. DB Systems with an initial storage size between 401-800 GB can be expanded up to 64 TB. DB Systems with an initial storage size between 801-1200 GB can be expanded up to 96 TB. DB Systems with an initial storage size of 1201 GB or more can be expanded up to 128 TB.

		It is not possible to decrease data storage size. You cannot set the maximum data storage size to less than either current DB System dataStorageSizeInGBs or allocatedStorageSizeInGBs. 
* `data_storage_size_in_gb` - DEPRECATED: User specified size of the data volume. May be less than current allocatedStorageSizeInGBs. Replaced by dataStorage.dataStorageSizeInGBs. 
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
* `fault_domain` - The fault domain on which to deploy the Read/Write endpoint. This defines the preferred primary instance.

	In a failover scenario, the Read/Write endpoint is redirected to one of the other fault domains and the MySQL instance in that domain is promoted to the primary instance. This redirection does not affect the IP address of the DB System in any way.

	For a standalone DB System, this defines the fault domain in which the DB System is placed. 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `heat_wave_cluster` - A summary of a HeatWave cluster. 
	* `cluster_size` - The number of analytics-processing compute instances, of the specified shape, in the HeatWave cluster. 
	* `is_lakehouse_enabled` - Lakehouse enabled status for the HeatWave cluster.
	* `shape_name` - The shape determines resources to allocate to the HeatWave nodes - CPU cores, memory. 
	* `state` - The current state of the MySQL HeatWave cluster.
	* `time_created` - The date and time the HeatWave cluster was created, as described by [RFC 3339](https://tools.ietf.org/rfc/rfc3339). 
	* `time_updated` - The time the HeatWave cluster was last updated, as described by [RFC 3339](https://tools.ietf.org/rfc/rfc3339). 
* `hostname_label` - The hostname for the primary endpoint of the DB System. Used for DNS. The value is the hostname portion of the primary private IP's fully qualified domain name (FQDN) (for example, "dbsystem-1" in FQDN "dbsystem-1.subnet123.vcn1.oraclevcn.com"). Must be unique across all VNICs in the subnet and comply with RFC 952 and RFC 1123. 
* `id` - The OCID of the DB System.
* `ip_address` - The IP address the DB System is configured to listen on. A private IP address of the primary endpoint of the DB System. Must be an available IP address within the subnet's CIDR. This will be a "dotted-quad" style IPv4 address. 
* `is_heat_wave_cluster_attached` - If the DB System has a HeatWave Cluster attached. 
* `is_highly_available` - Specifies if the DB System is highly available. 
* `lifecycle_details` - Additional information about the current lifecycleState.
* `maintenance` - The Maintenance Policy for the DB System or Read Replica that this model is included in. 
	* `window_start_time` - The start time of the maintenance window.

		This string is of the format: "{day-of-week} {time-of-day}".

		"{day-of-week}" is a case-insensitive string like "mon", "tue", &c.

		"{time-of-day}" is the "Time" portion of an RFC3339-formatted timestamp. Any second or sub-second time data will be truncated to zero.

		If you set the read replica maintenance window to "" or if not specified, the read replica is set same as the DB system maintenance window. 
* `mysql_version` - Name of the MySQL Version in use for the DB System.
* `point_in_time_recovery_details` - Point-in-time Recovery details like earliest and latest recovery time point for the DB System. 
	* `time_earliest_recovery_point` - Earliest recovery time point for the DB System, as described by [RFC 3339](https://tools.ietf.org/rfc/rfc3339). 
	* `time_latest_recovery_point` - Latest recovery time point for the DB System, as described by [RFC 3339](https://tools.ietf.org/rfc/rfc3339). 
* `port` - The port for primary endpoint of the DB System to listen on.
* `port_x` - The network port on which X Plugin listens for TCP/IP connections. This is the X Plugin equivalent of port. 
* `secure_connections` - Secure connection configuration details. 
	* `certificate_generation_type` - Select whether to use MySQL Database Service-managed certificate (SYSTEM) or your own certificate (BYOC). 
	* `certificate_id` - The OCID of the certificate to use.
* `shape_name` - The shape of the primary instances of the DB System. The shape determines resources allocated to a DB System - CPU cores and memory for VM shapes; CPU cores, memory and storage for non-VM (or bare metal) shapes. To get a list of shapes, use (the [ListShapes](https://docs.cloud.oracle.com/iaas/api/#/en/mysql/20181021/ShapeSummary/ListShapes) operation. 
* `source` - Parameters detailing how to provision the initial data of the DB System. 
	* `backup_id` - The OCID of the backup to be used as the source for the new DB System. 
	* `db_system_id` - The OCID of the DB System from which a backup shall be selected to be restored when creating the new DB System. Use this together with recovery point to perform a point in time recovery operation. 
	* `recovery_point` - The date and time, as per RFC 3339, of the change up to which the new DB System shall be restored to, using a backup and logs from the original DB System. In case no point in time is specified, then this new DB System shall be restored up to the latest change recorded for the original DB System. 
	* `source_type` - The specific source identifier. 
* `state` - The current state of the DB System.
* `subnet_id` - The OCID of the subnet the DB System is associated with. 
* `time_created` - The date and time the DB System was created.
* `time_updated` - The time the DB System was last updated.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 1 hours), when creating the Mysql Db System
	* `update` - (Defaults to 1 hours), when updating the Mysql Db System
	* `delete` - (Defaults to 1 hours), when destroying the Mysql Db System


## Import

MysqlDbSystems can be imported using the `id`, e.g.

```
$ terraform import oci_mysql_mysql_db_system.test_mysql_db_system "id"
```

