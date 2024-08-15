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

variable "compartment_ocid" {
}

variable "kms_key_id" {
}

variable "kms_vault_id" {
}

variable "ssh_public_keys" {
}

variable "compartment_id" {
}

variable "database_id" {
}

variable "subnet_id" {
}

variable "vcn_id" {
}

variable "source_connection_id"{
}

variable "source_connection_container_id"{
}

variable "target_connection_id"{
}

variable "ssh_key" {
}

variable "src_database_id" {
}

variable "tgt_database_id" {
}

variable "bucket_id" {
}

variable "source_connection_rds_id" {
}

variable "connection_string" {
  default = ""
}

variable "nsg_ids" {
  default = ""
}

variable "database_autonomous_id" {
  default = ""
}

variable "migration_id" {
  default = ""
}

variable "source_connection_oracle_id" {
  default = ""
}

variable "source_connection_container_oracle_id" {
  default = ""
}

variable "target_connection_oracle_id" {
  default = ""
}

variable "bucket_oracle_id" {
  default = ""
}

variable "connection_pdb_string" {
  default = ""
}

variable "connection_cdb_string" {
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


data "oci_database_migration_migrations" "test_migrations" {
  #Required
  migration_id = var.migration_id
}

data "oci_database_migration_migration_object_types" "test_migration_object_types" {
  connection_type = "ORACLE"
}

data "oci_identity_availability_domains" "test_availability_domains" {
  compartment_id = var.compartment_id
}

resource "random_string" "autonomous_database_admin_password" {
  length = 16
  min_numeric = 2
  min_lower = 1
  min_upper = 1
  min_special = 2
  special = true
  override_special = "-_#"
}

resource "oci_database_migration_connection" "test_connection_autonomous_target" {
  compartment_id = var.compartment_id
  display_name = "TF_display_test_autonomous_target"
  connection_type = "ORACLE"
  key_id = var.kms_key_id
  vault_id = var.kms_vault_id
  database_id = var.database_autonomous_id
  password = "BEstrO0ng_#11"
  technology_type = "OCI_AUTONOMOUS_DATABASE"
  username = "ggfe"
  replication_password="replicationPassword"
  replication_username="replicationUsername"
}

resource "oci_database_migration_connection" "test_connection_pdb_source" {
  compartment_id = var.compartment_id
  display_name = "TF_display_test_pdb_source"
  connection_type = "ORACLE"
  key_id = var.kms_key_id
  vault_id = var.kms_vault_id
  subnet_id = var.subnet_id
  database_id = var.database_id
  technology_type = "ORACLE_DATABASE"
  username = "ggfe"
  password = "BEstrO0ng_#11"
  replication_password="replicationPassword"
  replication_username="replicationUsername"
}

resource "oci_database_migration_connection" "test_connection_cdb_source" {
  compartment_id = var.compartment_id
  display_name = "TF_display_test_cdb_source"
  connection_type = "ORACLE"
  key_id = var.kms_key_id
  vault_id = var.kms_vault_id
  subnet_id = var.subnet_id
  database_id = var.database_id
  technology_type = "ORACLE_DATABASE"
  username = "ggfe"
  password = "BEstrO0ng_#11"
  replication_password="replicationPassword"
  replication_username="replicationUsername"
}


resource "oci_database_migration_migration" "test_oracle_migration" {
  compartment_id = var.compartment_id
  database_combination = "ORACLE"
  source_database_connection_id = oci_database_migration_connection.test_connection_pdb_source.id
  source_container_database_connection_id = oci_database_migration_connection.test_connection_cdb_source.id
  target_database_connection_id = oci_database_migration_connection.test_connection_autonomous_target.id
  advanced_parameters {
    data_type = "STRING"
    name = "DATAPUMPSETTINGS_METADATAONLY"
    value = "True"
  }
  data_transfer_medium_details {
    type = "OBJECT_STORAGE"
    object_storage_bucket {
      bucket = var.bucket_oracle_id
      namespace = "namespace"
    }
  }
  type = "ONLINE"
  display_name = "displayName"
}


output "password" {
  sensitive = true
  value = random_string.autonomous_database_admin_password.result
}