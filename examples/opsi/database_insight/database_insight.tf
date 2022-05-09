// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "enterprise_manager_bridge_ocid" {}
variable "enterprise_manager_entity_id" {}
variable "enterprise_manager_id" {}

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

variable "database_insight_database_type" {
  default = ["EXTERNAL-NONCDB"]
}

variable "database_insight_defined_tags_value" {
  default = "value"
}

variable "database_insight_entity_source" {
  default = "EM_MANAGED_EXTERNAL_DATABASE"
}

variable "database_insight_fields" {
  default = ["databaseName", "databaseType", "compartmentId", "databaseDisplayName", "freeformTags", "definedTags"]
}

variable "database_insight_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "resource_status" {
  default = "ENABLED"
}

// Create Database insight for EM managed External Database
resource "oci_opsi_database_insight" "test_database_insight" {
  compartment_id                       = var.compartment_ocid
  enterprise_manager_bridge_id         = var.enterprise_manager_bridge_ocid
  enterprise_manager_entity_identifier = var.enterprise_manager_entity_id
  enterprise_manager_identifier        = var.enterprise_manager_id
  entity_source                        = var.database_insight_entity_source
  defined_tags                         = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.database_insight_defined_tags_value}")}"
  freeform_tags                        = var.database_insight_freeform_tags
  status                               = var.resource_status
}

variable "database_insight_state" {
  default = ["ACTIVE"]
}

variable "database_insight_status" {
  default = ["ENABLED"]
}

// List EM managed database insights
data "oci_opsi_database_insights" "test_database_insights" {

  #Optional
  compartment_id               = var.compartment_ocid
  database_type                = var.database_insight_database_type
  enterprise_manager_bridge_id = var.enterprise_manager_bridge_ocid
  fields                       = var.database_insight_fields
  state                        = var.database_insight_state
  status                       = var.database_insight_status
}

// Get an EM managed database insight
data "oci_opsi_database_insight" "test_database_insight" {
  database_insight_id = oci_opsi_database_insight.test_database_insight.id
}

