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

variable "database_tools_connection_property_set_key" {
  default = "ORACLE_DATABASE_EXTERNAL_AUTHENTICATION"
}

variable "database_tools_connection_property_set_property_set_key" {
  default = "ORACLE_DATABASE_EXTERNAL_AUTHENTICATION"
}

variable "database_tools_connection_property_set_identity_provider_type" {
  default = "NONE"
}

provider "oci" {
  region              = var.region
  auth                = var.auth
  config_file_profile = var.config_file_profile
}

resource "oci_database_tools_runtime_database_tools_connection_property_set" "test_database_tools_connection_property_set" {
  #Required
  database_tools_connection_id = var.database_tools_connection_ocid
  key                          = var.database_tools_connection_property_set_key
  property_set_key             = var.database_tools_connection_property_set_property_set_key

  identity_provider {
    #Required
    type = var.database_tools_connection_property_set_identity_provider_type
  }
}

data "oci_database_tools_runtime_database_tools_connection_property_set" "test_database_tools_connection_property_set" {
  #Required
  database_tools_connection_id = var.database_tools_connection_ocid
  property_set_key             = var.database_tools_connection_property_set_property_set_key
}
