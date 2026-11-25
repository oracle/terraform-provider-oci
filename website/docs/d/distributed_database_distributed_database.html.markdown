---
subcategory: "Distributed Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_distributed_database_distributed_database"
sidebar_current: "docs-oci-datasource-distributed_database-distributed_database"
description: |-
  Provides details about a specific Distributed Database in Oracle Cloud Infrastructure Distributed Database service
---

# Data Source: oci_distributed_database_distributed_database
This data source provides details about a specific Distributed Database resource in Oracle Cloud Infrastructure Distributed Database service.

Gets the details of the Globally distributed database identified by given id.


## Example Usage

```hcl
data "oci_distributed_database_distributed_database" "test_distributed_database" {
	#Required
	distributed_database_id = oci_distributed_database_distributed_database.test_distributed_database.id

	#Optional
	metadata {
	}
}
```

## Argument Reference

The following arguments are supported:

* `distributed_database_id` - (Required) Globally distributed database identifier.
* `metadata` - (Optional) Comma-separated argument names for which metadata should be retrieved. 


## Attributes Reference

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
* `catalog_details` - Catalog details associated with the distributed database.
	* `data_guard_replicas` - The details of data guard replicas for the catalog.
		* `database` - Details of a distributed database catalog.
			* `availability_domain` - The name of the availability domain that the distributed database catalog will be located in.
			* `container_database_id` - the identifier of the container database for underlying supporting resource.
			* `db_home_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Home.
			* `db_storage_vault_details` - The Storage Vault for Distributed Database Resource.
				* `additional_flash_cache_in_percent` - The size of additional Flash Cache in percentage of High Capacity database storage.
				* `autoscale_limit_in_gbs` - The maximum limit, in gigabytes, to which the Vault storage size can automatically scale when auto scaling is enabled for the Database Storage Vault.
				* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for Db Storage Vault.
				* `db_storage_vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata Database Storage Vault.
				* `display_name` - The user-friendly name for the Exadata Database Storage Vault. The name does not need to be unique.
				* `high_capacity_database_storage` - Total storage capacity in GB for vault storage.
				* `is_autoscale_enabled` - Indicates if autoscale feature is enabled for the Database Storage Vault. The default value is `FALSE`.
				* `subscription_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription with which resource needs to be associated with.
			* `exascale_db_storage_vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the exascale db storage vault.
			* `kms_key_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations. 
			* `kms_key_version_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. 
			* `metadata` - Additional metadata related to Globally distributed database resources.
				* `map` - The map containing key-value pair of additional metadata.
			* `name` - The name of the catalog.
			* `protection_mode` - The protection mode used for the Data Guard association.
			* `source` - The source of the Globally distributed database type. Use XS_EXISTING_CLUSTER for a database on Exascale Infrastructure that uses a pre-existing cluster. Use XS_NEW_CLUSTER for a database on Exascale Infrastructure that requires a new cluster created from scratch. Use XS_NEW_VAULT_AND_CLUSTER for a database on Exascale Infrastructure that requires a new vault and cluster created from scratch. Use XD_EXISTING_CLUSTER for a database on Dedicated Infrastructure that uses a pre-existing cluster. Use EXISTING_DB_HOME for a database on a pre-existing database home. 
			* `status` - Status of the distributed database catalog.
			* `supporting_resource_id` - the identifier of the underlying supporting resource.
			* `time_created` - The time the catalog was created. An RFC3339 formatted datetime string.
			* `time_updated` - The time the catalog was last updated. An RFC3339 formatted datetime string.
			* `transport_type` - The transport type used for the Data Guard association.
			* `vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts).  This parameter and `kmsKeyId` are required for Customer Managed Keys. 
			* `vm_cluster_details` - The Exadata VM cluster for Distributed Database Resource.
				* `backup_network_nsg_ids` - A list of the [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that the backup network of this DB system belongs to.  Setting this to an empty array after the list is created removes the resource from all NSGs.  For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). Applicable only to Exadata systems. 
				* `backup_subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup network subnet associated with the Exadata VM cluster on Exascale Infrastructure. 
				* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for VM Cluster.
				* `display_name` - The user-friendly name for the Exadata VM cluster on Exascale Infrastructure. The name does not need to be unique.
				* `domain` - A domain name used for the Exadata VM cluster on Exascale Infrastructure.  If the Oracle-provided internet and VCN resolver is enabled for the specified subnet, then the domain name for the subnet is used (do not provide one). Otherwise, provide a valid DNS domain name.  Hyphens (-) are not permitted. Applies to Exadata Database Service on Exascale Infrastructure only. 
				* `enabled_ecpu_count` - The number of ECPUs to enable for an Exadata VM cluster on Exascale Infrastructure. 
				* `is_diagnostics_events_enabled` - Indicates whether diagnostic collection is enabled for the VM cluster.  Enabling diagnostic collection allows you to receive Events service notifications for guest VM issues.  Diagnostic collection also allows Oracle to provide enhanced service and proactive support for your Exadata system.  You can enable diagnostic collection during VM cluster provisioning. You can also disable or enable it at any time using the `UpdateVmCluster` API. 
				* `is_health_monitoring_enabled` - Indicates whether health monitoring is enabled for the VM cluster.  Enabling health monitoring allows Oracle to collect diagnostic data and share it with its operations and support personnel.  You may also receive notifications for some events. Collecting health diagnostics enables Oracle to provide proactive support and enhanced service for your system. Optionally enable health monitoring while provisioning a system. You can also disable or enable health monitoring anytime using the `UpdateVmCluster` API. 
				* `is_incident_logs_enabled` - Indicates whether incident logs and trace collection are enabled for the VM cluster.  Enabling incident logs collection allows Oracle to receive Events service notifications for guest VM issues, collect incident logs and traces, and use them to diagnose issues and resolve them. Optionally enable incident logs collection while provisioning a system. You can also disable or enable incident logs collection anytime using the `UpdateVmCluster` API. 
				* `license_model` - The Oracle license model that applies to the Exadata VM cluster on Exascale Infrastructure. The default is BRING_YOUR_OWN_LICENSE. 
				* `nsg_ids` - The list of [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the network security groups (NSGs) to which this resource belongs.  Setting this to an empty list removes all resources from all NSGs.  For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). 
				* `private_zone_id` - The private zone ID in which you want DNS records to be created. 
				* `ssh_public_keys` - The public key portion of one or more key pairs used for SSH access to the Exadata VM cluster on Exascale Infrastructure.
				* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet associated with the Exadata VM cluster on Exascale Infrastructure. 
				* `subscription_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription with which resource needs to be associated with.
				* `total_ecpu_count` - The number of Total ECPUs for an Exadata VM cluster on Exascale Infrastructure. 
				* `vm_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM cluster.
				* `vm_file_system_storage_size` - File System Storage Size in GBs for Exadata VM cluster. 
			* `vm_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Cloud VM Cluster.
		* `name` - The name of the catalog replica.  It must start with a letter and contain only letters, digits, and underscores.  Maximum length is 40 characters. 
	* `original_replica` - Details of a catalog replica.
		* `database` - Details of a distributed database catalog.
			* `availability_domain` - The name of the availability domain that the distributed database catalog will be located in.
			* `container_database_id` - the identifier of the container database for underlying supporting resource.
			* `db_home_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Home.
			* `db_storage_vault_details` - The Storage Vault for Distributed Database Resource.
				* `additional_flash_cache_in_percent` - The size of additional Flash Cache in percentage of High Capacity database storage.
				* `autoscale_limit_in_gbs` - The maximum limit, in gigabytes, to which the Vault storage size can automatically scale when auto scaling is enabled for the Database Storage Vault.
				* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for Db Storage Vault.
				* `db_storage_vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata Database Storage Vault.
				* `display_name` - The user-friendly name for the Exadata Database Storage Vault. The name does not need to be unique.
				* `high_capacity_database_storage` - Total storage capacity in GB for vault storage.
				* `is_autoscale_enabled` - Indicates if autoscale feature is enabled for the Database Storage Vault. The default value is `FALSE`.
				* `subscription_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription with which resource needs to be associated with.
			* `exascale_db_storage_vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the exascale db storage vault.
			* `kms_key_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations. 
			* `kms_key_version_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. 
			* `metadata` - Additional metadata related to Globally distributed database resources.
				* `map` - The map containing key-value pair of additional metadata.
			* `name` - The name of the catalog.
			* `protection_mode` - The protection mode used for the Data Guard association.
			* `source` - The source of the Globally distributed database type. Use XS_EXISTING_CLUSTER for a database on Exascale Infrastructure that uses a pre-existing cluster. Use XS_NEW_CLUSTER for a database on Exascale Infrastructure that requires a new cluster created from scratch. Use XS_NEW_VAULT_AND_CLUSTER for a database on Exascale Infrastructure that requires a new vault and cluster created from scratch. Use XD_EXISTING_CLUSTER for a database on Dedicated Infrastructure that uses a pre-existing cluster. Use EXISTING_DB_HOME for a database on a pre-existing database home. 
			* `status` - Status of the distributed database catalog.
			* `supporting_resource_id` - the identifier of the underlying supporting resource.
			* `time_created` - The time the catalog was created. An RFC3339 formatted datetime string.
			* `time_updated` - The time the catalog was last updated. An RFC3339 formatted datetime string.
			* `transport_type` - The transport type used for the Data Guard association.
			* `vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts).  This parameter and `kmsKeyId` are required for Customer Managed Keys. 
			* `vm_cluster_details` - The Exadata VM cluster for Distributed Database Resource.
				* `backup_network_nsg_ids` - A list of the [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that the backup network of this DB system belongs to.  Setting this to an empty array after the list is created removes the resource from all NSGs.  For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). Applicable only to Exadata systems. 
				* `backup_subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup network subnet associated with the Exadata VM cluster on Exascale Infrastructure. 
				* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for VM Cluster.
				* `display_name` - The user-friendly name for the Exadata VM cluster on Exascale Infrastructure. The name does not need to be unique.
				* `domain` - A domain name used for the Exadata VM cluster on Exascale Infrastructure.  If the Oracle-provided internet and VCN resolver is enabled for the specified subnet, then the domain name for the subnet is used (do not provide one). Otherwise, provide a valid DNS domain name.  Hyphens (-) are not permitted. Applies to Exadata Database Service on Exascale Infrastructure only. 
				* `enabled_ecpu_count` - The number of ECPUs to enable for an Exadata VM cluster on Exascale Infrastructure. 
				* `is_diagnostics_events_enabled` - Indicates whether diagnostic collection is enabled for the VM cluster.  Enabling diagnostic collection allows you to receive Events service notifications for guest VM issues.  Diagnostic collection also allows Oracle to provide enhanced service and proactive support for your Exadata system.  You can enable diagnostic collection during VM cluster provisioning. You can also disable or enable it at any time using the `UpdateVmCluster` API. 
				* `is_health_monitoring_enabled` - Indicates whether health monitoring is enabled for the VM cluster.  Enabling health monitoring allows Oracle to collect diagnostic data and share it with its operations and support personnel.  You may also receive notifications for some events. Collecting health diagnostics enables Oracle to provide proactive support and enhanced service for your system. Optionally enable health monitoring while provisioning a system. You can also disable or enable health monitoring anytime using the `UpdateVmCluster` API. 
				* `is_incident_logs_enabled` - Indicates whether incident logs and trace collection are enabled for the VM cluster.  Enabling incident logs collection allows Oracle to receive Events service notifications for guest VM issues, collect incident logs and traces, and use them to diagnose issues and resolve them. Optionally enable incident logs collection while provisioning a system. You can also disable or enable incident logs collection anytime using the `UpdateVmCluster` API. 
				* `license_model` - The Oracle license model that applies to the Exadata VM cluster on Exascale Infrastructure. The default is BRING_YOUR_OWN_LICENSE. 
				* `nsg_ids` - The list of [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the network security groups (NSGs) to which this resource belongs.  Setting this to an empty list removes all resources from all NSGs.  For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). 
				* `private_zone_id` - The private zone ID in which you want DNS records to be created. 
				* `ssh_public_keys` - The public key portion of one or more key pairs used for SSH access to the Exadata VM cluster on Exascale Infrastructure.
				* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet associated with the Exadata VM cluster on Exascale Infrastructure. 
				* `subscription_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription with which resource needs to be associated with.
				* `total_ecpu_count` - The number of Total ECPUs for an Exadata VM cluster on Exascale Infrastructure. 
				* `vm_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM cluster.
				* `vm_file_system_storage_size` - File System Storage Size in GBs for Exadata VM cluster. 
			* `vm_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Cloud VM Cluster.
		* `name` - The name of the catalog replica.  It must start with a letter and contain only letters, digits, and underscores.  Maximum length is 40 characters. 
