// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "database_mysql_hatewave_id" {
  default = ""
}

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

variable "migration_id" {
  default = ""
}

variable "database_autonomous_id" {
  default = ""
}

variable "database_target_mysql_id" {
  default = ""
}

variable "database_mysql_id" {
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
  connection_type = "MYSQL"
}

data "oci_identity_availability_domains" "test_availability_domains" {
  compartment_id = var.compartment_id
}

resource "oci_database_migration_connection" "test_connection_mysql_hatewave_target" {
  compartment_id = var.compartment_id
  database_id = var.database_mysql_hatewave_id
  display_name = "TF_display_test_create"

  connection_type = "MYSQL"
  key_id = var.kms_key_id
  vault_id = var.kms_vault_id
  password = "BEstrO0ng_#11"
  technology_type = "OCI_MYSQL"
  username = "ggfe"
  database_name = "ggfe"
  host = "254.249.0.0"
  port = "3306"
  replication_password="replicationPassword"
  replication_username="replicationUsername"
  security_protocol="PLAIN"
  ssh_host =   "sshHost"
  ssh_key = "sshKey"
  ssh_sudo_location = "sshSudoLocation"
  ssh_user = "sshUser"
  subnet_id = var.subnet_id
  wallet =  "wallet2"

}

resource "oci_database_migration_connection" "test_connection_mysql_server_source" {
  compartment_id = var.compartment_id
  display_name = "TF_display_test_create_source"
  connection_type = "MYSQL"
  key_id = var.kms_key_id
  vault_id = var.kms_vault_id
  password = "BEstrO0ng_#11"
  technology_type = "MYSQL_SERVER"
  username = "ggfe"
  database_name = "ggfe"
  host = "254.249.0.0"
  port = "3306"
  replication_password="replicationPassword"
  replication_username="replicationUsername"
  security_protocol="PLAIN"
  ssh_host =   "sshHost"
  ssh_key = "sshKey"
  ssh_sudo_location = "sshSudoLocation"
  ssh_user = "sshUser"
  wallet =  "wallet2"

}

resource "oci_database_migration_migration" "test_mysql_migration" {
  compartment_id = var.compartment_id
  database_combination = "MYSQL"
  source_database_connection_id = oci_database_migration_connection.test_connection_mysql_server_source.id
  target_database_connection_id = oci_database_migration_connection.test_connection_mysql_hatewave_target.id
  type = "ONLINE"
  display_name = "displayName"
}

resource "oci_database_migration_migration" "test_offline_mysql_migration" {
  compartment_id = var.compartment_id
  database_combination = "MYSQL"
  source_database_connection_id = oci_database_migration_connection.test_connection_mysql_server_source.id
  target_database_connection_id = oci_database_migration_connection.test_connection_mysql_hatewave_target.id
  type = "OFFLINE"
  display_name = "displayName"
}
