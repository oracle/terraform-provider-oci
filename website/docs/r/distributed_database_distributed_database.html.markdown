---
subcategory: "Distributed Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_distributed_database_distributed_database"
sidebar_current: "docs-oci-resource-distributed_database-distributed_database"
description: |-
  Provides the Distributed Database resource in Oracle Cloud Infrastructure Distributed Database service
---

# oci_distributed_database_distributed_database
This resource provides the Distributed Database resource in Oracle Cloud Infrastructure Distributed Database service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/globally-distributed-database/latest/DistributedDatabase

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/distributed_database

Creates a Globally distributed database.

  Patch operation to add, remove or update shards to the Globally distributed database topology. In single patch
operation, multiple shards can be either added, or removed or updated. Combination of inserts, update
and remove in single operation is not allowed.


## Example Usage

```hcl
resource "oci_distributed_database_distributed_database" "test_distributed_database" {
	#Required
	catalog_details {
		#Required
		admin_password = var.distributed_database_catalog_details_admin_password
		source = var.distributed_database_catalog_details_source
		vm_cluster_id = oci_database_vm_cluster.test_vm_cluster.id

		#Optional
		kms_key_id = oci_kms_key.test_key.id
		kms_key_version_id = oci_kms_key_version.test_key_version.id
		peer_details {
			#Required
			vm_cluster_id = oci_database_vm_cluster.test_vm_cluster.id

			#Optional
			protection_mode = var.distributed_database_catalog_details_peer_details_protection_mode
			transport_type = var.distributed_database_catalog_details_peer_details_transport_type
		}
		peer_vm_cluster_ids = var.distributed_database_catalog_details_peer_vm_cluster_ids
		shard_space = var.distributed_database_catalog_details_shard_space
		vault_id = oci_kms_vault.test_vault.id
	}
	character_set = var.distributed_database_character_set
	compartment_id = var.compartment_id
	database_version = var.distributed_database_database_version
	db_deployment_type = var.distributed_database_db_deployment_type
	display_name = var.distributed_database_display_name
	distributed_database_id = var.distributed_database_distributed_database_id
	listener_port = var.distributed_database_listener_port
	ncharacter_set = var.distributed_database_ncharacter_set
	ons_port_local = var.distributed_database_ons_port_local
	ons_port_remote = var.distributed_database_ons_port_remote
	prefix = var.distributed_database_prefix
	private_endpoint_ids = var.distributed_database_private_endpoint_ids
	shard_details {
		#Required
		admin_password = var.distributed_database_shard_details_admin_password
		source = var.distributed_database_shard_details_source
		vm_cluster_id = oci_database_vm_cluster.test_vm_cluster.id

		#Optional
		kms_key_id = oci_kms_key.test_key.id
		kms_key_version_id = oci_kms_key_version.test_key_version.id
		peer_details {
			#Required
			vm_cluster_id = oci_database_vm_cluster.test_vm_cluster.id

			#Optional
			protection_mode = var.distributed_database_shard_details_peer_details_protection_mode
			transport_type = var.distributed_database_shard_details_peer_details_transport_type
		}
		peer_vm_cluster_ids = var.distributed_database_shard_details_peer_vm_cluster_ids
		shard_space = var.distributed_database_shard_details_shard_space
		vault_id = oci_kms_vault.test_vault.id
	}
	sharding_method = var.distributed_database_sharding_method

	#Optional
	chunks = var.distributed_database_chunks
	db_backup_config {

		#Optional
		auto_backup_window = var.distributed_database_db_backup_config_auto_backup_window
		auto_full_backup_day = var.distributed_database_db_backup_config_auto_full_backup_day
		auto_full_backup_window = var.distributed_database_db_backup_config_auto_full_backup_window
		backup_deletion_policy = var.distributed_database_db_backup_config_backup_deletion_policy
		backup_destination_details {
			#Required
			type = var.distributed_database_db_backup_config_backup_destination_details_type

			#Optional
			dbrs_policy_id = oci_identity_policy.test_policy.id
			id = var.distributed_database_db_backup_config_backup_destination_details_id
			internet_proxy = var.distributed_database_db_backup_config_backup_destination_details_internet_proxy
			is_remote = var.distributed_database_db_backup_config_backup_destination_details_is_remote
			is_zero_data_loss_enabled = var.distributed_database_db_backup_config_backup_destination_details_is_zero_data_loss_enabled
			remote_region = var.distributed_database_db_backup_config_backup_destination_details_remote_region
			vpc_password = var.distributed_database_db_backup_config_backup_destination_details_vpc_password
			vpc_user = var.distributed_database_db_backup_config_backup_destination_details_vpc_user
		}
		can_run_immediate_full_backup = var.distributed_database_db_backup_config_can_run_immediate_full_backup
		is_auto_backup_enabled = var.distributed_database_db_backup_config_is_auto_backup_enabled
		is_remote_backup_enabled = var.distributed_database_db_backup_config_is_remote_backup_enabled
		recovery_window_in_days = var.distributed_database_db_backup_config_recovery_window_in_days
		remote_region = var.distributed_database_db_backup_config_remote_region
	}
	defined_tags = {"foo-namespace.bar-key"= "value"}
	freeform_tags = {"bar-key"= "value"}
	gsm_ssh_public_key = var.distributed_database_gsm_ssh_public_key
	listener_port_tls = var.distributed_database_listener_port_tls
	patch_operations {
		#Required
		operation = var.distributed_database_patch_operations_operation
		selection = var.distributed_database_patch_operations_selection

		#Optional
		value = var.distributed_database_patch_operations_value
	}
	replication_factor = var.distributed_database_replication_factor
	replication_method = var.distributed_database_replication_method
	replication_unit = var.distributed_database_replication_unit
}
```

