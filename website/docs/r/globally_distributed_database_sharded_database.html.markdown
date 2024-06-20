---
subcategory: "Globally Distributed Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_globally_distributed_database_sharded_database"
sidebar_current: "docs-oci-resource-globally_distributed_database-sharded_database"
description: |-
  Provides the Sharded Database resource in Oracle Cloud Infrastructure Globally Distributed Database service
---

# oci_globally_distributed_database_sharded_database
This resource provides the Sharded Database resource in Oracle Cloud Infrastructure Globally Distributed Database service.

Creates a Sharded Database.

  Patch operation to add, remove or update shards to the sharded database topology. In single patch
operation, multiple shards can be either added, or removed or updated. Combination of inserts, update
and remove in single operation is not allowed.


## Example Usage

```hcl
resource "oci_globally_distributed_database_sharded_database" "test_sharded_database" {
	#Required
	catalog_details {
		#Required
		admin_password = var.sharded_database_catalog_details_admin_password
		cloud_autonomous_vm_cluster_id = oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster.id
		compute_count = var.sharded_database_catalog_details_compute_count
		data_storage_size_in_gbs = var.sharded_database_catalog_details_data_storage_size_in_gbs
		is_auto_scaling_enabled = var.sharded_database_catalog_details_is_auto_scaling_enabled

		#Optional
		encryption_key_details {
			#Required
			kms_key_id = oci_kms_key.test_key.id
			vault_id = oci_kms_vault.test_vault.id

			#Optional
			kms_key_version_id = oci_kms_key_version.test_key_version.id
		}
		peer_cloud_autonomous_vm_cluster_id = oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster.id
	}
	character_set = var.sharded_database_character_set
	compartment_id = var.compartment_id
	db_deployment_type = var.sharded_database_db_deployment_type
	db_version = var.sharded_database_db_version
	db_workload = var.sharded_database_db_workload
	display_name = var.sharded_database_display_name
	listener_port = var.sharded_database_listener_port
	listener_port_tls = var.sharded_database_listener_port_tls
	ncharacter_set = var.sharded_database_ncharacter_set
	ons_port_local = var.sharded_database_ons_port_local
	ons_port_remote = var.sharded_database_ons_port_remote
	prefix = var.sharded_database_prefix
	shard_details {
		#Required
		admin_password = var.sharded_database_shard_details_admin_password
		cloud_autonomous_vm_cluster_id = oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster.id
		compute_count = var.sharded_database_shard_details_compute_count
		data_storage_size_in_gbs = var.sharded_database_shard_details_data_storage_size_in_gbs
		is_auto_scaling_enabled = var.sharded_database_shard_details_is_auto_scaling_enabled

		#Optional
		encryption_key_details {
			#Required
			kms_key_id = oci_kms_key.test_key.id
			vault_id = oci_kms_vault.test_vault.id

			#Optional
			kms_key_version_id = oci_kms_key_version.test_key_version.id
		}
		peer_cloud_autonomous_vm_cluster_id = oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster.id
		shard_space = var.sharded_database_shard_details_shard_space
	}
	sharded_database_id = var.sharded_database_sharded_database_id
	sharding_method = var.sharded_database_sharding_method

	#Optional
	chunks = var.sharded_database_chunks
	cluster_certificate_common_name = var.sharded_database_cluster_certificate_common_name
	defined_tags = {"foo-namespace.bar-key"= "value"}
	freeform_tags = {"bar-key"= "value"}
	patch_operations {
		#Required
		operation = var.sharded_database_patch_operations_operation
		selection = var.sharded_database_patch_operations_selection

		#Optional
		value = var.sharded_database_patch_operations_value
	}
}
```

## Argument Reference

The following arguments are supported:

