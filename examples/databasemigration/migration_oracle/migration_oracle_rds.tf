// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "nsg_mysql_id" {
  default = ""
}
variable "nsg_mysql_id2" {
  default = ""
}

resource "oci_database_migration_connection" "test_connection_rds_source" {
  compartment_id = var.compartment_id
  display_name = "TF_display_test_rds_source"
  connection_type = "ORACLE"
  key_id = var.kms_key_id
  vault_id = var.kms_vault_id
  connection_string = var.connection_string
  password = "BEstrO0ng_#11"
  technology_type = "AMAZON_RDS_ORACLE"
  username = "ggfe"
  nsg_ids = [var.nsg_mysql_id,var.nsg_mysql_id2]
  replication_password="replicationPassword"
  replication_username="replicationUsername"
}

resource "oci_database_migration_connection" "test_connection_rds_target" {
  compartment_id = var.compartment_id
  display_name = "TF_display_test_rds_target"
  connection_type = "ORACLE"
  key_id = var.kms_key_id
  vault_id = var.kms_vault_id
  database_id = var.database_autonomous_id
  password = "BEstrO0ng_#11"
  technology_type = "OCI_AUTONOMOUS_DATABASE"
  username = "ggfe"
  replication_password="replicationPassword"
  replication_username="replicationUsername"
  subnet_id = var.subnet_id
}

resource "oci_database_migration_migration" "test_oracle_rds_migration" {
  compartment_id = var.compartment_id
  database_combination = "ORACLE"
  source_database_connection_id = oci_database_migration_connection.test_connection_rds_source.id
  target_database_connection_id = oci_database_migration_connection.test_connection_rds_target.id

  data_transfer_medium_details {
    type = "AWS_S3"
    name = "rdsbucket"
    region = "us-east-1"
    secret_access_key = "12345/12345"
    access_key_id = "12345"
    object_storage_bucket {
      bucket = var.bucket_oracle_id
      namespace = "namespace"
    }
  }
  initial_load_settings {
    job_mode = "SCHEMA"
    export_directory_object {
      name = "name"
    }
    data_pump_parameters {
      estimate = "BLOCKS"
      is_cluster = "false"
      table_exists_action = "TRUNCATE"
    }
  }
  type = "ONLINE"
  display_name = "displayName"
}