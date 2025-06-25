// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "sp_compartment_ocid" {}

variable "occm_demand_signal_delivery_id" {
  default = "id"
}

variable "demand_signal_item_id" {}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_capacity_management_occm_demand_signal_deliveries" "test_occm_demand_signal_deliveries" {
  #Required
  compartment_id = var.sp_compartment_ocid

  #Optional
  id                         = var.occm_demand_signal_delivery_id
  occm_demand_signal_item_id = var.demand_signal_item_id
}

