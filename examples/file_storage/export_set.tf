// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_file_storage_export_set" "my_export_set_1" {
  # Required
  mount_target_id = oci_file_storage_mount_target.my_mount_target_1.id

  # Optional
  display_name      = var.export_set_name_1
  max_fs_stat_bytes = var.max_byte
  max_fs_stat_files = var.max_files
}

resource "oci_file_storage_export_set" "my_export_set_2" {
  # Required
  mount_target_id = oci_file_storage_mount_target.my_mount_target_2.id

  # Optional
  display_name      = var.export_set_name_2
  max_fs_stat_bytes = var.max_byte
  max_fs_stat_files = var.max_files
}

resource "oci_file_storage_export_set" "my_krb_export_set" {
  # Required
  mount_target_id = oci_file_storage_mount_target.my_krb_mount_target.id
}
