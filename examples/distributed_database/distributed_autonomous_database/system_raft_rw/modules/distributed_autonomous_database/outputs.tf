output "id" {
  value = oci_distributed_database_distributed_autonomous_database.this.id
}

output "display_name" {
  value = data.oci_distributed_database_distributed_autonomous_database.this.display_name
}

output "lifecycle_state" {
  value = try(data.oci_distributed_database_distributed_autonomous_database.this.state, null)
}

output "gds_control_node_details" {
  value = try(data.oci_distributed_database_distributed_autonomous_database.this.gds_control_node_details, [])
}
