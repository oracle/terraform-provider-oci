// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

variable "fusion_environment_refresh_activity_display_name" {
  default = "Manual_Execute_Refresh_2022-12-02T03:48:18.932Z"
}

variable "fusion_environment_refresh_activity_state" {
  default = "ACCEPTED"
}

variable "fusion_environment_refresh_activity_time_expected_finish_less_than_or_equal_to" {
  default = "2006-01-02T15:04:05Z"
}

variable "fusion_environment_refresh_activity_time_scheduled_start_greater_than_or_equal_to" {
  default = "2006-01-02T15:04:05Z"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_fusion_apps_fusion_environment_refresh_activity" "test_fusion_environment_refresh_activity" {
  #Required
  fusion_environment_id        = oci_fusion_apps_fusion_environment.test_fusion_environment.id
  source_fusion_environment_id = oci_fusion_apps_fusion_environment.test_fusion_environment.id
}

data "oci_fusion_apps_fusion_environment_refresh_activities" "test_fusion_environment_refresh_activities" {
  #Required
  fusion_environment_id = oci_fusion_apps_fusion_environment.test_fusion_environment.id

  #Optional
  display_name                                  = var.fusion_environment_refresh_activity_display_name
  state                                         = var.fusion_environment_refresh_activity_state
  time_expected_finish_less_than_or_equal_to    = var.fusion_environment_refresh_activity_time_expected_finish_less_than_or_equal_to
  time_scheduled_start_greater_than_or_equal_to = var.fusion_environment_refresh_activity_time_scheduled_start_greater_than_or_equal_to
}
