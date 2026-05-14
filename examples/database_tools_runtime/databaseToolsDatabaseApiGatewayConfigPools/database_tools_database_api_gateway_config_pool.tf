// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "region" {}

variable "auth" {
  default = "SecurityToken"
}

variable "config_file_profile" {
  default = "boat-oc1-session"
}

variable "database_tools_connection_ocid" {}

variable "database_tools_cloud_api_gateway_config_id" {}

variable "database_tools_database_api_gateway_config_pool_display_name" {
  default = "APIs for HR database"
}

variable "database_tools_database_api_gateway_config_pool_pool_route_value" {
  default = "poolRouteValue"
}

variable "database_tools_database_api_gateway_config_pool_type" {
  default = "DEFAULT"
}



provider "oci" {
  region              = var.region
  auth                = var.auth
  config_file_profile = var.config_file_profile
}

resource "oci_database_tools_runtime_database_tools_database_api_gateway_config_pool" "test_database_tools_database_api_gateway_config_pool" {
  #Required
  database_tools_connection_id                  = var.database_tools_connection_ocid
  database_tools_database_api_gateway_config_id = var.database_tools_cloud_api_gateway_config_id
  display_name                                  = var.database_tools_database_api_gateway_config_pool_display_name
  pool_route_value                              = var.database_tools_database_api_gateway_config_pool_pool_route_value
  type                                          = var.database_tools_database_api_gateway_config_pool_type
}

data "oci_database_tools_runtime_database_tools_database_api_gateway_config_pools" "test_database_tools_database_api_gateway_config_pools" {
  #Required
  database_tools_database_api_gateway_config_id = var.database_tools_cloud_api_gateway_config_id

  #Optional
  display_name = var.database_tools_database_api_gateway_config_pool_display_name
}
