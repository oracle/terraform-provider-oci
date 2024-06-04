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
  depends_on = [oci_database_autonomous_database.autonomous_database_cross_tenancy_dataguard_primary]
}

resource "oci_database_autonomous_database" "autonomous_database_cross_tenancy_dataguard_primary" {
  admin_password           = random_string.autonomous_database_admin_password.result
  compartment_id           = var.compartment_id
  cpu_core_count           = "1"
  data_storage_size_in_tbs = "1"
  db_name                  = "adbdb11ff6510"
  db_version               = "19c"
  license_model            = "LICENSE_INCLUDED"
  is_free_tier             = "false"
}

resource "oci_database_autonomous_database" "autonomous_database_cross_tenancy_dataguard_standby" {
  provider = oci.standby_tenancy
  compartment_id           = var.standby_compartment_id
  source    = "CROSS_TENANCY_DISASTER_RECOVERY"
  disaster_recovery_type = "ADG"
  source_id = oci_database_autonomous_database.autonomous_database_cross_tenancy_dataguard_primary.id
  db_name = oci_database_autonomous_database.autonomous_database_cross_tenancy_dataguard_primary.db_name
}

data "oci_database_autonomous_databases" "autonomous_databases" {
  filter {
    name   = "id"
    values = [oci_database_autonomous_database.autonomous_database_cross_tenancy_dataguard_standby.id]
  }

  filter {
    name   = "role"
    values = ["STANDBY"]
  }

  #Required
  compartment_id = var.standby_compartment_id

}

output "autonomous_databases" {
  value = data.oci_database_autonomous_databases.autonomous_databases.autonomous_databases
}

data "oci_database_autonomous_database_peers" "autonomous_database_peers" {
    autonomous_database_id = oci_database_autonomous_database.autonomous_database_cross_tenancy_dataguard_primary.id
}

output "autonomous_database_peers" {
    value = data.oci_database_autonomous_database_peers.autonomous_database_peers.autonomous_database_peer_collection
}



