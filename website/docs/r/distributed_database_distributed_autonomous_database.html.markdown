---
subcategory: "Distributed Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_distributed_database_distributed_autonomous_database"
sidebar_current: "docs-oci-resource-distributed_database-distributed_autonomous_database"
description: |-
  Provides the Distributed Autonomous Database resource in Oracle Cloud Infrastructure Distributed Database service
---

# oci_distributed_database_distributed_autonomous_database
This resource provides the Distributed Autonomous Database resource in Oracle Cloud Infrastructure Distributed Database service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/globally-distributed-database/latest/DistributedAutonomousDatabase

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/distributed_database

Creates a Globally distributed autonomous database.

  Patch operation to add, remove or update shards to the Globally distributed autonomous database topology. In single patch
operation, multiple shards can be either added, or removed or updated. Combination of inserts, update
and remove in single operation is not allowed.


## Example Usage

```hcl
resource "oci_distributed_database_distributed_autonomous_database" "test_distributed_autonomous_database" {
	#Required
	catalog_details {
		#Required
		admin_password = var.distributed_autonomous_database_catalog_details_admin_password
		cloud_autonomous_vm_cluster_id = oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster.id
		compute_count = var.distributed_autonomous_database_catalog_details_compute_count
		data_storage_size_in_gbs = var.distributed_autonomous_database_catalog_details_data_storage_size_in_gbs
		is_auto_scaling_enabled = var.distributed_autonomous_database_catalog_details_is_auto_scaling_enabled
		source = var.distributed_autonomous_database_catalog_details_source

		#Optional
		kms_key_id = oci_kms_key.test_key.id
		kms_key_version_id = oci_kms_key_version.test_key_version.id
		peer_cloud_autonomous_vm_cluster_ids = var.distributed_autonomous_database_catalog_details_peer_cloud_autonomous_vm_cluster_ids
		peer_details {
			#Required
			cloud_autonomous_vm_cluster_id = oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster.id

			#Optional
			fast_start_fail_over_lag_limit_in_seconds = var.distributed_autonomous_database_catalog_details_peer_details_fast_start_fail_over_lag_limit_in_seconds
			is_automatic_failover_enabled = var.distributed_autonomous_database_catalog_details_peer_details_is_automatic_failover_enabled
			protection_mode = var.distributed_autonomous_database_catalog_details_peer_details_protection_mode
			standby_maintenance_buffer_in_days = var.distributed_autonomous_database_catalog_details_peer_details_standby_maintenance_buffer_in_days
		}
		vault_id = oci_kms_vault.test_vault.id
	}
	character_set = var.distributed_autonomous_database_character_set
	compartment_id = var.compartment_id
	database_version = var.distributed_autonomous_database_database_version
	db_deployment_type = var.distributed_autonomous_database_db_deployment_type
	db_workload = var.distributed_autonomous_database_db_workload
	display_name = var.distributed_autonomous_database_display_name
	distributed_autonomous_database_id = var.distributed_autonomous_database_distributed_autonomous_database_id
	listener_port = var.distributed_autonomous_database_listener_port
	ncharacter_set = var.distributed_autonomous_database_ncharacter_set
	ons_port_local = var.distributed_autonomous_database_ons_port_local
	ons_port_remote = var.distributed_autonomous_database_ons_port_remote
	prefix = var.distributed_autonomous_database_prefix
	private_endpoint_ids = var.distributed_autonomous_database_private_endpoint_ids
	shard_details {
		#Required
		admin_password = var.distributed_autonomous_database_shard_details_admin_password
		cloud_autonomous_vm_cluster_id = oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster.id
		compute_count = var.distributed_autonomous_database_shard_details_compute_count
		data_storage_size_in_gbs = var.distributed_autonomous_database_shard_details_data_storage_size_in_gbs
		is_auto_scaling_enabled = var.distributed_autonomous_database_shard_details_is_auto_scaling_enabled
		source = var.distributed_autonomous_database_shard_details_source

		#Optional
		kms_key_id = oci_kms_key.test_key.id
		kms_key_version_id = oci_kms_key_version.test_key_version.id
		peer_cloud_autonomous_vm_cluster_ids = var.distributed_autonomous_database_shard_details_peer_cloud_autonomous_vm_cluster_ids
		peer_details {
			#Required
			cloud_autonomous_vm_cluster_id = oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster.id

			#Optional
			fast_start_fail_over_lag_limit_in_seconds = var.distributed_autonomous_database_shard_details_peer_details_fast_start_fail_over_lag_limit_in_seconds
			is_automatic_failover_enabled = var.distributed_autonomous_database_shard_details_peer_details_is_automatic_failover_enabled
			protection_mode = var.distributed_autonomous_database_shard_details_peer_details_protection_mode
			standby_maintenance_buffer_in_days = var.distributed_autonomous_database_shard_details_peer_details_standby_maintenance_buffer_in_days
		}
		shard_space = var.distributed_autonomous_database_shard_details_shard_space
		vault_id = oci_kms_vault.test_vault.id
	}
	sharding_method = var.distributed_autonomous_database_sharding_method

	#Optional
	chunks = var.distributed_autonomous_database_chunks
	db_backup_config {

		#Optional
		backup_destination_details {
			#Required
			type = var.distributed_autonomous_database_db_backup_config_backup_destination_details_type

			#Optional
			dbrs_policy_id = oci_identity_policy.test_policy.id
			id = var.distributed_autonomous_database_db_backup_config_backup_destination_details_id
			internet_proxy = var.distributed_autonomous_database_db_backup_config_backup_destination_details_internet_proxy
			is_remote = var.distributed_autonomous_database_db_backup_config_backup_destination_details_is_remote
			remote_region = var.distributed_autonomous_database_db_backup_config_backup_destination_details_remote_region
			vpc_password = var.distributed_autonomous_database_db_backup_config_backup_destination_details_vpc_password
			vpc_user = var.distributed_autonomous_database_db_backup_config_backup_destination_details_vpc_user
		}
		recovery_window_in_days = var.distributed_autonomous_database_db_backup_config_recovery_window_in_days
	}
	defined_tags = {"foo-namespace.bar-key"= "value"}
	freeform_tags = {"bar-key"= "value"}
	listener_port_tls = var.distributed_autonomous_database_listener_port_tls
	patch_operations {
		#Required
		operation = var.distributed_autonomous_database_patch_operations_operation
		selection = var.distributed_autonomous_database_patch_operations_selection

		#Optional
		value = var.distributed_autonomous_database_patch_operations_value
	}
	replication_factor = var.distributed_autonomous_database_replication_factor
	replication_method = var.distributed_autonomous_database_replication_method
	replication_unit = var.distributed_autonomous_database_replication_unit
}
```

