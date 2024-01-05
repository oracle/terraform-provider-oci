// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * This example onboards a namespace with Log Analytics
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

data "oci_log_analytics_namespaces" "test_namespaces" {
  compartment_id = var.tenancy_ocid
}

// will return NotAuthorizedOrNotFound 404 Error if tenancy not onboarded with log analytics
//data "oci_log_analytics_namespace" "test_namespace" {
//  namespace = data.oci_log_analytics_namespaces.test_namespaces.namespace_collection.0.items.0.namespace
//}

resource "oci_log_analytics_namespace" "test_namespace" {
  namespace = data.oci_log_analytics_namespaces.test_namespaces.namespace_collection.0.items.0.namespace
  is_onboarded = false
  compartment_id = var.tenancy_ocid
}

# Fetch namespace name from object store GET /n
data "oci_objectstorage_namespace" "ns" {
  compartment_id             = var.compartment_ocid
}

