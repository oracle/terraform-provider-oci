// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
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
  default = "CREATE TABLE IF NOT EXISTS test_mrtable(id INTEGER, name STRING, info JSON, PRIMARY KEY(id)) with schema frozen"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_nosql_table" "test_mrtable" {
  #Required
  compartment_id = var.compartment_ocid
  ddl_statement  = var.table_ddl_statement
  name           = "test_mrtable"

  table_limits {
    #Required
    max_read_units     = "50"
    max_write_units    = "50"
    max_storage_in_gbs = "1"
  }
}

resource "oci_nosql_table_replica" "replica_yul" {
  table_name_or_id = oci_nosql_table.test_mrtable.id
  region = "ca-montreal-1"

  #Optional
  max_read_units     = "60"
  max_write_units    = "60"
}

data "oci_nosql_table" "get_test_mrtable" {
  compartment_id = var.compartment_ocid
  table_name_or_id = oci_nosql_table.test_mrtable.id

  depends_on = [oci_nosql_table_replica.replica_yul]
}

output "test_mrtable_replicas" {
  value = [
    data.oci_nosql_table.get_test_mrtable.replicas
  ]
}
