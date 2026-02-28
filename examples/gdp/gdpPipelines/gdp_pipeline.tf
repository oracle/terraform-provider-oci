// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "buckets_compartment_id" {}
variable "pipeline_compartment_id" {}
variable "peered_region" {}


provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_objectstorage_namespace" "test_namespace" {
  compartment_id = "${var.tenancy_ocid}"
}

resource "oci_objectstorage_bucket" "test_source_bucket" {
  compartment_id = var.buckets_compartment_id
  name = "tf-test-source"
  namespace = data.oci_objectstorage_namespace.test_namespace.namespace
  object_events_enabled = true
}

resource "oci_objectstorage_bucket" "test_transfer_bucket" {
  compartment_id = var.buckets_compartment_id
  name = "tf-test-transfer"
  namespace = data.oci_objectstorage_namespace.test_namespace.namespace
}

resource "oci_objectstorage_bucket" "test_reject_bucket" {
  compartment_id = var.buckets_compartment_id
  name = "tf-test-reject"
  namespace = data.oci_objectstorage_namespace.test_namespace.namespace
}

resource "oci_gdp_gdp_pipeline" "test_gdp_pipeline" {
  #Required
  bucket_details {
    #Required
    bucket_type = "SOURCE"
    id          = oci_objectstorage_bucket.test_source_bucket.bucket_id
    name        = oci_objectstorage_bucket.test_source_bucket.name
    namespace   = oci_objectstorage_bucket.test_source_bucket.namespace
  }
  bucket_details {
    #Required
    bucket_type = "TRANSFER"
    id          = oci_objectstorage_bucket.test_transfer_bucket.bucket_id
    name        = oci_objectstorage_bucket.test_transfer_bucket.name
    namespace   = oci_objectstorage_bucket.test_transfer_bucket.namespace
  }
  bucket_details {
    #Required
    bucket_type = "REJECT"
    id          = oci_objectstorage_bucket.test_reject_bucket.bucket_id
    name        = oci_objectstorage_bucket.test_reject_bucket.name
    namespace   = oci_objectstorage_bucket.test_reject_bucket.namespace
  }
  compartment_id = var.pipeline_compartment_id
  display_name   = "tf-test-sender"
  peering_region = var.peered_region
  pipeline_type  = "SENDER"
}

data "oci_gdp_gdp_pipelines" "test_gdp_pipelines" {

  #Optional
  compartment_id  = var.pipeline_compartment_id
  display_name    = "tf-test-sender"
}

