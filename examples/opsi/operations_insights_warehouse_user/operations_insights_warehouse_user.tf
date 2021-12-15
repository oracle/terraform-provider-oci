// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0


variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}


provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_identity_tag_namespace" "tag-namespace1" {
  compartment_id = var.tenancy_ocid
  description    = "example tag namespace"
  name           = "examples-tag-namespace-all"
  is_retired     = false
}

resource "oci_identity_tag" "tag1" {
  description      = "example tag"
  name             = "example-tag"
  tag_namespace_id = oci_identity_tag_namespace.tag-namespace1.id
  is_cost_tracking = false
  is_retired       = false
}

variable "wh_user_defined_tags_value" {
  default = "wh_user_tag_value"
}

variable "wh_user_name" {
  default = "TestWarehouseUserName"
}

variable "wh_user_connection_password" {
  default = "connectionPassword1"
}

variable "wh_user_is_awr_data_access" {
  default = "false"
}

variable "wh_user_is_em_data_access" {
  default = "false"
}

variable "wh_user_is_opsi_data_access" {
  default = "false"
}

variable "wh_user_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "wh_user_state" {
  default = ["ACTIVE"]
}

variable "warehouse_defined_tags_value" {
  default = "warehouse_tag_value"
}

variable "warehouse_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "warehouse_display_name" {
  default = "TestWarehouseDisplayName"
}

variable "warehouse_cpu_allocated" {
  default = 1.0
}

variable "storage_allocated_in_gbs" {
  default = 1.0
}

// To Create a Warehouse
resource "oci_opsi_operations_insights_warehouse" "test_operations_insights_warehouse" {
  #Required
  compartment_id             = var.compartment_ocid
  cpu_allocated              = var.warehouse_cpu_allocated
  display_name               = var.warehouse_display_name

  #Optional
  defined_tags               = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.warehouse_defined_tags_value}")}"
  freeform_tags              = var.warehouse_freeform_tags
  storage_allocated_in_gbs 	 = var.storage_allocated_in_gbs
}

// To Create a wh_user
resource "oci_opsi_operations_insights_warehouse_user" "test_operations_insights_warehouse_user" {
  #Required
  compartment_id             = var.compartment_ocid
  connection_password        = var.wh_user_connection_password
  is_awr_data_access         = var.wh_user_is_awr_data_access
  name                       = var.wh_user_name
  operations_insights_warehouse_id = oci_opsi_operations_insights_warehouse.test_operations_insights_warehouse.id

  #Optional
  defined_tags               = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.wh_user_defined_tags_value}")}"
  freeform_tags              = var.wh_user_freeform_tags
  is_em_data_access          = var.wh_user_is_em_data_access
  is_opsi_data_access        = var.wh_user_is_opsi_data_access
}

output "operations_insights_warehouse_user_id" {
  value = oci_opsi_operations_insights_warehouse_user.test_operations_insights_warehouse_user.id
}

// List wh_user present under a compartment having state ACTIVE
data "oci_opsi_operations_insights_warehouse_users" "test_operations_insights_warehouse_users" {
  operations_insights_warehouse_id = oci_opsi_operations_insights_warehouse.test_operations_insights_warehouse.id
  compartment_id = var.compartment_ocid
  state          = var.wh_user_state
}

// Get wh_user for a particular id
data "oci_opsi_operations_insights_warehouse_user" "test_operations_insights_warehouse_user" {
  operations_insights_warehouse_user_id = oci_opsi_operations_insights_warehouse_user.test_operations_insights_warehouse_user.id
}
