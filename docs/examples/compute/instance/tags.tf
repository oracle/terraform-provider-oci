resource "oci_identity_tag_namespace" "tag-namespace1" {
  #Required
  compartment_id = "${var.compartment_ocid}"
  description = "tf example tag namespace"
  name = "tf-example-tag-namespace"
}

resource "oci_identity_tag" "tag1" {
  #Required
  description = "tf example tag"
  name = "tf-example-tag"
  tag_namespace_id = "${oci_identity_tag_namespace.tag-namespace1.id}"
}

resource "oci_identity_tag" "tag2" {
  #Required
  description = "tf example tag"
  name = "tf-example-tag-2"
  tag_namespace_id = "${oci_identity_tag_namespace.tag-namespace1.id}"
}