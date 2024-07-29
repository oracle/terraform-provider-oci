// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "internal_occ_handover_resource_block_handover_date_greater_than_or_equal_to" {
  default = "2023-08-05T17:17:14.816Z"
}

variable "internal_occ_handover_resource_block_handover_date_less_than_or_equal_to" {
  default = "2024-08-05T17:17:14.816Z"
}

variable "internal_occ_handover_resource_block_namespace" {
  default = "COMPUTE"
}

variable "occ_customer_group_id" {
  default = "customerGroupId"
}

variable "occ_handover_resource_block_id" {
  default = "handoverResourceBlockId"
}

variable "occ_handover_resource_resource_name" {
  default = "handoverResource"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_capacity_management_internal_occ_handover_resource_blocks" "test_internal_occ_handover_resource_blocks" {
  #Required
  compartment_id        = var.compartment_id
  namespace             = var.internal_occ_handover_resource_block_namespace
  occ_customer_group_id = var.occ_customer_group_id

  #Optional
  handover_date_greater_than_or_equal_to = var.internal_occ_handover_resource_block_handover_date_greater_than_or_equal_to
  handover_date_less_than_or_equal_to    = var.internal_occ_handover_resource_block_handover_date_less_than_or_equal_to
  handover_resource_name                 = var.occ_handover_resource_resource_name
  occ_handover_resource_block_id         = var.occ_handover_resource_block_id
}
