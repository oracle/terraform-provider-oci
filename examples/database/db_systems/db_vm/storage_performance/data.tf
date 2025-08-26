# $Header$
#
# Copyright (c) 2025, Oracle and/or its affiliates. All rights reserved.
#    NAME
#      data.tf - Shepherd Data Source file
#
#    USAGE
#      Use the following path for Example Test & Backward-Compatibility-Test: database/db_systems/db_vm/storage_performance
#    NOTES
#      Terraform Integration Test: TestDatabaseDbSystemStoragePerformanceResource_basic
#    FILES
#
#    DESCRIPTION
#
#    MODIFIED   MM/DD/YYYY
#    jufabian   05/06/2025 - Created

data "oci_database_db_system_storage_performances" "test_db_system_storage_performances" {
  #Required
  storage_management = var.storage_management

  #Optional
  shape_type = var.shape_type

  #Optional
  database_edition = var.database_edition

  #Optional
  compartment_id = var.compartment_id
}