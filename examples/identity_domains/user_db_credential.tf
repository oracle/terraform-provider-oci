// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "user_db_credential_user_db_credential_count" {
  default = 10
}

variable "user_db_credential_authorization" {
  default = "authorization"
}

variable "user_db_credential_db_password" {
  default = "dbPassword123456"
}

variable "user_db_credential_description" {
  default = "description"
}

variable "user_db_credential_expires_on" {
  default = "2030-01-01T00:00:00Z"
}

variable "user_db_credential_start_index" {
  default = 1
}

variable "user_db_credential_status" {
  default = "ACTIVE"
}

variable "user_db_credential_tags_key" {
  default = "key"
}

variable "user_db_credential_tags_value" {
  default = "value"
}

variable "user_db_credential_urnietfparamsscimschemasoracleidcsextensionself_change_user_allow_self_change" {
  default = false
}

resource "oci_identity_domains_user_db_credential" "test_user_db_credential" {
  #Required
  db_password   = var.user_db_credential_db_password
  idcs_endpoint = data.oci_identity_domain.test_domain.url
  schemas       = ["urn:ietf:params:scim:schemas:oracle:idcs:UserDbCredentials"]

  #Optional
  attribute_sets  = ["all"]
  attributes     = ""
  authorization  = var.user_db_credential_authorization
  description    = var.user_db_credential_description
  expires_on     = var.user_db_credential_expires_on
  #use the latest if not provided
  # resource_type_schema_version = var.user_db_credential_resource_type_schema_version
  status = var.user_db_credential_status
  tags {
    #Required
    key   = var.user_db_credential_tags_key
    value = var.user_db_credential_tags_value
  }
  urnietfparamsscimschemasoracleidcsextensionself_change_user {

    #Optional
    allow_self_change = var.user_db_credential_urnietfparamsscimschemasoracleidcsextensionself_change_user_allow_self_change
  }
  user {
    #Required
    #must be a user that exists
    value = oci_identity_domains_user.test_user.id

    #Optional
    #use the ocid of the same user set in value
    ocid = oci_identity_domains_user.test_user.ocid
  }

  lifecycle {
    ignore_changes = [
      // ignore fields that will never be returned
      status,
      urnietfparamsscimschemasoracleidcsextensionself_change_user,
      db_password
    ]
  }
}

data "oci_identity_domains_user_db_credentials" "test_user_db_credentials" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain.url

  #Optional
  user_db_credential_count  = var.user_db_credential_user_db_credential_count
  user_db_credential_filter = "user.value eq \"${oci_identity_domains_user.test_user.id}\""
  attribute_sets            = []
  attributes                = ""
  authorization             = var.user_db_credential_authorization
  #use the latest if not provided
  # resource_type_schema_version = var.user_db_credential_resource_type_schema_version
  start_index = var.user_db_credential_start_index
}

