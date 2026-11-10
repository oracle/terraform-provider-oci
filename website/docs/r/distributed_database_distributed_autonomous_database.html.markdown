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
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/


Creates a Globally distributed autonomous database.


## Managing Topology Changes

The topology blocks on this resource describe its initial configuration. Use the separate action resources for additions and removals:

* [INSERT](distributed_database_distributed_autonomous_databases_insert.html)
* [REMOVE](distributed_database_distributed_autonomous_databases_remove.html)
* [INSERTDG](distributed_database_distributed_autonomous_databases_insert_dg.html)
* [REMOVEDG](distributed_database_distributed_autonomous_databases_remove_dg.html)

Each action instance submits one to three operations in a single request. Resource labels such as `insert` and `insert_2` are user-defined; each instance submits its own operations. Keep successful instances unchanged and add another instance for additional operations. Changed action lists submit the entire list again; unchanged instances are not replayed. Destroying an action only clears its state and does not execute an inverse operation. Action-resource changes do not replace or destroy this main database resource.

Refresh preserves the initial topology inputs and exposes the current remote topology through the computed `observed_*` attributes below. For an imported resource without an initial topology snapshot, refresh initializes the topology inputs from the remote database.

After adding shards, invoke [distributed_autonomous_databases configure sharding](distributed_database_distributed_autonomous_databases_configure_sharding.html) after the successful action. Use explicit dependencies to order separate action instances targeting the same database.

## Example Usage

