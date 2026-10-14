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

variable "config_file_profile" {
  type        = string
  description = "OCI configuration profile containing the legacy SecurityToken session."
}

variable "kms_key_id" {
}

variable "kms_vault_id" {
}

variable "compartment_id" {
}

variable "database_id" {
}

variable "subnet_id" {
}


provider "oci" {
  tenancy_ocid        = var.tenancy_ocid
  user_ocid           = var.user_ocid
  fingerprint         = var.fingerprint
  private_key_path    = var.private_key_path
  auth                = "SecurityToken"
  config_file_profile = var.config_file_profile
  region              = var.region

}

resource "random_string" "autonomous_database_admin_password" {
  length           = 16
  min_numeric      = 2
  min_lower        = 1
  min_upper        = 1
  min_special      = 2
  special          = true
  override_special = "-_#"
}

resource "oci_core_subnet" "test_subnet" {
  cidr_block     = "10.0.0.0/24"
  compartment_id = var.compartment_id
  vcn_id         = oci_core_vcn.test_vcn.id
}

resource "oci_core_vcn" "test_vcn" {
  cidr_block     = "10.0.0.0/16"
  compartment_id = var.compartment_id
}

data "oci_database_migration_migration_object_types" "test_migration_object_types" {
  connection_type = "MYSQL"
}
variable "connection_string" {
  default = ""
}
resource "oci_database_migration_connection" "test_connection_rds_source" {
  compartment_id       = var.compartment_id
  display_name         = "TF_display_test_rds_source"
  connection_type      = "ORACLE"
  key_id               = var.kms_key_id
  vault_id             = var.kms_vault_id
  connection_string    = var.connection_string
  password             = var.test_connection_rds_source_password
  technology_type      = "AMAZON_RDS_ORACLE"
  username             = var.test_connection_rds_source_username
  subnet_id            = var.subnet_id
  replication_password = var.test_connection_rds_source_replication_password
  replication_username = var.test_connection_rds_source_replication_username
}

variable "database_autonomous_id" {
  default = ""
}
resource "oci_database_migration_connection" "test_connection_rds_target" {
  compartment_id       = var.compartment_id
  display_name         = "TF_display_test_rds_target"
  connection_type      = "ORACLE"
  key_id               = var.kms_key_id
  vault_id             = var.kms_vault_id
  database_id          = var.database_autonomous_id
  password             = var.test_connection_rds_target_password
  technology_type      = "OCI_AUTONOMOUS_DATABASE"
  username             = var.test_connection_rds_target_username
  subnet_id            = var.subnet_id
  replication_password = var.test_connection_rds_target_replication_password
  replication_username = var.test_connection_rds_target_replication_username
}

resource "oci_database_migration_connection" "test_connection_target" {
  compartment_id = var.compartment_id
  database_id    = var.database_id
  display_name   = "TF_display_test_create"

  connection_type   = "MYSQL"
  key_id            = var.kms_key_id
  vault_id          = var.kms_vault_id
  password          = var.test_connection_target_password
  technology_type   = "AMAZON_RDS_MYSQL"
  username          = var.test_connection_target_username
  database_name     = var.test_connection_target_database_name
  host              = var.test_connection_target_host
  port              = var.test_connection_target_port
  security_protocol = "PLAIN"
  ssh_host          = var.test_connection_target_ssh_host
  ssh_key           = var.test_connection_target_ssh_key
  ssh_sudo_location = var.test_connection_target_ssh_sudo_location
  ssh_user          = var.test_connection_target_ssh_user
  subnet_id         = var.subnet_id
  wallet            = var.test_connection_target_wallet
}

data "oci_identity_availability_domains" "test_availability_domains" {
  compartment_id = var.compartment_id
}

resource "oci_database_migration_connection" "test_connection_source" {
  compartment_id    = var.compartment_id
  display_name      = "TF_display_test_create_source"
  connection_type   = "MYSQL"
  key_id            = var.kms_key_id
  vault_id          = var.kms_vault_id
  password          = var.test_connection_source_password
  technology_type   = "AMAZON_RDS_MYSQL"
  username          = var.test_connection_source_username
  database_id       = var.database_mysql_id
  database_name     = var.test_connection_source_database_name
  host              = var.test_connection_source_host
  port              = var.test_connection_source_port
  security_protocol = "PLAIN"
  ssh_host          = var.test_connection_source_ssh_host
  ssh_key           = var.test_connection_source_ssh_key
  ssh_sudo_location = var.test_connection_source_ssh_sudo_location
  ssh_user          = var.test_connection_source_ssh_user
  subnet_id         = var.subnet_id
  wallet            = var.test_connection_source_wallet
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
  compartment_id                = var.compartment_id
  database_combination          = "MYSQL"
  source_database_connection_id = var.source_connection_mysql_id
  target_database_connection_id = var.target_connection_mysql_id
  type                          = "ONLINE"
  display_name                  = "displayName"
  data_transfer_medium_details {
    type = "OBJECT_STORAGE"
    object_storage_bucket {
      bucket    = var.bucket_oracle_id
      namespace = var.object_storage_namespace
    }
  }

  initial_load_settings {
    job_mode = "SCHEMA"
  }
}

