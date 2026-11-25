output "private_endpoint_id" {
  value = try(module.private_endpoint[0].id, null)
}

output "private_endpoint_ids" {
  value = local.private_endpoint_ids
}

output "distributed_autonomous_database_id" {
  value = module.distributed_autonomous_database.id
}

output "distributed_autonomous_database_lifecycle_state" {
  value = module.distributed_autonomous_database.lifecycle_state
}

output "gds_control_node_details" {
  value = module.distributed_autonomous_database.gds_control_node_details
}

output "action_work_request_ids" {
  value = {
    insert    = try(oci_distributed_database_distributed_autonomous_databases_insert.insert[0].work_request_id, null)
    insert_dg = try(oci_distributed_database_distributed_autonomous_databases_insert_dg.insert_dg[0].work_request_id, null)
    remove_dg = try(oci_distributed_database_distributed_autonomous_databases_remove_dg.remove_dg[0].work_request_id, null)
    remove    = try(oci_distributed_database_distributed_autonomous_databases_remove.remove[0].work_request_id, null)
  }
}

output "database_operational_state" {
  value = data.oci_distributed_database_distributed_autonomous_database.after_actions.operational_state
}

output "observed_topology" {
  sensitive = true
  value = {
    catalog_details                   = data.oci_distributed_database_distributed_autonomous_database.after_actions.catalog_details
    user_shard_spaces                 = data.oci_distributed_database_distributed_autonomous_database.after_actions.user_shard_spaces
    system_raft_clusters              = data.oci_distributed_database_distributed_autonomous_database.after_actions.system_raft_clusters
    composite_raft_shard_spaces       = data.oci_distributed_database_distributed_autonomous_database.after_actions.composite_raft_shard_spaces
    system_data_guard_databases       = data.oci_distributed_database_distributed_autonomous_database.after_actions.system_data_guard_databases
    composite_data_guard_shard_spaces = data.oci_distributed_database_distributed_autonomous_database.after_actions.composite_data_guard_shard_spaces
  }
}
