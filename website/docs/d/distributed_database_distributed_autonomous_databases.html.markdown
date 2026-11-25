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
	compartment_id_in_subtree = var.distributed_autonomous_database_compartment_id_in_subtree
	db_deployment_type = var.distributed_autonomous_database_db_deployment_type
	display_name = var.distributed_autonomous_database_display_name
	is_auto_resource_management_enabled = var.distributed_autonomous_database_is_auto_resource_management_enabled
	metadata {
	}
	private_endpoint_id = oci_distributed_database_distributed_database_private_endpoint.test_distributed_database_private_endpoint.id
	state = var.distributed_autonomous_database_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `compartment_id_in_subtree` - (Optional) When set to true, returns resources in the specified compartment and in all of its child compartments recursively. Defaults to false.
* `db_deployment_type` - (Optional) A filter to return only resources whose dbDeploymentType matches the specified value.
* `display_name` - (Optional) A filter to return only Globally distributed autonomous databases that match the entire name given. The match is not case sensitive.
* `is_auto_resource_management_enabled` - (Optional) A filter to return only resources where autoResourceManagement is enabled when set to `true`, or disabled when set to `false`. If omitted, no autoResourceManagement-based filtering is applied.
* `metadata` - (Optional) Comma-separated argument names for which metadata should be retrieved. 
* `private_endpoint_id` - (Optional) A filter to return only resources that are associated with the given privateEndpointId.
* `state` - (Optional) A filter to return only resources whose lifecycleState matches the specified value.


## Attributes Reference

The following attributes are exported:

* `distributed_autonomous_database_collection` - The list of distributed_autonomous_database_collection.

### DistributedAutonomousDatabase Reference

The following attributes are exported:

* `auto_resource_management_config` - Automatic resource management configuration details for the Globally distributed database. 
	* `action_type` - The action that will be taken when autoResourceManagement is triggered.
	* `cool_off_period_in_minutes` - Time period to wait for the database to get stable after autoResourceManagement event.
	* `is_enabled` - Flag indicating if autoResourceManagement is enabled or disabled.
	* `max_move_ru_attempts` - Maximum number of move replication unit attempts allowed within the configured period per database.
	* `max_move_ru_unit` - Time unit applicable to maxMoveRuAttempts.
	* `mode` - Mode of autoResourceManagement execution.
	* `notification_topic_ids` - The list of notification topic ids [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) to receive autoResourceManagement events.
	* `stream_ids` - The list of stream ids [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) that receive shard-related telemetry.
