// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "discovery_schedule_defined_tags_value" {
  default = "value"
}

variable "discovery_schedule_display_name" {
  default = "displayName"
}

variable "discovery_schedule_execution_recurrences" {
  default = "executionRecurrences"
}

variable "discovery_schedule_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "discovery_schedule_state" {
  default = "AVAILABLE"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_cloud_bridge_discovery_schedule" "test_discovery_schedule" {
  #Required
  compartment_id        = var.compartment_id
  execution_recurrences = var.discovery_schedule_execution_recurrences

  #Optional
  defined_tags  = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.discovery_schedule_defined_tags_value)
  display_name  = var.discovery_schedule_display_name
  freeform_tags = var.discovery_schedule_freeform_tags
}

data "oci_cloud_bridge_discovery_schedules" "test_discovery_schedules" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  discovery_schedule_id = oci_cloud_bridge_discovery_schedule.test_discovery_schedule.id
  display_name          = var.discovery_schedule_display_name
  state                 = var.discovery_schedule_state
}

