// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

variable "occ_handover_resource_block_id" {
  default = "handoverResourceBlockId"
}

variable "occ_handover_resource_host_id" {
  default = "hostId"
}


provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_capacity_management_occ_handover_resource_block_details" "test_occ_handover_resource_block_details" {
  #Required
  occ_handover_resource_block_id = var.occ_handover_resource_block_id

  #Optional
  host_id = var.occ_handover_resource_host_id
}
