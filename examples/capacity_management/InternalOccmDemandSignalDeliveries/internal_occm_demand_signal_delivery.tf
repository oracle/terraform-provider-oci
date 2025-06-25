// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "prod_compartment_ocid" {}

variable "internal_occm_demand_signal_delivery_accepted_quantity" {
  default = 10
}

variable "internal_occm_demand_signal_delivery_defined_tags_value" {
  default = "value"
}

variable "internal_occm_demand_signal_delivery_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "internal_occm_demand_signal_delivery_id" {
  default = "id"
}

variable "internal_occm_demand_signal_delivery_justification" {
  default = "justification"
}

variable "internal_occm_demand_signal_delivery_notes" {
  default = "notes"
}

variable "inprogress_demand_signal_id" {
}

variable "demand_signal_item_id" {}

variable "customergroup_id" {}


provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_capacity_management_internal_occm_demand_signal_delivery" "test_internal_occm_demand_signal_delivery" {
  #Required
  accepted_quantity     = var.internal_occm_demand_signal_delivery_accepted_quantity
  compartment_id        = var.prod_compartment_ocid
  demand_signal_id      = var.inprogress_demand_signal_id
  demand_signal_item_id = var.demand_signal_item_id
  occ_customer_group_id = var.customergroup_id

  #Optional
  freeform_tags = var.internal_occm_demand_signal_delivery_freeform_tags
  justification = var.internal_occm_demand_signal_delivery_justification
  notes         = var.internal_occm_demand_signal_delivery_notes
}

data "oci_capacity_management_internal_occm_demand_signal_deliveries" "test_internal_occm_demand_signal_deliveries" {
  #Required
  compartment_id        = var.prod_compartment_ocid
  occ_customer_group_id = var.customergroup_id

  #Optional
  occm_demand_signal_item_id = var.demand_signal_item_id
}

