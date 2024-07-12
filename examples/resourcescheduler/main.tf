// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "schedule_action" {
  default = "START_RESOURCE"
}

variable "schedule_defined_tags_value" {
  default = "value"
}

variable "schedule_description" {
  default = "description"
}

variable "schedule_display_name" {
  default = "displayName"
}

variable "schedule_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "schedule_recurrence_details" {
  default = "FREQ=DAILY;INTERVAL=1"
}

variable "schedule_recurrence_type" {
  default = "ICAL"
}

variable "schedule_resource_filters_attribute" {
  default = "COMPARTMENT_ID"
}

variable "schedule_resource_filters_condition" {
  default = "EQUAL"
}

variable "schedule_resource_filters_should_include_child_compartments" {
  default = false
}

variable "schedule_resource_filters_value_namespace" {
  default = "namespace"
}

variable "schedule_resource_filters_value_tag_key" {
  default = "tagKey"
}

variable "schedule_resource_filters_value_value" {
  default = "value"
}

variable "schedule_resources_id" {
  default = "id"
}

variable "schedule_resources_metadata" {
  default = "metadata"
}

variable "schedule_state" {
  default = "ACTIVE"
}

variable "schedule_time_ends" {
  default = "2024-07-23T17:45:44.408Z"
}

variable "schedule_time_starts" {
  default = "2024-07-13T17:45:44.408Z"
}


provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_resource_scheduler_schedule" "test_schedule" {
  #Required
  action             = var.schedule_action
  compartment_id     = var.compartment_id
  recurrence_details = var.schedule_recurrence_details
  recurrence_type    = var.schedule_recurrence_type

  resource_filters {
    # Required
    attribute = "DEFINED_TAGS"
    value {
      namespace="ResourceSchedulerCanary"
      tag_key="ScheduleTagFilterTestKey"
      value="foo"
    }
  }
  resource_filters {
    # Required
    attribute = "LIFECYCLE_STATE"
    value {
      value="running"
    }
    value {
      value="stopped"
    }
  }
  resource_filters {
    # Required
    attribute = "COMPARTMENT_ID"
    value {
      value=var.compartment_id
    }
  }

  #Optional
  defined_tags  = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.schedule_defined_tags_value)
  description   = var.schedule_description
  display_name  = var.schedule_display_name
  freeform_tags = var.schedule_freeform_tags
  time_ends   = var.schedule_time_ends
  time_starts = var.schedule_time_starts
}

data "oci_resource_scheduler_schedules" "test_schedules" {

  #Optional
  compartment_id = var.compartment_id
  display_name   = var.schedule_display_name
  schedule_id    = oci_resource_scheduler_schedule.test_schedule.id
  state          = var.schedule_state
}
