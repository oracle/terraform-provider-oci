---
subcategory: "MySQL Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_mysql_mysql_db_systems"
sidebar_current: "docs-oci-datasource-mysql-mysql_db_systems"
description: |-
  Provides the list of Mysql Db Systems in Oracle Cloud Infrastructure MySQL Database service
---

# Data Source: oci_mysql_mysql_db_systems
This data source provides the list of Mysql Db Systems in Oracle Cloud Infrastructure MySQL Database service.

Get a list of DB Systems in the specified compartment.
The default sort order is by timeUpdated, descending.


## Example Usage

```hcl
data "oci_mysql_mysql_db_systems" "test_mysql_db_systems" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	configuration_id = var.mysql_configuration_id
	database_management = var.mysql_db_system_database_management
	db_system_id = oci_mysql_mysql_db_system.test_db_system.id
	display_name = var.mysql_db_system_display_name
	is_heat_wave_cluster_attached = var.mysql_db_system_is_heat_wave_cluster_attached
	is_up_to_date = var.mysql_db_system_is_up_to_date
	state = var.mysql_db_system_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `configuration_id` - (Optional) The requested Configuration instance.
* `database_management` - (Optional) Filter DB Systems by their Database Management configuration. 
* `db_system_id` - (Optional) The DB System [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `display_name` - (Optional) A filter to return only the resource matching the given display name exactly.
* `is_heat_wave_cluster_attached` - (Optional) If true, return only DB Systems with a HeatWave cluster attached, if false return only DB Systems with no HeatWave cluster attached. If not present, return all DB Systems. 
* `is_up_to_date` - (Optional) Filter instances if they are using the latest revision of the Configuration they are associated with. 
* `state` - (Optional) DbSystem Lifecycle State


## Attributes Reference

The following attributes are exported:

* `db_systems` - The list of db_systems.

### MysqlDbSystem Reference

The following attributes are exported:

* `access_mode` - The access mode indicating if the database access is unrestricted (to all MySQL user accounts),  or restricted (to only certain users with specific privileges):
	* UNRESTRICTED: the access to the database is not restricted;
	* RESTRICTED: the access is allowed only to users with specific privileges;  RESTRICTED will correspond to setting the MySQL system variable  [offline_mode](https://dev.mysql.com/doc/en/server-system-variables.html#sysvar_offline_mode) to ON. 
* `availability_domain` - The availability domain on which to deploy the Read/Write endpoint. This defines the preferred primary instance.

	In a failover scenario, the Read/Write endpoint is redirected to one of the other availability domains and the MySQL instance in that domain is promoted to the primary instance. This redirection does not affect the IP address of the DB System in any way.

	For a standalone DB System, this defines the availability domain in which the DB System is placed. 
* `backup_policy` - The Backup policy for the DB System.
	* `copy_policies` - List of policies of a DB system to schedule cross-region DB system backup copy.

		The policy includes the name of the destination region to which the DB system backup will be copied, and an optional parameter which specifies the retention period of the copied DB system backup in days.

		**Note:** Currently, only one policy can be specified in the list. 
		* `backup_copy_retention_in_days` - Number of days to retain the copied DB system backup.
		* `copy_to_region` - The destination region name to which the DB system backup will be copied.
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
	* `soft_delete` - Retains the backup to be deleted due to the retention policy in DELETE SCHEDULED state for 7 days before permanently deleting it. 
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
	* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
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
* `customer_contacts` - The list of customer email addresses that receive information from Oracle about the specified Oracle Cloud Infrastructure DB System resource.  Oracle uses these email addresses to send notifications about planned and unplanned software maintenance updates, information about system hardware, and other information needed by administrators.  Up to 10 email addresses can be added to the customer contacts for a DB System. 
	* `email` - The email address used by Oracle to send notifications regarding the DB System. 
* `data_storage_size_in_gb` - Initial size of the data volume in GiBs that will be created and attached. 
* `database_management` - Whether to enable monitoring via the Database Management service. 
* `database_mode` - The database mode indicating the types of statements that are allowed to run in the the DB system. This mode applies only to statements run by user connections. Replicated write statements continue  to be allowed regardless of the DatabaseMode.
	* READ_WRITE: allow running read and write statements on the DB system;
	* READ_ONLY: only allow running read statements on the DB system. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `deletion_policy` - The Deletion policy for the DB System.
	* `automatic_backup_retention` - Specifies if any automatic backups created for a DB System should be retained or deleted when the DB System is deleted. 
	* `final_backup` - Specifies whether or not a backup is taken when the DB System is deleted. REQUIRE_FINAL_BACKUP: a backup is taken if the DB System is deleted. SKIP_FINAL_BACKUP: a backup is not taken if the DB System is deleted. 
	* `is_delete_protected` - Specifies whether the DB System can be deleted. Set to true to prevent deletion, false (default) to allow. 
* `description` - User-provided data about the DB System.
* `display_name` - The user-friendly name for the DB System. It does not have to be unique.
* `encrypt_data` - Encrypt data details. 
	* `key_generation_type` - Select whether to use Oracle-managed key (SYSTEM) or your own key (BYOK).
	* `key_id` - The OCID of the key to use.
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
	* `maintenance_schedule_type` - The maintenance schedule type of the DB system. EARLY:   Maintenance schedule follows a cycle where upgrades are performed when versions become deprecated. REGULAR: Maintenance schedule follows the normal cycle where upgrades are performed when versions become unavailable. 
	* `target_version` - The version that is expected to be targeted during the next scheduled maintenance run. 
	* `time_scheduled` - The time the scheduled maintenance is expected to start, as described by [RFC 3339](https://tools.ietf.org/rfc/rfc3339). 
	* `version_preference` - The preferred version to target when performing an automatic MySQL upgrade.

		OLDEST: Choose the oldest available MySQL version based on the current version of the DB System. SECOND_NEWEST: Choose the MySQL version before the newest for auto-upgrade. NEWEST: Choose the latest and greatest MySQL version available for auto-upgrade. 
	* `version_track_preference` - The preferred version track to target when performing an automatic MySQL upgrade. LONG_TERM_SUPPORT: No MySQL database behavior changes. INNOVATION:        Provides access to the latest features and all bug fixes. FOLLOW:            Follows the track of the current MySQL version. 
	* `window_start_time` - The start time of the maintenance window.

		This string is of the format: "{day-of-week} {time-of-day}".

		"{day-of-week}" is a case-insensitive string like "mon", "tue", &c.

		"{time-of-day}" is the "Time" portion of an RFC3339-formatted timestamp. Any second or sub-second time data will be truncated to zero.

		If you set the read replica maintenance window to "" or if not specified, the read replica is set same as the DB system maintenance window. 
* `mysql_version` - Name of the MySQL Version in use for the DB System.
* `nsg_ids` - Network Security Group OCIDs used for the VNIC attachment.
* `point_in_time_recovery_details` - Point-in-time Recovery details like earliest and latest recovery time point for the DB System. 
	* `time_earliest_recovery_point` - Earliest recovery time point for the DB System, as described by [RFC 3339](https://tools.ietf.org/rfc/rfc3339). 
	* `time_latest_recovery_point` - Latest recovery time point for the DB System, as described by [RFC 3339](https://tools.ietf.org/rfc/rfc3339). 
* `port` - The port for primary endpoint of the DB System to listen on.
* `port_x` - The network port on which X Plugin listens for TCP/IP connections. This is the X Plugin equivalent of port. 
* `read_endpoint` - The read endpoint of a DB System. 
	* `exclude_ips` - A list of IP addresses of read replicas that are excluded from serving read requests. 
	* `is_enabled` - Specifies if the DB System read endpoint is enabled or not. 
	* `read_endpoint_hostname_label` - The hostname for the read endpoint of the DB System. Used for DNS.

		The value is the hostname portion of the primary private IP's fully qualified domain name (FQDN)  (for example, "dbsystem-1" in FQDN "dbsystem-1.subnet123.vcn1.oraclevcn.com").

		Must be unique across all VNICs in the subnet and comply with RFC 952 and RFC 1123. 
	* `read_endpoint_ip_address` - The IP address the DB System read endpoint is configured to listen on. A private IP address of your choice to assign to the read endpoint of the DB System. Must be an available IP address within the subnet's CIDR. If you don't specify a value, Oracle automatically assigns a private IP address from the subnet. This should be a "dotted-quad" style IPv4 address. 
* `rest` - REST configuration details. 
	* `configuration` - Select how REST is configured across the DB System instances.
	* `port` - The port for REST to listen on. Supported port numbers are 443 and from 1024 to 65535.
* `secure_connections` - Secure connection configuration details. 
	* `certificate_generation_type` - Select whether to use MySQL Database Service-managed certificate (SYSTEM) or your own certificate (BYOC). 
	* `certificate_id` - The OCID of the certificate to use.
* `security_attributes` - Security Attributes for this resource. Each key is predefined and scoped to a namespace. For more information, see [ZPR Artifacts](https://docs.oracle.com/en-us/iaas/Content/zero-trust-packet-routing/zpr-artifacts.htm). Example: `{"Oracle-ZPR": {"MaxEgressCount": {"value": "42", "mode": "audit"}}}` 
* `shape_name` - The shape of the primary instances of the DB System. The shape determines resources allocated to a DB System - CPU cores and memory for VM shapes; CPU cores, memory and storage for non-VM (or bare metal) shapes. To get a list of shapes, use (the [ListShapes](https://docs.cloud.oracle.com/iaas/api/#/en/mysql/20181021/ShapeSummary/ListShapes) operation. 
* `source` - Parameters detailing how to provision the initial data of the DB System. 
	* `backup_id` - The OCID of the backup to be used as the source for the new DB System. 
	* `db_system_id` - The OCID of the DB System from which a backup shall be selected to be restored when creating the new DB System. Use this together with recovery point to perform a point in time recovery operation. 
	* `recovery_point` - The date and time, as per RFC 3339, of the change up to which the new DB System shall be restored to, using a backup and logs from the original DB System. In case no point in time is specified, then this new DB System shall be restored up to the latest change recorded for the original DB System. 
	* `source_type` - The specific source identifier. 
* `state` - The current state of the DB System.
* `subnet_id` - The OCID of the subnet the DB System is associated with. 
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the DB System was created.
* `time_updated` - The time the DB System was last updated.

