# Copyright (c) 2025, Oracle and/or its affiliates. All rights reserved.
#
#    NAME
#      datasource.tf
#
#    USAGE
#      Use the following path for the Example & Backward Compatibility tests: database/db_systems/db_vm/multicloud
#    NOTES
#      Terraform Integration Test: TestResourceDatabaseDBSystemMultiCloud
#
#    FILE(S)
#      database_db_system_resource_multicloud_test.go
#
#    MODIFIED   MM/DD/YY
#    escabrer   08/28/2025 - Created



data "oci_identity_availability_domain" "test_multicloud_availability_domain" {
  ad_number = "3"
  compartment_id = var.multicloud_compartment_id
}

data "oci_database_db_systems" "test_multicloud_db_systems" {
  compartment_id = var.multicloud_compartment_id
  filter {
    name = "id"
    values = [oci_database_db_system.test_multicloud_db_system.id]
  }
}




