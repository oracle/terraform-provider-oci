// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "my_support_account_my_support_account_count" {
  default = 10
}

variable "my_support_account_my_support_account_filter" {
  default = ""
}

variable "my_support_account_my_support_account_provider" {
  default = "MOS"
}

variable "my_support_account_authorization" {
  default = "authorization"
}

variable "my_support_account_start_index" {
  default = 1
}

variable "my_support_account_tags_key" {
  default = "key"
}

variable "my_support_account_tags_value" {
  default = "value"
}

#provide the token
variable "my_support_account_token" {
  default = ""
}

variable "my_support_account_user_display" {
  default = "display"
}

variable "my_support_account_user_name" {
  default = "name"
}

variable "my_support_account_user_value" {
  default = "value"
}


resource "oci_identity_domains_my_support_account" "test_my_support_account" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain_for_my_endpoint.url
  schemas       = ["urn:ietf:params:scim:schemas:oracle:idcs:supportAccount"]
  token         = var.my_support_account_token

  #Optional
  authorization = var.my_support_account_authorization
  #use the latest if not provided
  # resource_type_schema_version = var.my_support_account_resource_type_schema_version
  tags {
    #Required
    key   = var.my_support_account_tags_key
    value = var.my_support_account_tags_value
  }

  /* #for my_* resources, `user` can only be set to current user
  user {

    #Optional
    ocid    = var.my_support_account_user_ocid
    value   = var.my_support_account_user_value
  }
  */
  lifecycle {
    ignore_changes = [
      // ignore fields that will never be returned
      // my* resource will not return non-default fields
      tags
    ]
  }

}

data "oci_identity_domains_my_support_accounts" "test_my_support_accounts" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain_for_my_endpoint.url

  #Optional
  my_support_account_count  = var.my_support_account_my_support_account_count
  my_support_account_filter = var.my_support_account_my_support_account_filter
  authorization             = var.my_support_account_authorization
  #use the latest if not provided
  # resource_type_schema_version = var.my_support_account_resource_type_schema_version
  start_index = var.my_support_account_start_index
}

