// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "java_family_family_version" {
  default = "11"
}
variable "java_family_display_name" {}

data "oci_jms_java_families" "test_java_families" {
  display_name   = var.java_family_display_name
  family_version = var.java_family_family_version
}

data "oci_jms_java_family" "test_java_family" {
  family_version = var.java_family_family_version
}