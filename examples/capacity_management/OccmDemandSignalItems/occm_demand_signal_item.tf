// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "sp_compartment_ocid" {}

variable "occm_demand_signal_item_availability_domain" {
  default = "us-ashburn-1-ad-1"
}

variable "occm_demand_signal_item_demand_quantity" {
  default = 10
}

variable "occm_demand_signal_item_demand_signal_namespace" {
  default = "COMPUTE"
}

variable "occm_demand_signal_item_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "occm_demand_signal_item_notes" {
  default = "notes"
}

variable "occm_demand_signal_item_region" {
  default = "us-ashburn-1"
}

variable "occm_demand_signal_item_request_type" {
  default = "DEMAND"
}

variable "occm_demand_signal_resource_properties_items" {
  description = "A map of resource property key-value pairs for the demand signal item."
  type        = map(string) # Explicitly enforce map of string type
  default     = {
    "ocpu"   = "1"
    "memory" = "13"
    # You can add other common properties here
    # "nvme" = "NA"
  }
}

variable "occm_demand_signal_item_time_needed_before" {
  default = "2025-06-30T23:59:59Z"
}

variable "occm_demand_signal_id" {}

variable "demand_signal_catalog_resource_id" {}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_capacity_management_occm_demand_signal_item" "test_occm_demand_signal_item" {
  #Required
  compartment_id                    = var.sp_compartment_ocid
  demand_quantity                   = var.occm_demand_signal_item_demand_quantity
  demand_signal_catalog_resource_id = var.demand_signal_catalog_resource_id
  demand_signal_id                  = var.occm_demand_signal_id
  region                            = var.occm_demand_signal_item_region
  request_type                      = var.occm_demand_signal_item_request_type
  resource_properties               = var.occm_demand_signal_resource_properties_items
  time_needed_before                = var.occm_demand_signal_item_time_needed_before

  #Optional
  availability_domain   = var.occm_demand_signal_item_availability_domain
  freeform_tags         = var.occm_demand_signal_item_freeform_tags
  notes                 = var.occm_demand_signal_item_notes
  target_compartment_id = var.sp_compartment_ocid
}

data "oci_capacity_management_occm_demand_signal_items" "test_occm_demand_signal_items" {
  #Required
  compartment_id = var.sp_compartment_ocid

  #Optional
  demand_signal_namespace = var.occm_demand_signal_item_demand_signal_namespace
  occm_demand_signal_id   = var.occm_demand_signal_id
}

