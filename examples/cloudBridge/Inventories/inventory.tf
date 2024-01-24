// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "inventory_defined_tags_value" {
  default = "value"
}

variable "inventory_display_name" {
  default = "displayName"
}

variable "inventory_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "inventory_state" {
  default = "AVAILABLE"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_cloud_bridge_inventory" "test_inventory" {
  #Required
  compartment_id = var.compartment_id
  display_name   = var.inventory_display_name

  #Optional
  defined_tags  = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.inventory_defined_tags_value)
  freeform_tags = var.inventory_freeform_tags
}

data "oci_cloud_bridge_inventories" "test_inventories" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  state = var.inventory_state
}

