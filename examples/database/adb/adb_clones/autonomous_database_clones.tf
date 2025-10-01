// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
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
  compute_count           = "1"
  compute_model = "ECPU"
  data_storage_size_in_tbs = "1"
  db_name                  = "adbDatabaseName123"

  #Optional
  db_version                                     = data.oci_database_autonomous_db_versions.test_autonomous_db_versions.autonomous_db_versions[0].version
  db_workload                                    = var.autonomous_database_db_workload
  display_name                                   = "example_autonomous_database"
  freeform_tags                                  = var.autonomous_database_freeform_tags
  is_auto_scaling_enabled                        = "true"
  is_auto_scaling_for_storage_enabled            = "true"
  license_model                                  = var.autonomous_database_license_model
  is_preview_version_with_service_terms_accepted = "false"
  whitelisted_ips                                = ["1.1.1.1/28"]
  character_set                                  = "AL32UTF8"
  ncharacter_set                                 = "AL16UTF16"
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


data "oci_database_autonomous_databases" "autonomous_databases" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  display_name = oci_database_autonomous_database.autonomous_database.display_name
  db_workload  = var.autonomous_database_db_workload
}

data "oci_database_autonomous_database_refreshable_clones" "autonomous_database_refreshable_clones" {
  #Required
  autonomous_database_id = oci_database_autonomous_database.autonomous_database.id
}

output "autonomous_database_admin_password" {
  value = random_string.autonomous_database_admin_password.result
}


data "oci_database_autonomous_databases_clones" "test_autonomous_databases_clones" {
  #Required
  autonomous_database_id = oci_database_autonomous_database.test_autonomous_database.id
  compartment_id = var.compartment_ocid

  #Optional
  clone_type = "REFRESHABLE_CLONE"
}
