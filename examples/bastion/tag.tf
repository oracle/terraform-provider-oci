resource "oci_identity_tag_namespace" "bastion_tag_namespace1" {
  #Required
  compartment_id = var.tenancy_ocid
  description    = var.tag_namespace_description
  name           = var.tag_namespace_name
}

resource "oci_identity_tag" "bastion_tag1" {
  #Required
  description      = "tf example tag"
  name             = "tf-example-tag"
  tag_namespace_id = oci_identity_tag_namespace.bastion_tag_namespace1.id
}