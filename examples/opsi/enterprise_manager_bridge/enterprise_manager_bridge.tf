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
  default = "em_data_collection_bucket"
}

resource "oci_objectstorage_bucket" "test_bucket" {
  name = var.bucket_name
  compartment_id = var.compartment_ocid
  namespace = data.oci_objectstorage_namespace.test_namespace.namespace
}

resource "oci_identity_tag_namespace" "tag-namespace1" {
  compartment_id = var.tenancy_ocid
  description    = "example tag namespace"
  name           = "example-tag-namespace-all"
  is_retired = false
}


resource "oci_identity_tag" "tag1" {
  description      = "example tag"
  name             = "example-tag"
  tag_namespace_id = oci_identity_tag_namespace.tag-namespace1.id
  is_cost_tracking = false 
  is_retired       = false
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
  description 		     = var.enterprise_manager_bridge_description
}

output "enterprise_manager_bridge_id" {
  value = oci_opsi_enterprise_manager_bridge.test_enterprise_manager_bridge.id
}

// List EM Bridge present under a compartment having state ACTIVE
data "oci_opsi_enterprise_manager_bridges" "test_enterprise_manager_bridges" {
  compartment_id = var.compartment_ocid
  state          = var.enterprise_manager_bridge_state
}

// List EM Bridge present under a compartment having state ACTIVE in current and all subcompartments
data "oci_opsi_enterprise_manager_bridges" "test_enterprise_manager_bridges2" {
  compartment_id            = var.compartment_ocid
  compartment_id_in_subtree = var.compartment_id_in_subtree
  state                     = var.enterprise_manager_bridge_state
}

// Get EM Bridge for a particular id 
data "oci_opsi_enterprise_manager_bridge" "test_enterprise_manager_bridge" {
  enterprise_manager_bridge_id = oci_opsi_enterprise_manager_bridge.test_enterprise_manager_bridge.id 
}

