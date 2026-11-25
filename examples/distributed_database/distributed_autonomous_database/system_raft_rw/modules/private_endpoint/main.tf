resource "oci_distributed_database_distributed_database_private_endpoint" "this" {
  compartment_id = var.compartment_id
  display_name   = var.display_name
  subnet_id      = var.subnet_id

  description   = var.description
  freeform_tags = var.freeform_tags
  vcn_nsg_ids {
    vcn_id  = var.vcn_id
    nsg_ids = var.nsg_ids
  }
}
