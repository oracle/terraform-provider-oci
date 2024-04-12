// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

variable "compartment_ocid" {
}

variable "vcn_ocid" {
}

variable "subnet_ocid" {
}

variable "ip_inventory_vcn_overlap_compartment_list" {
  type = list(string)
  default = []
}

variable "ip_inventory_vcn_overlap_region_list" {
  type = list(string)
  default = []
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

// Get VCN Overlap
data "oci_core_ip_inventory_vcn_overlaps" "test_ip_inventory_vcn_overlaps" {
  #Required
  compartment_list = var.ip_inventory_vcn_overlap_compartment_list
  region_list      = var.ip_inventory_vcn_overlap_region_list
  vcn_id           = var.vcn_ocid
}

// Subnet CIDR Utilisation
data "oci_core_ip_inventory_subnet_cidr" "test_ip_inventory_subnet_cidrs" {
  #Required
  subnet_id = var.subnet_ocid
}

// Subnet Details
data "oci_core_ip_inventory_subnet" "test_ip_inventory_subnets" {
  #Required
  subnet_id = var.subnet_ocid
}
