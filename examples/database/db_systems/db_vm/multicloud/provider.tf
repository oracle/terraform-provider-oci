# Copyright (c) 2025, Oracle and/or its affiliates. All rights reserved.
#
#    NAME
#      provider.tf
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


provider "oci" {
  auth                = "SecurityToken"
  config_file_profile = "terraform-federation-test"
  region              = var.region
  tenancy_ocid        = var.compartment_id
}