// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * This example shows how to import custom content
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

# Update preferences
resource "oci_log_analytics_log_analytics_preferences_management" "preferences" {
    namespace = data.oci_objectstorage_namespace.ns.namespace
    items {
        name  = "DEFAULT_HOMEPAGE"
        value = "value1"
    }
}

# Fetch updated preferences
data "oci_log_analytics_log_analytics_preference" "preferenceList" {
  depends_on = [oci_log_analytics_log_analytics_preferences_management.preferences]
  namespace = data.oci_objectstorage_namespace.ns.namespace
}