* `catalog_details` - Catalog details associated with the distributed autonomous database.
	* `data_guard_replicas` - The details of data guard replicas for the catalog.
		* `database` - Details of the catalog database associated with the Globally distributed autonomous database.
			* `cloud_autonomous_vm_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud Autonomous VM Cluster.
			* `compute_count` - The compute count for the catalog database. It has to be in multiples of 2.
			* `container_database_id` - the identifier of the container database for underlying supporting resource.
			* `data_storage_size_in_gbs` - The data disk group size to be allocated in GBs for the catalog database.
			* `fast_start_fail_over_lag_limit_in_seconds` - The lag time preference based on data loss tolerance in seconds.
			* `is_auto_scaling_enabled` - Indicates if vertical auto scaling is enabled for the Autonomous AI Database CPU core count. The default value is `FALSE`. 
			* `kms_key_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
			* `kms_key_version_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. 
			* `metadata` - Additional metadata related to Globally distributed autonomous database resources.
				* `map` - The map containing key-value pair of additional metadata.
			* `name` - The name of the catalog.
			* `okv_end_point_group_name` - The OKV endpoint name.
			* `okv_key_store_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key store used to create the catalog.
			* `protection_mode` - The protection mode for the catalog peer.
			* `source` - The source of the Globally distributed autonomous database type. Use ADBD_EXISTING_CLUSTER for an autonomous database on Dedicated Infrastructure that uses a pre-existing cluster. 
			* `standby_maintenance_buffer_in_days` - The scheduling detail for the quarterly maintenance window of the standby Autonomous Container Database. This value represents the number of days before scheduled maintenance of the primary database. 
			* `status` - Status of the distributed autonomous database catalog.
			* `supporting_resource_id` - the identifier of the underlying supporting resource.
			* `time_created` - The time the catalog was created. An RFC3339 formatted datetime string.
			* `time_updated` - The time the catalog was last updated. An RFC3339 formatted datetime string.
			* `vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `kmsKeyId` are required for Customer Managed Keys.
		* `name` - The name of the catalog replica.  It must start with a letter and contain only letters, digits, and underscores.  Maximum length is 40 characters. 
	* `original_replica` - Details of a catalog replica.
		* `database` - Details of the catalog database associated with the Globally distributed autonomous database.
			* `cloud_autonomous_vm_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud Autonomous VM Cluster.
			* `compute_count` - The compute count for the catalog database. It has to be in multiples of 2.
			* `container_database_id` - the identifier of the container database for underlying supporting resource.
			* `data_storage_size_in_gbs` - The data disk group size to be allocated in GBs for the catalog database.
			* `fast_start_fail_over_lag_limit_in_seconds` - The lag time preference based on data loss tolerance in seconds.
			* `is_auto_scaling_enabled` - Indicates if vertical auto scaling is enabled for the Autonomous AI Database CPU core count. The default value is `FALSE`. 
			* `kms_key_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
			* `kms_key_version_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. 
			* `metadata` - Additional metadata related to Globally distributed autonomous database resources.
				* `map` - The map containing key-value pair of additional metadata.
			* `name` - The name of the catalog.
			* `okv_end_point_group_name` - The OKV endpoint name.
			* `okv_key_store_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key store used to create the catalog.
			* `protection_mode` - The protection mode for the catalog peer.
			* `source` - The source of the Globally distributed autonomous database type. Use ADBD_EXISTING_CLUSTER for an autonomous database on Dedicated Infrastructure that uses a pre-existing cluster. 
			* `standby_maintenance_buffer_in_days` - The scheduling detail for the quarterly maintenance window of the standby Autonomous Container Database. This value represents the number of days before scheduled maintenance of the primary database. 
			* `status` - Status of the distributed autonomous database catalog.
			* `supporting_resource_id` - the identifier of the underlying supporting resource.
			* `time_created` - The time the catalog was created. An RFC3339 formatted datetime string.
			* `time_updated` - The time the catalog was last updated. An RFC3339 formatted datetime string.
			* `vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `kmsKeyId` are required for Customer Managed Keys.
		* `name` - The name of the catalog replica.  It must start with a letter and contain only letters, digits, and underscores.  Maximum length is 40 characters. 
