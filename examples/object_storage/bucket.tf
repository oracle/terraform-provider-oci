// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

/*
 * This example shows how to manage a bucket
 */

resource "oci_objectstorage_bucket" "bucket1" {
  compartment_id = "${var.compartment_ocid}"
  namespace      = "${data.oci_objectstorage_namespace.ns.namespace}"
  name           = "tf-example-bucket"
  access_type    = "NoPublicAccess"
}

resource "oci_objectstorage_bucket" "bucket_with_versioning" {
  compartment_id = "${var.compartment_ocid}"
  namespace      = "${data.oci_objectstorage_namespace.ns.namespace}"
  name           = "bucketWithVersioning"
  access_type    = "NoPublicAccess"
  versioning     = "Suspended"
}

resource "oci_objectstorage_object_lifecycle_policy" "lifecyclePolicy1" {
  namespace = "${data.oci_objectstorage_namespace.ns.namespace}"
  bucket    = "${oci_objectstorage_bucket.bucket1.name}"

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
  bucket    = "${oci_objectstorage_bucket.bucket1.name}"

  depends_on = ["oci_objectstorage_object_lifecycle_policy.lifecyclePolicy1"]
}

output lifecyclePolicies1 {
  value = "${data.oci_objectstorage_object_lifecycle_policy.lifecyclePolicies1.rules}"
}

data "oci_objectstorage_object_versions" "test_object_versions1" {
  #Required
  bucket    = "${oci_objectstorage_bucket.bucket1.name}"
  namespace = "${data.oci_objectstorage_namespace.ns.namespace}"
}

output test_object_versions1 {
  value = "${data.oci_objectstorage_object_versions.test_object_versions1.items}"
}
