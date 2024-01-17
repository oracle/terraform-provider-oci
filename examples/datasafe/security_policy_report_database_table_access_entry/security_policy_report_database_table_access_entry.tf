// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "security_policy_report_ocid" {}
variable "data_safe_target_ocid" {}
variable "database_table_access_entrykey" {}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_data_safe_security_policy_report_database_table_access_entries" "test_security_policy_report_database_table_access_entries" {
  #Required
  security_policy_report_id = var.security_policy_report_ocid
}

data "oci_data_safe_security_policy_report_database_table_access_entry" "test_security_policy_report_database_table_access_entry" {
  #Required
  security_policy_report_id = var.security_policy_report_ocid
  database_table_access_entry_key = var.database_table_access_entrykey
}
