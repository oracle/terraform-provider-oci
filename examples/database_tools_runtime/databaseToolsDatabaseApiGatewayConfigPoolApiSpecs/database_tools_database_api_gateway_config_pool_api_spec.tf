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

variable "database_tools_database_api_gateway_config_pool_api_spec_content" {
  default = "{\"openapi\":\"3.0.0\",\"info\":{\"title\":\"unit test spec\",\"version\":\"1.0\"},\"paths\":{\"/test\":{\"get\":{\"responses\":{},\"x-dbtools-operation\":{\"sourceType\":\"query\",\"source\":\"select 1 from dual\"}}}}}"
}

variable "database_tools_database_api_gateway_config_pool_api_spec_display_name" {
  default = "Employees API"
}

variable "database_tools_database_api_gateway_config_pool_api_spec_type" {
  default = "DEFAULT"
}



provider "oci" {
  region              = var.region
  auth                = var.auth
  config_file_profile = var.config_file_profile
}

resource "oci_database_tools_runtime_database_tools_database_api_gateway_config_pool_api_spec" "test_database_tools_database_api_gateway_config_pool_api_spec" {
  #Required
  content                                       = var.database_tools_database_api_gateway_config_pool_api_spec_content
  database_tools_database_api_gateway_config_id = var.database_tools_cloud_api_gateway_config_id
  display_name                                  = var.database_tools_database_api_gateway_config_pool_api_spec_display_name
  pool_key                                      = var.database_tools_cloud_api_gateway_config_pool_key
  type                                          = var.database_tools_database_api_gateway_config_pool_api_spec_type
}

data "oci_database_tools_runtime_database_tools_database_api_gateway_config_pool_api_specs" "test_database_tools_database_api_gateway_config_pool_api_specs" {
  #Required
  database_tools_database_api_gateway_config_id = var.database_tools_cloud_api_gateway_config_id
  pool_key                                      = var.database_tools_cloud_api_gateway_config_pool_key

  #Optional
  display_name = var.database_tools_database_api_gateway_config_pool_api_spec_display_name
}
