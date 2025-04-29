# $Header$
#
# Copyright (c) 2025, Oracle and/or its affiliates. All rights reserved.
#    NAME
#      provider.tf - Shepherd Provider file
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



provider "oci" {
  auth                = "SecurityToken"
  config_file_profile = "terraform-federation-test"
  region              = var.region
  tenancy_ocid        = var.compartment_id
}