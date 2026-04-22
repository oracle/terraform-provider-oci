# $Header$
#
# Copyright (c) 2026, Oracle and/or its affiliates. All rights reserved.
#    NAME
#      variables.tf - Variables file
#
#    USAGE
#      Example & Backward Compatibility Path: database/db_systems/db_vm/db_backup_source_dbrs

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

variable "defined_tag_namespace_name" {
  default = ""
}
