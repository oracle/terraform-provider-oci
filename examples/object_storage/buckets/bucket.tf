// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}

variable "bucket_access_type" {
  default = "NoPublicAccess"
}

variable "bucket_auto_tiering" {
  default = "Disabled"
}

variable "bucket_bucket_scope" {
  default = "NAMESPACE"
}

variable "bucket_defined_tags_value" {
  default = "value"
}

variable "bucket_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "bucket_metadata" {
  default = { "content-type" = "text/plain" }
}

variable "bucket_name" {
  default = "testBucketName"
}

variable "bucket_object_events_enabled" {
  default = false
}

variable "bucket_storage_tier" {
  default = "Standard"
}

variable "bucket_versioning" {
  default = "Enabled"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_objectstorage_namespace" "os_namespace" {
  #Optional
  compartment_id = var.compartment_ocid
}

resource "oci_objectstorage_bucket" "test_bucket" {
  #Required
  compartment_id = var.compartment_ocid
  name           = var.bucket_name
  namespace      = data.oci_objectstorage_namespace.os_namespace.namespace

  #Optional
  access_type           = var.bucket_access_type
  auto_tiering          = var.bucket_auto_tiering
  bucket_scope          = var.bucket_bucket_scope
  freeform_tags         = var.bucket_freeform_tags
  metadata              = var.bucket_metadata
  object_events_enabled = var.bucket_object_events_enabled
  storage_tier          = var.bucket_storage_tier
  versioning            = var.bucket_versioning
}

resource "oci_objectstorage_bucket" "regional_test_bucket" {
  #Required
  compartment_id = var.compartment_ocid
  name           = "test-bucket-name-2"
  namespace      = data.oci_objectstorage_namespace.os_namespace.namespace

  #Optional
  access_type           = var.bucket_access_type
  auto_tiering          = var.bucket_auto_tiering
  bucket_scope          = "REGION"
  freeform_tags         = var.bucket_freeform_tags
  metadata              = var.bucket_metadata
  object_events_enabled = var.bucket_object_events_enabled
  storage_tier          = var.bucket_storage_tier
  versioning            = var.bucket_versioning
}

data "oci_objectstorage_bucket_summaries" "test_buckets" {
  #Required
  compartment_id = var.compartment_ocid
  namespace      = data.oci_objectstorage_namespace.os_namespace.namespace
  depends_on = [oci_objectstorage_bucket.regional_test_bucket, oci_objectstorage_bucket.test_bucket]
}

data "oci_objectstorage_bucket" "regional_test_bucket_data" {
  #Required
  name = "test-bucket-name-2"
  namespace = data.oci_objectstorage_namespace.os_namespace.namespace
  depends_on = [oci_objectstorage_bucket.regional_test_bucket]
}

output "bucket_scope_from_list" {
  value = [for summary in data.oci_objectstorage_bucket_summaries.test_buckets.bucket_summaries : summary.bucket_scope]
}

output "bucket_scope_region" {
  value = data.oci_objectstorage_bucket.regional_test_bucket_data.bucket_scope
}
