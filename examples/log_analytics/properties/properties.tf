// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * This example shows how to manage collection properties
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

locals {
  namespace        = data.oci_objectstorage_namespace.ns.namespace
  property_name    = data.oci_log_analytics_namespace_properties_metadata.badsql_retry_metadata.property_metadata_summary_collection[0].items[0].name
  display_text     = "badsql_retry"
  property_value   = "PT30M"
  tenancy_level     = "TENANCY"
  source_name      = "unifieddbauditlogfromdbsource121"
}

# Fetch namespace name from object store GET /n
data "oci_objectstorage_namespace" "ns" {
  compartment_id   = var.compartment_ocid
}

# Fetch the metadata details of all collection properties applicable at TENANCY level
data "oci_log_analytics_namespace_properties_metadata" "tenancy_level_metadata" {
  namespace        = local.namespace
  level            = local.tenancy_level
}

# Fetch the metadata details of badsql_retry property
data "oci_log_analytics_namespace_properties_metadata" "badsql_retry_metadata" {
  namespace        = local.namespace
  display_text     = local.display_text
}

# Set the tenant-level collection property value for badsql_retry
resource "oci_log_analytics_log_analytics_preferences_management" "badsql_retry_property" {
  namespace        = local.namespace
  items {
    name           = local.property_name
    value          = local.property_value
  }
}

# Fetch the effective value of badsql_retry property for a source
data "oci_log_analytics_namespace_effective_properties" "badsql_retry_property_value" {
  namespace        = local.namespace
  name             = local.property_name
  source_name      = local.source_name

  depends_on       = [oci_log_analytics_log_analytics_preferences_management.badsql_retry_property]
}