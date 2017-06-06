resource "baremetal_objectstorage_bucket" "t" {
  compartment_id = "${var.compartment_ocid}"
  name = "BucketOne"
  access_type = "ObjectRead" // or NoPublicAccess
  namespace = "${var.namespace_name}"
  metadata = {
    "foo" = "bar"
  }
}

resource "baremetal_objectstorage_bucket" "t_private" {
  compartment_id = "${var.compartment_ocid}"
  name = "PrivateBucket"
  //defaults to NoPublicAccess
  namespace = "${var.namespace_name}"
  metadata = {
    "foo" = "bar"
  }
}
