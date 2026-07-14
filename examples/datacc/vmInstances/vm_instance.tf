// Copyright (c) 2017, 2026, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "vm_instance_boot_storage_size_in_gbs" {
  default = 1.0
}

variable "base_server_id" {
  default = "serverId"
}

variable "vm_instance_cpus_enabled" {
  default = 10
}

variable "vm_instance_data_storage_size_in_gb" {
  default = 1.0
}

variable "vm_instance_defined_tags_value" {
  default = "value"
}

variable "vm_instance_description" {
  default = "description"
}

variable "vm_instance_display_name" {
  default = "displayName"
}

variable "vm_instance_dns_servers" {
  default = []
}

variable "domain_name" {
  default = "domainName"
}

variable "vm_instance_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "vm_instance_gateway" {
  default = "gateway"
}

variable "vm_instance_hostname" {
  default = "hostname"
}

variable "image_id" {
  default = "imageId"
}

variable "vm_instance_ip_address" {
  default = "ipAddress"
}

variable "vm_instance_memory_size_in_gbs" {
  default = 1.0
}

variable "vm_instance_metadata" {
  default = "metadata"
}

variable "vm_instance_netmask" {
  default = "netmask"
}

variable "vm_instance_ntp_servers" {
  default = []
}

variable "server_id" {
  default = "serverId"
}

variable "vm_instance_ssh_public_keys" {
  default = []
}

variable "vm_instance_state" {
  default = []
}

variable "vm_instance_system_tags" {
  default = "value"
}

variable "vm_instance_time_zone" {
  default = "timeZone"
}

variable "vm_instance_userdata" {
  default = "userdata"
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

resource "oci_datacc_vm_instance" "test_vm_instance" {
  #Required
  compartment_id    = var.compartment_id
  cpus_enabled      = var.vm_instance_cpus_enabled
  infrastructure_id = oci_datacc_infrastructure.test_infrastructure.id
  ssh_public_keys   = var.vm_instance_ssh_public_keys

  #Optional
  boot_storage_size_in_gbs = var.vm_instance_boot_storage_size_in_gbs
  data_storage_size_in_gb  = var.vm_instance_data_storage_size_in_gb
  defined_tags             = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.vm_instance_defined_tags_value)
  description              = var.vm_instance_description
  display_name             = var.vm_instance_display_name
  dns_servers              = var.vm_instance_dns_servers
  domain_name              = var.domain_name
  freeform_tags            = var.vm_instance_freeform_tags
  gateway                  = var.vm_instance_gateway
  hostname                 = var.vm_instance_hostname
  image_id                 = var.image_id
  ip_address               = var.vm_instance_ip_address
  memory_size_in_gbs       = var.vm_instance_memory_size_in_gbs
  metadata                 = var.vm_instance_metadata
  netmask                  = var.vm_instance_netmask
  ntp_servers              = var.vm_instance_ntp_servers
  server_id                = var.server_id
  system_tags              = var.vm_instance_system_tags
  time_zone                = var.vm_instance_time_zone
  userdata                 = var.vm_instance_userdata
  vlan_id                  = var.vlan_id
  vm_network_id            = oci_datacc_vm_network.test_vm_network.id
}

data "oci_datacc_vm_instances" "test_vm_instances" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  base_server_id    = var.base_server_id
  display_name      = var.vm_instance_display_name
  infrastructure_id = oci_datacc_infrastructure.test_infrastructure.id
  state             = var.vm_instance_state
}
