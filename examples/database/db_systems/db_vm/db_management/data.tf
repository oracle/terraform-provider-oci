# Copyright (c) 2025, Oracle and/or its affiliates. All rights reserved.
#    NAME
#      data.tf - Data Source file
#
#    USAGE
#      Use the following path for Example and Backward-Compatibility Tests: database/db_systems/db_vm/db_management
#    NOTES
#      Associated Integration Test: TestDatabaseCloudDatabaseManagementResource_basic
#    FILES
#
#    DESCRIPTION
#
#    MODIFIED   MM/DD/YY
#    escabrer   06/23/2025 - Created


data "oci_identity_availability_domains" "test_availability_domains" {
  compartment_id = var.compartment_id
}


data "oci_database_db_systems" "test_db_systems" {
  compartment_id = var.compartment_id
  filter {
    name = "id"
    values = [oci_database_db_system.test_db_system.id]
  }
}

data "oci_database_db_homes" "test_db_homes" {
  compartment_id = var.compartment_id
  db_system_id = oci_database_db_system.test_db_system.id
}

data "oci_database_databases" "test_databases" {
  compartment_id = var.compartment_id
  db_home_id = data.oci_database_db_homes.test_db_homes.db_homes.0.db_home_id
}

data "oci_database_database" "test_database" {
  database_id = data.oci_database_databases.test_databases.databases.0.id
}