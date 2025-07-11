# Copyright (c) 2025, Oracle and/or its affiliates. All rights reserved.
#
#    NAME
#      variables.tf
#
#    USAGE
#      Use the following path for the Example & Backward Compatibility tests: database/dataguard/db_vm/multicloud
#    NOTES
#      Terraform Integration Test: TestDatabaseDataGuardAssociationResourceMultiCloud
#
#    FILE(S)
#      database_data_guard_association_multicloud_test.go
#
#    MODIFIED   MM/DD/YY
#    escabrer   08/28/2025 - Created


variable "tenancy_ocid" {
  type = string
}

variable "compartment_id" {
  type = string
}

variable "region" {
  type = string
}

variable "ssh_public_key" {
  type = string
}

variable "admin_password" {
  type = string
}

variable "multicloud_compartment_id" {
  type = string
}

variable "multicloud_domain" {
  type = string
}

variable "multicloud_nsg_id" {
  type = string
}

variable "multicloud_subnet_id" {
  type = string
}

variable "multicloud_subscription_id" {
  type = string
}

variable "multicloud_cluster_placement_group_id" {
  type = string
}