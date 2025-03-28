# $Header$
#
# Copyright (c) 2024, Oracle and/or its affiliates. All rights reserved.
#
#    NAME
#      data.tf
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



data "oci_identity_availability_domains" "test_availability_domains" {
  compartment_id = var.tenancy_ocid
}