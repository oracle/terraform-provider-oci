// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}
variable "database_tools_mcp_server_id" {}

variable "database_tools_mcp_toolset_version_display_name" {
  default = "displayName"
}

provider "oci" {
  region              = var.region
  auth                = "SecurityToken"
  config_file_profile = "terraform-federation-test"
}

data "oci_database_tools_database_tools_mcp_toolset_versions" "test_database_tools_mcp_toolset_versions" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  database_tools_mcp_server_id = var.database_tools_mcp_server_id
  display_name                 = var.database_tools_mcp_toolset_version_display_name
}

