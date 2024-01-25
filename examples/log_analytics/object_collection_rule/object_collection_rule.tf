// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * This example shows how to manage log analytics object collection rule resource
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

variable "log_analytics_log_group_id" {}
variable "log_analytics_entity_id" {}
variable "object_collection_rule_bucket_name" {}
variable "object_collection_rule_bucket_name_log_events" {}

variable "object_collection_rule_name" {
  default = "tf-obj-coll-example-opt"
}
variable "object_collection_rule_name_log_events" {
  default = "tf-obj-coll-example-log-events"
}
variable "object_collection_rule_freeform_tags" {
  default = { "servicegroup" = "test", "Dept" = "Devops" }
}
variable "object_collection_rule_log_type" {
  default = "LOG"
}
variable "object_collection_rule_log_source_name" {
  default = "LinuxSyslogSource"
}
variable "object_collection_rule_description" {
  default = "Object Collection Rule with optional parameters"
}
variable "object_collection_rule_collection_type" {
  default = "HISTORIC"
}
variable "object_collection_rule_is_force_historic_collection" {
  default = "false"
}
variable "object_collection_rule_poll_since" {
  default = "2020-04-01T00:00:00.000Z"
}
variable "object_collection_rule_poll_till" {
  default = "2021-04-01T00:00:00.000Z"
}
variable "object_collection_rule_char_encoding" {
  default = "utf-8"
}
variable "object_collection_rule_log_source_override_match_value" {
  default = "db"
}
variable "object_collection_rule_log_source_override_property_value" {
  default = "DBAuditLogSource"
}
variable "object_collection_rule_char_encoding_override_match_value" {
  default = "db"
}
variable "object_collection_rule_char_encoding_override_property_value" {
  default = "utf-16"
}
variable "object_collection_rule_object_name_filter" {
  default = "*"
}
variable "object_collection_rule_log_set" {
  default = "logSet2"
}
variable "object_collection_rule_log_set_ext_regex" {
  default = "^(.+)\\/([^\\/]+)$"
}
variable "object_collection_rule_log_set_key" {
  default = "OBJECT_PATH"
}
variable "object_collection_rule_timezone" {
  default = "Asia/Dhaka"
}
variable "object_collection_rule_is_enabled" {
  default = "true"
}

# Create a object collection rule with required parameters
resource "oci_log_analytics_log_analytics_object_collection_rule" "objectCollectionRuleRequired" {
  compartment_id             = var.compartment_ocid
  namespace                  = data.oci_objectstorage_namespace.ns.namespace
  name                       = "tf-obj-coll-example-req"
  log_group_id               = var.log_analytics_log_group_id
  log_source_name            = "LinuxSyslogSource"
  os_bucket_name             = var.object_collection_rule_bucket_name
  os_namespace               = data.oci_objectstorage_namespace.ns.namespace
}

# Get details of above created object collection rule with required parameters
data "oci_log_analytics_log_analytics_object_collection_rule" "objectCollectionRuleRequiredDetails" {
  namespace                  = data.oci_objectstorage_namespace.ns.namespace
  log_analytics_object_collection_rule_id = oci_log_analytics_log_analytics_object_collection_rule.objectCollectionRuleRequired.id
}

# Create a object collection rule with optional parameters
resource "oci_log_analytics_log_analytics_object_collection_rule" "objectCollectionRuleOptional" {
  compartment_id             = var.compartment_ocid
  namespace                  = data.oci_objectstorage_namespace.ns.namespace
  name                       = var.object_collection_rule_name
  log_group_id               = var.log_analytics_log_group_id
  log_type                   = var.object_collection_rule_log_type
  log_source_name            = var.object_collection_rule_log_source_name
  os_bucket_name             = var.object_collection_rule_bucket_name
  os_namespace               = data.oci_objectstorage_namespace.ns.namespace
  description                = var.object_collection_rule_description
  collection_type            = var.object_collection_rule_collection_type
  is_force_historic_collection = var.object_collection_rule_is_force_historic_collection
  log_set                    = var.object_collection_rule_log_set
  log_set_ext_regex          = var.object_collection_rule_log_set_ext_regex
  log_set_key                = var.object_collection_rule_log_set_key
  timezone                   = var.object_collection_rule_timezone
  poll_since                 = var.object_collection_rule_poll_since
  poll_till                  = var.object_collection_rule_poll_till
  char_encoding              = var.object_collection_rule_char_encoding
  freeform_tags              = var.object_collection_rule_freeform_tags
  overrides {
    match_type       = "contains"
    match_value      = var.object_collection_rule_char_encoding_override_match_value
    property_name    = "charEncoding"
    property_value   = var.object_collection_rule_char_encoding_override_property_value
  }
  overrides {
    match_type       = "contains"
    match_value      = var.object_collection_rule_log_source_override_match_value
    property_name    = "logSourceName"
    property_value   = var.object_collection_rule_log_source_override_property_value
  }
  object_name_filters              = [var.object_collection_rule_object_name_filter]
  is_enabled = var.object_collection_rule_is_enabled
}

# Create a object collection rule with optional parameters
resource "oci_log_analytics_log_analytics_object_collection_rule" "objectCollectionRuleLogEventsType" {
  compartment_id             = var.compartment_ocid
  namespace                  = data.oci_objectstorage_namespace.ns.namespace
  name                       = var.object_collection_rule_name_log_events
  log_group_id               = var.log_analytics_log_group_id
  log_source_name            = var.object_collection_rule_log_source_name
  os_bucket_name             = var.object_collection_rule_bucket_name_log_events
  os_namespace               = data.oci_objectstorage_namespace.ns.namespace
  collection_type            = var.object_collection_rule_collection_type
  log_set                    = var.object_collection_rule_log_set
  poll_since                 = var.object_collection_rule_poll_since
  poll_till                  = var.object_collection_rule_poll_till
  freeform_tags              = var.object_collection_rule_freeform_tags
  log_type                   = "LOG_EVENTS"
  is_enabled = var.object_collection_rule_is_enabled
}

# Get details of above created object collection rule with optional parameters
data "oci_log_analytics_log_analytics_object_collection_rule" "objectCollectionRuleOptionalDetails" {
  namespace                  = data.oci_objectstorage_namespace.ns.namespace
  log_analytics_object_collection_rule_id = oci_log_analytics_log_analytics_object_collection_rule.objectCollectionRuleOptional.id
}

# Get the list of object collection rules in a compartment
data "oci_log_analytics_log_analytics_object_collection_rules" "objectCollectionRulessList" {
  compartment_id             = var.compartment_ocid
  namespace                  = data.oci_objectstorage_namespace.ns.namespace
}
