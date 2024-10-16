# $Header$
#
# Copyright (c) 2024, Oracle and/or its affiliates. All rights reserved.
#    NAME
#      datasources.tf - Shepherd Data Source file
#
#    USAGE
#
#    NOTES
#      Terraform Example: TestResourceDatabaseDBSystemFromDatabase
#    FILES
#
#    DESCRIPTION
#
#    MODIFIED   MM/DD/YY
#    escabrer   10/14/2024 - Created


data "oci_database_db_homes" "test_source_db_system" {
  compartment_id = var.compartment_id
  db_system_id = oci_database_db_system.test_source_db_system.id
}

data "oci_database_databases" "test_source_db_system" {
  compartment_id = var.compartment_id
  db_home_id = data.oci_database_db_homes.test_source_db_system.db_homes.0.db_home_id
}


data "oci_identity_availability_domains" "test_availability_domains" {
  compartment_id = var.tenancy_ocid
}