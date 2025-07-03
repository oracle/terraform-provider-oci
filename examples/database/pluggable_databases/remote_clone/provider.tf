# $Header$
#
# Copyright (c) 2025, Oracle and/or its affiliates. All rights reserved.
#    NAME
#      provider.tf - provider file
#
#    USAGE
#      Use the following path for Example Test & Backward-Compatibility-Test: database/pluggable_databases/remote_clone
#    NOTES
#      Terraform Example:
#    FILES
#
#    DESCRIPTION
#
#    MODIFIED   MM/DD/YY
#    escabrer   05/08/2025 - Created


provider "oci" {
    auth                = "SecurityToken"
    config_file_profile = "terraform-federation-test"
    region              = var.region
    tenancy_ocid        = var.compartment_id
}