* `catalog_details` - (Required) Collection of ATP-Dedicated catalogs that needs to be created.
	* `admin_password` - (Required) Admin password for the catalog database.
	* `cloud_autonomous_vm_cluster_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud Autonomous Exadata VM Cluster.
	* `compute_count` - (Required) The compute count for the catalog database. It has to be in multiple of 2.
	* `data_storage_size_in_gbs` - (Required) The data disk group size to be allocated in GBs for the catalog database.
	* `encryption_key_details` - (Optional) Details of encryption key to be used to encrypt data for shards and catalog for sharded database. For system-defined sharding type, all shards have to use same encryptionKeyDetails. For system-defined sharding, if encryptionKeyDetails are not specified for catalog, then Oracle managed key will be used for catalog. For user-defined sharding type, if encryptionKeyDetails are not provided for any shard or catalog, then Oracle managed key will be used for such shard or catalog. For system-defined or user-defined sharding type, if the shard or catalog has a peer in region other than primary shard or catalog region, then make sure to provide virtual vault for such shard or catalog, which is also replicated to peer region (the region where peer or standby shard or catalog exists). 
		* `kms_key_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the KMS key in vault identified by vaultId in customer tenancy  that is used as the master encryption key. 
		* `kms_key_version_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the KMS key version for key identified by kmsKeyId that is used in data encryption (TDE) operations. 
		* `vault_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the vault in customer tenancy where KMS key is present. For shard or catalog with cross-region data guard enabled, user needs to make sure to provide virtual private vault only, which is also replicated in the region of standby shard. 
	* `is_auto_scaling_enabled` - (Required) Determines the auto-scaling mode for the catalog database.
	* `peer_cloud_autonomous_vm_cluster_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the peer cloud Autonomous Exadata VM Cluster.
* `character_set` - (Required) The character set for the new shard database being created. Use database api ListAutonomousDatabaseCharacterSets to get the list of allowed character set for autonomous dedicated database. See documentation: https://docs.oracle.com/en-us/iaas/api/#/en/database/20160918/AutonomousDatabaseCharacterSets/ListAutonomousDatabaseCharacterSets 
* `chunks` - (Optional) The default number of unique chunks in a shardspace. The value of chunks must be greater than 2 times the size of the largest shardgroup in any shardspace. 
* `cluster_certificate_common_name` - (Optional) The certificate common name used in all cloudAutonomousVmClusters for the sharded database topology. Eg. Production. All the clusters used in one sharded database topology shall have same CABundle setup. Valid characterset for clusterCertificateCommonName include uppercase or lowercase letters, numbers, hyphens, underscores, and period. 
* `compartment_id` - (Required) (Updatable) Identifier of the compartment where sharded database is to be created.
* `db_deployment_type` - (Required) The database deployment type. 
* `db_version` - (Required) Oracle Database version of the Autonomous Container Database.
* `db_workload` - (Required) Possible workload types.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Required) (Updatable) Oracle sharded database display name.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `listener_port` - (Required) The listener port number for sharded database.
* `listener_port_tls` - (Required) The TLS listener port number for sharded database.
* `ncharacter_set` - (Required) The national character set for the new shard database being created. Use database api ListAutonomousDatabaseCharacterSets to get the list of allowed national character set for autonomous dedicated database. See documentation: https://docs.oracle.com/en-us/iaas/api/#/en/database/20160918/AutonomousDatabaseCharacterSets/ListAutonomousDatabaseCharacterSets 
* `ons_port_local` - (Required) Ons port local for sharded database.
* `ons_port_remote` - (Required) Ons remote port for sharded database.
* `patch_operations` - (Optional) (Updatable) 
	* `operation` - (Required) (Updatable) The operation can be one of these values: `INSERT`, `MERGE`, `REMOVE`
	* `selection` - (Required) (Updatable) 
	* `value` - (Required when operation=INSERT | MERGE) (Updatable) 
* `prefix` - (Required) Unique name prefix for the sharded databases. Only alpha-numeric values are allowed. First character has to be a letter followed by any combination of letter and number. 
* `shard_details` - (Required) Collection of ATP-Dedicated shards that needs to be created.
	* `admin_password` - (Required) Admin password for shard database.
	* `cloud_autonomous_vm_cluster_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud Autonomous Exadata VM Cluster.
	* `compute_count` - (Required) The compute count for the shard database. It has to be in multiples of 2.
	* `data_storage_size_in_gbs` - (Required) The data disk group size to be allocated in GBs for the shard database.
	* `encryption_key_details` - (Optional) Details of encryption key to be used to encrypt data for shards and catalog for sharded database. For system-defined sharding type, all shards have to use same encryptionKeyDetails. For system-defined sharding, if encryptionKeyDetails are not specified for catalog, then Oracle managed key will be used for catalog. For user-defined sharding type, if encryptionKeyDetails are not provided for any shard or catalog, then Oracle managed key will be used for such shard or catalog. For system-defined or user-defined sharding type, if the shard or catalog has a peer in region other than primary shard or catalog region, then make sure to provide virtual vault for such shard or catalog, which is also replicated to peer region (the region where peer or standby shard or catalog exists). 
		* `kms_key_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the KMS key in vault identified by vaultId in customer tenancy  that is used as the master encryption key. 
		* `kms_key_version_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the KMS key version for key identified by kmsKeyId that is used in data encryption (TDE) operations. 
		* `vault_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the vault in customer tenancy where KMS key is present. For shard or catalog with cross-region data guard enabled, user needs to make sure to provide virtual private vault only, which is also replicated in the region of standby shard. 
	* `is_auto_scaling_enabled` - (Required) Determines the auto-scaling mode for the shard database.
	* `peer_cloud_autonomous_vm_cluster_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the peer cloud Autonomous Exadata VM Cluster.
	* `shard_space` - (Optional) The shard space name for the shard database. Shard space for existing shard cannot be changed, once shard is created. Shard space name shall be used while creation of new shards. For User defined sharding, every shard must have a unique shard space name. For system defined sharding, shard space name is not required. 
