output "private_endpoint_id" {
  description = "Created distributed database private endpoint OCID. Null when create_private_endpoint is false."
  value       = try(module.private_endpoint[0].id, null)
}

output "private_endpoint_ids" {
  description = "Private endpoint OCIDs passed to the distributed database."
  value       = local.private_endpoint_ids
}

output "distributed_database_id" {
  description = "Distributed database OCID."
  value       = module.distributed_database.id
}

output "distributed_database_lifecycle_state" {
  description = "Distributed database lifecycle state."
  value       = module.distributed_database.lifecycle_state
}
output "gds_control_node_details" {
  description = "GDS control-node details returned by the distributed database."
  value       = module.distributed_database.gds_control_node_details
}

output "action_work_request_ids" {
  value = {
    insert    = try(oci_distributed_database_distributed_databases_insert.insert[0].work_request_id, null)
    insert_dg = try(oci_distributed_database_distributed_databases_insert_dg.insert_dg[0].work_request_id, null)
    remove_dg = try(oci_distributed_database_distributed_databases_remove_dg.remove_dg[0].work_request_id, null)
    remove    = try(oci_distributed_database_distributed_databases_remove.remove[0].work_request_id, null)
  }
}

output "database_operational_state" {
  value = data.oci_distributed_database_distributed_database.after_actions.operational_state
}

output "observed_topology" {
  sensitive = true
  value = {
    catalog_details                   = data.oci_distributed_database_distributed_database.after_actions.catalog_details
    user_shard_spaces                 = data.oci_distributed_database_distributed_database.after_actions.user_shard_spaces
    system_raft_clusters              = data.oci_distributed_database_distributed_database.after_actions.system_raft_clusters
    composite_raft_shard_spaces       = data.oci_distributed_database_distributed_database.after_actions.composite_raft_shard_spaces
    system_data_guard_databases       = data.oci_distributed_database_distributed_database.after_actions.system_data_guard_databases
    composite_data_guard_shard_spaces = data.oci_distributed_database_distributed_database.after_actions.composite_data_guard_shard_spaces
  }
}
