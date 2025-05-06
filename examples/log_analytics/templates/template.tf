// Copyright (c) 2017, 2025, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * This example shows how to use terraform feature of templates
 */

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "compartment_ocid" {}
variable "region" {}

provider "oci" {
  tenancy_ocid         = var.tenancy_ocid
  user_ocid            = var.user_ocid
  fingerprint          = var.fingerprint
  private_key_path     = var.private_key_path
  region               = var.region
}

# Fetch namespace name from object store GET /n
data "oci_objectstorage_namespace" "ns" {
  compartment_id       = var.compartment_ocid
}

# List all templates in a compartment with filters
data "oci_log_analytics_namespace_templates" "filtered_templates" {
  compartment_id        = var.tenancy_ocid
  namespace             = data.oci_objectstorage_namespace.ns.namespace
  name                  = "Linux ROOT Logins"
  state                 = "ACTIVE"
  template_display_text = "root"
  type                  = "Scheduled Search"
}

# Get details of a template
data "oci_log_analytics_namespace_template" "linux_root_login_template" {
  namespace             = data.oci_objectstorage_namespace.ns.namespace
  template_id           = data.oci_log_analytics_namespace_templates.filtered_templates.log_analytics_template_collection[0].items[0].id
}