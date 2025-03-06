# $Header$
#
# Copyright (c) 2024, Oracle and/or its affiliates. All rights reserved.
#    NAME
#      data.tf -  Data Source file
#
#    USAGE
#      Example & Backward Compatibility Path: database/db_systems/db_vm/db_backup
#    NOTES
#      Terraform Integration Test: TestDatabaseBackupResource_basic
#    FILES
#
#    DESCRIPTION
#
#    MODIFIED   MM/DD/YY
#    escabrer   11/1/2024 - Created


data "oci_identity_availability_domains" "test_availability_domains" {
  compartment_id = var.tenancy_ocid
}

data "oci_core_services" "test_services" {
  filter {
    name = "name"
    regex = "true"
    values = [".*Oracle.*Services.*Network"]
  }
}

data "oci_database_db_homes" "test_db_system" {
  compartment_id = var.compartment_id
  db_system_id = oci_database_db_system.test_db_system.id
  filter {
    name = "display_name"
    values = ["tfDbHome"]
  }
}

data "oci_database_databases" "test_db_system" {
  compartment_id = var.compartment_id
  db_home_id = data.oci_database_db_homes.test_db_system.db_homes.0.db_home_id
}