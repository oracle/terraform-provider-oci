# $Header$
#
# Copyright (c) 2024, Oracle and/or its affiliates. All rights reserved.
#
#    NAME
#      variables.tf
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



variable "tenancy_ocid" {
  type = string
}

variable "ssh_public_key" {
  type = string
}

variable "region" {
  type = string
}

variable "compartment_ocid" {
  type = string
}

variable "defined_tag_namespace_name" {
  default = ""
}

variable "kms_key_id" {
  type = string
}

variable "vault_id" {
  type = string
}

variable "admin_password" {
  type = string
}