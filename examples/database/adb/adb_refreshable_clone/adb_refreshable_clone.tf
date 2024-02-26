// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

provider "oci" {
  region           = var.region
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
}

data "oci_database_autonomous_db_versions" "test_autonomous_db_versions" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  db_workload    = var.autonomous_database_db_workload

  filter {
    name   = "version"
    values = ["19c"]
  }
}

data "oci_database_autonomous_db_versions" "test_autonomous_dw_versions" {
  compartment_id = var.compartment_id
  db_workload    = "DW"
}

resource "oci_database_autonomous_database" "test_autonomous_database_source" {
  admin_password           = "BEstrO0ng_#11"
  compartment_id           = var.compartment_id
  cpu_core_count           = "1"
  data_storage_size_in_tbs = "1"
  db_name                  = "rcB8w9HgKux1t1"
  db_version               = "19c"
  db_workload              = "OLTP"
  display_name             = "regular_source"
  is_dedicated             = "false"
  license_model            = "LICENSE_INCLUDED"
}

resource "oci_database_autonomous_database" "test_autonomous_database_refreshable_clone_manual" {
  compartment_id           = var.compartment_id
  db_name                  = "bjfjkXw4ZutTt2"
  cpu_core_count           = "1"
  data_storage_size_in_tbs = "1"
  is_dedicated             = "false"
  is_refreshable_clone     = "true"
  license_model            = "LICENSE_INCLUDED"
  refreshable_mode         = "MANUAL"
  source                   = "CLONE_TO_REFRESHABLE"
  source_id                = oci_database_autonomous_database.test_autonomous_database_source.id
}

resource "oci_database_autonomous_database" "test_autonomous_database_refreshable_clone_automatic" {
  compartment_id                                 = var.compartment_id
  db_name                                        = "bjfjkXw4ZutTt3"
  cpu_core_count                                 = "1"
  data_storage_size_in_tbs                       = "1"
  is_dedicated                                   = "false"
  is_refreshable_clone                           = "true"
  license_model                                  = "LICENSE_INCLUDED"
  refreshable_mode                               = "AUTOMATIC"
  auto_refresh_point_lag_in_seconds              = "5000"
  auto_refresh_frequency_in_seconds              = "6000"
  time_of_auto_refresh_start                     = formatdate("YYYY-MM-DD'T'hh:mm:ss'.000'Z", timeadd(timestamp(), "24h"))
  source                                         = "CLONE_TO_REFRESHABLE"
  source_id                                      = oci_database_autonomous_database.test_autonomous_database_source.id
}

data "oci_database_autonomous_database" "oci_database_autonomous_database_manual" {
  autonomous_database_id = oci_database_autonomous_database.test_autonomous_database_refreshable_clone_manual.id
}

data "oci_database_autonomous_databases" "oci_database_autonomous_databases_manual" {
  compartment_id = var.compartment_id

  filter {
    name   = "id"
    values = [oci_database_autonomous_database.test_autonomous_database_refreshable_clone_manual.id]
  }
}

output "autonomous_database_refreshable_clone_manual" {
  value = data.oci_database_autonomous_databases.oci_database_autonomous_databases_manual.autonomous_databases
}

data "oci_database_autonomous_database" "oci_database_autonomous_database_automatic" {
  autonomous_database_id = oci_database_autonomous_database.test_autonomous_database_refreshable_clone_automatic.id
}

data "oci_database_autonomous_databases" "oci_database_autonomous_databases_automatic" {
  compartment_id = var.compartment_id

  filter {
    name   = "id"
    values = [oci_database_autonomous_database.test_autonomous_database_refreshable_clone_automatic.id]
  }
}

output "autonomous_database_refreshable_clone_automatic" {
  value = data.oci_database_autonomous_databases.oci_database_autonomous_databases_automatic.autonomous_databases
}