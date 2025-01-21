// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

variable "compartment_ocid" {}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

variable "byoasn_asn" {
  default = 11
}

variable "byoasn_display_name" {
  default = "test_byoasn_display_name"
}

// Create a BYOASN resource
resource "oci_core_byoasn" "byoasn_test_resource" {
  asn            = var.byoasn_asn
  compartment_id = var.compartment_ocid
  display_name   = var.byoasn_display_name
}

// Get BYOASN resource details
data "oci_core_byoasn" "byoasn_test_resource" {
  byoasn_id            = oci_core_byoasn.byoasn_test_resource.id
}