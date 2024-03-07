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

data "oci_database_migration_jobs" "test_jobs" {
  display_name = "displayName"
  filter {
    name = "TF_id"
    values = [
      "jobId"]
  }
  migration_id = "migrationId"
  state = "Succeeded"
}

data "oci_database_migration_job" "test_job" {
  job_id = "jobId"
}

data "oci_database_migration_agent" "test_agent" {
  agent_id = "agentId"
}

data "oci_database_migration_migrations" "test_migrations" {
  #Required
  compartment_id =  var.compartment_id
}

data "oci_database_migration_job_advisor_report" "test_job_advisor_report" {
  job_id = "jobId"
}

data "oci_database_migration_job_output" "test_job_output" {
  job_id = "jobId"
}

data "oci_database_migration_migration_object_types" "test_migration_object_types" {
}

data "oci_database_migration_agent_images" "test_agent_images" {}

resource "oci_database_migration_connection" "test_connection_target" {
  admin_credentials {
    password = random_string.autonomous_database_admin_password.result
    username = "admin"
  }
  compartment_id = var.compartment_id
  database_id = var.database_id
  database_type = "AUTONOMOUS"
  display_name = "TF_display_test_create"
  private_endpoint {
    compartment_id = var.compartment_id
    subnet_id = var.subnet_id
    vcn_id = var.vcn_id
  }
  vault_details {
    compartment_id = var.compartment_id
    key_id = var.kms_key_id
    vault_id = var.kms_vault_id
  }
}

data "oci_identity_availability_domains" "test_availability_domains" {
  compartment_id = var.compartment_id
}

resource "oci_database_migration_connection" "test_connection_source" {
  admin_credentials {
    password = "ORcl##4567890"
    username = "admin"
  }
  compartment_id = var.compartment_id
  connect_descriptor {
    connect_string = "(description=(address=(port=1521)(host=10.2.2.17))(connect_data=(service_name=pdb0107svc.dbsubnet.gghubvcn.oraclevcn.com)))"
  }
  database_type = "MANUAL"
  display_name = "TF_display_test_create_source"
  ssh_details {
    host = "10.2.2.17"
    sshkey = var.ssh_key
    sudo_location = "/usr/bin/sudo"
    user = "opc"
  }
  vault_details {
    compartment_id = var.compartment_id
    key_id = var.kms_key_id
    vault_id = var.kms_vault_id
  }
}

resource "oci_database_migration_connection" "test_connection_source_rds" {
  admin_credentials {
    password = "ORcl##4567890"
    username = "admin"
  }
  compartment_id = var.compartment_id
  connect_descriptor {
    connect_string = "(description=(address=(port=1521)(host=10.2.2.17))(connect_data=(service_name=pdb0107svc.dbsubnet.gghubvcn.oraclevcn.com)))"
  }
  database_type = "MANUAL"
  manual_database_sub_type = "RDS_ORACLE"
  display_name = "TF_display_test_create_source_rds"
  vault_details {
    compartment_id = var.compartment_id
    key_id = var.kms_key_id
    vault_id = var.kms_vault_id
  }
}

resource "oci_database_migration_connection" "test_connection_source_no_ssh" {
  admin_credentials {
    password = "ORcl##4567890"
    username = "admin"
  }
  compartment_id = var.compartment_id
  database_type = "USER_MANAGED_OCI"
  database_id = var.src_database_id
  display_name = "TF_display_test_create_source"

  connect_descriptor {
    connect_string = "(description=(address=(port=1521)(host=10.0.0.42))(connect_data=(service_name=pdb.sub10311806420.vcntesttf.oraclevcn.com)))"
  }
  vault_details {
    compartment_id = var.compartment_id
    key_id = var.kms_key_id
    vault_id = var.kms_vault_id
  }
}

resource "oci_database_migration_connection" "test_connection_target_usr_managed_oci" {
  admin_credentials {
    password = random_string.autonomous_database_admin_password.result
    username = "admin"
  }
  compartment_id = var.compartment_id
  database_type = "USER_MANAGED_OCI"
  database_id = var.tgt_database_id
  display_name = "TF_display_test_create_target"

  connect_descriptor {
    connect_string = "(description=(address=(port=1521)(host=10.0.0.42))(connect_data=(service_name=pdb.sub10311806420.vcntesttf.oraclevcn.com)))"
  }
  vault_details {
    compartment_id = var.compartment_id
    key_id = var.kms_key_id
    vault_id = var.kms_vault_id
  }
}

