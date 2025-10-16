// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

data "oci_jms_java_downloads_java_licenses" "test_java_licenses" {

  #Optional
  display_name = "Oracle Technology Network"
  license_type = "OTN"
}
