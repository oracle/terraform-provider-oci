// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "region" {}
variable "compartment_id" {}

variable "database_tools_connection_ocid" {}
variable "domain_ocid" {}

variable "database_tools_mcp_server_access_token_expiry_in_seconds" {
  default = 60
}

variable "database_tools_mcp_server_custom_roles_description" {
  default = "description"
}

variable "database_tools_mcp_server_custom_roles_display_name" {
  default = "McpServer1"
}

variable "database_tools_mcp_server_description" {
  default = "description"
}

variable "database_tools_mcp_server_display_name" {
  default = "McpServer1"
}

variable "database_tools_mcp_server_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "database_tools_mcp_server_refresh_token_expiry_in_seconds" {
  default = 60
}

variable "database_tools_mcp_server_related_resource_identifier" {
  default = "relatedResourceIdentifier"
}

variable "database_tools_mcp_server_runtime_identity" {
  default = "AUTHENTICATED_PRINCIPAL"
}

variable "database_tools_mcp_server_state" {
  default = "ACTIVE"
}

variable "database_tools_mcp_server_storage_type" {
  default = "NONE"
}

variable "database_tools_mcp_server_type" {
  default = ["DEFAULT"]
}

provider "oci" {
  region              = var.region
  auth                = "SecurityToken"
  config_file_profile = "terraform-federation-test"
}

resource "oci_database_tools_database_tools_mcp_server" "test_database_tools_mcp_server" {
  #Required
  compartment_id               = var.compartment_id
  database_tools_connection_id = var.database_tools_connection_ocid
  display_name                 = var.database_tools_mcp_server_display_name
  domain_id                    = var.domain_ocid
  storage {
    #Required
    type = var.database_tools_mcp_server_storage_type
  }
  type = var.database_tools_mcp_server_type[0]

  #Optional
  access_token_expiry_in_seconds = var.database_tools_mcp_server_access_token_expiry_in_seconds
  custom_roles {
    #Required
    description  = var.database_tools_mcp_server_custom_roles_description
    display_name = var.database_tools_mcp_server_custom_roles_display_name
  }
  defined_tags                    = {}
  description                     = var.database_tools_mcp_server_description
  freeform_tags                   = var.database_tools_mcp_server_freeform_tags
  refresh_token_expiry_in_seconds = var.database_tools_mcp_server_refresh_token_expiry_in_seconds
  runtime_identity                = var.database_tools_mcp_server_runtime_identity
}

data "oci_database_tools_database_tools_mcp_servers" "test_database_tools_mcp_servers" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  database_tools_connection_id = var.database_tools_connection_ocid
  display_name                 = var.database_tools_mcp_server_display_name
  related_resource_identifier  = var.database_tools_mcp_server_related_resource_identifier
  state                        = var.database_tools_mcp_server_state
  type                         = var.database_tools_mcp_server_type
}