// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "smtp_credential_smtp_credential_count" {
  default = 10
}

variable "smtp_credential_authorization" {
  default = "authorization"
}

variable "smtp_credential_description" {
  default = "description"
}

variable "smtp_credential_expires_on" {
  default = "2030-01-01T00:00:00Z"
}

variable "smtp_credential_idcs_created_by_display" {
  default = "display"
}

variable "smtp_credential_start_index" {
  default = 1
}

variable "smtp_credential_status" {
  default = "ACTIVE"
}

variable "smtp_credential_tags_key" {
  default = "key"
}

variable "smtp_credential_tags_value" {
  default = "value"
}

variable "smtp_credential_urnietfparamsscimschemasoracleidcsextensionself_change_user_allow_self_change" {
  default = false
}

resource "oci_identity_domains_smtp_credential" "test_smtp_credential" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain.url
  schemas       = ["urn:ietf:params:scim:schemas:oracle:idcs:smtpCredential"]

  #Optional
  attribute_sets  = ["all"]
  attributes     = ""
  authorization  = var.smtp_credential_authorization
  description    = var.smtp_credential_description
  expires_on     = var.smtp_credential_expires_on
  #use the latest if not provided
  # resource_type_schema_version = var.smtp_credential_resource_type_schema_version
  status = var.smtp_credential_status
  tags {
    #Required
    key   = var.smtp_credential_tags_key
    value = var.smtp_credential_tags_value
  }
  urnietfparamsscimschemasoracleidcsextensionself_change_user {

    #Optional
    allow_self_change = var.smtp_credential_urnietfparamsscimschemasoracleidcsextensionself_change_user_allow_self_change
  }
  user {

    #Optional
    #use the ocid of the same user set in value
    ocid = oci_identity_domains_user.test_user.ocid
    #must be a user that exists
    value = oci_identity_domains_user.test_user.id
  }
  lifecycle {
    ignore_changes = [
      // ignore fields that will never be returned
      status,
      urnietfparamsscimschemasoracleidcsextensionself_change_user
    ]
  }
}

data "oci_identity_domains_smtp_credentials" "test_smtp_credentials" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain.url

  #Optional
  smtp_credential_count  = var.smtp_credential_smtp_credential_count
  smtp_credential_filter = "user.value eq \"${oci_identity_domains_user.test_user.id}\""
  attribute_sets         = []
  attributes             = ""
  authorization          = var.smtp_credential_authorization
  #use the latest if not provided
  # resource_type_schema_version = var.smtp_credential_resource_type_schema_version
  start_index = var.smtp_credential_start_index
}

