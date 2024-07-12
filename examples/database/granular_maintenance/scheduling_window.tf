// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "scheduling_policy_scheduling_window_defined_tags_value" {
  default = "definedTags"
}

variable "scheduling_policy_scheduling_window_display_name" {
  default = "displayName"
}

variable "scheduling_policy_scheduling_window_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "scheduling_policy_scheduling_window_state" {
  default = "AVAILABLE"
}

variable "scheduling_policy_scheduling_window_window_preference_days_of_week_name" {
  default = "MONDAY"
}

variable "scheduling_policy_scheduling_window_window_preference_duration" {
  default = 180
}

variable "scheduling_policy_scheduling_window_window_preference_is_enforced_duration" {
  default = false
}

variable "scheduling_policy_scheduling_window_window_preference_months_name1" {
  default = "FEBRUARY"
}

variable "scheduling_policy_scheduling_window_window_preference_months_name2" {
  default = "MAY"
}

variable "scheduling_policy_scheduling_window_window_preference_months_name3" {
  default = "AUGUST"
}

variable "scheduling_policy_scheduling_window_window_preference_months_name4" {
  default = "NOVEMBER"
}

variable "scheduling_policy_scheduling_window_window_preference_start_time" {
  default = "10:00"
}

variable "scheduling_policy_scheduling_window_window_preference_weeks_of_month" {
  default = ["1"]
}


resource "oci_database_scheduling_policy_scheduling_window" "test_scheduling_policy_scheduling_window" {
  #Required
  scheduling_policy_id = oci_database_scheduling_policy.test_scheduling_policy.id
  window_preference {
    #Required
    days_of_week {
      #Required
      name = var.scheduling_policy_scheduling_window_window_preference_days_of_week_name
    }
    duration             = var.scheduling_policy_scheduling_window_window_preference_duration
    is_enforced_duration = var.scheduling_policy_scheduling_window_window_preference_is_enforced_duration
    start_time           = var.scheduling_policy_scheduling_window_window_preference_start_time
    weeks_of_month       = var.scheduling_policy_scheduling_window_window_preference_weeks_of_month

    #Optional
    months {
      #Required
      name = var.scheduling_policy_scheduling_window_window_preference_months_name1
    }
    months {
      #Required
      name = var.scheduling_policy_scheduling_window_window_preference_months_name2
    }
    months {
      #Required
      name = var.scheduling_policy_scheduling_window_window_preference_months_name3
    }
    months {
      #Required
      name = var.scheduling_policy_scheduling_window_window_preference_months_name4
    }
  }

  #Optional
  compartment_id = var.compartment_id
  defined_tags   = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.scheduling_policy_scheduling_window_defined_tags_value)
  freeform_tags  = var.scheduling_policy_scheduling_window_freeform_tags
}

data "oci_database_scheduling_policy_scheduling_windows" "test_scheduling_policy_scheduling_windows" {
  #Required
  scheduling_policy_id = oci_database_scheduling_policy.test_scheduling_policy.id

  #Optional
  compartment_id = var.compartment_id
  display_name   = var.scheduling_policy_scheduling_window_display_name
  state          = var.scheduling_policy_scheduling_window_state
}