## Argument Reference

The following arguments are supported:

* `catalog_details` - (Required) Collection of catalog for the Globally distributed database.
	* `admin_password` - (Required) The admin password for the cataog associated with Globally distributed database.
	* `kms_key_id` - (Optional) The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
	* `kms_key_version_id` - (Optional) The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. 
	* `peer_details` - (Optional) The details required for creation of the peer for the ExadbXs infrastructure based catalog.
		* `protection_mode` - (Optional) The protectionMode for the catalog peer.
		* `transport_type` - (Optional) The redo transport type to use for this Data Guard association.
		* `vm_cluster_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM Cluster for the catalog peer.
	* `peer_vm_cluster_ids` - (Optional) This field is deprecated. This should not be used while creation of new distributed database. To set the peers on catalog of distributed database please use peerDetails. 
	* `shard_space` - (Optional) The shard space name for the Globally distributed database. Shard space for existing shard cannot be changed, once shard is created. Shard space name shall be used while creation of new shards. 
	* `source` - (Required) The source of Globally distributed database type: Use EXADB_XS for the Globally distributed database with Exascale based distributed database. 
	* `vault_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `kmsKeyId` are required for Customer Managed Keys.
	* `vm_cluster_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VmCluster.
* `character_set` - (Required) The character set for the database.
* `chunks` - (Optional) Number of chunks in a shardspace. The value of chunks must be greater than 2 times the size of the largest shardgroup in any shardspace. Chunks is required to be provided for distributed databases being created with SYSTEM shardingMethod. For USER shardingMethod, chunks should not be set in create payload. 
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Globally distributed database compartment.
* `database_version` - (Required) Oracle Database version for the shards and catalog used in Globally distributed database.
* `db_backup_config` - (Optional) Backup Options To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to an administrator. If you're an administrator who needs to write policies to give users access, see [Getting Started with Policies](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/policygetstarted.htm). 
	* `auto_backup_window` - (Optional) Time window selected for initiating automatic backup for the database system. There are twelve available two-hour time windows. If no option is selected, a start time between 12:00 AM to 7:00 AM in the region of the database is automatically chosen. For example, if the user selects SLOT_TWO from the enum list, the automatic backup job will start in between 2:00 AM (inclusive) to 4:00 AM (exclusive). Example: `SLOT_TWO` 
	* `auto_full_backup_day` - (Optional) Day of the week the full backup should be applied on the database system. If no option is selected, the value is null and we will default to Sunday.
	* `auto_full_backup_window` - (Optional) Time window selected for initiating full backup for the database system. There are twelve available two-hour time windows. If no option is selected, the value is null and a start time between 12:00 AM to 7:00 AM in the region of the database is automatically chosen. For example, if the user selects SLOT_TWO from the enum list, the automatic backup job will start in between 2:00 AM (inclusive) to 4:00 AM (exclusive). Example: `SLOT_TWO` 
	* `backup_deletion_policy` - (Optional) This defines when the backups will be deleted. - IMMEDIATE option keep the backup for predefined time i.e 72 hours and then delete permanently... - RETAIN will keep the backups as per the policy defined for database backups.
	* `backup_destination_details` - (Optional) Backup destination details.
		* `dbrs_policy_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DBRS policy used for backup.
		* `id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup destination.
		* `internet_proxy` - (Optional) Proxy URL to connect to object store.
		* `is_remote` - (Optional) Indicates whether the backup destination is cross-region or local region.
		* `is_zero_data_loss_enabled` - (Optional) Indicates whether Zero Data Loss functionality is enabled for a Recovery Appliance backup destination in an Autonomous Container Database. When enabled, the database automatically ships all redo logs in real-time to the Recovery Appliance for a Zero Data Loss recovery setup (sub-second RPO). Defaults to `TRUE` if no value is given.
		* `remote_region` - (Optional) The name of the remote region where the remote automatic incremental backups will be stored. For information about valid region names, see [Regions and Availability Domains](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/regions.htm). 
		* `type` - (Required) Type of the database backup destination.
		* `vpc_password` - (Optional) For a RECOVERY_APPLIANCE backup destination, the password for the VPC user that is used to access the Recovery Appliance.
		* `vpc_user` - (Optional) For a RECOVERY_APPLIANCE backup destination, the Virtual Private Catalog (VPC) user that is used to access the Recovery Appliance.
	* `can_run_immediate_full_backup` - (Optional) If set to true, configures automatic full backups in the local region (the region of the DB system) for the first backup run immediately.
	* `is_auto_backup_enabled` - (Optional) If set to true, configures automatic backups. If you previously used RMAN or dbcli to configure backups and then you switch to using the Console or the API for backups, a new backup configuration is created and associated with your database. This means that you can no longer rely on your previously configured unmanaged backups to work.
	* `is_remote_backup_enabled` - (Optional) If set to true, configures automatic incremental backups in the local region (the region of the DB system) and the remote region with a default frequency of 1 hour. If you previously used RMAN or dbcli to configure backups, using the Console or the API for manged backups creates a new backup configuration for your database. The new configuration replaces the configuration created with RMAN or dbcli. This means that you can no longer rely on your previously configured unmanaged backups to work. 
	* `recovery_window_in_days` - (Optional) Number of days between the current and the earliest point of recoverability covered by automatic backups. This value applies to automatic backups only. After a new automatic backup has been created, Oracle removes old automatic backups that are created before the window. When the value is updated, it is applied to all existing automatic backups. 
	* `remote_region` - (Optional) The name of the remote region where the remote automatic incremental backups will be stored. For information about valid region names, see [Regions and Availability Domains](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/regions.htm). 
* `db_deployment_type` - (Required) The distributed database deployment type. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Required) (Updatable) The display name of the Globally distributed database.
* `distributed_database_id` - (Required) 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `gsm_ssh_public_key` - (Optional) The SSH public key for Global service manager instances.
* `listener_port` - (Required) The listener port number for the Globally distributed database. The listener port number has to be unique for a customer tenancy across all distributed databases. Same port number should not be re-used for any other distributed database. 
* `listener_port_tls` - (Optional) The TLS listener port number for the Globally distributed database. The TLS listener port number has to be unique for a customer tenancy across all distributed databases. Same port number should not be re-used for any other distributed database. For BASE_DB and EXADB_XS based distributed databases, tls is not supported hence the listenerPortTls is not needed to be provided in create payload. 
* `ncharacter_set` - (Required) The national character set for the database.
* `ons_port_local` - (Required) The ons local port number for the Globally distributed database. The onsPortLocal has to be unique for a customer tenancy across all distributed databases. Same port number should not be re-used for any other distributed database. 
* `ons_port_remote` - (Required) The ons remote port number for the Globally distributed database. The onsPortRemote has to be unique for a customer tenancy across all distributed databases. Same port number should not be re-used for any other distributed database. 
* `patch_operations` - (Optional) (Updatable) 
	* `operation` - (Required) (Updatable) The operation can be one of these values: `INSERT`, `MERGE`, `REMOVE`
	* `selection` - (Required) (Updatable) 
	* `value` - (Required when operation=INSERT | MERGE) (Updatable) 
* `prefix` - (Required) Unique name prefix for the Globally distributed databases. Only alpha-numeric values are allowed. First character has to be a letter followed by any combination of letter and number. 
* `private_endpoint_ids` - (Required) The collection of [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the private endpoint associated with Globally distributed autonomous database.
* `replication_factor` - (Optional) The Replication factor for RAFT replication based Globally distributed database. Currently supported values are 3, 5 and 7. 
* `replication_method` - (Optional) The Replication method for Globally distributed database. Use RAFT for Raft based replication. With RAFT replication, shards cannot have peers details set on them. In case shards need to have peers, please do not set RAFT replicationMethod. For all non RAFT replication cases (with or without peers), please set replicationMethod as DG or do not set any value for replicationMethod. 
* `replication_unit` - (Optional) The replication unit count for RAFT based distributed database. For RAFT replication based Globally distributed database, the value should be at least twice the number of shards. 
* `shard_details` - (Required) Collection of shards for the Globally distributed database.
	* `admin_password` - (Required) The admin password for the shard associated with Globally distributed database.
	* `kms_key_id` - (Optional) The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
	* `kms_key_version_id` - (Optional) The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. 
	* `peer_details` - (Optional) The details required for creation of the peer for the ExadbXs infrastructure based shard.
		* `protection_mode` - (Optional) The protectionMode for the shard peer.
		* `transport_type` - (Optional) The redo transport type to use for this Data Guard association.
		* `vm_cluster_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM Cluster for the shard peer.
	* `peer_vm_cluster_ids` - (Optional) This field is deprecated. This should not be used while creation of new distributed database. To set the peers on new shards of distributed database please use peerDetails. 
	* `shard_space` - (Optional) The shard space name for the Globally distributed database. Shard space for existing shard cannot be changed, once shard is created. Shard space name shall be used while creation of new shards. 
	* `source` - (Required) The source of Globally distributed database type: Use EXADB_XS for the Globally distributed database with Exascale based distributed database. 
	* `vault_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `kmsKeyId` are required for Customer Managed Keys.
	* `vm_cluster_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VmCluster.
