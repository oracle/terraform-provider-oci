// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_jms_java_downloads_java_download_report" "test_java_download_report_content_data" {
  #Required
  compartment_id = var.tenancy_ocid
  format         = "CSV"

  #Optional
  time_end      = "2024-08-01T03:07:27Z"
  time_start    = "2023-08-01T03:07:27Z"
  freeform_tags =  { "bar-key" = "value" }

  # Create the Tag namespace in OCI before enabling
  # See user guide: https://docs.oracle.com/en-us/iaas/Content/Tagging/Tasks/managingtagsandtagnamespaces.htm
  # defined_tags  = { "example-tag-namespace-all.example-tag" = "value" }
}

data "oci_jms_java_downloads_java_download_reports" "test_java_download_reports" {
  #Required
  compartment_id = var.tenancy_ocid

  #Optional
  java_download_report_id = "id"
  state                   = "ACTIVE"
}

