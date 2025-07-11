# $Header$
#
# Copyright (c) 2024, Oracle and/or its affiliates. All rights reserved.
#    NAME
#      provider.tf - Shepherd Data Source file
#
#    USAGE
#
#    NOTES
#      Terraform Example: TestResourceDatabaseDBSystemBasic
#    FILES
#
#    DESCRIPTION
#
#    MODIFIED   MM/DD/YY
#    escabrer   10/14/2024 - Created, removed old auth dependencies and migrated to SecurityToken



provider "oci" {
  auth                = "SecurityToken"
  config_file_profile = "terraform-federation-test"
  region              = var.region
  tenancy_ocid        = var.compartment_id
}