* `sharded_database_id` - (Required) 
* `sharding_method` - (Required) Sharding Method.
* `configure_gsms_trigger` - (Optional) (Updatable) An optional property when incremented triggers Configure Gsms. Could be set to any integer value.
* `configure_sharding_trigger` - (Optional) (Updatable) An optional property when incremented triggers Configure Sharding. Could be set to any integer value.
* `download_gsm_certificate_signing_request_trigger` - (Optional) (Updatable) An optional property when incremented triggers Download Gsm Certificate Signing Request. Could be set to any integer value.
* `generate_gsm_certificate_signing_request_trigger` - (Optional) (Updatable) An optional property when incremented triggers Generate Gsm Certificate Signing Request. Could be set to any integer value.
* `generate_wallet_trigger` - (Optional) (Updatable) An optional property when incremented triggers Generate Wallet. Could be set to any integer value.
* `get_connection_string_trigger` - (Optional) (Updatable) An optional property when incremented triggers Get Connection String. Could be set to any integer value.
* `start_database_trigger` - (Optional) (Updatable) An optional property when incremented triggers Start Database. Could be set to any integer value.
* `stop_database_trigger` - (Optional) (Updatable) An optional property when incremented triggers Stop Database. Could be set to any integer value.
* `upload_signed_certificate_and_generate_wallet_trigger` - (Optional) (Updatable) An optional property when incremented triggers Upload Signed Certificate And Generate Wallet. Could be set to any integer value.
* `validate_network_trigger` - (Optional) (Updatable) An optional property when incremented triggers Validate Network. Could be set to any integer value.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `catalog_details` - Details of ATP-D based catalogs.
	* `cloud_autonomous_vm_cluster_id` - Identifier of the primary cloudAutonomousVmCluster for the catalog. 
	* `compute_count` - The compute amount available to the underlying autonomous database associated with shard or catalog.
	* `container_database_id` - Identifier of the underlying container database. 
	* `container_database_parent_id` - Identifier of the underlying container database parent. 
	* `data_storage_size_in_gbs` - The data disk group size to be allocated in GBs.
	* `encryption_key_details` - Details of encryption key to be used to encrypt data for shards and catalog for sharded database. For system-defined sharding type, all shards have to use same encryptionKeyDetails. For system-defined sharding, if encryptionKeyDetails are not specified for catalog, then Oracle managed key will be used for catalog. For user-defined sharding type, if encryptionKeyDetails are not provided for any shard or catalog, then Oracle managed key will be used for such shard or catalog. For system-defined or user-defined sharding type, if the shard or catalog has a peer in region other than primary shard or catalog region, then make sure to provide virtual vault for such shard or catalog, which is also replicated to peer region (the region where peer or standby shard or catalog exists). 
		* `kms_key_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the KMS key in vault identified by vaultId in customer tenancy  that is used as the master encryption key. 
		* `kms_key_version_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the KMS key version for key identified by kmsKeyId that is used in data encryption (TDE) operations. 
		* `vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the vault in customer tenancy where KMS key is present. For shard or catalog with cross-region data guard enabled, user needs to make sure to provide virtual private vault only, which is also replicated in the region of standby shard. 
	* `is_auto_scaling_enabled` - Determines the auto-scaling mode.
	* `metadata` - Additional metadata related to catalog's underlying supporting resource.
	* `name` - Catalog name
	* `peer_cloud_autonomous_vm_cluster_id` - Identifier of the peer cloudAutonomousVmCluster for the catalog. 
	* `shard_group` - Name of the shard-group to which the catalog belongs.
	* `status` - Status of shard or catalog or gsm for the sharded database.
	* `supporting_resource_id` - Identifier of the underlying supporting resource. 
	* `time_created` - The time the catalog was created. An RFC3339 formatted datetime string
	* `time_ssl_certificate_expires` - The time the ssl certificate associated with catalog expires. An RFC3339 formatted datetime string
	* `time_updated` - The time the catalog was last created. An RFC3339 formatted datetime string
* `character_set` - The character set for the database.
* `chunks` - The default number of unique chunks in a shardspace. The value of chunks must be greater than 2 times the size of the largest shardgroup in any shardspace. 
* `cluster_certificate_common_name` - The certificate common name used in all cloudAutonomousVmClusters for the sharded database topology. Eg. Production. All the clusters used in one sharded database topology shall have same CABundle setup. Valid characterset for clusterCertificateCommonName include uppercase or lowercase letters, numbers, hyphens, underscores, and period. 
* `compartment_id` - Identifier of the compartment in which sharded database exists.
* `connection_strings` - Details of sharded database connection String.
	* `all_connection_strings` - Collection of connection strings.
* `db_deployment_type` - The database deployment type. 
* `db_version` - Oracle Database version number.
* `db_workload` - Possible workload types.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - Oracle sharded database display name.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `gsms` - Details of GSM instances for the sharded database.
	* `compute_count` - The compute count for the GSM instance.
	* `data_storage_size_in_gbs` - The data disk group size to be allocated in GBs.
	* `metadata` - Additional metadata related to GSM's underlying supporting resource.
	* `name` - Name of the GSM instance
	* `status` - Status of shard or catalog or gsm for the sharded database.
	* `supporting_resource_id` - Identifier of the underlying supporting resource. 
	* `time_created` - The time the GSM instance was created. An RFC3339 formatted datetime string
	* `time_ssl_certificate_expires` - The time the ssl certificate associated with GSM expires. An RFC3339 formatted datetime string
	* `time_updated` - The time the GSM instance was last updated. An RFC3339 formatted datetime string
* `id` - Sharded Database identifier
* `lifecycle_state_details` - Detailed message for the lifecycle state.
* `listener_port` - The GSM listener port number.
* `listener_port_tls` - The TLS listener port number for sharded database.
* `ncharacter_set` - The national character set for the database.
* `ons_port_local` - Ons local port number.
* `ons_port_remote` - Ons remote port number.
* `prefix` - Unique prefix for the sharded database.
* `private_endpoint` - The OCID of private endpoint being used by the sharded database.
* `shard_details` - Details of ATP-D based shards.
	* `cloud_autonomous_vm_cluster_id` - Identifier of the primary cloudAutonomousVmCluster for the shard. 
	* `compute_count` - The compute amount available to the underlying autonomous database associated with shard.
	* `container_database_id` - Identifier of the underlying container database. 
	* `container_database_parent_id` - Identifier of the underlying container database parent. 
	* `data_storage_size_in_gbs` - The data disk group size to be allocated in GBs.
	* `encryption_key_details` - Details of encryption key to be used to encrypt data for shards and catalog for sharded database. For system-defined sharding type, all shards have to use same encryptionKeyDetails. For system-defined sharding, if encryptionKeyDetails are not specified for catalog, then Oracle managed key will be used for catalog. For user-defined sharding type, if encryptionKeyDetails are not provided for any shard or catalog, then Oracle managed key will be used for such shard or catalog. For system-defined or user-defined sharding type, if the shard or catalog has a peer in region other than primary shard or catalog region, then make sure to provide virtual vault for such shard or catalog, which is also replicated to peer region (the region where peer or standby shard or catalog exists). 
		* `kms_key_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the KMS key in vault identified by vaultId in customer tenancy  that is used as the master encryption key. 
		* `kms_key_version_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the KMS key version for key identified by kmsKeyId that is used in data encryption (TDE) operations. 
		* `vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the vault in customer tenancy where KMS key is present. For shard or catalog with cross-region data guard enabled, user needs to make sure to provide virtual private vault only, which is also replicated in the region of standby shard. 
	* `is_auto_scaling_enabled` - Determines the auto-scaling mode.
	* `metadata` - Additional metadata related to shard's underlying supporting resource.
	* `name` - Name of the shard.
	* `peer_cloud_autonomous_vm_cluster_id` - Identifier of the peer cloudAutonomousVmCluster for the shard. 
	* `shard_group` - Name of the shard-group to which the shard belongs.
	* `shard_space` - Shard space name.
	* `status` - Status of shard or catalog or gsm for the sharded database.
	* `supporting_resource_id` - Identifier of the underlying supporting resource. 
	* `time_created` - The time the the shard was created. An RFC3339 formatted datetime string
	* `time_ssl_certificate_expires` - The time the ssl certificate associated with shard expires. An RFC3339 formatted datetime string
	* `time_updated` - The time the shard was last updated. An RFC3339 formatted datetime string
* `sharding_method` - Sharding Method.
* `state` - Lifecycle states for sharded databases.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the the Sharded Database was created. An RFC3339 formatted datetime string
* `time_updated` - The time the Sharded Database was last updated. An RFC3339 formatted datetime string
* `time_zone` - Timezone associated with the sharded database.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Sharded Database
	* `update` - (Defaults to 20 minutes), when updating the Sharded Database
	* `delete` - (Defaults to 20 minutes), when destroying the Sharded Database


## Import

ShardedDatabases can be imported using the `id`, e.g.

```
$ terraform import oci_globally_distributed_database_sharded_database.test_sharded_database "id"
```

