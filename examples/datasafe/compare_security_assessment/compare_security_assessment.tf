// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}





provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}


resource "oci_data_safe_compare_security_assessment" "test_compare_security_assessment" {
    comparison_security_assessment_id =  oci_data_safe_security_assessment.test_security_assessment1.id
    security_assessment_id            =  oci_data_safe_security_assessment.test_security_assessment2.id
}
