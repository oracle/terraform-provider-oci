// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "compartment_id" {}

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
  default = "tf-testexamples-tag-namespace"
}

resource "oci_identity_tag_namespace" "tag-namespace1" {
  #Required
  compartment_id = var.compartment_id
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

output "tag_namespace1_name" {
  value = oci_identity_tag_namespace.tag-namespace1.name
}
output "tag1_name" {
  value = oci_identity_tag.tag1.name
}
output "tag2_name" {
  value = oci_identity_tag.tag2.name
}