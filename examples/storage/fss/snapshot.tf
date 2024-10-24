// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_file_storage_snapshot" "my_snapshot" {
  #Required
  file_system_id = oci_file_storage_file_system.my_fs_1.id
  name           = var.snapshot_name
  # defined_tags = {
  #   "example-tag-namespace-all.example-tag" = "value"
  # }

  # Optional
  freeform_tags = {
    "Department" = "Finance"
  }
  # Commented out expiration_time as the date given should be a time in the future
  # expiration_time = "2096-01-02T15:04:05Z"


  locks {
    #Required
    type = var.locks_type

    #Optional
    message = var.locks_message
  }
  is_lock_override = var.is_lock_override
}