```hcl
resource "oci_distributed_database_distributed_autonomous_database" "test_distributed_autonomous_database" {
	#Required
	catalog_details {
		#Required
		original_replica {
			#Required
			database {
				#Required
				cloud_autonomous_vm_cluster_id = oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster.id
				compute_count = var.distributed_autonomous_database_catalog_details_original_replica_database_compute_count
				data_storage_size_in_gbs = var.distributed_autonomous_database_catalog_details_original_replica_database_data_storage_size_in_gbs
				source = var.distributed_autonomous_database_catalog_details_original_replica_database_source

				#Optional
				admin_password = var.distributed_autonomous_database_catalog_details_original_replica_database_admin_password
				admin_password_secret_id = oci_vault_secret.test_secret.id
				admin_password_secret_version_number = var.distributed_autonomous_database_catalog_details_original_replica_database_admin_password_secret_version_number
				fast_start_fail_over_lag_limit_in_seconds = var.distributed_autonomous_database_catalog_details_original_replica_database_fast_start_fail_over_lag_limit_in_seconds
				is_auto_scaling_enabled = var.distributed_autonomous_database_catalog_details_original_replica_database_is_auto_scaling_enabled
				kms_key_id = oci_kms_key.test_key.id
				kms_key_version_id = oci_kms_key_version.test_key_version.id
				okv_end_point_group_name = oci_identity_group.test_group.name
				okv_key_store_id = oci_database_key_store.test_key_store.id
				protection_mode = var.distributed_autonomous_database_catalog_details_original_replica_database_protection_mode
				standby_maintenance_buffer_in_days = var.distributed_autonomous_database_catalog_details_original_replica_database_standby_maintenance_buffer_in_days
				vault_id = oci_kms_vault.test_vault.id
			}
			name = var.distributed_autonomous_database_catalog_details_original_replica_name
		}

		#Optional
		data_guard_replicas {
			#Required
			database {
				#Required
				cloud_autonomous_vm_cluster_id = oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster.id
				compute_count = var.distributed_autonomous_database_catalog_details_data_guard_replicas_database_compute_count
				data_storage_size_in_gbs = var.distributed_autonomous_database_catalog_details_data_guard_replicas_database_data_storage_size_in_gbs
				source = var.distributed_autonomous_database_catalog_details_data_guard_replicas_database_source

				#Optional
				admin_password = var.distributed_autonomous_database_catalog_details_data_guard_replicas_database_admin_password
				admin_password_secret_id = oci_vault_secret.test_secret.id
				admin_password_secret_version_number = var.distributed_autonomous_database_catalog_details_data_guard_replicas_database_admin_password_secret_version_number
				fast_start_fail_over_lag_limit_in_seconds = var.distributed_autonomous_database_catalog_details_data_guard_replicas_database_fast_start_fail_over_lag_limit_in_seconds
				is_auto_scaling_enabled = var.distributed_autonomous_database_catalog_details_data_guard_replicas_database_is_auto_scaling_enabled
				kms_key_id = oci_kms_key.test_key.id
				kms_key_version_id = oci_kms_key_version.test_key_version.id
				okv_end_point_group_name = oci_identity_group.test_group.name
				okv_key_store_id = oci_database_key_store.test_key_store.id
				protection_mode = var.distributed_autonomous_database_catalog_details_data_guard_replicas_database_protection_mode
				standby_maintenance_buffer_in_days = var.distributed_autonomous_database_catalog_details_data_guard_replicas_database_standby_maintenance_buffer_in_days
				vault_id = oci_kms_vault.test_vault.id
			}
			name = var.distributed_autonomous_database_catalog_details_data_guard_replicas_name
		}
	}
	character_set = var.distributed_autonomous_database_character_set
	compartment_id = var.compartment_id
	database_version = var.distributed_autonomous_database_database_version
	db_workload_type = var.distributed_autonomous_database_db_workload_type
	display_name = var.distributed_autonomous_database_display_name
	distributed_autonomous_database_id = var.distributed_autonomous_database_distributed_autonomous_database_id
	listener_port = var.distributed_autonomous_database_listener_port
	ncharacter_set = var.distributed_autonomous_database_ncharacter_set
	ons_port_local = var.distributed_autonomous_database_ons_port_local
	ons_port_remote = var.distributed_autonomous_database_ons_port_remote
	prefix = var.distributed_autonomous_database_prefix
	private_endpoint_ids = var.distributed_autonomous_database_private_endpoint_ids

	#Optional
	composite_data_guard_shard_spaces {
		#Required
		chunk_count = var.distributed_autonomous_database_composite_data_guard_shard_spaces_chunk_count
		name = var.distributed_autonomous_database_composite_data_guard_shard_spaces_name
		original_replica {
			#Required
			databases {
				#Required
				cloud_autonomous_vm_cluster_id = oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster.id
				compute_count = var.distributed_autonomous_database_composite_data_guard_shard_spaces_original_replica_databases_compute_count
				data_storage_size_in_gbs = var.distributed_autonomous_database_composite_data_guard_shard_spaces_original_replica_databases_data_storage_size_in_gbs
				source = var.distributed_autonomous_database_composite_data_guard_shard_spaces_original_replica_databases_source

				#Optional
				admin_password = var.distributed_autonomous_database_composite_data_guard_shard_spaces_original_replica_databases_admin_password
				admin_password_secret_id = oci_vault_secret.test_secret.id
				admin_password_secret_version_number = var.distributed_autonomous_database_composite_data_guard_shard_spaces_original_replica_databases_admin_password_secret_version_number
				fast_start_fail_over_lag_limit_in_seconds = var.distributed_autonomous_database_composite_data_guard_shard_spaces_original_replica_databases_fast_start_fail_over_lag_limit_in_seconds
				is_auto_scaling_enabled = var.distributed_autonomous_database_composite_data_guard_shard_spaces_original_replica_databases_is_auto_scaling_enabled
				kms_key_id = oci_kms_key.test_key.id
				kms_key_version_id = oci_kms_key_version.test_key_version.id
				okv_end_point_group_name = oci_identity_group.test_group.name
				okv_key_store_id = oci_database_key_store.test_key_store.id
				protection_mode = var.distributed_autonomous_database_composite_data_guard_shard_spaces_original_replica_databases_protection_mode
				shard_space_name = var.distributed_autonomous_database_composite_data_guard_shard_spaces_original_replica_databases_shard_space_name
				standby_maintenance_buffer_in_days = var.distributed_autonomous_database_composite_data_guard_shard_spaces_original_replica_databases_standby_maintenance_buffer_in_days
				vault_id = oci_kms_vault.test_vault.id
			}
			name = var.distributed_autonomous_database_composite_data_guard_shard_spaces_original_replica_name
		}

		#Optional
		data_guard_replicas {
			#Required
			databases {
				#Required
				cloud_autonomous_vm_cluster_id = oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster.id
				compute_count = var.distributed_autonomous_database_composite_data_guard_shard_spaces_data_guard_replicas_databases_compute_count
				data_storage_size_in_gbs = var.distributed_autonomous_database_composite_data_guard_shard_spaces_data_guard_replicas_databases_data_storage_size_in_gbs
				source = var.distributed_autonomous_database_composite_data_guard_shard_spaces_data_guard_replicas_databases_source

				#Optional
				admin_password = var.distributed_autonomous_database_composite_data_guard_shard_spaces_data_guard_replicas_databases_admin_password
				admin_password_secret_id = oci_vault_secret.test_secret.id
				admin_password_secret_version_number = var.distributed_autonomous_database_composite_data_guard_shard_spaces_data_guard_replicas_databases_admin_password_secret_version_number
				fast_start_fail_over_lag_limit_in_seconds = var.distributed_autonomous_database_composite_data_guard_shard_spaces_data_guard_replicas_databases_fast_start_fail_over_lag_limit_in_seconds
				is_auto_scaling_enabled = var.distributed_autonomous_database_composite_data_guard_shard_spaces_data_guard_replicas_databases_is_auto_scaling_enabled
				kms_key_id = oci_kms_key.test_key.id
				kms_key_version_id = oci_kms_key_version.test_key_version.id
				okv_end_point_group_name = oci_identity_group.test_group.name
				okv_key_store_id = oci_database_key_store.test_key_store.id
				protection_mode = var.distributed_autonomous_database_composite_data_guard_shard_spaces_data_guard_replicas_databases_protection_mode
				shard_space_name = var.distributed_autonomous_database_composite_data_guard_shard_spaces_data_guard_replicas_databases_shard_space_name
				standby_maintenance_buffer_in_days = var.distributed_autonomous_database_composite_data_guard_shard_spaces_data_guard_replicas_databases_standby_maintenance_buffer_in_days
				vault_id = oci_kms_vault.test_vault.id
			}
			name = var.distributed_autonomous_database_composite_data_guard_shard_spaces_data_guard_replicas_name
		}
	}
	composite_raft_shard_spaces {
		#Required
		chunk_count = var.distributed_autonomous_database_composite_raft_shard_spaces_chunk_count
		name = var.distributed_autonomous_database_composite_raft_shard_spaces_name
		raft_clusters {
			#Required
			databases {
				#Required
				cloud_autonomous_vm_cluster_id = oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster.id
				compute_count = var.distributed_autonomous_database_composite_raft_shard_spaces_raft_clusters_databases_compute_count
				data_storage_size_in_gbs = var.distributed_autonomous_database_composite_raft_shard_spaces_raft_clusters_databases_data_storage_size_in_gbs
				source = var.distributed_autonomous_database_composite_raft_shard_spaces_raft_clusters_databases_source

				#Optional
				admin_password = var.distributed_autonomous_database_composite_raft_shard_spaces_raft_clusters_databases_admin_password
				admin_password_secret_id = oci_vault_secret.test_secret.id
				admin_password_secret_version_number = var.distributed_autonomous_database_composite_raft_shard_spaces_raft_clusters_databases_admin_password_secret_version_number
				fast_start_fail_over_lag_limit_in_seconds = var.distributed_autonomous_database_composite_raft_shard_spaces_raft_clusters_databases_fast_start_fail_over_lag_limit_in_seconds
				is_auto_scaling_enabled = var.distributed_autonomous_database_composite_raft_shard_spaces_raft_clusters_databases_is_auto_scaling_enabled
				kms_key_id = oci_kms_key.test_key.id
				kms_key_version_id = oci_kms_key_version.test_key_version.id
				okv_end_point_group_name = oci_identity_group.test_group.name
				okv_key_store_id = oci_database_key_store.test_key_store.id
				protection_mode = var.distributed_autonomous_database_composite_raft_shard_spaces_raft_clusters_databases_protection_mode
				shard_space_name = var.distributed_autonomous_database_composite_raft_shard_spaces_raft_clusters_databases_shard_space_name
				standby_maintenance_buffer_in_days = var.distributed_autonomous_database_composite_raft_shard_spaces_raft_clusters_databases_standby_maintenance_buffer_in_days
				vault_id = oci_kms_vault.test_vault.id
			}
			name = var.distributed_autonomous_database_composite_raft_shard_spaces_raft_clusters_name
			replication_factor = var.distributed_autonomous_database_composite_raft_shard_spaces_raft_clusters_replication_factor
			ru_mode = var.distributed_autonomous_database_composite_raft_shard_spaces_raft_clusters_ru_mode
			witness_count = var.distributed_autonomous_database_composite_raft_shard_spaces_raft_clusters_witness_count
		}
		replication_unit_count = var.distributed_autonomous_database_composite_raft_shard_spaces_replication_unit_count
	}
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
		}
		recovery_window_in_days = var.distributed_autonomous_database_db_backup_config_recovery_window_in_days
	}
	defined_tags = {"foo-namespace.bar-key"= "value"}
	freeform_tags = {"bar-key"= "value"}
	listener_port_tls = var.distributed_autonomous_database_listener_port_tls
	notification_topic_ids = var.distributed_autonomous_database_notification_topic_ids
	system_chunk_count = var.distributed_autonomous_database_system_chunk_count
	system_data_guard_databases {
		#Required
		original_replica {
			#Required
			databases {
				#Required
				cloud_autonomous_vm_cluster_id = oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster.id
				compute_count = var.distributed_autonomous_database_system_data_guard_databases_original_replica_databases_compute_count
				data_storage_size_in_gbs = var.distributed_autonomous_database_system_data_guard_databases_original_replica_databases_data_storage_size_in_gbs
				source = var.distributed_autonomous_database_system_data_guard_databases_original_replica_databases_source

				#Optional
				admin_password = var.distributed_autonomous_database_system_data_guard_databases_original_replica_databases_admin_password
				admin_password_secret_id = oci_vault_secret.test_secret.id
				admin_password_secret_version_number = var.distributed_autonomous_database_system_data_guard_databases_original_replica_databases_admin_password_secret_version_number
				fast_start_fail_over_lag_limit_in_seconds = var.distributed_autonomous_database_system_data_guard_databases_original_replica_databases_fast_start_fail_over_lag_limit_in_seconds
				is_auto_scaling_enabled = var.distributed_autonomous_database_system_data_guard_databases_original_replica_databases_is_auto_scaling_enabled
				kms_key_id = oci_kms_key.test_key.id
				kms_key_version_id = oci_kms_key_version.test_key_version.id
				okv_end_point_group_name = oci_identity_group.test_group.name
				okv_key_store_id = oci_database_key_store.test_key_store.id
				protection_mode = var.distributed_autonomous_database_system_data_guard_databases_original_replica_databases_protection_mode
				shard_space_name = var.distributed_autonomous_database_system_data_guard_databases_original_replica_databases_shard_space_name
				standby_maintenance_buffer_in_days = var.distributed_autonomous_database_system_data_guard_databases_original_replica_databases_standby_maintenance_buffer_in_days
				vault_id = oci_kms_vault.test_vault.id
			}
			name = var.distributed_autonomous_database_system_data_guard_databases_original_replica_name
		}

		#Optional
		data_guard_replicas {
			#Required
			databases {
				#Required
				cloud_autonomous_vm_cluster_id = oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster.id
				compute_count = var.distributed_autonomous_database_system_data_guard_databases_data_guard_replicas_databases_compute_count
				data_storage_size_in_gbs = var.distributed_autonomous_database_system_data_guard_databases_data_guard_replicas_databases_data_storage_size_in_gbs
				source = var.distributed_autonomous_database_system_data_guard_databases_data_guard_replicas_databases_source

				#Optional
				admin_password = var.distributed_autonomous_database_system_data_guard_databases_data_guard_replicas_databases_admin_password
				admin_password_secret_id = oci_vault_secret.test_secret.id
				admin_password_secret_version_number = var.distributed_autonomous_database_system_data_guard_databases_data_guard_replicas_databases_admin_password_secret_version_number
				fast_start_fail_over_lag_limit_in_seconds = var.distributed_autonomous_database_system_data_guard_databases_data_guard_replicas_databases_fast_start_fail_over_lag_limit_in_seconds
				is_auto_scaling_enabled = var.distributed_autonomous_database_system_data_guard_databases_data_guard_replicas_databases_is_auto_scaling_enabled
				kms_key_id = oci_kms_key.test_key.id
				kms_key_version_id = oci_kms_key_version.test_key_version.id
				okv_end_point_group_name = oci_identity_group.test_group.name
				okv_key_store_id = oci_database_key_store.test_key_store.id
				protection_mode = var.distributed_autonomous_database_system_data_guard_databases_data_guard_replicas_databases_protection_mode
				shard_space_name = var.distributed_autonomous_database_system_data_guard_databases_data_guard_replicas_databases_shard_space_name
				standby_maintenance_buffer_in_days = var.distributed_autonomous_database_system_data_guard_databases_data_guard_replicas_databases_standby_maintenance_buffer_in_days
				vault_id = oci_kms_vault.test_vault.id
			}
			name = var.distributed_autonomous_database_system_data_guard_databases_data_guard_replicas_name
		}
	}
	system_raft_clusters {
		#Required
		databases {
			#Required
			cloud_autonomous_vm_cluster_id = oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster.id
			compute_count = var.distributed_autonomous_database_system_raft_clusters_databases_compute_count
			data_storage_size_in_gbs = var.distributed_autonomous_database_system_raft_clusters_databases_data_storage_size_in_gbs
			source = var.distributed_autonomous_database_system_raft_clusters_databases_source

			#Optional
			admin_password = var.distributed_autonomous_database_system_raft_clusters_databases_admin_password
			admin_password_secret_id = oci_vault_secret.test_secret.id
			admin_password_secret_version_number = var.distributed_autonomous_database_system_raft_clusters_databases_admin_password_secret_version_number
			fast_start_fail_over_lag_limit_in_seconds = var.distributed_autonomous_database_system_raft_clusters_databases_fast_start_fail_over_lag_limit_in_seconds
			is_auto_scaling_enabled = var.distributed_autonomous_database_system_raft_clusters_databases_is_auto_scaling_enabled
			kms_key_id = oci_kms_key.test_key.id
			kms_key_version_id = oci_kms_key_version.test_key_version.id
			okv_end_point_group_name = oci_identity_group.test_group.name
			okv_key_store_id = oci_database_key_store.test_key_store.id
			protection_mode = var.distributed_autonomous_database_system_raft_clusters_databases_protection_mode
			shard_space_name = var.distributed_autonomous_database_system_raft_clusters_databases_shard_space_name
			standby_maintenance_buffer_in_days = var.distributed_autonomous_database_system_raft_clusters_databases_standby_maintenance_buffer_in_days
			vault_id = oci_kms_vault.test_vault.id
		}
		name = var.distributed_autonomous_database_system_raft_clusters_name
		replication_factor = var.distributed_autonomous_database_system_raft_clusters_replication_factor
		ru_mode = var.distributed_autonomous_database_system_raft_clusters_ru_mode
		witness_count = var.distributed_autonomous_database_system_raft_clusters_witness_count
	}
	system_raft_replication_unit_count = var.distributed_autonomous_database_system_raft_replication_unit_count
	user_shard_spaces {
		#Required
		name = var.distributed_autonomous_database_user_shard_spaces_name
		original_replica {
			#Required
			database {
				#Required
				cloud_autonomous_vm_cluster_id = oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster.id
				compute_count = var.distributed_autonomous_database_user_shard_spaces_original_replica_database_compute_count
				data_storage_size_in_gbs = var.distributed_autonomous_database_user_shard_spaces_original_replica_database_data_storage_size_in_gbs
				source = var.distributed_autonomous_database_user_shard_spaces_original_replica_database_source

				#Optional
				admin_password = var.distributed_autonomous_database_user_shard_spaces_original_replica_database_admin_password
				admin_password_secret_id = oci_vault_secret.test_secret.id
				admin_password_secret_version_number = var.distributed_autonomous_database_user_shard_spaces_original_replica_database_admin_password_secret_version_number
				fast_start_fail_over_lag_limit_in_seconds = var.distributed_autonomous_database_user_shard_spaces_original_replica_database_fast_start_fail_over_lag_limit_in_seconds
				is_auto_scaling_enabled = var.distributed_autonomous_database_user_shard_spaces_original_replica_database_is_auto_scaling_enabled
				kms_key_id = oci_kms_key.test_key.id
				kms_key_version_id = oci_kms_key_version.test_key_version.id
				okv_end_point_group_name = oci_identity_group.test_group.name
				okv_key_store_id = oci_database_key_store.test_key_store.id
				protection_mode = var.distributed_autonomous_database_user_shard_spaces_original_replica_database_protection_mode
				shard_space_name = var.distributed_autonomous_database_user_shard_spaces_original_replica_database_shard_space_name
				standby_maintenance_buffer_in_days = var.distributed_autonomous_database_user_shard_spaces_original_replica_database_standby_maintenance_buffer_in_days
				vault_id = oci_kms_vault.test_vault.id
			}
			name = var.distributed_autonomous_database_user_shard_spaces_original_replica_name
		}

		#Optional
		data_guard_replicas {
			#Required
			database {
				#Required
				cloud_autonomous_vm_cluster_id = oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster.id
				compute_count = var.distributed_autonomous_database_user_shard_spaces_data_guard_replicas_database_compute_count
				data_storage_size_in_gbs = var.distributed_autonomous_database_user_shard_spaces_data_guard_replicas_database_data_storage_size_in_gbs
				source = var.distributed_autonomous_database_user_shard_spaces_data_guard_replicas_database_source

				#Optional
				admin_password = var.distributed_autonomous_database_user_shard_spaces_data_guard_replicas_database_admin_password
				admin_password_secret_id = oci_vault_secret.test_secret.id
				admin_password_secret_version_number = var.distributed_autonomous_database_user_shard_spaces_data_guard_replicas_database_admin_password_secret_version_number
				fast_start_fail_over_lag_limit_in_seconds = var.distributed_autonomous_database_user_shard_spaces_data_guard_replicas_database_fast_start_fail_over_lag_limit_in_seconds
				is_auto_scaling_enabled = var.distributed_autonomous_database_user_shard_spaces_data_guard_replicas_database_is_auto_scaling_enabled
				kms_key_id = oci_kms_key.test_key.id
				kms_key_version_id = oci_kms_key_version.test_key_version.id
				okv_end_point_group_name = oci_identity_group.test_group.name
				okv_key_store_id = oci_database_key_store.test_key_store.id
				protection_mode = var.distributed_autonomous_database_user_shard_spaces_data_guard_replicas_database_protection_mode
				shard_space_name = var.distributed_autonomous_database_user_shard_spaces_data_guard_replicas_database_shard_space_name
				standby_maintenance_buffer_in_days = var.distributed_autonomous_database_user_shard_spaces_data_guard_replicas_database_standby_maintenance_buffer_in_days
				vault_id = oci_kms_vault.test_vault.id
			}
			name = var.distributed_autonomous_database_user_shard_spaces_data_guard_replicas_name
		}
	}
	vcn_nsg_ids {
		#Required
		nsg_ids = var.distributed_autonomous_database_vcn_nsg_ids_nsg_ids
		vcn_id = oci_core_vcn.test_vcn.id
	}
}
```

