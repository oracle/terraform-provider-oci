// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * This example shows how to manage log analytics unprocessed data bucket resource
 */

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "compartment_ocid" {}
variable "region" {}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

# Fetch namespace name from object store GET /n
data "oci_objectstorage_namespace" "ns" {
  compartment_id             = var.compartment_ocid
}

variable "log_analytics_unprocessed_data_bucket_name" {
  default = "tf-unprocessed-data-bucket"
}

variable "log_analytics_unprocessed_data_bucket_enabled" {
  default = "true"
}

# Create a unprocessed data bucket
resource "oci_log_analytics_log_analytics_unprocessed_data_bucket_management" "unprocessedDataBucket" {
  namespace                     = data.oci_objectstorage_namespace.ns.namespace
  bucket                        = var.log_analytics_unprocessed_data_bucket_name
  is_enabled                    = var.log_analytics_unprocessed_data_bucket_enabled
}

# Get details of above created unprocessed data bucket
data "oci_log_analytics_log_analytics_unprocessed_data_bucket" "unprocessedDataBucketDetails" {
  depends_on                 = [oci_log_analytics_log_analytics_unprocessed_data_bucket_management.unprocessedDataBucket]
  namespace                  = data.oci_objectstorage_namespace.ns.namespace
}
