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
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/


Creates a Globally distributed database.


## Managing Topology Changes

The topology blocks on this resource describe its initial configuration. Use the separate action resources for additions and removals:

* [INSERT](distributed_database_distributed_databases_insert.html)
* [REMOVE](distributed_database_distributed_databases_remove.html)
* [INSERTDG](distributed_database_distributed_databases_insert_dg.html)
* [REMOVEDG](distributed_database_distributed_databases_remove_dg.html)

Each action instance submits one to three operations in a single request. Resource labels such as `insert` and `insert_2` are user-defined; each instance submits its own operations. Keep successful instances unchanged and add another instance for additional operations. Changed action lists submit the entire list again; unchanged instances are not replayed. Destroying an action only clears its state and does not execute an inverse operation. Action-resource changes do not replace or destroy this main database resource.

Refresh preserves the initial topology inputs and exposes the current remote topology through the computed `observed_*` attributes below. For an imported resource without an initial topology snapshot, refresh initializes the topology inputs from the remote database.

After adding shards, invoke [distributed_databases configure sharding](distributed_database_distributed_databases_configure_sharding.html) after the successful action. Use explicit dependencies to order separate action instances targeting the same database.

## Example Usage

```hcl
resource "oci_distributed_database_distributed_database" "test_distributed_database" {
	#Required
	catalog_details {
		#Required
		original_replica {
			#Required
			database {
				#Required
				db_home_id = oci_database_db_home.test_db_home.id
				source = var.distributed_database_catalog_details_original_replica_database_source

				#Optional
				admin_password = var.distributed_database_catalog_details_original_replica_database_admin_password
				admin_password_secret_id = oci_vault_secret.test_secret.id
				admin_password_secret_version_number = var.distributed_database_catalog_details_original_replica_database_admin_password_secret_version_number
				availability_domain = var.distributed_database_catalog_details_original_replica_database_availability_domain
				db_storage_vault_details {

					#Optional
					additional_flash_cache_in_percent = var.distributed_database_catalog_details_original_replica_database_db_storage_vault_details_additional_flash_cache_in_percent
					autoscale_limit_in_gbs = var.distributed_database_catalog_details_original_replica_database_db_storage_vault_details_autoscale_limit_in_gbs
					compartment_id = var.compartment_id
					high_capacity_database_storage = var.distributed_database_catalog_details_original_replica_database_db_storage_vault_details_high_capacity_database_storage
					is_autoscale_enabled = var.distributed_database_catalog_details_original_replica_database_db_storage_vault_details_is_autoscale_enabled
					subscription_id = oci_onesubscription_subscription.test_subscription.id
				}
				exascale_db_storage_vault_id = oci_database_exascale_db_storage_vault.test_exascale_db_storage_vault.id
				kms_key_id = oci_kms_key.test_key.id
				kms_key_version_id = oci_kms_key_version.test_key_version.id
				protection_mode = var.distributed_database_catalog_details_original_replica_database_protection_mode
				tde_wallet_password = var.distributed_database_catalog_details_original_replica_database_tde_wallet_password
				tde_wallet_password_secret_id = oci_vault_secret.test_secret.id
				tde_wallet_password_secret_version_number = var.distributed_database_catalog_details_original_replica_database_tde_wallet_password_secret_version_number
				transport_type = var.distributed_database_catalog_details_original_replica_database_transport_type
				vault_id = oci_kms_vault.test_vault.id
				vm_cluster_details {

					#Optional
					backup_network_nsg_ids = var.distributed_database_catalog_details_original_replica_database_vm_cluster_details_backup_network_nsg_ids
					backup_subnet_id = oci_core_subnet.test_subnet.id
					compartment_id = var.compartment_id
					domain = var.distributed_database_catalog_details_original_replica_database_vm_cluster_details_domain
					enabled_ecpu_count = var.distributed_database_catalog_details_original_replica_database_vm_cluster_details_enabled_ecpu_count
					is_diagnostics_events_enabled = var.distributed_database_catalog_details_original_replica_database_vm_cluster_details_is_diagnostics_events_enabled
					is_health_monitoring_enabled = var.distributed_database_catalog_details_original_replica_database_vm_cluster_details_is_health_monitoring_enabled
					is_incident_logs_enabled = var.distributed_database_catalog_details_original_replica_database_vm_cluster_details_is_incident_logs_enabled
					license_model = var.distributed_database_catalog_details_original_replica_database_vm_cluster_details_license_model
					nsg_ids = var.distributed_database_catalog_details_original_replica_database_vm_cluster_details_nsg_ids
					private_zone_id = oci_dns_zone.test_zone.id
					ssh_public_keys = var.distributed_database_catalog_details_original_replica_database_vm_cluster_details_ssh_public_keys
					subnet_id = oci_core_subnet.test_subnet.id
					subscription_id = oci_onesubscription_subscription.test_subscription.id
					total_ecpu_count = var.distributed_database_catalog_details_original_replica_database_vm_cluster_details_total_ecpu_count
					vm_file_system_storage_size = var.distributed_database_catalog_details_original_replica_database_vm_cluster_details_vm_file_system_storage_size
				}
				vm_cluster_id = oci_database_vm_cluster.test_vm_cluster.id
			}
			name = var.distributed_database_catalog_details_original_replica_name
		}

		#Optional
		data_guard_replicas {
			#Required
			database {
				#Required
				db_home_id = oci_database_db_home.test_db_home.id
				source = var.distributed_database_catalog_details_data_guard_replicas_database_source

				#Optional
				admin_password = var.distributed_database_catalog_details_data_guard_replicas_database_admin_password
				admin_password_secret_id = oci_vault_secret.test_secret.id
				admin_password_secret_version_number = var.distributed_database_catalog_details_data_guard_replicas_database_admin_password_secret_version_number
				availability_domain = var.distributed_database_catalog_details_data_guard_replicas_database_availability_domain
				db_storage_vault_details {

					#Optional
					additional_flash_cache_in_percent = var.distributed_database_catalog_details_data_guard_replicas_database_db_storage_vault_details_additional_flash_cache_in_percent
					autoscale_limit_in_gbs = var.distributed_database_catalog_details_data_guard_replicas_database_db_storage_vault_details_autoscale_limit_in_gbs
					compartment_id = var.compartment_id
					high_capacity_database_storage = var.distributed_database_catalog_details_data_guard_replicas_database_db_storage_vault_details_high_capacity_database_storage
					is_autoscale_enabled = var.distributed_database_catalog_details_data_guard_replicas_database_db_storage_vault_details_is_autoscale_enabled
					subscription_id = oci_onesubscription_subscription.test_subscription.id
				}
				exascale_db_storage_vault_id = oci_database_exascale_db_storage_vault.test_exascale_db_storage_vault.id
				kms_key_id = oci_kms_key.test_key.id
				kms_key_version_id = oci_kms_key_version.test_key_version.id
				protection_mode = var.distributed_database_catalog_details_data_guard_replicas_database_protection_mode
				tde_wallet_password = var.distributed_database_catalog_details_data_guard_replicas_database_tde_wallet_password
				tde_wallet_password_secret_id = oci_vault_secret.test_secret.id
				tde_wallet_password_secret_version_number = var.distributed_database_catalog_details_data_guard_replicas_database_tde_wallet_password_secret_version_number
				transport_type = var.distributed_database_catalog_details_data_guard_replicas_database_transport_type
				vault_id = oci_kms_vault.test_vault.id
				vm_cluster_details {

					#Optional
					backup_network_nsg_ids = var.distributed_database_catalog_details_data_guard_replicas_database_vm_cluster_details_backup_network_nsg_ids
					backup_subnet_id = oci_core_subnet.test_subnet.id
					compartment_id = var.compartment_id
					domain = var.distributed_database_catalog_details_data_guard_replicas_database_vm_cluster_details_domain
					enabled_ecpu_count = var.distributed_database_catalog_details_data_guard_replicas_database_vm_cluster_details_enabled_ecpu_count
					is_diagnostics_events_enabled = var.distributed_database_catalog_details_data_guard_replicas_database_vm_cluster_details_is_diagnostics_events_enabled
					is_health_monitoring_enabled = var.distributed_database_catalog_details_data_guard_replicas_database_vm_cluster_details_is_health_monitoring_enabled
					is_incident_logs_enabled = var.distributed_database_catalog_details_data_guard_replicas_database_vm_cluster_details_is_incident_logs_enabled
					license_model = var.distributed_database_catalog_details_data_guard_replicas_database_vm_cluster_details_license_model
					nsg_ids = var.distributed_database_catalog_details_data_guard_replicas_database_vm_cluster_details_nsg_ids
					private_zone_id = oci_dns_zone.test_zone.id
					ssh_public_keys = var.distributed_database_catalog_details_data_guard_replicas_database_vm_cluster_details_ssh_public_keys
					subnet_id = oci_core_subnet.test_subnet.id
					subscription_id = oci_onesubscription_subscription.test_subscription.id
					total_ecpu_count = var.distributed_database_catalog_details_data_guard_replicas_database_vm_cluster_details_total_ecpu_count
					vm_file_system_storage_size = var.distributed_database_catalog_details_data_guard_replicas_database_vm_cluster_details_vm_file_system_storage_size
				}
				vm_cluster_id = oci_database_vm_cluster.test_vm_cluster.id
			}
			name = var.distributed_database_catalog_details_data_guard_replicas_name
		}
	}
	character_set = var.distributed_database_character_set
	compartment_id = var.compartment_id
	database_version = var.distributed_database_database_version
	display_name = var.distributed_database_display_name
	distributed_database_id = var.distributed_database_distributed_database_id
	listener_port = var.distributed_database_listener_port
	ncharacter_set = var.distributed_database_ncharacter_set
	ons_port_local = var.distributed_database_ons_port_local
	ons_port_remote = var.distributed_database_ons_port_remote
	prefix = var.distributed_database_prefix
	private_endpoint_ids = var.distributed_database_private_endpoint_ids

	#Optional
	composite_data_guard_shard_spaces {
		#Required
		chunk_count = var.distributed_database_composite_data_guard_shard_spaces_chunk_count
		name = var.distributed_database_composite_data_guard_shard_spaces_name
		original_replica {
			#Required
			databases {
				#Required
				db_home_id = oci_database_db_home.test_db_home.id
				source = var.distributed_database_composite_data_guard_shard_spaces_original_replica_databases_source

				#Optional
				admin_password = var.distributed_database_composite_data_guard_shard_spaces_original_replica_databases_admin_password
				admin_password_secret_id = oci_vault_secret.test_secret.id
				admin_password_secret_version_number = var.distributed_database_composite_data_guard_shard_spaces_original_replica_databases_admin_password_secret_version_number
				availability_domain = var.distributed_database_composite_data_guard_shard_spaces_original_replica_databases_availability_domain
				db_storage_vault_details {

					#Optional
					additional_flash_cache_in_percent = var.distributed_database_composite_data_guard_shard_spaces_original_replica_databases_db_storage_vault_details_additional_flash_cache_in_percent
					autoscale_limit_in_gbs = var.distributed_database_composite_data_guard_shard_spaces_original_replica_databases_db_storage_vault_details_autoscale_limit_in_gbs
					compartment_id = var.compartment_id
					high_capacity_database_storage = var.distributed_database_composite_data_guard_shard_spaces_original_replica_databases_db_storage_vault_details_high_capacity_database_storage
					is_autoscale_enabled = var.distributed_database_composite_data_guard_shard_spaces_original_replica_databases_db_storage_vault_details_is_autoscale_enabled
					subscription_id = oci_onesubscription_subscription.test_subscription.id
				}
				exascale_db_storage_vault_id = oci_database_exascale_db_storage_vault.test_exascale_db_storage_vault.id
				kms_key_id = oci_kms_key.test_key.id
				kms_key_version_id = oci_kms_key_version.test_key_version.id
				protection_mode = var.distributed_database_composite_data_guard_shard_spaces_original_replica_databases_protection_mode
				tde_wallet_password = var.distributed_database_composite_data_guard_shard_spaces_original_replica_databases_tde_wallet_password
				tde_wallet_password_secret_id = oci_vault_secret.test_secret.id
				tde_wallet_password_secret_version_number = var.distributed_database_composite_data_guard_shard_spaces_original_replica_databases_tde_wallet_password_secret_version_number
				transport_type = var.distributed_database_composite_data_guard_shard_spaces_original_replica_databases_transport_type
				vault_id = oci_kms_vault.test_vault.id
				vm_cluster_details {

					#Optional
					backup_network_nsg_ids = var.distributed_database_composite_data_guard_shard_spaces_original_replica_databases_vm_cluster_details_backup_network_nsg_ids
					backup_subnet_id = oci_core_subnet.test_subnet.id
					compartment_id = var.compartment_id
					domain = var.distributed_database_composite_data_guard_shard_spaces_original_replica_databases_vm_cluster_details_domain
					enabled_ecpu_count = var.distributed_database_composite_data_guard_shard_spaces_original_replica_databases_vm_cluster_details_enabled_ecpu_count
					is_diagnostics_events_enabled = var.distributed_database_composite_data_guard_shard_spaces_original_replica_databases_vm_cluster_details_is_diagnostics_events_enabled
					is_health_monitoring_enabled = var.distributed_database_composite_data_guard_shard_spaces_original_replica_databases_vm_cluster_details_is_health_monitoring_enabled
					is_incident_logs_enabled = var.distributed_database_composite_data_guard_shard_spaces_original_replica_databases_vm_cluster_details_is_incident_logs_enabled
					license_model = var.distributed_database_composite_data_guard_shard_spaces_original_replica_databases_vm_cluster_details_license_model
					nsg_ids = var.distributed_database_composite_data_guard_shard_spaces_original_replica_databases_vm_cluster_details_nsg_ids
					private_zone_id = oci_dns_zone.test_zone.id
					ssh_public_keys = var.distributed_database_composite_data_guard_shard_spaces_original_replica_databases_vm_cluster_details_ssh_public_keys
					subnet_id = oci_core_subnet.test_subnet.id
					subscription_id = oci_onesubscription_subscription.test_subscription.id
					total_ecpu_count = var.distributed_database_composite_data_guard_shard_spaces_original_replica_databases_vm_cluster_details_total_ecpu_count
					vm_file_system_storage_size = var.distributed_database_composite_data_guard_shard_spaces_original_replica_databases_vm_cluster_details_vm_file_system_storage_size
				}
				vm_cluster_id = oci_database_vm_cluster.test_vm_cluster.id
			}
			name = var.distributed_database_composite_data_guard_shard_spaces_original_replica_name
		}

		#Optional
		data_guard_replicas {
			#Required
			databases {
				#Required
				db_home_id = oci_database_db_home.test_db_home.id
				source = var.distributed_database_composite_data_guard_shard_spaces_data_guard_replicas_databases_source

				#Optional
				admin_password = var.distributed_database_composite_data_guard_shard_spaces_data_guard_replicas_databases_admin_password
				admin_password_secret_id = oci_vault_secret.test_secret.id
				admin_password_secret_version_number = var.distributed_database_composite_data_guard_shard_spaces_data_guard_replicas_databases_admin_password_secret_version_number
				availability_domain = var.distributed_database_composite_data_guard_shard_spaces_data_guard_replicas_databases_availability_domain
				db_storage_vault_details {

					#Optional
					additional_flash_cache_in_percent = var.distributed_database_composite_data_guard_shard_spaces_data_guard_replicas_databases_db_storage_vault_details_additional_flash_cache_in_percent
					autoscale_limit_in_gbs = var.distributed_database_composite_data_guard_shard_spaces_data_guard_replicas_databases_db_storage_vault_details_autoscale_limit_in_gbs
					compartment_id = var.compartment_id
					high_capacity_database_storage = var.distributed_database_composite_data_guard_shard_spaces_data_guard_replicas_databases_db_storage_vault_details_high_capacity_database_storage
					is_autoscale_enabled = var.distributed_database_composite_data_guard_shard_spaces_data_guard_replicas_databases_db_storage_vault_details_is_autoscale_enabled
					subscription_id = oci_onesubscription_subscription.test_subscription.id
				}
				exascale_db_storage_vault_id = oci_database_exascale_db_storage_vault.test_exascale_db_storage_vault.id
				kms_key_id = oci_kms_key.test_key.id
				kms_key_version_id = oci_kms_key_version.test_key_version.id
				protection_mode = var.distributed_database_composite_data_guard_shard_spaces_data_guard_replicas_databases_protection_mode
				tde_wallet_password = var.distributed_database_composite_data_guard_shard_spaces_data_guard_replicas_databases_tde_wallet_password
				tde_wallet_password_secret_id = oci_vault_secret.test_secret.id
				tde_wallet_password_secret_version_number = var.distributed_database_composite_data_guard_shard_spaces_data_guard_replicas_databases_tde_wallet_password_secret_version_number
				transport_type = var.distributed_database_composite_data_guard_shard_spaces_data_guard_replicas_databases_transport_type
				vault_id = oci_kms_vault.test_vault.id
				vm_cluster_details {

					#Optional
					backup_network_nsg_ids = var.distributed_database_composite_data_guard_shard_spaces_data_guard_replicas_databases_vm_cluster_details_backup_network_nsg_ids
					backup_subnet_id = oci_core_subnet.test_subnet.id
					compartment_id = var.compartment_id
					domain = var.distributed_database_composite_data_guard_shard_spaces_data_guard_replicas_databases_vm_cluster_details_domain
					enabled_ecpu_count = var.distributed_database_composite_data_guard_shard_spaces_data_guard_replicas_databases_vm_cluster_details_enabled_ecpu_count
					is_diagnostics_events_enabled = var.distributed_database_composite_data_guard_shard_spaces_data_guard_replicas_databases_vm_cluster_details_is_diagnostics_events_enabled
					is_health_monitoring_enabled = var.distributed_database_composite_data_guard_shard_spaces_data_guard_replicas_databases_vm_cluster_details_is_health_monitoring_enabled
					is_incident_logs_enabled = var.distributed_database_composite_data_guard_shard_spaces_data_guard_replicas_databases_vm_cluster_details_is_incident_logs_enabled
					license_model = var.distributed_database_composite_data_guard_shard_spaces_data_guard_replicas_databases_vm_cluster_details_license_model
					nsg_ids = var.distributed_database_composite_data_guard_shard_spaces_data_guard_replicas_databases_vm_cluster_details_nsg_ids
					private_zone_id = oci_dns_zone.test_zone.id
					ssh_public_keys = var.distributed_database_composite_data_guard_shard_spaces_data_guard_replicas_databases_vm_cluster_details_ssh_public_keys
					subnet_id = oci_core_subnet.test_subnet.id
					subscription_id = oci_onesubscription_subscription.test_subscription.id
					total_ecpu_count = var.distributed_database_composite_data_guard_shard_spaces_data_guard_replicas_databases_vm_cluster_details_total_ecpu_count
					vm_file_system_storage_size = var.distributed_database_composite_data_guard_shard_spaces_data_guard_replicas_databases_vm_cluster_details_vm_file_system_storage_size
				}
				vm_cluster_id = oci_database_vm_cluster.test_vm_cluster.id
			}
			name = var.distributed_database_composite_data_guard_shard_spaces_data_guard_replicas_name
		}
	}
	composite_raft_shard_spaces {
		#Required
		chunk_count = var.distributed_database_composite_raft_shard_spaces_chunk_count
		name = var.distributed_database_composite_raft_shard_spaces_name
		raft_clusters {
			#Required
			databases {
				#Required
				db_home_id = oci_database_db_home.test_db_home.id
				source = var.distributed_database_composite_raft_shard_spaces_raft_clusters_databases_source

				#Optional
				admin_password = var.distributed_database_composite_raft_shard_spaces_raft_clusters_databases_admin_password
				admin_password_secret_id = oci_vault_secret.test_secret.id
				admin_password_secret_version_number = var.distributed_database_composite_raft_shard_spaces_raft_clusters_databases_admin_password_secret_version_number
				availability_domain = var.distributed_database_composite_raft_shard_spaces_raft_clusters_databases_availability_domain
				db_storage_vault_details {

					#Optional
					additional_flash_cache_in_percent = var.distributed_database_composite_raft_shard_spaces_raft_clusters_databases_db_storage_vault_details_additional_flash_cache_in_percent
					autoscale_limit_in_gbs = var.distributed_database_composite_raft_shard_spaces_raft_clusters_databases_db_storage_vault_details_autoscale_limit_in_gbs
					compartment_id = var.compartment_id
					high_capacity_database_storage = var.distributed_database_composite_raft_shard_spaces_raft_clusters_databases_db_storage_vault_details_high_capacity_database_storage
					is_autoscale_enabled = var.distributed_database_composite_raft_shard_spaces_raft_clusters_databases_db_storage_vault_details_is_autoscale_enabled
					subscription_id = oci_onesubscription_subscription.test_subscription.id
				}
				exascale_db_storage_vault_id = oci_database_exascale_db_storage_vault.test_exascale_db_storage_vault.id
				kms_key_id = oci_kms_key.test_key.id
				kms_key_version_id = oci_kms_key_version.test_key_version.id
				protection_mode = var.distributed_database_composite_raft_shard_spaces_raft_clusters_databases_protection_mode
				tde_wallet_password = var.distributed_database_composite_raft_shard_spaces_raft_clusters_databases_tde_wallet_password
				tde_wallet_password_secret_id = oci_vault_secret.test_secret.id
				tde_wallet_password_secret_version_number = var.distributed_database_composite_raft_shard_spaces_raft_clusters_databases_tde_wallet_password_secret_version_number
				transport_type = var.distributed_database_composite_raft_shard_spaces_raft_clusters_databases_transport_type
				vault_id = oci_kms_vault.test_vault.id
				vm_cluster_details {

					#Optional
					backup_network_nsg_ids = var.distributed_database_composite_raft_shard_spaces_raft_clusters_databases_vm_cluster_details_backup_network_nsg_ids
					backup_subnet_id = oci_core_subnet.test_subnet.id
					compartment_id = var.compartment_id
					domain = var.distributed_database_composite_raft_shard_spaces_raft_clusters_databases_vm_cluster_details_domain
					enabled_ecpu_count = var.distributed_database_composite_raft_shard_spaces_raft_clusters_databases_vm_cluster_details_enabled_ecpu_count
					is_diagnostics_events_enabled = var.distributed_database_composite_raft_shard_spaces_raft_clusters_databases_vm_cluster_details_is_diagnostics_events_enabled
					is_health_monitoring_enabled = var.distributed_database_composite_raft_shard_spaces_raft_clusters_databases_vm_cluster_details_is_health_monitoring_enabled
					is_incident_logs_enabled = var.distributed_database_composite_raft_shard_spaces_raft_clusters_databases_vm_cluster_details_is_incident_logs_enabled
					license_model = var.distributed_database_composite_raft_shard_spaces_raft_clusters_databases_vm_cluster_details_license_model
					nsg_ids = var.distributed_database_composite_raft_shard_spaces_raft_clusters_databases_vm_cluster_details_nsg_ids
					private_zone_id = oci_dns_zone.test_zone.id
					ssh_public_keys = var.distributed_database_composite_raft_shard_spaces_raft_clusters_databases_vm_cluster_details_ssh_public_keys
					subnet_id = oci_core_subnet.test_subnet.id
					subscription_id = oci_onesubscription_subscription.test_subscription.id
					total_ecpu_count = var.distributed_database_composite_raft_shard_spaces_raft_clusters_databases_vm_cluster_details_total_ecpu_count
					vm_file_system_storage_size = var.distributed_database_composite_raft_shard_spaces_raft_clusters_databases_vm_cluster_details_vm_file_system_storage_size
				}
				vm_cluster_id = oci_database_vm_cluster.test_vm_cluster.id
			}
			name = var.distributed_database_composite_raft_shard_spaces_raft_clusters_name
			replication_factor = var.distributed_database_composite_raft_shard_spaces_raft_clusters_replication_factor
			ru_mode = var.distributed_database_composite_raft_shard_spaces_raft_clusters_ru_mode
			witness_count = var.distributed_database_composite_raft_shard_spaces_raft_clusters_witness_count
		}
		replication_unit_count = var.distributed_database_composite_raft_shard_spaces_replication_unit_count
	}
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
			remote_region = var.distributed_database_db_backup_config_backup_destination_details_remote_region
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
	notification_topic_ids = var.distributed_database_notification_topic_ids
	scan_listener_port = var.distributed_database_scan_listener_port
	system_chunk_count = var.distributed_database_system_chunk_count
	system_data_guard_databases {
		#Required
		original_replica {
			#Required
			databases {
				#Required
				db_home_id = oci_database_db_home.test_db_home.id
				source = var.distributed_database_system_data_guard_databases_original_replica_databases_source

				#Optional
				admin_password = var.distributed_database_system_data_guard_databases_original_replica_databases_admin_password
				admin_password_secret_id = oci_vault_secret.test_secret.id
				admin_password_secret_version_number = var.distributed_database_system_data_guard_databases_original_replica_databases_admin_password_secret_version_number
				availability_domain = var.distributed_database_system_data_guard_databases_original_replica_databases_availability_domain
				db_storage_vault_details {

					#Optional
					additional_flash_cache_in_percent = var.distributed_database_system_data_guard_databases_original_replica_databases_db_storage_vault_details_additional_flash_cache_in_percent
					autoscale_limit_in_gbs = var.distributed_database_system_data_guard_databases_original_replica_databases_db_storage_vault_details_autoscale_limit_in_gbs
					compartment_id = var.compartment_id
					high_capacity_database_storage = var.distributed_database_system_data_guard_databases_original_replica_databases_db_storage_vault_details_high_capacity_database_storage
					is_autoscale_enabled = var.distributed_database_system_data_guard_databases_original_replica_databases_db_storage_vault_details_is_autoscale_enabled
					subscription_id = oci_onesubscription_subscription.test_subscription.id
				}
				exascale_db_storage_vault_id = oci_database_exascale_db_storage_vault.test_exascale_db_storage_vault.id
				kms_key_id = oci_kms_key.test_key.id
				kms_key_version_id = oci_kms_key_version.test_key_version.id
				protection_mode = var.distributed_database_system_data_guard_databases_original_replica_databases_protection_mode
				tde_wallet_password = var.distributed_database_system_data_guard_databases_original_replica_databases_tde_wallet_password
				tde_wallet_password_secret_id = oci_vault_secret.test_secret.id
				tde_wallet_password_secret_version_number = var.distributed_database_system_data_guard_databases_original_replica_databases_tde_wallet_password_secret_version_number
				transport_type = var.distributed_database_system_data_guard_databases_original_replica_databases_transport_type
				vault_id = oci_kms_vault.test_vault.id
				vm_cluster_details {

					#Optional
					backup_network_nsg_ids = var.distributed_database_system_data_guard_databases_original_replica_databases_vm_cluster_details_backup_network_nsg_ids
					backup_subnet_id = oci_core_subnet.test_subnet.id
					compartment_id = var.compartment_id
					domain = var.distributed_database_system_data_guard_databases_original_replica_databases_vm_cluster_details_domain
					enabled_ecpu_count = var.distributed_database_system_data_guard_databases_original_replica_databases_vm_cluster_details_enabled_ecpu_count
					is_diagnostics_events_enabled = var.distributed_database_system_data_guard_databases_original_replica_databases_vm_cluster_details_is_diagnostics_events_enabled
					is_health_monitoring_enabled = var.distributed_database_system_data_guard_databases_original_replica_databases_vm_cluster_details_is_health_monitoring_enabled
					is_incident_logs_enabled = var.distributed_database_system_data_guard_databases_original_replica_databases_vm_cluster_details_is_incident_logs_enabled
					license_model = var.distributed_database_system_data_guard_databases_original_replica_databases_vm_cluster_details_license_model
					nsg_ids = var.distributed_database_system_data_guard_databases_original_replica_databases_vm_cluster_details_nsg_ids
					private_zone_id = oci_dns_zone.test_zone.id
					ssh_public_keys = var.distributed_database_system_data_guard_databases_original_replica_databases_vm_cluster_details_ssh_public_keys
					subnet_id = oci_core_subnet.test_subnet.id
					subscription_id = oci_onesubscription_subscription.test_subscription.id
					total_ecpu_count = var.distributed_database_system_data_guard_databases_original_replica_databases_vm_cluster_details_total_ecpu_count
					vm_file_system_storage_size = var.distributed_database_system_data_guard_databases_original_replica_databases_vm_cluster_details_vm_file_system_storage_size
				}
				vm_cluster_id = oci_database_vm_cluster.test_vm_cluster.id
			}
			name = var.distributed_database_system_data_guard_databases_original_replica_name
		}

		#Optional
		data_guard_replicas {
			#Required
			databases {
				#Required
				db_home_id = oci_database_db_home.test_db_home.id
				source = var.distributed_database_system_data_guard_databases_data_guard_replicas_databases_source

				#Optional
				admin_password = var.distributed_database_system_data_guard_databases_data_guard_replicas_databases_admin_password
				admin_password_secret_id = oci_vault_secret.test_secret.id
				admin_password_secret_version_number = var.distributed_database_system_data_guard_databases_data_guard_replicas_databases_admin_password_secret_version_number
				availability_domain = var.distributed_database_system_data_guard_databases_data_guard_replicas_databases_availability_domain
				db_storage_vault_details {

					#Optional
					additional_flash_cache_in_percent = var.distributed_database_system_data_guard_databases_data_guard_replicas_databases_db_storage_vault_details_additional_flash_cache_in_percent
					autoscale_limit_in_gbs = var.distributed_database_system_data_guard_databases_data_guard_replicas_databases_db_storage_vault_details_autoscale_limit_in_gbs
					compartment_id = var.compartment_id
					high_capacity_database_storage = var.distributed_database_system_data_guard_databases_data_guard_replicas_databases_db_storage_vault_details_high_capacity_database_storage
					is_autoscale_enabled = var.distributed_database_system_data_guard_databases_data_guard_replicas_databases_db_storage_vault_details_is_autoscale_enabled
					subscription_id = oci_onesubscription_subscription.test_subscription.id
				}
				exascale_db_storage_vault_id = oci_database_exascale_db_storage_vault.test_exascale_db_storage_vault.id
				kms_key_id = oci_kms_key.test_key.id
				kms_key_version_id = oci_kms_key_version.test_key_version.id
				protection_mode = var.distributed_database_system_data_guard_databases_data_guard_replicas_databases_protection_mode
				tde_wallet_password = var.distributed_database_system_data_guard_databases_data_guard_replicas_databases_tde_wallet_password
				tde_wallet_password_secret_id = oci_vault_secret.test_secret.id
				tde_wallet_password_secret_version_number = var.distributed_database_system_data_guard_databases_data_guard_replicas_databases_tde_wallet_password_secret_version_number
				transport_type = var.distributed_database_system_data_guard_databases_data_guard_replicas_databases_transport_type
				vault_id = oci_kms_vault.test_vault.id
				vm_cluster_details {

					#Optional
					backup_network_nsg_ids = var.distributed_database_system_data_guard_databases_data_guard_replicas_databases_vm_cluster_details_backup_network_nsg_ids
					backup_subnet_id = oci_core_subnet.test_subnet.id
					compartment_id = var.compartment_id
					domain = var.distributed_database_system_data_guard_databases_data_guard_replicas_databases_vm_cluster_details_domain
					enabled_ecpu_count = var.distributed_database_system_data_guard_databases_data_guard_replicas_databases_vm_cluster_details_enabled_ecpu_count
					is_diagnostics_events_enabled = var.distributed_database_system_data_guard_databases_data_guard_replicas_databases_vm_cluster_details_is_diagnostics_events_enabled
					is_health_monitoring_enabled = var.distributed_database_system_data_guard_databases_data_guard_replicas_databases_vm_cluster_details_is_health_monitoring_enabled
					is_incident_logs_enabled = var.distributed_database_system_data_guard_databases_data_guard_replicas_databases_vm_cluster_details_is_incident_logs_enabled
					license_model = var.distributed_database_system_data_guard_databases_data_guard_replicas_databases_vm_cluster_details_license_model
					nsg_ids = var.distributed_database_system_data_guard_databases_data_guard_replicas_databases_vm_cluster_details_nsg_ids
					private_zone_id = oci_dns_zone.test_zone.id
					ssh_public_keys = var.distributed_database_system_data_guard_databases_data_guard_replicas_databases_vm_cluster_details_ssh_public_keys
					subnet_id = oci_core_subnet.test_subnet.id
					subscription_id = oci_onesubscription_subscription.test_subscription.id
					total_ecpu_count = var.distributed_database_system_data_guard_databases_data_guard_replicas_databases_vm_cluster_details_total_ecpu_count
					vm_file_system_storage_size = var.distributed_database_system_data_guard_databases_data_guard_replicas_databases_vm_cluster_details_vm_file_system_storage_size
				}
				vm_cluster_id = oci_database_vm_cluster.test_vm_cluster.id
			}
			name = var.distributed_database_system_data_guard_databases_data_guard_replicas_name
		}
	}
	system_raft_clusters {
		#Required
		databases {
			#Required
			db_home_id = oci_database_db_home.test_db_home.id
			source = var.distributed_database_system_raft_clusters_databases_source

			#Optional
			admin_password = var.distributed_database_system_raft_clusters_databases_admin_password
			admin_password_secret_id = oci_vault_secret.test_secret.id
			admin_password_secret_version_number = var.distributed_database_system_raft_clusters_databases_admin_password_secret_version_number
			availability_domain = var.distributed_database_system_raft_clusters_databases_availability_domain
			db_storage_vault_details {

				#Optional
				additional_flash_cache_in_percent = var.distributed_database_system_raft_clusters_databases_db_storage_vault_details_additional_flash_cache_in_percent
				autoscale_limit_in_gbs = var.distributed_database_system_raft_clusters_databases_db_storage_vault_details_autoscale_limit_in_gbs
				compartment_id = var.compartment_id
				high_capacity_database_storage = var.distributed_database_system_raft_clusters_databases_db_storage_vault_details_high_capacity_database_storage
				is_autoscale_enabled = var.distributed_database_system_raft_clusters_databases_db_storage_vault_details_is_autoscale_enabled
				subscription_id = oci_onesubscription_subscription.test_subscription.id
			}
			exascale_db_storage_vault_id = oci_database_exascale_db_storage_vault.test_exascale_db_storage_vault.id
			kms_key_id = oci_kms_key.test_key.id
			kms_key_version_id = oci_kms_key_version.test_key_version.id
			protection_mode = var.distributed_database_system_raft_clusters_databases_protection_mode
			tde_wallet_password = var.distributed_database_system_raft_clusters_databases_tde_wallet_password
			tde_wallet_password_secret_id = oci_vault_secret.test_secret.id
			tde_wallet_password_secret_version_number = var.distributed_database_system_raft_clusters_databases_tde_wallet_password_secret_version_number
			transport_type = var.distributed_database_system_raft_clusters_databases_transport_type
			vault_id = oci_kms_vault.test_vault.id
			vm_cluster_details {

				#Optional
				backup_network_nsg_ids = var.distributed_database_system_raft_clusters_databases_vm_cluster_details_backup_network_nsg_ids
				backup_subnet_id = oci_core_subnet.test_subnet.id
				compartment_id = var.compartment_id
				domain = var.distributed_database_system_raft_clusters_databases_vm_cluster_details_domain
				enabled_ecpu_count = var.distributed_database_system_raft_clusters_databases_vm_cluster_details_enabled_ecpu_count
				is_diagnostics_events_enabled = var.distributed_database_system_raft_clusters_databases_vm_cluster_details_is_diagnostics_events_enabled
				is_health_monitoring_enabled = var.distributed_database_system_raft_clusters_databases_vm_cluster_details_is_health_monitoring_enabled
				is_incident_logs_enabled = var.distributed_database_system_raft_clusters_databases_vm_cluster_details_is_incident_logs_enabled
				license_model = var.distributed_database_system_raft_clusters_databases_vm_cluster_details_license_model
				nsg_ids = var.distributed_database_system_raft_clusters_databases_vm_cluster_details_nsg_ids
				private_zone_id = oci_dns_zone.test_zone.id
				ssh_public_keys = var.distributed_database_system_raft_clusters_databases_vm_cluster_details_ssh_public_keys
				subnet_id = oci_core_subnet.test_subnet.id
				subscription_id = oci_onesubscription_subscription.test_subscription.id
				total_ecpu_count = var.distributed_database_system_raft_clusters_databases_vm_cluster_details_total_ecpu_count
				vm_file_system_storage_size = var.distributed_database_system_raft_clusters_databases_vm_cluster_details_vm_file_system_storage_size
			}
			vm_cluster_id = oci_database_vm_cluster.test_vm_cluster.id
		}
		name = var.distributed_database_system_raft_clusters_name
		replication_factor = var.distributed_database_system_raft_clusters_replication_factor
		ru_mode = var.distributed_database_system_raft_clusters_ru_mode
		witness_count = var.distributed_database_system_raft_clusters_witness_count
	}
	system_raft_replication_unit_count = var.distributed_database_system_raft_replication_unit_count
	user_shard_spaces {
		#Required
		name = var.distributed_database_user_shard_spaces_name
		original_replica {
			#Required
			database {
				#Required
				db_home_id = oci_database_db_home.test_db_home.id
				source = var.distributed_database_user_shard_spaces_original_replica_database_source

				#Optional
				admin_password = var.distributed_database_user_shard_spaces_original_replica_database_admin_password
				admin_password_secret_id = oci_vault_secret.test_secret.id
				admin_password_secret_version_number = var.distributed_database_user_shard_spaces_original_replica_database_admin_password_secret_version_number
				availability_domain = var.distributed_database_user_shard_spaces_original_replica_database_availability_domain
				db_storage_vault_details {

					#Optional
					additional_flash_cache_in_percent = var.distributed_database_user_shard_spaces_original_replica_database_db_storage_vault_details_additional_flash_cache_in_percent
					autoscale_limit_in_gbs = var.distributed_database_user_shard_spaces_original_replica_database_db_storage_vault_details_autoscale_limit_in_gbs
					compartment_id = var.compartment_id
					high_capacity_database_storage = var.distributed_database_user_shard_spaces_original_replica_database_db_storage_vault_details_high_capacity_database_storage
					is_autoscale_enabled = var.distributed_database_user_shard_spaces_original_replica_database_db_storage_vault_details_is_autoscale_enabled
					subscription_id = oci_onesubscription_subscription.test_subscription.id
				}
				exascale_db_storage_vault_id = oci_database_exascale_db_storage_vault.test_exascale_db_storage_vault.id
				kms_key_id = oci_kms_key.test_key.id
				kms_key_version_id = oci_kms_key_version.test_key_version.id
				protection_mode = var.distributed_database_user_shard_spaces_original_replica_database_protection_mode
				tde_wallet_password = var.distributed_database_user_shard_spaces_original_replica_database_tde_wallet_password
				tde_wallet_password_secret_id = oci_vault_secret.test_secret.id
				tde_wallet_password_secret_version_number = var.distributed_database_user_shard_spaces_original_replica_database_tde_wallet_password_secret_version_number
				transport_type = var.distributed_database_user_shard_spaces_original_replica_database_transport_type
				vault_id = oci_kms_vault.test_vault.id
				vm_cluster_details {

					#Optional
					backup_network_nsg_ids = var.distributed_database_user_shard_spaces_original_replica_database_vm_cluster_details_backup_network_nsg_ids
					backup_subnet_id = oci_core_subnet.test_subnet.id
					compartment_id = var.compartment_id
					domain = var.distributed_database_user_shard_spaces_original_replica_database_vm_cluster_details_domain
					enabled_ecpu_count = var.distributed_database_user_shard_spaces_original_replica_database_vm_cluster_details_enabled_ecpu_count
					is_diagnostics_events_enabled = var.distributed_database_user_shard_spaces_original_replica_database_vm_cluster_details_is_diagnostics_events_enabled
					is_health_monitoring_enabled = var.distributed_database_user_shard_spaces_original_replica_database_vm_cluster_details_is_health_monitoring_enabled
					is_incident_logs_enabled = var.distributed_database_user_shard_spaces_original_replica_database_vm_cluster_details_is_incident_logs_enabled
					license_model = var.distributed_database_user_shard_spaces_original_replica_database_vm_cluster_details_license_model
					nsg_ids = var.distributed_database_user_shard_spaces_original_replica_database_vm_cluster_details_nsg_ids
					private_zone_id = oci_dns_zone.test_zone.id
					ssh_public_keys = var.distributed_database_user_shard_spaces_original_replica_database_vm_cluster_details_ssh_public_keys
					subnet_id = oci_core_subnet.test_subnet.id
					subscription_id = oci_onesubscription_subscription.test_subscription.id
					total_ecpu_count = var.distributed_database_user_shard_spaces_original_replica_database_vm_cluster_details_total_ecpu_count
					vm_file_system_storage_size = var.distributed_database_user_shard_spaces_original_replica_database_vm_cluster_details_vm_file_system_storage_size
				}
				vm_cluster_id = oci_database_vm_cluster.test_vm_cluster.id
			}
			name = var.distributed_database_user_shard_spaces_original_replica_name
		}

		#Optional
		data_guard_replicas {
			#Required
			database {
				#Required
				db_home_id = oci_database_db_home.test_db_home.id
				source = var.distributed_database_user_shard_spaces_data_guard_replicas_database_source

				#Optional
				admin_password = var.distributed_database_user_shard_spaces_data_guard_replicas_database_admin_password
				admin_password_secret_id = oci_vault_secret.test_secret.id
				admin_password_secret_version_number = var.distributed_database_user_shard_spaces_data_guard_replicas_database_admin_password_secret_version_number
				availability_domain = var.distributed_database_user_shard_spaces_data_guard_replicas_database_availability_domain
				db_storage_vault_details {

					#Optional
					additional_flash_cache_in_percent = var.distributed_database_user_shard_spaces_data_guard_replicas_database_db_storage_vault_details_additional_flash_cache_in_percent
					autoscale_limit_in_gbs = var.distributed_database_user_shard_spaces_data_guard_replicas_database_db_storage_vault_details_autoscale_limit_in_gbs
					compartment_id = var.compartment_id
					high_capacity_database_storage = var.distributed_database_user_shard_spaces_data_guard_replicas_database_db_storage_vault_details_high_capacity_database_storage
					is_autoscale_enabled = var.distributed_database_user_shard_spaces_data_guard_replicas_database_db_storage_vault_details_is_autoscale_enabled
					subscription_id = oci_onesubscription_subscription.test_subscription.id
				}
				exascale_db_storage_vault_id = oci_database_exascale_db_storage_vault.test_exascale_db_storage_vault.id
				kms_key_id = oci_kms_key.test_key.id
				kms_key_version_id = oci_kms_key_version.test_key_version.id
				protection_mode = var.distributed_database_user_shard_spaces_data_guard_replicas_database_protection_mode
				tde_wallet_password = var.distributed_database_user_shard_spaces_data_guard_replicas_database_tde_wallet_password
				tde_wallet_password_secret_id = oci_vault_secret.test_secret.id
				tde_wallet_password_secret_version_number = var.distributed_database_user_shard_spaces_data_guard_replicas_database_tde_wallet_password_secret_version_number
				transport_type = var.distributed_database_user_shard_spaces_data_guard_replicas_database_transport_type
				vault_id = oci_kms_vault.test_vault.id
				vm_cluster_details {

					#Optional
					backup_network_nsg_ids = var.distributed_database_user_shard_spaces_data_guard_replicas_database_vm_cluster_details_backup_network_nsg_ids
					backup_subnet_id = oci_core_subnet.test_subnet.id
					compartment_id = var.compartment_id
					domain = var.distributed_database_user_shard_spaces_data_guard_replicas_database_vm_cluster_details_domain
					enabled_ecpu_count = var.distributed_database_user_shard_spaces_data_guard_replicas_database_vm_cluster_details_enabled_ecpu_count
					is_diagnostics_events_enabled = var.distributed_database_user_shard_spaces_data_guard_replicas_database_vm_cluster_details_is_diagnostics_events_enabled
					is_health_monitoring_enabled = var.distributed_database_user_shard_spaces_data_guard_replicas_database_vm_cluster_details_is_health_monitoring_enabled
					is_incident_logs_enabled = var.distributed_database_user_shard_spaces_data_guard_replicas_database_vm_cluster_details_is_incident_logs_enabled
					license_model = var.distributed_database_user_shard_spaces_data_guard_replicas_database_vm_cluster_details_license_model
					nsg_ids = var.distributed_database_user_shard_spaces_data_guard_replicas_database_vm_cluster_details_nsg_ids
					private_zone_id = oci_dns_zone.test_zone.id
					ssh_public_keys = var.distributed_database_user_shard_spaces_data_guard_replicas_database_vm_cluster_details_ssh_public_keys
					subnet_id = oci_core_subnet.test_subnet.id
					subscription_id = oci_onesubscription_subscription.test_subscription.id
					total_ecpu_count = var.distributed_database_user_shard_spaces_data_guard_replicas_database_vm_cluster_details_total_ecpu_count
					vm_file_system_storage_size = var.distributed_database_user_shard_spaces_data_guard_replicas_database_vm_cluster_details_vm_file_system_storage_size
				}
				vm_cluster_id = oci_database_vm_cluster.test_vm_cluster.id
			}
			name = var.distributed_database_user_shard_spaces_data_guard_replicas_name
		}
	}
	vcn_nsg_ids {
		#Required
		nsg_ids = var.distributed_database_vcn_nsg_ids_nsg_ids
		vcn_id = oci_core_vcn.test_vcn.id
	}
}
```

