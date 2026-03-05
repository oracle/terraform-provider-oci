// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

# 1) READ: get available maintenance schedule options
data "oci_lustre_file_storage_available_maintenance_schedule_start_times" "available_maintenance" {
  compartment_id      = var.compartment_ocid
  availability_domain = data.oci_identity_availability_domain.ad.name
}

output "maintenance_items" {
  value = data.oci_lustre_file_storage_available_maintenance_schedule_start_times.available_maintenance
}

output "chosen_maintenance_window" {
  value = {
    day  = local.chosen_day_obj.day_of_week
    time = local.chosen_time
  }
}

# 2) CREATE: create FS using the schedule returned above
resource "oci_lustre_file_storage_lustre_file_system" "maintenance_window_test_lfs" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  subnet_id = oci_core_subnet.my_subnet.id
  compartment_id      = var.compartment_ocid
  capacity_in_gbs     = var.lustre_file_system_capacity_in_gbs
  file_system_name    = var.lustre_file_system_name
  performance_tier    = var.lustre_file_system_performance_tier

  display_name               = var.lustre_file_system_display_name
  file_system_description    = var.lustre_file_system_file_system_description

  freeform_tags = {
    "Department" = "Finance"
  }

  root_squash_configuration {
    #Optional
    client_exceptions = var.lustre_file_system_root_squash_configuration_client_exceptions
    identity_squash   = var.lustre_file_system_root_squash_configuration_identity_squash
    squash_gid        = var.lustre_file_system_root_squash_configuration_squash_gid
    squash_uid        = var.lustre_file_system_root_squash_configuration_squash_uid
  }

  # Example: pick available maintenance from above read
  maintenance_window {
    day_of_week = local.chosen_day_obj.day_of_week
    time_start  = local.chosen_time
  }
}


