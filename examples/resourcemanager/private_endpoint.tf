// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0


variable "private_endpoint_defined_tags_value" {
  default = "value"
}

variable "private_endpoint_description" {
  default = "Example Private Endpoint"
}

variable "private_endpoint_display_name" {
  default = "My Private Endpoint"
}

variable "private_endpoint_dns_zones" {
  default = []
}

variable "private_endpoint_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "private_endpoint_is_used_with_configuration_source_provider" {
  default = false
}

variable "private_endpoint_nsg_id_list" {
  default = []
}


resource "oci_resourcemanager_private_endpoint" "test_private_endpoint" {
  #Required
  compartment_id = var.compartment_ocid
  display_name   = var.private_endpoint_display_name
  subnet_id      = oci_core_subnet.test_subnet.id
  vcn_id         = oci_core_vcn.test_vcn.id

  description                                = var.private_endpoint_description
  dns_zones                                  = var.private_endpoint_dns_zones
  freeform_tags                              = var.private_endpoint_freeform_tags
  is_used_with_configuration_source_provider = var.private_endpoint_is_used_with_configuration_source_provider
  nsg_id_list                                = var.private_endpoint_nsg_id_list
}

resource "oci_core_vcn" "test_vcn" {
  cidr_block     = local.vcn_cidr_block
  compartment_id = var.compartment_ocid
  display_name   = "test_vcn"
}

resource "oci_core_subnet" "test_subnet" {
  compartment_id             = var.compartment_ocid
  vcn_id                     = oci_core_vcn.test_vcn.id
  display_name               = "test_subnet"
  prohibit_public_ip_on_vnic = true
  cidr_block                 = cidrsubnet(local.vcn_cidr_block, 8, 1)
}

data "oci_resourcemanager_private_endpoints" "test_private_endpoints" {

  #Optional
  compartment_id      = var.compartment_ocid
  display_name        = var.private_endpoint_display_name
  private_endpoint_id = oci_resourcemanager_private_endpoint.test_private_endpoint.id
  vcn_id              = oci_core_vcn.test_vcn.id
}


