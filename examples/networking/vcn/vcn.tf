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

variable "compartment_ocid" {
}

variable "region" {
}

provider "oci" {

  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
  auth = "SecurityToken"
  config_file_profile = "terraform-federation-test"
  ignore_defined_tags      = ["testexamples-tag-namespace.tf-example-tag"]
}

resource "oci_core_vcn" "vcn" {
  cidr_blocks    = ["10.0.0.0/16","11.0.0.0/16"]
  dns_label      = "vcn"
  compartment_id = var.compartment_ocid
  display_name   = "vcn"
  security_attributes = {"sample-namespace.value": "examplevalue", "sample-namespace.mode": "examplemode"}
}

output "vcn_id" {
  value = oci_core_vcn.vcn.id
}

