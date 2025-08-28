// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "security_assessment_ocid" {}
variable "data_safe_target_ocid" {}

provider "oci" {
  auth                = "SecurityToken"
  config_file_profile = "terraform-federation-test"
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}


data "oci_data_safe_security_assessment_checks" "test_security_assessment_checks" {
  #Required
  security_assessment_id = var.security_assessment_ocid
}