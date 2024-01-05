// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * This example shows how to add lifeCyclePolicyRules to a bucket
 */

variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "region" {
}

variable "compartment_ocid" {
}

data "oci_objectstorage_namespace" "ns" {
  #Optional
  compartment_id = var.compartment_ocid
}

resource "oci_objectstorage_bucket" "bucket_with_lifeCyclePolicyRule" {
  compartment_id = var.compartment_ocid
  namespace      = data.oci_objectstorage_namespace.ns.namespace
  name           = "tf-example-bucket-lifecyclerule"
  access_type    = "NoPublicAccess"
}

resource "oci_objectstorage_object_lifecycle_policy" "lifecyclePolicy" {
  namespace = data.oci_objectstorage_namespace.ns.namespace
  bucket    = oci_objectstorage_bucket.bucket_with_lifeCyclePolicyRule.name

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

    target = "objects"
  }

  #Optional
  rules {
    #Required
    action      = "ABORT"
    is_enabled  = "true"
    name        = "test-rule-1"
    time_amount = "10"
    time_unit   = "DAYS"

    #Optional
    target = "multipart-uploads"
  }
}

data "oci_objectstorage_object_lifecycle_policy" "lifecyclePolicies" {
  namespace = data.oci_objectstorage_namespace.ns.namespace
  bucket    = oci_objectstorage_bucket.bucket_with_lifeCyclePolicyRule.name

  depends_on = [oci_objectstorage_object_lifecycle_policy.lifecyclePolicy]
}

output "lifecyclePolicies" {
  value = data.oci_objectstorage_object_lifecycle_policy.lifecyclePolicies.rules
}

