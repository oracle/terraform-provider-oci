# Copyright (c) 2025, Oracle and/or its affiliates. All rights reserved.
#
#    NAME
#      provider.tf
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


provider "oci" {
  auth                = "SecurityToken"
  config_file_profile = "terraform-federation-test"
  region              = var.region
  tenancy_ocid        = var.compartment_id
}