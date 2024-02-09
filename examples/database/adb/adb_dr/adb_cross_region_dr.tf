// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "random_string" "autonomous_database_admin_password" {
  length      = 16
  min_numeric = 1
  min_lower   = 1
  min_upper   = 1
  min_special = 1
}
resource "time_sleep" "wait_300_seconds" {
  destroy_duration = "5m"
  depends_on = [oci_database_autonomous_database.autonomous_database_cross_region_dr_primary]
}

resource "oci_database_autonomous_database" "autonomous_database_cross_region_dr_primary" {
  admin_password           = random_string.autonomous_database_admin_password.result
  compartment_id           = var.compartment_id
  cpu_core_count           = "1"
  data_storage_size_in_tbs = "1"
  db_name                  = "adbdb11ff2"
  db_version               = "19c"
  license_model            = "LICENSE_INCLUDED"
  is_free_tier             = "false"
}

resource "oci_database_autonomous_database" "autonomous_database_cross_region_dr_standby" {
  #Note: this should be provisioned in another region as the source database.
  provider = oci.peer_region

  #Required for cross-region standby
  compartment_id           = var.compartment_id
  source    = "CROSS_REGION_DISASTER_RECOVERY"
  source_id = oci_database_autonomous_database.autonomous_database_cross_region_dr_primary.id
  db_name = oci_database_autonomous_database.autonomous_database_cross_region_dr_primary.db_name
  is_replicate_automatic_backups = "true"
}

data "oci_database_autonomous_databases" "autonomous_databases" {
  filter {
    name   = "id"
    values = [oci_database_autonomous_database.autonomous_database_cross_region_dr_standby.id]
  }

  filter {
    name   = "peer_db_ids"
    values = [oci_database_autonomous_database.autonomous_database_cross_region_dr_primary.id]
  }

  filter {
    name   = "role"
    values = ["STANDBY"]
  }

  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name = oci_database_autonomous_database.autonomous_database_cross_region_dr_standby.display_name
}

output "autonomous_databases" {
  value = data.oci_database_autonomous_databases.autonomous_databases.autonomous_databases
}


