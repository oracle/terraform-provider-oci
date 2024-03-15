// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

variable "recovery_service_subnet_defined_tags_value" {
  default = "value"
}

variable "recovery_service_subnet_display_name" {
  default = "displayName"
}

variable "recovery_service_subnet_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "recovery_service_subnet_id" {
  default = "id"
}

variable "recovery_service_subnet_state" {
  default = "ACTIVE"
}

variable "recovery_service_subnet_nsg_ids" {
  default = []
}

resource "oci_core_vcn" "test_vcn" {
  cidr_block     = "10.0.0.0/16"
  compartment_id = var.compartment_id
  display_name   = "exampleVCN"
  dns_label      = "tfexamplevcn"
}

resource "oci_core_subnet" "test_subnet" {
  cidr_block        = "10.0.0.0/24"
  display_name      = "tfexampleSubnet"
  dns_label         = "tfexampleSubnet"
  compartment_id    = var.compartment_id
  vcn_id            = oci_core_vcn.test_vcn.id
  security_list_ids = [oci_core_vcn.test_vcn.default_security_list_id]
  route_table_id    = oci_core_vcn.test_vcn.default_route_table_id
  dhcp_options_id   = oci_core_vcn.test_vcn.default_dhcp_options_id
}


resource "oci_recovery_recovery_service_subnet" "test_recovery_service_subnet" {
  #Required
  compartment_id = var.compartment_id
  display_name   = var.recovery_service_subnet_display_name
  subnets        = tolist([oci_core_subnet.test_subnet.id])
  vcn_id         = oci_core_vcn.test_vcn.id
  nsg_ids        = var.recovery_service_subnet_nsg_ids

  #Optional
  freeform_tags = var.recovery_service_subnet_freeform_tags
}

data "oci_recovery_recovery_service_subnets" "test_recovery_service_subnets" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name = var.recovery_service_subnet_display_name
  id           = var.recovery_service_subnet_id
  state        = var.recovery_service_subnet_state
  vcn_id       = oci_core_vcn.test_vcn.id
}