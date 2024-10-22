# $Header$
#
# Copyright (c) 2024, Oracle and/or its affiliates. All rights reserved.
#    NAME
#      variables.tf - Shepherd Data Source file
#
#    USAGE
#      Use the following path for Example Test & Backward Compatibility Test: database/db_systems/db_vm/db_upgrade_from_database_software_image
#
#    NOTES
#      Terraform Example: TestDatabaseDatabaseUpgradeResource_DbSoftwareImage
#    FILES
#
#    DESCRIPTION
#
#    MODIFIED   MM/DD/YY
#    escabrer   10/23/2024 - Created


variable "tenancy_ocid" {
  type = string
}

variable "region" {
  type = string
}

variable defined_tag_namespace_name {
  type = string
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