## Argument Reference

The following arguments are supported:

* `catalog_details` - (Required) Collection of catalog for the Globally distributed autonomous database.
	* `admin_password` - (Required) Admin password for catalog database.
	* `cloud_autonomous_vm_cluster_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud Autonomous VM Cluster.
	* `compute_count` - (Required) The compute count for the catalog database. It has to be in multiples of 2.
	* `data_storage_size_in_gbs` - (Required) The data disk group size to be allocated in GBs for the catalog database.
	* `is_auto_scaling_enabled` - (Required) Determines the auto-scaling mode for the catalog database.
	* `kms_key_id` - (Optional) The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
	* `kms_key_version_id` - (Optional) The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. 
	* `peer_cloud_autonomous_vm_cluster_ids` - (Optional) This field is deprecated. This should not be used while creation of new distributed autonomous database. To set the peers on catalog of distributed autonomous database please use peerDetails. 
	* `peer_details` - (Optional) The details required for creation of the peer for the autonomous dedicated infrastructure based catalog.
		* `cloud_autonomous_vm_cluster_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud Autonomous VM Cluster for the peer catalog.
		* `fast_start_fail_over_lag_limit_in_seconds` - (Optional) The lag time preference based on data loss tolerance in seconds.
		* `is_automatic_failover_enabled` - (Optional) Indicates whether Automatic Failover is enabled for Autonomous Container Database Dataguard Association 
		* `protection_mode` - (Optional) The protectionMode for the catalog peer.
		* `standby_maintenance_buffer_in_days` - (Optional) The scheduling detail for the quarterly maintenance window of the standby Autonomous Container Database. This value represents the number of days before schedlued maintenance of the primary database. 
	* `source` - (Required) The source of Globally distributed autonomous database type: Use ADB_D for the Globally distributed autonomous database with autonomous dedicated cloudautonomousvmclusters. 
	* `vault_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `kmsKeyId` are required for Customer Managed Keys.
* `character_set` - (Required) The character set for the database.
* `chunks` - (Optional) Number of chunks in a shardspace. The value of chunks must be greater than 2 times the size of the largest shardgroup in any shardspace. Chunks is required to be provided for distributed autonomous databases being created with SYSTEM shardingMethod. For USER shardingMethod, chunks should not be set in create payload. 
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Globally distributed autonomous database compartment.
* `database_version` - (Required) Oracle Database version for the shards and catalog used in Globally distributed autonomous database.
* `db_backup_config` - (Optional) Backup options for the Distributed Autonomous Database. 
	* `backup_destination_details` - (Optional) Backup destination details.
		* `dbrs_policy_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DBRS policy used for backup.
		* `id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup destination.
		* `internet_proxy` - (Optional) Proxy URL to connect to object store.
		* `is_remote` - (Optional) Indicates whether the backup destination is cross-region or local region.
		* `remote_region` - (Optional) The name of the remote region where the remote automatic incremental backups will be stored. For information about valid region names, see [Regions and Availability Domains](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/regions.htm). 
		* `type` - (Required) Type of the database backup destination.
		* `vpc_password` - (Optional) For a RECOVERY_APPLIANCE backup destination, the password for the VPC user that is used to access the Recovery Appliance.
		* `vpc_user` - (Optional) For a RECOVERY_APPLIANCE backup destination, the Virtual Private Catalog (VPC) user that is used to access the Recovery Appliance.
	* `recovery_window_in_days` - (Optional) Number of days between the current and the earliest point of recoverability covered by automatic backups. This value applies to automatic backups. When the value is updated, it is applied to all existing automatic backups. If the number of specified days is 0 then there will be no backups. 
* `db_deployment_type` - (Required) The distributed autonomous database deployment type. 
* `db_workload` - (Required) Possible workload types. Currently only OLTP workload type is supported.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Required) (Updatable) The display name of the Globally distributed autonomous database.
* `distributed_autonomous_database_id` - (Required) 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `listener_port` - (Required) The listener port number for the Globally distributed autonomous database. The listener port number has to be unique for a customer tenancy across all distributed autonomous databases. Same port number should not be re-used for any other distributed autonomous database. 
* `listener_port_tls` - (Optional) The TLS listener port number for Globally distributed autonomous database. The TLS listener port number has to be unique for a customer tenancy across all distributed autonomous databases. Same port number should not be re-used for any other distributed autonomous database. The listenerPortTls is mandatory for dedicated infrastructure based distributed autonomous databases. 
* `ncharacter_set` - (Required) The national character set for the database.
* `ons_port_local` - (Required) Ons local port number for Globally distributed autonomous database. The onsPortLocal has to be unique for a customer tenancy across all distributed autonomous databases. Same port number should not be re-used for any other distributed autonomous database. 
* `ons_port_remote` - (Required) Ons remote port number for Globally distributed autonomous database. The onsPortRemote has to be unique for a customer tenancy across all distributed autonomous databases. Same port number should not be re-used for any other distributed autonomous database. 
* `patch_operations` - (Optional) (Updatable) 
	* `operation` - (Required) (Updatable) The operation can be one of these values: `INSERT`, `MERGE`, `REMOVE`
	* `selection` - (Required) (Updatable) 
	* `value` - (Required when operation=INSERT | MERGE) (Updatable) 
* `prefix` - (Required) Unique name prefix for the Globally distributed autonomous databases. Only alpha-numeric values are allowed. First character has to be a letter followed by any combination of letter and number. 
* `private_endpoint_ids` - (Required) The collection of [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the private endpoint associated with Globally distributed autonomous database.
* `replication_factor` - (Optional) The Replication factor for RAFT replication based Globally distributed autonomous database. Currently supported values are 3, 5 and 7. 
* `replication_method` - (Optional) The Replication method for Globally distributed autonomous database. Use RAFT for Raft based replication. With RAFT replication, shards cannot have peers details set on them. In case shards need to have peers, please do not set RAFT replicationMethod. For all non RAFT replication cases (with or without peers), please set replicationMethod as DG or do not set any value for replicationMethod. 
* `replication_unit` - (Optional) The replication unit count for RAFT based distributed autonomous database. For RAFT replication based Globally distributed autonomous database, the value should be at least twice the number of shards. 
* `shard_details` - (Required) Collection of shards for the Globally distributed autonomous database.
	* `admin_password` - (Required) Admin password for shard database.
	* `cloud_autonomous_vm_cluster_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud Autonomous Exadata VM Cluster.
	* `compute_count` - (Required) The compute count for the shard database. It has to be in multiples of 2.
	* `data_storage_size_in_gbs` - (Required) The data disk group size to be allocated in GBs for the shard database.
	* `is_auto_scaling_enabled` - (Required) Determines the auto-scaling mode for the shard database.
	* `kms_key_id` - (Optional) The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
	* `kms_key_version_id` - (Optional) The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. 
	* `peer_cloud_autonomous_vm_cluster_ids` - (Optional) This field is deprecated. This should not be used while creation of new distributed autonomous database. To set the peers on new shards of distributed autonomous database please use peerDetails. 
	* `peer_details` - (Optional) The details required for creation of the peer for the autonomous dedicated infrastructure based shard.
		* `cloud_autonomous_vm_cluster_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud Autonomous VM Cluster for the peer shard.
		* `fast_start_fail_over_lag_limit_in_seconds` - (Optional) The lag time preference based on data loss tolerance in seconds.
		* `is_automatic_failover_enabled` - (Optional) Indicates whether Automatic Failover is enabled for Autonomous Container Database Dataguard Association 
		* `protection_mode` - (Optional) The protectionMode for the shard peer.
		* `standby_maintenance_buffer_in_days` - (Optional) The scheduling detail for the quarterly maintenance window of the standby Autonomous Container Database. This value represents the number of days before schedlued maintenance of the primary database. 
	* `shard_space` - (Optional) The shard space name for the shard database. Shard space for existing shard cannot be changed, once shard is created. Shard space name shall be used while creation of new shards. For User defined sharding, every shard must have a unique shard space name. For system defined sharding, shard space name is not required. 
	* `source` - (Required) The source of Globally distributed autonomous database type: Use ADB_D for the Globally distributed autonomous database with autonomous dedicated cloudautonomousvmclusters. 
	* `vault_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `kmsKeyId` are required for Customer Managed Keys.
