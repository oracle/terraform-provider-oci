// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

variable "compartment_ocid" {}
variable "byoip_ocid" {}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}


// Get BYOIP resource details
data "oci_core_byoip_range" "test_byoip_resource" {
  byoip_range_id            = var.byoip_ocid
}

// List BYOIP resources details
data "oci_core_byoip_ranges" "test_byoip_resource" {
  compartment_id            = var.compartment_ocid
}

// GET byoip allocated ranges details
data "oci_core_byoip_allocated_ranges" "test_byoip_allocated_ranges" {
  byoip_range_id            = var.byoip_ocid
}