* `sharding_method` - (Required) Sharding Methods for the Globally distributed database.
* `state` - (Optional) (Updatable) The target state for the Distributed Database. Could be set to `ACTIVE` or `INACTIVE`. 
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

* `catalog_details` - Collection of catalogs associated with the Globally distributed database.
	* `container_database_id` - the identifier of the container database for underlying supporting resource.
	* `db_home_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Home.
	* `kms_key_id` - The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
	* `kms_key_version_id` - The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. 
	* `metadata` - Additional metadata related to Globally distributed database resources.
		* `map` - The map containing key-value pair of additional metadata.
	* `name` - The name of catalog.
	* `peer_details` - Peer details for the catalog.
		* `container_database_id` - the identifier of the container database for underlying supporting resource.
		* `metadata` - Additional metadata related to Globally distributed database resources.
			* `map` - The map containing key-value pair of additional metadata.
		* `protection_mode` - The protectionMode for the catalog peer.
		* `shard_group` - The name of the shardGroup for the peer.
		* `status` - Status of EXADB_XS based catalog peer.
		* `supporting_resource_id` - the identifier of the underlying supporting resource.
		* `time_created` - The time the catalog peer was created. An RFC3339 formatted datetime string
		* `time_updated` - The time the catalog peer was last updated. An RFC3339 formatted datetime string
		* `transport_type` - The redo transport type to use for this Data Guard association.
		* `vm_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VmCluster.
	* `shard_group` - The name of the shardGroup for the catalog.
	* `source` - The source of Globally distributed database type: Use EXADB_XS for the Globally distributed database with Exascale based distributed database. 
	* `status` - Status of EXADB_XS based catalog.
	* `supporting_resource_id` - the identifier of the underlying supporting resource.
	* `time_created` - The time the catalog was created. An RFC3339 formatted datetime string
	* `time_updated` - The time the catalog was last updated. An RFC3339 formatted datetime string
	* `vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `kmsKeyId` are required for Customer Managed Keys.
	* `vm_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VmCluster.
