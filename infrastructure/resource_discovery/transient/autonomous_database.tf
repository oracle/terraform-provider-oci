// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "random_string" "autonomous_database_admin_password_rd" {
  length      = 16
  min_numeric = 1
  min_lower   = 1
  min_upper   = 1
  min_special = 1
}

data "oci_database_autonomous_db_versions" "test_autonomous_db_versions_rd" {
  #Required
  compartment_id = "${var.compartment_ocid}"

  #Optional
  db_workload = "${var.autonomous_database_db_workload}"

  filter {
    name   = "version"
    values = ["19c"]
  }
}

resource "oci_database_autonomous_database" "autonomous_database_rd" {
  #Required
  admin_password           = "${random_string.autonomous_database_admin_password_rd.result}"
  compartment_id           = "${var.compartment_ocid}"
  cpu_core_count           = "1"
  data_storage_size_in_tbs = "1"
  db_name                  = "adbdbrd"

  #Optional
  db_version                                     = "${data.oci_database_autonomous_db_versions.test_autonomous_db_versions_rd.autonomous_db_versions.0.version}"
  db_workload                                    = "${var.autonomous_database_db_workload}"
  display_name                                   = "autonomousDadtabaseRD"
  freeform_tags                                  = "${var.autonomous_database_freeform_tags}"
  is_auto_scaling_enabled                        = "true"
  license_model                                  = "${var.autonomous_database_license_model}"
  is_preview_version_with_service_terms_accepted = "false"
  whitelisted_ips                                = ["1.1.1.1/28"]
}

resource "random_string" "autonomous_data_warehouse_admin_password_rd" {
  length      = 16
  min_numeric = 1
  min_lower   = 1
  min_upper   = 1
  min_special = 1
}

data "oci_database_autonomous_db_versions" "test_autonomous_dw_versions_rd" {
  #Required
  compartment_id = "${var.compartment_ocid}"

  #Optional
  db_workload = "${var.autonomous_data_warehouse_db_workload}"
}

resource "oci_database_autonomous_database" "autonomous_data_warehouse_rd" {
  #Required
  admin_password           = "${random_string.autonomous_data_warehouse_admin_password_rd.result}"
  compartment_id           = "${var.compartment_ocid}"
  cpu_core_count           = "1"
  data_storage_size_in_tbs = "1"
  db_name                  = "adbdwrd"

  #Optional
  db_version              = "${data.oci_database_autonomous_db_versions.test_autonomous_dw_versions_rd.autonomous_db_versions.0.version}"
  db_workload             = "${var.autonomous_data_warehouse_db_workload}"
  display_name            = "autonomousDataWarehouseRD"
  freeform_tags           = "${var.autonomous_database_freeform_tags}"
  is_auto_scaling_enabled = "false"
  license_model           = "${var.autonomous_database_license_model}"
}
