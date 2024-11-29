# $Header$
#
# Copyright (c) 2024, Oracle and/or its affiliates. All rights reserved.
#    NAME
#      variables.tf - Resources file
#
#    USAGE
#      Example & Backward Compatibility Path: database/db_systems/db_vm/db_backup
#    NOTES
#      Terraform Integration Test: TestDatabaseBackupResource_basic
#    FILES
#
#    DESCRIPTION
#
#    MODIFIED   MM/DD/YY
#    escabrer   11/1/2024 - Created


variable "tenancy_ocid" {
  type = string
}

variable "ssh_public_key" {
  type = string
}

variable "region" {
  type = string
}

variable "compartment_id" {
  type = string
}

variable defined_tag_namespace_name {
  default = ""
}