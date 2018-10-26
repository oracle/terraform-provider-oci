/*
 * This example shows how to manage a bucket
 */

resource "oci_objectstorage_bucket" "bucket1" {
  compartment_id = "${var.compartment_ocid}"
  namespace      = "${data.oci_objectstorage_namespace.ns.namespace}"
  name           = "tf-example-bucket"
  access_type    = "NoPublicAccess"
}

resource "oci_objectstorage_object_lifecycle_policy" "lifecyclePolicy1" {
  namespace = "${data.oci_objectstorage_namespace.ns.namespace}"
  bucket    = "tf-example-bucket"

  #Optional
  rules {
    #Required
    action      = "ARCHIVE"
    is_enabled  = "true"
    name        = "test-rule-1"
    time_amount = "10"
    time_unit   = "DAYS"

    #Optional
    object_name_filter {
      #Optional
      inclusion_prefixes = ["my-test"]
    }
  }
}

data "oci_objectstorage_bucket_summaries" "buckets1" {
  compartment_id = "${var.compartment_ocid}"
  namespace      = "${data.oci_objectstorage_namespace.ns.namespace}"

  filter {
    name   = "name"
    values = ["${oci_objectstorage_bucket.bucket1.name}"]
  }
}

output buckets {
  value = "${data.oci_objectstorage_bucket_summaries.buckets1.bucket_summaries}"
}

data "oci_objectstorage_object_lifecycle_policy" "lifecyclePolicies1" {
  namespace = "${data.oci_objectstorage_namespace.ns.namespace}"
  bucket    = "tf-example-bucket"

  depends_on = ["oci_objectstorage_object_lifecycle_policy.lifecyclePolicy1"]
}

output lifecyclePolicies1 {
  value = "${data.oci_objectstorage_object_lifecycle_policy.lifecyclePolicies1.rules}"
}
