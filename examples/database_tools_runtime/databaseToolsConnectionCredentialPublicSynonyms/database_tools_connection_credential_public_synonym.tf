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

variable "database_tools_connection_credential_public_synonym_credential_key" {}

variable "database_tools_connection_credential_public_synonym_key" {}



provider "oci" {
  region              = var.region
  auth                = var.auth
  config_file_profile = var.config_file_profile
}

resource "oci_database_tools_runtime_database_tools_connection_credential_public_synonym" "test_database_tools_connection_credential_public_synonym" {
  #Required
  credential_key               = var.database_tools_connection_credential_public_synonym_credential_key
  database_tools_connection_id = var.database_tools_connection_ocid
  key                          = var.database_tools_connection_credential_public_synonym_key
}

data "oci_database_tools_runtime_database_tools_connection_credential_public_synonyms" "test_database_tools_connection_credential_public_synonyms" {
  #Required
  credential_key               = var.database_tools_connection_credential_public_synonym_credential_key
  database_tools_connection_id = var.database_tools_connection_ocid
}
