# $Header$
#
# Copyright (c) 2026, Oracle and/or its affiliates. All rights reserved.
#    NAME
#      data.tf - Data source file
#
#    USAGE
#      Example & Backward Compatibility Path: database/db_systems/db_vm/db_backup_source_dbrs

data "oci_identity_availability_domains" "test_availability_domains" {
  compartment_id = var.tenancy_ocid
}

data "oci_core_services" "test_services" {
  filter {
    name   = "name"
    regex  = "true"
    values = [".*Oracle.*Services.*Network"]
  }
}

data "oci_database_db_homes" "test_db_system" {
  compartment_id = var.compartment_id
  db_system_id   = oci_database_db_system.test_db_system.id

  filter {
    name   = "display_name"
    values = ["tfDbHome"]
  }
}

data "oci_database_databases" "test_db_system" {
  compartment_id = var.compartment_id
  db_home_id     = data.oci_database_db_homes.test_db_system.db_homes.0.db_home_id
}
