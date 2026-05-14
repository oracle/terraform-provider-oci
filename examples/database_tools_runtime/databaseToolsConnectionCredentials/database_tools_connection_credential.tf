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

variable "database_tools_connection_credential_key" {
  default = "key"
}

variable "database_tools_connection_credential_password" {
  default = "BEstrO0ng_#11"
}

variable "database_tools_connection_credential_type" {
  default = "BASIC"
}

variable "database_tools_connection_credential_user_name" {}

provider "oci" {
  region              = var.region
  auth                = var.auth
  config_file_profile = var.config_file_profile
}

resource "oci_database_tools_runtime_database_tools_connection_credential" "test_database_tools_connection_credential" {
  #Required
  database_tools_connection_id = var.database_tools_connection_ocid
  key                          = var.database_tools_connection_credential_key
  password                     = var.database_tools_connection_credential_password
  type                         = var.database_tools_connection_credential_type
  user_name                    = var.database_tools_connection_credential_user_name
}

data "oci_database_tools_runtime_database_tools_connection_credentials" "test_database_tools_connection_credentials" {
  #Required
  database_tools_connection_id = var.database_tools_connection_ocid
}
