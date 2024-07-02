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


provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  auth             = "SecurityToken"
  region           = var.region

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

resource "oci_core_subnet" "test_subnet" {
  cidr_block = "10.0.0.0/24"
  compartment_id = var.compartment_id
  vcn_id = oci_core_vcn.test_vcn.id
}

resource "oci_core_vcn" "test_vcn" {
  cidr_block = "10.0.0.0/16"
  compartment_id = var.compartment_id
}

variable "jobId" {
  default = ""
}
data "oci_database_migration_job" "test_job" {
  job_id = var.jobId
}

data "oci_database_migration_agent" "test_agent" {
  agent_id = "agentId"
}

data "oci_database_migration_migrations" "test_migrations" {
  #Required
  compartment_id =  var.compartment_id
}

data "oci_database_migration_job_advisor_report" "test_job_advisor_report" {
  job_id = var.jobId
}

data "oci_database_migration_job_output" "test_job_output" {
  job_id = var.jobId
}

data "oci_database_migration_migration_object_types" "test_migration_object_types" {
  connection_type = "MYSQL"
}

data "oci_database_migration_agent_images" "test_agent_images" {}

resource "oci_database_migration_connection" "test_connection_target" {
  compartment_id = var.compartment_id
  database_id = var.database_id
  display_name = "TF_display_test_create"

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

data "oci_identity_availability_domains" "test_availability_domains" {
  compartment_id = var.compartment_id
}

resource "oci_database_migration_connection" "test_connection_source" {
  compartment_id = var.compartment_id
  display_name = "TF_display_test_create_source"
  connection_type = "MYSQL"
  key_id = var.kms_key_id
  vault_id = var.kms_vault_id
  password = "BEstrO0ng_#11"
  technology_type = "AMAZON_RDS_MYSQL"
  username = "ggfe"
  database_id = var.database_mysql_id
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

variable "database_mysql_id" {
  default = ""
}

variable "source_connection_mysql_id" {
  default = ""
}
variable "target_connection_mysql_id" {
  default = ""
}
resource "oci_database_migration_migration" "test_migration" {
  compartment_id = var.compartment_id
  database_combination = "MYSQL"
  source_database_connection_id = var.source_connection_mysql_id
  target_database_connection_id = var.target_connection_mysql_id
  type = "ONLINE"
  display_name = "displayName"
}

output "password" {
  sensitive = true
  value = random_string.autonomous_database_admin_password.result
}