// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "random_string" "autonomous_data_warehouse_admin_password" {
  length      = 16
  min_numeric = 1
  min_lower   = 1
  min_upper   = 1
  min_special = 1
}

data "oci_database_autonomous_db_versions" "test_autonomous_dw_versions" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  db_workload = var.autonomous_data_warehouse_db_workload
}

resource "oci_database_autonomous_database" "autonomous_data_warehouse" {
  #Required
  admin_password           = random_string.autonomous_data_warehouse_admin_password.result
  compartment_id           = var.compartment_ocid
  cpu_core_count           = "1"
  data_storage_size_in_tbs = "1"
  db_name                  = "adbdw1f"

  #Optional
//If a db version is passed, it must be valid
//db_version              = data.oci_database_autonomous_db_versions.test_autonomous_dw_versions.autonomous_db_versions[0].version
  db_workload             = var.autonomous_data_warehouse_db_workload
  display_name            = "example_autonomous_data_warehouse"
  freeform_tags           = var.autonomous_database_freeform_tags
  is_auto_scaling_enabled = "false"
  license_model           = var.autonomous_database_license_model
}

data "oci_database_autonomous_databases" "autonomous_data_warehouses" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  display_name = oci_database_autonomous_database.autonomous_data_warehouse.display_name
  db_workload  = var.autonomous_data_warehouse_db_workload
}

output "autonomous_data_warehouse_admin_password" {
  value = random_string.autonomous_data_warehouse_admin_password.result
}

output "autonomous_data_warehouse_high_connection_string" {
  value = lookup(
    oci_database_autonomous_database.autonomous_data_warehouse.connection_strings[0].all_connection_strings,
    "high",
    "unavailable",
  )
}

output "autonomous_data_warehouses" {
  value = data.oci_database_autonomous_databases.autonomous_data_warehouses.autonomous_databases
}

