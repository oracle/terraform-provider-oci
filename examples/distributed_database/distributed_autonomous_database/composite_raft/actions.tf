locals {
  gds_control_node_actions = var.gds_control_nodes
}

resource "oci_distributed_database_distributed_autonomous_databases_add_gds_control_node" "this" {
  count = length(local.gds_control_node_actions)

  distributed_autonomous_database_id = module.distributed_autonomous_database.id
  public_ssh_key                     = local.gds_control_node_actions[count.index].public_ssh_key
  subnet_id                          = try(local.gds_control_node_actions[count.index].subnet_id, null)

  depends_on = [module.distributed_autonomous_database]
}


locals {
  insert_operations    = var.actions.insert
  insert_dg_operations = var.actions.insert_dg
  remove_dg_operations = var.actions.remove_dg
  remove_operations    = var.actions.remove
}

resource "oci_distributed_database_distributed_autonomous_databases_insert" "insert" {
  count = length(local.insert_operations) == 0 ? 0 : 1

  distributed_autonomous_database_id = module.distributed_autonomous_database.id

  dynamic "insert" {
    for_each = var.actions.insert
    content {
      selection = insert.value.selection
      value     = jsonencode(jsondecode(insert.value.value))
    }
  }

  depends_on = [module.distributed_autonomous_database]
}

resource "oci_distributed_database_distributed_autonomous_databases_configure_sharding" "configure_sharding_insert" {
  count = length(local.insert_operations) > 0 ? 1 : 0

  distributed_autonomous_database_id = module.distributed_autonomous_database.id
  is_rebalance_required              = try(var.configure_sharding_is_rebalance_required, true)
  certificate_id                     = try(var.certificate_id, null)
  ca_bundle_id                       = try(var.ca_bundle_id, null)

  depends_on = [oci_distributed_database_distributed_autonomous_databases_insert.insert]

  lifecycle {
    replace_triggered_by = [oci_distributed_database_distributed_autonomous_databases_insert.insert]
  }
}



resource "oci_distributed_database_distributed_autonomous_databases_insert_dg" "insert_dg" {
  count = length(local.insert_dg_operations) == 0 ? 0 : 1

  distributed_autonomous_database_id = module.distributed_autonomous_database.id

  dynamic "insert_dg" {
    for_each = var.actions.insert_dg
    content {
      selection    = insert_dg.value.selection
      replica_name = insert_dg.value.replica_name
      database {
        cloud_autonomous_vm_cluster_id            = try(insert_dg.value.database.cloud_autonomous_vm_cluster_id, null)
        fast_start_fail_over_lag_limit_in_seconds = try(insert_dg.value.database.fast_start_fail_over_lag_limit_in_seconds, null)
        protection_mode                           = try(insert_dg.value.database.protection_mode, null)
        source                                    = insert_dg.value.database.source
        standby_maintenance_buffer_in_days        = try(insert_dg.value.database.standby_maintenance_buffer_in_days, null)
      }
    }
  }

  depends_on = [module.distributed_autonomous_database, oci_distributed_database_distributed_autonomous_databases_configure_sharding.configure_sharding_insert]
}

resource "oci_distributed_database_distributed_autonomous_databases_remove_dg" "remove_dg" {
  count = length(local.remove_dg_operations) == 0 ? 0 : 1

  distributed_autonomous_database_id = module.distributed_autonomous_database.id

  dynamic "remove_dg" {
    for_each = var.actions.remove_dg
    content {
      selection    = remove_dg.value.selection
      replica_name = remove_dg.value.replica_name
    }
  }

  depends_on = [module.distributed_autonomous_database, oci_distributed_database_distributed_autonomous_databases_insert_dg.insert_dg]
}

resource "oci_distributed_database_distributed_autonomous_databases_remove" "remove" {
  count = length(local.remove_operations) == 0 ? 0 : 1

  distributed_autonomous_database_id = module.distributed_autonomous_database.id

  dynamic "remove" {
    for_each = var.actions.remove
    content {
      selection = remove.value.selection
    }
  }

  depends_on = [module.distributed_autonomous_database, oci_distributed_database_distributed_autonomous_databases_remove_dg.remove_dg]
}


data "oci_distributed_database_distributed_autonomous_database" "after_actions" {
  distributed_autonomous_database_id = module.distributed_autonomous_database.id
  metadata_query                     = "VM_CLUSTER_INFO,ADDITIONAL_RESOURCE_INFO,RESOURCE_HIERARCHY"

  depends_on = [oci_distributed_database_distributed_autonomous_databases_remove.remove]
}
