// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "data_safe_target_ocid" {}

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

variable "masking_policy_access_level" {
  default = "ACCESSIBLE"
}

variable "masking_policy_column_source_column_source" {
  default = "SENSITIVE_DATA_MODEL"
}

variable "masking_policy_compartment_id_in_subtree" {
  default = true
}

variable "masking_policy_description" {
  default = "description"
}

variable "masking_policy_display_name" {
  default = "displayName"
}

variable "masking_policy_state" {
  default = "ACTIVE"
}

variable "masking_policy_time_created_greater_than_or_equal_to" {
  default = "2018-01-01T00:00:00.000Z"
}

variable "masking_policy_time_created_less_than" {
  default = "2038-01-01T00:00:00.000Z"
}

variable "difference_discovery_type" {
  default = "ALL"
}

variable "difference_job_state" {
  default = "ACTIVE"
}

variable "difference_freeform_tags" {
  default = { "Department" = "Finance" }
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

resource "oci_data_safe_masking_policy" "test_masking_policy" {
  #Required
  column_source {
    #Required
    column_source = var.masking_policy_column_source_column_source

    #Optional
    sensitive_data_model_id = oci_data_safe_sensitive_data_model.test_sensitive_data_model.id
  }
  compartment_id = var.compartment_ocid

  #Optional
  description                 = var.masking_policy_description
  display_name                = var.masking_policy_display_name
}

resource "oci_data_safe_sdm_masking_policy_difference" "test_sdm_masking_policy_difference" {
  #Required
  compartment_id          = var.compartment_ocid
  masking_policy_id = oci_data_safe_masking_policy.test_masking_policy.id

}

data "oci_data_safe_sdm_masking_policy_difference_difference_columns" "test_sdm_masking_policy_difference_difference_columns" {
  #Required
  sdm_masking_policy_difference_id = oci_data_safe_sdm_masking_policy_difference.test_sdm_masking_policy_difference.id
}

