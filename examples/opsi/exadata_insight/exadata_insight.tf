// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "enterprise_manager_bridge_ocid" {}
variable "em_exadata_enterprise_manager_entity_id" {}
variable "em_exadata_enterprise_manager_id" {}

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
  is_retired = false
}


resource "oci_identity_tag" "tag1" {
  description      = "example tag"
  name             = "example-tag"
  tag_namespace_id = oci_identity_tag_namespace.tag-namespace1.id
  is_cost_tracking = false
  is_retired       = false
}

variable "exadata_insight_exadata_type" {
  default = ["DBMACHINE"]
}

variable "exadata_insight_defined_tags_value" {
  default = "value"
}

variable "exadata_insight_entity_source" {
  default = "EM_MANAGED_EXTERNAL_EXADATA"
}

variable "exadata_insight_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "resource_status" {
  default = "ENABLED"
}

// Create Exadata insight for EM managed External Exadata
resource "oci_opsi_exadata_insight" "test_exadata_insight" {
  compartment_id                       = var.compartment_ocid
  enterprise_manager_bridge_id         = var.enterprise_manager_bridge_ocid
  enterprise_manager_entity_identifier = var.em_exadata_enterprise_manager_entity_id
  enterprise_manager_identifier        = var.em_exadata_enterprise_manager_id
  entity_source                        = var.exadata_insight_entity_source
  defined_tags                         = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.exadata_insight_defined_tags_value}")}"
  freeform_tags                        = var.exadata_insight_freeform_tags
  status                               = var.resource_status
}

variable "exadata_insight_state" {
  default = ["ACTIVE"]
}

variable "exadata_insight_status" {
  default = ["ENABLED"]
}

variable "compartment_id_in_subtree" {
  default = true
}

// List EM managed external exadata insights
data "oci_opsi_exadata_insights" "test_exadata_insights" {

  #Optional
  compartment_id               = var.compartment_ocid
  exadata_type                 = var.exadata_insight_exadata_type
  enterprise_manager_bridge_id = var.enterprise_manager_bridge_ocid
  state                        = var.exadata_insight_state
  status                       = var.exadata_insight_status
}

// List EM managed external exadata insights in current and subcompartments
data "oci_opsi_exadata_insights" "test_exadata_insights2" {

  #Optional
  compartment_id               = var.compartment_ocid
  compartment_id_in_subtree    = var.compartment_id_in_subtree
  exadata_type                 = var.exadata_insight_exadata_type
  enterprise_manager_bridge_id = var.enterprise_manager_bridge_ocid
  state                        = var.exadata_insight_state
  status                       = var.exadata_insight_status
}

// Get an EM managed exadata insight
data "oci_opsi_exadata_insight" "test_exadata_insight" {
  exadata_insight_id = oci_opsi_exadata_insight.test_exadata_insight.id
}

