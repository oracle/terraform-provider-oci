# $Header$
#
# Copyright (c) 2024, Oracle and/or its affiliates. All rights reserved.
#    NAME
#      datasources.tf - Shepherd Data Source file
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
#    escabrer   10/14/2024 - Created


data "oci_identity_availability_domains" "test_availability_domains" {
  compartment_id = var.tenancy_ocid
}