* `character_set` - The character set for the database.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Globally distributed database compartment.
* `composite_data_guard_shard_spaces` - Collection of composite data guard shard spaces.
	* `chunk_count` - Count of chunks associated with the shard space.
	* `data_guard_replicas` - The details of data guard replica associated with the shard.
		* `databases` - The details of composite data guard shard space replica databases.
			* `availability_domain` - The name of the availability domain that the distributed database shard will be located in.
			* `container_database_id` - the identifier of the container database for underlying supporting resource.
			* `db_home_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Home.
			* `db_storage_vault_details` - The Storage Vault for Distributed Database Resource.
				* `additional_flash_cache_in_percent` - The size of additional Flash Cache in percentage of High Capacity database storage.
				* `autoscale_limit_in_gbs` - The maximum limit, in gigabytes, to which the Vault storage size can automatically scale when auto scaling is enabled for the Database Storage Vault.
				* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for Db Storage Vault.
				* `db_storage_vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata Database Storage Vault.
				* `display_name` - The user-friendly name for the Exadata Database Storage Vault. The name does not need to be unique.
				* `high_capacity_database_storage` - Total storage capacity in GB for vault storage.
				* `is_autoscale_enabled` - Indicates if autoscale feature is enabled for the Database Storage Vault. The default value is `FALSE`.
				* `subscription_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription with which resource needs to be associated with.
			* `exascale_db_storage_vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the exascale db storage vault.
			* `kms_key_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations. 
			* `kms_key_version_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. 
			* `metadata` - Additional metadata related to Globally distributed database resources.
				* `map` - The map containing key-value pair of additional metadata.
			* `name` - Name of the shard.
			* `protection_mode` - The protection mode used for the Data Guard association.
			* `source` - The source of the Globally distributed database type. Use XS_EXISTING_CLUSTER for a database on Exascale Infrastructure that uses a pre-existing cluster. Use XS_NEW_CLUSTER for a database on Exascale Infrastructure that requires a new cluster created from scratch. Use XS_NEW_VAULT_AND_CLUSTER for a database on Exascale Infrastructure that requires a new vault and cluster created from scratch. Use XD_EXISTING_CLUSTER for a database on Dedicated Infrastructure that uses a pre-existing cluster. Use EXISTING_DB_HOME for a database on a pre-existing database home. 
			* `status` - Status of distributed database shard.
			* `supporting_resource_id` - the identifier of the underlying supporting resource.
			* `time_created` - The time the shard was created. An RFC3339 formatted datetime string.
			* `time_updated` - The time the shard was last updated. An RFC3339 formatted datetime string.
			* `transport_type` - The transport type used for the Data Guard association.
			* `vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts).  This parameter and `kmsKeyId` are required for Customer Managed Keys. 
			* `vm_cluster_details` - The Exadata VM cluster for Distributed Database Resource.
				* `backup_network_nsg_ids` - A list of the [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that the backup network of this DB system belongs to.  Setting this to an empty array after the list is created removes the resource from all NSGs.  For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). Applicable only to Exadata systems. 
				* `backup_subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup network subnet associated with the Exadata VM cluster on Exascale Infrastructure. 
				* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for VM Cluster.
				* `display_name` - The user-friendly name for the Exadata VM cluster on Exascale Infrastructure. The name does not need to be unique.
				* `domain` - A domain name used for the Exadata VM cluster on Exascale Infrastructure.  If the Oracle-provided internet and VCN resolver is enabled for the specified subnet, then the domain name for the subnet is used (do not provide one). Otherwise, provide a valid DNS domain name.  Hyphens (-) are not permitted. Applies to Exadata Database Service on Exascale Infrastructure only. 
				* `enabled_ecpu_count` - The number of ECPUs to enable for an Exadata VM cluster on Exascale Infrastructure. 
				* `is_diagnostics_events_enabled` - Indicates whether diagnostic collection is enabled for the VM cluster.  Enabling diagnostic collection allows you to receive Events service notifications for guest VM issues.  Diagnostic collection also allows Oracle to provide enhanced service and proactive support for your Exadata system.  You can enable diagnostic collection during VM cluster provisioning. You can also disable or enable it at any time using the `UpdateVmCluster` API. 
				* `is_health_monitoring_enabled` - Indicates whether health monitoring is enabled for the VM cluster.  Enabling health monitoring allows Oracle to collect diagnostic data and share it with its operations and support personnel.  You may also receive notifications for some events. Collecting health diagnostics enables Oracle to provide proactive support and enhanced service for your system. Optionally enable health monitoring while provisioning a system. You can also disable or enable health monitoring anytime using the `UpdateVmCluster` API. 
				* `is_incident_logs_enabled` - Indicates whether incident logs and trace collection are enabled for the VM cluster.  Enabling incident logs collection allows Oracle to receive Events service notifications for guest VM issues, collect incident logs and traces, and use them to diagnose issues and resolve them. Optionally enable incident logs collection while provisioning a system. You can also disable or enable incident logs collection anytime using the `UpdateVmCluster` API. 
				* `license_model` - The Oracle license model that applies to the Exadata VM cluster on Exascale Infrastructure. The default is BRING_YOUR_OWN_LICENSE. 
				* `nsg_ids` - The list of [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the network security groups (NSGs) to which this resource belongs.  Setting this to an empty list removes all resources from all NSGs.  For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). 
				* `private_zone_id` - The private zone ID in which you want DNS records to be created. 
				* `ssh_public_keys` - The public key portion of one or more key pairs used for SSH access to the Exadata VM cluster on Exascale Infrastructure.
				* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet associated with the Exadata VM cluster on Exascale Infrastructure. 
				* `subscription_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription with which resource needs to be associated with.
				* `total_ecpu_count` - The number of Total ECPUs for an Exadata VM cluster on Exascale Infrastructure. 
				* `vm_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM cluster.
				* `vm_file_system_storage_size` - File System Storage Size in GBs for Exadata VM cluster. 
			* `vm_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Cloud VM Cluster.
		* `name` - The name of shard space replica.  It must start with a letter and contain only letters, digits, and underscores.  Maximum length is 40 characters. 
	* `name` - The name of shard space.  It must start with a letter and contain only letters, digits, and underscores.  Maximum length is 40 characters. 
	* `original_replica` - Globally distributed database composite data guard replica details.
		* `databases` - The details of composite data guard shard space replica databases.
			* `availability_domain` - The name of the availability domain that the distributed database shard will be located in.
			* `container_database_id` - the identifier of the container database for underlying supporting resource.
			* `db_home_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Home.
			* `db_storage_vault_details` - The Storage Vault for Distributed Database Resource.
				* `additional_flash_cache_in_percent` - The size of additional Flash Cache in percentage of High Capacity database storage.
				* `autoscale_limit_in_gbs` - The maximum limit, in gigabytes, to which the Vault storage size can automatically scale when auto scaling is enabled for the Database Storage Vault.
				* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for Db Storage Vault.
				* `db_storage_vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata Database Storage Vault.
				* `display_name` - The user-friendly name for the Exadata Database Storage Vault. The name does not need to be unique.
				* `high_capacity_database_storage` - Total storage capacity in GB for vault storage.
				* `is_autoscale_enabled` - Indicates if autoscale feature is enabled for the Database Storage Vault. The default value is `FALSE`.
				* `subscription_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription with which resource needs to be associated with.
			* `exascale_db_storage_vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the exascale db storage vault.
			* `kms_key_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations. 
			* `kms_key_version_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. 
			* `metadata` - Additional metadata related to Globally distributed database resources.
				* `map` - The map containing key-value pair of additional metadata.
			* `name` - Name of the shard.
			* `protection_mode` - The protection mode used for the Data Guard association.
			* `source` - The source of the Globally distributed database type. Use XS_EXISTING_CLUSTER for a database on Exascale Infrastructure that uses a pre-existing cluster. Use XS_NEW_CLUSTER for a database on Exascale Infrastructure that requires a new cluster created from scratch. Use XS_NEW_VAULT_AND_CLUSTER for a database on Exascale Infrastructure that requires a new vault and cluster created from scratch. Use XD_EXISTING_CLUSTER for a database on Dedicated Infrastructure that uses a pre-existing cluster. Use EXISTING_DB_HOME for a database on a pre-existing database home. 
			* `status` - Status of distributed database shard.
			* `supporting_resource_id` - the identifier of the underlying supporting resource.
			* `time_created` - The time the shard was created. An RFC3339 formatted datetime string.
			* `time_updated` - The time the shard was last updated. An RFC3339 formatted datetime string.
			* `transport_type` - The transport type used for the Data Guard association.
			* `vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts).  This parameter and `kmsKeyId` are required for Customer Managed Keys. 
			* `vm_cluster_details` - The Exadata VM cluster for Distributed Database Resource.
				* `backup_network_nsg_ids` - A list of the [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that the backup network of this DB system belongs to.  Setting this to an empty array after the list is created removes the resource from all NSGs.  For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). Applicable only to Exadata systems. 
				* `backup_subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup network subnet associated with the Exadata VM cluster on Exascale Infrastructure. 
				* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for VM Cluster.
				* `display_name` - The user-friendly name for the Exadata VM cluster on Exascale Infrastructure. The name does not need to be unique.
				* `domain` - A domain name used for the Exadata VM cluster on Exascale Infrastructure.  If the Oracle-provided internet and VCN resolver is enabled for the specified subnet, then the domain name for the subnet is used (do not provide one). Otherwise, provide a valid DNS domain name.  Hyphens (-) are not permitted. Applies to Exadata Database Service on Exascale Infrastructure only. 
				* `enabled_ecpu_count` - The number of ECPUs to enable for an Exadata VM cluster on Exascale Infrastructure. 
				* `is_diagnostics_events_enabled` - Indicates whether diagnostic collection is enabled for the VM cluster.  Enabling diagnostic collection allows you to receive Events service notifications for guest VM issues.  Diagnostic collection also allows Oracle to provide enhanced service and proactive support for your Exadata system.  You can enable diagnostic collection during VM cluster provisioning. You can also disable or enable it at any time using the `UpdateVmCluster` API. 
				* `is_health_monitoring_enabled` - Indicates whether health monitoring is enabled for the VM cluster.  Enabling health monitoring allows Oracle to collect diagnostic data and share it with its operations and support personnel.  You may also receive notifications for some events. Collecting health diagnostics enables Oracle to provide proactive support and enhanced service for your system. Optionally enable health monitoring while provisioning a system. You can also disable or enable health monitoring anytime using the `UpdateVmCluster` API. 
				* `is_incident_logs_enabled` - Indicates whether incident logs and trace collection are enabled for the VM cluster.  Enabling incident logs collection allows Oracle to receive Events service notifications for guest VM issues, collect incident logs and traces, and use them to diagnose issues and resolve them. Optionally enable incident logs collection while provisioning a system. You can also disable or enable incident logs collection anytime using the `UpdateVmCluster` API. 
				* `license_model` - The Oracle license model that applies to the Exadata VM cluster on Exascale Infrastructure. The default is BRING_YOUR_OWN_LICENSE. 
				* `nsg_ids` - The list of [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the network security groups (NSGs) to which this resource belongs.  Setting this to an empty list removes all resources from all NSGs.  For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). 
				* `private_zone_id` - The private zone ID in which you want DNS records to be created. 
				* `ssh_public_keys` - The public key portion of one or more key pairs used for SSH access to the Exadata VM cluster on Exascale Infrastructure.
				* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet associated with the Exadata VM cluster on Exascale Infrastructure. 
				* `subscription_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription with which resource needs to be associated with.
				* `total_ecpu_count` - The number of Total ECPUs for an Exadata VM cluster on Exascale Infrastructure. 
				* `vm_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM cluster.
				* `vm_file_system_storage_size` - File System Storage Size in GBs for Exadata VM cluster. 
			* `vm_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Cloud VM Cluster.
		* `name` - The name of shard space replica.  It must start with a letter and contain only letters, digits, and underscores.  Maximum length is 40 characters. 
