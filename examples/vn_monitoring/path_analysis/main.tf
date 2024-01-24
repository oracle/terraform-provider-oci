// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}

variable "path_analysi_destination_endpoint_type" {
  default = "SUBNET"
}

variable "path_analysi_protocol" {
  default = 1
}

variable "path_analysi_protocol_parameters_destination_port" {
  default = 10
}

variable "path_analysi_protocol_parameters_icmp_code" {
  default = 10
}

variable "path_analysi_protocol_parameters_icmp_type" {
  default = 10
}

variable "path_analysi_protocol_parameters_source_port" {
  default = 10
}

variable "path_analysi_protocol_parameters_type" {
  default = "ICMP"
}

variable "path_analysi_query_options_is_bi_directional_analysis" {
  default = false
}

variable "path_analysi_source_endpoint_type" {
  default = "SUBNET"
}

variable "path_analysi_type" {
  default = "ADHOC_QUERY"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_core_vcn" "test_vcn" {
  cidr_block     = "10.1.0.0/16"
  compartment_id = var.compartment_ocid
  display_name   = "testVcn"
  dns_label      = "testvcn"
}

resource "oci_core_subnet" "test_subnet" {
  cidr_block          = "10.1.20.0/24"
  display_name        = "testSubnet1"
  dns_label           = "testsubnet1"
  compartment_id      = var.compartment_ocid
  vcn_id              = oci_core_vcn.test_vcn.id
}

resource "oci_core_subnet" "test_subnet2" {
  cidr_block          = "10.1.21.0/24"
  display_name        = "testSubnet2"
  dns_label           = "testsubnet2"
  compartment_id      = var.compartment_ocid
  vcn_id              = oci_core_vcn.test_vcn.id
}

resource "oci_vn_monitoring_path_analysi" "test_path_analysi" {
  #Required
  type = var.path_analysi_type

  #Optional
  compartment_id = var.compartment_ocid
  destination_endpoint {
    #Required
    type = var.path_analysi_destination_endpoint_type

    #Optional
    address                  = "10.1.20.55"
    subnet_id                = oci_core_subnet.test_subnet.id
  }
  protocol              = var.path_analysi_protocol
  protocol_parameters {
    #Required
    type = var.path_analysi_protocol_parameters_type

    #Optional
    destination_port = var.path_analysi_protocol_parameters_destination_port
    icmp_code        = var.path_analysi_protocol_parameters_icmp_code
    icmp_type        = var.path_analysi_protocol_parameters_icmp_type
    source_port      = var.path_analysi_protocol_parameters_source_port
  }
  query_options {

    #Optional
    is_bi_directional_analysis = var.path_analysi_query_options_is_bi_directional_analysis
  }
  source_endpoint {
    #Required
    type = var.path_analysi_source_endpoint_type

    #Optional
    address                  = "10.1.21.55"
    subnet_id                = oci_core_subnet.test_subnet2.id
  }
}

