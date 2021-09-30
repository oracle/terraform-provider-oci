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
  default = "awrhub_bucket"
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

variable "awrhub_defined_tags_value" {
  default = "awrhub_tag_value"
}

variable "awrhub_display_name" {
  default = "TestAwrhubDisplayName"
}

variable "awrhub_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "awrhub_state" {
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

// To Create a awrhub
resource "oci_opsi_awr_hub" "test_awr_hub" {
  #Required
  compartment_id             = var.compartment_ocid
  display_name               = var.awrhub_display_name
  object_storage_bucket_name = oci_objectstorage_bucket.test_bucket.name
  operations_insights_warehouse_id = oci_opsi_operations_insights_warehouse.test_operations_insights_warehouse.id

  #Optional
  defined_tags               = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.awrhub_defined_tags_value}")}"
  freeform_tags              = var.awrhub_freeform_tags
}

output "awr_hub_id" {
  value = oci_opsi_awr_hub.test_awr_hub.id
}

// List awrhub present under a compartment having state ACTIVE
data "oci_opsi_awr_hubs" "test_awr_hubs" {
  operations_insights_warehouse_id = oci_opsi_operations_insights_warehouse.test_operations_insights_warehouse.id
  compartment_id = var.compartment_ocid
  state          = var.awrhub_state
}

// Get awrhub for a particular id
data "oci_opsi_awr_hub" "test_awr_hub" {
  awr_hub_id = oci_opsi_awr_hub.test_awr_hub.id
}

// Get source summary for a particular AWR Hub id
data "oci_opsi_awr_hub_awr_sources_summary" "test_awr_hub_awr_sources_summary" {
  awr_hub_id = oci_opsi_awr_hub.test_awr_hub.id
}

output "source_summary_output" {
  value = length(data.oci_opsi_awr_hub_awr_sources_summary.test_awr_hub_awr_sources_summary)
}

variable "awr_source_database_identifier" {
  default = "12345"
}

// Get snapshots summary for a particular AWR Hub id
data "oci_opsi_awr_hub_awr_snapshots" "test_awr_hub_awr_snapshots" {
  awr_hub_id = oci_opsi_awr_hub.test_awr_hub.id
  awr_source_database_identifier = var.awr_source_database_identifier
}

output "snapshots_summary_output" {
  value = length(data.oci_opsi_awr_hub_awr_snapshots.test_awr_hub_awr_snapshots)
}
