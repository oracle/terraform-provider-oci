// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_nosql_table" "table_rd" {
  #Required
  compartment_id = "${var.compartment_ocid}"
  ddl_statement  = "${var.table_ddl_statement}"
  name           = "test_table"

  table_limits {
    #Required
    max_read_units     = "10"
    max_storage_in_gbs = "10"
    max_write_units    = "10"
  }
}

resource "oci_nosql_index" "index_rd" {
  #Required
  keys {
    #Required
    column_name = "${var.index_keys_column_name}"
  }

  name             = "test_index"
  table_name_or_id = "${oci_nosql_table.table_rd.id}"
}
