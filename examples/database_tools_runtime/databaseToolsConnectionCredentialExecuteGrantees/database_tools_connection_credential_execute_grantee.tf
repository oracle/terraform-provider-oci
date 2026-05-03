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

variable "database_tools_connection_credential_execute_grantee_credential_key" {
  default = "MY_TEST_CREDENTIAL"
}

variable "database_tools_connection_credential_execute_grantee_key" {
  default = "APEX_240200"
}



provider "oci" {
  region              = var.region
  auth                = var.auth
  config_file_profile = var.config_file_profile
}

resource "oci_database_tools_runtime_database_tools_connection_credential_execute_grantee" "test_database_tools_connection_credential_execute_grantee" {
  #Required
  credential_key               = var.database_tools_connection_credential_execute_grantee_credential_key
  database_tools_connection_id = var.database_tools_connection_ocid
  key                          = var.database_tools_connection_credential_execute_grantee_key
}

data "oci_database_tools_runtime_database_tools_connection_credential_execute_grantees" "test_database_tools_connection_credential_execute_grantees" {
  #Required
  credential_key               = var.database_tools_connection_credential_execute_grantee_credential_key
  database_tools_connection_id = var.database_tools_connection_ocid
}
