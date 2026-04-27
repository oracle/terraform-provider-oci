// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "database_tools_database_api_gateway_config_defined_tags_value" {
  default = "value"
}

variable "database_tools_database_api_gateway_config_display_name" {
  default = "MyDbApiConfig1"
}

variable "database_tools_database_api_gateway_config_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "database_tools_database_api_gateway_config_locks_message" {
  default = "message"
}


variable "database_tools_database_api_gateway_config_locks_type" {
  default = "FULL"
}

variable "database_tools_database_api_gateway_config_metadata_source" {
  default = "DATABASE"
}

variable "database_tools_database_api_gateway_config_state" {
  default = "ACTIVE"
}

variable "database_tools_database_api_gateway_config_type" {
  default = ["DEFAULT"]
}


provider "oci" {
  region           = var.region
  auth = "SecurityToken"
  config_file_profile = "terraform-federation-test"
}

resource "oci_database_tools_database_tools_database_api_gateway_config" "test_database_tools_database_api_gateway_config" {
  #Required
  compartment_id  = var.compartment_id
  display_name    = var.database_tools_database_api_gateway_config_display_name
  metadata_source = var.database_tools_database_api_gateway_config_metadata_source
  type            = var.database_tools_database_api_gateway_config_type[0]

  #Optional
  defined_tags  = {}
  freeform_tags = var.database_tools_database_api_gateway_config_freeform_tags
}

data "oci_database_tools_database_tools_database_api_gateway_configs" "test_database_tools_database_api_gateway_configs" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name = var.database_tools_database_api_gateway_config_display_name
  state        = var.database_tools_database_api_gateway_config_state
  type         = var.database_tools_database_api_gateway_config_type
}

