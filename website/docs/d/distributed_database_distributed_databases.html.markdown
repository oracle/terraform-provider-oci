---
subcategory: "Distributed Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_distributed_database_distributed_databases"
sidebar_current: "docs-oci-datasource-distributed_database-distributed_databases"
description: |-
  Provides the list of Distributed Databases in Oracle Cloud Infrastructure Distributed Database service
---

# Data Source: oci_distributed_database_distributed_databases
This data source provides the list of Distributed Databases in Oracle Cloud Infrastructure Distributed Database service.

List of Globally distributed databases.


## Example Usage

```hcl
data "oci_distributed_database_distributed_databases" "test_distributed_databases" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	db_deployment_type = var.distributed_database_db_deployment_type
	display_name = var.distributed_database_display_name
	metadata {
	}
	state = var.distributed_database_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `db_deployment_type` - (Optional) A filter to return only resources their dbDeploymentType matches the given dbDeploymentType.
* `display_name` - (Optional) A filter to return only Globally distributed databases that match the entire name given. The match is not case sensitive.
* `metadata` - (Optional) Comma separated names of argument corresponding to which metadata need to be retrived. 
* `state` - (Optional) A filter to return only resources their lifecycleState matches the given lifecycleState.


## Attributes Reference

The following attributes are exported:

* `distributed_database_collection` - The list of distributed_database_collection.

### DistributedDatabase Reference

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

