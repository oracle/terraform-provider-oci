// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

/*
  This example shows how to add a url list associated to a firewall policy.
*/

variable "network_firewall_policy_url_list_name" {
  default = "url_list_1"
}

variable "network_firewall_policy_url_list_urls_pattern" {
  default = ["oracle.com", "example.com"]
}

variable "network_firewall_policy_url_list_urls_type" {
  # only allowed value for type is 'SIMPLE'
  default = "SIMPLE"
}

resource "oci_network_firewall_network_firewall_policy_url_list" "test_network_firewall_policy_url_list" {
  #Required
  name                       = var.network_firewall_policy_url_list_name
  network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id
  dynamic "urls" {
    for_each = toset(var.network_firewall_policy_url_list_urls_pattern)
    content {
      #Required
      pattern = urls.value
      type    = var.network_firewall_policy_url_list_urls_type
    }
  }
}

data "oci_network_firewall_network_firewall_policy_url_lists" "test_network_firewall_policy_url_lists" {
  #Required
  network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id

  #Optional
  display_name = var.network_firewall_policy_url_list_name
}

data "oci_network_firewall_network_firewall_policy_url_list" "test_network_firewall_policy_url_list" {
  #Required
  network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id
  name = var.network_firewall_policy_url_list_name
}