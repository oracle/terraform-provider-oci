# $Header$
#
# Copyright (c) 2024, Oracle and/or its affiliates. All rights reserved.
#
#    NAME
#      variables.tf
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


variable "tenancy_ocid" {
  type = string
}

variable "region" {
  type = string
}

variable defined_tag_namespace_name {
  default = ""
}

variable "compartment_id" {
  type = string
}

variable "ssh_public_key" {
  type = string
}

variable "database_software_image_id" {
  type = string
}
