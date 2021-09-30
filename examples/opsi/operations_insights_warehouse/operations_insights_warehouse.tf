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

data "oci_objectstorage_namespace" "test_namespace" {
    compartment_id = var.compartment_ocid
}

variable "bucket_name" {
  default = "warehouse_bucket"
}

resource "oci_objectstorage_bucket" "test_bucket" {
  name           = var.bucket_name
  compartment_id = var.compartment_ocid
  namespace      = data.oci_objectstorage_namespace.test_namespace.namespace
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

variable "warehouse_defined_tags_value" {
  default = "warehouse_tag_value"
}

variable "warehouse_display_name" {
  default = "TestWarehouseDisplayName"
}

variable "warehouse_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "warehouse_cpu_allocated" {
  default = 1.0
}

variable "storage_allocated_in_gbs" {
  default = 1.0
}

variable "warehouse_state" {
  default = ["ACTIVE"]
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

output "operations_insights_warehouse_id" {
  value = oci_opsi_operations_insights_warehouse.test_operations_insights_warehouse.id
}

// List Warehouse present under a compartment having state ACTIVE
data "oci_opsi_operations_insights_warehouses" "test_operations_insights_warehouses" {
  compartment_id = var.compartment_ocid
  state          = var.warehouse_state
}

// Get Warehouse for a particular id
data "oci_opsi_operations_insights_warehouse" "test_operations_insights_warehouse" {
  operations_insights_warehouse_id = oci_opsi_operations_insights_warehouse.test_operations_insights_warehouse.id
}

// Get Resource Usage summary for a particular Opsi Warehouse Id
data "oci_opsi_operations_insights_warehouse_resource_usage_summary" "test_operations_insights_warehouse_resource_usage_summary" {
  operations_insights_warehouse_id = oci_opsi_operations_insights_warehouse.test_operations_insights_warehouse.id
}

output "resource_usage_summary_output" {
  value = length(data.oci_opsi_operations_insights_warehouse_resource_usage_summary.test_operations_insights_warehouse_resource_usage_summary)
}

// Rotate Wallet
data "oci_opsi_operations_insights_warehouse_rotate_warehouse_wallet" "test_operations_insights_warehouse_rotate_warehouse_wallet" {
  operations_insights_warehouse_id = oci_opsi_operations_insights_warehouse.test_operations_insights_warehouse.id
}

// Declare the password to be used.
variable "operations_insights_warehouse_wallet_password" {
  default = "Admin@1234"
}

// Download Wallet
data "oci_opsi_operations_insights_warehouse_download_warehouse_wallet" "test_operations_insights_warehouse_download_warehouse_wallet" {
  operations_insights_warehouse_id = oci_opsi_operations_insights_warehouse.test_operations_insights_warehouse.id
  operations_insights_warehouse_wallet_password = var.operations_insights_warehouse_wallet_password
}
