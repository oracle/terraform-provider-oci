/*locals {
  dadb_private_endpoint_ids = length(var.private_endpoint_ids) > 0 ? var.private_endpoint_ids : [oci_distributed_database_distributed_database_private_endpoint.pe.id]
}*/

locals {
  //dadb_private_endpoint_ids = length(var.private_endpoint_ids) > 0 ? var.private_endpoint_ids : [oci_distributed_database_distributed_database_private_endpoint.pe.id]
  dedb_private_endpoint_ids = var.private_endpoint_ids
  vnd                       = var.validate_network_details == null ? {} : var.validate_network_details
}
