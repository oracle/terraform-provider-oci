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

variable "table_ddl_statement" {
  default = "CREATE TABLE IF NOT EXISTS test_ondemand(id INTEGER, name STRING, age STRING, PRIMARY KEY(SHARD(id)))"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_nosql_table" "test_ondemand" {
  #Required
  compartment_id = var.compartment_ocid
  ddl_statement  = var.table_ddl_statement
  name           = "test_ondemand"

  table_limits {
    #Required
    max_read_units     = "0"
    max_write_units    = "0"
    max_storage_in_gbs = "25"
    capacity_mode      = "ON_DEMAND"
  }
}

data "oci_nosql_table" "test_ondemand_verify" {
  #Required
  compartment_id = var.compartment_ocid
  table_name_or_id = oci_nosql_table.test_ondemand.name
}

output "ondemand_table_info" {
  value = [
    data.oci_nosql_table.test_ondemand_verify.id,
    data.oci_nosql_table.test_ondemand_verify.lifecycle_details,
    data.oci_nosql_table.test_ondemand_verify.table_limits,
  ]
}