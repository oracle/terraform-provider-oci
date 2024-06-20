---
subcategory: "Globally Distributed Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_globally_distributed_database_sharded_database"
sidebar_current: "docs-oci-datasource-globally_distributed_database-sharded_database"
description: |-
  Provides details about a specific Sharded Database in Oracle Cloud Infrastructure Globally Distributed Database service
---

# Data Source: oci_globally_distributed_database_sharded_database
This data source provides details about a specific Sharded Database resource in Oracle Cloud Infrastructure Globally Distributed Database service.

Gets the details of the Sharded database identified by given id.


## Example Usage

```hcl
data "oci_globally_distributed_database_sharded_database" "test_sharded_database" {
	#Required
	sharded_database_id = oci_globally_distributed_database_sharded_database.test_sharded_database.id

	#Optional
	metadata = var.sharded_database_metadata
}
```

## Argument Reference

The following arguments are supported:

* `metadata` - (Optional) Comma separated names of argument corresponding to which metadata need to be retrived, namely VM_CLUSTER_INFO, ADDITIONAL_RESOURCE_INFO. An example is metadata=VM_CLUSTER_INFO,ADDITIONAL_RESOURCE_INFO. 
* `sharded_database_id` - (Required) Sharded Database identifier


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

