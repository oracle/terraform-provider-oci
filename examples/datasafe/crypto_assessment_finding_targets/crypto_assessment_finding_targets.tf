// Copyright (c) 2017, 2026, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

terraform {
  required_providers {
    oci = {
      source = "oracle/oci"
    }
  }
}

variable "tenancy_ocid" {}
variable "region" {}
variable "compartment_id" {}
variable "target_id" {}

provider "oci" {
  auth                = "SecurityToken"
  config_file_profile = "terraform-federation-test"
  tenancy_ocid        = var.tenancy_ocid
  region              = var.region
}

data "oci_data_safe_crypto_assessment_finding_targets" "test_finding_targets" {
  compartment_id             = var.compartment_id
  finding_key                = ["CONF.ALLOWED_WEAK_CERT_ALGORITHMS"]
  access_level               = "ACCESSIBLE"
  assessment_type            = "LATEST"
  compartment_id_in_subtree  = true
  is_quantum_readiness_check = false
  target_id                  = var.target_id
}
