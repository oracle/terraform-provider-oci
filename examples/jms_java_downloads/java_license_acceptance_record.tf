// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_jms_java_downloads_java_license_acceptance_record" "test_java_license_acceptance_record" {
  #Required
  compartment_id            = var.tenancy_ocid
  license_acceptance_status = "ACCEPTED"
  license_type              = "OTN"
  lifecycle {
    ignore_changes = [defined_tags, system_tags]
  }
}

data "oci_jms_java_downloads_java_license_acceptance_records" "test_java_license_acceptance_records" {
  #Required
  compartment_id = var.tenancy_ocid

  #Optional
  id             = "id"
  license_type   = "OTN"
  search_by_user = var.user_ocid
  status         = "ACCEPTED"
}

