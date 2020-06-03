// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_file_storage_snapshot" "my_snapshot" {
  #Required
  file_system_id = "${oci_file_storage_file_system.my_fs_1.id}"
  name           = "${var.snapshot_name}"
  defined_tags   = "${map("example-tag-namespace-all.example-tag", "value")}"

  freeform_tags = {
    "Department" = "Finance"
  }
}
