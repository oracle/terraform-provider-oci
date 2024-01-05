// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "my_customer_secret_key_my_customer_secret_key_count" {
  default = 10
}

variable "my_customer_secret_key_my_customer_secret_key_filter" {
  default = ""
}

variable "my_customer_secret_key_access_key" {
  default = "accessKey"
}

variable "my_customer_secret_key_authorization" {
  default = "authorization"
}

variable "my_customer_secret_key_description" {
  default = "description"
}

variable "my_customer_secret_key_display_name" {
  default = "displayName"
}

variable "my_customer_secret_key_expires_on" {
  default = "2030-01-01T00:00:00Z"
}

variable "my_customer_secret_key_start_index" {
  default = 1
}

variable "my_customer_secret_key_status" {
  default = "ACTIVE"
}

variable "my_customer_secret_key_tags_key" {
  default = "key"
}

variable "my_customer_secret_key_tags_value" {
  default = "value"
}

variable "my_customer_secret_key_user_ocid" {
  default = "ocid"
}

variable "my_customer_secret_key_user_value" {
  default = "value"
}


resource "oci_identity_domains_my_customer_secret_key" "test_my_customer_secret_key" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain_for_my_endpoint.url
  schemas       = ["urn:ietf:params:scim:schemas:oracle:idcs:customerSecretKey"]

  #Optional
  authorization = var.my_customer_secret_key_authorization
  description   = var.my_customer_secret_key_description
  display_name  = var.my_customer_secret_key_display_name
  expires_on    = var.my_customer_secret_key_expires_on
  #use the latest if not provided
  # resource_type_schema_version = var.my_customer_secret_key_resource_type_schema_version
  status = var.my_customer_secret_key_status
  tags {
    #Required
    key   = var.my_customer_secret_key_tags_key
    value = var.my_customer_secret_key_tags_value
  }
  /* #for my_* resources, `user` can only be set to current user
  user {

    #Optional
    ocid  = var.my_customer_secret_key_user_ocid
    value = var.my_customer_secret_key_user_value
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

data "oci_identity_domains_my_customer_secret_keys" "test_my_customer_secret_keys" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain_for_my_endpoint.url

  #Optional
  my_customer_secret_key_count  = var.my_customer_secret_key_my_customer_secret_key_count
  my_customer_secret_key_filter = var.my_customer_secret_key_my_customer_secret_key_filter
  authorization                 = var.my_customer_secret_key_authorization
  #use the latest if not provided
  # resource_type_schema_version  = var.my_customer_secret_key_resource_type_schema_version
  start_index = var.my_customer_secret_key_start_index
}

