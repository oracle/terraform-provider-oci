---
subcategory: "Distributed Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_distributed_database_distributed_autonomous_databases"
sidebar_current: "docs-oci-datasource-distributed_database-distributed_autonomous_databases"
description: |-
  Provides the list of Distributed Autonomous Databases in Oracle Cloud Infrastructure Distributed Database service
---

# Data Source: oci_distributed_database_distributed_autonomous_databases
This data source provides the list of Distributed Autonomous Databases in Oracle Cloud Infrastructure Distributed Database service.

List of Globally distributed autonomous databases.


## Example Usage

```hcl
data "oci_distributed_database_distributed_autonomous_databases" "test_distributed_autonomous_databases" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	db_deployment_type = var.distributed_autonomous_database_db_deployment_type
	display_name = var.distributed_autonomous_database_display_name
	metadata {
	}
	state = var.distributed_autonomous_database_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `db_deployment_type` - (Optional) A filter to return only resources their dbDeploymentType matches the given dbDeploymentType.
* `display_name` - (Optional) A filter to return only Globally distributed autonomous databases that match the entire name given. The match is not case sensitive.
* `metadata` - (Optional) Comma separated names of argument corresponding to which metadata need to be retrived. 
* `state` - (Optional) A filter to return only resources their lifecycleState matches the given lifecycleState.


## Attributes Reference

The following attributes are exported:

* `distributed_autonomous_database_collection` - The list of distributed_autonomous_database_collection.

### DistributedAutonomousDatabase Reference

The following attributes are exported:

* `catalog_details` - Collection of catalogs associated with the Globally distributed autonomous database.
	* `cloud_autonomous_vm_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloudAutonomousVmCluster.
	* `compute_count` - The compute count for the catalog database. It has to be in multiples of 2.
	* `container_database_id` - the identifier of the container database for underlying supporting resource.
	* `data_storage_size_in_gbs` - The data disk group size to be allocated in GBs for the catalog database.
	* `is_auto_scaling_enabled` - Determines the auto-scaling mode for the catalog database.
	* `kms_key_id` - The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
	* `kms_key_version_id` - The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. 
	* `metadata` - Additional metadata related to Globally distributed autonomous database resources.
		* `map` - The map containing key-value pair of additional metadata.
	* `name` - The name of catalog.
	* `peer_cloud_autonomous_vm_cluster_ids` - This field is deprecated. For catalog peer details please refer peerDetails attribute.
	* `peer_details` - Peer details for the catalog with dedicated infrastructure.
		* `cloud_autonomous_vm_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloudAutonomousVmCluster.
		* `container_database_id` - the identifier of the container database for underlying supporting resource.
		* `fast_start_fail_over_lag_limit_in_seconds` - The lag time for my preference based on data loss tolerance in seconds.
		* `is_automatic_failover_enabled` - Indicates whether Automatic Failover is enabled for Autonomous Container Database Dataguard Association 
		* `metadata` - Additional metadata related to Globally distributed autonomous database resources.
			* `map` - The map containing key-value pair of additional metadata.
		* `protection_mode` - The protectionMode for the shard peer.
		* `shard_group` - The name of the shardGroup for the peer.
		* `standby_maintenance_buffer_in_days` - The scheduling detail for the quarterly maintenance window of the standby Autonomous Container Database. This value represents the number of days before schedlued maintenance of the primary database. 
		* `status` - Status of catalog with dedicated infrastructure for the Globally distributed autonomous database.
		* `supporting_resource_id` - the identifier of the underlying supporting resource.
		* `time_created` - The time the catalog peer was created. An RFC3339 formatted datetime string
		* `time_updated` - The time the catalog peer was last updated. An RFC3339 formatted datetime string
	* `shard_group` - The name of the shardGroup for the catalog.
	* `source` - The source of Globally distributed autonomous database type: Use ADB_D for the Globally distributed autonomous database with autonomous dedicated cloudautonomousvmclusters. 
	* `status` - Status of catalog with dedicated infrastructure for the Globally distributed autonomous database.
	* `supporting_resource_id` - the identifier of the underlying supporting resource.
	* `time_created` - The time the catalog was created. An RFC3339 formatted datetime string
	* `time_updated` - The time the catalog was last updated. An RFC3339 formatted datetime string
	* `vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `kmsKeyId` are required for Customer Managed Keys.
