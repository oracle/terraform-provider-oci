// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "grant_grant_count" {
  default = 10
}

variable "grant_grant_filter" {
  default = ""
}

variable "grant_authorization" {
  default = "authorization"
}

variable "grant_grant_mechanism" {
  default = "IMPORT_APPROLE_MEMBERS"
}

variable "grant_grantee_type" {
  default = "User"
}

variable "grant_start_index" {
  default = 1
}

variable "grant_tags_key" {
  default = "key"
}

variable "grant_tags_value" {
  default = "value"
}

resource "oci_identity_domains_user" "test_grant_user" {
  # Required
  emails {
    value = "value@email.com"
    type = "work"
    primary = "true"
  }
  idcs_endpoint = data.oci_identity_domain.test_domain.url
  name {
    family_name = "testGrantFamilyName"
  }
  schemas = ["urn:ietf:params:scim:schemas:core:2.0:User"]
  user_name = "testGrantUserName"
  lifecycle {
    ignore_changes = [
      urnietfparamsscimschemasoracleidcsextension_oci_tags[0].defined_tags,
      emails,
      schemas,
    ]
  }
}

data "oci_identity_domains_apps" "test_grant_apps" {
  idcs_endpoint = data.oci_identity_domain.test_domain.url
  app_filter                   = "displayName sw \"GrantTestApp\""
}

resource "oci_identity_domains_grant" "test_grant" {
  #Required
  grant_mechanism = var.grant_grant_mechanism
  grantee {
    #Required
    type  = var.grant_grantee_type
    value = oci_identity_domains_user.test_grant_user.id
  }
  idcs_endpoint = data.oci_identity_domain.test_domain.url
  schemas       = ["urn:ietf:params:scim:schemas:oracle:idcs:Grant"]

  #Optional
  app {
    #Required
    value = data.oci_identity_domains_apps.test_grant_apps.apps.0.id
  }
  attribute_sets = ["all"]
  attributes     = ""
  authorization  = var.grant_authorization

  #use the latest version if not provided
  #resource_type_schema_version  = var.grant_resource_type_schema_version
  tags {
    #Required
    key   = var.grant_tags_key
    value = var.grant_tags_value
  }
}

data "oci_identity_domains_grants" "test_grants" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain.url

  #Optional
  grant_count                  = var.grant_grant_count
  grant_filter                 = var.grant_grant_filter
  attribute_sets               = ["all"]
  attributes                   = ""
  authorization                = var.grant_authorization
  start_index                  = var.grant_start_index
  #use the latest version if not provided
  #resource_type_schema_version = var.grant_resource_type_schema_version
}

