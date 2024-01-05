// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "data_safe_target_ocid" {}

variable "sensitive_data_model_access_level" {
  default = "ACCESSIBLE"
}

variable "sensitive_data_model_app_suite_name" {
  default = "appSuiteName"
}

variable "sensitive_data_model_compartment_id_in_subtree" {
  default = true
}

variable "sensitive_data_model_description" {
  default = "description"
}

variable "sensitive_data_model_display_name" {
  default = "displayName"
}

variable "sensitive_data_model_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "sensitive_data_model_is_app_defined_relation_discovery_enabled" {
  default = false
}

variable "sensitive_data_model_is_sample_data_collection_enabled" {
  default = false
}

variable "sensitive_data_model_schemas_for_discovery" {
  default = []
}

variable "sensitive_data_model_sensitive_type_ids_for_discovery" {
  default = []
}

variable "sensitive_data_model_state" {
  default = "ACTIVE"
}

variable "sensitive_data_model_time_created_less_than" {
  default = "2038-01-01T00:00:00.000Z"
}


provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_data_safe_sensitive_data_model" "test_sensitive_data_model" {
  #Required
  compartment_id = var.compartment_ocid
  target_id      = var.data_safe_target_ocid

  #Optional
  app_suite_name                            = var.sensitive_data_model_app_suite_name
  description                               = var.sensitive_data_model_description
  display_name                              = var.sensitive_data_model_display_name
  freeform_tags                             = var.sensitive_data_model_freeform_tags
  is_app_defined_relation_discovery_enabled = var.sensitive_data_model_is_app_defined_relation_discovery_enabled
  is_sample_data_collection_enabled         = var.sensitive_data_model_is_sample_data_collection_enabled
  schemas_for_discovery                     = var.sensitive_data_model_schemas_for_discovery
  sensitive_type_ids_for_discovery          = var.sensitive_data_model_sensitive_type_ids_for_discovery
}

data "oci_data_safe_sensitive_data_models" "test_sensitive_data_models" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  access_level                          = var.sensitive_data_model_access_level
  compartment_id_in_subtree             = var.sensitive_data_model_compartment_id_in_subtree
  display_name                          = var.sensitive_data_model_display_name
  sensitive_data_model_id               = oci_data_safe_sensitive_data_model.test_sensitive_data_model.id
  state                                 = var.sensitive_data_model_state
  time_created_less_than                = var.sensitive_data_model_time_created_less_than
}

