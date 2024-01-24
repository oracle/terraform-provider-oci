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
  default = "CREATE TABLE IF NOT EXISTS test_table(id INTEGER, name STRING, age STRING, info JSON, PRIMARY KEY(SHARD(id)))"
}

variable "childtable_ddl_statement" {
  default = "CREATE TABLE IF NOT EXISTS test_table.test_child(idc INTEGER, email STRING, address json, PRIMARY KEY(idc))"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_nosql_table" "test_table" {
  #Required
  compartment_id = var.compartment_ocid
  ddl_statement  = var.table_ddl_statement
  name           = "test_table"

  table_limits {
    #Required
    max_read_units     = "50"
    max_write_units    = "50"
    max_storage_in_gbs = "25"
  }
}

resource "oci_nosql_table" "child_table" {
  #Required
  compartment_id = var.compartment_ocid
  ddl_statement  = var.childtable_ddl_statement
  name           = "test_table.test_child"

  depends_on = [oci_nosql_table.test_table]
}

resource "oci_nosql_index" "child_index" {
  #Required
  keys {
    #Required
    column_name = "email"
  }
  keys {
    column_name = "address"
    json_field_type = "string"
    json_path = "zipCode"
  }

  name             = "idxEmailZipcode"
  table_name_or_id = oci_nosql_table.child_table.id
}

data "oci_nosql_table" "test_table_verify" {
  #Required
  compartment_id = var.compartment_ocid
  table_name_or_id = oci_nosql_table.test_table.name
}

output "table_info" {
  value = [
    data.oci_nosql_table.test_table_verify.id,
    data.oci_nosql_table.test_table_verify.lifecycle_details,
    data.oci_nosql_table.test_table_verify.table_limits,
  ]
}

data "oci_nosql_table" "child_table_verify" {
  #Required
  compartment_id = var.compartment_ocid
  table_name_or_id = oci_nosql_table.child_table.name
}

output "child_table_info" {
  value = [
    data.oci_nosql_table.child_table_verify.id,
    data.oci_nosql_table.child_table_verify.lifecycle_details,
    data.oci_nosql_table.child_table_verify.table_limits,
  ]
}

data "oci_nosql_index" "child_index_verify" {
  #Required
  compartment_id = var.compartment_ocid
  table_name_or_id = oci_nosql_table.child_table.id
  index_name = oci_nosql_index.child_index.name
}

output "child_index_email" {
  value = [
    data.oci_nosql_index.child_index_verify.table_name_or_id,
    data.oci_nosql_index.child_index_verify.index_name,
    data.oci_nosql_index.child_index_verify.keys,
    data.oci_nosql_index.child_index_verify.lifecycle_details,
  ]
}
