// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0
variable "compartment_id" {
}

variable "network_source_defined_tags_value" {
  default = "value"
}

variable "network_source_description" {
  default = "corporate ip ranges to be used for ip based authorization"
}

variable "network_source_freeform_tags" {
  default = {
    "Department" = "Finance"
  }
}

variable "network_source_name" {
  default = "corpnet"
}

variable "network_source_public_source_list" {
  default = ["128.2.13.5"]
}

variable "network_source_services" {
  default = ["all"]
}

variable "network_source_virtual_source_list" {
  default = []
}

resource "oci_core_vcn" "vcn1" {
  cidr_block     = "10.0.0.0/16"
  dns_label      = "vcn1"
  compartment_id = var.compartment_id
  display_name   = "vcn1"
}

resource "oci_identity_network_source" "test_network_source" {
  #Required
  compartment_id = var.tenancy_ocid
  description    = var.network_source_description
  name           = var.network_source_name

  #Optional
  freeform_tags      = var.network_source_freeform_tags
  public_source_list = var.network_source_public_source_list
  services           = var.network_source_services

  virtual_source_list {
    vcn_id    = oci_core_vcn.vcn1.id
    ip_ranges = ["10.0.0.0/16"]
  }
}

data "oci_identity_network_sources" "test_network_sources" {
  #Required
  compartment_id = var.tenancy_ocid

  filter {
    name   = "id"
    values = [oci_identity_network_source.test_network_source.id]
  }
}

