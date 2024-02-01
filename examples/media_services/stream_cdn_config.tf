// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0


variable "stream_cdn_config_config_edge_hostname" {
  default = "edgeHostname"
}

variable "stream_cdn_config_config_edge_path_prefix" {
  default = "edgePathPrefix"
}

variable "stream_cdn_config_config_edge_token_key" {
  default = "edgeTokenKey"
}

variable "stream_cdn_config_config_edge_token_salt" {
  default = "edgeTokenSalt"
}

variable "stream_cdn_config_config_is_edge_token_auth" {
  default = false
}

variable "stream_cdn_config_config_origin_auth_secret_key_a" {
  default = "originAuthSecretKeyA"
}

variable "stream_cdn_config_config_origin_auth_secret_key_b" {
  default = "originAuthSecretKeyB"
}

variable "stream_cdn_config_config_origin_auth_secret_key_nonce_a" {
  default = "originAuthSecretKeyNonceA"
}

variable "stream_cdn_config_config_origin_auth_secret_key_nonce_b" {
  default = "originAuthSecretKeyNonceB"
}

variable "stream_cdn_config_config_origin_auth_sign_encryption" {
  default = "SHA256-HMAC"
}

variable "stream_cdn_config_config_origin_auth_sign_type" {
  default = "ForwardURL"
}

variable "stream_cdn_config_config_type" {
  default = "EDGE"
}

variable "stream_cdn_config_is_enabled" {
  default = false
}

resource "oci_media_services_stream_cdn_config" "test_stream_cdn_config" {
  #Required
  config {
    #Required
    type = var.stream_cdn_config_config_type

    #Optional
    edge_hostname                  = var.stream_cdn_config_config_edge_hostname
    edge_path_prefix               = var.stream_cdn_config_config_edge_path_prefix
    edge_token_key                 = var.stream_cdn_config_config_edge_token_key
    edge_token_salt                = var.stream_cdn_config_config_edge_token_salt
    is_edge_token_auth             = var.stream_cdn_config_config_is_edge_token_auth
    origin_auth_secret_key_a       = var.stream_cdn_config_config_origin_auth_secret_key_a
    origin_auth_secret_key_b       = var.stream_cdn_config_config_origin_auth_secret_key_b
    origin_auth_secret_key_nonce_a = var.stream_cdn_config_config_origin_auth_secret_key_nonce_a
    origin_auth_secret_key_nonce_b = var.stream_cdn_config_config_origin_auth_secret_key_nonce_b
    origin_auth_sign_encryption    = var.stream_cdn_config_config_origin_auth_sign_encryption
    origin_auth_sign_type          = var.stream_cdn_config_config_origin_auth_sign_type
  }
  display_name            = var.display_name
  distribution_channel_id = oci_media_services_stream_distribution_channel.test_stream_distribution_channel.id

  #Optional
  defined_tags  = map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.defined_tags_value}")
  freeform_tags = var.freeform_tags
  is_enabled    = var.stream_cdn_config_is_enabled
  locks {
    #Required
    compartment_id = var.compartment_id
    type = var.locks_type

    #Optional
    message = var.locks_message
  }
  is_lock_override = var.is_lock_override
  lifecycle {
    ignore_changes = [defined_tags, locks, system_tags, lifecyle_details, config, is_lock_override]
  }
}

data "oci_media_services_stream_cdn_configs" "test_stream_cdn_configs" {
  #Required
  distribution_channel_id = oci_media_services_stream_distribution_channel.test_stream_distribution_channel.id

  #Optional
  display_name = var.display_name
  id           = var.id
  state        = var.active_state
}

