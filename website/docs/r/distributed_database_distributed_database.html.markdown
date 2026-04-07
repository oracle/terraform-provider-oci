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
For an EXADB_XS based distributed database, removing a shard with the parameter mustDeleteInfra set to true 
will also delete the associated VmCluster and DbStorageVault.


## Example Usage

```hcl
resource "oci_distributed_database_distributed_database" "test_distributed_database" {
	#Required
	catalog_details {
		#Required
		admin_password = var.distributed_database_catalog_details_admin_password
		source = var.distributed_database_catalog_details_source

		#Optional
		availability_domain = var.distributed_database_catalog_details_availability_domain
		db_storage_vault_details {

			#Optional
			additional_flash_cache_in_percent = var.distributed_database_catalog_details_db_storage_vault_details_additional_flash_cache_in_percent
			high_capacity_database_storage = var.distributed_database_catalog_details_db_storage_vault_details_high_capacity_database_storage
		}
		kms_key_id = oci_kms_key.test_key.id
		kms_key_version_id = oci_kms_key_version.test_key_version.id
		peer_details {

			#Optional
			availability_domain = var.distributed_database_catalog_details_peer_details_availability_domain
			db_storage_vault_details {

				#Optional
				additional_flash_cache_in_percent = var.distributed_database_catalog_details_peer_details_db_storage_vault_details_additional_flash_cache_in_percent
				high_capacity_database_storage = var.distributed_database_catalog_details_peer_details_db_storage_vault_details_high_capacity_database_storage
			}
			protection_mode = var.distributed_database_catalog_details_peer_details_protection_mode
			transport_type = var.distributed_database_catalog_details_peer_details_transport_type
			vm_cluster_details {

				#Optional
				backup_network_nsg_ids = var.distributed_database_catalog_details_peer_details_vm_cluster_details_backup_network_nsg_ids
				backup_subnet_id = oci_core_subnet.test_subnet.id
				domain = var.distributed_database_catalog_details_peer_details_vm_cluster_details_domain
				enabled_ecpu_count = var.distributed_database_catalog_details_peer_details_vm_cluster_details_enabled_ecpu_count
				is_diagnostics_events_enabled = var.distributed_database_catalog_details_peer_details_vm_cluster_details_is_diagnostics_events_enabled
				is_health_monitoring_enabled = var.distributed_database_catalog_details_peer_details_vm_cluster_details_is_health_monitoring_enabled
				is_incident_logs_enabled = var.distributed_database_catalog_details_peer_details_vm_cluster_details_is_incident_logs_enabled
				license_model = var.distributed_database_catalog_details_peer_details_vm_cluster_details_license_model
				nsg_ids = var.distributed_database_catalog_details_peer_details_vm_cluster_details_nsg_ids
				private_zone_id = oci_dns_zone.test_zone.id
				ssh_public_keys = var.distributed_database_catalog_details_peer_details_vm_cluster_details_ssh_public_keys
				subnet_id = oci_core_subnet.test_subnet.id
				total_ecpu_count = var.distributed_database_catalog_details_peer_details_vm_cluster_details_total_ecpu_count
				vm_file_system_storage_size = var.distributed_database_catalog_details_peer_details_vm_cluster_details_vm_file_system_storage_size
			}
			vm_cluster_id = oci_database_vm_cluster.test_vm_cluster.id
		}
		peer_vm_cluster_ids = var.distributed_database_catalog_details_peer_vm_cluster_ids
		shard_space = var.distributed_database_catalog_details_shard_space
		vault_id = oci_kms_vault.test_vault.id
		vm_cluster_details {

			#Optional
			backup_network_nsg_ids = var.distributed_database_catalog_details_vm_cluster_details_backup_network_nsg_ids
			backup_subnet_id = oci_core_subnet.test_subnet.id
			domain = var.distributed_database_catalog_details_vm_cluster_details_domain
			enabled_ecpu_count = var.distributed_database_catalog_details_vm_cluster_details_enabled_ecpu_count
			is_diagnostics_events_enabled = var.distributed_database_catalog_details_vm_cluster_details_is_diagnostics_events_enabled
			is_health_monitoring_enabled = var.distributed_database_catalog_details_vm_cluster_details_is_health_monitoring_enabled
			is_incident_logs_enabled = var.distributed_database_catalog_details_vm_cluster_details_is_incident_logs_enabled
			license_model = var.distributed_database_catalog_details_vm_cluster_details_license_model
			nsg_ids = var.distributed_database_catalog_details_vm_cluster_details_nsg_ids
			private_zone_id = oci_dns_zone.test_zone.id
			ssh_public_keys = var.distributed_database_catalog_details_vm_cluster_details_ssh_public_keys
			subnet_id = oci_core_subnet.test_subnet.id
			total_ecpu_count = var.distributed_database_catalog_details_vm_cluster_details_total_ecpu_count
			vm_file_system_storage_size = var.distributed_database_catalog_details_vm_cluster_details_vm_file_system_storage_size
		}
		vm_cluster_id = oci_database_vm_cluster.test_vm_cluster.id
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

		#Optional
		availability_domain = var.distributed_database_shard_details_availability_domain
		db_storage_vault_details {

			#Optional
			additional_flash_cache_in_percent = var.distributed_database_shard_details_db_storage_vault_details_additional_flash_cache_in_percent
			high_capacity_database_storage = var.distributed_database_shard_details_db_storage_vault_details_high_capacity_database_storage
		}
		kms_key_id = oci_kms_key.test_key.id
		kms_key_version_id = oci_kms_key_version.test_key_version.id
		peer_details {

			#Optional
			availability_domain = var.distributed_database_shard_details_peer_details_availability_domain
			db_storage_vault_details {

				#Optional
				additional_flash_cache_in_percent = var.distributed_database_shard_details_peer_details_db_storage_vault_details_additional_flash_cache_in_percent
				high_capacity_database_storage = var.distributed_database_shard_details_peer_details_db_storage_vault_details_high_capacity_database_storage
			}
			protection_mode = var.distributed_database_shard_details_peer_details_protection_mode
			transport_type = var.distributed_database_shard_details_peer_details_transport_type
			vm_cluster_details {

				#Optional
				backup_network_nsg_ids = var.distributed_database_shard_details_peer_details_vm_cluster_details_backup_network_nsg_ids
				backup_subnet_id = oci_core_subnet.test_subnet.id
				domain = var.distributed_database_shard_details_peer_details_vm_cluster_details_domain
				enabled_ecpu_count = var.distributed_database_shard_details_peer_details_vm_cluster_details_enabled_ecpu_count
				is_diagnostics_events_enabled = var.distributed_database_shard_details_peer_details_vm_cluster_details_is_diagnostics_events_enabled
				is_health_monitoring_enabled = var.distributed_database_shard_details_peer_details_vm_cluster_details_is_health_monitoring_enabled
				is_incident_logs_enabled = var.distributed_database_shard_details_peer_details_vm_cluster_details_is_incident_logs_enabled
				license_model = var.distributed_database_shard_details_peer_details_vm_cluster_details_license_model
				nsg_ids = var.distributed_database_shard_details_peer_details_vm_cluster_details_nsg_ids
				private_zone_id = oci_dns_zone.test_zone.id
				ssh_public_keys = var.distributed_database_shard_details_peer_details_vm_cluster_details_ssh_public_keys
				subnet_id = oci_core_subnet.test_subnet.id
				total_ecpu_count = var.distributed_database_shard_details_peer_details_vm_cluster_details_total_ecpu_count
				vm_file_system_storage_size = var.distributed_database_shard_details_peer_details_vm_cluster_details_vm_file_system_storage_size
			}
			vm_cluster_id = oci_database_vm_cluster.test_vm_cluster.id
		}
		peer_vm_cluster_ids = var.distributed_database_shard_details_peer_vm_cluster_ids
		shard_space = var.distributed_database_shard_details_shard_space
		vault_id = oci_kms_vault.test_vault.id
		vm_cluster_details {

			#Optional
			backup_network_nsg_ids = var.distributed_database_shard_details_vm_cluster_details_backup_network_nsg_ids
			backup_subnet_id = oci_core_subnet.test_subnet.id
			domain = var.distributed_database_shard_details_vm_cluster_details_domain
			enabled_ecpu_count = var.distributed_database_shard_details_vm_cluster_details_enabled_ecpu_count
			is_diagnostics_events_enabled = var.distributed_database_shard_details_vm_cluster_details_is_diagnostics_events_enabled
			is_health_monitoring_enabled = var.distributed_database_shard_details_vm_cluster_details_is_health_monitoring_enabled
			is_incident_logs_enabled = var.distributed_database_shard_details_vm_cluster_details_is_incident_logs_enabled
			license_model = var.distributed_database_shard_details_vm_cluster_details_license_model
			nsg_ids = var.distributed_database_shard_details_vm_cluster_details_nsg_ids
			private_zone_id = oci_dns_zone.test_zone.id
			ssh_public_keys = var.distributed_database_shard_details_vm_cluster_details_ssh_public_keys
			subnet_id = oci_core_subnet.test_subnet.id
			total_ecpu_count = var.distributed_database_shard_details_vm_cluster_details_total_ecpu_count
			vm_file_system_storage_size = var.distributed_database_shard_details_vm_cluster_details_vm_file_system_storage_size
		}
		vm_cluster_id = oci_database_vm_cluster.test_vm_cluster.id
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
	scan_listener_port = var.distributed_database_scan_listener_port
}
```

