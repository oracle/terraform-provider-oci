// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "log_saved_search_freeform_tags" {
  default = {
    "Department" = "Finance"
  }
}
variable "log_defined_tags_value" {
  default = "tf-value"
}
variable "log_group_defined_tags_value" {
  default = "tf-value-group"
}
variable "compartment_id" {}
variable "tag_namespace1_name" {}
variable "tag1_name" {}
variable "test_log_saved_search_name" {
  default = "tf-exampleLogSavedSearch"
}

resource "oci_logging_log_saved_search" "test_log_saved_search" {
  #Required
  compartment_id = var.compartment_id
  name           = var.test_log_saved_search_name
  query          = "exampleQuery"

  #Optional
  defined_tags = {
    "${var.tag_namespace1_name}.${var.tag1_name}" = var.log_group_defined_tags_value
  }
  description = "description"

  freeform_tags = var.log_saved_search_freeform_tags

  lifecycle {
    ignore_changes = [ defined_tags ]
  }
}

data "oci_logging_log_saved_searches" "test_log_saved_searches" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  name           = var.test_log_saved_search_name
  log_saved_search_id = oci_logging_log_saved_search.test_log_saved_search.id
}
