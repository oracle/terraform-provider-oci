// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
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

variable "awr_hub_source_name" {
  default = "tstHubSrc"
}

variable "awr_hub_source_type" {
  default = "ADW_S"
}

variable "awr_hub_source_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "awrhubsource_defined_tags_value" {
  default = "awrhubsource_tag_value"
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
  storage_allocated_in_gbs   = var.storage_allocated_in_gbs
}

// To Create a awrhub
resource "oci_opsi_awr_hub" "test_awr_hub" {
  #Required
  compartment_id             = var.compartment_ocid
  display_name               = var.awrhub_display_name
  operations_insights_warehouse_id = oci_opsi_operations_insights_warehouse.test_operations_insights_warehouse.id

  #Optional
  defined_tags               = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.awrhub_defined_tags_value}")}"
  freeform_tags              = var.awrhub_freeform_tags
}

output "awr_hub_id" {
  value = oci_opsi_awr_hub.test_awr_hub.id
}

resource "oci_opsi_awr_hub_source" "test_awr_hub_source" {
  #Required
  awr_hub_id     = oci_opsi_awr_hub.test_awr_hub.id
  compartment_id = var.compartment_ocid
  name           = var.awr_hub_source_name
  type           = var.awr_hub_source_type

  #Optional
  defined_tags               = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.awrhubsource_defined_tags_value}")}"
  freeform_tags              = var.awr_hub_source_freeform_tags
}

output "awr_hub_source_id" {
  value = oci_opsi_awr_hub_source.test_awr_hub_source.id
}

// List awrhub source present under a compartment having state ACTIVE
data "oci_opsi_awr_hub_sources" "test_awr_hub_sources" {
  awr_hub_id     = oci_opsi_awr_hub.test_awr_hub.id
  compartment_id = var.compartment_ocid
  state          = var.awrhub_state
}

// Get awrhub source for a particular id
data "oci_opsi_awr_hub_source" "test_awr_hub_source" {
  awr_hub_source_id = oci_opsi_awr_hub_source.test_awr_hub_source.id
}


