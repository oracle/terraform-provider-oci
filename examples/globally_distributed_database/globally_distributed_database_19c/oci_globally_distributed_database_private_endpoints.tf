resource "oci_globally_distributed_database_private_endpoint" "this" {
  #Required
  compartment_id = var.compartment_ocid
  display_name   = "GloballyDistributedDB-PrivateEndpoint-Example"
  subnet_id      = var.subnet_ocid

  #Optional
  #defined_tags  = var.oci_globally_distributed_database_defined_tags_value
  description = "Test OCI Globally Distributed Database Private Endpoint"
  #freeform_tags = var.oci_globally_distributed_database_freeform_tags
  nsg_ids = var.nsg_ocids
    
  lifecycle {
    ignore_changes = [
      sharded_databases,
    ]
  }
}