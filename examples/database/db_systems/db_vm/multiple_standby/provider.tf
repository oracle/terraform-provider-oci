
# $Header$
#
# Copyright (c) 2024, Oracle and/or its affiliates. All rights reserved.
#    NAME
#      provider.tf - Shepherd Data Source file
#
#    USAGE
#      Use the following path for the Example & Backward Compatibility tests: database/db_systems/db_vm/multiple_standby
#    NOTES
#      Terraform Example: TestResourceDbSystemDataGuardAssociation

provider "oci" {
auth                = "SecurityToken"
config_file_profile = "terraform-federation-test"
region              = var.region
tenancy_ocid        = var.compartment_id
}