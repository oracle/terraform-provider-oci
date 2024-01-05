// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}

# Fetch namespace name from object store GET /n
data "oci_objectstorage_namespace" "ns" {
  compartment_id             = var.compartment_ocid
}

data "oci_log_analytics_namespace_storage_encryption_key_info" "test_namespace_storage_encryption_key_info" {
  #Required
  namespace = data.oci_objectstorage_namespace.ns.namespace
}