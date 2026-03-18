

locals {
  dedb_private_endpoint_ids = var.private_endpoint_ids
  vnd                       = var.validate_network_details == null ? {} : var.validate_network_details
}
