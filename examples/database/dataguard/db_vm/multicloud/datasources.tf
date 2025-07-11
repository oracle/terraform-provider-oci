# Copyright (c) 2025, Oracle and/or its affiliates. All rights reserved.
#
#    NAME
#      datasource.tf
#
#    USAGE
#      Use the following path for the Example & Backward Compatibility tests: database/dataguard/db_vm/multicloud
#    NOTES
#      Terraform Integration Test: TestDatabaseDataGuardAssociationResourceMultiCloud
#
#    FILE(S)
#      database_data_guard_association_multicloud_test.go
#
#    MODIFIED   MM/DD/YY
#    escabrer   08/28/2025 - Created



data "oci_identity_availability_domain" "test_multicloud_availability_domain" {
  ad_number = "3"
  compartment_id = var.multicloud_compartment_id
}

data "oci_database_db_homes" "test_multicloud_db_homes" {
  compartment_id = var.multicloud_compartment_id
  db_system_id = oci_database_db_system.test_multicloud_db_system.id
}

data "oci_database_databases" "test_multicloud_databases" {
  compartment_id = var.multicloud_compartment_id
  db_home_id = data.oci_database_db_homes.test_multicloud_db_homes.db_homes.0.db_home_id
}

data "oci_database_data_guard_association" "test_multicloud_dataguard_association" {
  data_guard_association_id = oci_database_data_guard_association.test_multicloud_dataguard_association.id
  database_id = data.oci_database_databases.test_multicloud_databases.databases.0.id
}

data "oci_database_data_guard_associations" "test_multicloud_dataguard_associations" {
  database_id = data.oci_database_databases.test_multicloud_databases.databases.0.id
  filter {
    name = "id"
    values = [oci_database_data_guard_association.test_multicloud_dataguard_association.id]
  }
}