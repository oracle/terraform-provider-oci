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

provider "oci" {
  auth                = "SecurityToken"
  config_file_profile = "terraform-federation-test"
  tenancy_ocid        = var.tenancy_ocid
  region              = var.region
}

data "oci_data_safe_crypto_assessment_finding_analytics" "test_finding_analytics" {
  compartment_id            = var.compartment_id
  access_level              = "ACCESSIBLE"
  compartment_id_in_subtree = true
}
