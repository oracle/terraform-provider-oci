output "id" {
  description = "Distributed database OCID."
  value       = oci_distributed_database_distributed_database.exascale.id
}

output "display_name" {
  description = "Distributed database display name."
  value       = data.oci_distributed_database_distributed_database.exascale.display_name
}

output "lifecycle_state" {
  description = "Distributed database lifecycle state."
  value       = try(data.oci_distributed_database_distributed_database.exascale.state, oci_distributed_database_distributed_database.exascale.state, null)
}

output "metadata" {
  description = "Metadata returned by the GET DDB call."
  value       = try(data.oci_distributed_database_distributed_database.exascale.metadata, [])
}

output "gsm_details" {
  description = "GSM details returned by the GET DDB call."
  value       = try(data.oci_distributed_database_distributed_database.exascale.gsm_details, [])
}

output "gds_control_node_details" {
  description = "First GDS control node detail returned by the GET DDB call."
  value       = local.first_gds_control_node_details
}

output "gds_control_node_private_ip" {
  description = "GDSCTL private IP from selected GDS control node metadata."
  value       = local.gds_control_node_private_ip
}
