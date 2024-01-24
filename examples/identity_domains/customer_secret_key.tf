// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "customer_secret_key_customer_secret_key_count" {
  default = 10
}

variable "customer_secret_key_authorization" {
  default = "authorization"
}

variable "customer_secret_key_description" {
  default = "description"
}

variable "customer_secret_key_display_name" {
  default = "displayName"
}

variable "customer_secret_key_expires_on" {
  default = "2030-01-01T00:00:00Z"
}

variable "customer_secret_key_start_index" {
  default = 1
}

variable "customer_secret_key_status" {
  default = "ACTIVE"
}

variable "customer_secret_key_tags_key" {
  default = "key"
}

variable "customer_secret_key_tags_value" {
  default = "value"
}

variable "customer_secret_key_urnietfparamsscimschemasoracleidcsextensionself_change_user_allow_self_change" {
  default = false
}


resource "oci_identity_domains_customer_secret_key" "test_customer_secret_key" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain.url
  schemas       = ["urn:ietf:params:scim:schemas:oracle:idcs:customerSecretKey"]

  #Optional
  attribute_sets = ["all"]
  attributes     = ""
  authorization  = var.customer_secret_key_authorization
  description    = var.customer_secret_key_description
  display_name   = var.customer_secret_key_display_name
  expires_on     = var.customer_secret_key_expires_on
  #use the latest if not provided
  # resource_type_schema_version = var.customer_secret_key_resource_type_schema_version
  status = var.customer_secret_key_status
  tags {
    #Required
    key   = var.customer_secret_key_tags_key
    value = var.customer_secret_key_tags_value
  }
  urnietfparamsscimschemasoracleidcsextensionself_change_user {

    #Optional
    allow_self_change = var.customer_secret_key_urnietfparamsscimschemasoracleidcsextensionself_change_user_allow_self_change
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

data "oci_identity_domains_customer_secret_keys" "test_customer_secret_keys" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain.url

  #Optional
  customer_secret_key_count  = var.customer_secret_key_customer_secret_key_count
  customer_secret_key_filter = "user.value eq \"${oci_identity_domains_user.test_user.id}\""
  attribute_sets             = []
  attributes                 = ""
  authorization              = var.customer_secret_key_authorization
  #use the latest if not provided
  # resource_type_schema_version = var.customer_secret_key_resource_type_schema_version
  start_index = var.customer_secret_key_start_index
}

