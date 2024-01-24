// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "database_security_config_ocid" {}
variable "data_safe_target_ocid" {}

variable "database_security_config_access_level" {
  default = "RESTRICTED"
}

variable "database_security_config_compartment_id_in_subtree" {
  default = false
}

variable "database_security_config_defined_tags_value" {
  default = "value"
}

variable "database_security_config_description" {
  default = "updated-description"
}

variable "database_security_config_display_name" {
  default = "updated-name"
}

variable "database_security_config_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "database_security_config_status" {
  default = "ACTIVE"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_data_safe_database_security_config_management" "test_database_security_config_management" {
  #Required
  compartment_id = var.compartment_ocid
  target_id = var.data_safe_target_ocid

  #Optional
  description           = var.database_security_config_description
  display_name          = var.database_security_config_display_name
  freeform_tags         = var.database_security_config_freeform_tags
}

