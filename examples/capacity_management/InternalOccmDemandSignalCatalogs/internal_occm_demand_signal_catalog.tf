// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "prod_compartment_ocid" {}

variable "internal_occm_demand_signal_catalog_display_name" {
  default = "displayName"
}

variable "customergroup_id" {}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_capacity_management_internal_occm_demand_signal_catalogs" "test_internal_occm_demand_signal_catalogs" {
  #Required
  compartment_id        = var.prod_compartment_ocid
  occ_customer_group_id = var.customergroup_id

  #Optional
  display_name = var.internal_occm_demand_signal_catalog_display_name
}

