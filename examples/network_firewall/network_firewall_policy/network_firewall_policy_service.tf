// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

/*
  This example shows how to create a service associated to a firewall policy.
*/

variable "network_firewall_policy_service_name_udp" {
  default = "service_udp"
}

variable "network_firewall_policy_service_name_tcp" {
  default = "service_tcp"
}

variable "network_firewall_policy_service_port_ranges" {
  default = ["10-100", "1"]
}

variable "network_firewall_policy_service_type_tcp" {
  # allowed values for 'type' are 'TCP_SERVICE', 'UDP_SERVICE'
  default = "TCP_SERVICE"
}

variable "network_firewall_policy_service_type_udp" {
  # allowed values for 'type' are 'TCP_SERVICE', 'UDP_SERVICE'
  default = "UDP_SERVICE"
}

resource "oci_network_firewall_network_firewall_policy_service" "test_network_firewall_policy_service_tcp" {
  #Required
  name                       = var.network_firewall_policy_service_name_tcp
  network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id
  dynamic "port_ranges" {
    for_each = toset(var.network_firewall_policy_service_port_ranges)
    content {
      minimum_port = split("-", port_ranges.value)[0]
      #Optional
      maximum_port = length(split("-", port_ranges.value)) == 2 ? split("-", port_ranges.value)[1] : null
    }
  }
  type = var.network_firewall_policy_service_type_tcp
}

resource "oci_network_firewall_network_firewall_policy_service" "test_network_firewall_policy_service_udp" {
  #Required
  name                       = var.network_firewall_policy_service_name_udp
  network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id
  dynamic "port_ranges" {
    for_each = toset(var.network_firewall_policy_service_port_ranges)
    content {
      minimum_port = split("-", port_ranges.value)[0]
      #Optional
      maximum_port = length(split("-", port_ranges.value)) == 2 ? split("-", port_ranges.value)[1] : null
    }
  }
  type = var.network_firewall_policy_service_type_udp
}

data "oci_network_firewall_network_firewall_policy_services" "test_network_firewall_policy_services" {
  #Required
  network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id

  #Optional
  display_name = var.network_firewall_policy_service_name_tcp
}

data "oci_network_firewall_network_firewall_policy_service" "test_network_firewall_policy_service" {
  #Required
  network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id
  name = var.network_firewall_policy_service_name_tcp
}