// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

/*
  This example shows how to add a decryption rule associated to a firewall policy.
*/

variable "network_firewall_policy_decryption_rule_action_no_decrypt" {
  default = "NO_DECRYPT"
}

variable "network_firewall_policy_decryption_rule_action_decrypt" {
  default = "DECRYPT"
}

variable "network_firewall_policy_decryption_rule_decryption_profile" {
  default = "decryption_profile_1"
}

variable "network_firewall_policy_decryption_rule_decryption_rule_priority_order" {
  default = 1
}

variable "network_firewall_policy_decryption_rule_name_decrypt" {
  default = "decryption_rule_decrypt"
}

variable "network_firewall_policy_decryption_rule_name_no_decrypt" {
  default = "decryption_rule_no_decrypt"
}

variable "network_firewall_policy_decryption_rule_position_after_rule" {
  default = null
}

variable "network_firewall_policy_decryption_rule_position_before_rule" {
  default = null
}

variable "network_firewall_policy_decryption_rule_secret" {
  default = "mapped_secret_1"
}

resource "oci_network_firewall_network_firewall_policy_decryption_rule" "test_network_firewall_policy_decryption_rule_no_decryption" {
  #Required
  name                       = var.network_firewall_policy_decryption_rule_name_no_decrypt
  network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id
  action = var.network_firewall_policy_decryption_rule_action_no_decrypt
  condition {

    #Optional
    destination_address = [oci_network_firewall_network_firewall_policy_address_list.test_network_firewall_policy_address_list_ip.name]
    source_address      = [oci_network_firewall_network_firewall_policy_address_list.test_network_firewall_policy_address_list_fqdn.name]
  }

  #Optional
  position {

    # Optional - Either 'after_rule' or 'before_rule' parameter should be set. This should be set only if there is an existing rule before this.
    # after_rule  = var.network_firewall_policy_decryption_rule_position_after_rule
    # before_rule = var.network_firewall_policy_decryption_rule_position_before_rule
  }
}

resource "oci_network_firewall_network_firewall_policy_decryption_rule" "test_network_firewall_policy_decryption_rule_with_decryption" {
  #Required
  name                       = var.network_firewall_policy_decryption_rule_name_decrypt
  network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id
  action = var.network_firewall_policy_decryption_rule_action_decrypt
  condition {

    #Optional
    destination_address = [oci_network_firewall_network_firewall_policy_address_list.test_network_firewall_policy_address_list_ip.name]
    source_address      = [oci_network_firewall_network_firewall_policy_address_list.test_network_firewall_policy_address_list_fqdn.name]
  }
  #Required when action is chosen as "DECRYPT"
  decryption_profile = oci_network_firewall_network_firewall_policy_decryption_profile.test_network_firewall_policy_decryption_profile_inbound_inspection.name
  secret             = oci_network_firewall_network_firewall_policy_mapped_secret.test_network_firewall_policy_mapped_secret_in.name

  #Optional
  position {

    #Optional - Either 'after_rule' or 'before_rule' parameter should be set. These should be set only if there is an existing rule before this.
    after_rule = oci_network_firewall_network_firewall_policy_decryption_rule.test_network_firewall_policy_decryption_rule_no_decryption.name
    #    before_rule = var.network_firewall_policy_decryption_rule_position_before_rule
  }
}

data "oci_network_firewall_network_firewall_policy_decryption_rules" "test_network_firewall_policy_decryption_rules" {
  #Required
  network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id

  #Optional - Either 'display_name' or 'decryption_rule_priority_order' parameter should be set.
  display_name = var.network_firewall_policy_decryption_rule_name_decrypt
  #  decryption_rule_priority_order = var.network_firewall_policy_decryption_rule_decryption_rule_priority_order
}

data "oci_network_firewall_network_firewall_policy_decryption_rule" "test_network_firewall_policy_decryption_rule" {
  #Required
  network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id

  #Optional - Either 'display_name' or 'decryption_rule_priority_order' parameter should be set.
  name = var.network_firewall_policy_decryption_rule_name_no_decrypt
  #  decryption_rule_priority_order = var.network_firewall_policy_decryption_rule_decryption_rule_priority_order
}