// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}

variable "sensitive_type_access_level" {
  default = "ACCESSIBLE"
}

variable "sensitive_type_comment_pattern" {
  default = "commentPattern"
}

variable "sensitive_type_compartment_id_in_subtree" {
  default = false
}

variable "sensitive_type_data_pattern" {
  default = "dataPattern"
}

variable "sensitive_type_defined_tags_value" {
  default = "value"
}

variable "sensitive_type_description" {
  default = "description"
}

variable "sensitive_type_display_name" {
  default = "displayName"
}

variable "sensitive_type_entity_type" {
  default = "SENSITIVE_TYPE"
}

variable "sensitive_type_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "sensitive_type_name_pattern" {
  default = "namePattern"
}

variable "sensitive_type_search_type" {
  default = "OR"
}

variable "sensitive_type_sensitive_type_source" {
  default = "ORACLE"
}

variable "sensitive_type_short_name" {
  default = "shortName"
}

variable "sensitive_type_state" {
  default = "AVAILABLE"
}

variable "sensitive_type_time_created_greater_than_or_equal_to" {
  default = "2018-01-01T00:00:00.000Z"
}

variable "sensitive_type_time_created_less_than" {
  default = "2038-01-01T00:00:00.000Z"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_data_safe_sensitive_type" "test_sensitive_type" {
  #Required
  compartment_id = var.compartment_ocid
  entity_type    = var.sensitive_type_entity_type

  #Optional
  comment_pattern           = var.sensitive_type_comment_pattern
  data_pattern              = var.sensitive_type_data_pattern
  description               = var.sensitive_type_description
  display_name              = var.sensitive_type_display_name
  freeform_tags             = var.sensitive_type_freeform_tags
  name_pattern              = var.sensitive_type_name_pattern
  search_type               = var.sensitive_type_search_type
  short_name                = var.sensitive_type_short_name

  lifecycle {
    ignore_changes = [defined_tags, system_tags, freeform_tags]
  }
}

data "oci_data_safe_sensitive_types" "test_sensitive_types" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  access_level                          = var.sensitive_type_access_level
  compartment_id_in_subtree             = var.sensitive_type_compartment_id_in_subtree
  display_name                          = var.sensitive_type_display_name
  entity_type                           = var.sensitive_type_entity_type
  sensitive_type_id                     = oci_data_safe_sensitive_type.test_sensitive_type.id
  sensitive_type_source                 = var.sensitive_type_sensitive_type_source

}

