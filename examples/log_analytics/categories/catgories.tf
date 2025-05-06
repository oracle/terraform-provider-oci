// Copyright (c) 2017, 2025, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * This example shows how to fetch categories info
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

# Fetch all categories
data "oci_log_analytics_log_analytics_categories_list" "all_categories" {
  namespace = data.oci_objectstorage_namespace.ns.namespace
}

# Fetch categories of type VENDOR
data "oci_log_analytics_log_analytics_categories_list" "vendor_categories" {
  namespace = data.oci_objectstorage_namespace.ns.namespace
  category_type = "VENDOR"
}

# Fetch categories that have Oracle in their display text
data "oci_log_analytics_log_analytics_categories_list" "Oracle_displaytext_categories" {
  namespace = data.oci_objectstorage_namespace.ns.namespace
  category_display_text = "Oracle"
}

# Fetch category named oracle
data "oci_log_analytics_log_analytics_category" "oracle_category" {
  namespace = data.oci_objectstorage_namespace.ns.namespace
  name = "oracle"
}

# Manage category assignments of dashboard named VCN_DB1
resource "oci_log_analytics_log_analytics_resource_categories_management" "VCN_DB1_categories" {
    namespace = data.oci_objectstorage_namespace.ns.namespace
    resource_id = "VCN_DB1"
    resource_type = "DASHBOARD"
    resource_categories = ["oracle", "linux", "database"]
}

# Fetch all category assignments of dashboard named VCN_DB1
data "oci_log_analytics_log_analytics_resource_categories_list" "VCN_DB1_categories_list" {
    depends_on = [oci_log_analytics_log_analytics_resource_categories_management.VCN_DB1_categories]
    namespace = data.oci_objectstorage_namespace.ns.namespace
    resource_ids = "VCN_DB1"
    resource_types = "DASHBOARD"
}

# Create a lookup with categories
resource "oci_log_analytics_namespace_lookup" "TFLookup" {
  compartment_id       = var.compartment_ocid
  namespace            = data.oci_objectstorage_namespace.ns.namespace
  lookup_name          = "TFLookup"
  type                 = "Lookup"
  description          = "A simple lookup"
  register_lookup_file = "./files/vendor.csv"

  categories {
    name = "database"
    type = "TIER"
  }

  categories {
    name = "oracle"
    type = "VENDOR"
  }
}

# Fetch all category assignments in compartment
data "oci_log_analytics_log_analytics_resource_categories_list" "compartment_resource_categories_list" {
  depends_on = [oci_log_analytics_namespace_lookup.TFLookup]
  namespace = data.oci_objectstorage_namespace.ns.namespace
  compartment_id = var.compartment_ocid
}