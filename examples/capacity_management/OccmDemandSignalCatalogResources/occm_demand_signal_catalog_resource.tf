// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "sp_compartment_ocid" {}

variable "occm_demand_signal_catalog_resource_demand_signal_namespace" {
  default = "COMPUTE"
}

variable "occm_demand_signal_catalog_resource_name" {
  default = "name"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_capacity_management_occm_demand_signal_catalog_resources" "test_occm_demand_signal_catalog_resources" {
  #Required
  compartment_id = var.sp_compartment_ocid

  #Optional
  demand_signal_namespace = var.occm_demand_signal_catalog_resource_demand_signal_namespace
  name                    = var.occm_demand_signal_catalog_resource_name
}

