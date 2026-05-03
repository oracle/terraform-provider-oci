// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "region" {}

variable "auth" {
  default = "SecurityToken"
}

variable "config_file_profile" {
  default = "boat-oc1-session"
}

variable "database_tools_cloud_api_gateway_config_id" {}

variable "database_tools_cloud_api_gateway_config_pool_key" {}

variable "database_tools_database_api_gateway_config_pool_auto_api_spec_alias" {
  default = "emp"
}

variable "database_tools_database_api_gateway_config_pool_auto_api_spec_database_object_type" {
  default = "TABLE"
}

variable "database_tools_database_api_gateway_config_pool_auto_api_spec_description" {
  default = "description"
}

variable "database_tools_database_api_gateway_config_pool_auto_api_spec_display_name" {
  default = "Employees Auto API"
}

variable "database_tools_database_api_gateway_config_pool_auto_api_spec_operations" {
  default = []
}

variable "database_tools_database_api_gateway_config_pool_auto_api_spec_pool_key" {
  default = "f0575ce0-bb39-453e-812d-82e8cdce80eb"
}

variable "database_tools_database_api_gateway_config_pool_auto_api_spec_database_object_name" {
  default = "EMPLOYEES"
}

variable "database_tools_database_api_gateway_config_pool_auto_api_spec_roles" {
  default = []
}

variable "database_tools_database_api_gateway_config_pool_auto_api_spec_scope" {
  default = "scope"
}

variable "database_tools_database_api_gateway_config_pool_auto_api_spec_security_schemes" {
  default = []
}

variable "database_tools_database_api_gateway_config_pool_auto_api_spec_type" {
  default = "DEFAULT"
}



provider "oci" {
  region              = var.region
  auth                = var.auth
  config_file_profile = var.config_file_profile
}

resource "oci_database_tools_runtime_database_tools_database_api_gateway_config_pool_auto_api_spec" "test_database_tools_database_api_gateway_config_pool_auto_api_spec" {
  #Required
  database_object_name                          = var.database_tools_database_api_gateway_config_pool_auto_api_spec_database_object_name
  database_object_type                          = var.database_tools_database_api_gateway_config_pool_auto_api_spec_database_object_type
  database_tools_database_api_gateway_config_id = var.database_tools_cloud_api_gateway_config_id
  display_name                                  = var.database_tools_database_api_gateway_config_pool_auto_api_spec_display_name
  pool_key                                      = var.database_tools_cloud_api_gateway_config_pool_key
  type                                          = var.database_tools_database_api_gateway_config_pool_auto_api_spec_type

  #Optional
  alias            = var.database_tools_database_api_gateway_config_pool_auto_api_spec_alias
  description      = var.database_tools_database_api_gateway_config_pool_auto_api_spec_description
  operations       = var.database_tools_database_api_gateway_config_pool_auto_api_spec_operations
  roles            = var.database_tools_database_api_gateway_config_pool_auto_api_spec_roles
  scope            = var.database_tools_database_api_gateway_config_pool_auto_api_spec_scope
  security_schemes = var.database_tools_database_api_gateway_config_pool_auto_api_spec_security_schemes
}

data "oci_database_tools_runtime_database_tools_database_api_gateway_config_pool_auto_api_specs" "test_database_tools_database_api_gateway_config_pool_auto_api_specs" {
  #Required
  database_tools_database_api_gateway_config_id = var.database_tools_cloud_api_gateway_config_id
  pool_key                                      = oci_database_tools_runtime_database_tools_database_api_gateway_config_pool_auto_api_spec.test_database_tools_database_api_gateway_config_pool_auto_api_spec.pool_key

  #Optional
  display_name = var.database_tools_database_api_gateway_config_pool_auto_api_spec_display_name
}
