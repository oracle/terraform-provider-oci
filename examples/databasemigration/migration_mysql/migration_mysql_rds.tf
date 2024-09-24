// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0



resource "oci_database_migration_connection" "test_connection_mysql_rds_target" {
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

resource "oci_database_migration_connection" "test_connection_mysql_rds_source" {
  compartment_id = var.compartment_id
  display_name = "TF_display_test_create_source"
  connection_type = "MYSQL"
  key_id = var.kms_key_id
  vault_id = var.kms_vault_id
  password = "BEstrO0ng_#11"
  technology_type = "AMAZON_RDS_MYSQL"
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

resource "oci_database_migration_migration" "test_mysql_rds_migration" {
  compartment_id = var.compartment_id
  database_combination = "MYSQL"
  source_database_connection_id = oci_database_migration_connection.test_connection_mysql_rds_source.id
  target_database_connection_id = oci_database_migration_connection.test_connection_mysql_rds_target.id
  type = "ONLINE"
  display_name = "displayName"
}

resource "oci_database_migration_migration" "test_offline_mysql_rds_migration" {
  compartment_id = var.compartment_id
  database_combination = "MYSQL"
  source_database_connection_id = oci_database_migration_connection.test_connection_mysql_rds_source.id
  target_database_connection_id = oci_database_migration_connection.test_connection_mysql_rds_target.id
  type = "OFFLINE"
  display_name = "displayName"
}