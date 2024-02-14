// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

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

variable "config_file_profile" {
  default = ""
}

provider "oci" {
  region           = var.region
  auth = "SecurityToken"
  config_file_profile = var.config_file_profile
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

variable "on_prem_path_analyzer_test_destination_endpoint_type" {
  default = "ON_PREM"
}
variable "on_prem_path_analyzer_test_destination_endpoint_address" {
  default = "10.2.41.5"
}
variable "on_prem_path_analyzer_test_protocol" {
  default = 6
}
variable "on_prem_path_analyzer_test_source_endpoint_type" {
  default = "IP_ADDRESS"
}
variable "on_prem_path_analyzer_test_source_endpoint_address" {
  default = "100.130.10.100"
}
variable "on_prem_path_analyzer_test_display_name" {
  default = "On Prem Path Analysis"
}
variable "on_prem_path_analyzer_test_freeform_tags" {
  default = { "bar-key" = "value" }
}
variable "on_prem_path_analyzer_test_protocol_parameters_type" {
  default = "TCP"
}
variable "on_prem_path_analyzer_test_protocol_parameters_destination_port" {
  default = 85
}
variable "on_prem_path_analyzer_test_protocol_parameters_source_port" {
  default = 84
}
variable "on_prem_path_analyzer_test_query_options_is_bi_directional_analysis" {
  default = false
}
resource "oci_vn_monitoring_path_analysi" "on_prem_test_path_analysis" {
  #Required
  compartment_id = var.compartment_ocid
  type = var.path_analysi_type
  destination_endpoint {
    #Required
    type = var.on_prem_path_analyzer_test_destination_endpoint_type

    #Optional
    address = var.on_prem_path_analyzer_test_destination_endpoint_address

  }
  protocol = var.on_prem_path_analyzer_test_protocol
  source_endpoint {
    #Required
    type = var.on_prem_path_analyzer_test_source_endpoint_type

    #Optional
    address = var.on_prem_path_analyzer_test_source_endpoint_address
  }

  protocol_parameters {
    #Required
    type = var.on_prem_path_analyzer_test_protocol_parameters_type

    #Optional
    destination_port = var.on_prem_path_analyzer_test_protocol_parameters_destination_port
    source_port      = var.on_prem_path_analyzer_test_protocol_parameters_source_port
  }
  query_options {

    #Optional
    is_bi_directional_analysis = var.on_prem_path_analyzer_test_query_options_is_bi_directional_analysis
  }
}

