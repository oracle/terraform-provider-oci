// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * This example shows how to add lifeCyclePolicyRules to a bucket
 * We omit the object_name_filter for both types of targets here
 */


resource "oci_objectstorage_bucket" "bucket_with_lifeCyclePolicyRule_2" {
  compartment_id = var.compartment_ocid
  namespace      = data.oci_objectstorage_namespace.ns.namespace
  name           = "tf-example-bucket-lifecyclerule_2"
  access_type    = "NoPublicAccess"
}

resource "oci_objectstorage_object_lifecycle_policy" "lifecyclePolicy_2" {
  namespace = data.oci_objectstorage_namespace.ns.namespace
  bucket    = oci_objectstorage_bucket.bucket_with_lifeCyclePolicyRule_2.name

  #Optional
  rules {
    #Required
    action      = "ARCHIVE"
    is_enabled  = "true"
    name        = "test-rule-3"
    time_amount = "5"
    time_unit   = "DAYS"

    target = "objects"
  }

  #Optional
  rules {
    #Required
    action      = "ABORT"
    is_enabled  = "true"
    name        = "test-rule-4"
    time_amount = "4"
    time_unit   = "DAYS"

    #Optional
    target = "multipart-uploads"

  }
}

data "oci_objectstorage_object_lifecycle_policy" "lifecyclePolicies_2" {
  namespace = data.oci_objectstorage_namespace.ns.namespace
  bucket    = oci_objectstorage_bucket.bucket_with_lifeCyclePolicyRule_2.name

  depends_on = [oci_objectstorage_object_lifecycle_policy.lifecyclePolicy_2]
}

output "lifecyclePolicies_2" {
  value = data.oci_objectstorage_object_lifecycle_policy.lifecyclePolicies_2.rules
}