* `character_set` - The character set for the database.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Globally distributed autonomous database compartment.
* `composite_data_guard_shard_spaces` - Collection of composite data guard shard spaces.
	* `chunk_count` - Count of chunks associated with the shard space.
	* `data_guard_replicas` - The details of data guard replica associated with the shard.
		* `databases` - The details of composite data guard shard space replica databases.
			* `cloud_autonomous_vm_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud Autonomous Exadata VM Cluster.
			* `compute_count` - The compute count for the shard database. It has to be in multiples of 2.
			* `container_database_id` - the identifier of the container database for underlying supporting resource.
			* `data_storage_size_in_gbs` - The data disk group size to be allocated in GBs for the shard database.
			* `fast_start_fail_over_lag_limit_in_seconds` - The lag time preference based on data loss tolerance in seconds.
			* `is_auto_scaling_enabled` - Indicates if vertical auto scaling is enabled for the Autonomous AI Database CPU core count. The default value is `FALSE`. 
			* `kms_key_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
			* `kms_key_version_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. 
			* `metadata` - Additional metadata related to Globally distributed autonomous database resources.
				* `map` - The map containing key-value pair of additional metadata.
			* `name` - Name of the shard.
			* `okv_end_point_group_name` - The OKV endpoint name.
			* `okv_key_store_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key store used to create the shard.
			* `protection_mode` - The protection mode for the shard peer.
			* `source` - The source of the Globally distributed autonomous database type. Use ADBD_EXISTING_CLUSTER for an autonomous database on Dedicated Infrastructure that uses a pre-existing cluster. 
			* `standby_maintenance_buffer_in_days` - The scheduling detail for the quarterly maintenance window of the standby Autonomous Container Database. 
			* `status` - Status of the distributed autonomous database shard.
			* `supporting_resource_id` - the identifier of the underlying supporting resource.
			* `time_created` - The time the shard was created. An RFC3339 formatted datetime string.
			* `time_updated` - The time the shard was last updated. An RFC3339 formatted datetime string.
			* `vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `kmsKeyId` are required for Customer Managed Keys.
		* `name` - The name of shard space replica.  It must start with a letter and contain only letters, digits, and underscores.  Maximum length is 40 characters. 
	* `name` - The name of shard space.  It must start with a letter and contain only letters, digits, and underscores.  Maximum length is 40 characters. 
	* `original_replica` - Globally distributed database composite data guard replica details.
		* `databases` - The details of composite data guard shard space replica databases.
			* `cloud_autonomous_vm_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud Autonomous Exadata VM Cluster.
			* `compute_count` - The compute count for the shard database. It has to be in multiples of 2.
			* `container_database_id` - the identifier of the container database for underlying supporting resource.
			* `data_storage_size_in_gbs` - The data disk group size to be allocated in GBs for the shard database.
			* `fast_start_fail_over_lag_limit_in_seconds` - The lag time preference based on data loss tolerance in seconds.
			* `is_auto_scaling_enabled` - Indicates if vertical auto scaling is enabled for the Autonomous AI Database CPU core count. The default value is `FALSE`. 
			* `kms_key_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
			* `kms_key_version_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. 
			* `metadata` - Additional metadata related to Globally distributed autonomous database resources.
				* `map` - The map containing key-value pair of additional metadata.
			* `name` - Name of the shard.
			* `okv_end_point_group_name` - The OKV endpoint name.
			* `okv_key_store_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key store used to create the shard.
			* `protection_mode` - The protection mode for the shard peer.
			* `source` - The source of the Globally distributed autonomous database type. Use ADBD_EXISTING_CLUSTER for an autonomous database on Dedicated Infrastructure that uses a pre-existing cluster. 
			* `standby_maintenance_buffer_in_days` - The scheduling detail for the quarterly maintenance window of the standby Autonomous Container Database. 
			* `status` - Status of the distributed autonomous database shard.
			* `supporting_resource_id` - the identifier of the underlying supporting resource.
			* `time_created` - The time the shard was created. An RFC3339 formatted datetime string.
			* `time_updated` - The time the shard was last updated. An RFC3339 formatted datetime string.
			* `vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `kmsKeyId` are required for Customer Managed Keys.
		* `name` - The name of shard space replica.  It must start with a letter and contain only letters, digits, and underscores.  Maximum length is 40 characters. 