* `character_set` - The character set for the database.
* `chunks` - The default number of unique chunks in a shardspace. The value of chunks must be greater than 2 times the size of the largest shardgroup in any shardspace. 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Globally distributed database compartment.
* `connection_strings` - Details of Globally distributed database connection String.
	* `all_connection_strings` - Collection of connection strings.
* `database_version` - Oracle Database version for the shards and catalog used in Globally distributed database.
* `db_backup_config` - Backup Options To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to an administrator. If you're an administrator who needs to write policies to give users access, see [Getting Started with Policies](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/policygetstarted.htm). 
	* `auto_backup_window` - Time window selected for initiating automatic backup for the database system. There are twelve available two-hour time windows. If no option is selected, a start time between 12:00 AM to 7:00 AM in the region of the database is automatically chosen. For example, if the user selects SLOT_TWO from the enum list, the automatic backup job will start in between 2:00 AM (inclusive) to 4:00 AM (exclusive). Example: `SLOT_TWO` 
	* `auto_full_backup_day` - Day of the week the full backup should be applied on the database system. If no option is selected, the value is null and we will default to Sunday.
	* `auto_full_backup_window` - Time window selected for initiating full backup for the database system. There are twelve available two-hour time windows. If no option is selected, the value is null and a start time between 12:00 AM to 7:00 AM in the region of the database is automatically chosen. For example, if the user selects SLOT_TWO from the enum list, the automatic backup job will start in between 2:00 AM (inclusive) to 4:00 AM (exclusive). Example: `SLOT_TWO` 
	* `backup_deletion_policy` - This defines when the backups will be deleted. - IMMEDIATE option keep the backup for predefined time i.e 72 hours and then delete permanently... - RETAIN will keep the backups as per the policy defined for database backups.
	* `backup_destination_details` - Backup destination details.
		* `dbrs_policy_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DBRS policy used for backup.
		* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup destination.
		* `internet_proxy` - Proxy URL to connect to object store.
		* `is_remote` - Indicates whether the backup destination is cross-region or local region.
		* `is_zero_data_loss_enabled` - Indicates whether Zero Data Loss functionality is enabled for a Recovery Appliance backup destination in an Autonomous Container Database. When enabled, the database automatically ships all redo logs in real-time to the Recovery Appliance for a Zero Data Loss recovery setup (sub-second RPO). Defaults to `TRUE` if no value is given.
		* `remote_region` - The name of the remote region where the remote automatic incremental backups will be stored. For information about valid region names, see [Regions and Availability Domains](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/regions.htm). 
		* `type` - Type of the database backup destination.
		* `vpc_password` - For a RECOVERY_APPLIANCE backup destination, the password for the VPC user that is used to access the Recovery Appliance.
		* `vpc_user` - For a RECOVERY_APPLIANCE backup destination, the Virtual Private Catalog (VPC) user that is used to access the Recovery Appliance.
	* `can_run_immediate_full_backup` - If set to true, configures automatic full backups in the local region (the region of the DB system) for the first backup run immediately.
	* `is_auto_backup_enabled` - If set to true, configures automatic backups. If you previously used RMAN or dbcli to configure backups and then you switch to using the Console or the API for backups, a new backup configuration is created and associated with your database. This means that you can no longer rely on your previously configured unmanaged backups to work.
	* `is_remote_backup_enabled` - If set to true, configures automatic incremental backups in the local region (the region of the DB system) and the remote region with a default frequency of 1 hour. If you previously used RMAN or dbcli to configure backups, using the Console or the API for manged backups creates a new backup configuration for your database. The new configuration replaces the configuration created with RMAN or dbcli. This means that you can no longer rely on your previously configured unmanaged backups to work. 
	* `recovery_window_in_days` - Number of days between the current and the earliest point of recoverability covered by automatic backups. This value applies to automatic backups only. After a new automatic backup has been created, Oracle removes old automatic backups that are created before the window. When the value is updated, it is applied to all existing automatic backups. 
	* `remote_region` - The name of the remote region where the remote automatic incremental backups will be stored. For information about valid region names, see [Regions and Availability Domains](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/regions.htm). 
* `db_deployment_type` - The distributed database deployment type. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - The display name of the Globally distributed database.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `gsm_details` - Collection of catalogs associated with the Globally distributed database.
	* `compute_count` - The compute count for the Global service manager instance.
	* `data_storage_size_in_gbs` - The data disk group size to be allocated in GBs for the Global service manager instance.
	* `gsm_image_details` - The Global service manager image details
		* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Global service manager software image.
		* `version_number` - The version number associated with the image identified by id.
	* `metadata` - Additional metadata related to Globally distributed database resources.
		* `map` - The map containing key-value pair of additional metadata.
	* `name` - Name of the Global service manager instance
	* `status` - Status of the gsm.
	* `supporting_resource_id` - the identifier of the underlying supporting resource.
	* `time_created` - The time the Global service manager instance was created. An RFC3339 formatted datetime string
	* `time_ssl_certificate_expires` - The time the ssl certificate associated with Global service manager expires. An RFC3339 formatted datetime string
	* `time_updated` - The time the Global service manager instance was last updated. An RFC3339 formatted datetime string
* `gsm_ssh_public_key` - The SSH public key for Global service manager instances.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Globally distributed database.
* `latest_gsm_image_details` - The Global service manager image details
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Global service manager software image.
	* `version_number` - The version number associated with the image identified by id.
* `lifecycle_details` - The lifecycleDetails for the Globally distributed database.
* `listener_port` - The Global service manager listener port number for the Globally distributed database.
* `listener_port_tls` - The TLS listener port number for Globally distributed database.
* `metadata` - Additional metadata related to Globally distributed database resources.
	* `map` - The map containing key-value pair of additional metadata.
* `ncharacter_set` - The national character set for the database.
* `ons_port_local` - Ons local port number.
* `ons_port_remote` - Ons remote port number.
* `prefix` - Unique name prefix for the Globally distributed databases. Only alpha-numeric values are allowed. First character has to be a letter followed by any combination of letter and number. 
* `private_endpoint_ids` - The collection of [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the private endpoint associated with Globally distributed autonomous database.
* `replication_factor` - The Replication factor for RAFT replication based Globally distributed database. Currently supported values are 3, 5 and 7. 
* `replication_method` - The Replication method for Globally distributed database. Use RAFT for Raft replication, and DG for DataGuard. If replicationMethod is not provided, it defaults to DG. 
* `replication_unit` - The replication unit count for RAFT based distributed database. For RAFT replication based Globally distributed database, the value should be at least twice the number of shards. 
* `shard_details` - Collection of shards associated with the Globally distributed database.
	* `container_database_id` - the identifier of the container database for underlying supporting resource.
	* `db_home_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Home.
	* `kms_key_id` - The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
	* `kms_key_version_id` - The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. 
	* `metadata` - Additional metadata related to Globally distributed database resources.
		* `map` - The map containing key-value pair of additional metadata.
	* `name` - Name of the shard.
	* `peer_details` - Peer details for the shard.
		* `container_database_id` - the identifier of the container database for underlying supporting resource.
		* `metadata` - Additional metadata related to Globally distributed database resources.
			* `map` - The map containing key-value pair of additional metadata.
		* `protection_mode` - The protectionMode for the shard peer.
		* `shard_group` - The name of the shardGroup for the peer.
		* `status` - Status of EXADB_XS based shard peer.
		* `supporting_resource_id` - the identifier of the underlying supporting resource.
		* `time_created` - The time the shard peer was created. An RFC3339 formatted datetime string
		* `time_updated` - The time the shard peer was last updated. An RFC3339 formatted datetime string
		* `transport_type` - The redo transport type to use for this Data Guard association.
		* `vm_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VmCluster.
	* `shard_group` - The name of the shardGroup for the shard.
	* `shard_space` - The shard space name for the Globally distributed database. Shard space for existing shard cannot be changed, once shard is created. Shard space name shall be used while creation of new shards. 
	* `source` - The source of Globally distributed database type: Use EXADB_XS for the Globally distributed database with Exascale based distributed database. 
	* `status` - Status of EXADB_XS based shard.
	* `supporting_resource_id` - the identifier of the underlying supporting resource.
	* `time_created` - The time the shard was created. An RFC3339 formatted datetime string
	* `time_updated` - The time the shard was last updated. An RFC3339 formatted datetime string
	* `vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `kmsKeyId` are required for Customer Managed Keys.
	* `vm_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VmCluster.
* `sharding_method` - Sharding Methods for the Globally distributed database.
* `state` - Lifecycle states for the Globally distributed database.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the Globally distributed database was created. An RFC3339 formatted datetime string
* `time_updated` - The time the Globally distributed database was last updated. An RFC3339 formatted datetime string

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Distributed Database
	* `update` - (Defaults to 20 minutes), when updating the Distributed Database
	* `delete` - (Defaults to 20 minutes), when destroying the Distributed Database


## Import

DistributedDatabases can be imported using the `id`, e.g.

```
$ terraform import oci_distributed_database_distributed_database.test_distributed_database "id"
```

