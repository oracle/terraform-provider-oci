// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "region" {}

variable "auth" {
  default = "SecurityToken"
}

variable "config_file_profile" {
  default = "terraform-federation-test"
}

variable "database_tools_connection_ocid" {}

provider "oci" {
  region              = var.region
  auth                = var.auth
  config_file_profile = var.config_file_profile
}

data "oci_database_tools_runtime_database_tools_connection_user_credentials" "test_database_tools_connection_user_credentials" {
  #Required
  database_tools_connection_id = var.database_tools_connection_ocid
  user_key                     = "APEX_240200"
}
