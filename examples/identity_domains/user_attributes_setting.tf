// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "user_attributes_setting_attribute_sets" {
  default = []
}

variable "user_attributes_setting_attributes" {
  default = "attributes"
}

variable "user_attributes_setting_authorization" {
  default = "authorization"
}

variable "user_attributes_setting_idcs_endpoint" {
  default = "idcsEndpoint"
}

variable "user_attributes_setting_resource_type_schema_version" {
  default = "resourceTypeSchemaVersion"
}

data "oci_identity_domains_user_attributes_setting" "test_user_attributes_setting" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain.url
  user_attributes_setting_id = "UserAttributesSettings"

  #Optional
  attribute_sets               = ["all"]
  attributes                   = ""
  authorization                = var.user_attributes_setting_authorization
}

data "oci_identity_domains_user_attributes_settings" "test_user_attributes_settings" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain.url

  #Optional
  attribute_sets               = ["all"]
  attributes                   = ""
  authorization                = var.user_attributes_setting_authorization
}

