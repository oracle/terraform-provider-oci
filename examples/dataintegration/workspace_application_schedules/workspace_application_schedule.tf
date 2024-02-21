// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0



variable "workspace_application_schedule_application_key" {
  default = ""
}

variable "workspace_application_schedule_description" {
  default = "description"
}

variable "workspace_application_schedule_frequency_details_custom_expression" {
  default = "customExpression"
}

variable "workspace_application_schedule_frequency_details_day_of_week" {
  default = "SUNDAY"
}

variable "workspace_application_schedule_frequency_details_days" {
  default = []
}

variable "workspace_application_schedule_frequency_details_frequency" {
  default = "HOURLY"
}

variable "workspace_application_schedule_frequency_details_interval" {
  default = 1
}

variable "workspace_application_schedule_frequency_details_model_type" {
  default = "HOURLY"
}

variable "workspace_application_schedule_frequency_details_time_hour" {
  default = 0
}

variable "workspace_application_schedule_frequency_details_time_minute" {
  default = 10
}

variable "workspace_application_schedule_frequency_details_time_second" {
  default = 10
}

variable "workspace_application_schedule_frequency_details_week_of_month" {
  default = "FIRST"
}

variable "workspace_application_identifier" {
  default = "APPLICATION_TF_TEST_1"
}

variable "workspace_application_schedule_identifier" {
  default = ["TERSI_TEST_SCHEDULE001"]
}

variable "workspace_application_schedule_is_daylight_adjustment_enabled" {
  default = false
}

variable "workspace_application_schedule_key" {
  default = "key"
}

variable "workspace_application_schedule_model_version" {
  default = "20210409"
}

variable "workspace_application_schedule_name" {
  default = "TERSI_TEST_SCHEDULE007"
}

variable "workspace_application_schedule_object_status" {
  default = 0
}

variable "workspace_application_schedule_object_version" {
  default = 0
}

variable "workspace_application_schedule_registry_metadata_aggregator_key" {
  default = "aggregatorKey"
}

variable "workspace_application_schedule_registry_metadata_is_favorite" {
  default = false
}

variable "workspace_application_schedule_registry_metadata_key" {
  default = "key"
}

variable "workspace_application_schedule_registry_metadata_labels" {
  default = ["temp_LABEL"]
}

variable "workspace_application_schedule_registry_metadata_registry_version" {
  default = 10
}

variable "workspace_application_schedule_timezone" {
  default = "UTC"
}

variable "workspace_application_schedule_type" {
  default = ["SCHEDULE"]
}
variable "workspace_application_workspace_id"{
    default = ""
}

resource "oci_dataintegration_workspace_application_schedule" "test_workspace_application_schedule" {
  #Required
  application_key = var.workspace_application_schedule_application_key
  identifier      = var.workspace_application_identifier
  name            = var.workspace_application_schedule_name
  workspace_id    = var.workspace_application_workspace_id

  #Optional
  description = var.workspace_application_schedule_description
  frequency_details {
    #Required
    model_type = var.workspace_application_schedule_frequency_details_model_type

    #Optional
    #custom_expression = var.workspace_application_schedule_frequency_details_custom_expression
    day_of_week       = var.workspace_application_schedule_frequency_details_day_of_week
    days              = var.workspace_application_schedule_frequency_details_days
    frequency         = var.workspace_application_schedule_frequency_details_frequency
    interval          = var.workspace_application_schedule_frequency_details_interval
    time {

      #Optional
      hour   = var.workspace_application_schedule_frequency_details_time_hour
      minute = var.workspace_application_schedule_frequency_details_time_minute
      second = var.workspace_application_schedule_frequency_details_time_second
    }
    #week_of_month = var.workspace_application_schedule_frequency_details_week_of_month
  }
  is_daylight_adjustment_enabled = var.workspace_application_schedule_is_daylight_adjustment_enabled
  #key                            = var.workspace_application_schedule_key
  model_version                  = var.workspace_application_schedule_model_version
  object_status                  = var.workspace_application_schedule_object_status
  object_version                 = var.workspace_application_schedule_object_version

  timezone = var.workspace_application_schedule_timezone
}

data "oci_dataintegration_workspace_application_schedules" "test_workspace_application_schedules" {
  #Required
  application_key = var.workspace_application_schedule_application_key
  workspace_id    = var.workspace_application_workspace_id

  #Optional
  identifier = var.workspace_application_schedule_identifier
  #key        = var.workspace_application_schedule_key
  name       = var.workspace_application_schedule_name
  type       = var.workspace_application_schedule_type
}
