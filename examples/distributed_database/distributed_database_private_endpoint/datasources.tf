#############################################
# Distributed Database Private Endpoint - datasources
#############################################
data "oci_distributed_database_distributed_database_private_endpoint" "pe_by_id" {
  distributed_database_private_endpoint_id = oci_distributed_database_distributed_database_private_endpoint.private_endpoint.id
}

data "oci_distributed_database_distributed_database_private_endpoints" "pe_list" {
  compartment_id = var.compartment_id
  # Optional filters (uncomment if useful)
  # display_name = var.display_name
  # state        = "ACTIVE"
}
