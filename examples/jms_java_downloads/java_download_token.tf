// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_jms_java_downloads_java_download_token" "test_java_download_token" {
  #Required
  compartment_id = var.tenancy_ocid
  description    = "Example token description for script friendly download"
  display_name   = "Unique-displayName-in-a-tenancy"
  java_version   = "11"
  license_type   = ["OTN"]
  time_expires   = "2026-12-31T00:00:00.000Z" # must be a future date

  #Optional
  is_default = false
}

data "oci_jms_java_downloads_java_download_tokens" "test_java_download_tokens" {
  #Required
  compartment_id = var.tenancy_ocid

  #Optional
  display_name   = "Unique-displayName-in-a-tenancy"
  family_version = "11"
  id             = "id"
  search_by_user = var.user_ocid
  state          = "ACTIVE"
  value          = "value"
}
