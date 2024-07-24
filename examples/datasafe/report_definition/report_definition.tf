// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "report_ocid" {}

variable "report_definition_access_level" {
  default = "RESTRICTED"
}

variable "report_definition_category" {
  default = "CUSTOM_REPORTS"
}

variable "report_definition_column_filters_expressions" {
  default = ["expressions"]
}

variable "report_definition_column_filters_field_name" {
  default = "targetName"
}

variable "report_definition_column_filters_is_enabled" {
  default = false
}

variable "report_definition_column_filters_is_hidden" {
  default = false
}

variable "report_definition_column_filters_operator" {
  default = "IN"
}

variable "report_definition_column_info_data_type" {
  default = "String"
}

variable "report_definition_column_info_display_name" {
  default = "Target Id"
}

variable "report_definition_column_info_display_order" {
  default = 1
}

variable "report_definition_column_info_field_name" {
  default = "targetId"
}

variable "report_definition_column_info_is_hidden" {
  default = true
}

variable "report_definition_column_sortings_field_name" {
  default = "targetName"
}

variable "report_definition_column_sortings_is_ascending" {
  default = false
}

variable "report_definition_column_sortings_sorting_order" {
  default = 10
}

variable "report_definition_compartment_id_in_subtree" {
  default = false
}

variable "report_definition_data_source" {
  default = "EVENTS"
}

variable "report_definition_defined_tags_value" {
  default = "value"
}

variable "report_definition_description" {
  default = "description"
}

variable "report_definition_display_name" {
  default = "displayName99"
}

variable "report_definition_is_seeded" {
  default = true
}

variable "report_definition_state" {
  default = "ACTIVE"
}

variable "report_definition_summary_count_of" {
  default = "targetName"
}

variable "report_definition_summary_display_order" {
  default = 10
}

variable "report_definition_summary_group_by_field_name" {
  default = "operation"
}

variable "report_definition_summary_is_hidden" {
  default = false
}

variable "report_definition_summary_name" {
  default = "name"
}

variable "report_definition_summary_scim_filter" {
  default = "operation eq \"LOGIN\""
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_data_safe_report_definition" "test_report_definition" {
  #Required
  column_filters {
    #Required
    expressions = var.report_definition_column_filters_expressions
    field_name  = var.report_definition_column_filters_field_name
    is_enabled  = var.report_definition_column_filters_is_enabled
    is_hidden   = var.report_definition_column_filters_is_hidden
    operator    = var.report_definition_column_filters_operator
  }
  column_info {
    #Required
    display_name  = var.report_definition_column_info_display_name
    display_order = var.report_definition_column_info_display_order
    field_name    = var.report_definition_column_info_field_name
    is_hidden     = var.report_definition_column_info_is_hidden

    #Optional
    data_type = var.report_definition_column_info_data_type
  }
  column_sortings {
    #Required
    field_name    = var.report_definition_column_sortings_field_name
    is_ascending  = var.report_definition_column_sortings_is_ascending
    sorting_order = var.report_definition_column_sortings_sorting_order
  }
  compartment_id = var.compartment_ocid
  display_name   = var.report_definition_display_name
  parent_id      = var.report_ocid
  summary {
    #Required
    display_order = var.report_definition_summary_display_order
    name          = var.report_definition_summary_name

    #Optional
    count_of            = var.report_definition_summary_count_of
    group_by_field_name = var.report_definition_summary_group_by_field_name
    is_hidden           = var.report_definition_summary_is_hidden
    scim_filter         = var.report_definition_summary_scim_filter
  }

  #Optional
  description   = var.report_definition_description

  lifecycle {
      ignore_changes = [system_tags]
    }
}

data "oci_data_safe_report_definitions" "test_report_definitions" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  data_source               = var.report_definition_data_source
}

