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

variable "database_tools_database_api_gateway_config_global_global_key" {
  default = "SETTINGS"
}

variable "database_tools_database_api_gateway_config_global_type" {
  default = "DEFAULT"
}



provider "oci" {
  region              = var.region
  auth                = var.auth
  config_file_profile = var.config_file_profile
}

resource "oci_database_tools_runtime_database_tools_database_api_gateway_config_global" "test_database_tools_database_api_gateway_config_global" {
  #Required
  database_tools_database_api_gateway_config_id = var.database_tools_cloud_api_gateway_config_id
  global_key                                    = var.database_tools_database_api_gateway_config_global_global_key
  type                                          = var.database_tools_database_api_gateway_config_global_type
}
