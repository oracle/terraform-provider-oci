# $Header$
#
# Copyright (c) 2024, Oracle and/or its affiliates. All rights reserved.
#
#    NAME
#      provider.tf
#
#    USAGE
#      Use the following path for the Example & Backward Compatibility tests: database/db_systems/db_vm/db_upgrade_from_database_software_image
#    NOTES
#      Terraform Integration Test: TestDatabaseDatabaseUpgradeResource_DbSoftwareImage
#
#    FILE(S)
#      database_database_upgrade_resource_test.go
#
#    MODIFIED   MM/DD/YY
#    escabrer   12/16/2024 - Created


provider "oci" {
  auth                = "SecurityToken"
  config_file_profile = "terraform-federation-test"
  region              = var.region
  tenancy_ocid        = var.compartment_id
}