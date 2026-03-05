# $Header$
#
# Copyright (c) 2024, Oracle and/or its affiliates. All rights reserved.
#    NAME
#      variables.tf - Shepherd Data Source file
#
#    USAGE
#
#    NOTES
#      Terraform Example: TestResourceDbSystemDataGuardAssociation


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