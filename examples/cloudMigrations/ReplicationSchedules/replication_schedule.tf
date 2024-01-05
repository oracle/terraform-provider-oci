// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "replication_schedule_defined_tags_value" {
  default = "value"
}

variable "replication_schedule_display_name" {
  default = "displayName"
}

variable "replication_schedule_execution_recurrences" {
  default = "executionRecurrences"
}

variable "replication_schedule_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "replication_schedule_state" {
  default = "AVAILABLE"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_cloud_migrations_replication_schedule" "test_replication_schedule" {
  #Required
  compartment_id        = var.compartment_id
  display_name          = var.replication_schedule_display_name
  execution_recurrences = var.replication_schedule_execution_recurrences

  #Optional
  defined_tags  = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.replication_schedule_defined_tags_value)
  freeform_tags = var.replication_schedule_freeform_tags
}

data "oci_cloud_migrations_replication_schedules" "test_replication_schedules" {

  #Optional
  compartment_id          = var.compartment_id
  display_name            = var.replication_schedule_display_name
  replication_schedule_id = oci_cloud_migrations_replication_schedule.test_replication_schedule.id
  state                   = var.replication_schedule_state
}

