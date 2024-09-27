// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "desktop_pool_are_privileged_users" {
  description = ""
  default = false
}

variable "desktop_pool_description" {
  description = ""
  default = "Pool Description"
}

variable "desktop_pool_contact_details" {
  description = ""
  default = ""
}

variable "desktop_pool_device_policy_audio_mode" {
  description = ""
  default = "NONE"
}

variable "desktop_pool_device_policy_cdm_mode" {
  description = ""
  default = "NONE"
}

variable "desktop_pool_device_policy_clipboard_mode" {
  description = ""
  default = "NONE"
}

variable "desktop_pool_device_policy_is_display_enabled" {
  description = ""
  default = false
}

variable "desktop_pool_device_policy_is_keyboard_enabled" {
  description = ""
  default = false
}

variable "desktop_pool_device_policy_is_pointer_enabled" {
  description = ""
  default = false
}

variable "desktop_pool_device_policy_is_printing_enabled" {
  description = ""
  default = false
}

variable "desktop_pool_display_name" {
  description = ""
  default = "testPool1"
}

variable "desktop_pool_freeform_tags" {
  description = ""
  type = map(string)

  default = {
    Test = "Test"
  }
}

variable "desktop_pool_standby_size" {
  description = ""
  default = 2
}

variable "desktop_pool_maximum_size" {
  description = ""
  default = 10
}

variable "desktop_pool_nsg_ids" {
  description = ""
  type    = set(string)
  default = null
}

variable "desktop_pool_is_storage_enabled" {
  description = ""
  default = true
}

variable "desktop_pool_backup_policy_id" {
  description = ""
  default = ""
}

variable "desktop_pool_storage_size_in_gbs" {
  description = ""
  default = 50
}

variable "desktop_pool_vcn_id" {
  description = ""
}

variable "desktop_pool_subnet_id" {
  description = ""
}

variable "desktop_pool_shape_name" {
  description = ""
  default = "Flex Low"
}

variable "desktop_pool_image_id" {
  description = ""
}

variable "desktop_pool_image_name" {
  description = ""
}

variable "desktop_pool_start_schedule_cron_expr" {
  description = ""
  default = "0 10 8 ? * 1"
}

variable "desktop_pool_start_schedule_timezone" {
  description = ""
  default = "MST"
}

variable "desktop_pool_stop_schedule_cron_expr" {
  description = ""
  default = "0 20 18 ? * 5"
}

variable "desktop_pool_stop_schedule_timezone" {
  description = ""
  default = "MST"
}

variable "desktop_pool_session_lifecycle_actions_disconnect_action" {
  description = ""
  default = "NONE"
}

variable "desktop_pool_session_lifecycle_actions_inactivity_action" {
  description = ""
  default = "NONE"
}

variable "desktop_pool_session_lifecycle_actions_disconnect_grace_period_in_minutes" {
  description = ""
  default = 60
}

variable "desktop_pool_session_lifecycle_actions_inactivity_grace_period_in_minutes" {
  description = ""
  default = 60
}


variable "desktop_pool_state" {
  description = ""
  default = "ACTIVE"
}
