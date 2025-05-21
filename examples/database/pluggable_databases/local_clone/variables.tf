# $Header$
#
# Copyright (c) 2025, Oracle and/or its affiliates. All rights reserved.
#    NAME
#      variables.tf - variables file
#
#    USAGE
#      Use the following path for Example Test & Backward-Compatibility-Test: database/pluggable_databases/local_clone
#    NOTES
#      Terraform Example:
#    FILES
#
#    DESCRIPTION
#
#    MODIFIED   MM/DD/YY
#    escabrer   05/08/2025 - Created


variable "tenancy_ocid" {
}

variable "ssh_public_key" {
}

variable "region" {
}

variable "compartment_id" {
}

variable defined_tag_namespace_name {
    default = ""
}

variable "cpu_core_count" {
    default = "1"
}
