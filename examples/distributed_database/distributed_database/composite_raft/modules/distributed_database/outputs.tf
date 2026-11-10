output "id" {
  value       = oci_distributed_database_distributed_database.this.id
  description = "Distributed database OCID."
}

output "lifecycle_state" {
  value       = try(data.oci_distributed_database_distributed_database.this.state, oci_distributed_database_distributed_database.this.state, null)
  description = "Distributed database lifecycle state."
}

output "gds_control_node_details" {
  value       = try(data.oci_distributed_database_distributed_database.this.gds_control_node_details, [])
  description = "GDS control node details returned by the service."
}
