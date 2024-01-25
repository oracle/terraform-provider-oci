// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * This example shows how to manage log analytics ingest time rule resource
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

# Create an ingest time rule with minimal required parameters
resource "oci_log_analytics_namespace_ingest_time_rule" "ingest_time_rule_minimal" {
  compartment_id   = var.compartment_ocid
  namespace        = data.oci_objectstorage_namespace.ns.namespace
  display_name     = "displayName"
  conditions {
    kind           = "FIELD"
    field_name     = "mtag"
    field_operator = "EQUAL"
    field_value    = "cveexploitattempt"
  }
  actions {
    type           = "METRIC_EXTRACTION"
    compartment_id = var.compartment_ocid
    namespace      = "tfmetricnamespace"
    metric_name    = "tfmetriccve"
  }
}

# Create an ingest time rule with all parameters
resource "oci_log_analytics_namespace_ingest_time_rule" "ingest_time_rule_full" {
  compartment_id   = var.compartment_ocid
  namespace        = data.oci_objectstorage_namespace.ns.namespace
  display_name     = "displayName2"
  description      = "description2"
  conditions {
    kind           = "FIELD"
    field_name     = "mtag"
    field_operator = "EQUAL"
    field_value    = "cveexploitattempt"
    additional_conditions {
      condition_field    = "SOURCE_NAME"
      condition_operator = "EQUAL"
      condition_value    = "omc_ociAuditLogSource"
    }
    additional_conditions {
      condition_field    = "mtgttype"
      condition_operator = "EQUAL"
      condition_value    = "omc_host_linux"
    }
  }
  actions {
    type           = "METRIC_EXTRACTION"
    compartment_id = var.compartment_ocid
    namespace      = "tfmetricnamespace"
    metric_name    = "tfmetriccve"
    resource_group = "critical"
    dimensions     = ["SOURCE_NAME", "event"]
  }
}

# Disable an ingest time rule
resource "oci_log_analytics_namespace_ingest_time_rules_management" "ingest_time_rule_full_disable" {
  namespace               = data.oci_objectstorage_namespace.ns.namespace
  ingest_time_rule_id     = oci_log_analytics_namespace_ingest_time_rule.ingest_time_rule_full.id
  enable_ingest_time_rule = false

  depends_on = [oci_log_analytics_namespace_ingest_time_rule.ingest_time_rule_full]
}

# Enable an ingest time rule
resource "oci_log_analytics_namespace_ingest_time_rules_management" "ingest_time_rule_full_enable" {
  namespace               = data.oci_objectstorage_namespace.ns.namespace
  ingest_time_rule_id     = oci_log_analytics_namespace_ingest_time_rule.ingest_time_rule_full.id
  enable_ingest_time_rule = true

  depends_on = [oci_log_analytics_namespace_ingest_time_rules_management.ingest_time_rule_full_disable]
}

# Fetch an ingest time rule
data "oci_log_analytics_namespace_ingest_time_rule" "ingest_time_rule" {
  namespace        = data.oci_objectstorage_namespace.ns.namespace
  ingest_time_rule_id = oci_log_analytics_namespace_ingest_time_rule.ingest_time_rule_full.id
}

# Fetch all ingest time rules in a compartment (with options and filter)
data "oci_log_analytics_namespace_ingest_time_rules" "ingest_time_rules" {
  compartment_id   = var.compartment_ocid
  namespace        = data.oci_objectstorage_namespace.ns.namespace
  condition_kind   = "FIELD"
  display_name     = "displayName2"
  field_name       = "mtag"
  field_value      = "cveexploitattempt"
  state            = "ACTIVE"

  filter {
    name = "id"
    values = [oci_log_analytics_namespace_ingest_time_rule.ingest_time_rule_full.id]
  }
}

# Fetch all rules (INGEST_TIME and/or SAVED_SEARCH) in a compartment
data "oci_log_analytics_namespace_rules" "rules" {
  compartment_id   = var.compartment_ocid
  namespace        = data.oci_objectstorage_namespace.ns.namespace
  display_name     = "displayName"
  kind             = "INGEST_TIME"
  state            = "ACTIVE"
}

# Fetch all rules (INGEST_TIME and/or SAVED_SEARCH) in a compartment with target_service=MONITORING
data "oci_log_analytics_namespace_rules" "rules_target_monitoring" {
  compartment_id   = var.compartment_ocid
  namespace        = data.oci_objectstorage_namespace.ns.namespace
  display_name     = "displayName"
  kind             = "INGEST_TIME"
  target_service   = "MONITORING"
  state            = "ACTIVE"
}

# Fetch rules summary
data "oci_log_analytics_namespace_rules_summary" "rules_summary" {
  compartment_id   = var.compartment_ocid
  namespace        = data.oci_objectstorage_namespace.ns.namespace
}
