// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
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

resource "oci_file_storage_file_system" "my_file_system_with_fs_snapshot_policy" {
  #Required
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid

  #Optional
  display_name = var.file_system_with_snapshot_policy_display_name
  filesystem_snapshot_policy_id = oci_file_storage_filesystem_snapshot_policy.my_filesystem_snapshot_policy.id
  # defined_tags = {
  #   "example-tag-namespace-all.example-tag" = "value"
  # }

  freeform_tags = {
    "Department" = "Accounting"
  }
}

resource "oci_file_storage_file_system" "my_krb_file_system" {
  #Required
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid

  #Optional
  display_name = var.krb_file_system
  # defined_tags = {
  #   "example-tag-namespace-all.example-tag" = "value"
  # }

  freeform_tags = {
    "Department" = "Accounting"
  }
}