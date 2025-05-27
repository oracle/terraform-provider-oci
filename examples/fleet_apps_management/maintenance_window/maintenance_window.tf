// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "maintenance_window_description" {
  default = "description of maintenance window"
}

variable "maintenance_window_display_name" {
  default = "maintenanceWindowDisplayName"
}

# 2 Hours
variable "maintenance_window_duration" {
  default = "PT2H"
}

variable "maintenance_window_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "maintenance_window_is_outage" {
  default = false
}

variable "maintenance_window_is_recurring" {
  default = false
}

# Not needed unless MW is recurring.
#   A recurring window might be set to "FREQ=DAILY;INTERVAL=1;COUNT=5",
#   for a window every day for 5 days.
variable "maintenance_window_recurrences" {
  default = "FREQ=MONTHLY;BYMONTHDAY=1;INTERVAL=1;COUNT=10"
}

variable "maintenance_window_state" {
  default = "ACTIVE"
}

variable "maintenance_window_time_schedule_start" {
  default = "2026-01-02T12:15:00.000Z"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_fleet_apps_management_maintenance_window" "test_maintenance_window" {
  #Required
  compartment_id = var.compartment_id
  duration       = var.maintenance_window_duration

  #Optional
  description             = var.maintenance_window_description
  display_name            = var.maintenance_window_display_name
  freeform_tags           = var.maintenance_window_freeform_tags
  is_outage               = var.maintenance_window_is_outage
  is_recurring            = var.maintenance_window_is_recurring
  recurrences             = var.maintenance_window_recurrences
  time_schedule_start     = var.maintenance_window_time_schedule_start
}

data "oci_fleet_apps_management_maintenance_windows" "test_maintenance_windows" {

  #Optional
  compartment_id = var.compartment_id
  display_name   = var.maintenance_window_display_name
  state          = var.maintenance_window_state
}

