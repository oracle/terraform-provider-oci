locals {
  gds_control_node_actions = var.gds_control_nodes
}

resource "oci_distributed_database_distributed_databases_add_gds_control_node" "this" {
  count = length(local.gds_control_node_actions)

  distributed_database_id = module.distributed_database.id
  public_ssh_key          = local.gds_control_node_actions[count.index].public_ssh_key
  subnet_id               = try(local.gds_control_node_actions[count.index].subnet_id, null)

  depends_on = [module.distributed_database]
}


locals {
  insert_operations    = var.actions.insert
  insert_dg_operations = var.actions.insert_dg
  remove_dg_operations = var.actions.remove_dg
  remove_operations    = var.actions.remove
}

resource "oci_distributed_database_distributed_databases_insert" "insert" {
  count = length(local.insert_operations) == 0 ? 0 : 1

  distributed_database_id = module.distributed_database.id

  dynamic "insert" {
    for_each = var.actions.insert
    content {
      selection = insert.value.selection
      value     = jsonencode(jsondecode(insert.value.value))
    }
  }

  depends_on = [module.distributed_database]
}

resource "oci_distributed_database_distributed_databases_configure_sharding" "configure_sharding_insert" {
  count = length(local.insert_operations) > 0 ? 1 : 0

  distributed_database_id = module.distributed_database.id
  is_rebalance_required   = try(var.configure_sharding_is_rebalance_required, true)

  depends_on = [oci_distributed_database_distributed_databases_insert.insert]

  lifecycle {
    replace_triggered_by = [oci_distributed_database_distributed_databases_insert.insert]
  }
}



resource "oci_distributed_database_distributed_databases_insert_dg" "insert_dg" {
  count = length(local.insert_dg_operations) == 0 ? 0 : 1

  distributed_database_id = module.distributed_database.id

  dynamic "insert_dg" {
    for_each = var.actions.insert_dg
    content {
      selection    = insert_dg.value.selection
      replica_name = insert_dg.value.replica_name
      database {
        admin_password                       = try(insert_dg.value.database.admin_password, null)
        admin_password_secret_id             = try(insert_dg.value.database.admin_password_secret_id, null)
        admin_password_secret_version_number = try(insert_dg.value.database.admin_password_secret_version_number, null)
        availability_domain                  = try(insert_dg.value.database.availability_domain, null)
        db_home_id                           = try(insert_dg.value.database.db_home_id, null)
        exascale_db_storage_vault_id         = try(insert_dg.value.database.exascale_db_storage_vault_id, null)
        protection_mode                      = try(insert_dg.value.database.protection_mode, null)
        source                               = insert_dg.value.database.source
        transport_type                       = try(insert_dg.value.database.transport_type, null)
        vm_cluster_id                        = try(insert_dg.value.database.vm_cluster_id, null)
        dynamic "db_storage_vault_details" {
          for_each = insert_dg.value.database.db_storage_vault_details[*]
          content {
            additional_flash_cache_in_percent = try(db_storage_vault_details.value.additional_flash_cache_in_percent, null)
            autoscale_limit_in_gbs            = try(db_storage_vault_details.value.autoscale_limit_in_gbs, null)
            compartment_id                    = try(db_storage_vault_details.value.compartment_id, null)
            high_capacity_database_storage    = try(db_storage_vault_details.value.high_capacity_database_storage, null)
            is_autoscale_enabled              = try(db_storage_vault_details.value.is_autoscale_enabled, null)
            subscription_id                   = try(db_storage_vault_details.value.subscription_id, null)
          }
        }
        dynamic "vm_cluster_details" {
          for_each = insert_dg.value.database.vm_cluster_details[*]
          content {
            backup_network_nsg_ids        = try(vm_cluster_details.value.backup_network_nsg_ids, null)
            backup_subnet_id              = try(vm_cluster_details.value.backup_subnet_id, null)
            compartment_id                = try(vm_cluster_details.value.compartment_id, null)
            domain                        = try(vm_cluster_details.value.domain, null)
            enabled_ecpu_count            = try(vm_cluster_details.value.enabled_ecpu_count, null)
            is_diagnostics_events_enabled = try(vm_cluster_details.value.is_diagnostics_events_enabled, null)
            is_health_monitoring_enabled  = try(vm_cluster_details.value.is_health_monitoring_enabled, null)
            is_incident_logs_enabled      = try(vm_cluster_details.value.is_incident_logs_enabled, null)
            license_model                 = try(vm_cluster_details.value.license_model, null)
            nsg_ids                       = try(vm_cluster_details.value.nsg_ids, null)
            private_zone_id               = try(vm_cluster_details.value.private_zone_id, null)
            ssh_public_keys               = try(vm_cluster_details.value.ssh_public_keys, null)
            subnet_id                     = try(vm_cluster_details.value.subnet_id, null)
            subscription_id               = try(vm_cluster_details.value.subscription_id, null)
            total_ecpu_count              = try(vm_cluster_details.value.total_ecpu_count, null)
            vm_file_system_storage_size   = try(vm_cluster_details.value.vm_file_system_storage_size, null)
          }
        }
      }
    }
  }

  depends_on = [module.distributed_database, oci_distributed_database_distributed_databases_configure_sharding.configure_sharding_insert]
}

resource "oci_distributed_database_distributed_databases_remove_dg" "remove_dg" {
  count = length(local.remove_dg_operations) == 0 ? 0 : 1

  distributed_database_id = module.distributed_database.id

  dynamic "remove_dg" {
    for_each = var.actions.remove_dg
    content {
      selection    = remove_dg.value.selection
      replica_name = remove_dg.value.replica_name
    }
  }

  depends_on = [module.distributed_database, oci_distributed_database_distributed_databases_insert_dg.insert_dg]
}

resource "oci_distributed_database_distributed_databases_remove" "remove" {
  count = length(local.remove_operations) == 0 ? 0 : 1

  distributed_database_id = module.distributed_database.id

  dynamic "remove" {
    for_each = var.actions.remove
    content {
      selection = remove.value.selection
    }
  }

  depends_on = [module.distributed_database, oci_distributed_database_distributed_databases_remove_dg.remove_dg]
}


data "oci_distributed_database_distributed_database" "after_actions" {
  distributed_database_id = module.distributed_database.id
  metadata_query          = "VM_CLUSTER_INFO,ADDITIONAL_RESOURCE_INFO,RESOURCE_HIERARCHY"

  depends_on = [oci_distributed_database_distributed_databases_remove.remove]
}
