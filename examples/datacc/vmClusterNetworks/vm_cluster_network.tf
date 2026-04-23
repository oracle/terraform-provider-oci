// Copyright (c) 2017, 2026, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "vm_cluster_network_consumer_type" {
  default = "INSTANCE"
}

variable "vm_cluster_network_defined_tags_value" {
  default = "value"
}

variable "vm_cluster_network_display_name" {
  default = "displayName"
}

variable "vm_cluster_network_dns_servers" {
  default = []
}

variable "vm_cluster_network_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "vm_cluster_network_is_scan_enabled" {
  default = false
}

variable "vm_cluster_network_listener_port" {
  default = 10
}

variable "vm_cluster_network_listener_port_ssl" {
  default = 10
}

variable "vm_cluster_network_node_count" {
  default = 10
}

variable "vm_cluster_network_ntp_servers" {
  default = []
}

variable "vm_cluster_network_scans_hostname" {
  default = "hostname"
}

variable "vm_cluster_network_scans_ips" {
  default = []
}

variable "vm_cluster_network_state" {
  default = []
}

variable "vm_cluster_network_vm_network_consumer_type" {
  default = "INSTANCE"
}

variable "vm_cluster_network_vm_networks_domain_name" {
  default = "domainName"
}

variable "vm_cluster_network_vm_networks_gateway" {
  default = "gateway"
}

variable "vm_cluster_network_vm_networks_netmask" {
  default = "netmask"
}

variable "vm_cluster_network_vm_networks_network_type" {
  default = "CLIENT"
}

variable "vm_cluster_network_vm_networks_nodes_hostname" {
  default = "hostname"
}

variable "vm_cluster_network_vm_networks_nodes_ip" {
  default = "ip"
}

variable "vm_cluster_network_vm_networks_nodes_vip" {
  default = "vip"
}

variable "vm_cluster_network_vm_networks_nodes_vip_hostname" {
  default = "vipHostname"
}

variable "vm_cluster_network_vm_networks_prefix" {
  default = "prefix"
}

variable "vlan_id" {
  default = "vlanId"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_datacc_vm_cluster_network" "test_vm_cluster_network" {
  #Required
  compartment_id    = var.compartment_id
  display_name      = var.vm_cluster_network_display_name
  infrastructure_id = oci_datacc_infrastructure.test_infrastructure.id
  vm_networks {
    #Required
    domain_name  = var.vm_cluster_network_vm_networks_domain_name
    gateway      = var.vm_cluster_network_vm_networks_gateway
    netmask      = var.vm_cluster_network_vm_networks_netmask
    network_type = var.vm_cluster_network_vm_networks_network_type
    nodes {
      #Required
      hostname = var.vm_cluster_network_vm_networks_nodes_hostname
      ip       = var.vm_cluster_network_vm_networks_nodes_ip

      #Optional
      vip          = var.vm_cluster_network_vm_networks_nodes_vip
      vip_hostname = var.vm_cluster_network_vm_networks_nodes_vip_hostname
    }

    #Optional
    prefix  = var.vm_cluster_network_vm_networks_prefix
    vlan_id = var.vlan_id
  }

  #Optional
  consumer_type     = var.vm_cluster_network_consumer_type
  defined_tags      = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.vm_cluster_network_defined_tags_value)
  dns_servers       = var.vm_cluster_network_dns_servers
  freeform_tags     = var.vm_cluster_network_freeform_tags
  listener_port     = var.vm_cluster_network_listener_port
  listener_port_ssl = var.vm_cluster_network_listener_port_ssl
  node_count        = var.vm_cluster_network_node_count
  ntp_servers       = var.vm_cluster_network_ntp_servers
  scans {
    #Required
    hostname = var.vm_cluster_network_scans_hostname
    ips      = var.vm_cluster_network_scans_ips
  }
}

data "oci_datacc_vm_cluster_networks" "test_vm_cluster_networks" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name             = var.vm_cluster_network_display_name
  infrastructure_id        = oci_datacc_infrastructure.test_infrastructure.id
  is_scan_enabled          = var.vm_cluster_network_is_scan_enabled
  node_count               = var.vm_cluster_network_node_count
  state                    = var.vm_cluster_network_state
  vm_network_consumer_type = var.vm_cluster_network_vm_network_consumer_type
}

