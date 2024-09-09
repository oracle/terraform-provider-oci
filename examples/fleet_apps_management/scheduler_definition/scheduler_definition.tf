// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "test_active_fleet" {
  type = "string"
  description = "This should be the OCID of a Fleet with resources, in the ACTIVE state."
}

variable "test_runbook_ocid" {
  type = "string"
  description = "OCID of an active Runbook. The Oracle managed Runbooks are created for you when you onboard your tenancy."
}


variable "scheduler_definition_action_groups_lifecycle_operation" {
  default = "PATCH"
}

variable "scheduler_definition_action_groups_product" {
  default = "WEBLOGIC/JAVA"
}

variable "scheduler_definition_action_groups_subjects" {
  default = []
}

variable "scheduler_definition_action_groups_type" {
  default = "PRODUCT"
}

# Must be less than duration.
variable "scheduler_definition_activity_initiation_cut_off" {
  default = 1
}

variable "scheduler_definition_defined_tags_value" {
  default = "value"
}

variable "scheduler_definition_description" {
  default = "schedule description"
}

variable "scheduler_definition_display_name" {
  default = "schedulerDefinitionDisplayName"
}

variable "scheduler_definition_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "scheduler_definition_product" {
  default = "WEBLOGIC/JAVA"
}

variable "scheduler_definition_run_books_input_parameters_arguments_name" {
  default = "name"
}

variable "scheduler_definition_run_books_input_parameters_arguments_value" {
  default = "value"
}

variable "scheduler_definition_run_books_input_parameters_step_name" {
  default = "stepName"
}

# For example, 2hr.
variable "scheduler_definition_schedule_duration" {
  default = "PT2H"
}

variable "scheduler_definition_schedule_execution_startdate" {
  default = "2025-01-02T12:15:00.000Z"
}

# Only required for recurring schedules.
variable "scheduler_definition_schedule_recurrences" {
  default = ""
}

variable "scheduler_definition_schedule_type" {
  default = "CUSTOM"
}

variable "scheduler_definition_state" {
  default = "ACTIVE"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

## Supporting Resources

resource "oci_fleet_apps_management_maintenance_window" "test_maintenance_window" {
  #Required
  compartment_id = var.tenancy_ocid
  duration       = var.scheduler_definition_schedule_duration

  #Optional
  description             = "test maintenance window"
  display_name            = "schedulerExampleMaintenanceWindow"
  is_outage               = false
  is_recurring            = false
  maintenance_window_type = "OPEN_ENDED"
  task_initiation_cutoff  = 1
  time_schedule_start     = var.scheduler_definition_schedule_execution_startdate
}


#########


resource "oci_fleet_apps_management_scheduler_definition" "test_scheduler_definition" {
  #Required
  action_groups {
    #Required
    resource_id = var.test_active_fleet
    runbook_id  = var.test_runbook_ocid

    #Optional
    lifecycle_operation = var.scheduler_definition_action_groups_lifecycle_operation
    product             = var.scheduler_definition_action_groups_product
    type                = var.scheduler_definition_action_groups_type
  }
  compartment_id = var.tenancy_ocid
  schedule {
    #Required
    execution_startdate = var.scheduler_definition_schedule_execution_startdate
    type                = var.scheduler_definition_schedule_type

    #Optional
    duration              = var.scheduler_definition_schedule_duration
    maintenance_window_id = oci_fleet_apps_management_maintenance_window.test_maintenance_window.id
    recurrences           = var.scheduler_definition_schedule_recurrences
  }

  #Optional
  activity_initiation_cut_off = var.scheduler_definition_activity_initiation_cut_off
  description                 = var.scheduler_definition_description
  display_name                = var.scheduler_definition_display_name
  freeform_tags               = var.scheduler_definition_freeform_tags
  run_books {
    #Required
    id = var.test_runbook_ocid

    #Optional
    input_parameters {
      #Required
      step_name = var.scheduler_definition_run_books_input_parameters_step_name

      #Optional
      arguments {
        #Required
        name = var.scheduler_definition_run_books_input_parameters_arguments_name

        #Optional
        value = var.scheduler_definition_run_books_input_parameters_arguments_value
      }
    }
  }
}

data "oci_fleet_apps_management_scheduler_definitions" "test_scheduler_definitions" {

  #Optional
  compartment_id        = var.compartment_id
  display_name          = var.scheduler_definition_display_name
  fleet_id              = var.test_active_fleet
  maintenance_window_id = oci_fleet_apps_management_maintenance_window.test_maintenance_window.id
  product               = var.scheduler_definition_product
  state                 = var.scheduler_definition_state
}

