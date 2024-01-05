// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "data_safe_target_ocid" {}


variable "masking_policies_masking_column_masking_formats_format_entries_end_length" {
  default = 10
}

variable "masking_policies_masking_column_masking_formats_format_entries_start_length" {
  default = 10
}

variable "masking_policies_masking_column_masking_formats_format_entries_type" {
  default = "DELETE_ROWS"
}

variable "masking_policies_masking_column_object_var" {
  default = "LOCATIONS"
}

variable "masking_policies_masking_column_object_list" {
  type = list(string)
  default = ["LOCATIONS"]
}


variable "masking_policies_masking_column_schema_name_var" {
  default = "ADMIN"
}

variable "masking_policies_masking_column_schema_name_list" {
  type = list(string)
  default = ["ADMIN"]
}

variable "masking_policies_masking_column_column_name_var" {
  default = "STREET_ADDRESS"
}

variable "masking_policies_masking_column_column_name_list" {
  type = list(string)
  default = ["STREET_ADDRESS"]
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

variable "masking_policy_column_source_column_source" {
  default = "TARGET"
}
variable "masking_policy_description" {
  default = "description"
}

variable "masking_policy_display_name" {
  default = "displayName"
}

resource "oci_data_safe_masking_policy" "test_masking_policy" {
  #Required
  column_source {
    #Required
    column_source = var.masking_policy_column_source_column_source

    #Optional
    target_id               = var.data_safe_target_ocid
  }
  compartment_id = var.compartment_ocid

  #Optional
  description                 = var.masking_policy_description
  display_name                = var.masking_policy_display_name
}

resource "oci_data_safe_masking_policies_masking_column" "test_masking_policies_masking_column" {
  #Required
  column_name       = var.masking_policies_masking_column_column_name_var
  masking_policy_id = oci_data_safe_masking_policy.test_masking_policy.id
  object            = var.masking_policies_masking_column_object_var
  schema_name       = var.masking_policies_masking_column_schema_name_var

  #Optional
  masking_formats {
    #Required
    format_entries {
      #Required
      type = var.masking_policies_masking_column_masking_formats_format_entries_type

      #Optional
      end_length                = var.masking_policies_masking_column_masking_formats_format_entries_end_length
      start_length              = var.masking_policies_masking_column_masking_formats_format_entries_start_length
    }
  }
}

data "oci_data_safe_masking_policies_masking_columns" "test_masking_policies_masking_columns" {
  #Required
  masking_policy_id = oci_data_safe_masking_policy.test_masking_policy.id
}

