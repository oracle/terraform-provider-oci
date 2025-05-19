// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

/*
  This example shows how to add a nat rule associated to a firewall policy.
*/

variable "network_firewall_policy_nat_rule_nat_type" {
  default = "NATV4"
}

variable "network_firewall_policy_nat_rule_action" {
  default = "DIPP_SRC_NAT"
}
variable "network_firewall_policy_nat_rule_priority_order" {
  default = 1
}

variable "network_firewall_policy_nat_rule_name_1" {
  default = "nat_rule_1"
}

variable "network_firewall_policy_nat_rule_name_2" {
  default = "nat_rule_2"
}
variable "network_firewall_policy_nat_rule_empty_condition_name" {
  default = "nat_rule_empty_condition"
}

resource "oci_network_firewall_network_firewall_policy_nat_rule" "test_network_firewall_policy_nat_rule" {
  #Required
  name                       = var.network_firewall_policy_nat_rule_name_1
  network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id
  action = var.network_firewall_policy_nat_rule_action
  type = var.network_firewall_policy_nat_rule_nat_type
  condition {

    #Optional
    service = oci_network_firewall_network_firewall_policy_service.test_network_firewall_policy_service_tcp.name
    source_address = [oci_network_firewall_network_firewall_policy_address_list.test_network_firewall_policy_address_list_ip.name]
    destination_address = [oci_network_firewall_network_firewall_policy_address_list.test_network_firewall_policy_address_list_ip.name]
  }

  #Optional
  position {

    # Optional - Either 'after_rule' or 'before_rule' parameter should be set. This should be set only if there is an existing rule before this.
    # after_rule  = var.network_firewall_policy_nat_rule_name_1
    # before_rule = var.network_firewall_policy_nat_rule_name_2
  }
}

resource "oci_network_firewall_network_firewall_policy_nat_rule" "test_network_firewall_policy_nat_rule_2" {
  #Required
  name                       = var.network_firewall_policy_nat_rule_name_2
  network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id
  action = var.network_firewall_policy_nat_rule_action
  type = var.network_firewall_policy_nat_rule_nat_type
  condition {

    #Optional
    source_address      = []
    destination_address = []
  }

  #Optional
  position {

    #Optional - Either 'after_rule' or 'before_rule' parameter should be set. These should be set only if there is an existing rule before this.
    after_rule = oci_network_firewall_network_firewall_policy_nat_rule.test_network_firewall_policy_nat_rule.name
    #before_rule = var.network_firewall_policy_nat_rule_name_1
  }
}

resource "oci_network_firewall_network_firewall_policy_nat_rule" "test_network_firewall_policy_nat_rule_empty_condition" {
  #Required
  name                       = var.network_firewall_policy_nat_rule_empty_condition_name
  network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id
  action = var.network_firewall_policy_nat_rule_action
  type = var.network_firewall_policy_nat_rule_nat_type
  condition {}

  position {
    after_rule = oci_network_firewall_network_firewall_policy_nat_rule.test_network_firewall_policy_nat_rule.name
  }
}

data "oci_network_firewall_network_firewall_policy_nat_rules" "test_network_firewall_policy_nat_rules" {
  #Required
  network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id

  display_name = var.network_firewall_policy_nat_rule_name_1
  #  nat_rule_priority_order = var.network_firewall_policy_nat_rule_priority_order

}

data "oci_network_firewall_network_firewall_policy_nat_rule" "test_network_firewall_policy_nat_rule" {
  #Required
  network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id
  nat_rule_name = var.network_firewall_policy_nat_rule_name_1
}