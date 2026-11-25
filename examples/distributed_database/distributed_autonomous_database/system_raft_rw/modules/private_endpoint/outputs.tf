output "id" {
  description = "Distributed database private endpoint OCID."
  value       = oci_distributed_database_distributed_database_private_endpoint.this.id
}

output "private_ip" {
  description = "Private endpoint private IP."
  value       = oci_distributed_database_distributed_database_private_endpoint.this.private_ip
}

output "proxy_compute_instance_id" {
  description = "Private endpoint proxy compute instance OCID."
  value       = oci_distributed_database_distributed_database_private_endpoint.this.proxy_compute_instance_id
}

output "subnet_id" {
  description = "Private endpoint subnet OCID."
  value       = oci_distributed_database_distributed_database_private_endpoint.this.subnet_id
}

output "state" {
  description = "Private endpoint lifecycle state."
  value       = oci_distributed_database_distributed_database_private_endpoint.this.state
}
