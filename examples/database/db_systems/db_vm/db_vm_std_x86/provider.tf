# $Header$
#
# Copyright (c) 2024, Oracle and/or its affiliates. All rights reserved.
#
#    NAME
#      main.tf
#
#    USAGE
#      Use the following path for the Example & Backward Compatibility tests: database/db_systems/db_vm/db_vm_std_x86
#    NOTES
#      Terraform Integration Test: TestResourceDatabaseDBSystemVMStdx86
#
#    FILE(S)
#      database_db_system_resource_vm_std_x86_test.go
#
#    MODIFIED   MM/DD/YY
#    aavadhan   08/18/2025 - Created




provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}