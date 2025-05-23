# Copyright (c) 2025, Oracle and/or its affiliates. All rights reserved.
#    NAME
#      provider.tf - provider file
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



provider "oci" {
  auth                = "SecurityToken"
  config_file_profile = "terraform-federation-test"
  region              = var.region
  tenancy_ocid        = var.compartment_id
}