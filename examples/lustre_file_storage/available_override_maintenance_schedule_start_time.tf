// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

# Wait for Poller to update Maintenance Metadata for created FS
resource "time_sleep" "wait_before_override" {
  create_duration = "10m"
  depends_on            = [oci_lustre_file_storage_lustre_file_system.maintenance_window_test_lfs]
}

# GET the file system (after wait)
data "oci_lustre_file_storage_lustre_file_system" "fs" {
  lustre_file_system_id = oci_lustre_file_storage_lustre_file_system.maintenance_window_test_lfs.id
  depends_on            = [time_sleep.wait_before_override]
}

output "planned_date" {
  value = local.planned_date
}

output "computed_override_date_plus_1" {
  value = local.override_date
}

# READ override slots for computed date
data "oci_lustre_file_storage_available_override_maintenance_start_times" "override_slots" {
  id   = oci_lustre_file_storage_lustre_file_system.maintenance_window_test_lfs.id
  date = local.override_date

  depends_on = [data.oci_lustre_file_storage_lustre_file_system.fs]
}

output "override_slots_raw" {
  value = data.oci_lustre_file_storage_available_override_maintenance_start_times.override_slots
}


