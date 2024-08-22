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

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_data_safe_unset_security_assessment_baseline_management" "test_unset_security_assessment_baseline_management" {
    #Required
    security_assessment_id = oci_data_safe_set_security_assessment_baseline_management.test_set_security_assessment_baseline_management.security_assessment_id
    compartment_id = var.compartment_ocid
}