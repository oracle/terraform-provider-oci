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

resource "oci_core_vcn" "test_vcn" {
  cidr_block     = "10.0.0.0/16"
  dns_label      = "vcn1"
  compartment_id = var.compartment_ocid
  display_name   = "vcn1"
}

data "oci_core_vcn_dns_resolver_association" "test_vcn_dns_resolver_association" {
  #Required
  vcn_id = oci_core_vcn.test_vcn.id
}
