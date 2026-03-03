// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "kms_key_id" {}

variable "bucket_access_type" {
  default = "NoPublicAccess"
}

variable "bucket_auto_tiering" {
  default = "Disabled"
}

variable "bucket_is_bucket_key_enabled" {
  default = false
}

variable "bucket_metadata" {
  default = { "content-type" = "text/plain" }
}

variable "bucket_kms_key_id" {
  default = null
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

data "oci_objectstorage_namespace" "ns" {
  compartment_id = var.compartment_ocid
}

resource "oci_objectstorage_bucket" "test_bucket" {
  #Required
  compartment_id = var.compartment_ocid
  name           = "test-bucket-bucket-key"
  namespace      = data.oci_objectstorage_namespace.ns.namespace

  #Optional
  access_type           = var.bucket_access_type
  auto_tiering          = var.bucket_auto_tiering
  is_bucket_key_enabled = var.bucket_is_bucket_key_enabled
  kms_key_id 			= var.kms_key_id
  metadata              = var.bucket_metadata
  object_events_enabled = var.bucket_object_events_enabled
  storage_tier          = var.bucket_storage_tier
  versioning            = var.bucket_versioning
}

resource "oci_objectstorage_bucket" "test_bucket_bucket_key_enabled" {
  #Required
  compartment_id = var.compartment_ocid
  name           = "test-bucket-bucket-key-enabled"
  namespace      = data.oci_objectstorage_namespace.ns.namespace

  #Optional
  access_type           = var.bucket_access_type
  auto_tiering          = var.bucket_auto_tiering
  kms_key_id 			= var.kms_key_id
  is_bucket_key_enabled = true
  metadata              = var.bucket_metadata
  object_events_enabled = var.bucket_object_events_enabled
  storage_tier          = var.bucket_storage_tier
  versioning            = var.bucket_versioning
}

data "oci_objectstorage_bucket_summaries" "test_buckets" {
  #Required
  compartment_id = var.compartment_ocid
  namespace      = data.oci_objectstorage_namespace.ns.namespace
}