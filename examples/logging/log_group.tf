// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "log_group_defined_tags_value" {
  default = "value2"
}

variable "log_group_freeform_tags" {
  default = {
    "Department" = "Finance"
  }
}

variable "tag_namespace_description" {
  default = "Just a test"
}

variable "tag_namespace_name" {
  default = "testexamples-tag-namespace"
}

resource "oci_identity_tag_namespace" "tag-namespace1" {
  #Required
  compartment_id = var.tenancy_ocid
  description    = var.tag_namespace_description
  name           = var.tag_namespace_name
}

resource "oci_identity_tag" "tag1" {
  #Required
  description      = "tf example tag"
  name             = "tf-example-tag"
  tag_namespace_id = oci_identity_tag_namespace.tag-namespace1.id
}

resource "oci_identity_tag" "tag2" {
  #Required
  description      = "tf example tag 2"
  name             = "tf-example-tag-2"
  tag_namespace_id = oci_identity_tag_namespace.tag-namespace1.id
}

resource "oci_logging_log_group" "test_log_group" {
  #Required
  compartment_id = var.compartment_ocid
  display_name   = "exampleLogGroup"

  #Optional
  defined_tags = {
    "${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = var.log_group_defined_tags_value
  }
  description = "description"

  freeform_tags = var.log_group_freeform_tags
}

data "oci_logging_log_groups" "test_log_groups" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  display_name                 = "exampleLogGroup"
  is_compartment_id_in_subtree = "false"
}