## Argument Reference

The following arguments are supported:

* `catalog_details` - (Required) Collection of catalog for the Globally distributed database.
	* `admin_password` - (Required) The admin password for the catalog associated with Globally distributed database.
	* `availability_domain` - (Required when source=NEW_VAULT_AND_CLUSTER) The name of the availability domain that the distributed database shard will be located in.
	* `db_storage_vault_details` - (Required when source=NEW_VAULT_AND_CLUSTER) Details of the request to create exascale db vault storage for shard or catalog of the distributed database. 
		* `additional_flash_cache_in_percent` - (Applicable when source=NEW_VAULT_AND_CLUSTER) The size of additional Flash Cache in percentage of High Capacity database storage.
		* `high_capacity_database_storage` - (Required when source=NEW_VAULT_AND_CLUSTER) Total storage capacity in GB for vault storage.
	* `kms_key_id` - (Optional) The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
	* `kms_key_version_id` - (Optional) The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. 
	* `peer_details` - (Optional) The details required for creation of the peer for the ExadbXs infrastructure based catalog.
		* `availability_domain` - (Required when source=NEW_VAULT_AND_CLUSTER) The name of the availability domain that the distributed database shard will be located in.
		* `db_storage_vault_details` - (Required when source=NEW_VAULT_AND_CLUSTER) Details of the request to create exascale db vault storage for shard or catalog of the distributed database. 
			* `additional_flash_cache_in_percent` - (Applicable when source=NEW_VAULT_AND_CLUSTER) The size of additional Flash Cache in percentage of High Capacity database storage.
			* `high_capacity_database_storage` - (Required when source=NEW_VAULT_AND_CLUSTER) Total storage capacity in GB for vault storage.
		* `protection_mode` - (Applicable when source=EXADB_XS | NEW_VAULT_AND_CLUSTER) The protectionMode for the catalog peer.
		* `transport_type` - (Applicable when source=EXADB_XS | NEW_VAULT_AND_CLUSTER) The redo transport type to use for this Data Guard association.
		* `vm_cluster_details` - (Required when source=NEW_VAULT_AND_CLUSTER) Details of the request to create exadb vm cluster for shard or catalog of the distributed database. 
			* `backup_network_nsg_ids` - (Applicable when source=NEW_VAULT_AND_CLUSTER) A list of the [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that the backup network of this DB system belongs to.  Setting this to an empty array after the list is created removes the resource from all NSGs.  For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). Applicable only to Exadata systems. 
			* `backup_subnet_id` - (Required when source=NEW_VAULT_AND_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup network subnet associated with the Exadata VM cluster on Exascale Infrastructure. 
			* `domain` - (Applicable when source=NEW_VAULT_AND_CLUSTER) A domain name used for the Exadata VM cluster on Exascale Infrastructure.  If the Oracle-provided internet and VCN resolver is enabled for the specified subnet, then the domain name for the subnet is used (do not provide one). Otherwise, provide a valid DNS domain name.  Hyphens (-) are not permitted. Applies to Exadata Database Service on Exascale Infrastructure only. 
			* `enabled_ecpu_count` - (Required when source=NEW_VAULT_AND_CLUSTER) The number of ECPUs to enable for an Exadata VM cluster on Exascale Infrastructure. 
			* `is_diagnostics_events_enabled` - (Applicable when source=NEW_VAULT_AND_CLUSTER) Indicates whether diagnostic collection is enabled for the VM cluster.  Enabling diagnostic collection allows you to receive Events service notifications for guest VM issues.  Diagnostic collection also allows Oracle to provide enhanced service and proactive support for your Exadata system.  You can enable diagnostic collection during VM cluster provisioning. You can also disable or enable it at any time using the `UpdateVmCluster` API. 
			* `is_health_monitoring_enabled` - (Applicable when source=NEW_VAULT_AND_CLUSTER) Indicates whether health monitoring is enabled for the VM cluster.  Enabling health monitoring allows Oracle to collect diagnostic data and share it with its operations and support personnel.  You may also receive notifications for some events. Collecting health diagnostics enables Oracle to provide proactive support and enhanced service for your system. Optionally enable health monitoring while provisioning a system. You can also disable or enable health monitoring anytime using the `UpdateVmCluster` API. 
			* `is_incident_logs_enabled` - (Applicable when source=NEW_VAULT_AND_CLUSTER) Indicates whether incident logs and trace collection are enabled for the VM cluster.  Enabling incident logs collection allows Oracle to receive Events service notifications for guest VM issues, collect incident logs and traces, and use them to diagnose issues and resolve them. Optionally enable incident logs collection while provisioning a system. You can also disable or enable incident logs collection anytime using the `UpdateVmCluster` API. 
			* `license_model` - (Applicable when source=NEW_VAULT_AND_CLUSTER) The Oracle license model that applies to the Exadata VM cluster on Exascale Infrastructure. The default is BRING_YOUR_OWN_LICENSE. 
			* `nsg_ids` - (Applicable when source=NEW_VAULT_AND_CLUSTER) The list of [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the network security groups (NSGs) to which this resource belongs.  Setting this to an empty list removes all resources from all NSGs.  For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). 
			* `private_zone_id` - (Applicable when source=NEW_VAULT_AND_CLUSTER) The private zone ID in which you want DNS records to be created. 
			* `ssh_public_keys` - (Required when source=NEW_VAULT_AND_CLUSTER) The public key portion of one or more key pairs used for SSH access to the Exadata VM cluster on Exascale Infrastructure.
			* `subnet_id` - (Required when source=NEW_VAULT_AND_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet associated with the Exadata VM cluster on Exascale Infrastructure. 
			* `total_ecpu_count` - (Applicable when source=NEW_VAULT_AND_CLUSTER) The number of Total ECPUs for an Exadata VM cluster on Exascale Infrastructure. 
			* `vm_file_system_storage_size` - (Applicable when source=NEW_VAULT_AND_CLUSTER) File System Storage Size in GBs for Exadata VM cluster. 
		* `vm_cluster_id` - (Required when source=EXADB_XS) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM Cluster for the catalog peer.
	* `peer_vm_cluster_ids` - (Applicable when source=EXADB_XS) This field is deprecated. This should not be used while creation of new distributed database. To set the peers on catalog of distributed database please use peerDetails. 
	* `shard_space` - (Optional) The shard space name for the Globally distributed database. Shard space for existing shard cannot be changed, once shard is created. Shard space name shall be used while creation of new shards. 
	* `source` - (Required) Type of Globally distributed database Shard or Catalog. Use NEW_VAULT_AND_CLUSTER for a Globally distributed database on Exascale with new vaults and clusters created from scratch. Use EXISTING_CLUSTER for a Globally distributed database on Exascale based on pre-existing clusters. EXADB_XS is currently the same as EXISTING_CLUSTER and will be deprecated after the deprecation cycle. 
	* `vault_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `kmsKeyId` are required for Customer Managed Keys.
	* `vm_cluster_details` - (Required when source=NEW_VAULT_AND_CLUSTER) Details of the request to create exadb vm cluster for shard or catalog of the distributed database. 
		* `backup_network_nsg_ids` - (Applicable when source=NEW_VAULT_AND_CLUSTER) A list of the [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that the backup network of this DB system belongs to.  Setting this to an empty array after the list is created removes the resource from all NSGs.  For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). Applicable only to Exadata systems. 
		* `backup_subnet_id` - (Required when source=NEW_VAULT_AND_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup network subnet associated with the Exadata VM cluster on Exascale Infrastructure. 
		* `domain` - (Applicable when source=NEW_VAULT_AND_CLUSTER) A domain name used for the Exadata VM cluster on Exascale Infrastructure.  If the Oracle-provided internet and VCN resolver is enabled for the specified subnet, then the domain name for the subnet is used (do not provide one). Otherwise, provide a valid DNS domain name.  Hyphens (-) are not permitted. Applies to Exadata Database Service on Exascale Infrastructure only. 
		* `enabled_ecpu_count` - (Required when source=NEW_VAULT_AND_CLUSTER) The number of ECPUs to enable for an Exadata VM cluster on Exascale Infrastructure. 
		* `is_diagnostics_events_enabled` - (Applicable when source=NEW_VAULT_AND_CLUSTER) Indicates whether diagnostic collection is enabled for the VM cluster.  Enabling diagnostic collection allows you to receive Events service notifications for guest VM issues.  Diagnostic collection also allows Oracle to provide enhanced service and proactive support for your Exadata system.  You can enable diagnostic collection during VM cluster provisioning. You can also disable or enable it at any time using the `UpdateVmCluster` API. 
		* `is_health_monitoring_enabled` - (Applicable when source=NEW_VAULT_AND_CLUSTER) Indicates whether health monitoring is enabled for the VM cluster.  Enabling health monitoring allows Oracle to collect diagnostic data and share it with its operations and support personnel.  You may also receive notifications for some events. Collecting health diagnostics enables Oracle to provide proactive support and enhanced service for your system. Optionally enable health monitoring while provisioning a system. You can also disable or enable health monitoring anytime using the `UpdateVmCluster` API. 
		* `is_incident_logs_enabled` - (Applicable when source=NEW_VAULT_AND_CLUSTER) Indicates whether incident logs and trace collection are enabled for the VM cluster.  Enabling incident logs collection allows Oracle to receive Events service notifications for guest VM issues, collect incident logs and traces, and use them to diagnose issues and resolve them. Optionally enable incident logs collection while provisioning a system. You can also disable or enable incident logs collection anytime using the `UpdateVmCluster` API. 
		* `license_model` - (Applicable when source=NEW_VAULT_AND_CLUSTER) The Oracle license model that applies to the Exadata VM cluster on Exascale Infrastructure. The default is BRING_YOUR_OWN_LICENSE. 
		* `nsg_ids` - (Applicable when source=NEW_VAULT_AND_CLUSTER) The list of [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the network security groups (NSGs) to which this resource belongs.  Setting this to an empty list removes all resources from all NSGs.  For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). 
		* `private_zone_id` - (Applicable when source=NEW_VAULT_AND_CLUSTER) The private zone ID in which you want DNS records to be created. 
		* `ssh_public_keys` - (Required when source=NEW_VAULT_AND_CLUSTER) The public key portion of one or more key pairs used for SSH access to the Exadata VM cluster on Exascale Infrastructure.
		* `subnet_id` - (Required when source=NEW_VAULT_AND_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet associated with the Exadata VM cluster on Exascale Infrastructure. 
		* `total_ecpu_count` - (Applicable when source=NEW_VAULT_AND_CLUSTER) The number of Total ECPUs for an Exadata VM cluster on Exascale Infrastructure. 
		* `vm_file_system_storage_size` - (Applicable when source=NEW_VAULT_AND_CLUSTER) File System Storage Size in GBs for Exadata VM cluster. 
	* `vm_cluster_id` - (Required when source=EXADB_XS) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VmCluster.
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
* `scan_listener_port` - (Optional) The TCP Single Client Access Name (SCAN) port for clusters created for Globally distributed database. The scanListenerPort number should only be provided if shard and catalog have source type NEW_VAULT_AND_CLUSTER. If shard and catalog have source type NEW_VAULT_AND_CLUSTER and scanListenerPort is not provided then the scanListenerPort will default to value 1521. 
* `shard_details` - (Required) Collection of shards for the Globally distributed database.
	* `admin_password` - (Required) The admin password for the shard associated with Globally distributed database.
	* `availability_domain` - (Required when source=NEW_VAULT_AND_CLUSTER) The name of the availability domain that the distributed database shard will be located in.
	* `db_storage_vault_details` - (Required when source=NEW_VAULT_AND_CLUSTER) Details of the request to create exascale db vault storage for shard or catalog of the distributed database. 
		* `additional_flash_cache_in_percent` - (Applicable when source=NEW_VAULT_AND_CLUSTER) The size of additional Flash Cache in percentage of High Capacity database storage.
		* `high_capacity_database_storage` - (Required when source=NEW_VAULT_AND_CLUSTER) Total storage capacity in GB for vault storage.
	* `kms_key_id` - (Optional) The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
	* `kms_key_version_id` - (Optional) The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. 
	* `peer_details` - (Optional) The details required for creation of the peer for the ExadbXs infrastructure based shard.
		* `availability_domain` - (Required when source=NEW_VAULT_AND_CLUSTER) The name of the availability domain that the distributed database shard will be located in.
		* `db_storage_vault_details` - (Required when source=NEW_VAULT_AND_CLUSTER) Details of the request to create exascale db vault storage for shard or catalog of the distributed database. 
			* `additional_flash_cache_in_percent` - (Applicable when source=NEW_VAULT_AND_CLUSTER) The size of additional Flash Cache in percentage of High Capacity database storage.
			* `high_capacity_database_storage` - (Required when source=NEW_VAULT_AND_CLUSTER) Total storage capacity in GB for vault storage.
		* `protection_mode` - (Applicable when source=EXADB_XS | NEW_VAULT_AND_CLUSTER) The protectionMode for the shard peer.
		* `transport_type` - (Applicable when source=EXADB_XS | NEW_VAULT_AND_CLUSTER) The redo transport type to use for this Data Guard association.
		* `vm_cluster_details` - (Required when source=NEW_VAULT_AND_CLUSTER) Details of the request to create exadb vm cluster for shard or catalog of the distributed database. 
			* `backup_network_nsg_ids` - (Applicable when source=NEW_VAULT_AND_CLUSTER) A list of the [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that the backup network of this DB system belongs to.  Setting this to an empty array after the list is created removes the resource from all NSGs.  For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). Applicable only to Exadata systems. 
			* `backup_subnet_id` - (Required when source=NEW_VAULT_AND_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup network subnet associated with the Exadata VM cluster on Exascale Infrastructure. 
			* `domain` - (Applicable when source=NEW_VAULT_AND_CLUSTER) A domain name used for the Exadata VM cluster on Exascale Infrastructure.  If the Oracle-provided internet and VCN resolver is enabled for the specified subnet, then the domain name for the subnet is used (do not provide one). Otherwise, provide a valid DNS domain name.  Hyphens (-) are not permitted. Applies to Exadata Database Service on Exascale Infrastructure only. 
			* `enabled_ecpu_count` - (Required when source=NEW_VAULT_AND_CLUSTER) The number of ECPUs to enable for an Exadata VM cluster on Exascale Infrastructure. 
			* `is_diagnostics_events_enabled` - (Applicable when source=NEW_VAULT_AND_CLUSTER) Indicates whether diagnostic collection is enabled for the VM cluster.  Enabling diagnostic collection allows you to receive Events service notifications for guest VM issues.  Diagnostic collection also allows Oracle to provide enhanced service and proactive support for your Exadata system.  You can enable diagnostic collection during VM cluster provisioning. You can also disable or enable it at any time using the `UpdateVmCluster` API. 
			* `is_health_monitoring_enabled` - (Applicable when source=NEW_VAULT_AND_CLUSTER) Indicates whether health monitoring is enabled for the VM cluster.  Enabling health monitoring allows Oracle to collect diagnostic data and share it with its operations and support personnel.  You may also receive notifications for some events. Collecting health diagnostics enables Oracle to provide proactive support and enhanced service for your system. Optionally enable health monitoring while provisioning a system. You can also disable or enable health monitoring anytime using the `UpdateVmCluster` API. 
			* `is_incident_logs_enabled` - (Applicable when source=NEW_VAULT_AND_CLUSTER) Indicates whether incident logs and trace collection are enabled for the VM cluster.  Enabling incident logs collection allows Oracle to receive Events service notifications for guest VM issues, collect incident logs and traces, and use them to diagnose issues and resolve them. Optionally enable incident logs collection while provisioning a system. You can also disable or enable incident logs collection anytime using the `UpdateVmCluster` API. 
			* `license_model` - (Applicable when source=NEW_VAULT_AND_CLUSTER) The Oracle license model that applies to the Exadata VM cluster on Exascale Infrastructure. The default is BRING_YOUR_OWN_LICENSE. 
			* `nsg_ids` - (Applicable when source=NEW_VAULT_AND_CLUSTER) The list of [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the network security groups (NSGs) to which this resource belongs.  Setting this to an empty list removes all resources from all NSGs.  For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). 
			* `private_zone_id` - (Applicable when source=NEW_VAULT_AND_CLUSTER) The private zone ID in which you want DNS records to be created. 
			* `ssh_public_keys` - (Required when source=NEW_VAULT_AND_CLUSTER) The public key portion of one or more key pairs used for SSH access to the Exadata VM cluster on Exascale Infrastructure.
			* `subnet_id` - (Required when source=NEW_VAULT_AND_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet associated with the Exadata VM cluster on Exascale Infrastructure. 
			* `total_ecpu_count` - (Applicable when source=NEW_VAULT_AND_CLUSTER) The number of Total ECPUs for an Exadata VM cluster on Exascale Infrastructure. 
			* `vm_file_system_storage_size` - (Applicable when source=NEW_VAULT_AND_CLUSTER) File System Storage Size in GBs for Exadata VM cluster. 
		* `vm_cluster_id` - (Required when source=EXADB_XS) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM Cluster for the shard peer.
	* `peer_vm_cluster_ids` - (Applicable when source=EXADB_XS) This field is deprecated. This should not be used while creation of new distributed database. To set the peers on new shards of distributed database please use peerDetails. 
	* `shard_space` - (Optional) The shard space name for the Globally distributed database. Shard space for existing shard cannot be changed, once shard is created. Shard space name shall be used while creation of new shards. 
	* `source` - (Required) Type of Globally distributed database Shard or Catalog. Use NEW_VAULT_AND_CLUSTER for a Globally distributed database on Exascale with new vaults and clusters created from scratch. Use EXISTING_CLUSTER for a Globally distributed database on Exascale based on pre-existing clusters. EXADB_XS is currently the same as EXISTING_CLUSTER and will be deprecated after the deprecation cycle. 
	* `vault_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `kmsKeyId` are required for Customer Managed Keys.
	* `vm_cluster_details` - (Required when source=NEW_VAULT_AND_CLUSTER) Details of the request to create exadb vm cluster for shard or catalog of the distributed database. 
		* `backup_network_nsg_ids` - (Applicable when source=NEW_VAULT_AND_CLUSTER) A list of the [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that the backup network of this DB system belongs to.  Setting this to an empty array after the list is created removes the resource from all NSGs.  For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). Applicable only to Exadata systems. 
		* `backup_subnet_id` - (Required when source=NEW_VAULT_AND_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup network subnet associated with the Exadata VM cluster on Exascale Infrastructure. 
		* `domain` - (Applicable when source=NEW_VAULT_AND_CLUSTER) A domain name used for the Exadata VM cluster on Exascale Infrastructure.  If the Oracle-provided internet and VCN resolver is enabled for the specified subnet, then the domain name for the subnet is used (do not provide one). Otherwise, provide a valid DNS domain name.  Hyphens (-) are not permitted. Applies to Exadata Database Service on Exascale Infrastructure only. 
		* `enabled_ecpu_count` - (Required when source=NEW_VAULT_AND_CLUSTER) The number of ECPUs to enable for an Exadata VM cluster on Exascale Infrastructure. 
		* `is_diagnostics_events_enabled` - (Applicable when source=NEW_VAULT_AND_CLUSTER) Indicates whether diagnostic collection is enabled for the VM cluster.  Enabling diagnostic collection allows you to receive Events service notifications for guest VM issues.  Diagnostic collection also allows Oracle to provide enhanced service and proactive support for your Exadata system.  You can enable diagnostic collection during VM cluster provisioning. You can also disable or enable it at any time using the `UpdateVmCluster` API. 
		* `is_health_monitoring_enabled` - (Applicable when source=NEW_VAULT_AND_CLUSTER) Indicates whether health monitoring is enabled for the VM cluster.  Enabling health monitoring allows Oracle to collect diagnostic data and share it with its operations and support personnel.  You may also receive notifications for some events. Collecting health diagnostics enables Oracle to provide proactive support and enhanced service for your system. Optionally enable health monitoring while provisioning a system. You can also disable or enable health monitoring anytime using the `UpdateVmCluster` API. 
		* `is_incident_logs_enabled` - (Applicable when source=NEW_VAULT_AND_CLUSTER) Indicates whether incident logs and trace collection are enabled for the VM cluster.  Enabling incident logs collection allows Oracle to receive Events service notifications for guest VM issues, collect incident logs and traces, and use them to diagnose issues and resolve them. Optionally enable incident logs collection while provisioning a system. You can also disable or enable incident logs collection anytime using the `UpdateVmCluster` API. 
		* `license_model` - (Applicable when source=NEW_VAULT_AND_CLUSTER) The Oracle license model that applies to the Exadata VM cluster on Exascale Infrastructure. The default is BRING_YOUR_OWN_LICENSE. 
		* `nsg_ids` - (Applicable when source=NEW_VAULT_AND_CLUSTER) The list of [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the network security groups (NSGs) to which this resource belongs.  Setting this to an empty list removes all resources from all NSGs.  For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). 
		* `private_zone_id` - (Applicable when source=NEW_VAULT_AND_CLUSTER) The private zone ID in which you want DNS records to be created. 
		* `ssh_public_keys` - (Required when source=NEW_VAULT_AND_CLUSTER) The public key portion of one or more key pairs used for SSH access to the Exadata VM cluster on Exascale Infrastructure.
		* `subnet_id` - (Required when source=NEW_VAULT_AND_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet associated with the Exadata VM cluster on Exascale Infrastructure. 
		* `total_ecpu_count` - (Applicable when source=NEW_VAULT_AND_CLUSTER) The number of Total ECPUs for an Exadata VM cluster on Exascale Infrastructure. 
		* `vm_file_system_storage_size` - (Applicable when source=NEW_VAULT_AND_CLUSTER) File System Storage Size in GBs for Exadata VM cluster. 
	* `vm_cluster_id` - (Required when source=EXADB_XS) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VmCluster.
* `sharding_method` - (Required) Sharding Methods for the Globally distributed database.
* `state` - (Optional) (Updatable) The target state for the Distributed Database. Could be set to `ACTIVE` or `INACTIVE`. 
* `change_db_backup_config_trigger` - (Optional) (Updatable) An optional property when incremented triggers Change Db Backup Config. Could be set to any integer value.
* `configure_sharding_trigger` - (Optional) (Updatable) An optional property when incremented triggers Configure Sharding. Could be set to any integer value.
* `download_gsm_certificate_signing_request_trigger` - (Optional) (Updatable) An optional property when incremented triggers Download Gsm Certificate Signing Request. Could be set to any integer value.
* `generate_gsm_certificate_signing_request_trigger` - (Optional) (Updatable) An optional property when incremented triggers Generate Gsm Certificate Signing Request. Could be set to any integer value.
* `generate_wallet_trigger` - (Optional) (Updatable) An optional property when incremented triggers Generate Wallet. Could be set to any integer value.
* `move_replication_unit_trigger` - (Optional) (Updatable) An optional property when incremented triggers Move Replication Unit. Could be set to any integer value.
* `recreate_failed_resource_trigger` - (Optional) (Updatable) An optional property when incremented triggers Recreate Failed Resource. Could be set to any integer value.
* `upload_signed_certificate_and_generate_wallet_trigger` - (Optional) (Updatable) An optional property when incremented triggers Upload Signed Certificate And Generate Wallet. Could be set to any integer value.
* `validate_network_trigger` - (Optional) (Updatable) An optional property when incremented triggers Validate Network. Could be set to any integer value.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `catalog_details` - Collection of catalogs associated with the Globally distributed database.
	* `availability_domain` - The name of the availability domain that the distributed database catalog will be located in.
	* `container_database_id` - the identifier of the container database for underlying supporting resource.
	* `db_home_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Home.
	* `db_storage_vault_details` - The Storage Vault for Distributed Database Resource
		* `additional_flash_cache_in_percent` - The size of additional Flash Cache in percentage of High Capacity database storage.
		* `db_storage_vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Vault Storage.
		* `display_name` - The user-friendly name for the Exadata Database Storage Vault. The name does not need to be unique.
		* `high_capacity_database_storage` - Total storage capacity in GB for vault storage.
	* `kms_key_id` - The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
	* `kms_key_version_id` - The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. 
	* `metadata` - Additional metadata related to Globally distributed database resources.
		* `map` - The map containing key-value pair of additional metadata.
	* `name` - The name of catalog.
	* `peer_details` - Peer details for the catalog.
		* `availability_domain` - The name of the availability domain that the distributed database shard will be located in.
		* `container_database_id` - the identifier of the container database for underlying supporting resource.
		* `db_storage_vault_details` - The Storage Vault for Distributed Database Resource
			* `additional_flash_cache_in_percent` - The size of additional Flash Cache in percentage of High Capacity database storage.
			* `db_storage_vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Vault Storage.
			* `display_name` - The user-friendly name for the Exadata Database Storage Vault. The name does not need to be unique.
			* `high_capacity_database_storage` - Total storage capacity in GB for vault storage.
		* `metadata` - Additional metadata related to Globally distributed database resources.
			* `map` - The map containing key-value pair of additional metadata.
		* `protection_mode` - The protectionMode for the catalog peer.
		* `shard_group` - The name of the shardGroup for the peer.
		* `status` - Status of EXADB_XS based catalog peer.
		* `supporting_resource_id` - the identifier of the underlying supporting resource.
		* `time_created` - The time the catalog peer was created. An RFC3339 formatted datetime string
		* `time_updated` - The time the catalog peer was last updated. An RFC3339 formatted datetime string
		* `transport_type` - The redo transport type to use for this Data Guard association.
		* `vm_cluster_details` - The Exadata VM cluster for Distributed Database Resource
			* `backup_network_nsg_ids` - A list of the [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that the backup network of this DB system belongs to.  Setting this to an empty array after the list is created removes the resource from all NSGs.  For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). Applicable only to Exadata systems. 
			* `backup_subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup network subnet associated with the Exadata VM cluster on Exascale Infrastructure. 
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
			* `total_ecpu_count` - The number of Total ECPUs for an Exadata VM cluster on Exascale Infrastructure. 
			* `vm_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata VM cluster on Exascale Infrastructure.
			* `vm_file_system_storage_size` - File System Storage Size in GBs for Exadata VM cluster. 
		* `vm_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VmCluster.
	* `shard_group` - The name of the shardGroup for the catalog.
	* `source` - Type of Globally distributed database Shard or Catalog. Use NEW_VAULT_AND_CLUSTER for a Globally distributed database on Exascale with new vaults and clusters created from scratch. Use EXISTING_CLUSTER for a Globally distributed database on Exascale based on pre-existing clusters. EXADB_XS is currently the same as EXISTING_CLUSTER and will be deprecated after the deprecation cycle. 
	* `status` - Status of EXADB_XS based catalog.
	* `supporting_resource_id` - the identifier of the underlying supporting resource.
	* `time_created` - The time the catalog was created. An RFC3339 formatted datetime string
	* `time_updated` - The time the catalog was last updated. An RFC3339 formatted datetime string
	* `vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `kmsKeyId` are required for Customer Managed Keys.
	* `vm_cluster_details` - The Exadata VM cluster for Distributed Database Resource
		* `backup_network_nsg_ids` - A list of the [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that the backup network of this DB system belongs to.  Setting this to an empty array after the list is created removes the resource from all NSGs.  For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). Applicable only to Exadata systems. 
		* `backup_subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup network subnet associated with the Exadata VM cluster on Exascale Infrastructure. 
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
		* `total_ecpu_count` - The number of Total ECPUs for an Exadata VM cluster on Exascale Infrastructure. 
		* `vm_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata VM cluster on Exascale Infrastructure.
		* `vm_file_system_storage_size` - File System Storage Size in GBs for Exadata VM cluster. 
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
* `scan_listener_port` - The TCP Single Client Access Name (SCAN) port for Globally distributed database clusters.
* `shard_details` - Collection of shards associated with the Globally distributed database.
	* `availability_domain` - The name of the availability domain that the distributed database shard will be located in.
	* `container_database_id` - the identifier of the container database for underlying supporting resource.
	* `db_home_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Home.
	* `db_storage_vault_details` - The Storage Vault for Distributed Database Resource
		* `additional_flash_cache_in_percent` - The size of additional Flash Cache in percentage of High Capacity database storage.
		* `db_storage_vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Vault Storage.
		* `display_name` - The user-friendly name for the Exadata Database Storage Vault. The name does not need to be unique.
		* `high_capacity_database_storage` - Total storage capacity in GB for vault storage.
	* `kms_key_id` - The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
	* `kms_key_version_id` - The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. 
	* `metadata` - Additional metadata related to Globally distributed database resources.
		* `map` - The map containing key-value pair of additional metadata.
	* `name` - Name of the shard.
	* `peer_details` - Peer details for the shard.
		* `availability_domain` - The name of the availability domain that the distributed database shard will be located in.
		* `container_database_id` - the identifier of the container database for underlying supporting resource.
		* `db_storage_vault_details` - The Storage Vault for Distributed Database Resource
			* `additional_flash_cache_in_percent` - The size of additional Flash Cache in percentage of High Capacity database storage.
			* `db_storage_vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Vault Storage.
			* `display_name` - The user-friendly name for the Exadata Database Storage Vault. The name does not need to be unique.
			* `high_capacity_database_storage` - Total storage capacity in GB for vault storage.
		* `metadata` - Additional metadata related to Globally distributed database resources.
			* `map` - The map containing key-value pair of additional metadata.
		* `protection_mode` - The protectionMode for the shard peer.
		* `shard_group` - The name of the shardGroup for the peer.
		* `status` - Status of EXADB_XS based shard peer.
		* `supporting_resource_id` - the identifier of the underlying supporting resource.
		* `time_created` - The time the shard peer was created. An RFC3339 formatted datetime string
		* `time_updated` - The time the shard peer was last updated. An RFC3339 formatted datetime string
		* `transport_type` - The redo transport type to use for this Data Guard association.
		* `vm_cluster_details` - The Exadata VM cluster for Distributed Database Resource
			* `backup_network_nsg_ids` - A list of the [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that the backup network of this DB system belongs to.  Setting this to an empty array after the list is created removes the resource from all NSGs.  For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). Applicable only to Exadata systems. 
			* `backup_subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup network subnet associated with the Exadata VM cluster on Exascale Infrastructure. 
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
			* `total_ecpu_count` - The number of Total ECPUs for an Exadata VM cluster on Exascale Infrastructure. 
			* `vm_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata VM cluster on Exascale Infrastructure.
			* `vm_file_system_storage_size` - File System Storage Size in GBs for Exadata VM cluster. 
		* `vm_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VmCluster.
	* `shard_group` - The name of the shardGroup for the shard.
	* `shard_space` - The shard space name for the Globally distributed database. Shard space for existing shard cannot be changed, once shard is created. Shard space name shall be used while creation of new shards. 
	* `source` - Type of Globally distributed database Shard or Catalog. Use NEW_VAULT_AND_CLUSTER for a Globally distributed database on Exascale with new vaults and clusters created from scratch. Use EXISTING_CLUSTER for a Globally distributed database on Exascale based on pre-existing clusters. EXADB_XS is currently the same as EXISTING_CLUSTER and will be deprecated after the deprecation cycle. 
	* `status` - Status of EXADB_XS based shard.
	* `supporting_resource_id` - the identifier of the underlying supporting resource.
	* `time_created` - The time the shard was created. An RFC3339 formatted datetime string
	* `time_updated` - The time the shard was last updated. An RFC3339 formatted datetime string
	* `vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `kmsKeyId` are required for Customer Managed Keys.
	* `vm_cluster_details` - The Exadata VM cluster for Distributed Database Resource
		* `backup_network_nsg_ids` - A list of the [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that the backup network of this DB system belongs to.  Setting this to an empty array after the list is created removes the resource from all NSGs.  For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). Applicable only to Exadata systems. 
		* `backup_subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup network subnet associated with the Exadata VM cluster on Exascale Infrastructure. 
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
		* `total_ecpu_count` - The number of Total ECPUs for an Exadata VM cluster on Exascale Infrastructure. 
		* `vm_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata VM cluster on Exascale Infrastructure.
		* `vm_file_system_storage_size` - File System Storage Size in GBs for Exadata VM cluster. 
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

