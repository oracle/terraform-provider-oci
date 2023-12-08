// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "data_safe_target_ocid" {}

variable "discovery_analytic_compartment_id_in_subtree" {
  default = true
}

variable "discovery_analytic_group_by" {
  default = "targetId"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

variable "sensitive_data_model_display_name" {
  default = "displayName"
}

variable "sensitive_data_model_schemas_for_discovery" {
  default = []
}

variable "sensitive_data_model_sensitive_type_ids_for_discovery" {
  default = []
}

variable "sensitive_type_entity_type" {
  default = "SENSITIVE_TYPE"
}

variable "sensitive_type_comment_pattern" {
  default = "commentPattern"
}

variable "sensitive_type_data_pattern" {
  default = "dataPattern"
}

variable "sensitive_type_defined_tags_value" {
  default = "value"
}

variable "sensitive_type_description" {
  default = "description"
}

variable "sensitive_type_display_name" {
  default = "displayName"
}

variable "sensitive_type_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "sensitive_type_name_pattern" {
  default = "namePattern"
}

variable "sensitive_type_search_type" {
  default = "OR"
}

variable "sensitive_type_sensitive_type_source" {
  default = "ORACLE"
}

variable "sensitive_type_short_name" {
  default = "shortName"
}

variable "sensitive_type_state" {
  default = "AVAILABLE"
}

variable "sensitive_type_time_created_greater_than_or_equal_to" {
  default = "2018-01-01T00:00:00.000Z"
}

variable "sensitive_type_time_created_less_than" {
  default = "2038-01-01T00:00:00.000Z"
}

resource "oci_data_safe_sensitive_data_model" "test_sensitive_data_model" {
  #Required
  compartment_id = var.compartment_ocid
  target_id      = var.data_safe_target_ocid

  #Optional
  display_name                              = var.sensitive_data_model_display_name
  schemas_for_discovery                     = var.sensitive_data_model_schemas_for_discovery
  sensitive_type_ids_for_discovery          = var.sensitive_data_model_sensitive_type_ids_for_discovery
}

resource "oci_data_safe_sensitive_type" "test_sensitive_type" {
  #Required
  compartment_id = var.compartment_ocid
  entity_type    = var.sensitive_type_entity_type

  #Optional
  comment_pattern           = var.sensitive_type_comment_pattern
  data_pattern              = var.sensitive_type_data_pattern
  description               = var.sensitive_type_description
  display_name              = var.sensitive_type_display_name
  freeform_tags             = var.sensitive_type_freeform_tags
  name_pattern              = var.sensitive_type_name_pattern
  search_type               = var.sensitive_type_search_type
  short_name                = var.sensitive_type_short_name
}

data "oci_data_safe_discovery_analytics" "test_discovery_analytics" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  compartment_id_in_subtree = var.discovery_analytic_compartment_id_in_subtree
  group_by                  = var.discovery_analytic_group_by
  sensitive_data_model_id   = oci_data_safe_sensitive_data_model.test_sensitive_data_model.id
  sensitive_type_id         = oci_data_safe_sensitive_type.test_sensitive_type.id
}

