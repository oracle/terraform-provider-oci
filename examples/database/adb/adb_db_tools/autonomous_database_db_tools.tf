// Copyright (c) 2026, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

// Example for creating an Autonomous JSON Database with DB Tools configured.
resource "random_string" "autonomous_database_admin_password" {
  length      = 16
  min_numeric = 1
  min_lower   = 1
  min_upper   = 1
  min_special = 1
}


resource "oci_database_autonomous_database" "test_autonomous_database_db_tools_ecpu" {
  admin_password                       = random_string.autonomous_database_admin_password.result
  compartment_id                       = var.compartment_ocid
  compute_count                        = "8.0"
  compute_model                        = "ECPU"
  data_storage_size_in_tbs             = "1"
  backup_retention_period_in_days      = "15"
  db_name                              = "adbdbtoolsEcpuOltp6"
  db_version                           = "19c"
  db_workload                          = "OLTP"
  license_model                        = "LICENSE_INCLUDED"
  is_free_tier                         = "false"
  db_tools_details {
                name = "APEX"
                is_enabled = false
                compute_count = 0
                max_idle_time_in_minutes = 0
              }
              db_tools_details {
                name = "DATA_TRANSFORMS"
                is_enabled = false
                compute_count = 0
                max_idle_time_in_minutes = 0
              }
              db_tools_details {
                name = "DATABASE_ACTIONS"
                is_enabled = false
                compute_count = 0
                max_idle_time_in_minutes = 0
              }
              db_tools_details {
                name = "GRAPH_STUDIO"
                is_enabled = true
                compute_count = 10
                max_idle_time_in_minutes = 120
              }
              db_tools_details {
                name = "MONGODB_API"
                is_enabled = false
                compute_count = 0
                max_idle_time_in_minutes = 0
              }
              db_tools_details {
                name = "OML"
                is_enabled = false
                compute_count = 0
                max_idle_time_in_minutes = 0
              }
              db_tools_details {
                name = "ORDS"
                is_enabled = false
                compute_count = 0
                max_idle_time_in_minutes = 0
              }
              db_tools_details {
                name = "SPATIAL_STUDIO"
                is_enabled = true
                compute_count = 8
                max_idle_time_in_minutes = 60
              }
              db_tools_details {
                  name = "DLA"
                  is_enabled = false
                  compute_count = 0
                  max_idle_time_in_minutes = 0
               }
               db_tools_details {
                  name = "ESSBASE"
                  is_enabled = false
                }
  }

data "oci_database_autonomous_databases_clones" "test_autonomous_databases_clones" {
  #Required
  autonomous_database_id = oci_database_autonomous_database.test_autonomous_database_db_tools_ecpu.id
  compartment_id = var.compartment_ocid

  #Optional
  clone_type = "REFRESHABLE_CLONE"
}



