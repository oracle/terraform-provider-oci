// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "my_oauth2client_credential_my_oauth2client_credential_count" {
  default = 10
}

variable "my_oauth2client_credential_my_oauth2client_credential_filter" {
  default = ""
}

variable "my_oauth2client_credential_authorization" {
  default = "authorization"
}

variable "my_oauth2client_credential_description" {
  default = "description"
}

variable "my_oauth2client_credential_expires_on" {
  default = "2030-01-01T00:00:00Z"
}

variable "my_oauth2client_credential_is_reset_secret" {
  default = false
}

variable "my_oauth2client_credential_name" {
  default = "name"
}

variable "my_oauth2client_credential_scopes_audience" {
  default = "audience"
}

variable "my_oauth2client_credential_scopes_scope" {
  default = "scope"
}

variable "my_oauth2client_credential_start_index" {
  default = 1
}

variable "my_oauth2client_credential_status" {
  default = "ACTIVE"
}

variable "my_oauth2client_credential_tags_key" {
  default = "key"
}

variable "my_oauth2client_credential_tags_value" {
  default = "value"
}

variable "my_oauth2client_credential_user_ocid" {
  default = "ocid"
}

variable "my_oauth2client_credential_user_value" {
  default = "value"
}


resource "oci_identity_domains_my_oauth2client_credential" "test_my_oauth2client_credential" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain_for_my_endpoint.url
  name          = var.my_oauth2client_credential_name
  schemas       = ["urn:ietf:params:scim:schemas:oracle:idcs:oauth2ClientCredential"]
  scopes {
    #Required
    audience = var.my_oauth2client_credential_scopes_audience
    scope    = var.my_oauth2client_credential_scopes_scope
  }

  #Optional
  authorization   = var.my_oauth2client_credential_authorization
  description     = var.my_oauth2client_credential_description
  expires_on      = var.my_oauth2client_credential_expires_on
  is_reset_secret = var.my_oauth2client_credential_is_reset_secret
  #use the latest if not provided
  # resource_type_schema_version = var.my_oauth2client_credential_resource_type_schema_version
  status = var.my_oauth2client_credential_status
  tags {
    #Required
    key   = var.my_oauth2client_credential_tags_key
    value = var.my_oauth2client_credential_tags_value
  }
  /* #for my_* resources, `user` can only be set to current user
  user {

    #Optional
    ocid  = var.my_oauth2client_credential_user_ocid
    value = var.my_oauth2client_credential_user_value
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

data "oci_identity_domains_my_oauth2client_credentials" "test_my_oauth2client_credentials" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain_for_my_endpoint.url

  #Optional
  my_oauth2client_credential_count  = var.my_oauth2client_credential_my_oauth2client_credential_count
  my_oauth2client_credential_filter = var.my_oauth2client_credential_my_oauth2client_credential_filter
  authorization                     = var.my_oauth2client_credential_authorization
  #use the latest if not provided
  # resource_type_schema_version      = var.my_oauth2client_credential_resource_type_schema_version
  start_index = var.my_oauth2client_credential_start_index
}