* `composite_raft_shard_spaces` - Collection of composite raft shards.
	* `chunk_count` - Count of chunks associated with the shard space.
	* `name` - The name of shard space.  It must start with a letter and contain only letters, digits, and underscores.  Maximum length is 40 characters. 
	* `raft_clusters` - The details of raft clusters associated with the shard space.
		* `databases` - The details of databases associated with the composite shard space raft cluster.
			* `cloud_autonomous_vm_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud Autonomous Exadata VM Cluster.
			* `compute_count` - The compute count for the shard database. It has to be in multiples of 2.
			* `container_database_id` - the identifier of the container database for underlying supporting resource.
			* `data_storage_size_in_gbs` - The data disk group size to be allocated in GBs for the shard database.
			* `fast_start_fail_over_lag_limit_in_seconds` - The lag time preference based on data loss tolerance in seconds.
			* `is_auto_scaling_enabled` - Indicates if vertical auto scaling is enabled for the Autonomous AI Database CPU core count. The default value is `FALSE`. 
			* `kms_key_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
			* `kms_key_version_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. 
			* `metadata` - Additional metadata related to Globally distributed autonomous database resources.
				* `map` - The map containing key-value pair of additional metadata.
			* `name` - Name of the shard.
			* `okv_end_point_group_name` - The OKV endpoint name.
			* `okv_key_store_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key store used to create the shard.
			* `protection_mode` - The protection mode for the shard peer.
			* `source` - The source of the Globally distributed autonomous database type. Use ADBD_EXISTING_CLUSTER for an autonomous database on Dedicated Infrastructure that uses a pre-existing cluster. 
			* `standby_maintenance_buffer_in_days` - The scheduling detail for the quarterly maintenance window of the standby Autonomous Container Database. 
			* `status` - Status of the distributed autonomous database shard.
			* `supporting_resource_id` - the identifier of the underlying supporting resource.
			* `time_created` - The time the shard was created. An RFC3339 formatted datetime string.
			* `time_updated` - The time the shard was last updated. An RFC3339 formatted datetime string.
			* `vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `kmsKeyId` are required for Customer Managed Keys.
		* `name` - The name of raft cluster.  It must start with a letter and contain only letters, digits, and underscores.  Maximum length is 40 characters. 
		* `replication_factor` - Replication factor associated with the raft cluster.
		* `ru_mode` - Replication factor associated with the raft cluster.
		* `witness_count` - Witnesses RU count associated with the raft cluster.
	* `replication_unit_count` - Number of replication units in the shard space.
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
	* `recovery_window_in_days` - Number of days between the current and the earliest point of recoverability covered by automatic backups. This value applies to automatic backups. When the value is updated, it is applied to all existing automatic backups. If the number of specified days is 0 then there will be no backups. 
* `db_deployment_type` - The distributed autonomous database deployment type. 
* `db_workload_type` - Possible workload types. Currently only OLTP workload type is supported.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - The display name of the Globally distributed autonomous database.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `gds_control_node_details` - Global Database Services Control(GDS CTL) instances associated with the distributed autonomous database.
	* `availability_domain` - The name of the availability domain where the Global Database Services Control instance is located in.
	* `gds_control_node_image_details` - The Global Database Services Control(GDS CTL) image details.
		* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Global Database Services Control(GDS CTL) software image.
		* `version` - The version associated with the image identified by id.
	* `metadata` - Additional metadata related to Globally distributed autonomous database resources.
		* `map` - The map containing key-value pair of additional metadata.
	* `name` - Name of the Global Database Services Control instance
	* `private_ip` - IP address of the Global Database Services Control instance.
	* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet associated with the Global Database Services Control instance. 
	* `supporting_resource_id` - The identifier of the underlying supporting resource.
	* `time_created` - The time the Global Database Services Control instance was created. An RFC3339 formatted datetime string.
	* `time_updated` - The time the Global Database Services Control instance was last updated. An RFC3339 formatted datetime string.
* `gsm_details` - Global Service Manager (GSM) instances associated with the distributed autonomous database.
	* `compute_count` - The compute count for the Global service manager instance.
	* `data_storage_size_in_gbs` - The data disk group size to be allocated in GBs for the Global service manager instance.
	* `gsm_image_details` - The Global service manager image details.
		* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Global service manager software image.
		* `version` - The version associated with the image identified by id.
	* `metadata` - Additional metadata related to Globally distributed autonomous database resources.
		* `map` - The map containing key-value pair of additional metadata.
	* `name` - Name of the Global service manager instance.
	* `status` - Status of the gsm for the Globally distributed autonomous database.
	* `supporting_resource_id` - the identifier of the underlying supporting resource.
	* `time_created` - The time the Global service manager instance was created. An RFC3339 formatted datetime string.
	* `time_ssl_certificate_expires` - The time the ssl certificate associated with Global service manager expires. An RFC3339 formatted datetime string.
	* `time_updated` - The time the Global service manager instance was last updated. An RFC3339 formatted datetime string.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Globally distributed autonomous database.
* `latest_gsm_image` - The Global service manager image details.
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Global service manager software image.
	* `version` - The version associated with the image identified by id.
* `lifecycle_details` - The lifecycleDetails for the Globally distributed autonomous database.
* `listener_port` - The listener port number for the Globally distributed autonomous database.
* `listener_port_tls` - The TLS listener port number for Globally distributed autonomous database.
* `metadata` - Additional metadata related to Globally distributed autonomous database resources.
	* `map` - The map containing key-value pair of additional metadata.
* `ncharacter_set` - The national character set for the database.
* `notification_topic_ids` - The collection of [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the notification topics associated with the globally distributed autonomous database.
* `ons_port_local` - Ons local port number for Globally distributed autonomous database.
* `ons_port_remote` - Ons remote port number for Globally distributed autonomous database.
* `prefix` - Unique name prefix for the Globally distributed autonomous databases. Only alpha-numeric values are allowed. First character has to be a letter followed by any combination of letter and number. 
* `private_endpoint_ids` - The collection of [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the private endpoint associated with Globally distributed autonomous database.
* `replication_method` - The Replication method for Globally distributed Autonomous database. Use RAFT for Raft based replication. With RAFT replication, shards cannot have peers details set on them. In case shards need to have peers, please do not set RAFT replicationMethod. For all non RAFT replication cases (with or without peers), please set replicationMethod as DG or do not set any value for replicationMethod. 
* `sharding_method` - Sharding methods for the Globally distributed autonomous database.
* `state` - Lifecycle states for the Globally distributed autonomous database.
* `system_chunk_count` - Count of chunks associated with system raft clusters or system data guard databases.
* `system_data_guard_databases` - System data guard shard database details.
	* `data_guard_replicas` - The system data guard replica details.
		* `databases` - Details of the databases associated with the system data guard replica.
			* `cloud_autonomous_vm_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud Autonomous Exadata VM Cluster.
			* `compute_count` - The compute count for the shard database. It has to be in multiples of 2.
			* `container_database_id` - the identifier of the container database for underlying supporting resource.
			* `data_storage_size_in_gbs` - The data disk group size to be allocated in GBs for the shard database.
			* `fast_start_fail_over_lag_limit_in_seconds` - The lag time preference based on data loss tolerance in seconds.
			* `is_auto_scaling_enabled` - Indicates if vertical auto scaling is enabled for the Autonomous AI Database CPU core count. The default value is `FALSE`. 
			* `kms_key_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
			* `kms_key_version_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. 
			* `metadata` - Additional metadata related to Globally distributed autonomous database resources.
				* `map` - The map containing key-value pair of additional metadata.
			* `name` - Name of the shard.
			* `okv_end_point_group_name` - The OKV endpoint name.
			* `okv_key_store_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key store used to create the shard.
			* `protection_mode` - The protection mode for the shard peer.
			* `source` - The source of the Globally distributed autonomous database type. Use ADBD_EXISTING_CLUSTER for an autonomous database on Dedicated Infrastructure that uses a pre-existing cluster. 
			* `standby_maintenance_buffer_in_days` - The scheduling detail for the quarterly maintenance window of the standby Autonomous Container Database. 
			* `status` - Status of the distributed autonomous database shard.
			* `supporting_resource_id` - the identifier of the underlying supporting resource.
			* `time_created` - The time the shard was created. An RFC3339 formatted datetime string.
			* `time_updated` - The time the shard was last updated. An RFC3339 formatted datetime string.
			* `vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `kmsKeyId` are required for Customer Managed Keys.
		* `name` - The name of data guard replica.  It must start with a letter and contain only letters, digits, and underscores.  Maximum length is 40 characters. 
	* `original_replica` - System data guard replica details.
		* `databases` - Details of the databases associated with the system data guard replica.
			* `cloud_autonomous_vm_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud Autonomous Exadata VM Cluster.
			* `compute_count` - The compute count for the shard database. It has to be in multiples of 2.
			* `container_database_id` - the identifier of the container database for underlying supporting resource.
			* `data_storage_size_in_gbs` - The data disk group size to be allocated in GBs for the shard database.
			* `fast_start_fail_over_lag_limit_in_seconds` - The lag time preference based on data loss tolerance in seconds.
			* `is_auto_scaling_enabled` - Indicates if vertical auto scaling is enabled for the Autonomous AI Database CPU core count. The default value is `FALSE`. 
			* `kms_key_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
			* `kms_key_version_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. 
			* `metadata` - Additional metadata related to Globally distributed autonomous database resources.
				* `map` - The map containing key-value pair of additional metadata.
			* `name` - Name of the shard.
			* `okv_end_point_group_name` - The OKV endpoint name.
			* `okv_key_store_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key store used to create the shard.
			* `protection_mode` - The protection mode for the shard peer.
			* `source` - The source of the Globally distributed autonomous database type. Use ADBD_EXISTING_CLUSTER for an autonomous database on Dedicated Infrastructure that uses a pre-existing cluster. 
			* `standby_maintenance_buffer_in_days` - The scheduling detail for the quarterly maintenance window of the standby Autonomous Container Database. 
			* `status` - Status of the distributed autonomous database shard.
			* `supporting_resource_id` - the identifier of the underlying supporting resource.
			* `time_created` - The time the shard was created. An RFC3339 formatted datetime string.
			* `time_updated` - The time the shard was last updated. An RFC3339 formatted datetime string.
			* `vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `kmsKeyId` are required for Customer Managed Keys.
		* `name` - The name of data guard replica.  It must start with a letter and contain only letters, digits, and underscores.  Maximum length is 40 characters. 
* `system_raft_clusters` - Collection of system raft clusters.
	* `databases` - Details of the databases associated with the system raft cluster.
		* `cloud_autonomous_vm_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud Autonomous Exadata VM Cluster.
		* `compute_count` - The compute count for the shard database. It has to be in multiples of 2.
		* `container_database_id` - the identifier of the container database for underlying supporting resource.
		* `data_storage_size_in_gbs` - The data disk group size to be allocated in GBs for the shard database.
		* `fast_start_fail_over_lag_limit_in_seconds` - The lag time preference based on data loss tolerance in seconds.
		* `is_auto_scaling_enabled` - Indicates if vertical auto scaling is enabled for the Autonomous AI Database CPU core count. The default value is `FALSE`. 
		* `kms_key_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
		* `kms_key_version_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. 
		* `metadata` - Additional metadata related to Globally distributed autonomous database resources.
			* `map` - The map containing key-value pair of additional metadata.
		* `name` - Name of the shard.
		* `okv_end_point_group_name` - The OKV endpoint name.
		* `okv_key_store_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key store used to create the shard.
		* `protection_mode` - The protection mode for the shard peer.
		* `source` - The source of the Globally distributed autonomous database type. Use ADBD_EXISTING_CLUSTER for an autonomous database on Dedicated Infrastructure that uses a pre-existing cluster. 
		* `standby_maintenance_buffer_in_days` - The scheduling detail for the quarterly maintenance window of the standby Autonomous Container Database. 
		* `status` - Status of the distributed autonomous database shard.
		* `supporting_resource_id` - the identifier of the underlying supporting resource.
		* `time_created` - The time the shard was created. An RFC3339 formatted datetime string.
		* `time_updated` - The time the shard was last updated. An RFC3339 formatted datetime string.
		* `vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `kmsKeyId` are required for Customer Managed Keys.
	* `name` - The name of raft cluster.  It must start with a letter and contain only letters, digits, and underscores.  Maximum length is 40 characters. 
	* `replication_factor` - Replication factor associated with the raft cluster.
	* `ru_mode` - Replication factor associated with the raft cluster.
	* `witness_count` - Witnesses RU count associated with the raft cluster.
