// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tag_namespace_description" {
  default = "Just a test"
}

variable "tag_namespace_name" {
  default = "tf-testexamples-tag-namespace"
}
variable "defined_tags_value" {
  default = "tf-value"
}

variable "freeform_tags_value" {
  default = {
    "Department" = "Finance log"
  }
}

variable "tag_namespace1_name" {}
variable "tag2_name" {}

variable "compartment_id" {}

variable "log_group_name" {
  default = "tf-exampleLogGroup"
}

resource "oci_logging_log_group" "test_log_group" {
  #Required
  compartment_id = var.compartment_id
  display_name   = var.log_group_name

  #Optional
  description = "description"
  defined_tags = {
    "${var.tag_namespace1_name}.${var.tag2_name}" = var.defined_tags_value
  }
  freeform_tags      = var.freeform_tags_value

  lifecycle {
    ignore_changes = [ defined_tags ]
  }
}

data "oci_logging_log_groups" "test_log_groups" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name                 = var.log_group_name
  is_compartment_id_in_subtree = "false"
}

output "test_log_group_id" {
  value = oci_logging_log_group.test_log_group.id
}
