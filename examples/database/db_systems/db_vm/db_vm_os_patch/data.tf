# Copyright (c) 2026, Oracle and/or its affiliates. All rights reserved.

data "oci_identity_availability_domains" "test_availability_domains" {
  compartment_id = var.tenancy_ocid
}

data "oci_database_db_systems" "test_db_system_os_patch" {
  compartment_id = var.compartment_id
  filter {
    name   = "id"
    values = [oci_database_db_system.test_db_system_os_patch.id]
  }
}

data "oci_database_db_homes" "test_db_system_os_patch" {
  compartment_id = var.compartment_id
  db_system_id   = oci_database_db_system.test_db_system_os_patch.id
}

data "oci_database_db_home" "test_db_system_os_patch" {
  db_home_id = data.oci_database_db_homes.test_db_system_os_patch.db_homes.0.db_home_id
}

data "oci_database_databases" "test_db_system_os_patch" {
  compartment_id = var.compartment_id
  db_home_id     = data.oci_database_db_homes.test_db_system_os_patch.db_homes.0.db_home_id
}

data "oci_database_database" "test_db_system_os_patch" {
  database_id = data.oci_database_databases.test_db_system_os_patch.databases.0.id
}

data "oci_database_db_system_os_patch_history_entries" "test_db_system_os_patch_history_entries" {
  #Required
  db_system_id = oci_database_db_system.test_db_system_os_patch.id

  #Optional
  action = var.db_system_os_patch_history_entry_action

  # Ensure the patch precheck update is executed before reading history.
  depends_on = [oci_database_db_system.test_db_system_os_patch]

}

