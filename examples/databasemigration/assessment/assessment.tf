// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "region" {
}

variable "private_key_path" {
}

variable "compartment_id" {
}

variable "source_connection_oracle_id" {
  default = ""
}

variable "target_connection_oracle_id" {
  default = ""
}


provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  auth             = "SecurityToken"
  region           = var.region

}


resource "oci_database_migration_assessment" "test_assessment" {
  compartment_id                    = var.compartment_id
  source_database_connection {
    id = var.source_connection_oracle_id
  }
  target_database_connection {
    id = var.target_connection_oracle_id
  }
  display_name                      = "TF_display_test_rds_source"
  acceptable_downtime               = "LESS_THAN_10_MINUTES"
  database_combination              = "ORACLE"
  database_data_size                = "LESS_THAN_1GB"
  ddl_expectation                   = "DDL_EXPECTED"
  network_speed_megabit_per_second  = "MBPS_10"
  creation_type                     = "CREATE_ONLY"
  description                       = "description"
  exclude_objects {
    object = ".*"
    is_omit_excluded_table_from_replication = false
    owner = "owner"
    schema = "schema"
    type = "ALL"
  }

}
