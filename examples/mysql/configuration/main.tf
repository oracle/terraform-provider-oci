// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.

variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "region" {
  // Define the region where destination backup will be created.
}

variable "compartment_ocid" {
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_mysql_mysql_configuration" "test_mysql_configuration" {
  #Required
  compartment_id = var.compartment_ocid
  shape_name = "MySQL.VM.Standard.E3.1.8GB"

  #Optional
  description = "test configuration created by terraform"
  display_name = "terraform test configuration"
  parent_configuration_id = data.oci_mysql_mysql_configurations.test_mysql_configurations.configurations[0].id
  variables {
    #Optional
    block_encryption_mode                   = "aes-128-ecb"
    binlog_group_commit_sync_delay          = 10
    binlog_group_commit_sync_no_delay_count = 10
    replica_net_timeout                     = 60
    require_secure_transport                = false
    local_infile                            = true
    thread_pool_query_threads_per_group     = 10
    thread_pool_transaction_delay           = 10
    innodb_redo_log_capacity                = 104857600
    explain_format                          = "TRADITIONAL"
    explicit_defaults_for_timestamp         = true
    sql_generate_invisible_primary_key      = false
    temptable_max_ram                       = 1073741824
    innodb_change_buffering                 = "NONE"
    innodb_adaptive_hash_index              = true
    innodb_undo_log_truncate                = true
    table_definition_cache                  = 400
    table_open_cache                        = 4000
    relay_log_space_limit                   = 10
    optimizer_switch                        = "batched_key_access=off"
    replica_type_conversions                = "ALL_LOSSY"
    replica_parallel_workers                = 10
    auto_increment_increment                = 10
    auto_increment_offset                   = 10
    skip_name_resolve                       = false
    max_user_connections                    = 10
    join_buffer_size                        = 262144
    max_seeks_for_key                       = 10
    range_optimizer_max_mem_size            = 8388608
    innodb_autoinc_lock_mode                = 2
    innodb_rollback_on_timeout              = false
    innodb_online_alter_log_max_size        = 134217728
    innodb_sort_buffer_size                 = 1048576
    thread_pool_size                        = 16
    innodb_numa_interleave                  = false
    long_query_time                         = 10
  }
}

data "oci_mysql_mysql_configurations" "test_mysql_configurations" {
  compartment_id = var.compartment_ocid

  #Optional
  state        = "ACTIVE"
  shape_name   = "MySQL.VM.Standard.E3.1.8GB"
}

data "oci_identity_availability_domains" "test_availability_domains" {
  compartment_id = var.tenancy_ocid
}
