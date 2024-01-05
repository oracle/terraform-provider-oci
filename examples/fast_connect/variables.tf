// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "cross_connect_display_name" {
  default = "displayName"
}

variable "cross_connect_state" {
  default = "AVAILABLE"
}

variable "virtual_circuit_type_private" {
  default = "PRIVATE"
}

variable "virtual_circuit_type_public" {
  default = "PUBLIC"
}

variable "virtual_circuit_bandwidth_shape_name" {
  default = "1 Gbps"
}

variable "virtual_circuit_cross_connect_mappings_customer_bgp_peering_ip" {
  default = "10.0.0.18/31"
}

variable "virtual_circuit_cross_connect_mappings_oracle_bgp_peering_ip" {
  default = "10.0.0.19/31"
}

variable "virtual_circuit_cross_connect_mappings_vlan" {
  default = 200
}

variable "virtual_circuit_customer_asn" {
  default = 10
}

variable "virtual_circuit_display_name" {
  default = "displayName"
}

variable "virtual_circuit_public_prefixes_cidr_block" {
  default = "11.0.0.0/24"
}

variable "virtual_circuit_public_prefixes_cidr_block2" {
  default = "11.0.1.0/24"
}

variable "virtual_circuit_public_prefixes_cidr_block3" {
  default = "11.0.2.0/24"
}

variable "virtual_circuit_region" {
  default = "us-phoenix-1"
}

variable "virtual_circuit_state" {
  default = "PROVISIONED"
}

variable "virtual_circuit_cross_connect_mappings_vlan_public" {
  default = 300
}

variable "virtual_circuit_public_prefix_verification_state" {
  default = "COMPLETED"
}

