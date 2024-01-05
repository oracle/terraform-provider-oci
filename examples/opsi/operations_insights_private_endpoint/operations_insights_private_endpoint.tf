// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "vcn_id" {}
variable "subnet_id" {}

variable "operations_insights_private_endpoint_compartment_id_in_subtree" {
  default = false
}

variable "operations_insights_private_endpoint_defined_tags_value" {
  default = "value"
}

variable "operations_insights_private_endpoint_description" {
  default = "TestDescription"
}

variable "operations_insights_private_endpoint_display_name" {
  default = "TestPrivateEndpoint"
}

variable "operations_insights_private_endpoint_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "operations_insights_private_endpoint_is_used_for_rac_dbs" {
  default = false
}

variable "operations_insights_private_endpoint_nsg_ids" {
  default = []
}

variable "operations_insights_private_endpoint_state" {
  default = ["ACTIVE"]
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
  is_retired = false
}

resource "oci_identity_tag" "tag1" {
  description      = "example tag"
  name             = "example-tag"
  tag_namespace_id = oci_identity_tag_namespace.tag-namespace1.id
  is_cost_tracking = false
  is_retired       = false
}

resource "oci_opsi_operations_insights_private_endpoint" "test_operations_insights_private_endpoint" {
  compartment_id      = var.compartment_ocid
  display_name        = var.operations_insights_private_endpoint_display_name
  description         = var.operations_insights_private_endpoint_description
  is_used_for_rac_dbs = var.operations_insights_private_endpoint_is_used_for_rac_dbs
  subnet_id           = var.subnet_id
  vcn_id              = var.vcn_id
  defined_tags        = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.operations_insights_private_endpoint_defined_tags_value}")}"
  freeform_tags       = var.operations_insights_private_endpoint_freeform_tags
  nsg_ids             = var.operations_insights_private_endpoint_nsg_ids
}

data "oci_opsi_operations_insights_private_endpoints" "test_operations_insights_private_endpoints" {

  #Optional
  compartment_id            = var.compartment_ocid
  compartment_id_in_subtree = var.operations_insights_private_endpoint_compartment_id_in_subtree
  display_name              = var.operations_insights_private_endpoint_display_name
  is_used_for_rac_dbs       = var.operations_insights_private_endpoint_is_used_for_rac_dbs
  opsi_private_endpoint_id  = oci_opsi_operations_insights_private_endpoint.test_operations_insights_private_endpoint.id
  state                     = var.operations_insights_private_endpoint_state
  vcn_id                    = var.vcn_id
}

