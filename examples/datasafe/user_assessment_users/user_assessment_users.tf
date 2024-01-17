// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "user_assessment_ocid" {}
variable "data_safe_target_ocid" {}


variable "description" {
  default = "description"
}

variable "display_name" {
  default = "UA_1"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_data_safe_user_assessment" "test_user_assessment" {
  #Required
  compartment_id      = var.compartment_ocid
  target_id           = var.data_safe_target_ocid

  #Optional
  description   = var.description
  display_name  = var.display_name
}

data "oci_data_safe_user_assessment_users" "test_user_assessment_users" {
  #Required
  user_assessment_id = var.user_assessment_ocid
}