## Argument Reference

The following arguments are supported:

* `catalog_details` - (Required) Collection of catalog details for the Globally distributed database.
	* `data_guard_replicas` - (Optional) Details required to create data guard replicas for the catalog.
		* `database` - (Required) Details for creating a distributed database catalog.
			* `admin_password` - (Optional) The admin password for the catalog associated with Globally distributed database.
			* `admin_password_secret_id` - (Optional) The Oracle Cloud Infrastructure vault secret [/Content/General/Concepts/identifiers.htm]OCID. This cannot be used in conjunction with adminPassword.
			* `admin_password_secret_version_number` - (Optional) The version of the vault secret. If no version is specified, the latest version will be used.
			* `availability_domain` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The name of the availability domain that the distributed database catalog will be located in.
			* `db_home_id` - (Required when source=EXISTING_DB_HOME) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Home.
			* `db_storage_vault_details` - (Required when source=XS_NEW_VAULT_AND_CLUSTER) Details of the request to create exascale db vault storage for shard or catalog of the distributed database. 
				* `additional_flash_cache_in_percent` - (Applicable when source=XS_NEW_VAULT_AND_CLUSTER) The size of additional Flash Cache in percentage of High Capacity database storage.
				* `autoscale_limit_in_gbs` - (Applicable when source=XS_NEW_VAULT_AND_CLUSTER) The maximum limit, in gigabytes, to which the Vault storage size can automatically scale when auto scaling is enabled for the Database Storage Vault.
				* `compartment_id` - (Required when source=XS_NEW_VAULT_AND_CLUSTER) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for Db Storage Vault.
				* `high_capacity_database_storage` - (Required when source=XS_NEW_VAULT_AND_CLUSTER) Total storage capacity in GB for vault storage.
				* `is_autoscale_enabled` - (Applicable when source=XS_NEW_VAULT_AND_CLUSTER) Indicates if autoscale feature is enabled for the Database Storage Vault. The default value is `FALSE`.
				* `subscription_id` - (Applicable when source=XS_NEW_VAULT_AND_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription with which resource needs to be associated with.
			* `exascale_db_storage_vault_id` - (Required when source=XS_NEW_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the exascale db storage vault.
			* `kms_key_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations. 
			* `kms_key_version_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. 
			* `protection_mode` - (Optional) The protection mode used for the Data Guard association.
			* `source` - (Required) The source of the Globally distributed database type. Use XS_EXISTING_CLUSTER for a database on Exascale Infrastructure that uses a pre-existing cluster. Use XS_NEW_CLUSTER for a database on Exascale Infrastructure that requires a new cluster created from scratch. Use XS_NEW_VAULT_AND_CLUSTER for a database on Exascale Infrastructure that requires a new vault and cluster created from scratch. Use XD_EXISTING_CLUSTER for a database on Dedicated Infrastructure that uses a pre-existing cluster. Use EXISTING_DB_HOME for a database on a pre-existing database home. 
			* `tde_wallet_password` - (Optional) The password to open the TDE wallet. The password must be 9 to 255 characters and contain at least two uppercase letters, two lowercase letters, two numeric characters, and two special characters. The allowed special characters are _, \#, and -. 
			* `tde_wallet_password_secret_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure vault secret. This cannot be used in conjunction with tdeWalletPassword.
			* `tde_wallet_password_secret_version_number` - (Optional) The version of the vault secret. If no version is specified, the latest version will be used.
			* `transport_type` - (Optional) The transport type used for the Data Guard association.
			* `vault_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts).  This parameter and `kmsKeyId` are required for Customer Managed Keys. 
			* `vm_cluster_details` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) Details of the request to create exadb vm cluster for shard or catalog of the distributed database. 
				* `backup_network_nsg_ids` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) A list of the [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that the backup network of this DB system belongs to.  Setting this to an empty array after the list is created removes the resource from all NSGs.  For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). Applicable only to Exadata systems. 
				* `backup_subnet_id` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup network subnet associated with the Exadata VM cluster on Exascale Infrastructure. 
				* `compartment_id` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for VM Cluster.
				* `domain` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) A domain name used for the Exadata VM cluster on Exascale Infrastructure.  If the Oracle-provided internet and VCN resolver is enabled for the specified subnet, then the domain name for the subnet is used (do not provide one). Otherwise, provide a valid DNS domain name.  Hyphens (-) are not permitted. Applies to Exadata Database Service on Exascale Infrastructure only. 
				* `enabled_ecpu_count` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The number of ECPUs to enable for an Exadata VM cluster on Exascale Infrastructure. 
				* `is_diagnostics_events_enabled` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) Indicates whether diagnostic collection is enabled for the VM cluster.  Enabling diagnostic collection allows you to receive Events service notifications for guest VM issues.  Diagnostic collection also allows Oracle to provide enhanced service and proactive support for your Exadata system.  You can enable diagnostic collection during VM cluster provisioning. You can also disable or enable it at any time using the `UpdateVmCluster` API. 
				* `is_health_monitoring_enabled` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) Indicates whether health monitoring is enabled for the VM cluster.  Enabling health monitoring allows Oracle to collect diagnostic data and share it with its operations and support personnel.  You may also receive notifications for some events. Collecting health diagnostics enables Oracle to provide proactive support and enhanced service for your system. Optionally enable health monitoring while provisioning a system. You can also disable or enable health monitoring anytime using the `UpdateVmCluster` API. 
				* `is_incident_logs_enabled` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) Indicates whether incident logs and trace collection are enabled for the VM cluster.  Enabling incident logs collection allows Oracle to receive Events service notifications for guest VM issues, collect incident logs and traces, and use them to diagnose issues and resolve them. Optionally enable incident logs collection while provisioning a system. You can also disable or enable incident logs collection anytime using the `UpdateVmCluster` API. 
				* `license_model` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The Oracle license model that applies to the Exadata VM cluster on Exascale Infrastructure. The default is BRING_YOUR_OWN_LICENSE. 
				* `nsg_ids` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The list of [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the network security groups (NSGs) to which this resource belongs.  Setting this to an empty list removes all resources from all NSGs.  For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). 
				* `private_zone_id` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The private zone ID in which you want DNS records to be created. 
				* `ssh_public_keys` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The public key portion of one or more key pairs used for SSH access to the Exadata VM cluster on Exascale Infrastructure.
				* `subnet_id` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet associated with the Exadata VM cluster on Exascale Infrastructure. 
				* `subscription_id` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription with which resource needs to be associated with.
				* `total_ecpu_count` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The number of Total ECPUs for an Exadata VM cluster on Exascale Infrastructure. 
				* `vm_file_system_storage_size` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) File System Storage Size in GBs for Exadata VM cluster. 
			* `vm_cluster_id` - (Required when source=EXISTING_DB_HOME | XD_EXISTING_CLUSTER | XS_EXISTING_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Cloud VM Cluster.
		* `name` - (Required) The name of the catalog replica.  It must start with a letter and contain only letters, digits, and underscores.  Maximum length is 40 characters. 
	* `original_replica` - (Required) Details required to create a catalog replica.
		* `database` - (Required) Details for creating a distributed database catalog.
			* `admin_password` - (Optional) The admin password for the catalog associated with Globally distributed database.
			* `admin_password_secret_id` - (Optional) The Oracle Cloud Infrastructure vault secret [/Content/General/Concepts/identifiers.htm]OCID. This cannot be used in conjunction with adminPassword.
			* `admin_password_secret_version_number` - (Optional) The version of the vault secret. If no version is specified, the latest version will be used.
			* `availability_domain` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The name of the availability domain that the distributed database catalog will be located in.
			* `db_home_id` - (Required when source=EXISTING_DB_HOME) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Home.
			* `db_storage_vault_details` - (Required when source=XS_NEW_VAULT_AND_CLUSTER) Details of the request to create exascale db vault storage for shard or catalog of the distributed database. 
				* `additional_flash_cache_in_percent` - (Applicable when source=XS_NEW_VAULT_AND_CLUSTER) The size of additional Flash Cache in percentage of High Capacity database storage.
				* `autoscale_limit_in_gbs` - (Applicable when source=XS_NEW_VAULT_AND_CLUSTER) The maximum limit, in gigabytes, to which the Vault storage size can automatically scale when auto scaling is enabled for the Database Storage Vault.
				* `compartment_id` - (Required when source=XS_NEW_VAULT_AND_CLUSTER) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for Db Storage Vault.
				* `high_capacity_database_storage` - (Required when source=XS_NEW_VAULT_AND_CLUSTER) Total storage capacity in GB for vault storage.
				* `is_autoscale_enabled` - (Applicable when source=XS_NEW_VAULT_AND_CLUSTER) Indicates if autoscale feature is enabled for the Database Storage Vault. The default value is `FALSE`.
				* `subscription_id` - (Applicable when source=XS_NEW_VAULT_AND_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription with which resource needs to be associated with.
			* `exascale_db_storage_vault_id` - (Required when source=XS_NEW_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the exascale db storage vault.
			* `kms_key_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations. 
			* `kms_key_version_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. 
			* `protection_mode` - (Optional) The protection mode used for the Data Guard association.
			* `source` - (Required) The source of the Globally distributed database type. Use XS_EXISTING_CLUSTER for a database on Exascale Infrastructure that uses a pre-existing cluster. Use XS_NEW_CLUSTER for a database on Exascale Infrastructure that requires a new cluster created from scratch. Use XS_NEW_VAULT_AND_CLUSTER for a database on Exascale Infrastructure that requires a new vault and cluster created from scratch. Use XD_EXISTING_CLUSTER for a database on Dedicated Infrastructure that uses a pre-existing cluster. Use EXISTING_DB_HOME for a database on a pre-existing database home. 
			* `tde_wallet_password` - (Optional) The password to open the TDE wallet. The password must be 9 to 255 characters and contain at least two uppercase letters, two lowercase letters, two numeric characters, and two special characters. The allowed special characters are _, \#, and -. 
			* `tde_wallet_password_secret_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure vault secret. This cannot be used in conjunction with tdeWalletPassword.
			* `tde_wallet_password_secret_version_number` - (Optional) The version of the vault secret. If no version is specified, the latest version will be used.
			* `transport_type` - (Optional) The transport type used for the Data Guard association.
			* `vault_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts).  This parameter and `kmsKeyId` are required for Customer Managed Keys. 
			* `vm_cluster_details` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) Details of the request to create exadb vm cluster for shard or catalog of the distributed database. 
				* `backup_network_nsg_ids` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) A list of the [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that the backup network of this DB system belongs to.  Setting this to an empty array after the list is created removes the resource from all NSGs.  For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). Applicable only to Exadata systems. 
				* `backup_subnet_id` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup network subnet associated with the Exadata VM cluster on Exascale Infrastructure. 
				* `compartment_id` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for VM Cluster.
				* `domain` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) A domain name used for the Exadata VM cluster on Exascale Infrastructure.  If the Oracle-provided internet and VCN resolver is enabled for the specified subnet, then the domain name for the subnet is used (do not provide one). Otherwise, provide a valid DNS domain name.  Hyphens (-) are not permitted. Applies to Exadata Database Service on Exascale Infrastructure only. 
				* `enabled_ecpu_count` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The number of ECPUs to enable for an Exadata VM cluster on Exascale Infrastructure. 
				* `is_diagnostics_events_enabled` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) Indicates whether diagnostic collection is enabled for the VM cluster.  Enabling diagnostic collection allows you to receive Events service notifications for guest VM issues.  Diagnostic collection also allows Oracle to provide enhanced service and proactive support for your Exadata system.  You can enable diagnostic collection during VM cluster provisioning. You can also disable or enable it at any time using the `UpdateVmCluster` API. 
				* `is_health_monitoring_enabled` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) Indicates whether health monitoring is enabled for the VM cluster.  Enabling health monitoring allows Oracle to collect diagnostic data and share it with its operations and support personnel.  You may also receive notifications for some events. Collecting health diagnostics enables Oracle to provide proactive support and enhanced service for your system. Optionally enable health monitoring while provisioning a system. You can also disable or enable health monitoring anytime using the `UpdateVmCluster` API. 
				* `is_incident_logs_enabled` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) Indicates whether incident logs and trace collection are enabled for the VM cluster.  Enabling incident logs collection allows Oracle to receive Events service notifications for guest VM issues, collect incident logs and traces, and use them to diagnose issues and resolve them. Optionally enable incident logs collection while provisioning a system. You can also disable or enable incident logs collection anytime using the `UpdateVmCluster` API. 
				* `license_model` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The Oracle license model that applies to the Exadata VM cluster on Exascale Infrastructure. The default is BRING_YOUR_OWN_LICENSE. 
				* `nsg_ids` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The list of [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the network security groups (NSGs) to which this resource belongs.  Setting this to an empty list removes all resources from all NSGs.  For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). 
				* `private_zone_id` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The private zone ID in which you want DNS records to be created. 
				* `ssh_public_keys` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The public key portion of one or more key pairs used for SSH access to the Exadata VM cluster on Exascale Infrastructure.
				* `subnet_id` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet associated with the Exadata VM cluster on Exascale Infrastructure. 
				* `subscription_id` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription with which resource needs to be associated with.
				* `total_ecpu_count` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The number of Total ECPUs for an Exadata VM cluster on Exascale Infrastructure. 
				* `vm_file_system_storage_size` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) File System Storage Size in GBs for Exadata VM cluster. 
			* `vm_cluster_id` - (Required when source=EXISTING_DB_HOME | XD_EXISTING_CLUSTER | XS_EXISTING_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Cloud VM Cluster.
		* `name` - (Required) The name of the catalog replica.  It must start with a letter and contain only letters, digits, and underscores.  Maximum length is 40 characters. 
* `character_set` - (Required) The character set for the database.
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Globally distributed database compartment.
* `composite_data_guard_shard_spaces` - (Optional) Collection of composite data guard shard spaces.
	* `chunk_count` - (Required) Count of chunks associated with the shard space.
	* `data_guard_replicas` - (Optional) The details of data guard replica for the shard being created.
		* `databases` - (Required) The details of composite data guard shard space replica databases.
			* `admin_password` - (Optional) The admin password for the shard associated with Globally distributed database.
			* `admin_password_secret_id` - (Optional) The Oracle Cloud Infrastructure vault secret [/Content/General/Concepts/identifiers.htm]OCID. This cannot be used in conjunction with adminPassword.
			* `admin_password_secret_version_number` - (Optional) The version of the vault secret. If no version is specified, the latest version will be used.
			* `availability_domain` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The name of the availability domain that the distributed database shard will be located in  like yKoY:US-ASHBURN-AD-3.
			* `db_home_id` - (Required when source=EXISTING_DB_HOME) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Home.
			* `db_storage_vault_details` - (Required when source=XS_NEW_VAULT_AND_CLUSTER) Details of the request to create exascale db vault storage for shard or catalog of the distributed database. 
				* `additional_flash_cache_in_percent` - (Applicable when source=XS_NEW_VAULT_AND_CLUSTER) The size of additional Flash Cache in percentage of High Capacity database storage.
				* `autoscale_limit_in_gbs` - (Applicable when source=XS_NEW_VAULT_AND_CLUSTER) The maximum limit, in gigabytes, to which the Vault storage size can automatically scale when auto scaling is enabled for the Database Storage Vault.
				* `compartment_id` - (Required when source=XS_NEW_VAULT_AND_CLUSTER) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for Db Storage Vault.
				* `high_capacity_database_storage` - (Required when source=XS_NEW_VAULT_AND_CLUSTER) Total storage capacity in GB for vault storage.
				* `is_autoscale_enabled` - (Applicable when source=XS_NEW_VAULT_AND_CLUSTER) Indicates if autoscale feature is enabled for the Database Storage Vault. The default value is `FALSE`.
				* `subscription_id` - (Applicable when source=XS_NEW_VAULT_AND_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription with which resource needs to be associated with.
			* `exascale_db_storage_vault_id` - (Required when source=XS_NEW_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the exascale db storage vault.
			* `kms_key_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations. 
			* `kms_key_version_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. 
			* `protection_mode` - (Optional) The protection mode used for the Data Guard association.
			* `source` - (Required) The source of the Globally distributed database type. Use XS_EXISTING_CLUSTER for a database on Exascale Infrastructure that uses a pre-existing cluster. Use XS_NEW_CLUSTER for a database on Exascale Infrastructure that requires a new cluster created from scratch. Use XS_NEW_VAULT_AND_CLUSTER for a database on Exascale Infrastructure that requires a new vault and cluster created from scratch. Use XD_EXISTING_CLUSTER for a database on Dedicated Infrastructure that uses a pre-existing cluster. Use EXISTING_DB_HOME for a database on a pre-existing database home. 
			* `tde_wallet_password` - (Optional) The password to open the TDE wallet. The password must be 9 to 255 characters and contain at least two uppercase letters, two lowercase letters, two numeric characters, and two special characters. The allowed special characters are _, \#, and -. 
			* `tde_wallet_password_secret_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure vault secret. This cannot be used in conjunction with tdeWalletPassword.
			* `tde_wallet_password_secret_version_number` - (Optional) The version of the vault secret. If no version is specified, the latest version will be used.
			* `transport_type` - (Optional) The transport type used for the Data Guard association.
			* `vault_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts).  This parameter and `kmsKeyId` are required for Customer Managed Keys. 
			* `vm_cluster_details` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) Details of the request to create exadb vm cluster for shard or catalog of the distributed database. 
				* `backup_network_nsg_ids` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) A list of the [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that the backup network of this DB system belongs to.  Setting this to an empty array after the list is created removes the resource from all NSGs.  For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). Applicable only to Exadata systems. 
				* `backup_subnet_id` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup network subnet associated with the Exadata VM cluster on Exascale Infrastructure. 
				* `compartment_id` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for VM Cluster.
				* `domain` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) A domain name used for the Exadata VM cluster on Exascale Infrastructure.  If the Oracle-provided internet and VCN resolver is enabled for the specified subnet, then the domain name for the subnet is used (do not provide one). Otherwise, provide a valid DNS domain name.  Hyphens (-) are not permitted. Applies to Exadata Database Service on Exascale Infrastructure only. 
				* `enabled_ecpu_count` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The number of ECPUs to enable for an Exadata VM cluster on Exascale Infrastructure. 
				* `is_diagnostics_events_enabled` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) Indicates whether diagnostic collection is enabled for the VM cluster.  Enabling diagnostic collection allows you to receive Events service notifications for guest VM issues.  Diagnostic collection also allows Oracle to provide enhanced service and proactive support for your Exadata system.  You can enable diagnostic collection during VM cluster provisioning. You can also disable or enable it at any time using the `UpdateVmCluster` API. 
				* `is_health_monitoring_enabled` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) Indicates whether health monitoring is enabled for the VM cluster.  Enabling health monitoring allows Oracle to collect diagnostic data and share it with its operations and support personnel.  You may also receive notifications for some events. Collecting health diagnostics enables Oracle to provide proactive support and enhanced service for your system. Optionally enable health monitoring while provisioning a system. You can also disable or enable health monitoring anytime using the `UpdateVmCluster` API. 
				* `is_incident_logs_enabled` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) Indicates whether incident logs and trace collection are enabled for the VM cluster.  Enabling incident logs collection allows Oracle to receive Events service notifications for guest VM issues, collect incident logs and traces, and use them to diagnose issues and resolve them. Optionally enable incident logs collection while provisioning a system. You can also disable or enable incident logs collection anytime using the `UpdateVmCluster` API. 
				* `license_model` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The Oracle license model that applies to the Exadata VM cluster on Exascale Infrastructure. The default is BRING_YOUR_OWN_LICENSE. 
				* `nsg_ids` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The list of [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the network security groups (NSGs) to which this resource belongs.  Setting this to an empty list removes all resources from all NSGs.  For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). 
				* `private_zone_id` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The private zone ID in which you want DNS records to be created. 
				* `ssh_public_keys` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The public key portion of one or more key pairs used for SSH access to the Exadata VM cluster on Exascale Infrastructure.
				* `subnet_id` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet associated with the Exadata VM cluster on Exascale Infrastructure. 
				* `subscription_id` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription with which resource needs to be associated with.
				* `total_ecpu_count` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The number of Total ECPUs for an Exadata VM cluster on Exascale Infrastructure. 
				* `vm_file_system_storage_size` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) File System Storage Size in GBs for Exadata VM cluster. 
			* `vm_cluster_id` - (Required when source=EXISTING_DB_HOME | XD_EXISTING_CLUSTER | XS_EXISTING_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Cloud VM Cluster.
		* `name` - (Required) The name of shard space replica.  It must start with a letter and contain only letters, digits, and underscores.  Maximum length is 40 characters. 
	* `name` - (Required) The name of shard space.  It must start with a letter and contain only letters, digits, and underscores.  Maximum length is 40 characters. 
	* `original_replica` - (Required) Globally distributed database composite data guard shard space replica creation details.
		* `databases` - (Required) The details of composite data guard shard space replica databases.
			* `admin_password` - (Optional) The admin password for the shard associated with Globally distributed database.
			* `admin_password_secret_id` - (Optional) The Oracle Cloud Infrastructure vault secret [/Content/General/Concepts/identifiers.htm]OCID. This cannot be used in conjunction with adminPassword.
			* `admin_password_secret_version_number` - (Optional) The version of the vault secret. If no version is specified, the latest version will be used.
			* `availability_domain` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The name of the availability domain that the distributed database shard will be located in  like yKoY:US-ASHBURN-AD-3.
			* `db_home_id` - (Required when source=EXISTING_DB_HOME) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Home.
			* `db_storage_vault_details` - (Required when source=XS_NEW_VAULT_AND_CLUSTER) Details of the request to create exascale db vault storage for shard or catalog of the distributed database. 
				* `additional_flash_cache_in_percent` - (Applicable when source=XS_NEW_VAULT_AND_CLUSTER) The size of additional Flash Cache in percentage of High Capacity database storage.
				* `autoscale_limit_in_gbs` - (Applicable when source=XS_NEW_VAULT_AND_CLUSTER) The maximum limit, in gigabytes, to which the Vault storage size can automatically scale when auto scaling is enabled for the Database Storage Vault.
				* `compartment_id` - (Required when source=XS_NEW_VAULT_AND_CLUSTER) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for Db Storage Vault.
				* `high_capacity_database_storage` - (Required when source=XS_NEW_VAULT_AND_CLUSTER) Total storage capacity in GB for vault storage.
				* `is_autoscale_enabled` - (Applicable when source=XS_NEW_VAULT_AND_CLUSTER) Indicates if autoscale feature is enabled for the Database Storage Vault. The default value is `FALSE`.
				* `subscription_id` - (Applicable when source=XS_NEW_VAULT_AND_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription with which resource needs to be associated with.
			* `exascale_db_storage_vault_id` - (Required when source=XS_NEW_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the exascale db storage vault.
			* `kms_key_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations. 
			* `kms_key_version_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. 
			* `protection_mode` - (Optional) The protection mode used for the Data Guard association.
			* `source` - (Required) The source of the Globally distributed database type. Use XS_EXISTING_CLUSTER for a database on Exascale Infrastructure that uses a pre-existing cluster. Use XS_NEW_CLUSTER for a database on Exascale Infrastructure that requires a new cluster created from scratch. Use XS_NEW_VAULT_AND_CLUSTER for a database on Exascale Infrastructure that requires a new vault and cluster created from scratch. Use XD_EXISTING_CLUSTER for a database on Dedicated Infrastructure that uses a pre-existing cluster. Use EXISTING_DB_HOME for a database on a pre-existing database home. 
			* `tde_wallet_password` - (Optional) The password to open the TDE wallet. The password must be 9 to 255 characters and contain at least two uppercase letters, two lowercase letters, two numeric characters, and two special characters. The allowed special characters are _, \#, and -. 
			* `tde_wallet_password_secret_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure vault secret. This cannot be used in conjunction with tdeWalletPassword.
			* `tde_wallet_password_secret_version_number` - (Optional) The version of the vault secret. If no version is specified, the latest version will be used.
			* `transport_type` - (Optional) The transport type used for the Data Guard association.
			* `vault_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts).  This parameter and `kmsKeyId` are required for Customer Managed Keys. 
			* `vm_cluster_details` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) Details of the request to create exadb vm cluster for shard or catalog of the distributed database. 
				* `backup_network_nsg_ids` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) A list of the [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that the backup network of this DB system belongs to.  Setting this to an empty array after the list is created removes the resource from all NSGs.  For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). Applicable only to Exadata systems. 
				* `backup_subnet_id` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup network subnet associated with the Exadata VM cluster on Exascale Infrastructure. 
				* `compartment_id` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for VM Cluster.
				* `domain` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) A domain name used for the Exadata VM cluster on Exascale Infrastructure.  If the Oracle-provided internet and VCN resolver is enabled for the specified subnet, then the domain name for the subnet is used (do not provide one). Otherwise, provide a valid DNS domain name.  Hyphens (-) are not permitted. Applies to Exadata Database Service on Exascale Infrastructure only. 
				* `enabled_ecpu_count` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The number of ECPUs to enable for an Exadata VM cluster on Exascale Infrastructure. 
				* `is_diagnostics_events_enabled` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) Indicates whether diagnostic collection is enabled for the VM cluster.  Enabling diagnostic collection allows you to receive Events service notifications for guest VM issues.  Diagnostic collection also allows Oracle to provide enhanced service and proactive support for your Exadata system.  You can enable diagnostic collection during VM cluster provisioning. You can also disable or enable it at any time using the `UpdateVmCluster` API. 
				* `is_health_monitoring_enabled` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) Indicates whether health monitoring is enabled for the VM cluster.  Enabling health monitoring allows Oracle to collect diagnostic data and share it with its operations and support personnel.  You may also receive notifications for some events. Collecting health diagnostics enables Oracle to provide proactive support and enhanced service for your system. Optionally enable health monitoring while provisioning a system. You can also disable or enable health monitoring anytime using the `UpdateVmCluster` API. 
				* `is_incident_logs_enabled` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) Indicates whether incident logs and trace collection are enabled for the VM cluster.  Enabling incident logs collection allows Oracle to receive Events service notifications for guest VM issues, collect incident logs and traces, and use them to diagnose issues and resolve them. Optionally enable incident logs collection while provisioning a system. You can also disable or enable incident logs collection anytime using the `UpdateVmCluster` API. 
				* `license_model` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The Oracle license model that applies to the Exadata VM cluster on Exascale Infrastructure. The default is BRING_YOUR_OWN_LICENSE. 
				* `nsg_ids` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The list of [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the network security groups (NSGs) to which this resource belongs.  Setting this to an empty list removes all resources from all NSGs.  For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). 
				* `private_zone_id` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The private zone ID in which you want DNS records to be created. 
				* `ssh_public_keys` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The public key portion of one or more key pairs used for SSH access to the Exadata VM cluster on Exascale Infrastructure.
				* `subnet_id` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet associated with the Exadata VM cluster on Exascale Infrastructure. 
				* `subscription_id` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription with which resource needs to be associated with.
				* `total_ecpu_count` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The number of Total ECPUs for an Exadata VM cluster on Exascale Infrastructure. 
				* `vm_file_system_storage_size` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) File System Storage Size in GBs for Exadata VM cluster. 
			* `vm_cluster_id` - (Required when source=EXISTING_DB_HOME | XD_EXISTING_CLUSTER | XS_EXISTING_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Cloud VM Cluster.
		* `name` - (Required) The name of shard space replica.  It must start with a letter and contain only letters, digits, and underscores.  Maximum length is 40 characters. 
* `composite_raft_shard_spaces` - (Optional) Collection of composite raft shards.
	* `chunk_count` - (Required) Count of chunks associated with the shard space.
	* `name` - (Required) The name of shard space.  It must start with a letter and contain only letters, digits, and underscores.  Maximum length is 40 characters. 
	* `raft_clusters` - (Required) The details of raft clusters associated with the shard space.
		* `databases` - (Required) The details of databases associated with the composite shard space raft cluster.
			* `admin_password` - (Optional) The admin password for the shard associated with Globally distributed database.
			* `admin_password_secret_id` - (Optional) The Oracle Cloud Infrastructure vault secret [/Content/General/Concepts/identifiers.htm]OCID. This cannot be used in conjunction with adminPassword.
			* `admin_password_secret_version_number` - (Optional) The version of the vault secret. If no version is specified, the latest version will be used.
			* `availability_domain` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The name of the availability domain that the distributed database shard will be located in  like yKoY:US-ASHBURN-AD-3.
			* `db_home_id` - (Required when source=EXISTING_DB_HOME) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Home.
			* `db_storage_vault_details` - (Required when source=XS_NEW_VAULT_AND_CLUSTER) Details of the request to create exascale db vault storage for shard or catalog of the distributed database. 
				* `additional_flash_cache_in_percent` - (Applicable when source=XS_NEW_VAULT_AND_CLUSTER) The size of additional Flash Cache in percentage of High Capacity database storage.
				* `autoscale_limit_in_gbs` - (Applicable when source=XS_NEW_VAULT_AND_CLUSTER) The maximum limit, in gigabytes, to which the Vault storage size can automatically scale when auto scaling is enabled for the Database Storage Vault.
				* `compartment_id` - (Required when source=XS_NEW_VAULT_AND_CLUSTER) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for Db Storage Vault.
				* `high_capacity_database_storage` - (Required when source=XS_NEW_VAULT_AND_CLUSTER) Total storage capacity in GB for vault storage.
				* `is_autoscale_enabled` - (Applicable when source=XS_NEW_VAULT_AND_CLUSTER) Indicates if autoscale feature is enabled for the Database Storage Vault. The default value is `FALSE`.
				* `subscription_id` - (Applicable when source=XS_NEW_VAULT_AND_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription with which resource needs to be associated with.
			* `exascale_db_storage_vault_id` - (Required when source=XS_NEW_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the exascale db storage vault.
			* `kms_key_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations. 
			* `kms_key_version_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. 
			* `protection_mode` - (Optional) The protection mode used for the Data Guard association.
			* `source` - (Required) The source of the Globally distributed database type. Use XS_EXISTING_CLUSTER for a database on Exascale Infrastructure that uses a pre-existing cluster. Use XS_NEW_CLUSTER for a database on Exascale Infrastructure that requires a new cluster created from scratch. Use XS_NEW_VAULT_AND_CLUSTER for a database on Exascale Infrastructure that requires a new vault and cluster created from scratch. Use XD_EXISTING_CLUSTER for a database on Dedicated Infrastructure that uses a pre-existing cluster. Use EXISTING_DB_HOME for a database on a pre-existing database home. 
			* `tde_wallet_password` - (Optional) The password to open the TDE wallet. The password must be 9 to 255 characters and contain at least two uppercase letters, two lowercase letters, two numeric characters, and two special characters. The allowed special characters are _, \#, and -. 
			* `tde_wallet_password_secret_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure vault secret. This cannot be used in conjunction with tdeWalletPassword.
			* `tde_wallet_password_secret_version_number` - (Optional) The version of the vault secret. If no version is specified, the latest version will be used.
			* `transport_type` - (Optional) The transport type used for the Data Guard association.
			* `vault_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts).  This parameter and `kmsKeyId` are required for Customer Managed Keys. 
			* `vm_cluster_details` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) Details of the request to create exadb vm cluster for shard or catalog of the distributed database. 
				* `backup_network_nsg_ids` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) A list of the [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that the backup network of this DB system belongs to.  Setting this to an empty array after the list is created removes the resource from all NSGs.  For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). Applicable only to Exadata systems. 
				* `backup_subnet_id` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup network subnet associated with the Exadata VM cluster on Exascale Infrastructure. 
				* `compartment_id` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for VM Cluster.
				* `domain` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) A domain name used for the Exadata VM cluster on Exascale Infrastructure.  If the Oracle-provided internet and VCN resolver is enabled for the specified subnet, then the domain name for the subnet is used (do not provide one). Otherwise, provide a valid DNS domain name.  Hyphens (-) are not permitted. Applies to Exadata Database Service on Exascale Infrastructure only. 
				* `enabled_ecpu_count` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The number of ECPUs to enable for an Exadata VM cluster on Exascale Infrastructure. 
				* `is_diagnostics_events_enabled` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) Indicates whether diagnostic collection is enabled for the VM cluster.  Enabling diagnostic collection allows you to receive Events service notifications for guest VM issues.  Diagnostic collection also allows Oracle to provide enhanced service and proactive support for your Exadata system.  You can enable diagnostic collection during VM cluster provisioning. You can also disable or enable it at any time using the `UpdateVmCluster` API. 
				* `is_health_monitoring_enabled` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) Indicates whether health monitoring is enabled for the VM cluster.  Enabling health monitoring allows Oracle to collect diagnostic data and share it with its operations and support personnel.  You may also receive notifications for some events. Collecting health diagnostics enables Oracle to provide proactive support and enhanced service for your system. Optionally enable health monitoring while provisioning a system. You can also disable or enable health monitoring anytime using the `UpdateVmCluster` API. 
				* `is_incident_logs_enabled` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) Indicates whether incident logs and trace collection are enabled for the VM cluster.  Enabling incident logs collection allows Oracle to receive Events service notifications for guest VM issues, collect incident logs and traces, and use them to diagnose issues and resolve them. Optionally enable incident logs collection while provisioning a system. You can also disable or enable incident logs collection anytime using the `UpdateVmCluster` API. 
				* `license_model` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The Oracle license model that applies to the Exadata VM cluster on Exascale Infrastructure. The default is BRING_YOUR_OWN_LICENSE. 
				* `nsg_ids` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The list of [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the network security groups (NSGs) to which this resource belongs.  Setting this to an empty list removes all resources from all NSGs.  For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). 
				* `private_zone_id` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The private zone ID in which you want DNS records to be created. 
				* `ssh_public_keys` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The public key portion of one or more key pairs used for SSH access to the Exadata VM cluster on Exascale Infrastructure.
				* `subnet_id` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet associated with the Exadata VM cluster on Exascale Infrastructure. 
				* `subscription_id` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription with which resource needs to be associated with.
				* `total_ecpu_count` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The number of Total ECPUs for an Exadata VM cluster on Exascale Infrastructure. 
				* `vm_file_system_storage_size` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) File System Storage Size in GBs for Exadata VM cluster. 
			* `vm_cluster_id` - (Required when source=EXISTING_DB_HOME | XD_EXISTING_CLUSTER | XS_EXISTING_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Cloud VM Cluster.
		* `name` - (Required) The name of raft cluster.  It must start with a letter and contain only letters, digits, and underscores.  Maximum length is 40 characters. 
		* `replication_factor` - (Required) Replication factor associated with the raft cluster.
		* `ru_mode` - (Required) Replication factor associated with the raft cluster.
		* `witness_count` - (Required) Witnesses RU count associated with the raft cluster.
	* `replication_unit_count` - (Required) Number of replication units in the shard space.
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
		* `remote_region` - (Optional) The name of the remote region where the remote automatic incremental backups will be stored. For information about valid region names, see [Regions and Availability Domains](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/regions.htm). 
		* `type` - (Required) Type of the database backup destination.
	* `can_run_immediate_full_backup` - (Optional) If set to true, configures automatic full backups in the local region (the region of the DB system) for the first backup run immediately.
	* `is_auto_backup_enabled` - (Optional) If set to true, configures automatic backups. If you previously used RMAN or dbcli to configure backups and then you switch to using the Console or the API for backups, a new backup configuration is created and associated with your database. This means that you can no longer rely on your previously configured unmanaged backups to work.
	* `is_remote_backup_enabled` - (Optional) If set to true, configures automatic incremental backups in the local region (the region of the DB system) and the remote region with a default frequency of 1 hour. If you previously used RMAN or dbcli to configure backups, using the Console or the API for manged backups creates a new backup configuration for your database. The new configuration replaces the configuration created with RMAN or dbcli. This means that you can no longer rely on your previously configured unmanaged backups to work. 
	* `recovery_window_in_days` - (Optional) Number of days between the current and the earliest point of recoverability covered by automatic backups. This value applies to automatic backups only. After a new automatic backup has been created, Oracle removes old automatic backups that are created before the window. When the value is updated, it is applied to all existing automatic backups. 
	* `remote_region` - (Optional) The name of the remote region where the remote automatic incremental backups will be stored. For information about valid region names, see [Regions and Availability Domains](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/regions.htm). 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Required) (Updatable) The display name of the Globally distributed database.
