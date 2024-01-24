// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "data_safe_target_ocid" {}
variable "data_safe_db_user_name" {}

variable "sql_collection_compartment_id_in_subtree" {
  default = false
}

variable "sql_collection_access_level" {
  default = "RESTRICTED"
}

variable "sql_collection_display_name" {
  default = "displayName"
}

variable "sql_collection_sql_level" {
  default = "ALL_SQL"
}

variable "sql_collection_status" {
  default = "DISABLED"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_data_safe_sql_collection" "test_sql_collection" {
  #Required
  compartment_id = var.compartment_ocid
  db_user_name = var.data_safe_db_user_name
  target_id = var.data_safe_target_ocid

  #Optional
  display_name = var.sql_collection_display_name
  sql_level = var.sql_collection_sql_level
  status = var.sql_collection_status
}

data "oci_data_safe_sql_collections" "test_sql_collections" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  sql_collection_id = oci_data_safe_sql_collection.test_sql_collection.id
  compartment_id_in_subtree = var.sql_collection_compartment_id_in_subtree
  access_level = var.sql_collection_access_level
}