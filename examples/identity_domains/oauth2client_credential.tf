// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "oauth2client_credential_oauth2client_credential_count" {
  default = 10
}

variable "oauth2client_credential_authorization" {
  default = "authorization"
}

variable "oauth2client_credential_description" {
  default = "description"
}

variable "oauth2client_credential_expires_on" {
  default = "2030-01-01T00:00:00Z"
}

variable "oauth2client_credential_is_reset_secret" {
  default = false
}

variable "oauth2client_credential_name" {
  default = "name"
}

variable "oauth2client_credential_scopes_audience" {
  default = "audience"
}

variable "oauth2client_credential_scopes_scope" {
  default = "scope"
}

variable "oauth2client_credential_start_index" {
  default = 1
}

variable "oauth2client_credential_status" {
  default = "ACTIVE"
}

variable "oauth2client_credential_tags_key" {
  default = "key"
}

variable "oauth2client_credential_tags_value" {
  default = "value"
}

variable "oauth2client_credential_urnietfparamsscimschemasoracleidcsextensionself_change_user_allow_self_change" {
  default = false
}

resource "oci_identity_domains_oauth2client_credential" "test_oauth2client_credential" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain.url
  name          = var.oauth2client_credential_name
  schemas       = ["urn:ietf:params:scim:schemas:oracle:idcs:oauth2ClientCredential"]
  scopes {
    #Required
    audience = var.oauth2client_credential_scopes_audience
    scope    = var.oauth2client_credential_scopes_scope
  }

  #Optional
  attribute_sets  = ["all"]
  attributes      = ""
  authorization   = var.oauth2client_credential_authorization
  description     = var.oauth2client_credential_description
  expires_on      = var.oauth2client_credential_expires_on
  is_reset_secret = var.oauth2client_credential_is_reset_secret
  #use the latest if not provided
  # resource_type_schema_version = var.oauth2client_credential_resource_type_schema_version
  status = var.oauth2client_credential_status
  tags {
    #Required
    key   = var.oauth2client_credential_tags_key
    value = var.oauth2client_credential_tags_value
  }
  urnietfparamsscimschemasoracleidcsextensionself_change_user {

    #Optional
    allow_self_change = var.oauth2client_credential_urnietfparamsscimschemasoracleidcsextensionself_change_user_allow_self_change
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

data "oci_identity_domains_oauth2client_credentials" "test_oauth2client_credentials" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain.url

  #Optional
  oauth2client_credential_count  = var.oauth2client_credential_oauth2client_credential_count
  oauth2client_credential_filter = "user.value eq \"${oci_identity_domains_user.test_user.id}\""
  attribute_sets                 = []
  attributes                     = ""
  authorization                  = var.oauth2client_credential_authorization
  #use the latest if not provided
  # resource_type_schema_version   = var.oauth2client_credential_resource_type_schema_version
  start_index = var.oauth2client_credential_start_index
}

