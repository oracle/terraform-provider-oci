// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "ccc_infrastructure_access_level" {
  default = "RESTRICTED"
}

variable "ccc_infrastructure_compartment_id_in_subtree" {
  default = false
}

variable "ccc_infrastructure_connection_details" {
  default = "connectionDetails"
}

variable "ccc_infrastructure_connection_state" {
  default = "REJECT"
}

variable "ccc_infrastructure_defined_tags_value" {
  default = "value"
}

variable "ccc_infrastructure_description" {
  default = "Datacenter 231"
}

variable "ccc_infrastructure_display_name" {
  default = "example_cccInfrastructure"
}

variable "ccc_infrastructure_display_name_contains" {
  default = "displayNameContains"
}

variable "ccc_infrastructure_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "ccc_infrastructure_state" {
  default = "ACTIVE"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_compute_cloud_at_customer_ccc_upgrade_schedule" "test_ccc_upgrade_schedule" {
compartment_id = var.compartment_id
display_name = "example_cccUpgradeSchedule"
events {
description = "description"
schedule_event_duration = "PT49H"
schedule_event_recurrences = "FREQ=MONTHLY;INTERVAL=3;"
time_start = "2023-09-09T16:10:25Z"
}
}

resource "oci_core_vcn" "test_vcn" {
cidr_block = "10.0.0.0/16"
compartment_id = var.compartment_id
lifecycle {
ignore_changes = ["defined_tags"]
}
}

resource "oci_core_subnet" "test_subnet" {
cidr_block = "10.0.0.0/24"
compartment_id = var.compartment_id
lifecycle {
ignore_changes = ["defined_tags"]
}
vcn_id = oci_core_vcn.test_vcn.id
}

variable defined_tag_namespace_name { default = "" }
resource "oci_identity_tag_namespace" "tag-namespace1" {
  		#Required
		compartment_id = var.compartment_id
  		description = "example tag namespace"
  		name = "${var.defined_tag_namespace_name != "" ? var.defined_tag_namespace_name : "example-tag-namespace-all"}"

		is_retired = false
    lifecycle {
        ignore_changes = []
    }
}

resource "oci_identity_tag" "tag1" {
  		#Required
  		description = "example tag"
  		name = "example-tag"
        tag_namespace_id = oci_identity_tag_namespace.tag-namespace1.id

		is_retired = false
}

resource "oci_compute_cloud_at_customer_ccc_infrastructure" "test_ccc_infrastructure" {
  #Required
  compartment_id = var.compartment_id
  display_name   = var.ccc_infrastructure_display_name
  subnet_id      = oci_core_subnet.test_subnet.id

  #Optional
  ccc_upgrade_schedule_id = oci_compute_cloud_at_customer_ccc_upgrade_schedule.test_ccc_upgrade_schedule.id
  connection_details      = var.ccc_infrastructure_connection_details
  connection_state        = var.ccc_infrastructure_connection_state
  defined_tags            = map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", var.ccc_infrastructure_defined_tags_value)
  description             = var.ccc_infrastructure_description
  freeform_tags           = var.ccc_infrastructure_freeform_tags
}

data "oci_compute_cloud_at_customer_ccc_infrastructures" "test_ccc_infrastructures" {

  #Optional
  access_level              = var.ccc_infrastructure_access_level
  ccc_infrastructure_id     = oci_compute_cloud_at_customer_ccc_infrastructure.test_ccc_infrastructure.id
  compartment_id            = var.compartment_id
  compartment_id_in_subtree = var.ccc_infrastructure_compartment_id_in_subtree
  display_name              = var.ccc_infrastructure_display_name
  display_name_contains     = var.ccc_infrastructure_display_name_contains
  state                     = var.ccc_infrastructure_state
}

