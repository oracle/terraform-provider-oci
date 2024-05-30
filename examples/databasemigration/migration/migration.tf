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



variable "migration_id" {
  default = ""
}
data "oci_database_migration_migrations" "test_migrations" {
  #Required
  migration_id = var.migration_id
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
variable "connection_string" {
  default = ""
}
variable "nsg_ids" {
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
  nsg_ids = var.nsg_ids
  replication_password="replicationPassword"
  replication_username="replicationUsername"
}

variable "database_autonomous_id" {
  default = ""
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
}

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

resource "oci_database_migration_migration" "test_offline_migration" {
  compartment_id = var.compartment_id
  database_combination = "MYSQL"
  source_database_connection_id = var.source_connection_mysql_id
  target_database_connection_id = var.target_connection_mysql_id
  type = "OFFLINE"
  display_name = "displayName"
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
resource "oci_database_migration_migration" "test_oracle_migration" {
  compartment_id = var.compartment_id
  database_combination = "ORACLE"
  source_database_connection_id = var.source_connection_oracle_id
  source_container_database_connection_id = var.source_connection_container_oracle_id
  target_database_connection_id = var.target_connection_oracle_id
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
  type = "ONLINE"
  display_name = "displayName"
}

output "password" {
  sensitive = true
  value = random_string.autonomous_database_admin_password.result
}