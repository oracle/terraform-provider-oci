// Copyright (c) 2017, 2025, Oracle and/or its affiliates. All rights reserved.
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
  namespace              = data.oci_objectstorage_namespace.ns.namespace
  property_name          = data.oci_log_analytics_namespace_properties_metadata.rest_api_timezone_metadata.property_metadata_summary_collection[0].items[0].name
  rest_api_timezone_name = "management_agent.rest_api.timezone"
  property_value         = "PST"
  tenancy_level          = "TENANCY"
  source_name            = "ociIdcsAuditEventsRestLogSource"
  pattern_id_long        = 836081380904045451
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

# Fetch the metadata details of management_agent.rest_api.timezone property
data "oci_log_analytics_namespace_properties_metadata" "rest_api_timezone_metadata" {
  namespace        = local.namespace
  name             = local.rest_api_timezone_name
}

# Set the tenant-level collection property value for management_agent.rest_api.timezone
resource "oci_log_analytics_log_analytics_preferences_management" "rest_api_timezone_property" {
  namespace        = local.namespace
  items {
    name           = local.property_name
    value          = local.property_value
  }
}

# Fetch the effective value of management_agent.rest_api.timezone property for a source
data "oci_log_analytics_namespace_effective_properties" "rest_api_timezone_property_value_source" {
  namespace        = local.namespace
  name             = local.property_name
  source_name      = local.source_name

  depends_on       = [oci_log_analytics_log_analytics_preferences_management.rest_api_timezone_property]
}

# Fetch the effective value of management_agent.rest_api.timezone property for a pattern
data "oci_log_analytics_namespace_effective_properties" "rest_api_timezone_property_value_pattern" {
  namespace        = local.namespace
  name             = local.property_name
  source_name      = local.source_name
  pattern_id_long  = local.pattern_id_long

  depends_on       = [oci_log_analytics_log_analytics_preferences_management.rest_api_timezone_property]
}