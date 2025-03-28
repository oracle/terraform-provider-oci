# $Header$
#
# Copyright (c) 2024, Oracle and/or its affiliates. All rights reserved.
#
#    NAME
#      data.tf
#
#    USAGE
#      Use the following path for the Example & Backward Compatibility tests: database/db_systems/db_vm/db_upgrade_from_database_software_image
#    NOTES
#      Terraform Integration Test: TestDatabaseDatabaseUpgradeResource_DbSoftwareImage
#
#    FILE(S)
#      database_database_upgrade_resource_test.go
#
#    MODIFIED   MM/DD/YY
#    escabrer   12/16/2024 - Created


data "oci_identity_availability_domains" "test_availability_domains" {
  compartment_id = var.tenancy_ocid
}

data "oci_database_db_systems" "test_db_system_for_upgrade" {
  compartment_id = var.compartment_id
  filter {
    name = "id"
    values = [oci_database_db_system.test_db_system_for_upgrade.id]
  }
}

data "oci_database_db_homes" "test_db_system_for_upgrade" {
  compartment_id = var.compartment_id
  db_system_id = oci_database_db_system.test_db_system_for_upgrade.id
}

data "oci_database_db_home" "test_db_system_for_upgrade" {
  db_home_id = data.oci_database_db_homes.test_db_system_for_upgrade.db_homes.0.db_home_id
}

data "oci_database_databases" "test_db_system_for_upgrade" {
  compartment_id = var.compartment_id
  db_home_id = data.oci_database_db_homes.test_db_system_for_upgrade.db_homes.0.db_home_id
}

data "oci_database_database" "test_db_system_for_upgrade" {
  database_id = data.oci_database_databases.test_db_system_for_upgrade.databases.0.id
}