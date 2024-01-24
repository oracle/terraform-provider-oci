// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "auth_token_auth_token_count" {
  default = 10
}

variable "auth_token_authorization" {
  default = "authorization"
}

variable "auth_token_description" {
  default = "description"
}

variable "auth_token_expires_on" {
  default = "2030-01-01T00:00:00Z"
}

variable "auth_token_start_index" {
  default = 1
}

variable "auth_token_status" {
  default = "ACTIVE"
}

variable "auth_token_tags_key" {
  default = "key"
}

variable "auth_token_tags_value" {
  default = "value"
}

variable "auth_token_urnietfparamsscimschemasoracleidcsextensionself_change_user_allow_self_change" {
  default = false
}


resource "oci_identity_domains_auth_token" "test_auth_token" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain.url
  schemas       = ["urn:ietf:params:scim:schemas:oracle:idcs:authToken"]

  #Optional
  attribute_sets = ["all"]
  attributes     = ""
  authorization  = var.auth_token_authorization
  description    = var.auth_token_description
  expires_on     = var.auth_token_expires_on
  #use the latest if not provided
  # resource_type_schema_version = var.auth_token_resource_type_schema_version
  status = var.auth_token_status
  tags {
    #Required
    key   = var.auth_token_tags_key
    value = var.auth_token_tags_value
  }
  urnietfparamsscimschemasoracleidcsextensionself_change_user {

    #Optional
    allow_self_change = var.auth_token_urnietfparamsscimschemasoracleidcsextensionself_change_user_allow_self_change
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

data "oci_identity_domains_auth_tokens" "test_auth_tokens" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain.url

  #Optional
  auth_token_count  = var.auth_token_auth_token_count
  auth_token_filter = "user.value eq \"${oci_identity_domains_user.test_user.id}\""
  attribute_sets    = []
  attributes        = ""
  authorization     = var.auth_token_authorization
  #use the latest if not provided
  # resource_type_schema_version = var.auth_token_resource_type_schema_version
  start_index = var.auth_token_start_index
}

