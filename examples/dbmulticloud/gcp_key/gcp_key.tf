// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

# variable "tenancy_ocid" {}
# variable "user_ocid" {}
# variable "fingerprint" {}
# variable "private_key_path" {}
# variable "region" {}
variable "compartment_id" {}
variable "oracle_db_gcp_key_ring_id" {
}
variable "oracle_db_gcp_key_display_name" {
  default = "displayName"
}

variable "oracle_db_gcp_key_id" {
  default = ""
}



# provider "oci" {
#   tenancy_ocid     = var.tenancy_ocid
#   user_ocid        = var.user_ocid
#   fingerprint      = var.fingerprint
#   private_key_path = var.private_key_path
#   region           = var.region
# }

data "oci_dbmulticloud_oracle_db_gcp_keys" "test_oracle_db_gcp_keys" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  #oracle_db_gcp_key_id      = var.oracle_db_gcp_key_id
  display_name              = var.oracle_db_gcp_key_display_name
  oracle_db_gcp_key_ring_id = var.oracle_db_gcp_key_ring_id
#   state                     = var.oracle_db_gcp_key_state
}