* `sharding_method` - (Required) Sharding Methods for the Globally distributed autonomous database.
* `state` - (Optional) (Updatable) The target state for the Distributed Autonomous Database. Could be set to `ACTIVE` or `INACTIVE`. 
* `change_db_backup_config_trigger` - (Optional) (Updatable) An optional property when incremented triggers Change Db Backup Config. Could be set to any integer value.
* `configure_sharding_trigger` - (Optional) (Updatable) An optional property when incremented triggers Configure Sharding. Could be set to any integer value.
* `download_gsm_certificate_signing_request_trigger` - (Optional) (Updatable) An optional property when incremented triggers Download Gsm Certificate Signing Request. Could be set to any integer value.
* `generate_gsm_certificate_signing_request_trigger` - (Optional) (Updatable) An optional property when incremented triggers Generate Gsm Certificate Signing Request. Could be set to any integer value.
* `generate_wallet_trigger` - (Optional) (Updatable) An optional property when incremented triggers Generate Wallet. Could be set to any integer value.
* `upload_signed_certificate_and_generate_wallet_trigger` - (Optional) (Updatable) An optional property when incremented triggers Upload Signed Certificate And Generate Wallet. Could be set to any integer value.
* `validate_network_trigger` - (Optional) (Updatable) An optional property when incremented triggers Validate Network. Could be set to any integer value.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Distributed Autonomous Database
	* `update` - (Defaults to 20 minutes), when updating the Distributed Autonomous Database
	* `delete` - (Defaults to 20 minutes), when destroying the Distributed Autonomous Database


## Import

DistributedAutonomousDatabases can be imported using the `id`, e.g.

```
$ terraform import oci_distributed_database_distributed_autonomous_database.test_distributed_autonomous_database "id"
```

