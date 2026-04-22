# $Header$
#
# Copyright (c) 2026, Oracle and/or its affiliates. All rights reserved.
#    NAME
#      provider.tf - Provider file
#
#    USAGE
#      Example & Backward Compatibility Path: database/db_systems/db_vm/db_backup_source_dbrs
#    NOTES
#      Terraform Integration Test: TestResourceDatabaseDBSystemSource

provider "oci" {
  auth                = "SecurityToken"
  config_file_profile = "terraform-federation-test"
  region              = var.region
  tenancy_ocid        = var.compartment_id
}
