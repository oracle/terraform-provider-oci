# $Header$
#
# Copyright (c) 2024, Oracle and/or its affiliates. All rights reserved.
#    NAME
#      variables.tf - Shepherd Data Source file
#
#    USAGE
#
#    NOTES
#      Terraform Example: TestDatabaseDataGuardAssociationResource_basic
#    FILES
#
#    DESCRIPTION
#
#    MODIFIED   MM/DD/YY
#    escabrer   10/14/2024 - Created


variable "tenancy_ocid" {
}

variable "region" {
}

variable defined_tag_namespace_name {
  default = ""
}

variable "ssh_public_key" {
}

variable "compartment_id" {
}
