# $Header$
#
# Copyright (c) 2026, Oracle and/or its affiliates. All rights reserved.
#    USAGE
#      Example & Backward Compatibility Path: database/db_systems/db_vm/db_vm_std_x86_dbrs_tags

provider "oci" {
  auth                = "SecurityToken"
  config_file_profile = "terraform-federation-test"
  region              = var.region
  tenancy_ocid        = var.compartment_id
}
