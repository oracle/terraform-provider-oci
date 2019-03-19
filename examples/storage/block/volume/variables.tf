// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}

variable "volume_backup_policy_display_name" {
  default = "custom"
}

variable "volume_backup_policy_freeform_tags" {
  default = {
    "Department" = "Finance"
  }
}

variable "volume_backup_policy_schedules_backup_type" {
  default = "INCREMENTAL"
}

variable "volume_backup_policy_schedules_offset_seconds" {
  default = 46800
}

variable "volume_backup_policy_schedules_period" {
  default = "ONE_WEEK"
}

variable "volume_backup_policy_schedules_retention_seconds" {
  default = 2592000
}

variable "volume_backup_policy_schedules_time_zone" {
  default = "UTC"
}

variable "DBSize" {
  default = "50" // size in GBs, min: 50, max 16384
}
