# DG CREATE accepts either plaintext or a Vault secret reference.
# Forward admin password secret IDs and versions for every database block.
resource "oci_distributed_database_distributed_database" "this" {
  operational_state    = try(var.operational_state, null)
  compartment_id       = var.compartment_id
  display_name         = var.display_name
  database_version     = try(var.database_version, "26ai")
  character_set        = try(var.character_set, "AL32UTF8")
  ncharacter_set       = try(var.ncharacter_set, "AL16UTF16")
  prefix               = var.prefix
  private_endpoint_ids = var.private_endpoint_ids
  listener_port        = var.listener_port
  listener_port_tls    = try(var.listener_port_tls, null)
  ons_port_local       = var.ons_port_local
  ons_port_remote      = var.ons_port_remote
  scan_listener_port   = try(var.scan_listener_port, null)
  freeform_tags        = try(var.freeform_tags, {})

  system_chunk_count                 = try(var.system_chunk_count, null)
  system_raft_replication_unit_count = try(var.system_raft_replication_unit_count, null)

  dynamic "db_backup_config" {
    for_each = var.db_backup_config[*]
    content {
      is_auto_backup_enabled        = try(db_backup_config.value.is_auto_backup_enabled, null)
      auto_backup_window            = try(db_backup_config.value.auto_backup_window, null)
      auto_full_backup_day          = try(db_backup_config.value.auto_full_backup_day, null)
      auto_full_backup_window       = try(db_backup_config.value.auto_full_backup_window, null)
      backup_deletion_policy        = try(db_backup_config.value.backup_deletion_policy, null)
      recovery_window_in_days       = try(db_backup_config.value.recovery_window_in_days, null)
      can_run_immediate_full_backup = try(db_backup_config.value.can_run_immediate_full_backup, null)
      is_remote_backup_enabled      = try(db_backup_config.value.is_remote_backup_enabled, null)
      remote_region                 = try(db_backup_config.value.remote_region, null)

      dynamic "backup_destination_details" {
        for_each = db_backup_config.value.backup_destination_details
        content {
          dbrs_policy_id = try(backup_destination_details.value.dbrs_policy_id, null)
          id             = try(backup_destination_details.value.id, null)
          internet_proxy = try(backup_destination_details.value.internet_proxy, null)
          is_remote      = try(backup_destination_details.value.is_remote, null)
          remote_region  = try(backup_destination_details.value.remote_region, null)
          type           = backup_destination_details.value.type
        }
      }
    }
  }

  dynamic "catalog_details" {
    for_each = var.catalog_details
    content {
      original_replica {
        name = catalog_details.value.original_replica.name
        database {
          admin_password                       = try(catalog_details.value.original_replica.database.admin_password, null)
          admin_password_secret_id             = try(catalog_details.value.original_replica.database.admin_password_secret_id, null)
          admin_password_secret_version_number = try(catalog_details.value.original_replica.database.admin_password_secret_version_number, null)
          source                               = catalog_details.value.original_replica.database.source
          vm_cluster_id                        = try(catalog_details.value.original_replica.database.vm_cluster_id, null)
          db_home_id                           = try(catalog_details.value.original_replica.database.db_home_id, null)
          vault_id                             = try(catalog_details.value.original_replica.database.vault_id, null)
          kms_key_id                           = try(catalog_details.value.original_replica.database.kms_key_id, null)
          protection_mode                      = try(catalog_details.value.original_replica.database.protection_mode, null)
          transport_type                       = try(catalog_details.value.original_replica.database.transport_type, null)
        }
      }

      dynamic "data_guard_replicas" {
        for_each = catalog_details.value.data_guard_replicas
        content {
          name = data_guard_replicas.value.name
          database {
            admin_password                       = try(data_guard_replicas.value.database.admin_password, null)
            admin_password_secret_id             = try(data_guard_replicas.value.database.admin_password_secret_id, null)
            admin_password_secret_version_number = try(data_guard_replicas.value.database.admin_password_secret_version_number, null)
            source                               = data_guard_replicas.value.database.source
            vm_cluster_id                        = try(data_guard_replicas.value.database.vm_cluster_id, null)
            db_home_id                           = try(data_guard_replicas.value.database.db_home_id, null)
            vault_id                             = try(data_guard_replicas.value.database.vault_id, null)
            kms_key_id                           = try(data_guard_replicas.value.database.kms_key_id, null)
            protection_mode                      = try(data_guard_replicas.value.database.protection_mode, null)
            transport_type                       = try(data_guard_replicas.value.database.transport_type, null)
          }
        }
      }
    }
  }

  dynamic "composite_data_guard_shard_spaces" {
    for_each = var.composite_data_guard_shard_spaces
    content {
      name        = composite_data_guard_shard_spaces.value.name
      chunk_count = composite_data_guard_shard_spaces.value.chunk_count

      original_replica {
        name = composite_data_guard_shard_spaces.value.original_replica.name
        dynamic "databases" {
          for_each = composite_data_guard_shard_spaces.value.original_replica.databases
          content {
            admin_password                       = try(databases.value.admin_password, null)
            admin_password_secret_id             = try(databases.value.admin_password_secret_id, null)
            admin_password_secret_version_number = try(databases.value.admin_password_secret_version_number, null)
            source                               = databases.value.source
            vm_cluster_id                        = try(databases.value.vm_cluster_id, null)
            db_home_id                           = try(databases.value.db_home_id, null)
            vault_id                             = try(databases.value.vault_id, null)
            kms_key_id                           = try(databases.value.kms_key_id, null)
            protection_mode                      = try(databases.value.protection_mode, null)
            transport_type                       = try(databases.value.transport_type, null)
          }
        }
      }

      dynamic "data_guard_replicas" {
        for_each = composite_data_guard_shard_spaces.value.data_guard_replicas
        content {
          name = data_guard_replicas.value.name
          dynamic "databases" {
            for_each = data_guard_replicas.value.databases
            content {
              admin_password                       = try(databases.value.admin_password, null)
              admin_password_secret_id             = try(databases.value.admin_password_secret_id, null)
              admin_password_secret_version_number = try(databases.value.admin_password_secret_version_number, null)
              source                               = databases.value.source
              vm_cluster_id                        = try(databases.value.vm_cluster_id, null)
              db_home_id                           = try(databases.value.db_home_id, null)
              vault_id                             = try(databases.value.vault_id, null)
              kms_key_id                           = try(databases.value.kms_key_id, null)
              protection_mode                      = try(databases.value.protection_mode, null)
              transport_type                       = try(databases.value.transport_type, null)
            }
          }
        }
      }
    }
  }

  validate_network_trigger = try(var.validate_network_trigger, null)
}

resource "oci_distributed_database_distributed_databases_configure_sharding" "this" {
  count = try(var.configure_sharding_enabled, true) ? 1 : 0

  distributed_database_id = oci_distributed_database_distributed_database.this.id
  is_rebalance_required   = try(var.configure_sharding_is_rebalance_required, true)
}

data "oci_distributed_database_distributed_database" "this" {
  distributed_database_id = oci_distributed_database_distributed_database.this.id

  depends_on = [
    oci_distributed_database_distributed_database.this,
    oci_distributed_database_distributed_databases_configure_sharding.this,
  ]
}
