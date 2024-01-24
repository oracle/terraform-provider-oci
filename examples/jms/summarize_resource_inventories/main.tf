// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}

variable "summarize_resource_inventory_time_end" {}
variable "summarize_resource_inventory_time_start" {}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_jms_summarize_resource_inventory" "test_summarize_resource_inventories" {

  #Optional
  compartment_id = var.compartment_ocid
  time_end       = var.summarize_resource_inventory_time_end
  time_start     = var.summarize_resource_inventory_time_start
}
