// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}



provider "oci" {
  tenancy_ocid = var.tenancy_ocid
  user_ocid = var.user_ocid
  fingerprint = var.fingerprint
  private_key_path = var.private_key_path
  region = var.region
}

####################### SQL Plan Management #########################

variable "compartment_id" {  default = "compartment.ocid"}

variable "managed_database_id" {
   default = "<database.ocid>"
}

variable "managed_database_sql_plan_baseline_is_accepted" {
  default = false
}

variable "managed_database_sql_plan_baseline_is_adaptive" {
  default = false
}

variable "managed_database_sql_plan_baseline_is_enabled" {
  default = true
}

variable "managed_database_sql_plan_baseline_is_fixed" {
  default = false
}

variable "managed_database_sql_plan_baseline_is_reproduced" {
  default = false
}

variable "managed_database_sql_plan_baseline_origin" {
  default = "AUTO_CAPTURE"
}

variable "managed_database_sql_plan_baseline_plan_name" {
  default = "planName"
}

variable "managed_database_sql_plan_baseline_sql_handle" {
  default = "sqlHandle"
}

variable "managed_database_sql_plan_baseline_sql_text" {
  default = "sqlText"
}

variable "managed_database_sql_plan_baseline_job_name" {
  default = "TestJobName"
}

# Get SQL Plan Baseline configuration details for the managed database
data "oci_database_management_managed_database_sql_plan_baseline_configuration" "test_managed_database_sql_plan_baseline_configuration" {
  #Required
  managed_database_id = var.managed_database_id
}

# List SQL Plan Baselines
data "oci_database_management_managed_database_sql_plan_baselines" "test_managed_database_sql_plan_baselines" {
  #Required
  managed_database_id = var.managed_database_id

  #Optional
  origin        = var.managed_database_sql_plan_baseline_origin
  plan_name     = var.managed_database_sql_plan_baseline_plan_name
  is_enabled    = var.managed_database_sql_plan_baseline_is_enabled
  is_accepted   = var.managed_database_sql_plan_baseline_is_accepted
  is_adaptive   = var.managed_database_sql_plan_baseline_is_adaptive
  is_fixed      = var.managed_database_sql_plan_baseline_is_fixed
  is_reproduced = var.managed_database_sql_plan_baseline_is_reproduced
  sql_handle    = var.managed_database_sql_plan_baseline_sql_handle
  sql_text      = var.managed_database_sql_plan_baseline_sql_text
}

data "oci_database_management_managed_database_sql_plan_baseline_jobs" "test_managed_database_sql_plan_baseline_jobs" {
  #Required
  managed_database_id = var.managed_database_id

  #Optional
  name = var.managed_database_sql_plan_baseline_job_name
}

data "oci_database_management_managed_database_cursor_cache_statements" "test_managed_database_cursor_cache_statements" {
  #Required
  managed_database_id = var.managed_database_id

  #Optional
  sql_text = var.managed_database_sql_plan_baseline_sql_text
}


