// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * This example shows how to manage a bucket
 */

resource "oci_objectstorage_bucket" "bucket1" {
  compartment_id = var.compartment_ocid
  namespace      = data.oci_objectstorage_namespace.ns.namespace
  name           = "tf-example-bucket"
  access_type    = "NoPublicAccess"
  auto_tiering = "Disabled"
}

resource "oci_objectstorage_bucket" "bucket_with_versioning" {
  compartment_id = var.compartment_ocid
  namespace      = data.oci_objectstorage_namespace.ns.namespace
  name           = "bucket-with-versioning"
  access_type    = "NoPublicAccess"
  versioning     = "Enabled"
}

data "oci_objectstorage_bucket_summaries" "buckets1" {
  compartment_id = var.compartment_ocid
  namespace      = data.oci_objectstorage_namespace.ns.namespace

  filter {
    name   = "name"
    values = [oci_objectstorage_bucket.bucket1.name]
  }
}

output "buckets" {
  value = data.oci_objectstorage_bucket_summaries.buckets1.bucket_summaries
}

data "oci_objectstorage_object_versions" "test_object_versions1" {
  #Required
  bucket    = oci_objectstorage_bucket.bucket1.name
  namespace = data.oci_objectstorage_namespace.ns.namespace
}

output "test_object_versions1" {
  value = data.oci_objectstorage_object_versions.test_object_versions1.items
}

