// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "my_api_key_my_api_key_count" {
  default = 10
}

variable "my_api_key_my_api_key_filter" {
  default = ""
}

variable "my_api_key_authorization" {
  default = "authorization"
}

variable "my_api_key_description" {
  default = "description"
}

variable "my_api_key_fingerprint" {
  default = "fingerprint"
}

#provide the public key
variable "my_api_key_key" {
  default = ""
}

variable "my_api_key_start_index" {
  default = 1
}

variable "my_api_key_tags_key" {
  default = "key"
}

variable "my_api_key_tags_value" {
  default = "value"
}

variable "my_api_key_user_ocid" {
  default = "ocid"
}

variable "my_api_key_user_value" {
  default = "value"
}


resource "oci_identity_domains_my_api_key" "test_my_api_key" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain_for_my_endpoint.url
  key           = var.my_api_key_key
  schemas       = ["urn:ietf:params:scim:schemas:oracle:idcs:apikey"]

  #Optional
  authorization = var.my_api_key_authorization
  description   = var.my_api_key_description
  #use the latest if not provided
  # resource_type_schema_version = var.my_api_key_resource_type_schema_version
  tags {
    #Required
    key   = var.my_api_key_tags_key
    value = var.my_api_key_tags_value
  }

  /* #for my_* resources, `user` can only be set to current user
  user {

    #Optional
    ocid  = var.my_api_key_user_ocid
    value = var.my_api_key_user_value
  }
  */

  lifecycle {
    ignore_changes = [
      // ignore fields that will never be returned
      tags
    ]
  }
}

data "oci_identity_domains_my_api_keys" "test_my_api_keys" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain_for_my_endpoint.url

  #Optional
  my_api_key_count  = var.my_api_key_my_api_key_count
  my_api_key_filter = var.my_api_key_my_api_key_filter
  authorization     = var.my_api_key_authorization
  #use the latest if not provided
  # resource_type_schema_version = var.my_api_key_resource_type_schema_version
  start_index = var.my_api_key_start_index
}

