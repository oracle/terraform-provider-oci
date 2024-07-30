// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

/*
  This example shows how to add a tunnel inspection rule associated to a firewall policy.
*/

variable "network_firewall_policy_tunnel_inspection_rule_action_inspect" {
  default = "INSPECT"
}

variable "network_firewall_policy_tunnel_inspection_rule_action_inspect_and_capture_log" {
  default = "INSPECT_AND_CAPTURE_LOG"
}

variable "network_firewall_policy_tunnel_inspection_rule_protocol" {
  default = "VXLAN"
}

variable "network_firewall_policy_tunnel_inspection_rule_priority_order" {
  default = 1
}

variable "network_firewall_policy_tunnel_inspection_rule_name_1" {
  default = "tunnel_rule_1"
}

variable "network_firewall_policy_tunnel_inspection_rule_name_2" {
  default = "tunnel_rule_2"
}

resource "oci_network_firewall_network_firewall_policy_tunnel_inspection_rule" "test_network_firewall_policy_tunnel_inspection_rule_inspect" {
  #Required
  name                       = var.network_firewall_policy_tunnel_inspection_rule_name_1
  network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id
  action = var.network_firewall_policy_tunnel_inspection_rule_action_inspect
  protocol = "VXLAN"
  condition {

    #Optional
    destination_address = []
    source_address      = []
  }

  #Optional
  position {

    # Optional - Either 'after_rule' or 'before_rule' parameter should be set. This should be set only if there is an existing rule before this.
    # after_rule  = var.network_firewall_policy_tunnel_inspection_rule_name_1
    # before_rule = var.network_firewall_policy_tunnel_inspection_rule_name_2
  }
}

resource "oci_network_firewall_network_firewall_policy_tunnel_inspection_rule" "test_network_firewall_policy_tunnel_inspection_rule_2" {
  #Required
  name                       = var.network_firewall_policy_tunnel_inspection_rule_name_2
  network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id
  action = var.network_firewall_policy_tunnel_inspection_rule_action_inspect_and_capture_log
  condition {

    #Optional
    destination_address = [oci_network_firewall_network_firewall_policy_address_list.test_network_firewall_policy_address_list_ip.name]
    source_address      = []
  }
  protocol = "VXLAN"

  #Optional
  position {

    #Optional - Either 'after_rule' or 'before_rule' parameter should be set. These should be set only if there is an existing rule before this.
    after_rule = oci_network_firewall_network_firewall_policy_tunnel_inspection_rule.test_network_firewall_policy_tunnel_inspection_rule_inspect.name
    #before_rule = var.network_firewall_policy_tunnel_inspection_rule_name_1
  }
}

data "oci_network_firewall_network_firewall_policy_tunnel_inspection_rules" "test_network_firewall_policy_tunnel_inspection_rules" {
  #Required
  network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id

  #Optional - Either 'display_name' or 'tunnel_inspection_rule_priority_order' parameter should be set.
  display_name = var.network_firewall_policy_tunnel_inspection_rule_name_1
  #  tunnel_inspection_rule_priority_order = var.network_firewall_policy_tunnel_inspection_rule_priority_order
}

data "oci_network_firewall_network_firewall_policy_tunnel_inspection_rule" "test_network_firewall_policy_tunnel_inspection_rule" {
  #Required
  network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id
  tunnel_inspection_rule_name = var.network_firewall_policy_tunnel_inspection_rule_name_1
}