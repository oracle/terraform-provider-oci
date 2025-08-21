variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "region" {
}

variable "compartment_ocid" {
}

variable "data_safe_adg_target_ocid" {
}

variable "db_system_id" {
}

variable "service_name" {
}

variable "listener_port" {
  default = 1521
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_data_safe_target_database_peer_target_database" "test_target_database_peer_target_database" {
  #Required
  target_database_id = var.data_safe_adg_target_ocid

  database_details {
    database_type       = "DATABASE_CLOUD_SERVICE"
    infrastructure_type = "ORACLE_CLOUD"
    db_system_id        = var.db_system_id
    listener_port       = var.listener_port
    service_name        = var.service_name
  }

  #Optional
  description  = "peer target database description"
  display_name = "peerTargetDatabase1"
}

data "oci_data_safe_target_database_peer_target_databases" "test_target_database_peer_target_databases" {
  target_database_id = var.data_safe_adg_target_ocid
}