* `composite_raft_shard_spaces` - Collection of composite raft shards.
	* `chunk_count` - Count of chunks associated with the shard space.
	* `name` - The name of shard space.  It must start with a letter and contain only letters, digits, and underscores.  Maximum length is 40 characters. 
	* `raft_clusters` - The details of raft clusters associated with the shard space.
		* `databases` - The details of databases associated with the composite shard space raft cluster.
			* `availability_domain` - The name of the availability domain that the distributed database shard will be located in.
			* `container_database_id` - the identifier of the container database for underlying supporting resource.
			* `db_home_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Home.
			* `db_storage_vault_details` - The Storage Vault for Distributed Database Resource.
				* `additional_flash_cache_in_percent` - The size of additional Flash Cache in percentage of High Capacity database storage.
				* `autoscale_limit_in_gbs` - The maximum limit, in gigabytes, to which the Vault storage size can automatically scale when auto scaling is enabled for the Database Storage Vault.
				* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for Db Storage Vault.
				* `db_storage_vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata Database Storage Vault.
				* `display_name` - The user-friendly name for the Exadata Database Storage Vault. The name does not need to be unique.
				* `high_capacity_database_storage` - Total storage capacity in GB for vault storage.
				* `is_autoscale_enabled` - Indicates if autoscale feature is enabled for the Database Storage Vault. The default value is `FALSE`.
				* `subscription_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription with which resource needs to be associated with.
			* `exascale_db_storage_vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the exascale db storage vault.
			* `kms_key_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations. 
			* `kms_key_version_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. 
			* `metadata` - Additional metadata related to Globally distributed database resources.
				* `map` - The map containing key-value pair of additional metadata.
			* `name` - Name of the shard.
			* `protection_mode` - The protection mode used for the Data Guard association.
			* `source` - The source of the Globally distributed database type. Use XS_EXISTING_CLUSTER for a database on Exascale Infrastructure that uses a pre-existing cluster. Use XS_NEW_CLUSTER for a database on Exascale Infrastructure that requires a new cluster created from scratch. Use XS_NEW_VAULT_AND_CLUSTER for a database on Exascale Infrastructure that requires a new vault and cluster created from scratch. Use XD_EXISTING_CLUSTER for a database on Dedicated Infrastructure that uses a pre-existing cluster. Use EXISTING_DB_HOME for a database on a pre-existing database home. 
			* `status` - Status of distributed database shard.
			* `supporting_resource_id` - the identifier of the underlying supporting resource.
			* `time_created` - The time the shard was created. An RFC3339 formatted datetime string.
			* `time_updated` - The time the shard was last updated. An RFC3339 formatted datetime string.
			* `transport_type` - The transport type used for the Data Guard association.
			* `vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts).  This parameter and `kmsKeyId` are required for Customer Managed Keys. 
			* `vm_cluster_details` - The Exadata VM cluster for Distributed Database Resource.
				* `backup_network_nsg_ids` - A list of the [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that the backup network of this DB system belongs to.  Setting this to an empty array after the list is created removes the resource from all NSGs.  For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). Applicable only to Exadata systems. 
				* `backup_subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup network subnet associated with the Exadata VM cluster on Exascale Infrastructure. 
				* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for VM Cluster.
				* `display_name` - The user-friendly name for the Exadata VM cluster on Exascale Infrastructure. The name does not need to be unique.
				* `domain` - A domain name used for the Exadata VM cluster on Exascale Infrastructure.  If the Oracle-provided internet and VCN resolver is enabled for the specified subnet, then the domain name for the subnet is used (do not provide one). Otherwise, provide a valid DNS domain name.  Hyphens (-) are not permitted. Applies to Exadata Database Service on Exascale Infrastructure only. 
				* `enabled_ecpu_count` - The number of ECPUs to enable for an Exadata VM cluster on Exascale Infrastructure. 
				* `is_diagnostics_events_enabled` - Indicates whether diagnostic collection is enabled for the VM cluster.  Enabling diagnostic collection allows you to receive Events service notifications for guest VM issues.  Diagnostic collection also allows Oracle to provide enhanced service and proactive support for your Exadata system.  You can enable diagnostic collection during VM cluster provisioning. You can also disable or enable it at any time using the `UpdateVmCluster` API. 
				* `is_health_monitoring_enabled` - Indicates whether health monitoring is enabled for the VM cluster.  Enabling health monitoring allows Oracle to collect diagnostic data and share it with its operations and support personnel.  You may also receive notifications for some events. Collecting health diagnostics enables Oracle to provide proactive support and enhanced service for your system. Optionally enable health monitoring while provisioning a system. You can also disable or enable health monitoring anytime using the `UpdateVmCluster` API. 
				* `is_incident_logs_enabled` - Indicates whether incident logs and trace collection are enabled for the VM cluster.  Enabling incident logs collection allows Oracle to receive Events service notifications for guest VM issues, collect incident logs and traces, and use them to diagnose issues and resolve them. Optionally enable incident logs collection while provisioning a system. You can also disable or enable incident logs collection anytime using the `UpdateVmCluster` API. 
				* `license_model` - The Oracle license model that applies to the Exadata VM cluster on Exascale Infrastructure. The default is BRING_YOUR_OWN_LICENSE. 
				* `nsg_ids` - The list of [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the network security groups (NSGs) to which this resource belongs.  Setting this to an empty list removes all resources from all NSGs.  For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). 
				* `private_zone_id` - The private zone ID in which you want DNS records to be created. 
				* `ssh_public_keys` - The public key portion of one or more key pairs used for SSH access to the Exadata VM cluster on Exascale Infrastructure.
				* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet associated with the Exadata VM cluster on Exascale Infrastructure. 
				* `subscription_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription with which resource needs to be associated with.
				* `total_ecpu_count` - The number of Total ECPUs for an Exadata VM cluster on Exascale Infrastructure. 
				* `vm_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM cluster.
				* `vm_file_system_storage_size` - File System Storage Size in GBs for Exadata VM cluster. 
			* `vm_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Cloud VM Cluster.
		* `name` - The name of raft cluster.  It must start with a letter and contain only letters, digits, and underscores.  Maximum length is 40 characters. 
		* `replication_factor` - Replication factor associated with the raft cluster.
		* `ru_mode` - Replication factor associated with the raft cluster.
		* `witness_count` - Witnesses RU count associated with the raft cluster.
	* `replication_unit_count` - Number of replication units in the shard space.
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
		* `remote_region` - The name of the remote region where the remote automatic incremental backups will be stored. For information about valid region names, see [Regions and Availability Domains](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/regions.htm). 
		* `type` - Type of the database backup destination.
	* `can_run_immediate_full_backup` - If set to true, configures automatic full backups in the local region (the region of the DB system) for the first backup run immediately.
	* `is_auto_backup_enabled` - If set to true, configures automatic backups. If you previously used RMAN or dbcli to configure backups and then you switch to using the Console or the API for backups, a new backup configuration is created and associated with your database. This means that you can no longer rely on your previously configured unmanaged backups to work.
	* `is_remote_backup_enabled` - If set to true, configures automatic incremental backups in the local region (the region of the DB system) and the remote region with a default frequency of 1 hour. If you previously used RMAN or dbcli to configure backups, using the Console or the API for manged backups creates a new backup configuration for your database. The new configuration replaces the configuration created with RMAN or dbcli. This means that you can no longer rely on your previously configured unmanaged backups to work. 
	* `recovery_window_in_days` - Number of days between the current and the earliest point of recoverability covered by automatic backups. This value applies to automatic backups only. After a new automatic backup has been created, Oracle removes old automatic backups that are created before the window. When the value is updated, it is applied to all existing automatic backups. 
	* `remote_region` - The name of the remote region where the remote automatic incremental backups will be stored. For information about valid region names, see [Regions and Availability Domains](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/regions.htm). 
* `db_deployment_type` - The distributed database deployment type. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - The display name of the Globally distributed database.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `gds_control_node_details` - Global Database Services Control(GDS CTL) instances associated with the distributed database.
	* `availability_domain` - The name of the availability domain where the Global Database Services Control instance is located in.
	* `gds_control_node_image_details` - The Global Database Services Control(GDS CTL) image details.
		* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Global Database Services Control(GDS CTL) software image.
		* `version` - The version associated with the image identified by id.
	* `metadata` - Additional metadata related to Globally distributed database resources.
		* `map` - The map containing key-value pair of additional metadata.
	* `name` - Name of the Global Database Services Control instance
	* `private_ip` - IP address of the Global Database Services Control instance.
	* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet associated with the Global Database Services Control instance. 
	* `supporting_resource_id` - The identifier of the underlying supporting resource.
	* `time_created` - The time the Global Database Services Control instance was created. An RFC3339 formatted datetime string.
	* `time_updated` - The time the Global Database Services Control instance was last updated. An RFC3339 formatted datetime string.
* `gsm_details` - Global Service Manager (GSM) instances associated with the distributed database.
	* `compute_count` - The compute count for the Global service manager instance.
	* `data_storage_size_in_gbs` - The data disk group size to be allocated in GBs for the Global service manager instance.
	* `gsm_image_details` - The Global service manager image details
		* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Global service manager software image.
		* `version` - The version associated with the image identified by id.
	* `metadata` - Additional metadata related to Globally distributed database resources.
		* `map` - The map containing key-value pair of additional metadata.
	* `name` - Name of the Global service manager instance
	* `status` - Status of the gsm.
	* `supporting_resource_id` - the identifier of the underlying supporting resource.
	* `time_created` - The time the Global service manager instance was created. An RFC3339 formatted datetime string.
	* `time_ssl_certificate_expires` - The time the ssl certificate associated with Global service manager expires. An RFC3339 formatted datetime string
	* `time_updated` - The time the Global service manager instance was last updated. An RFC3339 formatted datetime string.
* `gsm_ssh_public_key` - The SSH public key for Global service manager instances.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Globally distributed database.
* `latest_gsm_image_details` - The Global service manager image details
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Global service manager software image.
	* `version` - The version associated with the image identified by id.
* `lifecycle_details` - The lifecycleDetails for the Globally distributed database.
* `listener_port` - The Global service manager listener port number for the Globally distributed database.
* `listener_port_tls` - The TLS listener port number for Globally distributed database.
* `metadata` - Additional metadata related to Globally distributed database resources.
	* `map` - The map containing key-value pair of additional metadata.
* `ncharacter_set` - The national character set for the database.
* `notification_topic_ids` - The collection of [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the notification topics associated with the globally distributed database.
* `ons_port_local` - Ons local port number.
* `ons_port_remote` - Ons remote port number.
* `prefix` - Unique name prefix for the Globally distributed databases. Only alpha-numeric values are allowed. First character has to be a letter followed by any combination of letter and number. 
* `private_endpoint_ids` - The collection of [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the private endpoint associated with Globally distributed autonomous database.
* `replication_method` - The Replication method for Globally distributed database. Use RAFT for Raft based replication. With RAFT replication, shards cannot have peers details set on them. In case shards need to have peers, please do not set RAFT replicationMethod. For all non RAFT replication cases (with or without peers), please set replicationMethod as DG or do not set any value for replicationMethod. 
* `scan_listener_port` - The TCP SCAN listener port for database clusters with source type XS_NEW_VAULT_AND_CLUSTER or XS_NEW_CLUSTER.
* `sharding_method` - Sharding methods for the Globally distributed database.
* `state` - Lifecycle states for the Globally distributed database.
* `system_chunk_count` - Count of chunks associated with system raft clusters or system data guard databases.
* `system_data_guard_databases` - System data guard shard database details.
	* `data_guard_replicas` - The system data guard replica details.
		* `databases` - Details of the databases associated with the system data guard replica.
			* `availability_domain` - The name of the availability domain that the distributed database shard will be located in.
			* `container_database_id` - the identifier of the container database for underlying supporting resource.
			* `db_home_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Home.
			* `db_storage_vault_details` - The Storage Vault for Distributed Database Resource.
				* `additional_flash_cache_in_percent` - The size of additional Flash Cache in percentage of High Capacity database storage.
				* `autoscale_limit_in_gbs` - The maximum limit, in gigabytes, to which the Vault storage size can automatically scale when auto scaling is enabled for the Database Storage Vault.
				* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for Db Storage Vault.
				* `db_storage_vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata Database Storage Vault.
				* `display_name` - The user-friendly name for the Exadata Database Storage Vault. The name does not need to be unique.
				* `high_capacity_database_storage` - Total storage capacity in GB for vault storage.
				* `is_autoscale_enabled` - Indicates if autoscale feature is enabled for the Database Storage Vault. The default value is `FALSE`.
				* `subscription_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription with which resource needs to be associated with.
			* `exascale_db_storage_vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the exascale db storage vault.
			* `kms_key_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations. 
			* `kms_key_version_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. 
			* `metadata` - Additional metadata related to Globally distributed database resources.
				* `map` - The map containing key-value pair of additional metadata.
			* `name` - Name of the shard.
			* `protection_mode` - The protection mode used for the Data Guard association.
			* `source` - The source of the Globally distributed database type. Use XS_EXISTING_CLUSTER for a database on Exascale Infrastructure that uses a pre-existing cluster. Use XS_NEW_CLUSTER for a database on Exascale Infrastructure that requires a new cluster created from scratch. Use XS_NEW_VAULT_AND_CLUSTER for a database on Exascale Infrastructure that requires a new vault and cluster created from scratch. Use XD_EXISTING_CLUSTER for a database on Dedicated Infrastructure that uses a pre-existing cluster. Use EXISTING_DB_HOME for a database on a pre-existing database home. 
			* `status` - Status of distributed database shard.
			* `supporting_resource_id` - the identifier of the underlying supporting resource.
			* `time_created` - The time the shard was created. An RFC3339 formatted datetime string.
			* `time_updated` - The time the shard was last updated. An RFC3339 formatted datetime string.
			* `transport_type` - The transport type used for the Data Guard association.
			* `vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts).  This parameter and `kmsKeyId` are required for Customer Managed Keys. 
			* `vm_cluster_details` - The Exadata VM cluster for Distributed Database Resource.
				* `backup_network_nsg_ids` - A list of the [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that the backup network of this DB system belongs to.  Setting this to an empty array after the list is created removes the resource from all NSGs.  For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). Applicable only to Exadata systems. 
				* `backup_subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup network subnet associated with the Exadata VM cluster on Exascale Infrastructure. 
				* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for VM Cluster.
				* `display_name` - The user-friendly name for the Exadata VM cluster on Exascale Infrastructure. The name does not need to be unique.
				* `domain` - A domain name used for the Exadata VM cluster on Exascale Infrastructure.  If the Oracle-provided internet and VCN resolver is enabled for the specified subnet, then the domain name for the subnet is used (do not provide one). Otherwise, provide a valid DNS domain name.  Hyphens (-) are not permitted. Applies to Exadata Database Service on Exascale Infrastructure only. 
				* `enabled_ecpu_count` - The number of ECPUs to enable for an Exadata VM cluster on Exascale Infrastructure. 
				* `is_diagnostics_events_enabled` - Indicates whether diagnostic collection is enabled for the VM cluster.  Enabling diagnostic collection allows you to receive Events service notifications for guest VM issues.  Diagnostic collection also allows Oracle to provide enhanced service and proactive support for your Exadata system.  You can enable diagnostic collection during VM cluster provisioning. You can also disable or enable it at any time using the `UpdateVmCluster` API. 
				* `is_health_monitoring_enabled` - Indicates whether health monitoring is enabled for the VM cluster.  Enabling health monitoring allows Oracle to collect diagnostic data and share it with its operations and support personnel.  You may also receive notifications for some events. Collecting health diagnostics enables Oracle to provide proactive support and enhanced service for your system. Optionally enable health monitoring while provisioning a system. You can also disable or enable health monitoring anytime using the `UpdateVmCluster` API. 
				* `is_incident_logs_enabled` - Indicates whether incident logs and trace collection are enabled for the VM cluster.  Enabling incident logs collection allows Oracle to receive Events service notifications for guest VM issues, collect incident logs and traces, and use them to diagnose issues and resolve them. Optionally enable incident logs collection while provisioning a system. You can also disable or enable incident logs collection anytime using the `UpdateVmCluster` API. 
				* `license_model` - The Oracle license model that applies to the Exadata VM cluster on Exascale Infrastructure. The default is BRING_YOUR_OWN_LICENSE. 
				* `nsg_ids` - The list of [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the network security groups (NSGs) to which this resource belongs.  Setting this to an empty list removes all resources from all NSGs.  For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). 
				* `private_zone_id` - The private zone ID in which you want DNS records to be created. 
				* `ssh_public_keys` - The public key portion of one or more key pairs used for SSH access to the Exadata VM cluster on Exascale Infrastructure.
				* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet associated with the Exadata VM cluster on Exascale Infrastructure. 
				* `subscription_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription with which resource needs to be associated with.
				* `total_ecpu_count` - The number of Total ECPUs for an Exadata VM cluster on Exascale Infrastructure. 
				* `vm_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM cluster.
				* `vm_file_system_storage_size` - File System Storage Size in GBs for Exadata VM cluster. 
			* `vm_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Cloud VM Cluster.
		* `name` - The name of data guard replica.  It must start with a letter and contain only letters, digits, and underscores.  Maximum length is 40 characters. 
	* `original_replica` - System data guard replica details.
		* `databases` - Details of the databases associated with the system data guard replica.
			* `availability_domain` - The name of the availability domain that the distributed database shard will be located in.
			* `container_database_id` - the identifier of the container database for underlying supporting resource.
			* `db_home_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Home.
			* `db_storage_vault_details` - The Storage Vault for Distributed Database Resource.
				* `additional_flash_cache_in_percent` - The size of additional Flash Cache in percentage of High Capacity database storage.
				* `autoscale_limit_in_gbs` - The maximum limit, in gigabytes, to which the Vault storage size can automatically scale when auto scaling is enabled for the Database Storage Vault.
				* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for Db Storage Vault.
				* `db_storage_vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata Database Storage Vault.
				* `display_name` - The user-friendly name for the Exadata Database Storage Vault. The name does not need to be unique.
				* `high_capacity_database_storage` - Total storage capacity in GB for vault storage.
				* `is_autoscale_enabled` - Indicates if autoscale feature is enabled for the Database Storage Vault. The default value is `FALSE`.
				* `subscription_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription with which resource needs to be associated with.
			* `exascale_db_storage_vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the exascale db storage vault.
			* `kms_key_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations. 
			* `kms_key_version_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. 
			* `metadata` - Additional metadata related to Globally distributed database resources.
				* `map` - The map containing key-value pair of additional metadata.
			* `name` - Name of the shard.
			* `protection_mode` - The protection mode used for the Data Guard association.
			* `source` - The source of the Globally distributed database type. Use XS_EXISTING_CLUSTER for a database on Exascale Infrastructure that uses a pre-existing cluster. Use XS_NEW_CLUSTER for a database on Exascale Infrastructure that requires a new cluster created from scratch. Use XS_NEW_VAULT_AND_CLUSTER for a database on Exascale Infrastructure that requires a new vault and cluster created from scratch. Use XD_EXISTING_CLUSTER for a database on Dedicated Infrastructure that uses a pre-existing cluster. Use EXISTING_DB_HOME for a database on a pre-existing database home. 
			* `status` - Status of distributed database shard.
			* `supporting_resource_id` - the identifier of the underlying supporting resource.
			* `time_created` - The time the shard was created. An RFC3339 formatted datetime string.
			* `time_updated` - The time the shard was last updated. An RFC3339 formatted datetime string.
			* `transport_type` - The transport type used for the Data Guard association.
			* `vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts).  This parameter and `kmsKeyId` are required for Customer Managed Keys. 
			* `vm_cluster_details` - The Exadata VM cluster for Distributed Database Resource.
				* `backup_network_nsg_ids` - A list of the [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that the backup network of this DB system belongs to.  Setting this to an empty array after the list is created removes the resource from all NSGs.  For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). Applicable only to Exadata systems. 
				* `backup_subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup network subnet associated with the Exadata VM cluster on Exascale Infrastructure. 
				* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for VM Cluster.
				* `display_name` - The user-friendly name for the Exadata VM cluster on Exascale Infrastructure. The name does not need to be unique.
				* `domain` - A domain name used for the Exadata VM cluster on Exascale Infrastructure.  If the Oracle-provided internet and VCN resolver is enabled for the specified subnet, then the domain name for the subnet is used (do not provide one). Otherwise, provide a valid DNS domain name.  Hyphens (-) are not permitted. Applies to Exadata Database Service on Exascale Infrastructure only. 
				* `enabled_ecpu_count` - The number of ECPUs to enable for an Exadata VM cluster on Exascale Infrastructure. 
				* `is_diagnostics_events_enabled` - Indicates whether diagnostic collection is enabled for the VM cluster.  Enabling diagnostic collection allows you to receive Events service notifications for guest VM issues.  Diagnostic collection also allows Oracle to provide enhanced service and proactive support for your Exadata system.  You can enable diagnostic collection during VM cluster provisioning. You can also disable or enable it at any time using the `UpdateVmCluster` API. 
				* `is_health_monitoring_enabled` - Indicates whether health monitoring is enabled for the VM cluster.  Enabling health monitoring allows Oracle to collect diagnostic data and share it with its operations and support personnel.  You may also receive notifications for some events. Collecting health diagnostics enables Oracle to provide proactive support and enhanced service for your system. Optionally enable health monitoring while provisioning a system. You can also disable or enable health monitoring anytime using the `UpdateVmCluster` API. 
				* `is_incident_logs_enabled` - Indicates whether incident logs and trace collection are enabled for the VM cluster.  Enabling incident logs collection allows Oracle to receive Events service notifications for guest VM issues, collect incident logs and traces, and use them to diagnose issues and resolve them. Optionally enable incident logs collection while provisioning a system. You can also disable or enable incident logs collection anytime using the `UpdateVmCluster` API. 
				* `license_model` - The Oracle license model that applies to the Exadata VM cluster on Exascale Infrastructure. The default is BRING_YOUR_OWN_LICENSE. 
				* `nsg_ids` - The list of [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the network security groups (NSGs) to which this resource belongs.  Setting this to an empty list removes all resources from all NSGs.  For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). 
				* `private_zone_id` - The private zone ID in which you want DNS records to be created. 
				* `ssh_public_keys` - The public key portion of one or more key pairs used for SSH access to the Exadata VM cluster on Exascale Infrastructure.
				* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet associated with the Exadata VM cluster on Exascale Infrastructure. 
				* `subscription_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription with which resource needs to be associated with.
				* `total_ecpu_count` - The number of Total ECPUs for an Exadata VM cluster on Exascale Infrastructure. 
				* `vm_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM cluster.
				* `vm_file_system_storage_size` - File System Storage Size in GBs for Exadata VM cluster. 
			* `vm_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Cloud VM Cluster.
		* `name` - The name of data guard replica.  It must start with a letter and contain only letters, digits, and underscores.  Maximum length is 40 characters. 
* `system_raft_clusters` - Collection of system raft clusters.
	* `databases` - Details of the databases associated with the system raft cluster.
		* `availability_domain` - The name of the availability domain that the distributed database shard will be located in.
		* `container_database_id` - the identifier of the container database for underlying supporting resource.
		* `db_home_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Home.
		* `db_storage_vault_details` - The Storage Vault for Distributed Database Resource.
			* `additional_flash_cache_in_percent` - The size of additional Flash Cache in percentage of High Capacity database storage.
			* `autoscale_limit_in_gbs` - The maximum limit, in gigabytes, to which the Vault storage size can automatically scale when auto scaling is enabled for the Database Storage Vault.
			* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for Db Storage Vault.
			* `db_storage_vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata Database Storage Vault.
			* `display_name` - The user-friendly name for the Exadata Database Storage Vault. The name does not need to be unique.
			* `high_capacity_database_storage` - Total storage capacity in GB for vault storage.
			* `is_autoscale_enabled` - Indicates if autoscale feature is enabled for the Database Storage Vault. The default value is `FALSE`.
			* `subscription_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription with which resource needs to be associated with.
		* `exascale_db_storage_vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the exascale db storage vault.
		* `kms_key_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations. 
		* `kms_key_version_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. 
		* `metadata` - Additional metadata related to Globally distributed database resources.
			* `map` - The map containing key-value pair of additional metadata.
		* `name` - Name of the shard.
		* `protection_mode` - The protection mode used for the Data Guard association.
		* `source` - The source of the Globally distributed database type. Use XS_EXISTING_CLUSTER for a database on Exascale Infrastructure that uses a pre-existing cluster. Use XS_NEW_CLUSTER for a database on Exascale Infrastructure that requires a new cluster created from scratch. Use XS_NEW_VAULT_AND_CLUSTER for a database on Exascale Infrastructure that requires a new vault and cluster created from scratch. Use XD_EXISTING_CLUSTER for a database on Dedicated Infrastructure that uses a pre-existing cluster. Use EXISTING_DB_HOME for a database on a pre-existing database home. 
		* `status` - Status of distributed database shard.
		* `supporting_resource_id` - the identifier of the underlying supporting resource.
		* `time_created` - The time the shard was created. An RFC3339 formatted datetime string.
		* `time_updated` - The time the shard was last updated. An RFC3339 formatted datetime string.
		* `transport_type` - The transport type used for the Data Guard association.
		* `vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts).  This parameter and `kmsKeyId` are required for Customer Managed Keys. 
		* `vm_cluster_details` - The Exadata VM cluster for Distributed Database Resource.
			* `backup_network_nsg_ids` - A list of the [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that the backup network of this DB system belongs to.  Setting this to an empty array after the list is created removes the resource from all NSGs.  For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). Applicable only to Exadata systems. 
			* `backup_subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup network subnet associated with the Exadata VM cluster on Exascale Infrastructure. 
			* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for VM Cluster.
			* `display_name` - The user-friendly name for the Exadata VM cluster on Exascale Infrastructure. The name does not need to be unique.
			* `domain` - A domain name used for the Exadata VM cluster on Exascale Infrastructure.  If the Oracle-provided internet and VCN resolver is enabled for the specified subnet, then the domain name for the subnet is used (do not provide one). Otherwise, provide a valid DNS domain name.  Hyphens (-) are not permitted. Applies to Exadata Database Service on Exascale Infrastructure only. 
			* `enabled_ecpu_count` - The number of ECPUs to enable for an Exadata VM cluster on Exascale Infrastructure. 
			* `is_diagnostics_events_enabled` - Indicates whether diagnostic collection is enabled for the VM cluster.  Enabling diagnostic collection allows you to receive Events service notifications for guest VM issues.  Diagnostic collection also allows Oracle to provide enhanced service and proactive support for your Exadata system.  You can enable diagnostic collection during VM cluster provisioning. You can also disable or enable it at any time using the `UpdateVmCluster` API. 
			* `is_health_monitoring_enabled` - Indicates whether health monitoring is enabled for the VM cluster.  Enabling health monitoring allows Oracle to collect diagnostic data and share it with its operations and support personnel.  You may also receive notifications for some events. Collecting health diagnostics enables Oracle to provide proactive support and enhanced service for your system. Optionally enable health monitoring while provisioning a system. You can also disable or enable health monitoring anytime using the `UpdateVmCluster` API. 
			* `is_incident_logs_enabled` - Indicates whether incident logs and trace collection are enabled for the VM cluster.  Enabling incident logs collection allows Oracle to receive Events service notifications for guest VM issues, collect incident logs and traces, and use them to diagnose issues and resolve them. Optionally enable incident logs collection while provisioning a system. You can also disable or enable incident logs collection anytime using the `UpdateVmCluster` API. 
			* `license_model` - The Oracle license model that applies to the Exadata VM cluster on Exascale Infrastructure. The default is BRING_YOUR_OWN_LICENSE. 
			* `nsg_ids` - The list of [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the network security groups (NSGs) to which this resource belongs.  Setting this to an empty list removes all resources from all NSGs.  For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). 
			* `private_zone_id` - The private zone ID in which you want DNS records to be created. 
			* `ssh_public_keys` - The public key portion of one or more key pairs used for SSH access to the Exadata VM cluster on Exascale Infrastructure.
			* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet associated with the Exadata VM cluster on Exascale Infrastructure. 
			* `subscription_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription with which resource needs to be associated with.
			* `total_ecpu_count` - The number of Total ECPUs for an Exadata VM cluster on Exascale Infrastructure. 
			* `vm_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM cluster.
			* `vm_file_system_storage_size` - File System Storage Size in GBs for Exadata VM cluster. 
		* `vm_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Cloud VM Cluster.
	* `name` - The name of raft cluster.  It must start with a letter and contain only letters, digits, and underscores.  Maximum length is 40 characters. 
	* `replication_factor` - Replication factor associated with the raft cluster.
	* `ru_mode` - Replication factor associated with the raft cluster.
	* `witness_count` - Witnesses RU count associated with the raft cluster.