## Argument Reference

The following arguments are supported:

* `catalog_details` - (Required) Collection of catalog for the Globally distributed autonomous database.
	* `data_guard_replicas` - (Optional) Details required to create data guard replicas for the catalog.
		* `database` - (Required) Details for creating a distributed autonomous database catalog.
			* `admin_password` - (Optional) The admin password for the catalog associated with the distributed autonomous database.
			* `admin_password_secret_id` - (Optional) The Oracle Cloud Infrastructure vault secret [/Content/General/Concepts/identifiers.htm]OCID. This cannot be used in conjunction with adminPassword.
			* `admin_password_secret_version_number` - (Optional) The version of the vault secret. If no version is specified, the latest version will be used.
			* `cloud_autonomous_vm_cluster_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud Autonomous VM Cluster.
			* `compute_count` - (Required) The compute count for the catalog database. It has to be in multiples of 2.
			* `data_storage_size_in_gbs` - (Required) The data disk group size to be allocated in GBs for the catalog database.
			* `fast_start_fail_over_lag_limit_in_seconds` - (Optional) The lag time preference based on data loss tolerance in seconds.
			* `is_auto_scaling_enabled` - (Optional) Indicates if vertical auto scaling is enabled for the Autonomous AI Database CPU core count. The default value is `FALSE`. 
			* `kms_key_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
			* `kms_key_version_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. 
			* `okv_end_point_group_name` - (Optional) The OKV endpoint name.
			* `okv_key_store_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key store used to create the catalog.
			* `protection_mode` - (Optional) The protectionMode for the catalog peer.
			* `source` - (Required) The source of the Globally distributed autonomous database type. Use ADBD_EXISTING_CLUSTER for an autonomous database on Dedicated Infrastructure that uses a pre-existing cluster. 
			* `standby_maintenance_buffer_in_days` - (Optional) The scheduling detail for the quarterly maintenance window of the standby Autonomous Container Database. This value represents the number of days before scheduled maintenance of the primary database. 
			* `vault_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts).  This parameter and `kmsKeyId` are required for Customer Managed Keys. 
		* `name` - (Required) The name of the catalog replica.  It must start with a letter and contain only letters, digits, and underscores.  Maximum length is 40 characters. 
	* `original_replica` - (Required) Details required to create a catalog replica.
		* `database` - (Required) Details for creating a distributed autonomous database catalog.
			* `admin_password` - (Optional) The admin password for the catalog associated with the distributed autonomous database.
			* `admin_password_secret_id` - (Optional) The Oracle Cloud Infrastructure vault secret [/Content/General/Concepts/identifiers.htm]OCID. This cannot be used in conjunction with adminPassword.
			* `admin_password_secret_version_number` - (Optional) The version of the vault secret. If no version is specified, the latest version will be used.
			* `cloud_autonomous_vm_cluster_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud Autonomous VM Cluster.
			* `compute_count` - (Required) The compute count for the catalog database. It has to be in multiples of 2.
			* `data_storage_size_in_gbs` - (Required) The data disk group size to be allocated in GBs for the catalog database.
			* `fast_start_fail_over_lag_limit_in_seconds` - (Optional) The lag time preference based on data loss tolerance in seconds.
			* `is_auto_scaling_enabled` - (Optional) Indicates if vertical auto scaling is enabled for the Autonomous AI Database CPU core count. The default value is `FALSE`. 
			* `kms_key_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
			* `kms_key_version_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. 
			* `okv_end_point_group_name` - (Optional) The OKV endpoint name.
			* `okv_key_store_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key store used to create the catalog.
			* `protection_mode` - (Optional) The protectionMode for the catalog peer.
			* `source` - (Required) The source of the Globally distributed autonomous database type. Use ADBD_EXISTING_CLUSTER for an autonomous database on Dedicated Infrastructure that uses a pre-existing cluster. 
			* `standby_maintenance_buffer_in_days` - (Optional) The scheduling detail for the quarterly maintenance window of the standby Autonomous Container Database. This value represents the number of days before scheduled maintenance of the primary database. 
			* `vault_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts).  This parameter and `kmsKeyId` are required for Customer Managed Keys. 
		* `name` - (Required) The name of the catalog replica.  It must start with a letter and contain only letters, digits, and underscores.  Maximum length is 40 characters. 
