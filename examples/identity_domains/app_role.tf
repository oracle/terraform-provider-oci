// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "app_role_app_role_count" {
  default = 10
}

variable "app_role_app_role_filter" {
  default = ""
}

variable "app_role_admin_role" {
  default = false
}

variable "app_role_authorization" {
  default = "authorization"
}

variable "app_role_available_to_clients" {
  default = false
}

variable "app_role_available_to_groups" {
  default = false
}

variable "app_role_available_to_users" {
  default = false
}

variable "app_role_description" {
  default = "description"
}

variable "app_role_display_name" {
  default = "displayName"
}

variable "app_role_public" {
  default = false
}

variable "app_role_resource_type_schema_version" {
  default = "resourceTypeSchemaVersion"
}

variable "app_role_start_index" {
  default = 1
}

variable "app_role_tags_key" {
  default = "key"
}

variable "app_role_tags_value" {
  default = "value"
}


resource "oci_identity_domains_app_role" "test_app_role" {
  #Required
  app {
    #Required
    value = oci_identity_domains_app.test_app.id
  }
  display_name  = var.app_role_display_name
  idcs_endpoint = data.oci_identity_domain.test_domain.url
  schemas       = ["urn:ietf:params:scim:schemas:oracle:idcs:AppRole"]

  #Optional
  admin_role                   = var.app_role_admin_role
  attribute_sets               = ["all"]
  attributes                   = ""
  authorization                = var.app_role_authorization
  available_to_clients         = var.app_role_available_to_clients
  available_to_groups          = var.app_role_available_to_groups
  available_to_users           = var.app_role_available_to_users
  description                  = var.app_role_description
  legacy_group_name            = "legacyGroupName"
  public                       = var.app_role_public
  #use the latest if not provided
  # resource_type_schema_version = var.app_resource_type_schema_version
  tags {
    #Required
    key   = var.app_role_tags_key
    value = var.app_role_tags_value
  }
}

data "oci_identity_domains_app_roles" "test_app_roles" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain.url

  #Optional
  app_role_count               = var.app_role_app_role_count
  app_role_filter              = var.app_role_app_role_filter
  attribute_sets               = ["all"]
  attributes                   = ""
  authorization                = var.app_role_authorization
  #use the latest if not provided
  # resource_type_schema_version = var.app_role_resource_type_schema_version
  start_index                  = var.app_role_start_index
}

