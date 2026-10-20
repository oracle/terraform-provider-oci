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
variable "crypto_assessment_id" {}
variable "target_id" {}

provider "oci" {
  auth                = "SecurityToken"
  config_file_profile = "terraform-federation-test"
  tenancy_ocid        = var.tenancy_ocid
  region              = var.region
}

data "oci_data_safe_crypto_assessment_wallets" "test_wallets" {
  compartment_id            = var.compartment_id
  access_level              = "ACCESSIBLE"
  assessment_id             = var.crypto_assessment_id
  assessment_type           = "LATEST"
  compartment_id_in_subtree = true
  target_id                 = var.target_id
}
