// Copyright (c) 2017, 2025, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "apm_domain_id" {}

variable "log_log_key" {
  default = "logKey"
}

variable "log_time_log_ended_less_than" {
  default = "2025-05-23T00:00:00Z"
}

variable "log_time_log_started_greater_than_or_equal_to" {
  default = "2025-05-18T00:00:00Z"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_apm_traces_log" "test_log" {
  #Required
  apm_domain_id                             = var.apm_domain_id
  log_key                                   = var.log_log_key
  time_log_ended_less_than                  = var.log_time_log_ended_less_than
  time_log_started_greater_than_or_equal_to = var.log_time_log_started_greater_than_or_equal_to
}
