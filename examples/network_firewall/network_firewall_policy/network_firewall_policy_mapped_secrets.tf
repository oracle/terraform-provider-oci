// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

/*
  This example shows how to create a mapped secret associated to a firewall policy.
*/


variable "vault_secret_id" {
}

variable "network_firewall_policy_mapped_secret_name_fp" {
  default = "mapped_secret_fp"
}

variable "network_firewall_policy_mapped_secret_name_in" {
  default = "mapped_secret_in"
}

variable "network_firewall_policy_mapped_secret_source" {
  # only allowed value for source is 'OCI_VAULT'
  default = "OCI_VAULT"
}

variable "network_firewall_policy_mapped_secret_type_in" {
  default = "SSL_INBOUND_INSPECTION"
}

variable "network_firewall_policy_mapped_secret_type_fp" {
  default = "SSL_FORWARD_PROXY"
}


variable "network_firewall_policy_mapped_secret_version_number" {
  default = 1
}

resource "oci_network_firewall_network_firewall_policy_mapped_secret" "test_network_firewall_policy_mapped_secret_fp" {
  #Required
  name                       = var.network_firewall_policy_mapped_secret_name_fp
  network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id
  source                     = var.network_firewall_policy_mapped_secret_source
  type                       = var.network_firewall_policy_mapped_secret_type_fp
  vault_secret_id            = var.vault_secret_id
  version_number             = var.network_firewall_policy_mapped_secret_version_number
}

resource "oci_network_firewall_network_firewall_policy_mapped_secret" "test_network_firewall_policy_mapped_secret_in" {
  #Required
  name                       = var.network_firewall_policy_mapped_secret_name_in
  network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id
  source                     = var.network_firewall_policy_mapped_secret_source
  type                       = var.network_firewall_policy_mapped_secret_type_in
  vault_secret_id            = var.vault_secret_id
  version_number             = var.network_firewall_policy_mapped_secret_version_number
}

data "oci_network_firewall_network_firewall_policy_mapped_secrets" "test_network_firewall_policy_mapped_secrets" {
  #Required
  network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id

  #Optional
  display_name = var.network_firewall_policy_mapped_secret_name_fp
}

data "oci_network_firewall_network_firewall_policy_mapped_secret" "test_network_firewall_policy_mapped_secret" {
  #Required
  network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id
  name = var.network_firewall_policy_mapped_secret_name_fp
}