# Copyright (c) 2025, Oracle and/or its affiliates. All rights reserved.
#    NAME
#      provider.tf - provider file
#
#    USAGE
#      Use the following path for Example and Backward-Compatibility Tests: database/db_systems/db_vm/db_management
#    NOTES
#      Associated Integration Test: TestDatabaseCloudDatabaseManagementResource_basic
#    FILES
#
#    DESCRIPTION
#
#    MODIFIED   MM/DD/YY
#    escabrer   06/23/2025 - Created



provider "oci" {
  auth                = "SecurityToken"
  config_file_profile = "terraform-federation-test"
  region              = var.region
  tenancy_ocid        = var.compartment_id
}