resource "oci_database_migration_migration" "test_offline_migration" {
  compartment_id                = var.compartment_id
  database_combination          = "MYSQL"
  source_database_connection_id = var.source_connection_mysql_id
  target_database_connection_id = var.target_connection_mysql_id
  type                          = "OFFLINE"
  display_name                  = "displayName"
  data_transfer_medium_details {
    type = "OBJECT_STORAGE"
    object_storage_bucket {
      bucket    = var.bucket_oracle_id
      namespace = var.object_storage_namespace
    }
  }

  initial_load_settings {
    job_mode = "SCHEMA"
  }
}

variable "source_connection_oracle_id" {
  default = ""
}

variable "export_directory_name" {
  description = "Oracle source export directory object name for the legacy Oracle migration."
}

variable "export_directory_path" {
  description = "Oracle source export directory path for the legacy Oracle migration."
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
  compartment_id                          = var.compartment_id
  database_combination                    = "ORACLE"
  source_database_connection_id           = var.source_connection_oracle_id
  source_container_database_connection_id = var.source_connection_container_oracle_id
  target_database_connection_id           = var.target_connection_oracle_id
  data_transfer_medium_details {
    type = "OBJECT_STORAGE"
    object_storage_bucket {
      bucket    = var.bucket_oracle_id
      namespace = var.object_storage_namespace
    }
  }
  initial_load_settings {
    job_mode = "SCHEMA"
    export_directory_object {
      name = var.export_directory_name
      path = var.export_directory_path
    }
  }
  type         = "ONLINE"
  display_name = "displayName"
}

resource "oci_database_migration_migration" "test_oracle_rds_migration" {
  compartment_id                = var.compartment_id
  database_combination          = "ORACLE"
  source_database_connection_id = oci_database_migration_connection.test_connection_rds_source.id
  target_database_connection_id = oci_database_migration_connection.test_connection_rds_target.id

  data_transfer_medium_details {
    type              = "AWS_S3"
    name              = var.rds_s3_bucket
    region            = var.rds_s3_region
    access_key_id     = var.rds_s3_access_key_id
    secret_access_key = var.rds_s3_secret_access_key
    object_storage_bucket {
      bucket    = var.bucket_oracle_id
      namespace = var.object_storage_namespace
    }
  }
  initial_load_settings {
    job_mode = "SCHEMA"
    export_directory_object {
      name = var.rds_export_directory_name
    }
    data_pump_parameters {
      estimate            = "BLOCKS"
      is_cluster          = "false"
      table_exists_action = "TRUNCATE"
    }
  }
  type         = "ONLINE"
  display_name = "displayName"
}

output "password" {
  sensitive = true
  value     = random_string.autonomous_database_admin_password.result
}

variable "source_connection_display_name" {
  description = "Display name of the existing source Database Migration connection to use"
  default     = ""
}

variable "target_connection_display_name" {
  description = "Display name of the existing target Database Migration connection to use"
  default     = ""
}

# Lookup source/target connections using the collection data source
data "oci_database_migration_connections" "ds_source" {
  compartment_id = var.compartment_id
  display_name   = var.source_connection_display_name
}

data "oci_database_migration_connections" "ds_target" {
  compartment_id = var.compartment_id
  display_name   = var.target_connection_display_name
}

# Read and list the migration created by this example.  A migration job is not
# created by the migration resource, so job/report/output lookups belong in
# their own examples and are intentionally not part of this lifecycle example.
data "oci_database_migration_migration" "test_migration" {
  migration_id = oci_database_migration_migration.test_migration.id
}

data "oci_database_migration_jobs" "jobs_for_ds_migration" {
  migration_id = oci_database_migration_migration.test_migration.id
}

