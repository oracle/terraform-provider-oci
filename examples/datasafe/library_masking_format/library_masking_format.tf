// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}

variable "library_masking_format_access_level" {
  default = "ACCESSIBLE"
}

variable "library_masking_format_compartment_id_in_subtree" {
  default = true
}

variable "library_masking_format_defined_tags_value" {
  default = "value"
}

variable "library_masking_format_description" {
  default = "description"
}

variable "library_masking_format_display_name" {
  default = "displayName"
}

variable "library_masking_format_format_entries_column_name" {
  default = "columnName"
}

variable "library_masking_format_format_entries_description" {
  default = "description"
}

variable "library_masking_format_format_entries_end_date" {
  default = "endDate"
}

variable "library_masking_format_format_entries_end_length" {
  default = 10
}

variable "library_masking_format_format_entries_end_value" {
  default = 1.0
}

variable "library_masking_format_format_entries_fixed_number" {
  default = 1.0
}

variable "library_masking_format_format_entries_fixed_string" {
  default = "fixedString"
}

variable "library_masking_format_format_entries_grouping_columns" {
  default = []
}

variable "library_masking_format_format_entries_length" {
  default = 10
}

variable "library_masking_format_format_entries_post_processing_function" {
  default = "postProcessingFunction"
}

variable "library_masking_format_format_entries_random_list" {
  default = []
}

variable "library_masking_format_format_entries_regular_expression" {
  default = "regularExpression"
}

variable "library_masking_format_format_entries_replace_with" {
  default = "replaceWith"
}

variable "library_masking_format_format_entries_schema_name" {
  default = "schemaName"
}

variable "library_masking_format_format_entries_sql_expression" {
  default = "sqlExpression"
}

variable "library_masking_format_format_entries_start_date" {
  default = "startDate"
}

variable "library_masking_format_format_entries_start_length" {
  default = 10
}

variable "library_masking_format_format_entries_start_position" {
  default = 10
}

variable "library_masking_format_format_entries_start_value" {
  default = 1.0
}

variable "library_masking_format_format_entries_type" {
  default = "DELETE_ROWS"
}

variable "library_masking_format_format_entries_user_defined_function" {
  default = "userDefinedFunction"
}

variable "library_masking_format_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "library_masking_format_library_masking_format_source" {
  default = "ORACLE"
}

variable "library_masking_format_sensitive_type_ids" {
  default = []
}

variable "library_masking_format_state" {
  default = "ACTIVE"
}

variable "library_masking_format_time_created_greater_than_or_equal_to" {
  default = "2018-01-01T00:00:00.000Z"
}

variable "library_masking_format_time_created_less_than" {
  default = "2038-01-01T00:00:00.000Z"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_data_safe_library_masking_format" "test_library_masking_format" {
  #Required
  compartment_id = var.compartment_ocid
  format_entries {
    #Required
    type = var.library_masking_format_format_entries_type

    #Optional
    column_name               = var.library_masking_format_format_entries_column_name
    description               = var.library_masking_format_format_entries_description
    end_date                  = var.library_masking_format_format_entries_end_date
    end_length                = var.library_masking_format_format_entries_end_length
    end_value                 = var.library_masking_format_format_entries_end_value
    fixed_number              = var.library_masking_format_format_entries_fixed_number
    fixed_string              = var.library_masking_format_format_entries_fixed_string
    grouping_columns          = var.library_masking_format_format_entries_grouping_columns
    length                    = var.library_masking_format_format_entries_length
    post_processing_function  = var.library_masking_format_format_entries_post_processing_function
    random_list               = var.library_masking_format_format_entries_random_list
    regular_expression        = var.library_masking_format_format_entries_regular_expression
    replace_with              = var.library_masking_format_format_entries_replace_with
    schema_name               = var.library_masking_format_format_entries_schema_name
    sql_expression            = var.library_masking_format_format_entries_sql_expression
    start_date                = var.library_masking_format_format_entries_start_date
    start_length              = var.library_masking_format_format_entries_start_length
    start_position            = var.library_masking_format_format_entries_start_position
    start_value               = var.library_masking_format_format_entries_start_value
    user_defined_function     = var.library_masking_format_format_entries_user_defined_function
  }

  #Optional
  description        = var.library_masking_format_description
  display_name       = var.library_masking_format_display_name
  freeform_tags      = var.library_masking_format_freeform_tags
  sensitive_type_ids = var.library_masking_format_sensitive_type_ids
}

data "oci_data_safe_library_masking_formats" "test_library_masking_formats" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  display_name                          = var.library_masking_format_display_name
  library_masking_format_id             = oci_data_safe_library_masking_format.test_library_masking_format.id
}

#datasafe/library_masking_format