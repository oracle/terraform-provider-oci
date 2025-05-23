# Copyright (c) 2025, Oracle and/or its affiliates. All rights reserved.
#    NAME
#      data.tf - Data Source file
#
#    USAGE
#      Use the following path for Example Test & Backward-Compatibility-Test: database/db_systems/db_vm/patches
#    NOTES
#      Terraform Integration Test: TestDatabaseDbSystemPatchResource_basic
#    FILES
#
#    DESCRIPTION
#
#    MODIFIED   MM/DD/YY
#    escabrer   05/23/2025 - Created


data "oci_identity_availability_domains" "test_availability_domains" {
  compartment_id = var.compartment_id
}

data "oci_database_db_system_patches" "test_db_system_patches" {
  #Required
  db_system_id = oci_database_db_system.test_db_system.id
}