* `system_raft_replication_unit_count` - Number of replication units associated with system raft clusters.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the Globally distributed database was created. An RFC3339 formatted datetime string.
* `time_updated` - The time the Globally distributed database was last updated. An RFC3339 formatted datetime string.
* `user_shard_spaces` - Collection of user defined shard spaces.
	* `data_guard_replicas` - The details of data guard replica for the user shard space.
		* `database` - Details of a distributed database shard.
			* `availability_domain` - The name of the availability domain that the distributed database shard will be located in.
			* `container_database_id` - the identifier of the container database for underlying supporting resource.
			* `db_home_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Home.
			* `db_storage_vault_details` - The Storage Vault for Distributed Database Resource.
				* `additional_flash_cache_in_percent` - The size of additional Flash Cache in percentage of High Capacity database storage.
				* `autoscale_limit_in_gbs` - The maximum limit, in gigabytes, to which the Vault storage size can automatically scale when auto scaling is enabled for the Database Storage Vault.
				* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for Db Storage Vault.
				* `db_storage_vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata Database Storage Vault.
				* `display_name` - The user-friendly name for the Exadata Database Storage Vault. The name does not need to be unique.
				* `high_capacity_database_storage` - Total storage capacity in GB for vault storage.
				* `is_autoscale_enabled` - Indicates if autoscale feature is enabled for the Database Storage Vault. The default value is `FALSE`.
				* `subscription_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription with which resource needs to be associated with.
			* `exascale_db_storage_vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the exascale db storage vault.
			* `kms_key_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations. 
			* `kms_key_version_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. 
			* `metadata` - Additional metadata related to Globally distributed database resources.
				* `map` - The map containing key-value pair of additional metadata.
			* `name` - Name of the shard.
			* `protection_mode` - The protection mode used for the Data Guard association.
			* `source` - The source of the Globally distributed database type. Use XS_EXISTING_CLUSTER for a database on Exascale Infrastructure that uses a pre-existing cluster. Use XS_NEW_CLUSTER for a database on Exascale Infrastructure that requires a new cluster created from scratch. Use XS_NEW_VAULT_AND_CLUSTER for a database on Exascale Infrastructure that requires a new vault and cluster created from scratch. Use XD_EXISTING_CLUSTER for a database on Dedicated Infrastructure that uses a pre-existing cluster. Use EXISTING_DB_HOME for a database on a pre-existing database home. 
			* `status` - Status of distributed database shard.
			* `supporting_resource_id` - the identifier of the underlying supporting resource.
			* `time_created` - The time the shard was created. An RFC3339 formatted datetime string.
			* `time_updated` - The time the shard was last updated. An RFC3339 formatted datetime string.
			* `transport_type` - The transport type used for the Data Guard association.
			* `vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts).  This parameter and `kmsKeyId` are required for Customer Managed Keys. 
			* `vm_cluster_details` - The Exadata VM cluster for Distributed Database Resource.
				* `backup_network_nsg_ids` - A list of the [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that the backup network of this DB system belongs to.  Setting this to an empty array after the list is created removes the resource from all NSGs.  For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). Applicable only to Exadata systems. 
				* `backup_subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup network subnet associated with the Exadata VM cluster on Exascale Infrastructure. 
				* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for VM Cluster.
				* `display_name` - The user-friendly name for the Exadata VM cluster on Exascale Infrastructure. The name does not need to be unique.
				* `domain` - A domain name used for the Exadata VM cluster on Exascale Infrastructure.  If the Oracle-provided internet and VCN resolver is enabled for the specified subnet, then the domain name for the subnet is used (do not provide one). Otherwise, provide a valid DNS domain name.  Hyphens (-) are not permitted. Applies to Exadata Database Service on Exascale Infrastructure only. 
				* `enabled_ecpu_count` - The number of ECPUs to enable for an Exadata VM cluster on Exascale Infrastructure. 
				* `is_diagnostics_events_enabled` - Indicates whether diagnostic collection is enabled for the VM cluster.  Enabling diagnostic collection allows you to receive Events service notifications for guest VM issues.  Diagnostic collection also allows Oracle to provide enhanced service and proactive support for your Exadata system.  You can enable diagnostic collection during VM cluster provisioning. You can also disable or enable it at any time using the `UpdateVmCluster` API. 
				* `is_health_monitoring_enabled` - Indicates whether health monitoring is enabled for the VM cluster.  Enabling health monitoring allows Oracle to collect diagnostic data and share it with its operations and support personnel.  You may also receive notifications for some events. Collecting health diagnostics enables Oracle to provide proactive support and enhanced service for your system. Optionally enable health monitoring while provisioning a system. You can also disable or enable health monitoring anytime using the `UpdateVmCluster` API. 
				* `is_incident_logs_enabled` - Indicates whether incident logs and trace collection are enabled for the VM cluster.  Enabling incident logs collection allows Oracle to receive Events service notifications for guest VM issues, collect incident logs and traces, and use them to diagnose issues and resolve them. Optionally enable incident logs collection while provisioning a system. You can also disable or enable incident logs collection anytime using the `UpdateVmCluster` API. 
				* `license_model` - The Oracle license model that applies to the Exadata VM cluster on Exascale Infrastructure. The default is BRING_YOUR_OWN_LICENSE. 
				* `nsg_ids` - The list of [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the network security groups (NSGs) to which this resource belongs.  Setting this to an empty list removes all resources from all NSGs.  For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). 
				* `private_zone_id` - The private zone ID in which you want DNS records to be created. 
				* `ssh_public_keys` - The public key portion of one or more key pairs used for SSH access to the Exadata VM cluster on Exascale Infrastructure.
				* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet associated with the Exadata VM cluster on Exascale Infrastructure. 
				* `subscription_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription with which resource needs to be associated with.
				* `total_ecpu_count` - The number of Total ECPUs for an Exadata VM cluster on Exascale Infrastructure. 
				* `vm_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM cluster.
				* `vm_file_system_storage_size` - File System Storage Size in GBs for Exadata VM cluster. 
			* `vm_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Cloud VM Cluster.
		* `name` - The name of user shard space replica.  It must start with a letter and contain only letters, digits, and underscores.  Maximum length is 40 characters. 
	* `name` - The name of user shard space.  It must start with a letter and contain only letters, digits, and underscores.  Maximum length is 40 characters. 
	* `original_replica` - Details of  user shard space.
		* `database` - Details of a distributed database shard.
			* `availability_domain` - The name of the availability domain that the distributed database shard will be located in.
			* `container_database_id` - the identifier of the container database for underlying supporting resource.
			* `db_home_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Home.
			* `db_storage_vault_details` - The Storage Vault for Distributed Database Resource.
				* `additional_flash_cache_in_percent` - The size of additional Flash Cache in percentage of High Capacity database storage.
				* `autoscale_limit_in_gbs` - The maximum limit, in gigabytes, to which the Vault storage size can automatically scale when auto scaling is enabled for the Database Storage Vault.
				* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for Db Storage Vault.
				* `db_storage_vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata Database Storage Vault.
				* `display_name` - The user-friendly name for the Exadata Database Storage Vault. The name does not need to be unique.
				* `high_capacity_database_storage` - Total storage capacity in GB for vault storage.
				* `is_autoscale_enabled` - Indicates if autoscale feature is enabled for the Database Storage Vault. The default value is `FALSE`.
				* `subscription_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription with which resource needs to be associated with.
			* `exascale_db_storage_vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the exascale db storage vault.
			* `kms_key_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations. 
			* `kms_key_version_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. 
			* `metadata` - Additional metadata related to Globally distributed database resources.
				* `map` - The map containing key-value pair of additional metadata.
			* `name` - Name of the shard.
			* `protection_mode` - The protection mode used for the Data Guard association.
			* `source` - The source of the Globally distributed database type. Use XS_EXISTING_CLUSTER for a database on Exascale Infrastructure that uses a pre-existing cluster. Use XS_NEW_CLUSTER for a database on Exascale Infrastructure that requires a new cluster created from scratch. Use XS_NEW_VAULT_AND_CLUSTER for a database on Exascale Infrastructure that requires a new vault and cluster created from scratch. Use XD_EXISTING_CLUSTER for a database on Dedicated Infrastructure that uses a pre-existing cluster. Use EXISTING_DB_HOME for a database on a pre-existing database home. 
			* `status` - Status of distributed database shard.
			* `supporting_resource_id` - the identifier of the underlying supporting resource.
			* `time_created` - The time the shard was created. An RFC3339 formatted datetime string.
			* `time_updated` - The time the shard was last updated. An RFC3339 formatted datetime string.
			* `transport_type` - The transport type used for the Data Guard association.
			* `vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts).  This parameter and `kmsKeyId` are required for Customer Managed Keys. 
			* `vm_cluster_details` - The Exadata VM cluster for Distributed Database Resource.
				* `backup_network_nsg_ids` - A list of the [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that the backup network of this DB system belongs to.  Setting this to an empty array after the list is created removes the resource from all NSGs.  For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). Applicable only to Exadata systems. 
				* `backup_subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup network subnet associated with the Exadata VM cluster on Exascale Infrastructure. 
				* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for VM Cluster.
				* `display_name` - The user-friendly name for the Exadata VM cluster on Exascale Infrastructure. The name does not need to be unique.
				* `domain` - A domain name used for the Exadata VM cluster on Exascale Infrastructure.  If the Oracle-provided internet and VCN resolver is enabled for the specified subnet, then the domain name for the subnet is used (do not provide one). Otherwise, provide a valid DNS domain name.  Hyphens (-) are not permitted. Applies to Exadata Database Service on Exascale Infrastructure only. 
				* `enabled_ecpu_count` - The number of ECPUs to enable for an Exadata VM cluster on Exascale Infrastructure. 
				* `is_diagnostics_events_enabled` - Indicates whether diagnostic collection is enabled for the VM cluster.  Enabling diagnostic collection allows you to receive Events service notifications for guest VM issues.  Diagnostic collection also allows Oracle to provide enhanced service and proactive support for your Exadata system.  You can enable diagnostic collection during VM cluster provisioning. You can also disable or enable it at any time using the `UpdateVmCluster` API. 
				* `is_health_monitoring_enabled` - Indicates whether health monitoring is enabled for the VM cluster.  Enabling health monitoring allows Oracle to collect diagnostic data and share it with its operations and support personnel.  You may also receive notifications for some events. Collecting health diagnostics enables Oracle to provide proactive support and enhanced service for your system. Optionally enable health monitoring while provisioning a system. You can also disable or enable health monitoring anytime using the `UpdateVmCluster` API. 
				* `is_incident_logs_enabled` - Indicates whether incident logs and trace collection are enabled for the VM cluster.  Enabling incident logs collection allows Oracle to receive Events service notifications for guest VM issues, collect incident logs and traces, and use them to diagnose issues and resolve them. Optionally enable incident logs collection while provisioning a system. You can also disable or enable incident logs collection anytime using the `UpdateVmCluster` API. 
				* `license_model` - The Oracle license model that applies to the Exadata VM cluster on Exascale Infrastructure. The default is BRING_YOUR_OWN_LICENSE. 
				* `nsg_ids` - The list of [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the network security groups (NSGs) to which this resource belongs.  Setting this to an empty list removes all resources from all NSGs.  For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). 
				* `private_zone_id` - The private zone ID in which you want DNS records to be created. 
				* `ssh_public_keys` - The public key portion of one or more key pairs used for SSH access to the Exadata VM cluster on Exascale Infrastructure.
				* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet associated with the Exadata VM cluster on Exascale Infrastructure. 
				* `subscription_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription with which resource needs to be associated with.
				* `total_ecpu_count` - The number of Total ECPUs for an Exadata VM cluster on Exascale Infrastructure. 
				* `vm_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM cluster.
				* `vm_file_system_storage_size` - File System Storage Size in GBs for Exadata VM cluster. 
			* `vm_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Cloud VM Cluster.
		* `name` - The name of user shard space replica.  It must start with a letter and contain only letters, digits, and underscores.  Maximum length is 40 characters. 
* `vcn_nsg_ids` - The list of network security group (NSG) details associated with the distributed database.
	* `nsg_ids` - The list of [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) associated with resources in this VCN. For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). 
	* `vcn_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the customer Virtual Cloud Network.

