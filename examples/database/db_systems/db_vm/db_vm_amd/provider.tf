# $Header$
#
# Copyright (c) 2024, Oracle and/or its affiliates. All rights reserved.
#
#    NAME
#      provider.tf
#
#    USAGE
#      Use the following path for the Example & Backward Compatibility tests: database/db_systems/db_vm/db_vm_amd
#    NOTES
#      Terraform Integration Test: TestResourceDatabaseDBSystemAmdVM
#
#    FILE(S)
#      database_db_system_resource_amd_vm_test.go
#
#    MODIFIED   MM/DD/YY
#    escabrer   12/12/2024 - Created




provider "oci" {
  auth                = "SecurityToken"
  config_file_profile = "terraform-federation-test"
  region              = var.region
  tenancy_ocid        = var.compartment_ocid
}