// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
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
    resource_categories = ["oracle", "oci", "network"]
}

# Fetch all category assignments of dashboard named VCN_DB1
data "oci_log_analytics_log_analytics_resource_categories_list" "VCN_DB1_categories_list" {
    depends_on = [oci_log_analytics_log_analytics_resource_categories_management.VCN_DB1_categories]
    namespace = data.oci_objectstorage_namespace.ns.namespace
    resource_ids = "VCN_DB1"
    resource_types = "DASHBOARD"
}