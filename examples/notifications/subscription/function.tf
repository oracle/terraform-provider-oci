// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "config" {
  default = {
    "MY_FUNCTION_CONFIG" = "ConfVal"
  }
}

variable "image" {
}

variable "image_digest" {
}

data "oci_identity_availability_domains" "test_availability_domains" {
  compartment_id = var.tenancy_ocid
}

resource "oci_core_vcn" "test_vcn" {
  cidr_block     = "10.0.0.0/16"
  compartment_id = var.compartment_ocid
  display_name   = "tf-vcn"
  dns_label      = "dnslabel"
}

resource "oci_core_internet_gateway" "test_network_entity" {
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.test_vcn.id
  display_name   = "-tf-internet-gateway"
}

resource "oci_core_route_table" "test_route_table" {
  compartment_id = var.compartment_ocid

  route_rules {
    cidr_block        = "0.0.0.0/0"
    network_entity_id = oci_core_internet_gateway.test_network_entity.id
  }

  vcn_id = oci_core_vcn.test_vcn.id
}

resource "oci_core_subnet" "test_subnet" {
  availability_domain = lower(
    data.oci_identity_availability_domains.test_availability_domains.availability_domains[0].name,
  )
  cidr_block                 = "10.0.0.0/16"
  compartment_id             = var.compartment_ocid
  dhcp_options_id            = oci_core_vcn.test_vcn.default_dhcp_options_id
  display_name               = "tf-subnet"
  dns_label                  = "dnslabel"
  prohibit_public_ip_on_vnic = "false"
  route_table_id             = oci_core_route_table.test_route_table.id
  security_list_ids          = [oci_core_vcn.test_vcn.default_security_list_id]
  vcn_id                     = oci_core_vcn.test_vcn.id
}

resource "oci_functions_application" "test_application" {
  #Required
  compartment_id = var.compartment_ocid
  display_name   = "example-application"
  subnet_ids     = [oci_core_subnet.test_subnet.id]

  #Optional
  config = var.config
}

resource "oci_functions_function" "test_function" {
  #Required
  application_id = oci_functions_application.test_application.id
  display_name   = "example-function"
  image          = var.image
  memory_in_mbs  = "128"

  #Optional
  config             = var.config
  image_digest       = var.image_digest
  timeout_in_seconds = "30"
}

