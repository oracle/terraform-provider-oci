// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "region" {
}

variable "compartment_ocid" {
}

variable "data_safe_target_ocid" {
}
variable "security_assessment_ocid" {

}
variable "comp_security_assessment_ocid" {

}
provider "oci" {
  version          = "7.12.0"
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_data_safe_security_assessment_comparison" "test_security_assessment_comparison" {
  #Required
  comparison_security_assessment_id = var.comp_security_assessment_ocid
  security_assessment_id = var.security_assessment_ocid
}