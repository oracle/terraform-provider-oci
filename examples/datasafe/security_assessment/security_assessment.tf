// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "security_assessment_id" {}
variable "data_safe_target_ocid" {}


variable "display_name" {
  default = "SA_1"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_data_safe_security_assessment" "oci_data_safe_security_assessment" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  target_id = var.data_safe_target_ocid
  display_name = var.display_name

  lifecycle {
    ignore_changes = [system_tags]
  }
}

data "oci_data_safe_security_assessments" "test_security_assessments" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  target_id = var.data_safe_target_ocid

}
