// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

/*
  This example shows how to create an application group associated to a firewall policy.
*/

variable "network_firewall_policy_application_group_name" {
  default = "application_group_1"
}

resource "oci_network_firewall_network_firewall_policy_application_group" "test_network_firewall_policy_application_group" {
  #Required
  apps                       = [oci_network_firewall_network_firewall_policy_application.test_network_firewall_policy_application_icmp.name]
  name                       = var.network_firewall_policy_application_group_name
  network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id
}

data "oci_network_firewall_network_firewall_policy_application_groups" "test_network_firewall_policy_application_groups" {
  #Required
  network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id

  #Optional
  display_name = var.network_firewall_policy_application_group_name
}

data "oci_network_firewall_network_firewall_policy_application_group" "test_network_firewall_policy_application_group" {
  #Required
  network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id
  name = var.network_firewall_policy_application_group_name
}