variable "test_connection_rds_source_password" {
  type        = string
  description = "password for the legacy test_connection_rds_source connection."
  sensitive   = true
}

variable "test_connection_rds_source_username" {
  type        = string
  description = "username for the legacy test_connection_rds_source connection."
}

variable "test_connection_rds_source_replication_password" {
  type        = string
  description = "replication_password for the legacy test_connection_rds_source connection."
  sensitive   = true
}

variable "test_connection_rds_source_replication_username" {
  type        = string
  description = "replication_username for the legacy test_connection_rds_source connection."
}

variable "test_connection_rds_target_password" {
  type        = string
  description = "password for the legacy test_connection_rds_target connection."
  sensitive   = true
}

variable "test_connection_rds_target_username" {
  type        = string
  description = "username for the legacy test_connection_rds_target connection."
}

variable "test_connection_rds_target_replication_password" {
  type        = string
  description = "replication_password for the legacy test_connection_rds_target connection."
  sensitive   = true
}

variable "test_connection_rds_target_replication_username" {
  type        = string
  description = "replication_username for the legacy test_connection_rds_target connection."
}

variable "test_connection_target_password" {
  type        = string
  description = "password for the legacy test_connection_target connection."
  sensitive   = true
}

variable "test_connection_target_username" {
  type        = string
  description = "username for the legacy test_connection_target connection."
}

variable "test_connection_target_database_name" {
  type        = string
  description = "database_name for the legacy test_connection_target connection."
}

variable "test_connection_target_host" {
  type        = string
  description = "host for the legacy test_connection_target connection."
}

variable "test_connection_target_port" {
  type        = number
  description = "port for the legacy test_connection_target connection."
}

variable "test_connection_target_ssh_host" {
  type        = string
  description = "ssh_host for the legacy test_connection_target connection."
  default     = null
}

variable "test_connection_target_ssh_key" {
  type        = string
  description = "ssh_key for the legacy test_connection_target connection."
  default     = null
  sensitive   = true
}

variable "test_connection_target_ssh_sudo_location" {
  type        = string
  description = "ssh_sudo_location for the legacy test_connection_target connection."
  default     = null
}

variable "test_connection_target_ssh_user" {
  type        = string
  description = "ssh_user for the legacy test_connection_target connection."
  default     = null
}

variable "test_connection_target_wallet" {
  type        = string
  description = "wallet for the legacy test_connection_target connection."
  default     = null
  sensitive   = true
}

variable "test_connection_source_password" {
  type        = string
  description = "password for the legacy test_connection_source connection."
  sensitive   = true
}

variable "test_connection_source_username" {
  type        = string
  description = "username for the legacy test_connection_source connection."
}

variable "test_connection_source_database_name" {
  type        = string
  description = "database_name for the legacy test_connection_source connection."
}

variable "test_connection_source_host" {
  type        = string
  description = "host for the legacy test_connection_source connection."
}

variable "test_connection_source_port" {
  type        = number
  description = "port for the legacy test_connection_source connection."
}

variable "test_connection_source_ssh_host" {
  type        = string
  description = "ssh_host for the legacy test_connection_source connection."
  default     = null
}

variable "test_connection_source_ssh_key" {
  type        = string
  description = "ssh_key for the legacy test_connection_source connection."
  default     = null
  sensitive   = true
}

variable "test_connection_source_ssh_sudo_location" {
  type        = string
  description = "ssh_sudo_location for the legacy test_connection_source connection."
  default     = null
}

variable "test_connection_source_ssh_user" {
  type        = string
  description = "ssh_user for the legacy test_connection_source connection."
  default     = null
}

variable "test_connection_source_wallet" {
  type        = string
  description = "wallet for the legacy test_connection_source connection."
  default     = null
  sensitive   = true
}

variable "object_storage_namespace" {
  type        = string
  description = "Object Storage namespace from the target tenancy."
}

variable "rds_s3_bucket" {
  type        = string
  description = "Existing S3 bucket used for the Oracle RDS migration."
}

variable "rds_s3_region" {
  type        = string
  description = "AWS region containing the Oracle RDS S3 bucket."
}

variable "rds_s3_access_key_id" {
  type        = string
  description = "AWS access key ID authorized for the S3 transfer."
  sensitive   = true
}

variable "rds_s3_secret_access_key" {
  type        = string
  description = "AWS secret access key authorized for the S3 transfer."
  sensitive   = true
}

variable "rds_export_directory_name" {
  type        = string
  description = "Oracle directory object used for the RDS export."
}
