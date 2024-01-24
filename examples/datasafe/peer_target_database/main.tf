// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

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

variable "data_safe_target_ocid" {}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_data_safe_data_safe_configuration" "test_data_safe_configuration" {
  is_enabled = "true"
}

variable "peer_target_database_description" {
  default = "description"
}

variable "peer_target_database_display_name" {
  default = "peerTargetDatabase1"
}

variable "peerdb_ocid" {}

variable "peerdb_port" {}

variable "service_name" {}


resource "random_string" "autonomous_database_admin_password" {
  length      = 16
  min_numeric = 1
  min_lower   = 1
  min_upper   = 1
  min_special = 1
}
variable "autonomous_database_db_workload" {
  default = "OLTP"
}

variable "autonomous_database_freeform_tags" {
  default = {
    "Department" = "Finance"
  }
}

resource "oci_data_safe_target_database_peer_target_database" "test_target_database_peer_target_database" {
#Required
    target_database_id = var.data_safe_target_ocid
    database_details {
        database_type = "DATABASE_CLOUD_SERVICE"
        infrastructure_type = "ORACLE_CLOUD"
        db_system_id = var.peerdb_ocid
        listener_port = var.peerdb_port
        service_name = var.service_name
    }
}