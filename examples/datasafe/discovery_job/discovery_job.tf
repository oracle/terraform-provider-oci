// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "data_safe_target_ocid" {}


variable "discovery_job_defined_tags_value" {
  default = "value"
}

variable "discovery_job_discovery_type" {
  default = "ALL"
}

variable "discovery_job_display_name" {
  default = "displayName"
}

variable "discovery_job_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "discovery_job_is_app_defined_relation_discovery_enabled" {
  default = false
}

variable "discovery_job_is_include_all_schemas" {
  default = true
}

variable "discovery_job_is_include_all_sensitive_types" {
  default = true
}

variable "discovery_job_state" {
  default = "ACTIVE"
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

resource "oci_data_safe_sensitive_data_model" "test_sensitive_data_model" {
  #Required
  compartment_id = var.compartment_ocid
  target_id      = var.data_safe_target_ocid

  #Optional
  display_name                              = var.sensitive_data_model_display_name
  schemas_for_discovery                     = var.sensitive_data_model_schemas_for_discovery
  sensitive_type_ids_for_discovery          = var.sensitive_data_model_sensitive_type_ids_for_discovery
}

resource "oci_data_safe_discovery_job" "test_discovery_job" {
  #Required
  compartment_id          = var.compartment_ocid
  sensitive_data_model_id = oci_data_safe_sensitive_data_model.test_sensitive_data_model.id

  #Optional
  discovery_type                            = var.discovery_job_discovery_type
  freeform_tags                             = var.discovery_job_freeform_tags
  is_app_defined_relation_discovery_enabled = var.discovery_job_is_app_defined_relation_discovery_enabled
  is_include_all_schemas                    = var.discovery_job_is_include_all_schemas
  is_include_all_sensitive_types            = var.discovery_job_is_include_all_sensitive_types
}

data "oci_data_safe_discovery_jobs" "test_discovery_jobs" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  discovery_job_id          = oci_data_safe_discovery_job.test_discovery_job.id
  sensitive_data_model_id   = oci_data_safe_sensitive_data_model.test_sensitive_data_model.id
  state                     = var.discovery_job_state
}

