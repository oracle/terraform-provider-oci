// Copyright (c) 2017, 2025, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * This example shows how to manage log analytics lookup resource
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
  compartment_id   = var.compartment_ocid
}

// Create a simple lookup
resource "oci_log_analytics_namespace_lookup" "TFLookup" {
  compartment_id       = var.tenancy_ocid
  namespace            = data.oci_objectstorage_namespace.ns.namespace
  lookup_name          = "TFLookup"
  type                 = "Lookup"
  register_lookup_file = "./files/vendor.csv"

  description          = "A simple lookup"
  char_encoding        = "UTF-8"
  is_hidden            = false
  freeform_tags        = {"servicegroup" = "test", "Dept" = "Devops"}
  defined_tags         = {"Oracle-Recommended-Tags.ResourceAuthorizedUser" = "test"}
}

# Create a simple lookup with tags
resource "oci_log_analytics_namespace_lookup" "TFLookupWithTags" {
  compartment_id       = var.tenancy_ocid
  namespace            = data.oci_objectstorage_namespace.ns.namespace
  lookup_name          = "TFLookupWithTags"
  type                 = "Lookup"
  register_lookup_file = "./files/vendor.csv"

  description          = "A simple lookup with tags"
  freeform_tags        = {"servicegroup" = "test", "Dept" = "Devops"}
  defined_tags         = {"Oracle-Recommended-Tags.ResourceAuthorizedUser" = "test"}
}

# Create a simple lookup with all attributes like categories and fields
resource "oci_log_analytics_namespace_lookup" "TFLookupAll" {
  compartment_id       = var.tenancy_ocid
  namespace            = data.oci_objectstorage_namespace.ns.namespace
  lookup_name          = "TFLookupAll"
  type                 = "Lookup"
  register_lookup_file = "./files/vendor.csv"

  description          = "A simple lookup with all attributes"
  freeform_tags        = {"servicegroup" = "test", "Dept" = "Devops"}
  defined_tags         = {"Oracle-Recommended-Tags.ResourceAuthorizedUser" = "test"}

  char_encoding        = "UTF-8"
  default_match_value  = "WILDCARD"
  is_hidden            = false
  max_matches          = 25

  categories {
    name = "database"
    type = "TIER"
  }

  categories {
    name = "oracle"
    type = "VENDOR"
  }

  fields {
    name = "name"
    match_operator = "WILDCARD"
  }

  fields {
    name = "profitmodel"
    match_operator = "WILDCARD"
  }

  lifecycle {
    ignore_changes = [
      fields, defined_tags
    ]
  }
}

# Append data to a lookup
resource "oci_log_analytics_namespace_lookups_append_data_management" "appendData" {
  namespace            = data.oci_objectstorage_namespace.ns.namespace
  lookup_name          = oci_log_analytics_namespace_lookup.TFLookup.lookup_name
  append_lookup_file   = "./files/append.csv"

  depends_on = [oci_log_analytics_namespace_lookup.TFLookup]
}

# Update data of a lookup
resource "oci_log_analytics_namespace_lookups_update_data_management" "updateData" {
  namespace            = data.oci_objectstorage_namespace.ns.namespace
  lookup_name          = oci_log_analytics_namespace_lookup.TFLookup.lookup_name
  update_lookup_file   = "./files/update.csv"

  depends_on = [oci_log_analytics_namespace_lookups_append_data_management.appendData]
}

# Get details of a lookup
data "oci_log_analytics_namespace_lookup" "BEALookup" {
  namespace            = data.oci_objectstorage_namespace.ns.namespace
  lookup_name          = "omc_beaerrmsgs"
}