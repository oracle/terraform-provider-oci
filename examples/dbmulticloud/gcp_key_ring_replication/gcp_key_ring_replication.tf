// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}

variable "oracle_db_gcp_connector_id" {
 type = string 
}

variable "oracle_db_gcp_key_ring_id" {
  type = string
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_dbmulticloud_oracle_db_gcp_key_ring" "existing" {
  compartment_id         = var.compartment_ocid
  display_name           = "Tersi_Example"
  oracle_db_connector_id = var.oracle_db_gcp_connector_id
    // Replication
  action        = "DELETE"
  target_region = "us-boardman-1"
}

data "oci_dbmulticloud_oracle_db_gcp_key_ring" "existing" {
   oracle_db_gcp_key_ring_id  = var.oracle_db_gcp_key_ring_id
}

output "gcp_key_ring_id" {
  value = oci_dbmulticloud_oracle_db_gcp_key_ring.existing.id
}

