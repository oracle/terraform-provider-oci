// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
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

variable "compartment_id" {  default = "<compartment.ocid>"}

variable "managed_database_id" {
   default = "<database.ocid>"
}

variable "sta_te_sql_execution_id" {
  default = "1"
}

variable "sta_te_sql_object_id" {
  default = "1"
}

variable "sta_te_task_id" {
  default = "1"
}
variable "managed_database_sql_tuning_advisor_task_name" {
  default = "name"
}

variable "managed_database_sql_tuning_advisor_task_status" {
  default = "INITIAL"
}

variable "managed_database_sql_tuning_advisor_task_time_greater_than_or_equal_to" {
  #default = "2023-11-27T05:49:00Z"
}

variable "managed_database_sql_tuning_advisor_task_time_less_than_or_equal_to" {
  #default = "2023-11-27T05:49:00Z"
}

variable "managed_database_sql_tuning_advisor_tasks_finding_finding_filter" {
  default = "none"
}

variable "managed_database_sql_tuning_advisor_tasks_finding_index_hash_filter" {
  default = "indexHashFilter"
}

variable "managed_database_sql_tuning_advisor_tasks_finding_search_period" {
  default = "LAST_24HR"
}

variable "managed_database_sql_tuning_advisor_tasks_finding_stats_hash_filter" {
  default = "statsHashFilter"
}

variable "managed_database_sql_tuning_advisor_tasks_sql_execution_plan_attribute" {
  default = "ORIGINAL"
}

variable "managed_database_sql_tuning_advisor_tasks_summary_report_begin_exec_id_greater_than_or_equal_to" {
  default = 1
}

variable "managed_database_sql_tuning_advisor_tasks_summary_report_end_exec_id_less_than_or_equal_to" {
  default = 10
}

variable "managed_database_sql_tuning_advisor_tasks_summary_report_search_period" {
  default = "LAST_24HR"
}

variable "managed_database_sql_tuning_advisor_tasks_summary_report_time_greater_than_or_equal_to" {
  #default = "timeGreaterThanOrEqualTo"
}

variable "managed_database_sql_tuning_advisor_tasks_summary_report_time_less_than_or_equal_to" {
  #default = "timeLessThanOrEqualTo"
}


data "oci_database_management_managed_database_sql_tuning_advisor_tasks" "test_managed_database_sql_tuning_advisor_tasks" {
  #Required
  managed_database_id = var.managed_database_id

  #Optional
  name                          = var.managed_database_sql_tuning_advisor_task_name
  status                        = var.managed_database_sql_tuning_advisor_task_status
  time_greater_than_or_equal_to = var.managed_database_sql_tuning_advisor_task_time_greater_than_or_equal_to
  time_less_than_or_equal_to    = var.managed_database_sql_tuning_advisor_task_time_less_than_or_equal_to
}


data "oci_database_management_managed_database_sql_tuning_advisor_tasks_execution_plan_stats_comparision" "test_managed_database_sql_tuning_advisor_tasks_execution_plan_stats_comparision" {
  #Required
  execution_id               = var.sta_te_sql_execution_id
  managed_database_id        = var.managed_database_id
  sql_object_id              = var.sta_te_sql_object_id
  sql_tuning_advisor_task_id = var.sta_te_task_id
}

data "oci_database_management_managed_database_sql_tuning_advisor_tasks_findings" "test_managed_database_sql_tuning_advisor_tasks_findings" {
  #Required
  managed_database_id        = var.managed_database_id
  sql_tuning_advisor_task_id = var.sta_te_task_id

  #Optional
  begin_exec_id     = 1
  end_exec_id       = 1
  finding_filter    = var.managed_database_sql_tuning_advisor_tasks_finding_finding_filter
  index_hash_filter = var.managed_database_sql_tuning_advisor_tasks_finding_index_hash_filter
  search_period     = var.managed_database_sql_tuning_advisor_tasks_finding_search_period
  stats_hash_filter = var.managed_database_sql_tuning_advisor_tasks_finding_stats_hash_filter
}

data "oci_database_management_managed_database_sql_tuning_advisor_tasks_recommendations" "test_managed_database_sql_tuning_advisor_tasks_recommendations" {
  #Required
  execution_id               = var.sta_te_sql_execution_id
  managed_database_id        = var.managed_database_id
  sql_object_id              = var.sta_te_sql_object_id
  sql_tuning_advisor_task_id = var.sta_te_task_id
}

data "oci_database_management_managed_database_sql_tuning_advisor_tasks_sql_execution_plan" "test_managed_database_sql_tuning_advisor_tasks_sql_execution_plan" {
  #Required
  attribute                  = var.managed_database_sql_tuning_advisor_tasks_sql_execution_plan_attribute
  managed_database_id        = var.managed_database_id
  sql_object_id              = var.sta_te_sql_object_id
  sql_tuning_advisor_task_id = var.sta_te_task_id
}

data "oci_database_management_managed_database_sql_tuning_advisor_tasks_summary_report" "test_managed_database_sql_tuning_advisor_tasks_summary_report" {
  #Required
  managed_database_id        = var.managed_database_id
  sql_tuning_advisor_task_id = var.sta_te_task_id

  #Optional
  #begin_exec_id_greater_than_or_equal_to = var.managed_database_sql_tuning_advisor_tasks_summary_report_begin_exec_id_greater_than_or_equal_to
  #end_exec_id_less_than_or_equal_to      = var.managed_database_sql_tuning_advisor_tasks_summary_report_end_exec_id_less_than_or_equal_to
  search_period                          = var.managed_database_sql_tuning_advisor_tasks_summary_report_search_period
  time_greater_than_or_equal_to          = var.managed_database_sql_tuning_advisor_tasks_summary_report_time_greater_than_or_equal_to
  time_less_than_or_equal_to             = var.managed_database_sql_tuning_advisor_tasks_summary_report_time_less_than_or_equal_to
}