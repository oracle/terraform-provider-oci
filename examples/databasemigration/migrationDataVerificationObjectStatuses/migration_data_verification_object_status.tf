// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}
variable "migration_id" {
  description = "OCID of a migration in any state after its initial load has completed."
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_database_migration_migration_data_verification_object_statuses" "test_migration_data_verification_object_statuses" {
  #Required
  compartment_id = var.compartment_id
  migration_id   = var.migration_id

  #Optional
  is_match                                         = null
  migration_data_verification_object_status_filter = null
  name                                             = null
  object_type                                      = null
  owner                                            = null
}

