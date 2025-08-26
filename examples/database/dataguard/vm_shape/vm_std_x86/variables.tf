# $Header$
#
# Copyright (c) 2024, Oracle and/or its affiliates. All rights reserved.
#    NAME
#      variables.tf - Shepherd Data Source file
#
#    USAGE
#
#    NOTES
#      Terraform Example: TestDatabaseDataGuardAssociationResourceVmStdx86_basic
#    FILES
#
#    DESCRIPTION
#
#    MODIFIED   MM/DD/YY
#    aavadhan   08/18/2025 - Created


variable "tenancy_ocid" {
}

variable "region" {
}

variable defined_tag_namespace_name {
  default = ""
}

variable "ssh_public_key" {
  type = string
}

variable "compartment_id" {
}
