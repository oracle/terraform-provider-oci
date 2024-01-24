// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "my_auth_token_my_auth_token_count" {
  default = 10
}

variable "my_auth_token_my_auth_token_filter" {
  default = ""
}

variable "my_auth_token_authorization" {
  default = "authorization"
}

variable "my_auth_token_description" {
  default = "description"
}

variable "my_auth_token_expires_on" {
  default = "2030-01-01T00:00:00Z"
}

variable "my_auth_token_start_index" {
  default = 1
}

variable "my_auth_token_status" {
  default = "ACTIVE"
}

variable "my_auth_token_tags_key" {
  default = "key"
}

variable "my_auth_token_tags_value" {
  default = "value"
}

variable "my_auth_token_user_ocid" {
  default = "ocid"
}

variable "my_auth_token_user_value" {
  default = "value"
}


resource "oci_identity_domains_my_auth_token" "test_my_auth_token" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain_for_my_endpoint.url
  schemas       = ["urn:ietf:params:scim:schemas:oracle:idcs:authToken"]

  #Optional
  authorization = var.my_auth_token_authorization
  description   = var.my_auth_token_description
  expires_on    = var.my_auth_token_expires_on
  #use the latest if not provided
  # resource_type_schema_version = var.my_auth_token_resource_type_schema_version
  status = var.my_auth_token_status
  tags {
    #Required
    key   = var.my_auth_token_tags_key
    value = var.my_auth_token_tags_value
  }

  /* #for my_* resources, `user` can only be set to current user
  user {

    #Optional
    ocid  = var.my_auth_token_user_ocid
    value = var.my_auth_token_user_value
  }
  */

  lifecycle {
    ignore_changes = [
      // ignore fields that will never be returned
      // my* resource will not return non-default fields
      tags,
      status
    ]
  }
}

data "oci_identity_domains_my_auth_tokens" "test_my_auth_tokens" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain_for_my_endpoint.url

  #Optional
  my_auth_token_count  = var.my_auth_token_my_auth_token_count
  my_auth_token_filter = var.my_auth_token_my_auth_token_filter
  authorization        = var.my_auth_token_authorization
  #use the latest if not provided
  # resource_type_schema_version = var.my_auth_token_resource_type_schema_version
  start_index = var.my_auth_token_start_index
}

