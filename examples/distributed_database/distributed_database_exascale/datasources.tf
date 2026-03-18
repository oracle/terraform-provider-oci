#############################################
# Private Endpoint - datasources
#############################################
/*data "oci_distributed_database_distributed_database_private_endpoint" "pe_by_id" {
  distributed_database_private_endpoint_id = oci_distributed_database_distributed_database_private_endpoint.private_endpoint.id
}*/
/*data "oci_distributed_database_distributed_database_private_endpoints" "pe_list" {
  compartment_id = var.compartment_ocid
  # Optional filters (uncomment if useful)
  # display_name = var.private_endpoint_display_name
  # state        = "ACTIVE"
}*/
#############################################
# Exascale Distributed Database - datasources
#############################################
data "oci_distributed_database_distributed_database" "gdd_by_id" {
  distributed_database_id = oci_distributed_database_distributed_database.ddb.id
}
data "oci_distributed_database_distributed_databases" "gdd_list" {
  compartment_id = var.compartment_id
  # Optional filters (uncomment if useful)
  # db_deployment_type = "EXADB_XS"
  # display_name       = var.gdd_display_name
  # state              = "AVAILABLE"
}