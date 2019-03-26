// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

resource "oci_file_storage_snapshot" "my_snapshot" {
  #Required
  file_system_id = "${oci_file_storage_file_system.my_fs_1.id}"
  name           = "${var.snapshot_name}"
}
