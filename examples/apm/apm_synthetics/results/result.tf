// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "apm_domain_id" {}
variable "monitor_id" {}

variable "result_execution_time" {
  default = "executionTime"
}

variable "result_result_content_type" {
  default = "resultContentType"
}

variable "result_result_type" {
  default = "resultType"
}

variable "result_vantage_point" {
  default = "vantagePoint"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
  apm_domain_id    = var.apm_domain_id
  monitor_id       = var.monitor_id
}

data "oci_apm_synthetics_results" "test_results" {
  #Required
  apm_domain_id       = var.apm_domain_id
  execution_time      = var.result_execution_time
  monitor_id          = var.monitor_id
  result_content_type = var.result_result_content_type
  result_type         = var.result_result_type
  vantage_point       = var.result_vantage_point
}

