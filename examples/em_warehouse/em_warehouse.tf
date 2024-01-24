// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}

variable "em_warehouse_defined_tags_value" {
  default = "value"
}

variable "em_warehouse_display_name" {
  default = "displayName"
}

variable "em_warehouse_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "em_warehouse_state" {
  default = "ACTIVE"
}



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

// To Create a OPSI Warehouse
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

data "oci_objectstorage_namespace" "test_namespace" {
    compartment_id = var.compartment_ocid
}

variable "bucket_name" {
  default = "em_data_collection_bucket"
}

resource "oci_objectstorage_bucket" "test_bucket" {
  name = var.bucket_name
  compartment_id = var.compartment_ocid
  namespace = data.oci_objectstorage_namespace.test_namespace.namespace
}

variable "enterprise_manager_bridge_defined_tags_value" {
  default = "embridge_tag_value"
}

variable "enterprise_manager_bridge_description" {
  default = "Test EM Bridge Description"
}

variable "enterprise_manager_bridge_display_name" {
  default = "TestEMManagedBridgeName"
}

variable "enterprise_manager_bridge_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "enterprise_manager_bridge_state" {
  default = ["ACTIVE"]
}

variable "compartment_id_in_subtree" {
  default = true
}

// To Create a Enterprise Manager Bridge
resource "oci_opsi_enterprise_manager_bridge" "test_enterprise_manager_bridge" {
  #Required
  compartment_id             = var.compartment_ocid
  display_name               = var.enterprise_manager_bridge_display_name
  object_storage_bucket_name = oci_objectstorage_bucket.test_bucket.name

  #Optional
  defined_tags               = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.enterprise_manager_bridge_defined_tags_value}")}"
  freeform_tags              = var.enterprise_manager_bridge_freeform_tags
  description 		           = var.enterprise_manager_bridge_description
}

// To Create a Enterprise Manager Warehouse
resource "oci_em_warehouse_em_warehouse" "test_em_warehouse" {
  #Required
  compartment_id                   = var.compartment_ocid
  em_bridge_id                     = oci_opsi_enterprise_manager_bridge.test_enterprise_manager_bridge.id
  operations_insights_warehouse_id = oci_opsi_operations_insights_warehouse.test_operations_insights_warehouse.id

  #Optional
  defined_tags  = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.em_warehouse_defined_tags_value}")}"
  display_name  = var.em_warehouse_display_name
  freeform_tags = var.em_warehouse_freeform_tags
}

// List emWarehouses present under a compartment
data "oci_em_warehouse_em_warehouses" "test_em_warehouses" {
  #Required
  compartment_id                   = var.compartment_ocid

  #Optional
  display_name                     = var.em_warehouse_display_name
  id                               = oci_em_warehouse_em_warehouse.test_em_warehouse.id
  operations_insights_warehouse_id = oci_opsi_operations_insights_warehouse.test_operations_insights_warehouse.id
  state                            = var.em_warehouse_state
}

// Get emWarehouse for a particular id
data "oci_em_warehouse_em_warehouse" "test_em_warehouse" {
  #Required
  em_warehouse_id = oci_em_warehouse_em_warehouse.test_em_warehouse.id
}

// Get ETL Runs for a particular emWarehouseId
data "oci_em_warehouse_em_warehouse_etl_runs" "test_em_warehouse_etl_runs" {
  #Required
  em_warehouse_id = oci_em_warehouse_em_warehouse.test_em_warehouse.id
}

// Get Resource Usage for a particular emWarehouseId
data "oci_em_warehouse_em_warehouse_resource_usage" "test_em_warehouse_resource_usage" {
  #Required
  em_warehouse_id = oci_em_warehouse_em_warehouse.test_em_warehouse.id
}