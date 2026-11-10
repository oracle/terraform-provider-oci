# DG CREATE accepts either plaintext or a Vault secret reference.
# Forward admin password secret IDs and versions for every database block.
resource "oci_distributed_database_distributed_database" "exascale" {
  operational_state                  = try(var.operational_state, null)
  compartment_id                     = var.compartment_id
  database_version                   = try(var.database_version, "26ai")
  display_name                       = var.display_name
  listener_port                      = var.listener_port
  ons_port_local                     = var.ons_port_local
  ons_port_remote                    = var.ons_port_remote
  prefix                             = var.prefix
  character_set                      = try(var.character_set, "AL32UTF8")
  ncharacter_set                     = try(var.ncharacter_set, "AL16UTF16")
  private_endpoint_ids               = var.private_endpoint_ids
  scan_listener_port                 = try(var.scan_listener_port, null)
  freeform_tags                      = try(var.freeform_tags, {})
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

  dynamic "system_raft_clusters" {
    for_each = var.system_raft_clusters
    content {
      name               = system_raft_clusters.value.name
      replication_factor = system_raft_clusters.value.replication_factor
      witness_count      = system_raft_clusters.value.witness_count
      ru_mode            = system_raft_clusters.value.ru_mode

      dynamic "databases" {
        for_each = system_raft_clusters.value.databases
        content {
          db_home_id                           = databases.value.source == "EXISTING_DB_HOME" ? databases.value.db_home_id : null
          source                               = databases.value.source
          vm_cluster_id                        = databases.value.vm_cluster_id
          admin_password                       = try(coalesce(databases.value.admin_password, databases.value.admin_password_secret_id == null ? var.admin_password : null), null)
          admin_password_secret_id             = try(databases.value.admin_password_secret_id, null)
          admin_password_secret_version_number = try(databases.value.admin_password_secret_version_number, null)
          vault_id                             = try(coalesce(databases.value.vault_id, var.vault_id), null)
          kms_key_id                           = try(coalesce(databases.value.kms_key_id, var.kms_key_id), null)
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
          source                               = catalog_details.value.original_replica.database.source
          db_home_id                           = catalog_details.value.original_replica.database.source == "EXISTING_DB_HOME" ? catalog_details.value.original_replica.database.db_home_id : null
          vm_cluster_id                        = catalog_details.value.original_replica.database.vm_cluster_id
          admin_password                       = try(coalesce(catalog_details.value.original_replica.database.admin_password, catalog_details.value.original_replica.database.admin_password_secret_id == null ? var.admin_password : null), null)
          admin_password_secret_id             = try(catalog_details.value.original_replica.database.admin_password_secret_id, null)
          admin_password_secret_version_number = try(catalog_details.value.original_replica.database.admin_password_secret_version_number, null)
          vault_id                             = try(coalesce(catalog_details.value.original_replica.database.vault_id, var.vault_id), null)
          kms_key_id                           = try(coalesce(catalog_details.value.original_replica.database.kms_key_id, var.kms_key_id), null)
        }
      }

      dynamic "data_guard_replicas" {
        for_each = catalog_details.value.data_guard_replicas
        content {
          name = data_guard_replicas.value.name
          database {
            source                               = data_guard_replicas.value.database.source
            db_home_id                           = data_guard_replicas.value.database.source == "EXISTING_DB_HOME" ? data_guard_replicas.value.database.db_home_id : null
            vm_cluster_id                        = data_guard_replicas.value.database.vm_cluster_id
            admin_password                       = try(coalesce(data_guard_replicas.value.database.admin_password, data_guard_replicas.value.database.admin_password_secret_id == null ? var.admin_password : null), null)
            admin_password_secret_id             = try(data_guard_replicas.value.database.admin_password_secret_id, null)
            admin_password_secret_version_number = try(data_guard_replicas.value.database.admin_password_secret_version_number, null)
            # Data Guard inherits the primary database encryption settings.
            vault_id        = try(coalesce(catalog_details.value.original_replica.database.vault_id, var.vault_id), null)
            kms_key_id      = try(coalesce(catalog_details.value.original_replica.database.kms_key_id, var.kms_key_id), null)
            protection_mode = try(data_guard_replicas.value.database.protection_mode, null)
            transport_type  = try(data_guard_replicas.value.database.transport_type, null)
          }
        }
      }
    }
  }

  validate_network_trigger = try(var.validate_network_trigger, null)
}

resource "oci_distributed_database_distributed_databases_configure_sharding" "configure_sharding" {
  count = try(var.configure_sharding_enabled, true) ? 1 : 0

  distributed_database_id = oci_distributed_database_distributed_database.exascale.id
  is_rebalance_required   = try(var.configure_sharding_is_rebalance_required, true)
}

data "oci_distributed_database_distributed_database" "exascale" {
  distributed_database_id = oci_distributed_database_distributed_database.exascale.id

  depends_on = [
    oci_distributed_database_distributed_database.exascale,
    oci_distributed_database_distributed_databases_configure_sharding.configure_sharding,
  ]
}

locals {
  gds_control_node_details       = coalesce(try(data.oci_distributed_database_distributed_database.exascale.gds_control_node_details, null), [])
  first_gds_control_node_details = length(local.gds_control_node_details) == 0 ? null : local.gds_control_node_details[0]
  gds_control_node_private_ip    = try(local.first_gds_control_node_details.private_ip, null)
}
