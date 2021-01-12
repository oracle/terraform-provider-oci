// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_file_storage_file_system" "my_fs_1" {
  #Required
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid

  #Optional
  display_name = var.file_system_1_display_name
  # defined_tags = {
  #   "example-tag-namespace-all.example-tag" = "value"
  # }

  freeform_tags = {
    "Department" = "Finance"
  }
}

resource "oci_file_storage_file_system" "my_fs_2" {
  #Required
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid

  #Optional
  display_name = var.file_system_2_display_name
  # defined_tags = {
  #   "example-tag-namespace-all.example-tag" = "value"
  # }

  freeform_tags = {
    "Department" = "Accounting"
  }
}

