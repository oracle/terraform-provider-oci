// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

/*
  This example shows how to add a service list associated to a firewall policy.
*/

variable "network_firewall_policy_service_list_name" {
  default = "service_list_1"
}

resource "oci_network_firewall_network_firewall_policy_service_list" "test_network_firewall_policy_service_list" {
  #Required
  name                       = var.network_firewall_policy_service_list_name
  network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id
  services                   = [oci_network_firewall_network_firewall_policy_service.test_network_firewall_policy_service_tcp.name]
}

data "oci_network_firewall_network_firewall_policy_service_lists" "test_network_firewall_policy_service_lists" {
  #Required
  network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id

  #Optional
  display_name = var.network_firewall_policy_service_list_name
}

data "oci_network_firewall_network_firewall_policy_service_list" "test_network_firewall_policy_service_list" {
  #Required
  network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id
  name = var.network_firewall_policy_service_list_name
}