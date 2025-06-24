# Copyright (c) 2025, Oracle and/or its affiliates. All rights reserved.
#    NAME
#      variables.tf - Variables file
#
#    USAGE
#      Use the following path for Example and Backward-Compatibility Tests: database/db_systems/db_vm/db_management
#    NOTES
#      Associated Integration Test: TestDatabaseCloudDatabaseManagementResource_basic
#    FILES
#
#    DESCRIPTION
#
#    MODIFIED   MM/DD/YY
#    escabrer   06/23/2025 - Created



variable "tenancy_ocid" {
}

variable "compartment_id" {
}

variable "region" {
}

variable "ssh_public_key" {
}

variable "admin_password" {
}

variable "kms_key_id" {
}

variable "kms_key_version_id" {
}

variable "vault_id" {
}

variable "ssl_secret_id" {
  default = "test_secret_id"
}
