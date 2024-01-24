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

variable "index_keys_column_name" {
  default = "name"
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

resource "oci_nosql_index" "test_index" {
  #Required
  keys {
    #Required
    column_name = var.index_keys_column_name
  }

  name             = "test_index"
  table_name_or_id = oci_nosql_table.test_table.id
}

data "oci_nosql_tables" "test_tables" {
  #Required
  compartment_id = var.compartment_ocid

  filter {
    name   = "id"
    values = [oci_nosql_table.test_table.id]
  }
}

output "table_name" {
  value = [
    data.oci_nosql_tables.test_tables.table_collection[0].name,
  ]
}

data "oci_nosql_indexes" "test_indexes" {
  #Required
  table_name_or_id = oci_nosql_table.test_table.id

  filter {
    name   = "name"
    values = [oci_nosql_index.test_index.name]
  }
}

output "index_name" {
  value = [
    data.oci_nosql_indexes.test_indexes.index_collection[0].name,
  ]
}

