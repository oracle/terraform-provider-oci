// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "resource_id_for_maintwin" {}

variable "maintenance_window_name" {
  default = "TestMaintenanceWindows"
}

variable "maintenance_window_schedule_schedule_type" {
  default = "ONE_TIME"
}

variable "maintenance_window_schedule_time_maintenance_window_start" {
  default = "2024-10-25T16:00:01.001Z"
}

variable "maintenance_window_schedule_time_maintenance_window_end" {
  default = "2024-10-26T16:00:01.001Z"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_stack_monitoring_maintenance_window" "test_maintenance_window_example" {
	#Required
	compartment_id  = var.compartment_ocid
	name            = var.maintenance_window_name

	resources {
        resource_id           = var.resource_id_for_maintwin
        are_members_included  = true
    }

    schedule {
        schedule_type                   = var.maintenance_window_schedule_schedule_type
        time_maintenance_window_start   = var.maintenance_window_schedule_time_maintenance_window_start
        time_maintenance_window_end     = var.maintenance_window_schedule_time_maintenance_window_end
  }

}
