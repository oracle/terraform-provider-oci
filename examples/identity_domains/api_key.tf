// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "api_key_api_key_count" {
  default = 10
}

variable "api_key_authorization" {
  default = "authorization"
}

variable "api_key_description" {
  default = "description"
}

variable "api_key_fingerprint" {
  default = "fingerprint"
}

#provide the public key
variable "api_key_key" {
  default = ""
}

variable "api_key_start_index" {
  default = 1
}

variable "api_key_tags_key" {
  default = "key"
}

variable "api_key_tags_value" {
  default = "value"
}

variable "api_key_urnietfparamsscimschemasoracleidcsextensionself_change_user_allow_self_change" {
  default = false
}


resource "oci_identity_domains_api_key" "test_api_key" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain.url
  key           = var.api_key_key
  schemas       = ["urn:ietf:params:scim:schemas:oracle:idcs:apikey"]

  #Optional
  attribute_sets = ["all"]
  attributes     = ""
  authorization  = var.api_key_authorization
  description    = var.api_key_description
  tags {
    #Required
    key   = var.api_key_tags_key
    value = var.api_key_tags_value
  }
  urnietfparamsscimschemasoracleidcsextensionself_change_user {

    #Optional
    allow_self_change = var.api_key_urnietfparamsscimschemasoracleidcsextensionself_change_user_allow_self_change
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
      urnietfparamsscimschemasoracleidcsextensionself_change_user
    ]
  }
}

data "oci_identity_domains_api_keys" "test_api_keys" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain.url

  #Optional
  api_key_count  = var.api_key_api_key_count
  api_key_filter = "user.value eq \"${oci_identity_domains_user.test_user.id}\""
  attribute_sets = []
  attributes     = ""
  authorization  = var.api_key_authorization
  #use the latest if not provided
  # resource_type_schema_version = var.user_resource_type_schema_version
  start_index = var.api_key_start_index
}