* `character_set` - (Required) The character set for the database.
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Globally distributed autonomous database compartment.
* `composite_data_guard_shard_spaces` - (Optional) Collection of composite data guard shard spaces.
	* `chunk_count` - (Required) Count of chunks associated with the shard space.
	* `data_guard_replicas` - (Optional) The details of data guard replica for the shard being created.
		* `databases` - (Required) The details of composite data guard shard space replica databases.
			* `admin_password` - (Optional) Admin password for shard database.
			* `admin_password_secret_id` - (Optional) The Oracle Cloud Infrastructure vault secret [/Content/General/Concepts/identifiers.htm]OCID. This cannot be used in conjunction with adminPassword.
			* `admin_password_secret_version_number` - (Optional) The version of the vault secret. If no version is specified, the latest version will be used.
			* `cloud_autonomous_vm_cluster_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud Autonomous Exadata VM Cluster.
			* `compute_count` - (Required) The compute count for the shard database. It has to be in multiples of 2.
			* `data_storage_size_in_gbs` - (Required) The data disk group size to be allocated in GBs for the shard database.
			* `fast_start_fail_over_lag_limit_in_seconds` - (Optional) The lag time preference based on data loss tolerance in seconds.
			* `is_auto_scaling_enabled` - (Optional) Indicates if vertical auto scaling is enabled for the Autonomous AI Database CPU core count. The default value is `FALSE`. 
			* `kms_key_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
			* `kms_key_version_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. 
			* `okv_end_point_group_name` - (Optional) The OKV endpoint name.
			* `okv_key_store_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key store used to create the shard.
			* `protection_mode` - (Optional) The protection mode for the shard peer.
			* `shard_space_name` - (Optional) The shard space name for the shard database. Shard space for an existing shard cannot be changed once the shard is created. The shard space name is used when creating new shards. For user-defined sharding, every shard must have a unique shard space name. For system-defined sharding, a shard space name is not required. 
			* `source` - (Required) The source of the Globally distributed autonomous database type. Use ADBD_EXISTING_CLUSTER for an autonomous database on Dedicated Infrastructure that uses a pre-existing cluster. 
			* `standby_maintenance_buffer_in_days` - (Optional) The scheduling detail for the quarterly maintenance window of the standby Autonomous Container Database. 
			* `vault_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `kmsKeyId` are required for Customer Managed Keys.
		* `name` - (Required) The name of shard space replica.  It must start with a letter and contain only letters, digits, and underscores.  Maximum length is 40 characters. 
	* `name` - (Required) The name of shard space.  It must start with a letter and contain only letters, digits, and underscores.  Maximum length is 40 characters. 
	* `original_replica` - (Required) Details required to create a composite data guard shard space replica for a Globally distributed autonomous database.
		* `databases` - (Required) The details of composite data guard shard space replica databases.
			* `admin_password` - (Optional) Admin password for shard database.
			* `admin_password_secret_id` - (Optional) The Oracle Cloud Infrastructure vault secret [/Content/General/Concepts/identifiers.htm]OCID. This cannot be used in conjunction with adminPassword.
			* `admin_password_secret_version_number` - (Optional) The version of the vault secret. If no version is specified, the latest version will be used.
			* `cloud_autonomous_vm_cluster_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud Autonomous Exadata VM Cluster.
			* `compute_count` - (Required) The compute count for the shard database. It has to be in multiples of 2.
			* `data_storage_size_in_gbs` - (Required) The data disk group size to be allocated in GBs for the shard database.
			* `fast_start_fail_over_lag_limit_in_seconds` - (Optional) The lag time preference based on data loss tolerance in seconds.
			* `is_auto_scaling_enabled` - (Optional) Indicates if vertical auto scaling is enabled for the Autonomous AI Database CPU core count. The default value is `FALSE`. 
			* `kms_key_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
			* `kms_key_version_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. 
			* `okv_end_point_group_name` - (Optional) The OKV endpoint name.
			* `okv_key_store_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key store used to create the shard.
			* `protection_mode` - (Optional) The protection mode for the shard peer.
			* `shard_space_name` - (Optional) The shard space name for the shard database. Shard space for an existing shard cannot be changed once the shard is created. The shard space name is used when creating new shards. For user-defined sharding, every shard must have a unique shard space name. For system-defined sharding, a shard space name is not required. 
			* `source` - (Required) The source of the Globally distributed autonomous database type. Use ADBD_EXISTING_CLUSTER for an autonomous database on Dedicated Infrastructure that uses a pre-existing cluster. 
			* `standby_maintenance_buffer_in_days` - (Optional) The scheduling detail for the quarterly maintenance window of the standby Autonomous Container Database. 
			* `vault_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `kmsKeyId` are required for Customer Managed Keys.
		* `name` - (Required) The name of shard space replica.  It must start with a letter and contain only letters, digits, and underscores.  Maximum length is 40 characters. 
