resource "oci_objectstorage_bucket" "bucket1" {
  compartment_id = var.compartment_ocid
  namespace      = data.oci_objectstorage_namespace.ns.namespace
  name           = "oslink-bucket-tf"
  access_type    = "NoPublicAccess"
  auto_tiering = "Disabled"
}


data "oci_objectstorage_namespace" "ns" {
  #Optional
  compartment_id = var.compartment_ocid
}

output "namespace" {
  value = data.oci_objectstorage_namespace.ns.namespace
}