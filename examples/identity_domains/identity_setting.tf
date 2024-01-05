// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "identity_setting_posix_gid_manual_assignment_ends_at" {
  default = "1000"
}

variable "identity_setting_posix_gid_manual_assignment_starts_from" {
  default = 10
}

variable "identity_setting_posix_uid_manual_assignment_ends_at" {
  default = "1000"
}

variable "identity_setting_posix_uid_manual_assignment_starts_from" {
  default = 10
}

variable "identity_setting_authorization" {
  default = "authorization"
}

variable "identity_setting_emit_locked_message_when_user_is_locked" {
  default = false
}

variable "identity_setting_id" {
  default = "IdentitySettings"
}

variable "identity_setting_primary_email_required" {
  default = false
}

variable "identity_setting_tags_key" {
  default = "key"
}

variable "identity_setting_tags_value" {
  default = "value"
}

variable "identity_setting_user_allowed_to_set_recovery_email" {
  default = false
}

resource "oci_identity_domains_identity_setting" "test_identity_setting" {
  #Required
  idcs_endpoint       = data.oci_identity_domain.test_domain.url
  identity_setting_id = var.identity_setting_id
  schemas             = ["urn:ietf:params:scim:schemas:oracle:idcs:IdentitySettings"]

  #Optional
  posix_gid {

    #Optional
    manual_assignment_ends_at     = var.identity_setting_posix_gid_manual_assignment_ends_at
    manual_assignment_starts_from = var.identity_setting_posix_gid_manual_assignment_starts_from
  }
  posix_uid {

    #Optional
    manual_assignment_ends_at     = var.identity_setting_posix_uid_manual_assignment_ends_at
    manual_assignment_starts_from = var.identity_setting_posix_uid_manual_assignment_starts_from
  }
  attribute_sets                          = ["all"]
  attributes                              = ""
  authorization                           = var.identity_setting_authorization
  emit_locked_message_when_user_is_locked = var.identity_setting_emit_locked_message_when_user_is_locked
  external_id                             = "externalId"
  primary_email_required                  = var.identity_setting_primary_email_required
  tags {
    #Required
    key   = var.identity_setting_tags_key
    value = var.identity_setting_tags_value
  }
  tokens {
    #Required
    type = "emailVerification"
    expires_after = "10"
  }
  tokens {
    #Required
    type = "passwordReset"
    expires_after = "120"
  }
  tokens {
    #Required
    type = "createUser"
    expires_after = "2500"
  }
  my_profile {

    #Optional
    allow_end_users_to_change_their_password          = true
    allow_end_users_to_link_their_support_account     = true
    allow_end_users_to_manage_their_capabilities      = true
    allow_end_users_to_update_their_security_settings = true
  }
  user_allowed_to_set_recovery_email = true
  lifecycle {
    ignore_changes = [my_profile]
  }
}

data "oci_identity_domains_identity_settings" "test_identity_settings" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain.url

  #Optional
  attribute_sets               = ["all"]
  attributes                   = ""
  authorization                = var.identity_setting_authorization
}

