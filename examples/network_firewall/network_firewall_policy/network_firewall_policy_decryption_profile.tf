// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

/*
  This example shows how to create a decryption profile associated to a firewall policy.
*/

variable "network_firewall_policy_decryption_profile_fp_name" {
  default = "decryption_profile_fp"
}

variable "network_firewall_policy_decryption_profile_in_name" {
  default = "decryption_profile_in"
}

variable "network_firewall_policy_decryption_profile_type_forward_proxy" {
  default = "SSL_FORWARD_PROXY"
}

variable "network_firewall_policy_decryption_profile_type_inbound_inspection" {
  default = "SSL_INBOUND_INSPECTION"
}

variable "network_firewall_policy_decryption_profile_are_certificate_extensions_restricted" {
  default = false
}

variable "network_firewall_policy_decryption_profile_is_auto_include_alt_name" {
  default = false
}

variable "network_firewall_policy_decryption_profile_is_expired_certificate_blocked" {
  default = false
}

variable "network_firewall_policy_decryption_profile_is_out_of_capacity_blocked" {
  default = false
}

variable "network_firewall_policy_decryption_profile_is_revocation_status_timeout_blocked" {
  default = false
}

variable "network_firewall_policy_decryption_profile_is_unknown_revocation_status_blocked" {
  default = false
}

variable "network_firewall_policy_decryption_profile_is_unsupported_cipher_blocked" {
  default = false
}

variable "network_firewall_policy_decryption_profile_is_unsupported_version_blocked" {
  default = false
}

variable "network_firewall_policy_decryption_profile_is_untrusted_issuer_blocked" {
  default = false
}

resource "oci_network_firewall_network_firewall_policy_decryption_profile" "test_network_firewall_policy_decryption_profile_forward_proxy" {
  #Required
  name                       = var.network_firewall_policy_decryption_profile_fp_name
  network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id
  type                       = var.network_firewall_policy_decryption_profile_type_forward_proxy

  #Optional
  are_certificate_extensions_restricted = var.network_firewall_policy_decryption_profile_are_certificate_extensions_restricted
  is_auto_include_alt_name              = var.network_firewall_policy_decryption_profile_is_auto_include_alt_name
  is_expired_certificate_blocked        = var.network_firewall_policy_decryption_profile_is_expired_certificate_blocked
  is_out_of_capacity_blocked            = var.network_firewall_policy_decryption_profile_is_out_of_capacity_blocked
  is_revocation_status_timeout_blocked  = var.network_firewall_policy_decryption_profile_is_revocation_status_timeout_blocked
  is_unknown_revocation_status_blocked  = var.network_firewall_policy_decryption_profile_is_unknown_revocation_status_blocked
  is_unsupported_cipher_blocked         = var.network_firewall_policy_decryption_profile_is_unsupported_cipher_blocked
  is_unsupported_version_blocked        = var.network_firewall_policy_decryption_profile_is_unsupported_version_blocked
  is_untrusted_issuer_blocked           = var.network_firewall_policy_decryption_profile_is_untrusted_issuer_blocked
}

resource "oci_network_firewall_network_firewall_policy_decryption_profile" "test_network_firewall_policy_decryption_profile_inbound_inspection" {
  #Required
  name                       = var.network_firewall_policy_decryption_profile_in_name
  network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id
  type                       = var.network_firewall_policy_decryption_profile_type_inbound_inspection

  #Optional - for type 'SSL_INBOUND_INSPECTION', only three optional parameters can be set
  is_out_of_capacity_blocked     = var.network_firewall_policy_decryption_profile_is_out_of_capacity_blocked
  is_unsupported_cipher_blocked  = var.network_firewall_policy_decryption_profile_is_unsupported_cipher_blocked
  is_unsupported_version_blocked = var.network_firewall_policy_decryption_profile_is_unsupported_version_blocked
}

data "oci_network_firewall_network_firewall_policy_decryption_profiles" "test_network_firewall_policy_decryption_profiles" {
  #Required
  network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id

  #Optional
  display_name = var.network_firewall_policy_decryption_profile_fp_name
}

data "oci_network_firewall_network_firewall_policy_decryption_profile" "test_network_firewall_policy_decryption_profile" {
  #Required
  network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id
  name = var.network_firewall_policy_decryption_profile_in_name
}