* `distributed_database_id` - (Required) 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `gsm_ssh_public_key` - (Optional) The SSH public key for Global service manager instances.
* `listener_port` - (Required) The listener port number for the Globally distributed database. The listener port number has to be unique for a customer tenancy across all distributed databases. Same port number should not be re-used for any other distributed database. 
* `listener_port_tls` - (Optional) The TLS listener port number for the Globally distributed database. The TLS listener port number has to be unique for a customer tenancy across all distributed databases. Same port number should not be re-used for any other distributed database. For EXADB based distributed databases, tls is not supported hence the listenerPortTls is not needed to be provided in create payload. 
* `ncharacter_set` - (Required) The national character set for the database.
* `notification_topic_ids` - (Optional) (Updatable) The collection of [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the notification topics associated with the globally distributed database.
* `ons_port_local` - (Required) The ons local port number for the Globally distributed database. The onsPortLocal has to be unique for a customer tenancy across all distributed databases. Same port number should not be re-used for any other distributed database. 
* `ons_port_remote` - (Required) The ons remote port number for the Globally distributed database. The onsPortRemote has to be unique for a customer tenancy across all distributed databases. Same port number should not be re-used for any other distributed database. 
* `patch_operations` - (Optional) Legacy schema retained for compatibility; this field does not execute topology actions on the main resource. Use the separate INSERT, REMOVE, INSERTDG, and REMOVEDG resources documented above.
	* `operation` - (Required) (Updatable) The operation can be one of these values: `INSERT`, `MERGE`, `REMOVE`
	* `selection` - (Required) (Updatable) 
	* `value` - (Required when operation=INSERT | MERGE) (Updatable) 
* `prefix` - (Required) Unique name prefix for the Globally distributed databases. Only alpha-numeric values are allowed. First character has to be a letter followed by any combination of letter and number. 
* `private_endpoint_ids` - (Required) The collection of [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the private endpoints associated with the Globally distributed database.
* `scan_listener_port` - (Optional) TCP SCAN listener port for new clusters to be created. Applicable only when shard and catalog source types are XS_NEW_VAULT_AND_CLUSTER or XS_NEW_CLUSTER. Defaults to 1521 when not provided. 
* `system_chunk_count` - (Optional) Count of chunks associated with system raft clusters or system data guard databases.
* `system_data_guard_databases` - (Optional) Details required to create system data guard shard.
	* `data_guard_replicas` - (Optional) The details required to create system data guard replicas.
		* `databases` - (Required) Details of the databases associated with the system data guard replica.
			* `admin_password` - (Optional) The admin password for the shard associated with Globally distributed database.
			* `admin_password_secret_id` - (Optional) The Oracle Cloud Infrastructure vault secret [/Content/General/Concepts/identifiers.htm]OCID. This cannot be used in conjunction with adminPassword.
			* `admin_password_secret_version_number` - (Optional) The version of the vault secret. If no version is specified, the latest version will be used.
			* `availability_domain` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The name of the availability domain that the distributed database shard will be located in  like yKoY:US-ASHBURN-AD-3.
			* `db_home_id` - (Required when source=EXISTING_DB_HOME) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Home.
			* `db_storage_vault_details` - (Required when source=XS_NEW_VAULT_AND_CLUSTER) Details of the request to create exascale db vault storage for shard or catalog of the distributed database. 
				* `additional_flash_cache_in_percent` - (Applicable when source=XS_NEW_VAULT_AND_CLUSTER) The size of additional Flash Cache in percentage of High Capacity database storage.
				* `autoscale_limit_in_gbs` - (Applicable when source=XS_NEW_VAULT_AND_CLUSTER) The maximum limit, in gigabytes, to which the Vault storage size can automatically scale when auto scaling is enabled for the Database Storage Vault.
				* `compartment_id` - (Required when source=XS_NEW_VAULT_AND_CLUSTER) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for Db Storage Vault.
				* `high_capacity_database_storage` - (Required when source=XS_NEW_VAULT_AND_CLUSTER) Total storage capacity in GB for vault storage.
				* `is_autoscale_enabled` - (Applicable when source=XS_NEW_VAULT_AND_CLUSTER) Indicates if autoscale feature is enabled for the Database Storage Vault. The default value is `FALSE`.
				* `subscription_id` - (Applicable when source=XS_NEW_VAULT_AND_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription with which resource needs to be associated with.
			* `exascale_db_storage_vault_id` - (Required when source=XS_NEW_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the exascale db storage vault.
			* `kms_key_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations. 
			* `kms_key_version_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. 
			* `protection_mode` - (Optional) The protection mode used for the Data Guard association.
			* `source` - (Required) The source of the Globally distributed database type. Use XS_EXISTING_CLUSTER for a database on Exascale Infrastructure that uses a pre-existing cluster. Use XS_NEW_CLUSTER for a database on Exascale Infrastructure that requires a new cluster created from scratch. Use XS_NEW_VAULT_AND_CLUSTER for a database on Exascale Infrastructure that requires a new vault and cluster created from scratch. Use XD_EXISTING_CLUSTER for a database on Dedicated Infrastructure that uses a pre-existing cluster. Use EXISTING_DB_HOME for a database on a pre-existing database home. 
			* `tde_wallet_password` - (Optional) The password to open the TDE wallet. The password must be 9 to 255 characters and contain at least two uppercase letters, two lowercase letters, two numeric characters, and two special characters. The allowed special characters are _, \#, and -. 
			* `tde_wallet_password_secret_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure vault secret. This cannot be used in conjunction with tdeWalletPassword.
			* `tde_wallet_password_secret_version_number` - (Optional) The version of the vault secret. If no version is specified, the latest version will be used.
			* `transport_type` - (Optional) The transport type used for the Data Guard association.
			* `vault_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts).  This parameter and `kmsKeyId` are required for Customer Managed Keys. 
			* `vm_cluster_details` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) Details of the request to create exadb vm cluster for shard or catalog of the distributed database. 
				* `backup_network_nsg_ids` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) A list of the [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that the backup network of this DB system belongs to.  Setting this to an empty array after the list is created removes the resource from all NSGs.  For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). Applicable only to Exadata systems. 
				* `backup_subnet_id` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup network subnet associated with the Exadata VM cluster on Exascale Infrastructure. 
				* `compartment_id` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for VM Cluster.
				* `domain` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) A domain name used for the Exadata VM cluster on Exascale Infrastructure.  If the Oracle-provided internet and VCN resolver is enabled for the specified subnet, then the domain name for the subnet is used (do not provide one). Otherwise, provide a valid DNS domain name.  Hyphens (-) are not permitted. Applies to Exadata Database Service on Exascale Infrastructure only. 
				* `enabled_ecpu_count` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The number of ECPUs to enable for an Exadata VM cluster on Exascale Infrastructure. 
				* `is_diagnostics_events_enabled` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) Indicates whether diagnostic collection is enabled for the VM cluster.  Enabling diagnostic collection allows you to receive Events service notifications for guest VM issues.  Diagnostic collection also allows Oracle to provide enhanced service and proactive support for your Exadata system.  You can enable diagnostic collection during VM cluster provisioning. You can also disable or enable it at any time using the `UpdateVmCluster` API. 
				* `is_health_monitoring_enabled` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) Indicates whether health monitoring is enabled for the VM cluster.  Enabling health monitoring allows Oracle to collect diagnostic data and share it with its operations and support personnel.  You may also receive notifications for some events. Collecting health diagnostics enables Oracle to provide proactive support and enhanced service for your system. Optionally enable health monitoring while provisioning a system. You can also disable or enable health monitoring anytime using the `UpdateVmCluster` API. 
				* `is_incident_logs_enabled` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) Indicates whether incident logs and trace collection are enabled for the VM cluster.  Enabling incident logs collection allows Oracle to receive Events service notifications for guest VM issues, collect incident logs and traces, and use them to diagnose issues and resolve them. Optionally enable incident logs collection while provisioning a system. You can also disable or enable incident logs collection anytime using the `UpdateVmCluster` API. 
				* `license_model` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The Oracle license model that applies to the Exadata VM cluster on Exascale Infrastructure. The default is BRING_YOUR_OWN_LICENSE. 
				* `nsg_ids` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The list of [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the network security groups (NSGs) to which this resource belongs.  Setting this to an empty list removes all resources from all NSGs.  For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). 
				* `private_zone_id` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The private zone ID in which you want DNS records to be created. 
				* `ssh_public_keys` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The public key portion of one or more key pairs used for SSH access to the Exadata VM cluster on Exascale Infrastructure.
				* `subnet_id` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet associated with the Exadata VM cluster on Exascale Infrastructure. 
				* `subscription_id` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription with which resource needs to be associated with.
				* `total_ecpu_count` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The number of Total ECPUs for an Exadata VM cluster on Exascale Infrastructure. 
				* `vm_file_system_storage_size` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) File System Storage Size in GBs for Exadata VM cluster. 
			* `vm_cluster_id` - (Required when source=EXISTING_DB_HOME | XD_EXISTING_CLUSTER | XS_EXISTING_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Cloud VM Cluster.
		* `name` - (Required) The name of data guard replica.  It must start with a letter and contain only letters, digits, and underscores.  Maximum length is 40 characters. 
	* `original_replica` - (Required) Details required to create system data guard replica.
		* `databases` - (Required) Details of the databases associated with the system data guard replica.
			* `admin_password` - (Optional) The admin password for the shard associated with Globally distributed database.
			* `admin_password_secret_id` - (Optional) The Oracle Cloud Infrastructure vault secret [/Content/General/Concepts/identifiers.htm]OCID. This cannot be used in conjunction with adminPassword.
			* `admin_password_secret_version_number` - (Optional) The version of the vault secret. If no version is specified, the latest version will be used.
			* `availability_domain` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The name of the availability domain that the distributed database shard will be located in  like yKoY:US-ASHBURN-AD-3.
			* `db_home_id` - (Required when source=EXISTING_DB_HOME) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Home.
			* `db_storage_vault_details` - (Required when source=XS_NEW_VAULT_AND_CLUSTER) Details of the request to create exascale db vault storage for shard or catalog of the distributed database. 
				* `additional_flash_cache_in_percent` - (Applicable when source=XS_NEW_VAULT_AND_CLUSTER) The size of additional Flash Cache in percentage of High Capacity database storage.
				* `autoscale_limit_in_gbs` - (Applicable when source=XS_NEW_VAULT_AND_CLUSTER) The maximum limit, in gigabytes, to which the Vault storage size can automatically scale when auto scaling is enabled for the Database Storage Vault.
				* `compartment_id` - (Required when source=XS_NEW_VAULT_AND_CLUSTER) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for Db Storage Vault.
				* `high_capacity_database_storage` - (Required when source=XS_NEW_VAULT_AND_CLUSTER) Total storage capacity in GB for vault storage.
				* `is_autoscale_enabled` - (Applicable when source=XS_NEW_VAULT_AND_CLUSTER) Indicates if autoscale feature is enabled for the Database Storage Vault. The default value is `FALSE`.
				* `subscription_id` - (Applicable when source=XS_NEW_VAULT_AND_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription with which resource needs to be associated with.
			* `exascale_db_storage_vault_id` - (Required when source=XS_NEW_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the exascale db storage vault.
			* `kms_key_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations. 
			* `kms_key_version_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. 
			* `protection_mode` - (Optional) The protection mode used for the Data Guard association.
			* `source` - (Required) The source of the Globally distributed database type. Use XS_EXISTING_CLUSTER for a database on Exascale Infrastructure that uses a pre-existing cluster. Use XS_NEW_CLUSTER for a database on Exascale Infrastructure that requires a new cluster created from scratch. Use XS_NEW_VAULT_AND_CLUSTER for a database on Exascale Infrastructure that requires a new vault and cluster created from scratch. Use XD_EXISTING_CLUSTER for a database on Dedicated Infrastructure that uses a pre-existing cluster. Use EXISTING_DB_HOME for a database on a pre-existing database home. 
			* `tde_wallet_password` - (Optional) The password to open the TDE wallet. The password must be 9 to 255 characters and contain at least two uppercase letters, two lowercase letters, two numeric characters, and two special characters. The allowed special characters are _, \#, and -. 
			* `tde_wallet_password_secret_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure vault secret. This cannot be used in conjunction with tdeWalletPassword.
			* `tde_wallet_password_secret_version_number` - (Optional) The version of the vault secret. If no version is specified, the latest version will be used.
			* `transport_type` - (Optional) The transport type used for the Data Guard association.
			* `vault_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts).  This parameter and `kmsKeyId` are required for Customer Managed Keys. 
			* `vm_cluster_details` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) Details of the request to create exadb vm cluster for shard or catalog of the distributed database. 
				* `backup_network_nsg_ids` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) A list of the [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that the backup network of this DB system belongs to.  Setting this to an empty array after the list is created removes the resource from all NSGs.  For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). Applicable only to Exadata systems. 
				* `backup_subnet_id` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup network subnet associated with the Exadata VM cluster on Exascale Infrastructure. 
				* `compartment_id` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for VM Cluster.
				* `domain` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) A domain name used for the Exadata VM cluster on Exascale Infrastructure.  If the Oracle-provided internet and VCN resolver is enabled for the specified subnet, then the domain name for the subnet is used (do not provide one). Otherwise, provide a valid DNS domain name.  Hyphens (-) are not permitted. Applies to Exadata Database Service on Exascale Infrastructure only. 
				* `enabled_ecpu_count` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The number of ECPUs to enable for an Exadata VM cluster on Exascale Infrastructure. 
				* `is_diagnostics_events_enabled` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) Indicates whether diagnostic collection is enabled for the VM cluster.  Enabling diagnostic collection allows you to receive Events service notifications for guest VM issues.  Diagnostic collection also allows Oracle to provide enhanced service and proactive support for your Exadata system.  You can enable diagnostic collection during VM cluster provisioning. You can also disable or enable it at any time using the `UpdateVmCluster` API. 
				* `is_health_monitoring_enabled` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) Indicates whether health monitoring is enabled for the VM cluster.  Enabling health monitoring allows Oracle to collect diagnostic data and share it with its operations and support personnel.  You may also receive notifications for some events. Collecting health diagnostics enables Oracle to provide proactive support and enhanced service for your system. Optionally enable health monitoring while provisioning a system. You can also disable or enable health monitoring anytime using the `UpdateVmCluster` API. 
				* `is_incident_logs_enabled` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) Indicates whether incident logs and trace collection are enabled for the VM cluster.  Enabling incident logs collection allows Oracle to receive Events service notifications for guest VM issues, collect incident logs and traces, and use them to diagnose issues and resolve them. Optionally enable incident logs collection while provisioning a system. You can also disable or enable incident logs collection anytime using the `UpdateVmCluster` API. 
				* `license_model` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The Oracle license model that applies to the Exadata VM cluster on Exascale Infrastructure. The default is BRING_YOUR_OWN_LICENSE. 
				* `nsg_ids` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The list of [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the network security groups (NSGs) to which this resource belongs.  Setting this to an empty list removes all resources from all NSGs.  For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). 
				* `private_zone_id` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The private zone ID in which you want DNS records to be created. 
				* `ssh_public_keys` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The public key portion of one or more key pairs used for SSH access to the Exadata VM cluster on Exascale Infrastructure.
				* `subnet_id` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet associated with the Exadata VM cluster on Exascale Infrastructure. 
				* `subscription_id` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription with which resource needs to be associated with.
				* `total_ecpu_count` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The number of Total ECPUs for an Exadata VM cluster on Exascale Infrastructure. 
				* `vm_file_system_storage_size` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) File System Storage Size in GBs for Exadata VM cluster. 
			* `vm_cluster_id` - (Required when source=EXISTING_DB_HOME | XD_EXISTING_CLUSTER | XS_EXISTING_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Cloud VM Cluster.
		* `name` - (Required) The name of data guard replica.  It must start with a letter and contain only letters, digits, and underscores.  Maximum length is 40 characters. 
* `system_raft_clusters` - (Optional) Collection of system raft clusters.
	* `databases` - (Required) Details of the databases associated with the system raft cluster.
		* `admin_password` - (Optional) The admin password for the shard associated with Globally distributed database.
		* `admin_password_secret_id` - (Optional) The Oracle Cloud Infrastructure vault secret [/Content/General/Concepts/identifiers.htm]OCID. This cannot be used in conjunction with adminPassword.
		* `admin_password_secret_version_number` - (Optional) The version of the vault secret. If no version is specified, the latest version will be used.
		* `availability_domain` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The name of the availability domain that the distributed database shard will be located in  like yKoY:US-ASHBURN-AD-3.
		* `db_home_id` - (Required when source=EXISTING_DB_HOME) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Home.
		* `db_storage_vault_details` - (Required when source=XS_NEW_VAULT_AND_CLUSTER) Details of the request to create exascale db vault storage for shard or catalog of the distributed database. 
			* `additional_flash_cache_in_percent` - (Applicable when source=XS_NEW_VAULT_AND_CLUSTER) The size of additional Flash Cache in percentage of High Capacity database storage.
			* `autoscale_limit_in_gbs` - (Applicable when source=XS_NEW_VAULT_AND_CLUSTER) The maximum limit, in gigabytes, to which the Vault storage size can automatically scale when auto scaling is enabled for the Database Storage Vault.
			* `compartment_id` - (Required when source=XS_NEW_VAULT_AND_CLUSTER) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for Db Storage Vault.
			* `high_capacity_database_storage` - (Required when source=XS_NEW_VAULT_AND_CLUSTER) Total storage capacity in GB for vault storage.
			* `is_autoscale_enabled` - (Applicable when source=XS_NEW_VAULT_AND_CLUSTER) Indicates if autoscale feature is enabled for the Database Storage Vault. The default value is `FALSE`.
			* `subscription_id` - (Applicable when source=XS_NEW_VAULT_AND_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription with which resource needs to be associated with.
		* `exascale_db_storage_vault_id` - (Required when source=XS_NEW_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the exascale db storage vault.
		* `kms_key_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations. 
		* `kms_key_version_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. 
		* `protection_mode` - (Optional) The protection mode used for the Data Guard association.
		* `source` - (Required) The source of the Globally distributed database type. Use XS_EXISTING_CLUSTER for a database on Exascale Infrastructure that uses a pre-existing cluster. Use XS_NEW_CLUSTER for a database on Exascale Infrastructure that requires a new cluster created from scratch. Use XS_NEW_VAULT_AND_CLUSTER for a database on Exascale Infrastructure that requires a new vault and cluster created from scratch. Use XD_EXISTING_CLUSTER for a database on Dedicated Infrastructure that uses a pre-existing cluster. Use EXISTING_DB_HOME for a database on a pre-existing database home. 
		* `tde_wallet_password` - (Optional) The password to open the TDE wallet. The password must be 9 to 255 characters and contain at least two uppercase letters, two lowercase letters, two numeric characters, and two special characters. The allowed special characters are _, \#, and -. 
		* `tde_wallet_password_secret_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure vault secret. This cannot be used in conjunction with tdeWalletPassword.
		* `tde_wallet_password_secret_version_number` - (Optional) The version of the vault secret. If no version is specified, the latest version will be used.
		* `transport_type` - (Optional) The transport type used for the Data Guard association.
		* `vault_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts).  This parameter and `kmsKeyId` are required for Customer Managed Keys. 
		* `vm_cluster_details` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) Details of the request to create exadb vm cluster for shard or catalog of the distributed database. 
			* `backup_network_nsg_ids` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) A list of the [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that the backup network of this DB system belongs to.  Setting this to an empty array after the list is created removes the resource from all NSGs.  For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). Applicable only to Exadata systems. 
			* `backup_subnet_id` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup network subnet associated with the Exadata VM cluster on Exascale Infrastructure. 
			* `compartment_id` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for VM Cluster.
			* `domain` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) A domain name used for the Exadata VM cluster on Exascale Infrastructure.  If the Oracle-provided internet and VCN resolver is enabled for the specified subnet, then the domain name for the subnet is used (do not provide one). Otherwise, provide a valid DNS domain name.  Hyphens (-) are not permitted. Applies to Exadata Database Service on Exascale Infrastructure only. 
			* `enabled_ecpu_count` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The number of ECPUs to enable for an Exadata VM cluster on Exascale Infrastructure. 
			* `is_diagnostics_events_enabled` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) Indicates whether diagnostic collection is enabled for the VM cluster.  Enabling diagnostic collection allows you to receive Events service notifications for guest VM issues.  Diagnostic collection also allows Oracle to provide enhanced service and proactive support for your Exadata system.  You can enable diagnostic collection during VM cluster provisioning. You can also disable or enable it at any time using the `UpdateVmCluster` API. 
			* `is_health_monitoring_enabled` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) Indicates whether health monitoring is enabled for the VM cluster.  Enabling health monitoring allows Oracle to collect diagnostic data and share it with its operations and support personnel.  You may also receive notifications for some events. Collecting health diagnostics enables Oracle to provide proactive support and enhanced service for your system. Optionally enable health monitoring while provisioning a system. You can also disable or enable health monitoring anytime using the `UpdateVmCluster` API. 
			* `is_incident_logs_enabled` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) Indicates whether incident logs and trace collection are enabled for the VM cluster.  Enabling incident logs collection allows Oracle to receive Events service notifications for guest VM issues, collect incident logs and traces, and use them to diagnose issues and resolve them. Optionally enable incident logs collection while provisioning a system. You can also disable or enable incident logs collection anytime using the `UpdateVmCluster` API. 
			* `license_model` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The Oracle license model that applies to the Exadata VM cluster on Exascale Infrastructure. The default is BRING_YOUR_OWN_LICENSE. 
			* `nsg_ids` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The list of [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the network security groups (NSGs) to which this resource belongs.  Setting this to an empty list removes all resources from all NSGs.  For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). 
			* `private_zone_id` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The private zone ID in which you want DNS records to be created. 
			* `ssh_public_keys` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The public key portion of one or more key pairs used for SSH access to the Exadata VM cluster on Exascale Infrastructure.
			* `subnet_id` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet associated with the Exadata VM cluster on Exascale Infrastructure. 
			* `subscription_id` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription with which resource needs to be associated with.
			* `total_ecpu_count` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The number of Total ECPUs for an Exadata VM cluster on Exascale Infrastructure. 
			* `vm_file_system_storage_size` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) File System Storage Size in GBs for Exadata VM cluster. 
		* `vm_cluster_id` - (Required when source=EXISTING_DB_HOME | XD_EXISTING_CLUSTER | XS_EXISTING_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Cloud VM Cluster.
	* `name` - (Required) The name of raft cluster.  It must start with a letter and contain only letters, digits, and underscores.  Maximum length is 40 characters. 
	* `replication_factor` - (Required) Replication factor associated with the raft cluster.
	* `ru_mode` - (Required) Replication factor associated with the raft cluster.
	* `witness_count` - (Required) Witnesses RU count associated with the raft cluster.
* `system_raft_replication_unit_count` - (Optional) Number of replication units associated with system raft clusters.
* `user_shard_spaces` - (Optional) Collection of user defined shard spaces.
	* `data_guard_replicas` - (Optional) The details of data guard replica for the user shard space being created.
		* `database` - (Required) Details for creating a distributed database shard.
			* `admin_password` - (Optional) The admin password for the shard associated with Globally distributed database.
			* `admin_password_secret_id` - (Optional) The Oracle Cloud Infrastructure vault secret [/Content/General/Concepts/identifiers.htm]OCID. This cannot be used in conjunction with adminPassword.
			* `admin_password_secret_version_number` - (Optional) The version of the vault secret. If no version is specified, the latest version will be used.
			* `availability_domain` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The name of the availability domain that the distributed database shard will be located in  like yKoY:US-ASHBURN-AD-3.
			* `db_home_id` - (Required when source=EXISTING_DB_HOME) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Home.
			* `db_storage_vault_details` - (Required when source=XS_NEW_VAULT_AND_CLUSTER) Details of the request to create exascale db vault storage for shard or catalog of the distributed database. 
				* `additional_flash_cache_in_percent` - (Applicable when source=XS_NEW_VAULT_AND_CLUSTER) The size of additional Flash Cache in percentage of High Capacity database storage.
				* `autoscale_limit_in_gbs` - (Applicable when source=XS_NEW_VAULT_AND_CLUSTER) The maximum limit, in gigabytes, to which the Vault storage size can automatically scale when auto scaling is enabled for the Database Storage Vault.
				* `compartment_id` - (Required when source=XS_NEW_VAULT_AND_CLUSTER) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for Db Storage Vault.
				* `high_capacity_database_storage` - (Required when source=XS_NEW_VAULT_AND_CLUSTER) Total storage capacity in GB for vault storage.
				* `is_autoscale_enabled` - (Applicable when source=XS_NEW_VAULT_AND_CLUSTER) Indicates if autoscale feature is enabled for the Database Storage Vault. The default value is `FALSE`.
				* `subscription_id` - (Applicable when source=XS_NEW_VAULT_AND_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription with which resource needs to be associated with.
			* `exascale_db_storage_vault_id` - (Required when source=XS_NEW_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the exascale db storage vault.
			* `kms_key_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations. 
			* `kms_key_version_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. 
			* `protection_mode` - (Optional) The protection mode used for the Data Guard association.
			* `source` - (Required) The source of the Globally distributed database type. Use XS_EXISTING_CLUSTER for a database on Exascale Infrastructure that uses a pre-existing cluster. Use XS_NEW_CLUSTER for a database on Exascale Infrastructure that requires a new cluster created from scratch. Use XS_NEW_VAULT_AND_CLUSTER for a database on Exascale Infrastructure that requires a new vault and cluster created from scratch. Use XD_EXISTING_CLUSTER for a database on Dedicated Infrastructure that uses a pre-existing cluster. Use EXISTING_DB_HOME for a database on a pre-existing database home. 
			* `tde_wallet_password` - (Optional) The password to open the TDE wallet. The password must be 9 to 255 characters and contain at least two uppercase letters, two lowercase letters, two numeric characters, and two special characters. The allowed special characters are _, \#, and -. 
			* `tde_wallet_password_secret_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure vault secret. This cannot be used in conjunction with tdeWalletPassword.
			* `tde_wallet_password_secret_version_number` - (Optional) The version of the vault secret. If no version is specified, the latest version will be used.
			* `transport_type` - (Optional) The transport type used for the Data Guard association.
			* `vault_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts).  This parameter and `kmsKeyId` are required for Customer Managed Keys. 
			* `vm_cluster_details` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) Details of the request to create exadb vm cluster for shard or catalog of the distributed database. 
				* `backup_network_nsg_ids` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) A list of the [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that the backup network of this DB system belongs to.  Setting this to an empty array after the list is created removes the resource from all NSGs.  For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). Applicable only to Exadata systems. 
				* `backup_subnet_id` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup network subnet associated with the Exadata VM cluster on Exascale Infrastructure. 
				* `compartment_id` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for VM Cluster.
				* `domain` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) A domain name used for the Exadata VM cluster on Exascale Infrastructure.  If the Oracle-provided internet and VCN resolver is enabled for the specified subnet, then the domain name for the subnet is used (do not provide one). Otherwise, provide a valid DNS domain name.  Hyphens (-) are not permitted. Applies to Exadata Database Service on Exascale Infrastructure only. 
				* `enabled_ecpu_count` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The number of ECPUs to enable for an Exadata VM cluster on Exascale Infrastructure. 
				* `is_diagnostics_events_enabled` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) Indicates whether diagnostic collection is enabled for the VM cluster.  Enabling diagnostic collection allows you to receive Events service notifications for guest VM issues.  Diagnostic collection also allows Oracle to provide enhanced service and proactive support for your Exadata system.  You can enable diagnostic collection during VM cluster provisioning. You can also disable or enable it at any time using the `UpdateVmCluster` API. 
				* `is_health_monitoring_enabled` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) Indicates whether health monitoring is enabled for the VM cluster.  Enabling health monitoring allows Oracle to collect diagnostic data and share it with its operations and support personnel.  You may also receive notifications for some events. Collecting health diagnostics enables Oracle to provide proactive support and enhanced service for your system. Optionally enable health monitoring while provisioning a system. You can also disable or enable health monitoring anytime using the `UpdateVmCluster` API. 
				* `is_incident_logs_enabled` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) Indicates whether incident logs and trace collection are enabled for the VM cluster.  Enabling incident logs collection allows Oracle to receive Events service notifications for guest VM issues, collect incident logs and traces, and use them to diagnose issues and resolve them. Optionally enable incident logs collection while provisioning a system. You can also disable or enable incident logs collection anytime using the `UpdateVmCluster` API. 
				* `license_model` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The Oracle license model that applies to the Exadata VM cluster on Exascale Infrastructure. The default is BRING_YOUR_OWN_LICENSE. 
				* `nsg_ids` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The list of [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the network security groups (NSGs) to which this resource belongs.  Setting this to an empty list removes all resources from all NSGs.  For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). 
				* `private_zone_id` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The private zone ID in which you want DNS records to be created. 
				* `ssh_public_keys` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The public key portion of one or more key pairs used for SSH access to the Exadata VM cluster on Exascale Infrastructure.
				* `subnet_id` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet associated with the Exadata VM cluster on Exascale Infrastructure. 
				* `subscription_id` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription with which resource needs to be associated with.
				* `total_ecpu_count` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The number of Total ECPUs for an Exadata VM cluster on Exascale Infrastructure. 
				* `vm_file_system_storage_size` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) File System Storage Size in GBs for Exadata VM cluster. 
			* `vm_cluster_id` - (Required when source=EXISTING_DB_HOME | XD_EXISTING_CLUSTER | XS_EXISTING_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Cloud VM Cluster.
		* `name` - (Required) The name of user shard space replica.  It must start with a letter and contain only letters, digits, and underscores.  Maximum length is 40 characters. 
	* `name` - (Required) The name of user shard space.  It must start with a letter and contain only letters, digits, and underscores.  Maximum length is 40 characters. 
	* `original_replica` - (Required) Details required to create user shard space replica.
		* `database` - (Required) Details for creating a distributed database shard.
			* `admin_password` - (Optional) The admin password for the shard associated with Globally distributed database.
			* `admin_password_secret_id` - (Optional) The Oracle Cloud Infrastructure vault secret [/Content/General/Concepts/identifiers.htm]OCID. This cannot be used in conjunction with adminPassword.
			* `admin_password_secret_version_number` - (Optional) The version of the vault secret. If no version is specified, the latest version will be used.
			* `availability_domain` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The name of the availability domain that the distributed database shard will be located in  like yKoY:US-ASHBURN-AD-3.
			* `db_home_id` - (Required when source=EXISTING_DB_HOME) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Home.
			* `db_storage_vault_details` - (Required when source=XS_NEW_VAULT_AND_CLUSTER) Details of the request to create exascale db vault storage for shard or catalog of the distributed database. 
				* `additional_flash_cache_in_percent` - (Applicable when source=XS_NEW_VAULT_AND_CLUSTER) The size of additional Flash Cache in percentage of High Capacity database storage.
				* `autoscale_limit_in_gbs` - (Applicable when source=XS_NEW_VAULT_AND_CLUSTER) The maximum limit, in gigabytes, to which the Vault storage size can automatically scale when auto scaling is enabled for the Database Storage Vault.
				* `compartment_id` - (Required when source=XS_NEW_VAULT_AND_CLUSTER) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for Db Storage Vault.
				* `high_capacity_database_storage` - (Required when source=XS_NEW_VAULT_AND_CLUSTER) Total storage capacity in GB for vault storage.
				* `is_autoscale_enabled` - (Applicable when source=XS_NEW_VAULT_AND_CLUSTER) Indicates if autoscale feature is enabled for the Database Storage Vault. The default value is `FALSE`.
				* `subscription_id` - (Applicable when source=XS_NEW_VAULT_AND_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription with which resource needs to be associated with.
			* `exascale_db_storage_vault_id` - (Required when source=XS_NEW_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the exascale db storage vault.
			* `kms_key_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations. 
			* `kms_key_version_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. 
			* `protection_mode` - (Optional) The protection mode used for the Data Guard association.
			* `source` - (Required) The source of the Globally distributed database type. Use XS_EXISTING_CLUSTER for a database on Exascale Infrastructure that uses a pre-existing cluster. Use XS_NEW_CLUSTER for a database on Exascale Infrastructure that requires a new cluster created from scratch. Use XS_NEW_VAULT_AND_CLUSTER for a database on Exascale Infrastructure that requires a new vault and cluster created from scratch. Use XD_EXISTING_CLUSTER for a database on Dedicated Infrastructure that uses a pre-existing cluster. Use EXISTING_DB_HOME for a database on a pre-existing database home. 
			* `tde_wallet_password` - (Optional) The password to open the TDE wallet. The password must be 9 to 255 characters and contain at least two uppercase letters, two lowercase letters, two numeric characters, and two special characters. The allowed special characters are _, \#, and -. 
			* `tde_wallet_password_secret_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure vault secret. This cannot be used in conjunction with tdeWalletPassword.
			* `tde_wallet_password_secret_version_number` - (Optional) The version of the vault secret. If no version is specified, the latest version will be used.
			* `transport_type` - (Optional) The transport type used for the Data Guard association.
			* `vault_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts).  This parameter and `kmsKeyId` are required for Customer Managed Keys. 
			* `vm_cluster_details` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) Details of the request to create exadb vm cluster for shard or catalog of the distributed database. 
				* `backup_network_nsg_ids` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) A list of the [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that the backup network of this DB system belongs to.  Setting this to an empty array after the list is created removes the resource from all NSGs.  For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). Applicable only to Exadata systems. 
				* `backup_subnet_id` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup network subnet associated with the Exadata VM cluster on Exascale Infrastructure. 
				* `compartment_id` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for VM Cluster.
				* `domain` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) A domain name used for the Exadata VM cluster on Exascale Infrastructure.  If the Oracle-provided internet and VCN resolver is enabled for the specified subnet, then the domain name for the subnet is used (do not provide one). Otherwise, provide a valid DNS domain name.  Hyphens (-) are not permitted. Applies to Exadata Database Service on Exascale Infrastructure only. 
				* `enabled_ecpu_count` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The number of ECPUs to enable for an Exadata VM cluster on Exascale Infrastructure. 
				* `is_diagnostics_events_enabled` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) Indicates whether diagnostic collection is enabled for the VM cluster.  Enabling diagnostic collection allows you to receive Events service notifications for guest VM issues.  Diagnostic collection also allows Oracle to provide enhanced service and proactive support for your Exadata system.  You can enable diagnostic collection during VM cluster provisioning. You can also disable or enable it at any time using the `UpdateVmCluster` API. 
				* `is_health_monitoring_enabled` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) Indicates whether health monitoring is enabled for the VM cluster.  Enabling health monitoring allows Oracle to collect diagnostic data and share it with its operations and support personnel.  You may also receive notifications for some events. Collecting health diagnostics enables Oracle to provide proactive support and enhanced service for your system. Optionally enable health monitoring while provisioning a system. You can also disable or enable health monitoring anytime using the `UpdateVmCluster` API. 
				* `is_incident_logs_enabled` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) Indicates whether incident logs and trace collection are enabled for the VM cluster.  Enabling incident logs collection allows Oracle to receive Events service notifications for guest VM issues, collect incident logs and traces, and use them to diagnose issues and resolve them. Optionally enable incident logs collection while provisioning a system. You can also disable or enable incident logs collection anytime using the `UpdateVmCluster` API. 
				* `license_model` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The Oracle license model that applies to the Exadata VM cluster on Exascale Infrastructure. The default is BRING_YOUR_OWN_LICENSE. 
				* `nsg_ids` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The list of [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the network security groups (NSGs) to which this resource belongs.  Setting this to an empty list removes all resources from all NSGs.  For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). 
				* `private_zone_id` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The private zone ID in which you want DNS records to be created. 
				* `ssh_public_keys` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The public key portion of one or more key pairs used for SSH access to the Exadata VM cluster on Exascale Infrastructure.
				* `subnet_id` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet associated with the Exadata VM cluster on Exascale Infrastructure. 
				* `subscription_id` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription with which resource needs to be associated with.
				* `total_ecpu_count` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The number of Total ECPUs for an Exadata VM cluster on Exascale Infrastructure. 
				* `vm_file_system_storage_size` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) File System Storage Size in GBs for Exadata VM cluster. 
			* `vm_cluster_id` - (Required when source=EXISTING_DB_HOME | XD_EXISTING_CLUSTER | XS_EXISTING_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Cloud VM Cluster.
		* `name` - (Required) The name of user shard space replica.  It must start with a letter and contain only letters, digits, and underscores.  Maximum length is 40 characters. 
* `vcn_nsg_ids` - (Optional) The list of network security group (NSG) details to be associated with the distributed database.
	* `nsg_ids` - (Required) The list of [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) associated with resources in this VCN. For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). 
	* `vcn_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the customer Virtual Cloud Network.
* `add_gds_control_node_trigger` - (Optional) (Updatable) An optional property when incremented triggers Add Gds Control Node. Could be set to any integer value.
* `change_db_backup_config_trigger` - (Optional) (Updatable) An optional property when incremented triggers Change Db Backup Config. Could be set to any integer value.
* `configure_auto_resource_management_trigger` - (Optional) (Updatable) An optional property when incremented triggers Configure Auto Resource Management. Could be set to any integer value.
* `configure_sharding_trigger` - (Optional) (Updatable) An optional property when incremented triggers Configure Sharding. Could be set to any integer value.
* `generate_wallet_trigger` - (Optional) (Updatable) An optional property when incremented triggers Generate Wallet. Could be set to any integer value.
* `move_replication_unit_trigger` - (Optional) (Updatable) An optional property when incremented triggers Move Replication Unit. Could be set to any integer value.
* `recreate_failed_resource_trigger` - (Optional) (Updatable) An optional property when incremented triggers Recreate Failed Resource. Could be set to any integer value.
* `scale_gsms_trigger` - (Optional) (Updatable) An optional property when incremented triggers Scale Gsms. Could be set to any integer value.
* `operational_state` - (Optional) (Updatable) Set to `AVAILABLE` to start or `STOPPED` to stop an existing database. Omit during creation or to leave start/stop unmanaged. Start runs before other updates; Stop runs after successful updates. Completion is verified using node and catalog metadata from DDB GET: `openMode` (`READ_WRITE`/`MOUNTED`). Parent DDB lifecycle does not control this check. Nodes and catalog replicas explicitly marked `DELETED` are excluded. Missing operational metadata on remaining databases produces `UNKNOWN`. Do not combine this argument with incremented start/stop counters.
* `start_database_trigger` - (Optional) (Updatable) An optional property when incremented triggers Start Database. Could be set to any integer value.
* `stop_database_trigger` - (Optional) (Updatable) An optional property when incremented triggers Stop Database. Could be set to any integer value.
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

