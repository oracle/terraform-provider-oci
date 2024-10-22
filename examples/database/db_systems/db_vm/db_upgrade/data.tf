# $Header$
#
# Copyright (c) 2024, Oracle and/or its affiliates. All rights reserved.
#    NAME
#      data.tf - Datasources
#
#    USAGE
#      Use the following path for Example Test & Backward Compatibility Test: database/db_systems/db_vm/db_upgrade
#
#    NOTES
#      Terraform Example: TestDatabaseDatabaseUpgradeResource_basic
#    FILES
#
#    DESCRIPTION
#
#    MODIFIED   MM/DD/YY
#    escabrer   10/23/2024 - Created


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




