// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "database_tools_sql_report_columns_description" {
  default = "Simple SQL query"
}

variable "database_tools_sql_report_columns_name" {
  default = "name"
}

variable "database_tools_sql_report_columns_type" {
  default = "ORACLE_DATABASE"
}

variable "database_tools_sql_report_defined_tags_value" {
  default = "value"
}

variable "database_tools_sql_report_description" {
  default = "Simple SQL query"
}

variable "database_tools_sql_report_display_name" {
  default = "SQL Report1"
}

variable "database_tools_sql_report_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "database_tools_sql_report_instructions" {
  default = "instructions"
}

variable "database_tools_sql_report_locks_message" {
  default = "message"
}

variable "database_tools_sql_report_locks_time_created" {
  default = "timeCreated"
}

variable "database_tools_sql_report_locks_type" {
  default = "ORACLE_DATABASE"
}

variable "database_tools_sql_report_purpose" {
  default = "purpose"
}

variable "database_tools_sql_report_source" {
  default = "SELECT * FROM SYS.DBA_HIST_SYSTEM_EVENT WHERE SNAP_ID IN(SELECT SNAP_ID FROM SYS.DBA_HIST_SNAPSHOT WHERE BEGIN_INTERVAL_TIME>SYSDATE-1)"
}

variable "database_tools_sql_report_state" {
  default = "ACTIVE"
}

variable "database_tools_sql_report_type" {
  default = ["ORACLE_DATABASE"]
}

variable "database_tools_sql_report_variables_description" {
  default = "Simple SQL query"
}

variable "database_tools_sql_report_variables_name" {
  default = "name"
}

variable "database_tools_sql_report_variables_type" {
  default = "ORACLE_DATABASE"
}


provider "oci" {
  region              = var.region
  auth                = "SecurityToken"
  config_file_profile = "terraform-federation-test"
}

resource "oci_database_tools_database_tools_sql_report" "test_database_tools_sql_report" {
  #Required
  compartment_id = var.compartment_id
  display_name   = var.database_tools_sql_report_display_name
  source         = var.database_tools_sql_report_source
  type           = var.database_tools_sql_report_type[0]

  #Optional
  columns {
    #Required
    description = var.database_tools_sql_report_columns_description
    name        = var.database_tools_sql_report_columns_name
    type        = var.database_tools_sql_report_columns_type
  }
  defined_tags  = {}
  description   = var.database_tools_sql_report_description
  freeform_tags = var.database_tools_sql_report_freeform_tags
  instructions  = var.database_tools_sql_report_instructions
  purpose = var.database_tools_sql_report_purpose
  variables {
    #Required
    description = var.database_tools_sql_report_variables_description
    name        = var.database_tools_sql_report_variables_name
    type        = var.database_tools_sql_report_variables_type
  }
}

data "oci_database_tools_database_tools_sql_reports" "test_database_tools_sql_reports" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name = var.database_tools_sql_report_display_name
  state        = var.database_tools_sql_report_state
  type         = var.database_tools_sql_report_type
}

