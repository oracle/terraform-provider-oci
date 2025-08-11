# $Header$
#
# Copyright (c) 2025, Oracle and/or its affiliates. All rights reserved.
#    NAME
#      variables.tf - Shepherd Variables file
#
#    USAGE
#      Use the following path for Example Test & Backward-Compatibility-Test: database/db_systems/db_vm/storage_performance
#    NOTES
#      Terraform Integration Test: TestDatabaseDbSystemStoragePerformanceResource_basic
#    FILES
#
#    DESCRIPTION
#
#    MODIFIED   MM/DD/YYYY
#    jufabian   05/06/2025 - Created



variable "tenancy_ocid" {
}

variable "region" {
}

variable "compartment_id" {
}

variable "storage_management" {
  default = "LVM"
}

variable "shape_type" {
   default = "AMPERE_FLEX_A1"
}

variable "database_edition" {
  default = "ENTERPRISE_EDITION_EXTREME_PERFORMANCE"
}