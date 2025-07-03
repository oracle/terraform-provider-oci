resource "oci_identity_tag_namespace" "tag_namespace" {
  #Required
  compartment_id = var.compartment_id
  description = "example tag namespace"
  name = var.defined_tag_namespace_name != "" ? var.defined_tag_namespace_name : "example-tag-namespace-all"

  is_retired = false
}

resource "oci_identity_tag" "tag" {
  #Required
  description = "example tag"
  name = "example-tag"
  tag_namespace_id = oci_identity_tag_namespace.tag_namespace.id

  is_retired = false
}