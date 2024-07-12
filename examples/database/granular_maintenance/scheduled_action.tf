// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "scheduled_action_action_members_estimated_time_in_mins" {
  default = 90
}

variable "scheduled_action_action_members_member_order" {
  default = 1
}

variable "scheduled_action_action_params" {
  default = {}
}

variable "scheduled_action_action_type" {
  default = "DB_SERVER_FULL_SOFTWARE_UPDATE"
}

variable "scheduled_action_defined_tags_value" {
  default = "definedTags"
}

variable "scheduled_action_display_name" {
  default = "displayName"
}

variable "scheduled_action_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "scheduled_action_id" {
  default = "id"
}

variable "scheduled_action_service_type" {
  default = "serviceType"
}

variable "scheduled_action_state" {
  default = "AVAILABLE"
}


resource "oci_database_scheduled_action" "test_scheduled_action" {
  #Required
  action_type          = var.scheduled_action_action_type
  compartment_id       = var.compartment_id
  scheduling_plan_id   = oci_database_scheduling_plan.test_scheduling_plan.id
  scheduling_window_id = oci_database_scheduling_policy_scheduling_window.test_scheduling_policy_scheduling_window.id

  #Optional
//  action_members {
//    #Required
//    member_id    = oci_database_member.test_member.id
//    member_order = var.scheduled_action_action_members_member_order
//
//    #Optional
//    estimated_time_in_mins = var.scheduled_action_action_members_estimated_time_in_mins
//  }
//  action_params = var.scheduled_action_action_params
  defined_tags  = var.scheduled_action_defined_tags_value
  freeform_tags = var.scheduled_action_freeform_tags
}

data "oci_database_scheduled_actions" "test_scheduled_actions" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name       = var.scheduled_action_display_name
  id                 = var.scheduled_action_id
  scheduling_plan_id = oci_database_scheduling_plan.test_scheduling_plan.id
  service_type       = var.scheduled_action_service_type
  state              = var.scheduled_action_state
}