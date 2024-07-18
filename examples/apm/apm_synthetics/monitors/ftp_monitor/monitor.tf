// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "apm_domain_id" {}

variable "apm_domain_description" {
  default = "description"
}

variable "apm_domain_display_name" {
  default = "displayName"
}

variable "apm_domain_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "apm_domain_is_free_tier" {
  default = false
}

variable "apm_domain_state" {
  default = "ACTIVE"
}

variable "data_key_data_key_type" {
  default = "PRIVATE"
}

variable "monitor_configuration_config_type" {
  default = "FTP_CONFIG"
}

variable "monitor_configuration_query" {
  default = "query"
}

variable "monitor_configuration_is_failure_retried" {
  default = false
}

variable "monitor_configuration_ftp_protocol" {
  default = "FTP"
}

variable "monitor_configuration_ftp_request_type" {
  default = "LIST"
}

variable "monitor_configuration_ftp_basic_authentication_details_username" {
  default = "username"
}

variable "monitor_configuration_ftp_basic_authentication_details_password_password_type" {
  default = "IN_TEXT"
}

variable "monitor_configuration_ftp_basic_authentication_details_password_password" {
  default = "BEstrO0ng_#11"
}

variable "monitor_display_name" {
  default = "displayName"
}

variable "monitor_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "monitor_monitor_type" {
  default = "FTP"
}

variable "monitor_repeat_interval_in_seconds" {
  default = 600
}

variable "monitor_is_run_once" {
  default = false
}

variable "monitor_is_run_now" {
  default = false
}

variable "monitor_scheduling_policy" {
  default = "ALL"
}

variable "monitor_status" {
  default = "ENABLED"
}

variable "monitor_target" {
  default = "www.oracle.com:80"
}

variable "monitor_timeout_in_seconds" {
  default = 60
}

variable "monitor_tag_name" {
  default =  "tagName"
}

variable "monitor_tag_value" {
  default =  "tagValue"
}

variable "monitor_availability_configuration_max_allowed_failures_per_interval" {
  default = 0
}

variable "monitor_availability_configuration_min_allowed_runs_per_interval" {
  default = 1
}

variable "monitor_maintenance_window_schedule_time_ended" {
  default = "2025-02-12T22:47:12.613Z"
}

variable "monitor_maintenance_window_schedule_time_started" {
  default = "2024-12-18T22:47:12.654Z"
}

variable "monitor_vantage_points_name" {
  default = "us-phoenix-internal"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_apm_synthetics_monitor" "test_monitor" {
  #Required
  apm_domain_id              = oci_apm_apm_domain.test_apm_domain.id
  display_name               = var.monitor_display_name
  monitor_type               = var.monitor_monitor_type
  repeat_interval_in_seconds = var.monitor_repeat_interval_in_seconds
  vantage_points {
    #Required
    name  = var.monitor_vantage_points_name
  }

  configuration {

    config_type                       = var.monitor_configuration_config_type
    is_failure_retried                = var.monitor_configuration_is_failure_retried
    ftp_protocol                      = var.monitor_configuration_ftp_protocol
    ftp_request_type                  = var.monitor_configuration_ftp_request_type

    ftp_basic_authentication_details {
      username           = var.monitor_configuration_ftp_basic_authentication_details_username
      password {
        password         = var.monitor_configuration_ftp_basic_authentication_details_password_password
        password_type    = var.monitor_configuration_ftp_basic_authentication_details_password_password_type
      }
    }

  }
  freeform_tags = var.monitor_freeform_tags

  status             = var.monitor_status
  target             = var.monitor_target
  timeout_in_seconds = var.monitor_timeout_in_seconds
  is_run_once        = var.monitor_is_run_once
  is_run_now         = var.monitor_is_run_now
  scheduling_policy  = var.monitor_scheduling_policy
  #Optional
  availability_configuration {
    max_allowed_failures_per_interval  = var.monitor_availability_configuration_max_allowed_failures_per_interval
    min_allowed_runs_per_interval      = var.monitor_availability_configuration_min_allowed_runs_per_interval
  }
  maintenance_window_schedule {
    time_ended   = var.monitor_maintenance_window_schedule_time_ended
    time_started = var.monitor_maintenance_window_schedule_time_started
  }
}

data "oci_apm_synthetics_monitors" "test_monitors" {
  #Required
  apm_domain_id = oci_apm_apm_domain.test_apm_domain.id

  #Optional
  display_name = var.monitor_display_name
  monitor_type = var.monitor_monitor_type
  status       = var.monitor_status
}


resource "oci_apm_apm_domain" "test_apm_domain" {
  #Required
  compartment_id = var.compartment_ocid
  display_name   = var.apm_domain_display_name

  #Optional
  description   = var.apm_domain_description
  freeform_tags = var.apm_domain_freeform_tags
  is_free_tier  = var.apm_domain_is_free_tier
}


