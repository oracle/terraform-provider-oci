# $Header$
#
# Copyright (c) 2024, Oracle and/or its affiliates. All rights reserved.
#    NAME
#      variables.tf
#
#    USAGE
#      Use the following path for the Example & Backward Compatibility tests: database/db_systems/db_vm/db_upgrade
#
#    NOTES
#      Terraform Integration Test: TestDatabaseDatabaseUpgradeResource_basic
#
#    FILE(S)
#      database_database_upgrade_test.go
#
#    MODIFIED   MM/DD/YY
#    escabrer   12/13/2024 - Created


variable "tenancy_ocid" {
  type = string
}

variable "region" {
  type = string
}

variable "defined_tag_namespace_name" {
  default = ""
}

variable "compartment_id" {
  type = string
}

variable "ssh_public_key" {
  type = string
}

variable "admin_password" {
  type = string
}