* `composite_raft_shard_spaces` - (Optional) Collection of composite raft shards.
	* `chunk_count` - (Required) Count of chunks associated with the shard space.
	* `name` - (Required) The name of shard space.  It must start with a letter and contain only letters, digits, and underscores.  Maximum length is 40 characters. 
	* `raft_clusters` - (Required) The details of raft clusters associated with the shard space.
		* `databases` - (Required) The details of databases associated with the composite shard space raft cluster.
			* `admin_password` - (Optional) Admin password for shard database.
			* `admin_password_secret_id` - (Optional) The Oracle Cloud Infrastructure vault secret [/Content/General/Concepts/identifiers.htm]OCID. This cannot be used in conjunction with adminPassword.
			* `admin_password_secret_version_number` - (Optional) The version of the vault secret. If no version is specified, the latest version will be used.
			* `cloud_autonomous_vm_cluster_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud Autonomous Exadata VM Cluster.
			* `compute_count` - (Required) The compute count for the shard database. It has to be in multiples of 2.
			* `data_storage_size_in_gbs` - (Required) The data disk group size to be allocated in GBs for the shard database.
			* `fast_start_fail_over_lag_limit_in_seconds` - (Optional) The lag time preference based on data loss tolerance in seconds.
			* `is_auto_scaling_enabled` - (Optional) Indicates if vertical auto scaling is enabled for the Autonomous AI Database CPU core count. The default value is `FALSE`. 
			* `kms_key_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
			* `kms_key_version_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. 
			* `okv_end_point_group_name` - (Optional) The OKV endpoint name.
			* `okv_key_store_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key store used to create the shard.
			* `protection_mode` - (Optional) The protection mode for the shard peer.
			* `shard_space_name` - (Optional) The shard space name for the shard database. Shard space for an existing shard cannot be changed once the shard is created. The shard space name is used when creating new shards. For user-defined sharding, every shard must have a unique shard space name. For system-defined sharding, a shard space name is not required. 
			* `source` - (Required) The source of the Globally distributed autonomous database type. Use ADBD_EXISTING_CLUSTER for an autonomous database on Dedicated Infrastructure that uses a pre-existing cluster. 
			* `standby_maintenance_buffer_in_days` - (Optional) The scheduling detail for the quarterly maintenance window of the standby Autonomous Container Database. 
			* `vault_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `kmsKeyId` are required for Customer Managed Keys.
		* `name` - (Required) The name of raft cluster.  It must start with a letter and contain only letters, digits, and underscores.  Maximum length is 40 characters. 
		* `replication_factor` - (Required) Replication factor associated with the raft cluster.
		* `ru_mode` - (Required) Replication factor associated with the raft cluster.
		* `witness_count` - (Required) Witnesses RU count associated with the raft cluster.
	* `replication_unit_count` - (Required) Number of replication units in the shard space.
* `database_version` - (Required) Oracle Database version for the shards and catalog used in Globally distributed autonomous database.
* `db_backup_config` - (Optional) Backup options for the Distributed Autonomous Database. 
	* `backup_destination_details` - (Optional) Backup destination details.
		* `dbrs_policy_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DBRS policy used for backup.
		* `id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup destination.
		* `internet_proxy` - (Optional) Proxy URL to connect to object store.
		* `is_remote` - (Optional) Indicates whether the backup destination is cross-region or local region.
		* `remote_region` - (Optional) The name of the remote region where the remote automatic incremental backups will be stored. For information about valid region names, see [Regions and Availability Domains](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/regions.htm). 
		* `type` - (Required) Type of the database backup destination.
	* `recovery_window_in_days` - (Optional) Number of days between the current and the earliest point of recoverability covered by automatic backups. This value applies to automatic backups. When the value is updated, it is applied to all existing automatic backups. If the number of specified days is 0 then there will be no backups. 
* `db_workload_type` - (Required) Possible workload types. Currently only OLTP workload type is supported.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Required) (Updatable) The display name of the Globally distributed autonomous database.
* `distributed_autonomous_database_id` - (Required) 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `listener_port` - (Required) The listener port number for the Globally distributed autonomous database. The listener port number has to be unique for a customer tenancy across all distributed autonomous databases. Same port number should not be re-used for any other distributed autonomous database. 
* `listener_port_tls` - (Optional) The TLS listener port number for Globally distributed autonomous database. The TLS listener port number has to be unique for a customer tenancy across all distributed autonomous databases. Same port number should not be re-used for any other distributed autonomous database. The listenerPortTls is mandatory for dedicated infrastructure based distributed autonomous databases. 
* `ncharacter_set` - (Required) The national character set for the database.
* `notification_topic_ids` - (Optional) (Updatable) The collection of [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the notification topics associated with the globally distributed autonomous database.
* `ons_port_local` - (Required) Ons local port number for Globally distributed autonomous database. The onsPortLocal has to be unique for a customer tenancy across all distributed autonomous databases. Same port number should not be re-used for any other distributed autonomous database. 
* `ons_port_remote` - (Required) Ons remote port number for Globally distributed autonomous database. The onsPortRemote has to be unique for a customer tenancy across all distributed autonomous databases. Same port number should not be re-used for any other distributed autonomous database. 
* `patch_operations` - (Optional) Legacy schema retained for compatibility; this field does not execute topology actions on the main resource. Use the separate INSERT, REMOVE, INSERTDG, and REMOVEDG resources documented above.
	* `operation` - (Required) (Updatable) The operation can be one of these values: `INSERT`, `MERGE`, `REMOVE`
	* `selection` - (Required) (Updatable) 
	* `value` - (Required when operation=INSERT | MERGE) (Updatable) 
* `prefix` - (Required) Unique name prefix for the Globally distributed autonomous databases. Only alpha-numeric values are allowed. First character has to be a letter followed by any combination of letter and number. 
* `private_endpoint_ids` - (Required) The collection of [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the private endpoint associated with Globally distributed autonomous database.
* `system_chunk_count` - (Optional) Count of chunks associated with system raft clusters or system data guard databases.
* `system_data_guard_databases` - (Optional) Details required to create a system data guard shard.
	* `data_guard_replicas` - (Optional) The details required to create system data guard replicas.
		* `databases` - (Required) Details of the databases associated with the system data guard replica.
			* `admin_password` - (Optional) Admin password for shard database.
			* `admin_password_secret_id` - (Optional) The Oracle Cloud Infrastructure vault secret [/Content/General/Concepts/identifiers.htm]OCID. This cannot be used in conjunction with adminPassword.
			* `admin_password_secret_version_number` - (Optional) The version of the vault secret. If no version is specified, the latest version will be used.
			* `cloud_autonomous_vm_cluster_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud Autonomous Exadata VM Cluster.
			* `compute_count` - (Required) The compute count for the shard database. It has to be in multiples of 2.
			* `data_storage_size_in_gbs` - (Required) The data disk group size to be allocated in GBs for the shard database.
			* `fast_start_fail_over_lag_limit_in_seconds` - (Optional) The lag time preference based on data loss tolerance in seconds.
			* `is_auto_scaling_enabled` - (Optional) Indicates if vertical auto scaling is enabled for the Autonomous AI Database CPU core count. The default value is `FALSE`. 
			* `kms_key_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
			* `kms_key_version_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. 
			* `okv_end_point_group_name` - (Optional) The OKV endpoint name.
			* `okv_key_store_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key store used to create the shard.
			* `protection_mode` - (Optional) The protection mode for the shard peer.
			* `shard_space_name` - (Optional) The shard space name for the shard database. Shard space for an existing shard cannot be changed once the shard is created. The shard space name is used when creating new shards. For user-defined sharding, every shard must have a unique shard space name. For system-defined sharding, a shard space name is not required. 
			* `source` - (Required) The source of the Globally distributed autonomous database type. Use ADBD_EXISTING_CLUSTER for an autonomous database on Dedicated Infrastructure that uses a pre-existing cluster. 
			* `standby_maintenance_buffer_in_days` - (Optional) The scheduling detail for the quarterly maintenance window of the standby Autonomous Container Database. 
			* `vault_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `kmsKeyId` are required for Customer Managed Keys.
		* `name` - (Required) The name of the data guard replica.  It must start with a letter and contain only letters, digits, and underscores.  Maximum length is 40 characters. 
	* `original_replica` - (Required) Details required to create a system data guard replica.
		* `databases` - (Required) Details of the databases associated with the system data guard replica.
			* `admin_password` - (Optional) Admin password for shard database.
			* `admin_password_secret_id` - (Optional) The Oracle Cloud Infrastructure vault secret [/Content/General/Concepts/identifiers.htm]OCID. This cannot be used in conjunction with adminPassword.
			* `admin_password_secret_version_number` - (Optional) The version of the vault secret. If no version is specified, the latest version will be used.
			* `cloud_autonomous_vm_cluster_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud Autonomous Exadata VM Cluster.
			* `compute_count` - (Required) The compute count for the shard database. It has to be in multiples of 2.
			* `data_storage_size_in_gbs` - (Required) The data disk group size to be allocated in GBs for the shard database.
			* `fast_start_fail_over_lag_limit_in_seconds` - (Optional) The lag time preference based on data loss tolerance in seconds.
			* `is_auto_scaling_enabled` - (Optional) Indicates if vertical auto scaling is enabled for the Autonomous AI Database CPU core count. The default value is `FALSE`. 
			* `kms_key_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
			* `kms_key_version_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. 
			* `okv_end_point_group_name` - (Optional) The OKV endpoint name.
			* `okv_key_store_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key store used to create the shard.
			* `protection_mode` - (Optional) The protection mode for the shard peer.
			* `shard_space_name` - (Optional) The shard space name for the shard database. Shard space for an existing shard cannot be changed once the shard is created. The shard space name is used when creating new shards. For user-defined sharding, every shard must have a unique shard space name. For system-defined sharding, a shard space name is not required. 
			* `source` - (Required) The source of the Globally distributed autonomous database type. Use ADBD_EXISTING_CLUSTER for an autonomous database on Dedicated Infrastructure that uses a pre-existing cluster. 
			* `standby_maintenance_buffer_in_days` - (Optional) The scheduling detail for the quarterly maintenance window of the standby Autonomous Container Database. 
			* `vault_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `kmsKeyId` are required for Customer Managed Keys.
		* `name` - (Required) The name of the data guard replica.  It must start with a letter and contain only letters, digits, and underscores.  Maximum length is 40 characters. 
* `system_raft_clusters` - (Optional) Collection of system raft clusters.
	* `databases` - (Required) Details of the databases associated with the system raft cluster.
		* `admin_password` - (Optional) Admin password for shard database.
		* `admin_password_secret_id` - (Optional) The Oracle Cloud Infrastructure vault secret [/Content/General/Concepts/identifiers.htm]OCID. This cannot be used in conjunction with adminPassword.
		* `admin_password_secret_version_number` - (Optional) The version of the vault secret. If no version is specified, the latest version will be used.
		* `cloud_autonomous_vm_cluster_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud Autonomous Exadata VM Cluster.
		* `compute_count` - (Required) The compute count for the shard database. It has to be in multiples of 2.
		* `data_storage_size_in_gbs` - (Required) The data disk group size to be allocated in GBs for the shard database.
		* `fast_start_fail_over_lag_limit_in_seconds` - (Optional) The lag time preference based on data loss tolerance in seconds.
		* `is_auto_scaling_enabled` - (Optional) Indicates if vertical auto scaling is enabled for the Autonomous AI Database CPU core count. The default value is `FALSE`. 
		* `kms_key_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
		* `kms_key_version_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. 
		* `okv_end_point_group_name` - (Optional) The OKV endpoint name.
		* `okv_key_store_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key store used to create the shard.
		* `protection_mode` - (Optional) The protection mode for the shard peer.
		* `shard_space_name` - (Optional) The shard space name for the shard database. Shard space for an existing shard cannot be changed once the shard is created. The shard space name is used when creating new shards. For user-defined sharding, every shard must have a unique shard space name. For system-defined sharding, a shard space name is not required. 
		* `source` - (Required) The source of the Globally distributed autonomous database type. Use ADBD_EXISTING_CLUSTER for an autonomous database on Dedicated Infrastructure that uses a pre-existing cluster. 
		* `standby_maintenance_buffer_in_days` - (Optional) The scheduling detail for the quarterly maintenance window of the standby Autonomous Container Database. 
		* `vault_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `kmsKeyId` are required for Customer Managed Keys.
	* `name` - (Required) The name of raft cluster.  It must start with a letter and contain only letters, digits, and underscores.  Maximum length is 40 characters. 
	* `replication_factor` - (Required) Replication factor associated with the raft cluster.
	* `ru_mode` - (Required) Replication factor associated with the raft cluster.
	* `witness_count` - (Required) Witnesses RU count associated with the raft cluster.
