// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0


variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "sp_compartment_ocid" {}

variable "occm_demand_signal_description" {
  default = "description"
}

variable "occm_demand_signal_display_name" {
  default = "displayName"
}

variable "occm_demand_signal_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "occm_demand_signal_id" {
  default = "id"
}

variable "occm_demand_signal_lifecycle_details" {
  default = "lifecycleDetails"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_capacity_management_occm_demand_signal" "test_occm_demand_signal" {
  #Required
  compartment_id = var.sp_compartment_ocid
  display_name   = var.occm_demand_signal_display_name

  #Optional
  description   = var.occm_demand_signal_description
  freeform_tags = var.occm_demand_signal_freeform_tags
}

data "oci_capacity_management_occm_demand_signals" "test_occm_demand_signals" {
  #Required
  compartment_id = var.sp_compartment_ocid

  #Optional
  display_name      = var.occm_demand_signal_display_name
  id                = var.occm_demand_signal_id
}

