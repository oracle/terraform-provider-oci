// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}

variable "alert_analytic_group_by" {
  default = ["targetIds"]
}

variable "alert_analytic_summary_field" {
  default = []
}

variable "alert_analytic_time_ended" {
  default = "2022-01-31T16:02:08.000Z"
}

variable "alert_analytic_time_started" {
  default = "2022-01-30T16:02:08.000Z"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_data_safe_alert_analytic" "test_alert_analytic" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  group_by                  = var.alert_analytic_group_by
  summary_field             = var.alert_analytic_summary_field
  time_ended                = var.alert_analytic_time_ended
  time_started              = var.alert_analytic_time_started
}

