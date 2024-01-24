// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "java_download_report_id" {
  default = "id"
}

variable "java_download_report_format" {
  default = "CSV"
}

variable "java_download_report_state" {
  default = "ACTIVE"
}

variable "java_download_report_time_end" {
  default = "2024-08-01T03:07:27Z"
}

variable "java_download_report_time_start" {
  default = "2023-08-01T03:07:27Z"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_jms_java_downloads_java_download_report" "test_java_download_report_content_data" {
  #Required
  compartment_id = var.tenancy_ocid
  format         = var.java_download_report_format

  #Optional
  time_end   = var.java_download_report_time_end
  time_start = var.java_download_report_time_start
}

data "oci_jms_java_downloads_java_download_reports" "test_java_download_reports" {
  #Required
  compartment_id = var.tenancy_ocid

  #Optional
  java_download_report_id = var.java_download_report_id
  state                   = var.java_download_report_state
}

