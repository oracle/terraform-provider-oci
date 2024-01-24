// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "region" {
}

variable "compartment_ocid" {
}

variable "autonomous_database_db_workload" {
  default = "OLTP"
}

variable "autonomous_database_defined_tags_value" {
  default = "value"
}

variable "autonomous_database_freeform_tags" {
  default = {
    "Department" = "Finance"
  }
}

variable "autonomous_database_license_model" {
  default = "LICENSE_INCLUDED"
}

variable "autonomous_database_is_dedicated" {
  default = false
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "random_string" "autonomous_database_admin_password" {
  length      = 16
  min_numeric = 1
  min_lower   = 1
  min_upper   = 1
  min_special = 1
}

resource "oci_database_autonomous_database" "autonomous_database" {
  #Required
  admin_password           = random_string.autonomous_database_admin_password.result
  compartment_id           = var.compartment_ocid
  cpu_core_count           = "1"
  data_storage_size_in_tbs = "1"
  db_name                  = "adbdb1"

  #Optional
  db_workload                                    = var.autonomous_database_db_workload
  display_name                                   = "example_autonomous_database"
  freeform_tags                                  = var.autonomous_database_freeform_tags
  is_auto_scaling_enabled                        = "true"
  license_model                                  = var.autonomous_database_license_model
  is_preview_version_with_service_terms_accepted = "false"
}

data "oci_database_autonomous_databases" "autonomous_databases" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  display_name = oci_database_autonomous_database.autonomous_database.display_name
  db_workload  = var.autonomous_database_db_workload
}

resource "oci_database_autonomous_database_instance_wallet_management" "test_autonomous_database_instance_wallet_management" {
  autonomous_database_id = oci_database_autonomous_database.autonomous_database.id
}

data "oci_database_autonomous_database_instance_wallet_management" "test_autonomous_database_instance_wallet_management" {
  #Required
  autonomous_database_id = oci_database_autonomous_database.autonomous_database.id
}

output "autonomous_database_admin_password" {
  value = random_string.autonomous_database_admin_password.result
}

output "autonomous_database_high_connection_string" {
  value = lookup(
    oci_database_autonomous_database.autonomous_database.connection_strings[0].all_connection_strings,
    "high",
    "unavailable",
  )
}

output "autonomous_databases" {
  value = data.oci_database_autonomous_databases.autonomous_databases.autonomous_databases
}