* `system_raft_replication_unit_count` - (Optional) Number of replication units associated with system raft clusters.
* `user_shard_spaces` - (Optional) Collection of user defined shard spaces.
	* `data_guard_replicas` - (Optional) Details required to create data guard replicas for the user shard space.
		* `database` - (Required) Details for creating a distributed autonomous database shard.
			* `admin_password` - (Optional) Admin password for shard database.
			* `admin_password_secret_id` - (Optional) The Oracle Cloud Infrastructure vault secret [/Content/General/Concepts/identifiers.htm]OCID. This cannot be used in conjunction with adminPassword.
			* `admin_password_secret_version_number` - (Optional) The version of the vault secret. If no version is specified, the latest version will be used.
			* `cloud_autonomous_vm_cluster_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud Autonomous Exadata VM Cluster.
			* `compute_count` - (Required) The compute count for the shard database. It has to be in multiples of 2.
			* `data_storage_size_in_gbs` - (Required) The data disk group size to be allocated in GBs for the shard database.
			* `fast_start_fail_over_lag_limit_in_seconds` - (Optional) The lag time preference based on data loss tolerance in seconds.
			* `is_auto_scaling_enabled` - (Optional) Indicates if vertical auto scaling is enabled for the Autonomous AI Database CPU core count. The default value is `FALSE`. 
			* `kms_key_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
			* `kms_key_version_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. 
			* `okv_end_point_group_name` - (Optional) The OKV endpoint name.
			* `okv_key_store_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key store used to create the shard.
			* `protection_mode` - (Optional) The protection mode for the shard peer.
			* `shard_space_name` - (Optional) The shard space name for the shard database. Shard space for an existing shard cannot be changed once the shard is created. The shard space name is used when creating new shards. For user-defined sharding, every shard must have a unique shard space name. For system-defined sharding, a shard space name is not required. 
			* `source` - (Required) The source of the Globally distributed autonomous database type. Use ADBD_EXISTING_CLUSTER for an autonomous database on Dedicated Infrastructure that uses a pre-existing cluster. 
			* `standby_maintenance_buffer_in_days` - (Optional) The scheduling detail for the quarterly maintenance window of the standby Autonomous Container Database. 
			* `vault_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `kmsKeyId` are required for Customer Managed Keys.
		* `name` - (Required) The name of the user shard space replica.  It must start with a letter and contain only letters, digits, and underscores.  Maximum length is 40 characters. 
	* `name` - (Required) The name of the user shard space.  It must start with a letter and contain only letters, digits, and underscores.  Maximum length is 40 characters. 
	* `original_replica` - (Required) Details required to create a user shard space replica.
		* `database` - (Required) Details for creating a distributed autonomous database shard.
			* `admin_password` - (Optional) Admin password for shard database.
			* `admin_password_secret_id` - (Optional) The Oracle Cloud Infrastructure vault secret [/Content/General/Concepts/identifiers.htm]OCID. This cannot be used in conjunction with adminPassword.
			* `admin_password_secret_version_number` - (Optional) The version of the vault secret. If no version is specified, the latest version will be used.
			* `cloud_autonomous_vm_cluster_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud Autonomous Exadata VM Cluster.
			* `compute_count` - (Required) The compute count for the shard database. It has to be in multiples of 2.
			* `data_storage_size_in_gbs` - (Required) The data disk group size to be allocated in GBs for the shard database.
			* `fast_start_fail_over_lag_limit_in_seconds` - (Optional) The lag time preference based on data loss tolerance in seconds.
			* `is_auto_scaling_enabled` - (Optional) Indicates if vertical auto scaling is enabled for the Autonomous AI Database CPU core count. The default value is `FALSE`. 
			* `kms_key_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
			* `kms_key_version_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. 
			* `okv_end_point_group_name` - (Optional) The OKV endpoint name.
			* `okv_key_store_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key store used to create the shard.
			* `protection_mode` - (Optional) The protection mode for the shard peer.
			* `shard_space_name` - (Optional) The shard space name for the shard database. Shard space for an existing shard cannot be changed once the shard is created. The shard space name is used when creating new shards. For user-defined sharding, every shard must have a unique shard space name. For system-defined sharding, a shard space name is not required. 
			* `source` - (Required) The source of the Globally distributed autonomous database type. Use ADBD_EXISTING_CLUSTER for an autonomous database on Dedicated Infrastructure that uses a pre-existing cluster. 
			* `standby_maintenance_buffer_in_days` - (Optional) The scheduling detail for the quarterly maintenance window of the standby Autonomous Container Database. 
			* `vault_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `kmsKeyId` are required for Customer Managed Keys.
		* `name` - (Required) The name of the user shard space replica.  It must start with a letter and contain only letters, digits, and underscores.  Maximum length is 40 characters. 
