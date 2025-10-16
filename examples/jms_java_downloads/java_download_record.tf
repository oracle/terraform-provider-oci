// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

data "oci_jms_java_downloads_java_download_records" "test_java_download_records" {
  #Required
  compartment_id = var.tenancy_ocid
}
