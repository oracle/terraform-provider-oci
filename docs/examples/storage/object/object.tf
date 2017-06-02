resource "baremetal_objectstorage_bucket" "t" {
  compartment_id = "${var.compartment_ocid}"
  name = "${var.BucketName}"
  access_type = "ObjectRead" // or NoPublicAccess
  namespace = "${var.namespace_name}"
  metadata = {
    "foo" = "bar"
  }
}