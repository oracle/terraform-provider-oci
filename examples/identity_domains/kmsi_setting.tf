// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "kmsi_setting_authorization" {
  default = "authorization"
}

variable "kmsi_setting_kmsi_feature_enabled" {
  default = false
}

variable "kmsi_setting_kmsi_prompt_enabled" {
  default = false
}

variable "kmsi_setting_last_enabled_on" {
  default = "2022-01-01T00:00:00Z"
}

variable "kmsi_setting_last_used_validity_in_days" {
  default = 10
}

variable "kmsi_setting_max_allowed_sessions" {
  default = 10
}

variable "kmsi_setting_schemas" {
  default = []
}

variable "kmsi_setting_tags_key" {
  default = "key"
}

variable "kmsi_setting_tags_value" {
  default = "value"
}

variable "kmsi_setting_tenancy_ocid" {
  default = "tenancyOcid"
}

variable "kmsi_setting_token_validity_in_days" {
  default = 10
}

variable "kmsi_setting_tou_prompt_disabled" {
  default = false
}


resource "oci_identity_domains_kmsi_setting" "test_kmsi_setting" {
  #Required
  idcs_endpoint   = data.oci_identity_domain.test_domain.url
  kmsi_setting_id = "KmsiSettings"
  schemas         = ["urn:ietf:params:scim:schemas:oracle:idcs:KmsiSettings"]

  #Optional
  attribute_sets             = ["all"]
  attributes                 = ""
  authorization              = var.kmsi_setting_authorization
  external_id                = "externalId"
  kmsi_feature_enabled       = var.kmsi_setting_kmsi_feature_enabled
  kmsi_prompt_enabled        = var.kmsi_setting_kmsi_prompt_enabled
  last_enabled_on            = var.kmsi_setting_last_enabled_on
  last_used_validity_in_days = var.kmsi_setting_last_used_validity_in_days
  max_allowed_sessions       = var.kmsi_setting_max_allowed_sessions
  #use the latest if not provided
  # resource_type_schema_version = var.kmsi_setting_resource_type_schema_version
  tags {
    #Required
    key   = var.kmsi_setting_tags_key
    value = var.kmsi_setting_tags_value
  }
  token_validity_in_days = var.kmsi_setting_token_validity_in_days
  tou_prompt_disabled    = var.kmsi_setting_tou_prompt_disabled
}

data "oci_identity_domains_kmsi_settings" "test_kmsi_settings" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain.url

  #Optional
  attribute_sets = []
  attributes     = ""
  authorization  = var.kmsi_setting_authorization
  #use the latest if not provided
  # resource_type_schema_version = var.kmsi_setting_resource_type_schema_version
}

