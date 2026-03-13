#############################################
# Resource outputs
#############################################

output "distributed_database_private_endpoint_id" {
  value       = oci_distributed_database_distributed_database_private_endpoint.private_endpoint.id
  description = "OCID of the created Distributed Database Private Endpoint"
}

output "distributed_database_private_endpoint_state" {
  value       = oci_distributed_database_distributed_database_private_endpoint.private_endpoint.state
  description = "Lifecycle state of the Distributed Database Private Endpoint"
}

output "distributed_database_private_endpoint_private_ip" {
  value       = oci_distributed_database_distributed_database_private_endpoint.private_endpoint.private_ip
  description = "Private IP address assigned to the Distributed Database Private Endpoint"
}

output "distributed_database_private_endpoint_proxy_compute_instance_id" {
  value       = oci_distributed_database_distributed_database_private_endpoint.private_endpoint.proxy_compute_instance_id
  description = "Proxy compute instance OCID backing the private endpoint"
}

output "distributed_database_private_endpoint_vcn_id" {
  value       = oci_distributed_database_distributed_database_private_endpoint.private_endpoint.vcn_id
  description = "VCN OCID where the private endpoint resides"
}

#############################################
# Singular datasource outputs
#############################################

output "ds_distributed_database_private_endpoint_by_id" {
  description = "Distributed Database Private Endpoint singular datasource fetched by OCID."
  value       = data.oci_distributed_database_distributed_database_private_endpoint.pe_by_id
}

#############################################
# Plural datasource outputs
#############################################

output "ds_distributed_database_private_endpoints_list" {
  description = "Distributed Database Private Endpoints plural datasource (list in compartment)."
  value       = data.oci_distributed_database_distributed_database_private_endpoints.pe_list
}

#############################################
# Optional: JSON-friendly outputs
#############################################

output "resource_distributed_database_private_endpoint_json" {
  description = "Distributed Database Private Endpoint resource as JSON string."
  value       = jsonencode(oci_distributed_database_distributed_database_private_endpoint.private_endpoint)
}