resource "oci_database_migration_migration" "test_migration" {
  compartment_id = var.compartment_id

  #csvText - Optional
  //csv_text = "MY_BIZZ,SRC_CITY,TABLE,EXCLUDE"
  golden_gate_service_details {
    settings {
      acceptable_lag = "10"
      extract {
        long_trans_duration = "10"
        performance_profile = "LOW"
      }
    }
  }
  data_transfer_medium_details {
    object_storage_details {
      bucket = "bucket"
      namespace = "namespace"
    }
  }
  data_transfer_medium_details_v2 {
    type = "AWS_S3"
    access_key_id = "AKIA4XMMNYWWF4PYZ3EW"
    object_storage_bucket {
      bucket = "bucket"
      namespace = "namespace"
    }
    name = "AWS-S3"
    region = "Ashburn"
    secret_access_key = "3xYJLMQkRDROe7/QzrZTgRDzeq2akfkn+Hb+C95D"
  }
  datapump_settings {
    export_directory_object {
      name = "test_export_dir"
      path = "/u01/app/oracle/product/19.0.0.0/dbhome_1/rdbms/log"
    }
    metadata_remaps {
      new_value = "DATA"
      old_value = "USERS"
      type = "TABLESPACE"
    }
  }
  exclude_objects {
    object = ".*"
    owner  = "owner"
    is_omit_excluded_table_from_replication = "false"
    type = "ALL"
  }
  golden_gate_details {
    hub {
      rest_admin_credentials {
        password = random_string.autonomous_database_admin_password.result
        username = "oggadmin"
      }
      source_container_db_admin_credentials {
        password = random_string.autonomous_database_admin_password.result
        username = "c##ggadmin"
      }
      source_db_admin_credentials {
        password = random_string.autonomous_database_admin_password.result
        username = "ggadmin"
      }
      source_microservices_deployment_name = "Target"
      target_db_admin_credentials {
        password = random_string.autonomous_database_admin_password.result
        username = "ggadmin"
      }
      target_microservices_deployment_name = "Target"
      url = "https://10.0.0.0"
    }
  }
  source_database_connection_id = var.source_connection_id
  source_container_database_connection_id = var.source_connection_container_id
  target_database_connection_id = var.target_connection_id
  type = "ONLINE"
  vault_details {
    compartment_id = var.compartment_id
    key_id = var.kms_key_id
    vault_id = var.kms_vault_id
  }
}

resource "oci_database_migration_migration" "test_migration_rds" {
  compartment_id = var.compartment_id

  golden_gate_service_details {
    settings {
      acceptable_lag = "10"
      extract {
        long_trans_duration = "10"
        performance_profile = "LOW"
      }
    }
  }
  data_transfer_medium_details_v2 {
    type = "OBJECT_STORAGE"
  }
  datapump_settings {
    export_directory_object {
      name = "test_export_dir"
      path = "/u01/app/oracle/product/19.0.0.0/dbhome_1/rdbms/log"
    }
    metadata_remaps {
      new_value = "DATA"
      old_value = "USERS"
      type = "TABLESPACE"
    }
  }
  exclude_objects {
    object = ".*"
    owner  = "owner"
    is_omit_excluded_table_from_replication = "false"
    type = "ALL"
  }
  source_database_connection_id = var.source_connection_rds_id
  target_database_connection_id = var.target_connection_id
  type = "ONLINE"
  vault_details {
    compartment_id = var.compartment_id
    key_id = var.kms_key_id
    vault_id = var.kms_vault_id
  }
}

resource "oci_database_migration_migration" "test_no_ssh_migration" {
  compartment_id = var.compartment_id
  source_database_connection_id = oci_database_migration_connection.test_connection_source_no_ssh.id
  target_database_connection_id = oci_database_migration_connection.test_connection_target_usr_managed_oci.id
  type = "OFFLINE"
  data_transfer_medium_details {
    object_storage_details {
      bucket = var.bucket_id
      namespace = "namespace"
    }
  }
  datapump_settings {
    export_directory_object {
      name = "test_export_dir"
      path = "/u01/app/oracle/product/19.0.0.0/dbhome_1/rdbms/log"
    }
    import_directory_object {
      name = "test_export_dir"
      path = "/u01/app/oracle/product/19.0.0.0/dbhome_1/rdbms/log"
    }
  }
  vault_details {
    compartment_id = var.compartment_id
    key_id = var.kms_key_id
    vault_id = var.kms_vault_id
  }
  dump_transfer_details {
    source {
      kind = "OCI_CLI"
      oci_home = "ociHome"
      wallet_location =  "wallet_location"
    }
    target {
      kind = "OCI_CLI"
      oci_home = "ociHome"
      wallet_location =  "wallet_location"
    }
  }
}

output "password" {
  sensitive = true
  value = random_string.autonomous_database_admin_password.result
}
