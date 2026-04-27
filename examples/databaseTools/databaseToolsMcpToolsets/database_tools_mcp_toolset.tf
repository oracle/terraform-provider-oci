// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "region" {}
variable "compartment_id" {}

variable "database_tools_mcp_server_id" {}
variable "generative_ai_semantic_store_id" {}
variable "database_tools_sql_report_id" {}

variable "database_tools_mcp_toolset_allowed_roles" {
  default = []
}

variable "database_tools_mcp_toolset_default_execution_type" {
  default = "SYNCHRONOUS"
}

variable "database_tools_mcp_toolset_description" {
  default = "description"
}

variable "database_tools_mcp_toolset_display_name" {
  default = "toolset1"
}

variable "database_tools_mcp_toolset_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "database_tools_mcp_toolset_reports_allowed_roles" {
  default = []
}

variable "database_tools_mcp_toolset_source_type" {
  default = "INLINE"
}

variable "database_tools_mcp_toolset_source_value" {
  default = "SELECT 1 FROM dual"
}

variable "database_tools_mcp_toolset_state" {
  default = "ACTIVE"
}

variable "database_tools_mcp_toolset_tool_description" {
  default = "toolDescription"
}

variable "database_tools_mcp_toolset_tool_name" {
  default = "myTool"
}

variable "database_tools_mcp_toolset_tools_allowed_roles" {
  default = []
}

variable "database_tools_mcp_toolset_tools_name" {
  default = "name"
}

variable "database_tools_mcp_toolset_tools_status" {
  default = "ENABLED"
}

variable "database_tools_mcp_toolset_type" {
  default = ["CUSTOM_SQL_TOOL"]
}

variable "database_tools_mcp_toolset_version" {
  default = 1
}

provider "oci" {
  region              = var.region
  auth                = "SecurityToken"
  config_file_profile = "terraform-federation-test"
}

resource "oci_database_tools_database_tools_mcp_toolset" "test_database_tools_mcp_toolset" {
  #Required
  compartment_id               = var.compartment_id
  database_tools_mcp_server_id = var.database_tools_mcp_server_id
  display_name                 = var.database_tools_mcp_toolset_display_name
  type                         = var.database_tools_mcp_toolset_type[0]
  version                      = var.database_tools_mcp_toolset_version

  #Optional
  allowed_roles                   = var.database_tools_mcp_toolset_allowed_roles
  default_execution_type          = var.database_tools_mcp_toolset_default_execution_type
  defined_tags                    = {}
  description                     = var.database_tools_mcp_toolset_description
  freeform_tags                   = var.database_tools_mcp_toolset_freeform_tags
  generative_ai_semantic_store_id = var.generative_ai_semantic_store_id
  reports {
    #Optional
    allowed_roles                = var.database_tools_mcp_toolset_reports_allowed_roles
    database_tools_sql_report_id = var.database_tools_sql_report_id
  }
  source {
    #Optional
    type  = var.database_tools_mcp_toolset_source_type
    value = var.database_tools_mcp_toolset_source_value
  }
  tool_description = var.database_tools_mcp_toolset_tool_description
  tool_name        = var.database_tools_mcp_toolset_tool_name
  tools {
    #Optional
    allowed_roles = var.database_tools_mcp_toolset_tools_allowed_roles
    name          = var.database_tools_mcp_toolset_tools_name
    status        = var.database_tools_mcp_toolset_tools_status
  }
}

data "oci_database_tools_database_tools_mcp_toolsets" "test_database_tools_mcp_toolsets" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  database_tools_mcp_server_id = var.database_tools_mcp_server_id
  display_name                 = var.database_tools_mcp_toolset_display_name
  state                        = var.database_tools_mcp_toolset_state
  type                         = var.database_tools_mcp_toolset_type
}