* `vcn_nsg_ids` - (Optional) The list of network security group (NSG) details to be associated with the distributed autonomous database.
	* `nsg_ids` - (Required) The list of [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) associated with resources in this VCN. For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). 
	* `vcn_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the customer Virtual Cloud Network.
* `add_gds_control_node_trigger` - (Optional) (Updatable) An optional property when incremented triggers Add Gds Control Node. Could be set to any integer value.
* `change_db_backup_config_trigger` - (Optional) (Updatable) An optional property when incremented triggers Change Db Backup Config. Could be set to any integer value.
* `configure_auto_resource_management_trigger` - (Optional) (Updatable) An optional property when incremented triggers Configure Auto Resource Management. Could be set to any integer value.
* `configure_gsm_wallet_trigger` - (Optional) (Updatable) An optional property when incremented triggers Configure Gsm Wallet. Could be set to any integer value.
* `configure_sharding_trigger` - (Optional) (Updatable) An optional property when incremented triggers Configure Sharding. Could be set to any integer value.
* `generate_wallet_trigger` - (Optional) (Updatable) An optional property when incremented triggers Generate Wallet. Could be set to any integer value.
* `move_replication_unit_trigger` - (Optional) (Updatable) An optional property when incremented triggers Move Replication Unit. Could be set to any integer value.
* `recreate_failed_resource_trigger` - (Optional) (Updatable) An optional property when incremented triggers Recreate Failed Resource. Could be set to any integer value.
* `scale_gsms_trigger` - (Optional) (Updatable) An optional property when incremented triggers Scale Gsms. Could be set to any integer value.
* `operational_state` - (Optional) (Updatable) Set to `AVAILABLE` to start or `STOPPED` to stop an existing database. Omit during creation or to leave start/stop unmanaged. Start runs before other updates; Stop runs after successful updates. Completion is verified using node and catalog metadata from DDB GET: `supportingResourceLifecycleState` (`AVAILABLE`/`STOPPED`). Parent DDB lifecycle does not control this check. Nodes and catalog replicas explicitly marked `DELETED` are excluded. Missing operational metadata on remaining databases produces `UNKNOWN`. Do not combine this argument with incremented start/stop counters.
* `start_database_trigger` - (Optional) (Updatable) An optional property when incremented triggers Start Database. Could be set to any integer value.
* `stop_database_trigger` - (Optional) (Updatable) An optional property when incremented triggers Stop Database. Could be set to any integer value.
* `validate_ca_bundle_trigger` - (Optional) (Updatable) An optional property when incremented triggers Validate Ca Bundle. Could be set to any integer value.
* `validate_network_trigger` - (Optional) (Updatable) An optional property when incremented triggers Validate Network. Could be set to any integer value.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `observed_catalog_details` - Current remote `catalog_details` topology, including changes made by separate action resources. Computed only; has the corresponding topology block structure.
* `observed_user_shard_spaces` - Current remote `user_shard_spaces` topology, including changes made by separate action resources. Computed only; has the corresponding topology block structure.
* `observed_system_raft_clusters` - Current remote `system_raft_clusters` topology, including changes made by separate action resources. Computed only; has the corresponding topology block structure.
* `observed_composite_raft_shard_spaces` - Current remote `composite_raft_shard_spaces` topology, including changes made by separate action resources. Computed only; has the corresponding topology block structure.
* `observed_system_data_guard_databases` - Current remote `system_data_guard_databases` topology, including changes made by separate action resources. Computed only; has the corresponding topology block structure.
* `observed_composite_data_guard_shard_spaces` - Current remote `composite_data_guard_shard_spaces` topology, including changes made by separate action resources. Computed only; has the corresponding topology block structure.

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

