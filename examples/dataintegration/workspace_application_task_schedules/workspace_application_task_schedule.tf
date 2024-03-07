// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0


variable "workspace_application_task_schedule_application_key" {
  default = ""
}

variable "workspace_application_task_schedule_auth_mode" {
  default = "RESOURCE_PRINCIPAL"
}

variable "workspace_application_task_schedule_config_provider_delegate" {
  default = "{\"bindings\":{\"PARAMETER_20240118_172132\":{\"simpleValue\":\"abcdef\"}},\"childProviders\":{}}"
}

variable "workspace_application_task_schedule_description" {
  default = "description"
}

variable "workspace_application_task_schedule_end_time_millis" {
  default = 10
}

variable "workspace_application_task_schedule_expected_duration" {
  default = 1.0
}

variable "workspace_application_task_schedule_expected_duration_unit" {
  default = "SECONDS"
}

variable "workspace_application_task_schedule_identifier" {
  default = "REST_VIVEK_20240207_111223_3365404"
}

variable "workspace_application_task_schedule_identifier_array" {
  default = ["REST_TASK_20240115_003958_20240116_003525_525100"]
  }


variable "workspace_application_task_schedule_is_backfill_enabled" {
  default = false
}

variable "workspace_application_task_schedule_is_concurrent_allowed" {
  default = false
}

variable "workspace_application_task_schedule_is_enabled" {
  default = false
}

variable "workspace_application_task_schedule_key" {
  default = "key"
}

variable "workspace_application_task_schedule_model_version" {
  default = "20210408"
}

variable "workspace_application_task_schedule_name" {
  default = "REST_VIVEK_20240207_111223_3365404"
}

variable "workspace_application_task_workspace_id" {
  default = ""
}

variable "workspace_application_task_schedule_number_of_retries" {
  default = 2
}

variable "workspace_application_task_schedule_object_status" {
  default = 8
}

variable "workspace_application_task_schedule_object_version" {
  default = 1
}

variable "workspace_application_task_schedule_registry_metadata_aggregator_key" {
  default = ""
}

variable "workspace_application_task_schedule_registry_metadata_is_favorite" {
  default = false
}

variable "workspace_application_task_schedule_registry_metadata_registry_version" {
  default = 0
}

variable "workspace_application_task_schedule_retry_delay" {
  default = 1
}

variable "workspace_application_task_schedule_retry_delay_unit" {
  default = "SECONDS"
}

variable "workspace_application_task_schedule_schedule_ref_description" {
  default = "description2"
}

variable "workspace_application_task_schedule_schedule_ref_identifier" {
  default = "SCHEDULE_TEMP"
}

variable "workspace_application_task_schedule_schedule_ref_key" {
  default = ""
}

variable "workspace_application_task_schedule_schedule_ref_model_type" {
  default = "SCHEDULE"
}

variable "workspace_application_task_schedule_schedule_ref_name" {
  default = "SCHEDULE_TEMP"
}

variable "workspace_application_task_schedule_schedule_ref_object_status" {
  default = 10
}

variable "workspace_application_task_schedule_schedule_ref_timezone" {
  default = "GMT"
}

resource "oci_dataintegration_workspace_application_task_schedule" "test_workspace_application_task_schedule" {
  #Required
  application_key = var.workspace_application_task_schedule_application_key
  identifier      = var.workspace_application_task_schedule_identifier
  name            = var.workspace_application_task_schedule_name
  workspace_id    = var.workspace_application_task_workspace_id

  #Optional
  auth_mode                = var.workspace_application_task_schedule_auth_mode
  config_provider_delegate = var.workspace_application_task_schedule_config_provider_delegate
  description              = var.workspace_application_task_schedule_description
  expected_duration        = var.workspace_application_task_schedule_expected_duration
  expected_duration_unit   = var.workspace_application_task_schedule_expected_duration_unit
  is_backfill_enabled      = var.workspace_application_task_schedule_is_backfill_enabled
  is_concurrent_allowed    = var.workspace_application_task_schedule_is_concurrent_allowed
  is_enabled               = var.workspace_application_task_schedule_is_enabled
  model_version            = var.workspace_application_task_schedule_model_version
  number_of_retries        = var.workspace_application_task_schedule_number_of_retries
  object_status            = var.workspace_application_task_schedule_object_status
  registry_metadata {

    #Optional
    aggregator_key   = var.workspace_application_task_schedule_registry_metadata_aggregator_key
    is_favorite      = var.workspace_application_task_schedule_registry_metadata_is_favorite
    registry_version = var.workspace_application_task_schedule_registry_metadata_registry_version
  }
  retry_delay      = var.workspace_application_task_schedule_retry_delay
  retry_delay_unit = var.workspace_application_task_schedule_retry_delay_unit
  schedule_ref {

    #Optional
    identifier                     = var.workspace_application_task_schedule_schedule_ref_identifier
    key                            = var.workspace_application_task_schedule_schedule_ref_key
    model_type     = var.workspace_application_task_schedule_schedule_ref_model_type
    name           = var.workspace_application_task_schedule_schedule_ref_name
    timezone = var.workspace_application_task_schedule_schedule_ref_timezone
  }
}

data "oci_dataintegration_workspace_application_task_schedules" "test_workspace_application_task_schedules" {
  #Required
  application_key = var.workspace_application_task_schedule_application_key
  workspace_id    = var.workspace_application_task_workspace_id

  #Optional
  identifier = var.workspace_application_task_schedule_identifier_array
  is_enabled = var.workspace_application_task_schedule_is_enabled
  name       = var.workspace_application_task_schedule_name
}
