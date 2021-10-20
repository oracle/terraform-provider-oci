// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "random_string" "autonomous_database_admin_password" {
  length      = 16
  min_numeric = 1
  min_lower   = 1
  min_upper   = 1
  min_special = 1
}

data "oci_database_autonomous_db_versions" "test_autonomous_db_versions" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  db_workload = var.autonomous_database_db_workload

  filter {
    name   = "version"
    values = ["19c"]
  }
}

resource "oci_database_autonomous_database" "autonomous_database" {
  #Required
  admin_password           = random_string.autonomous_database_admin_password.result
  compartment_id           = var.compartment_ocid
  cpu_core_count           = "1"
  data_storage_size_in_tbs = "1"
  db_name                  = "adbdb1f"

  #Optional
  db_version                                     = data.oci_database_autonomous_db_versions.test_autonomous_db_versions.autonomous_db_versions[0].version
  db_workload                                    = var.autonomous_database_db_workload
  display_name                                   = "example_autonomous_database"
  freeform_tags                                  = var.autonomous_database_freeform_tags
  is_auto_scaling_enabled                        = "true"
  license_model                                  = var.autonomous_database_license_model
  is_preview_version_with_service_terms_accepted = "false"
  whitelisted_ips                                = ["1.1.1.1/28"]
}

resource "oci_database_autonomous_database" "test_autonomous_database" {
  admin_password                       = random_string.autonomous_database_admin_password.result
  compartment_id                       = var.compartment_ocid
  cpu_core_count                       = "1"
  data_storage_size_in_tbs             = "1"
  db_name                              = "adbdb11f"
  db_version                           = "19c"
  db_workload                          = "AJD"
  license_model                        = "LICENSE_INCLUDED"
  is_free_tier                         = "false"
  autonomous_maintenance_schedule_type = var.autonomous_database_autonomous_maintenance_schedule_type
}

resource "oci_database_autonomous_database" "test_autonomous_database_apex" {
  admin_password           = random_string.autonomous_database_admin_password.result
  compartment_id           = var.compartment_ocid
  cpu_core_count           = "1"
  data_storage_size_in_tbs = "1"
  db_name                  = "adbapex"
  db_version               = "19c"
  db_workload              = "APEX"
  license_model            = "LICENSE_INCLUDED"
  is_free_tier             = "false"
}

resource "oci_database_autonomous_database" "test_autonomous_database_actions" {
  admin_password           = random_string.autonomous_database_admin_password.result
  compartment_id           = var.compartment_ocid
  cpu_core_count           = "1"
  data_storage_size_in_tbs = "1"
  db_name                  = "adbdb11ff"
  db_version               = "19c"
  db_workload              = "AJD"
  license_model            = "LICENSE_INCLUDED"
  is_free_tier             = "false"
  state                    = "STOPPED"
}

resource "oci_database_autonomous_database_backup" "autonomous_database_backup" {
  #Required
  autonomous_database_id = oci_database_autonomous_database.autonomous_database.id
  display_name           = var.autonomous_database_backup_display_name
}

resource "oci_database_autonomous_database" "autonomous_database_oneway_tls_connection" {
  admin_password              = random_string.autonomous_database_admin_password.result
  compartment_id              = var.compartment_ocid
  cpu_core_count              = "1"
  data_storage_size_in_tbs    = "1"
  db_name                     = "adbOneWay"

  whitelisted_ips             = ["1.1.1.1"]
  is_mtls_connection_required    = "true"
}

resource "oci_database_autonomous_database" "autonomous_database_dbms_status" {
  admin_password              = random_string.autonomous_database_admin_password.result
  compartment_id              = var.compartment_ocid
  cpu_core_count              = "1"
  data_storage_size_in_tbs    = "1"
  db_name                     = "adbdbms"
  database_management_status  = "ENABLED"
}

// Per service, we need to pass in a back up that is at least 2 hours old
/*
resource "oci_database_autonomous_database" "autonomous_database_from_backup_id" {
  #Required
  admin_password           = random_string.autonomous_database_admin_password.result
  compartment_id           = var.compartment_ocid
  cpu_core_count           = "1"
  data_storage_size_in_tbs = "1"
  db_name                  = "adbdb2"

  clone_type                    = "FULL"
  source                        = "BACKUP_FROM_ID"
  autonomous_database_backup_id = oci_database_autonomous_database_backup.autonomous_database_backup.id
}

resource "oci_database_autonomous_database" "autonomous_database_from_backup_timestamp" {
  #Required
  admin_password           = random_string.autonomous_database_admin_password.result
  compartment_id           = var.compartment_ocid
  cpu_core_count           = "1"
  data_storage_size_in_tbs = "1"
  db_name                  = "adbdb3"

  clone_type             = "FULL"
  source                 = "BACKUP_FROM_TIMESTAMP"
  autonomous_database_id = oci_database_autonomous_database_backup.autonomous_database_backup.autonomous_database_id
  timestamp              = oci_database_autonomous_database_backup.autonomous_database_backup.time_ended
}
*/

data "oci_database_autonomous_databases" "autonomous_databases" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  display_name = oci_database_autonomous_database.autonomous_database.display_name
  db_workload  = var.autonomous_database_db_workload
}

output "autonomous_database_admin_password" {
  value = random_string.autonomous_database_admin_password.result
}

output "autonomous_database_high_connection_string" {
  value = lookup(
    oci_database_autonomous_database.autonomous_database.connection_strings[0],
    "high",
    "unavailable",
  )
}

output "autonomous_databases" {
  value = data.oci_database_autonomous_databases.autonomous_databases.autonomous_databases
}

