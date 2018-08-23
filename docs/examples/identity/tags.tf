resource "oci_identity_tag_namespace" "tag-namespace1" {
  #Required
  compartment_id = "${var.compartment_ocid}"
  description    = "Just a test"
  name           = "exampletagns"

  is_retired = false
}

resource "oci_identity_tag" "tag1" {
  #Required
  description      = "tf example tag"
  name             = "tf-example-tag"
  tag_namespace_id = "${oci_identity_tag_namespace.tag-namespace1.id}"

  is_retired = false
}

output "tag_namespaces" {
  value = "${oci_identity_tag_namespace.tag-namespace1.id}"
}

output "tags" {
  value = "${oci_identity_tag.tag1.id}"
}

output "resource_defined_tags_key" {
  value = "${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}"
}
