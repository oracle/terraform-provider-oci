# $Header$
#
# Copyright (c) 2025, Oracle and/or its affiliates. All rights reserved.
#    NAME
#      data.tf - data source file
#
#    USAGE
#      Use the following path for Example Test & Backward-Compatibility-Test: database/pluggable_databases/local_clone
#    NOTES
#      Terraform Example:
#    FILES
#
#    DESCRIPTION
#
#    MODIFIED   MM/DD/YY
#    escabrer   05/08/2025 - Created


data "oci_identity_availability_domains" "test_availability_domain" {
  compartment_id = var.compartment_id
}

data "oci_database_db_systems" "test_db_system" {
  compartment_id = var.compartment_id
  filter {
    name   = "id"
    values = [oci_database_db_system.test_db_system.id]
  }
}

data "oci_database_db_homes" "test_db_homes" {
  compartment_id = var.compartment_id
  db_system_id = oci_database_db_system.test_db_system.id
  filter {
    name   = "db_system_id"
    values = [oci_database_db_system.test_db_system.id]
  }
}

data "oci_database_databases" "test_databases" {
  compartment_id = var.compartment_id
  db_home_id = data.oci_database_db_homes.test_db_homes.db_homes.0.db_home_id
  filter {
    name   = "db_name"
    values = [oci_database_db_system.test_db_system.db_home.0.database.0.db_name]
  }
}

data "oci_database_database" "test_database" {
  database_id = data.oci_database_databases.test_databases.databases.0.id
}

data "oci_database_pluggable_database" "test_pluggable_database" {
  pluggable_database_id = oci_database_pluggable_database.test_pluggable_database.id
}
