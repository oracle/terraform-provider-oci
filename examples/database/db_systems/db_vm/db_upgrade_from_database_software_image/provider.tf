# $Header$
#
# Copyright (c) 2024, Oracle and/or its affiliates. All rights reserved.
#    NAME
#      provider.tf - Shepherd Data Source file
#
#    USAGE
#      Use the following path for Example Test & Backward Compatibility Test: database/db_systems/db_vm/db_upgrade_from_database_software_image
#    NOTES
#      Terraform Example: TestDatabaseDatabaseUpgradeResource_DbSoftwareImage
#    FILES
#
#    DESCRIPTION
#
#    MODIFIED   MM/DD/YY
#    escabrer   10/23/2024 - Created, removed old auth dependencies and migrated to SecurityToken


provider "oci" {
  auth                = "SecurityToken"
  config_file_profile = "terraform-federation-test"
  region              = var.region
  tenancy_ocid        = var.compartment_id
}