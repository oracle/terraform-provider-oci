// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

/*
  This example shows how to add address lists associated to a firewall policy.
*/

variable "network_firewall_policy_address_list_ip_addresses" {
  default = ["10.180.1.0/24", "10.180.2.0/25"]
}

variable "network_firewall_policy_address_list_fqdn_addresses" {
  default = ["oracle.com", "example.com"]
}

variable "network_firewall_policy_address_list_name_ip" {
  default = "address_list_1_ip"
}

variable "network_firewall_policy_address_list_name_fqdn" {
  default = "address_list_1_fqdn"
}

variable "network_firewall_policy_address_list_ip_type" {
  default = "IP"
}

variable "network_firewall_policy_address_list_fqdn_type" {
  default = "FQDN"
}

resource "oci_network_firewall_network_firewall_policy_address_list" "test_network_firewall_policy_address_list_ip" {
  #Required
  name                       = var.network_firewall_policy_address_list_name_ip
  network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id
  type                       = var.network_firewall_policy_address_list_ip_type
  addresses                  = var.network_firewall_policy_address_list_ip_addresses
}

resource "oci_network_firewall_network_firewall_policy_address_list" "test_network_firewall_policy_address_list_fqdn" {
  #Required
  name                       = var.network_firewall_policy_address_list_name_fqdn
  network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id
  type                       = var.network_firewall_policy_address_list_fqdn_type
  addresses                  = var.network_firewall_policy_address_list_fqdn_addresses
}

data "oci_network_firewall_network_firewall_policy_address_lists" "test_network_firewall_policy_address_lists" {
  #Required
  network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id

  #Optional
  display_name = var.network_firewall_policy_address_list_name_ip
}

data "oci_network_firewall_network_firewall_policy_address_list" "test_network_firewall_policy_address_list" {
  #Required
  network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id
  name = var.network_firewall_policy_address_list_name_ip
}
