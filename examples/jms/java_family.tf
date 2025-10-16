// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

data "oci_jms_java_families" "test_java_families" {
  display_name   = "JDK 11"
  family_version = "11"
}

data "oci_jms_java_family" "test_java_family" {
  family_version = "11"
}