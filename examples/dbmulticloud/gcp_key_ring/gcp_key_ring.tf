// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

# variable "tenancy_ocid" {}
# variable "user_ocid" {}
# variable "fingerprint" {}
# variable "private_key_path" {}
# variable "region" {}
variable "compartment_id" {}

variable "oracle_db_gcp_oracle_db_connector_id" {
}
variable "oracle_db_gcp_key_ring_display_name" {
  default = "TestDbGcpVaultExample"
}

variable "oracle_db_gcp_key_ring_location" {
  default = "location"
}

variable "oracle_db_gcp_key_ring_properties" {
  default = ""
}

variable "oracle_db_gcp_key_ring_state" {
  default = ""
}

variable "oracle_db_gcp_key_ring_type" {
  default = "type"
}

variable "gcp_key_ring_id" {
  default = "test key ring"
}

# provider "oci" {
#   tenancy_ocid     = var.tenancy_ocid
#   user_ocid        = var.user_ocid
#   fingerprint      = var.fingerprint
#   private_key_path = var.private_key_path
#   region           = var.region
# }

resource "oci_dbmulticloud_oracle_db_gcp_key_ring" "test_oracle_db_gcp_key_ring" {
  #Required
  compartment_id         = var.compartment_id
  display_name           = var.oracle_db_gcp_key_ring_display_name
  oracle_db_connector_id = var.oracle_db_gcp_oracle_db_connector_id

  gcp_key_ring_id = var.gcp_key_ring_id
  location        = var.oracle_db_gcp_key_ring_location
#   properties      = var.oracle_db_gcp_key_ring_properties
  type            = var.oracle_db_gcp_key_ring_type
}

data "oci_dbmulticloud_oracle_db_gcp_key_rings" "test_oracle_db_gcp_key_rings" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name               = var.oracle_db_gcp_key_ring_display_name
  oracle_db_gcp_connector_id = var.oracle_db_gcp_oracle_db_connector_id
  oracle_db_gcp_key_ring_id  = var.gcp_key_ring_id
#   state                      = var.oracle_db_gcp_key_ring_state
}

output "oracle_db_gcp_key_ring_id" {
  value = oci_dbmulticloud_oracle_db_gcp_key_ring.test_oracle_db_gcp_key_ring.id
}
