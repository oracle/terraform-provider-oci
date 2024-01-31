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

variable "managed_database_id" {
   default = "<database.ocid>"
}

variable "managed_database_optimizer_statistics_advisor_execution_script_execution_name" {
  default = "EXEC_9584"
}

variable "managed_database_optimizer_statistics_advisor_execution_script_task_name" {
  default = "AUTO_STATS_ADVISOR_TASK"
}

variable "managed_database_optimizer_statistics_collection_aggregation_group_type" {
  default = "TASK_STATUS"
}

variable "managed_database_optimizer_statistics_advisor_execution_end_time_less_than_or_equal_to" {
  #default = "2023-11-27T05:49:00Z"
}

variable "managed_database_optimizer_statistics_advisor_execution_start_time_greater_than_or_equal_to" {
  #default = "2023-11-27T05:49:00Z"
}

variable "managed_database_optimizer_statistics_collection_aggregation_end_time_less_than_or_equal_to" {
  #default = "2023-11-27T05:49:00Z"
}

variable "managed_database_optimizer_statistics_collection_aggregation_start_time_greater_than_or_equal_to" {
  #default = "2023-11-27T05:49:00Z"
}

variable "managed_database_optimizer_statistics_collection_aggregation_task_type" {
  default = "ALL"
}

variable "managed_database_optimizer_statistics_collection_operation_end_time_less_than_or_equal_to" {
  #default = "2023-11-27T05:49:00Z"
}

variable "managed_database_optimizer_statistics_collection_operation_filter_by" {
  #default = "filterBy"
}

variable "managed_database_optimizer_statistics_collection_operation_start_time_greater_than_or_equal_to" {
  #default = "2023-11-27T05:49:00Z"
}

variable "managed_database_optimizer_statistics_collection_operation_task_type" {
  default = "ALL"
}

data "oci_database_management_managed_database_table_statistics" "test_managed_database_table_statistics" {
  #Required
  managed_database_id = var.managed_database_id
}

data "oci_database_management_managed_database_optimizer_statistics_advisor_executions" "test_managed_database_optimizer_statistics_advisor_executions" {
  #Required
  managed_database_id = var.managed_database_id

  #Optional
  end_time_less_than_or_equal_to      = var.managed_database_optimizer_statistics_advisor_execution_end_time_less_than_or_equal_to
  start_time_greater_than_or_equal_to = var.managed_database_optimizer_statistics_advisor_execution_start_time_greater_than_or_equal_to
}

data "oci_database_management_managed_database_optimizer_statistics_advisor_execution_script" "test_managed_database_optimizer_statistics_advisor_execution_scripts" {
  #Required
  execution_name      = var.managed_database_optimizer_statistics_advisor_execution_script_execution_name
  managed_database_id = var.managed_database_id
  task_name           = var.managed_database_optimizer_statistics_advisor_execution_script_task_name
}

data "oci_database_management_managed_database_optimizer_statistics_collection_aggregations" "test_managed_database_optimizer_statistics_collection_aggregations" {
  #Required
  group_type          = var.managed_database_optimizer_statistics_collection_aggregation_group_type
  managed_database_id = var.managed_database_id

  #Optional
  end_time_less_than_or_equal_to      = var.managed_database_optimizer_statistics_collection_aggregation_end_time_less_than_or_equal_to
  start_time_greater_than_or_equal_to = var.managed_database_optimizer_statistics_collection_aggregation_start_time_greater_than_or_equal_to
  task_type                           = var.managed_database_optimizer_statistics_collection_aggregation_task_type
}

data "oci_database_management_managed_database_optimizer_statistics_collection_operations" "test_managed_database_optimizer_statistics_collection_operations" {
  #Required
  managed_database_id = var.managed_database_id

  #Optional
  end_time_less_than_or_equal_to      = var.managed_database_optimizer_statistics_collection_operation_end_time_less_than_or_equal_to
  filter_by                           = var.managed_database_optimizer_statistics_collection_operation_filter_by
  start_time_greater_than_or_equal_to = var.managed_database_optimizer_statistics_collection_operation_start_time_greater_than_or_equal_to
  task_type                           = var.managed_database_optimizer_statistics_collection_operation_task_type
}