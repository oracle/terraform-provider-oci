// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "data_safe_target_ocid" {}
variable "data_safe_discovery_job_ocid" {}
variable "data_safe_discovery_job_result_key" {}

variable "discovery_jobs_result_column_name" {
  default = []
}

variable "discovery_jobs_result_discovery_type" {
  default = "NEW"
}

variable "discovery_jobs_result_is_result_applied" {
  default = false
}

variable "discovery_jobs_result_planned_action" {
  default = "NONE"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_data_safe_discovery_jobs_results" "test_discovery_jobs_results" {
  #Required
  discovery_job_id = var.data_safe_discovery_job_ocid

  #Optional
  discovery_type    = var.discovery_jobs_result_discovery_type
  is_result_applied = var.discovery_jobs_result_is_result_applied
}

