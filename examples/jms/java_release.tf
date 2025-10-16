// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

data "oci_jms_java_releases" "test_java_releases" {
  family_version      = "17"
  jre_security_status = "UPDATE_REQUIRED"
  license_type        = "OTN"
  release_type        = "CPU"
  release_version     = "17.0.2"
}

data "oci_jms_java_release" "test_java_release" {
  release_version     = "17.0.2"
}