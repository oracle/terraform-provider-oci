// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

/*
  This example shows how to add a security rule associated to a firewall policy.
*/

variable "network_firewall_policy_security_rule_action_allow" {
  # allowed values for 'action' are 'ALLOW', 'REJECT', 'DROP', 'INSPECT'
  default = "ALLOW"
}

variable "network_firewall_policy_security_rule_action_inspect" {
  # allowed values for 'action' are 'ALLOW', 'REJECT', 'DROP', 'INSPECT'
  default = "INSPECT"
}

variable "network_firewall_policy_security_rule_inspection" {
  # allowed values for 'inspection' are 'INTRUSION_DETECTION', 'INTRUSION_PREVENTION'
  default = "INTRUSION_DETECTION"
}

variable "network_firewall_policy_security_rule_allow_name" {
  default = "security_rule_allow"
}

variable "network_firewall_policy_security_rule_inspect_name" {
  default = "security_rule_inspect"
}

variable "network_firewall_policy_security_rule_empty_condition_name" {
  default = "security_rule_empty_condition"
}

variable "network_firewall_policy_security_rule_position_after_rule" {
  default = null
}

variable "network_firewall_policy_security_rule_position_before_rule" {
  default = null
}

variable "network_firewall_policy_security_rule_security_rule_priority_order" {
  default = 1
}

resource "oci_network_firewall_network_firewall_policy_security_rule" "test_network_firewall_policy_security_rule_allow" {
  #Required
  name                       = var.network_firewall_policy_security_rule_allow_name
  network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id
  action = var.network_firewall_policy_security_rule_action_allow
  condition {

    #Optional
    application         = [oci_network_firewall_network_firewall_policy_application_group.test_network_firewall_policy_application_group.name]
    destination_address = [oci_network_firewall_network_firewall_policy_address_list.test_network_firewall_policy_address_list_ip.name]
    service             = [oci_network_firewall_network_firewall_policy_service_list.test_network_firewall_policy_service_list.name]
    source_address      = [oci_network_firewall_network_firewall_policy_address_list.test_network_firewall_policy_address_list_ip.name]
    url                 = [oci_network_firewall_network_firewall_policy_url_list.test_network_firewall_policy_url_list.name]
  }

  position {
    # Optional - #Optional - Either 'after_rule' or 'before_rule' parameter should be set. These should be set only if there is an existing rule before this.
    # after_rule  = var.network_firewall_policy_security_rule_position_after_rule
    # before_rule = var.network_firewall_policy_security_rule_position_before_rule
  }
}

resource "oci_network_firewall_network_firewall_policy_security_rule" "test_network_firewall_policy_security_rule_inspect" {
  #Required
  name                       = var.network_firewall_policy_security_rule_inspect_name
  network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id
  action = var.network_firewall_policy_security_rule_action_inspect
  condition {

    #Optional
    application         = []
    destination_address = []
    service             = []
    source_address      = []
    url                 = []
  }

  #Optional - provide 'inspection' value only when action is selected as 'INSPECT'
  inspection = var.network_firewall_policy_security_rule_inspection
  position {
    after_rule  = oci_network_firewall_network_firewall_policy_security_rule.test_network_firewall_policy_security_rule_allow.name
  }
}

resource "oci_network_firewall_network_firewall_policy_security_rule" "test_network_firewall_policy_security_rule_empty_condition" {
  #Required
  name                       = var.network_firewall_policy_security_rule_empty_condition_name
  network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id
  action = var.network_firewall_policy_security_rule_action_inspect
  condition {}

  #Optional - provide 'inspection' value only when action is selected as 'INSPECT'
  inspection = var.network_firewall_policy_security_rule_inspection
  position {
    after_rule  = oci_network_firewall_network_firewall_policy_security_rule.test_network_firewall_policy_security_rule_allow.name
  }
}

data "oci_network_firewall_network_firewall_policy_security_rules" "test_network_firewall_policy_security_rules" {
  #Required
  network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id

  #Optional - Either 'display_name' or 'security_rule_priority_order' parameter should be set.
  display_name = var.network_firewall_policy_security_rule_allow_name
  # security_rule_priority_order = var.network_firewall_policy_security_rule_security_rule_priority_order
}

data "oci_network_firewall_network_firewall_policy_security_rule" "test_network_firewall_policy_security_rule" {
  #Required
  network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id
  name = var.network_firewall_policy_security_rule_inspect_name
}