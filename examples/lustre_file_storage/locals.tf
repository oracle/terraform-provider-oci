# Compute next day from active/next planned maintenance date
locals {
  # FS metadata date like "2026-02-13"
  planned_date = data.oci_lustre_file_storage_lustre_file_system.fs.maintenance_window_metadata[0].active_or_next_planned_maintenance[0].date

  # Convert YYYY-MM-DD
  planned_date_formated = format("%sT00:00:00Z", local.planned_date)
  override_date         = substr(timeadd(local.planned_date_formated, "24h"), 0, 10)

  # Maintenance schedule slots
  maintenance_items = data.oci_lustre_file_storage_available_maintenance_schedule_start_times.available_maintenance.available_maintenance_schedule_start_time_collection[0].items

  # pick one of the slots returned by API (first day entry)
  chosen_day_obj = local.maintenance_items[0]

  # pick one of the times returned by API (first time)
  chosen_time = local.chosen_day_obj.start_times[0]

  # Override maintenance slots
  override_items = data.oci_lustre_file_storage_available_override_maintenance_start_times.override_slots.available_override_maintenance_start_time_collection[0].items

  chosen_override_item = local.override_items[0]

  # pick one of the start times returned by the API (first one)
  chosen_override_time = local.chosen_override_item.start_times[0]

  chosen_override_date = local.chosen_override_item.time_date_available
}
