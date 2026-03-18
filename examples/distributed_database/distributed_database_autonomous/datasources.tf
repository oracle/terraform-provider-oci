#############################################
# Private Endpoint - datasources
#############################################
/*data "oci_distributed_database_distributed_database_private_endpoint" "pe_by_id" {
  distributed_database_private_endpoint_id = oci_distributed_database_distributed_database_private_endpoint.private_endpoint.id
}*/
/*data "oci_distributed_database_distributed_database_private_endpoints" "pe_list" {
  compartment_id = var.compartment_id
  # Optional filters (uncomment if useful)
  # display_name = var.private_endpoint_display_name
  # state        = "ACTIVE"
}*/
#############################################
# Autonomous Distributed Database - datasources
#############################################
data "oci_distributed_database_distributed_autonomous_database" "dadb_by_id" {
  distributed_autonomous_database_id = oci_distributed_database_distributed_autonomous_database.dadb.id
}

data "oci_distributed_database_distributed_autonomous_databases" "dadb_list" {
  compartment_id = var.compartment_id
  # Optional filters (uncomment if useful)
  # db_deployment_type = "ADB_D"
  # display_name       = var.display_name
  # state              = "ACTIVE"
}
