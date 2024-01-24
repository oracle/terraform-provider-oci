// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0
resource "oci_file_storage_filesystem_snapshot_policy" "my_filesystem_snapshot_policy" {
  #Required
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid

  #Optional
  # defined_tags  = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.filesystem_snapshot_policy_defined_tags_value)
  display_name  = var.filesystem_snapshot_policy_display_name
  freeform_tags = var.filesystem_snapshot_policy_freeform_tags
  policy_prefix = var.filesystem_snapshot_policy_policy_prefix
  schedules {
    #Required
    period    = "YEARLY"
    time_zone = var.filesystem_snapshot_policy_schedules_time_zone

    #Optional
    day_of_month                  = var.filesystem_snapshot_policy_schedules_day_of_month
    hour_of_day                   = var.filesystem_snapshot_policy_schedules_hour_of_day
    month                         = var.filesystem_snapshot_policy_schedules_month
    retention_duration_in_seconds = var.filesystem_snapshot_policy_schedules_retention_duration_in_seconds
    schedule_prefix               = "yearly-schedule"
    # Commented out time_schedule_start as the date given should be a time in the future
    # time_schedule_start           = "2096-01-02T15:04:05Z"
  }
}
