// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" { default = "eu-frankfurt-1" }
variable "compartment_id" {}

variable "discovery_schedule_display_name" {
  default = "displayName"
}

variable "discovery_schedule_execution_recurrences" {
  default = "FREQ=DAILY;BYHOUR=6"
}

variable "discovery_schedule_state" {
  default = "ACTIVE"
}


provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_cloud_bridge_discovery_schedule" "test_discovery_schedule" {
  compartment_id        = var.compartment_id
  execution_recurrences = var.discovery_schedule_execution_recurrences
  display_name  = var.discovery_schedule_display_name
}

data "oci_cloud_bridge_discovery_schedules" "test_discovery_schedules" {
  compartment_id = var.compartment_id
  discovery_schedule_id = oci_cloud_bridge_discovery_schedule.test_discovery_schedule.id
  display_name          = var.discovery_schedule_display_name
  state                 = var.discovery_schedule_state
}