* `character_set` - The character set for the database.
* `chunks` - The default number of unique chunks in a shardspace. The value of chunks must be greater than 2 times the size of the largest shardgroup in any shardspace. 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Globally distributed autonomous database compartment.
* `connection_strings` - Details of Globally distributed autonomous database connection String.
	* `all_connection_strings` - Collection of connection strings.
* `database_version` - Oracle Database version for the shards and catalog used in Globally distributed autonomous database.
* `db_backup_config` - Backup options for the Distributed Autonomous Database. 
	* `backup_destination_details` - Backup destination details.
		* `dbrs_policy_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DBRS policy used for backup.
		* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup destination.
		* `internet_proxy` - Proxy URL to connect to object store.
		* `is_remote` - Indicates whether the backup destination is cross-region or local region.
		* `remote_region` - The name of the remote region where the remote automatic incremental backups will be stored. For information about valid region names, see [Regions and Availability Domains](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/regions.htm). 
		* `type` - Type of the database backup destination.
		* `vpc_password` - For a RECOVERY_APPLIANCE backup destination, the password for the VPC user that is used to access the Recovery Appliance.
		* `vpc_user` - For a RECOVERY_APPLIANCE backup destination, the Virtual Private Catalog (VPC) user that is used to access the Recovery Appliance.
	* `recovery_window_in_days` - Number of days between the current and the earliest point of recoverability covered by automatic backups. This value applies to automatic backups. When the value is updated, it is applied to all existing automatic backups. If the number of specified days is 0 then there will be no backups. 
* `db_deployment_type` - The distributed autonomous database deployment type. 
* `db_workload` - Possible workload types. Currently only OLTP workload type is supported.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - The display name of the Globally distributed autonomous database.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `gsm_details` - Collection of catalogs associated with the Globally distributed autonomous database.
	* `compute_count` - The compute count for the Global service manager instance.
	* `data_storage_size_in_gbs` - The data disk group size to be allocated in GBs for the Global service manager instance.
	* `gsm_image_details` - The Global service manager image details.
		* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Global service manager software image.
		* `version_number` - The version number associated with the image identified by id.
	* `metadata` - Additional metadata related to Globally distributed autonomous database resources.
		* `map` - The map containing key-value pair of additional metadata.
	* `name` - Name of the Global service manager instance
	* `status` - Status of the gsm for the Globally distributed autonomous database.
	* `supporting_resource_id` - the identifier of the underlying supporting resource.
	* `time_created` - The time the Global service manager instance was created. An RFC3339 formatted datetime string
	* `time_ssl_certificate_expires` - The time the ssl certificate associated with Global service manager expires. An RFC3339 formatted datetime string
	* `time_updated` - The time the Global service manager instance was last updated. An RFC3339 formatted datetime string
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Globally distributed autonomous database.
* `latest_gsm_image` - The Global service manager image details.
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Global service manager software image.
	* `version_number` - The version number associated with the image identified by id.
* `lifecycle_details` - The lifecycleDetails for the Globally distributed autonomous database.
* `listener_port` - The listener port number for the Globally distributed autonomous database.
* `listener_port_tls` - The TLS listener port number for Globally distributed autonomous database.
* `metadata` - Additional metadata related to Globally distributed autonomous database resources.
	* `map` - The map containing key-value pair of additional metadata.
* `ncharacter_set` - The national character set for the database.
* `ons_port_local` - Ons local port number for Globally distributed autonomous database.
* `ons_port_remote` - Ons remote port number for Globally distributed autonomous database.
* `prefix` - Unique name prefix for the Globally distributed autonomous databases. Only alpha-numeric values are allowed. First character has to be a letter followed by any combination of letter and number. 
* `private_endpoint_ids` - The collection of [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the private endpoint associated with Globally distributed autonomous database.
* `replication_factor` - The Replication factor for RAFT replication based Globally distributed autonomous database. Currently supported values are 3, 5 and 7. 
* `replication_method` - The Replication method for Globally distributed autonomous database. Use RAFT for Raft replication, and DG for DataGuard. If replicationMethod is not provided, it defaults to DG. 
* `replication_unit` - The replication unit count for RAFT based distributed autonomous database. For RAFT replication based Globally distributed autonomous database, the value should be at least twice the number of shards. 
* `shard_details` - Collection of shards associated with the Globally distributed autonomous database.
	* `cloud_autonomous_vm_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloudAutonomousVmCluster.
	* `compute_count` - The compute count for the shard database. It has to be in multiples of 2.
	* `container_database_id` - the identifier of the container database for underlying supporting resource.
	* `data_storage_size_in_gbs` - The data disk group size to be allocated in GBs for the shard database.
	* `is_auto_scaling_enabled` - Determines the auto-scaling mode for the shard database.
	* `kms_key_id` - The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
	* `kms_key_version_id` - The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. 
	* `metadata` - Additional metadata related to Globally distributed autonomous database resources.
		* `map` - The map containing key-value pair of additional metadata.
	* `name` - Name of the shard.
	* `peer_cloud_autonomous_vm_cluster_ids` - This field is deprecated. For shard peer details please refer peerDetails attribute.
	* `peer_details` - Peer details for the shard with dedicated infrastructure.
		* `cloud_autonomous_vm_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloudAutonomousVmCluster.
		* `container_database_id` - the identifier of the container database for underlying supporting resource.
		* `fast_start_fail_over_lag_limit_in_seconds` - The lag time for my preference based on data loss tolerance in seconds.
		* `is_automatic_failover_enabled` - Indicates whether Automatic Failover is enabled for Autonomous Container Database Dataguard Association 
		* `metadata` - Additional metadata related to Globally distributed autonomous database resources.
			* `map` - The map containing key-value pair of additional metadata.
		* `protection_mode` - The protectionMode for the shard peer.
		* `shard_group` - The name of the shardGroup for the peer.
		* `standby_maintenance_buffer_in_days` - The scheduling detail for the quarterly maintenance window of the standby Autonomous Container Database. This value represents the number of days before schedlued maintenance of the primary database. 
		* `status` - Status of shard with dedicated infrastructure for the Globally distributed autonomous database.
		* `supporting_resource_id` - the identifier of the underlying supporting resource.
		* `time_created` - The time the shard peer was created. An RFC3339 formatted datetime string
		* `time_updated` - The time the shard peer was last updated. An RFC3339 formatted datetime string
	* `shard_group` - The name of the shardGroup for the shard.
	* `shard_space` - The shard space name for the Globally distributed autonomous database. Shard space for existing shard cannot be changed, once shard is created. Shard space name shall be used while creation of new shards. 
	* `source` - The source of Globally distributed autonomous database type: Use ADB_D for the Globally distributed autonomous database with autonomous dedicated cloudautonomousvmclusters. 
	* `status` - Status of shard with dedicated infrastructure for the Globally distributed autonomous database.
	* `supporting_resource_id` - the identifier of the underlying supporting resource.
	* `time_created` - The time the shard was created. An RFC3339 formatted datetime string
	* `time_updated` - The time the shard was last updated. An RFC3339 formatted datetime string
	* `vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `kmsKeyId` are required for Customer Managed Keys.
* `sharding_method` - Sharding Methods for the Globally distributed autonomous database.
* `state` - Lifecycle states for the Globally distributed autonomous database.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the Globally distributed autonomous database was created. An RFC3339 formatted datetime string
* `time_updated` - The time the Globally distributed autonomous database was last updated. An RFC3339 formatted datetime string

