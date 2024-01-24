// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

/*
  This example shows how to add an application associated to a firewall policy.
*/

variable "network_firewall_policy_application_icmp_code" {
  default = 10
}

variable "network_firewall_policy_application_icmp_type" {
  default = 10
}

variable "network_firewall_policy_application_name_icmp" {
  default = "application_icmp"
}

variable "network_firewall_policy_application_name_icmp_v6" {
  default = "application_icmp_v6"
}


variable "network_firewall_policy_application_type_icmp" {
  default = "ICMP"
}

variable "network_firewall_policy_application_type_icmp_v6" {
  default = "ICMP_V6"
}

resource "oci_network_firewall_network_firewall_policy_application" "test_network_firewall_policy_application_icmp" {
  #Required
  icmp_type                  = var.network_firewall_policy_application_icmp_type
  name                       = var.network_firewall_policy_application_name_icmp
  type                       = var.network_firewall_policy_application_type_icmp
  network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id

  #Optional
  icmp_code = var.network_firewall_policy_application_icmp_code
}

resource "oci_network_firewall_network_firewall_policy_application" "test_network_firewall_policy_application_icmp_v6" {
  #Required
  icmp_type                  = var.network_firewall_policy_application_icmp_type
  name                       = var.network_firewall_policy_application_name_icmp_v6
  type                       = var.network_firewall_policy_application_type_icmp_v6
  network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id

  #Optional
  icmp_code = var.network_firewall_policy_application_icmp_code
}

data "oci_network_firewall_network_firewall_policy_applications" "test_network_firewall_policy_applications" {
  #Required
  network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id

  #Optional
  display_name = var.network_firewall_policy_application_name_icmp
}

data "oci_network_firewall_network_firewall_policy_application" "test_network_firewall_policy_application" {
  #Required
  network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id
  name = var.network_firewall_policy_application_name_icmp
}