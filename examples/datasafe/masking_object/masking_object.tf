// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "data_safe_target_ocid" {}


variable "masking_policies_masking_column_column_name_var" {
  default = "FIRST_NAME"
}

variable "masking_policies_masking_column_schema_name_var" {
  default = "HCM"
}

variable "masking_policies_masking_column_object_var" {
  default = "EMPLOYEES"
}

// ---------------


variable "masking_policy_masking_object_object_list" {
  type = list(string)
  default = ["EMPLOYEES"]
}


provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

variable "masking_policy_display_name" {
  default = "displayName"
}

variable "masking_policy_column_source_column_source" {
  default = "TARGET"
}
variable "masking_policy_description" {
  default = "description"
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
}

data "oci_data_safe_masking_policy_masking_objects" "test_masking_policy_masking_objects" {
  #Required
  masking_policy_id = oci_data_safe_masking_policy.test_masking_policy.id
   object                           = var.masking_policy_masking_object_object_list
}


