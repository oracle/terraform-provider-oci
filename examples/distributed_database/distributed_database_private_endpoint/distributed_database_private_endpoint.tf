resource "oci_distributed_database_distributed_database_private_endpoint" "private_endpoint" {
  compartment_id = var.compartment_id
  display_name   = var.display_name
  subnet_id      = var.subnet_id

  description = var.description
  nsg_ids     = local.distributed_database_private_endpoint_nsg_ids

  defined_tags  = var.defined_tags
  freeform_tags = var.freeform_tags

  reinstate_proxy_instance_trigger = var.reinstate_proxy_instance_trigger
}
