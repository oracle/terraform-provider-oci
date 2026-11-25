# Forward admin password secret IDs and versions for every database block.
resource "oci_distributed_database_distributed_autonomous_database" "this" {
  operational_state    = try(var.operational_state, null)
  compartment_id       = var.compartment_id
  display_name         = var.display_name
  database_version     = try(var.database_version, "26ai")
  db_workload_type     = try(var.db_workload_type, "OLTP")
  character_set        = try(var.character_set, "AL32UTF8")
  ncharacter_set       = try(var.ncharacter_set, "AL16UTF16")
  prefix               = var.prefix
  private_endpoint_ids = var.private_endpoint_ids
  listener_port        = var.listener_port
  listener_port_tls    = var.listener_port_tls
  ons_port_local       = var.ons_port_local
  ons_port_remote      = var.ons_port_remote
  freeform_tags        = try(var.freeform_tags, {})

  dynamic "db_backup_config" {
    for_each = var.db_backup_config[*]
    content {
      recovery_window_in_days = try(db_backup_config.value.recovery_window_in_days, null)

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
          cloud_autonomous_vm_cluster_id       = catalog_details.value.original_replica.database.cloud_autonomous_vm_cluster_id
          compute_count                        = catalog_details.value.original_replica.database.compute_count
          data_storage_size_in_gbs             = catalog_details.value.original_replica.database.data_storage_size_in_gbs
          source                               = catalog_details.value.original_replica.database.source
          is_auto_scaling_enabled              = try(catalog_details.value.original_replica.database.is_auto_scaling_enabled, null)
          vault_id                             = try(catalog_details.value.original_replica.database.vault_id, null)
          kms_key_id                           = try(catalog_details.value.original_replica.database.kms_key_id, null)
          protection_mode                      = try(catalog_details.value.original_replica.database.protection_mode, null)
        }
      }

      dynamic "data_guard_replicas" {
        for_each = catalog_details.value.data_guard_replicas
        content {
          name = data_guard_replicas.value.name
          database {
            cloud_autonomous_vm_cluster_id       = data_guard_replicas.value.database.cloud_autonomous_vm_cluster_id
            compute_count                        = data_guard_replicas.value.database.compute_count
            data_storage_size_in_gbs             = data_guard_replicas.value.database.data_storage_size_in_gbs
            admin_password                       = try(data_guard_replicas.value.database.admin_password, null)
            admin_password_secret_id             = try(data_guard_replicas.value.database.admin_password_secret_id, null)
            admin_password_secret_version_number = try(data_guard_replicas.value.database.admin_password_secret_version_number, null)
            source                               = data_guard_replicas.value.database.source
            is_auto_scaling_enabled              = try(data_guard_replicas.value.database.is_auto_scaling_enabled, null)
            vault_id                             = try(data_guard_replicas.value.database.vault_id, null)
            kms_key_id                           = try(data_guard_replicas.value.database.kms_key_id, null)
            protection_mode                      = try(data_guard_replicas.value.database.protection_mode, null)
          }
        }
      }
    }
  }

  dynamic "user_shard_spaces" {
    for_each = var.user_shard_spaces
    content {
      name = user_shard_spaces.value.name
      original_replica {
        name = user_shard_spaces.value.original_replica.name
        database {
          admin_password                       = try(user_shard_spaces.value.original_replica.database.admin_password, null)
          admin_password_secret_id             = try(user_shard_spaces.value.original_replica.database.admin_password_secret_id, null)
          admin_password_secret_version_number = try(user_shard_spaces.value.original_replica.database.admin_password_secret_version_number, null)
          cloud_autonomous_vm_cluster_id       = user_shard_spaces.value.original_replica.database.cloud_autonomous_vm_cluster_id
          compute_count                        = user_shard_spaces.value.original_replica.database.compute_count
          data_storage_size_in_gbs             = user_shard_spaces.value.original_replica.database.data_storage_size_in_gbs
          source                               = user_shard_spaces.value.original_replica.database.source
          shard_space_name                     = try(user_shard_spaces.value.original_replica.database.shard_space_name, null)
          is_auto_scaling_enabled              = try(user_shard_spaces.value.original_replica.database.is_auto_scaling_enabled, null)
          vault_id                             = try(user_shard_spaces.value.original_replica.database.vault_id, null)
          kms_key_id                           = try(user_shard_spaces.value.original_replica.database.kms_key_id, null)
          protection_mode                      = try(user_shard_spaces.value.original_replica.database.protection_mode, null)
        }
      }

      dynamic "data_guard_replicas" {
        for_each = user_shard_spaces.value.data_guard_replicas
        content {
          name = data_guard_replicas.value.name
          database {
            cloud_autonomous_vm_cluster_id       = data_guard_replicas.value.database.cloud_autonomous_vm_cluster_id
            compute_count                        = data_guard_replicas.value.database.compute_count
            data_storage_size_in_gbs             = data_guard_replicas.value.database.data_storage_size_in_gbs
            admin_password                       = try(data_guard_replicas.value.database.admin_password, null)
            admin_password_secret_id             = try(data_guard_replicas.value.database.admin_password_secret_id, null)
            admin_password_secret_version_number = try(data_guard_replicas.value.database.admin_password_secret_version_number, null)
            source                               = data_guard_replicas.value.database.source
            shard_space_name                     = try(data_guard_replicas.value.database.shard_space_name, null)
            is_auto_scaling_enabled              = try(data_guard_replicas.value.database.is_auto_scaling_enabled, null)
            vault_id                             = try(data_guard_replicas.value.database.vault_id, null)
            kms_key_id                           = try(data_guard_replicas.value.database.kms_key_id, null)
            protection_mode                      = try(data_guard_replicas.value.database.protection_mode, null)
          }
        }
      }
    }
  }
  validate_network_trigger        = try(var.validate_network_trigger, null)
  change_db_backup_config_trigger = try(var.change_db_backup_config_trigger, null)
}

resource "oci_distributed_database_distributed_autonomous_databases_configure_sharding" "this" {
  count = try(var.configure_sharding_enabled, true) ? 1 : 0

  distributed_autonomous_database_id = oci_distributed_database_distributed_autonomous_database.this.id
  certificate_id                     = var.certificate_id
  ca_bundle_id                       = var.ca_bundle_id
  is_rebalance_required              = try(var.configure_sharding_is_rebalance_required, true)
}

data "oci_distributed_database_distributed_autonomous_database" "this" {
  distributed_autonomous_database_id = oci_distributed_database_distributed_autonomous_database.this.id

  depends_on = [
    oci_distributed_database_distributed_autonomous_database.this,
    oci_distributed_database_distributed_autonomous_databases_configure_sharding.this,
  ]
}