* `system_raft_replication_unit_count` - Number of replication units associated with system raft clusters.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the Globally distributed autonomous database was created. An RFC3339 formatted datetime string.
* `time_updated` - The time the Globally distributed autonomous database was last updated. An RFC3339 formatted datetime string.
* `user_shard_spaces` - Collection of user defined shard spaces.
	* `data_guard_replicas` - Details of data guard replicas for the user shard space.
		* `database` - Globally distributed autonomous database shard.
			* `cloud_autonomous_vm_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud Autonomous Exadata VM Cluster.
			* `compute_count` - The compute count for the shard database. It has to be in multiples of 2.
			* `container_database_id` - the identifier of the container database for underlying supporting resource.
			* `data_storage_size_in_gbs` - The data disk group size to be allocated in GBs for the shard database.
			* `fast_start_fail_over_lag_limit_in_seconds` - The lag time preference based on data loss tolerance in seconds.
			* `is_auto_scaling_enabled` - Indicates if vertical auto scaling is enabled for the Autonomous AI Database CPU core count. The default value is `FALSE`. 
			* `kms_key_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
			* `kms_key_version_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. 
			* `metadata` - Additional metadata related to Globally distributed autonomous database resources.
				* `map` - The map containing key-value pair of additional metadata.
			* `name` - Name of the shard.
			* `okv_end_point_group_name` - The OKV endpoint name.
			* `okv_key_store_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key store used to create the shard.
			* `protection_mode` - The protection mode for the shard peer.
			* `source` - The source of the Globally distributed autonomous database type. Use ADBD_EXISTING_CLUSTER for an autonomous database on Dedicated Infrastructure that uses a pre-existing cluster. 
			* `standby_maintenance_buffer_in_days` - The scheduling detail for the quarterly maintenance window of the standby Autonomous Container Database. 
			* `status` - Status of the distributed autonomous database shard.
			* `supporting_resource_id` - the identifier of the underlying supporting resource.
			* `time_created` - The time the shard was created. An RFC3339 formatted datetime string.
			* `time_updated` - The time the shard was last updated. An RFC3339 formatted datetime string.
			* `vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `kmsKeyId` are required for Customer Managed Keys.
		* `name` - The name of the user shard space replica.  It must start with a letter and contain only letters, digits, and underscores.  Maximum length is 40 characters. 
	* `name` - The name of the user shard space.  It must start with a letter and contain only letters, digits, and underscores.  Maximum length is 40 characters. 
	* `original_replica` - User shard space replica details.
		* `database` - Globally distributed autonomous database shard.
			* `cloud_autonomous_vm_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud Autonomous Exadata VM Cluster.
			* `compute_count` - The compute count for the shard database. It has to be in multiples of 2.
			* `container_database_id` - the identifier of the container database for underlying supporting resource.
			* `data_storage_size_in_gbs` - The data disk group size to be allocated in GBs for the shard database.
			* `fast_start_fail_over_lag_limit_in_seconds` - The lag time preference based on data loss tolerance in seconds.
			* `is_auto_scaling_enabled` - Indicates if vertical auto scaling is enabled for the Autonomous AI Database CPU core count. The default value is `FALSE`. 
			* `kms_key_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
			* `kms_key_version_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. 
			* `metadata` - Additional metadata related to Globally distributed autonomous database resources.
				* `map` - The map containing key-value pair of additional metadata.
			* `name` - Name of the shard.
			* `okv_end_point_group_name` - The OKV endpoint name.
			* `okv_key_store_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key store used to create the shard.
			* `protection_mode` - The protection mode for the shard peer.
			* `source` - The source of the Globally distributed autonomous database type. Use ADBD_EXISTING_CLUSTER for an autonomous database on Dedicated Infrastructure that uses a pre-existing cluster. 
			* `standby_maintenance_buffer_in_days` - The scheduling detail for the quarterly maintenance window of the standby Autonomous Container Database. 
			* `status` - Status of the distributed autonomous database shard.
			* `supporting_resource_id` - the identifier of the underlying supporting resource.
			* `time_created` - The time the shard was created. An RFC3339 formatted datetime string.
			* `time_updated` - The time the shard was last updated. An RFC3339 formatted datetime string.
			* `vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `kmsKeyId` are required for Customer Managed Keys.
		* `name` - The name of the user shard space replica.  It must start with a letter and contain only letters, digits, and underscores.  Maximum length is 40 characters. 
* `vcn_nsg_ids` - The list of network security group (NSG) details associated with the distributed autonomous database.
	* `nsg_ids` - The list of [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) associated with resources in this VCN. For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). 
	* `vcn_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the customer Virtual Cloud Network.

