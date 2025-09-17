# $Header$
#
# Copyright (c) 2024, Oracle and/or its affiliates. All rights reserved.
#    NAME
#      provider.tf - Shepherd Data Source file
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



provider "oci" {
  auth                = "SecurityToken"
  config_file_profile = "terraform-federation-test"
  region              = var.region
  tenancy_ocid        = var.compartment_id
}