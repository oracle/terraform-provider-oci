// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

variable "inventory_display_name" {
  default = "displayName"
}

variable "inventory_state" {
  default = "ACTIVE"
}


provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_cloud_bridge_inventory" "test_inventory" {
  compartment_id = var.tenancy_ocid
  display_name   = var.inventory_display_name
}

data "oci_cloud_bridge_inventories" "test_inventories" {
  compartment_id = var.tenancy_ocid
  state = var.inventory_state
}
