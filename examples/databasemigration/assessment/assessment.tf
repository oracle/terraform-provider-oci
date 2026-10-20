// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "region" {
}

variable "config_file_profile" {
}

variable "compartment_ocid" {
}

variable "source_connection_oracle_id" {
  default = ""
}

variable "target_connection_oracle_id" {
  default = ""
}


provider "oci" {
  auth                = "SecurityToken"
  config_file_profile = var.config_file_profile
  region              = var.region
}


resource "oci_database_migration_assessment" "test_assessment" {
  compartment_id = var.compartment_ocid
  source_database_connection {
    id = var.source_connection_oracle_id
  }
  target_database_connection {
    id = var.target_connection_oracle_id
  }
  display_name                     = "TF_display_test_rds_source"
  migration_scope                  = "SCHEMA"
  acceptable_downtime              = "LESS_THAN_10_MINUTES"
  database_combination             = "ORACLE"
  database_data_size               = "LESS_THAN_1GB"
  ddl_expectation                  = "DDL_EXPECTED"
  network_speed_megabit_per_second = "MBPS_10"
  creation_type                    = "CREATE_ONLY"
  description                      = "description"
  exclude_objects {
    object                                  = ".*"
    is_omit_excluded_table_from_replication = false
    owner                                   = "owner"
    schema                                  = "schema"
    type                                    = "ALL"
  }

}
