// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "scheduling_policy_cadence" {
  default = "QUARTERLY"
}

variable "scheduling_policy_cadence_start_month_name" {
  default = "FEBRUARY"
}

variable "scheduling_policy_defined_tags_value" {
  default = "definedTags"
}

variable "scheduling_policy_display_name" {
  default = "tstSchedulingPolicy1"
}

variable "scheduling_policy_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "scheduling_policy_state" {
  default = "AVAILABLE"
}

resource "oci_database_scheduling_policy" "test_scheduling_policy" {
  #Required
  cadence        = var.scheduling_policy_cadence
  compartment_id = var.compartment_id
  display_name   = var.scheduling_policy_display_name

  #Optional
  cadence_start_month {
    #Required
    name = var.scheduling_policy_cadence_start_month_name
  }
  defined_tags  = var.scheduling_policy_defined_tags_value
  freeform_tags = var.scheduling_policy_freeform_tags
}

data "oci_database_scheduling_policies" "test_scheduling_policies" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name = var.scheduling_policy_display_name
  state        = var.scheduling_policy_state
}
