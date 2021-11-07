// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "managed_database_group_description" {
  default = "Sales test database group"
}

variable "managed_database_group_id" {
  default = "id"
}

variable "managed_database_group_name" {
  default = "TestGroup"
}

variable "managed_database_group_state" {
  default = "ACTIVE"
}

variable "managed_database_id" {
  default = "testManagedDatabase0"
}

variable "managed_databases_database_parameter_credentials_username" {
  default = "sys"
}

variable "managed_databases_database_parameter_credentials_password" {
  default = "sys"
}

variable "managed_databases_database_parameter_credentials_role" {
  default = "NORMAL"
}

variable "managed_databases_database_parameter_parameters_name" {
  default = "open_cursors"
}

variable "managed_databases_database_parameter_parameters_value" {
  default = "305"
}

variable "managed_databases_database_parameter_update_comment" {
  default = "Terraform update of open cursors"
}

variable "managed_databases_database_parameter_scope" {
  default = "BOTH"
}

variable "managed_databases_database_parameter_is_allowed_values_included" {
  default = "false"
}

variable "managed_databases_database_parameter_source" {
  default = "CURRENT"
}

variable "db_management_private_endpoint_name" {
  default = "TestPrivateEndpoint"
}

variable "db_management_private_endpoint_description" {
  default = "Test private endpoint"
}

variable "db_management_private_endpoint_state" {
  default = "ACTIVE"
}

variable "db_management_private_endpoint_is_cluster" {
  default = false
}

variable "managed_database_sql_tuning_advisor_task_name" {
  default = "name"
}

variable "managed_database_sql_tuning_advisor_task_status" {
  default = "INITIAL"
}

variable "managed_database_sql_tuning_advisor_task_time_greater_than_or_equal_to" {
  default = "timeGreaterThanOrEqualTo"
}

variable "managed_database_sql_tuning_advisor_task_time_less_than_or_equal_to" {
  default = "timeLessThanOrEqualTo"
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
  default = 10
}

variable "managed_database_sql_tuning_advisor_tasks_summary_report_end_exec_id_less_than_or_equal_to" {
  default = 10
}

variable "managed_database_sql_tuning_advisor_tasks_summary_report_search_period" {
  default = "LAST_24HR"
}

variable "managed_database_sql_tuning_advisor_tasks_summary_report_time_greater_than_or_equal_to" {
  default = "timeGreaterThanOrEqualTo"
}

variable "managed_database_sql_tuning_advisor_tasks_summary_report_time_less_than_or_equal_to" {
  default = "timeLessThanOrEqualTo"
}

provider "oci" {
  tenancy_ocid = var.tenancy_ocid
  user_ocid = var.user_ocid
  fingerprint = var.fingerprint
  private_key_path = var.private_key_path
  region = var.region
}

resource "oci_database_management_managed_database_group" "test_managed_database_group" {
  #Required
  compartment_id = var.compartment_id
  name = var.managed_database_group_name

  #Optional
  description = var.managed_database_group_description
  managed_databases {
    id = var.managed_database_id
  }
}

data "oci_database_management_managed_database_groups" "test_managed_database_groups_with_id" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  id = oci_database_management_managed_database_group.test_managed_database_group.id
  state = var.managed_database_group_state
}

data "oci_database_management_managed_database_groups" "test_managed_database_groups_with_name" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  name = var.managed_database_group_name
  state = var.managed_database_group_state
}

resource "oci_database_management_managed_databases_change_database_parameter" "test_managed_databases_change_database_parameter" {
  #Required
  credentials {

    #Optional
    password = var.managed_databases_database_parameter_credentials_password
    role = var.managed_databases_database_parameter_credentials_role
    user_name = var.managed_databases_database_parameter_credentials_username
  }
  managed_database_id = var.managed_database_id
  parameters {
    #Required
    name = var.managed_databases_database_parameter_parameters_name
    value = var.managed_databases_database_parameter_parameters_value

    #Optional
    update_comment = var.managed_databases_database_parameter_update_comment
  }
  scope = var.managed_databases_database_parameter_scope
}

resource "oci_database_management_managed_databases_reset_database_parameter" "test_managed_databases_reset_database_parameter" {
  #Required
  credentials {

    #Optional
    password = var.managed_databases_database_parameter_credentials_password
    role = var.managed_databases_database_parameter_credentials_role
    user_name = var.managed_databases_database_parameter_credentials_username
  }
  managed_database_id = var.managed_database_id
  parameters = [var.managed_databases_database_parameter_parameters_name]
  scope = var.managed_databases_database_parameter_scope
}

data "oci_database_management_managed_databases_database_parameter" "test_managed_databases_database_parameter" {
  #Required
  managed_database_id = var.managed_database_id

  #Optional
  is_allowed_values_included = var.managed_databases_database_parameter_is_allowed_values_included
  name = var.managed_databases_database_parameter_parameters_name
  source = var.managed_databases_database_parameter_source
}

resource "oci_core_vcn" "test_vcn" {
  cidr_block     = "10.0.0.0/16"
  compartment_id = var.compartment_id
  display_name   = "exampleVCN"
  dns_label      = "tfexamplevcn"
}

resource "oci_core_subnet" "test_subnet" {
  cidr_block     = "10.0.0.0/24"
  display_name   = "regionalSubnet"
  dns_label      = "regionalsubnet"
  compartment_id = var.compartment_id
  vcn_id         = oci_core_vcn.test_vcn.id
}

resource "oci_core_network_security_group" "test_network_security_group" {
  #Required
  compartment_id = var.compartment_id
  vcn_id         = oci_core_vcn.test_vcn.id
}

resource "oci_database_management_db_management_private_endpoint" "test_db_management_private_endpoint" {
  #Required
  compartment_id = var.compartment_id
  name = var.db_management_private_endpoint_name
  subnet_id = oci_core_subnet.test_subnet.id

  #Optional
  description = var.db_management_private_endpoint_description
  nsg_ids   = [oci_core_network_security_group.test_network_security_group.id]
  is_cluster  = var.db_management_private_endpoint_is_cluster
}

data "oci_database_management_db_management_private_endpoint" "test_db_management_private_endpoint" {
  db_management_private_endpoint_id = oci_database_management_db_management_private_endpoint.test_db_management_private_endpoint.id
}

data "oci_database_management_db_management_private_endpoints" "test_db_management_private_endpoints" {
  #Required
  compartment_id = var.compartment_id
}

data "oci_database_management_db_management_private_endpoints" "test_db_management_private_endpoints_with_name" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  name = var.db_management_private_endpoint_name
  vcn_id = oci_core_vcn.test_vcn.id
  state = var.db_management_private_endpoint_state
  is_cluster = var.db_management_private_endpoint_is_cluster
}

data "oci_database_management_job_executions_status" "test_job_executions_status" {
  #Required
  compartment_id = var.compartment_id
  start_time = formatdate("YYYY-MM-DD'T'hh:mm:ss'.000'Z", timeadd(timestamp(), "-12h"))
  end_time = formatdate("YYYY-MM-DD'T'hh:mm:ss'.000'Z", timestamp())

  #Optional
  managed_database_id = var.managed_database_id
}

data "oci_database_management_managed_database_sql_tuning_advisor_tasks" "test_managed_database_sql_tuning_advisor_tasks" {
  #Required
  managed_database_id = oci_database_management_managed_database.test_managed_database.id

  #Optional
  name                          = var.managed_database_sql_tuning_advisor_task_name
  status                        = var.managed_database_sql_tuning_advisor_task_status
  time_greater_than_or_equal_to = var.managed_database_sql_tuning_advisor_task_time_greater_than_or_equal_to
  time_less_than_or_equal_to    = var.managed_database_sql_tuning_advisor_task_time_less_than_or_equal_to
}

data "oci_database_management_managed_database_sql_tuning_advisor_tasks_execution_plan_stats_comparision" "test_managed_database_sql_tuning_advisor_tasks_execution_plan_stats_comparision" {
  #Required
  execution_id               = oci_database_management_execution.test_execution.id
  managed_database_id        = oci_database_management_managed_database.test_managed_database.id
  sql_object_id              = oci_objectstorage_object.test_object.id
  sql_tuning_advisor_task_id = oci_database_management_sql_tuning_advisor_task.test_sql_tuning_advisor_task.id
}

data "oci_database_management_managed_database_sql_tuning_advisor_tasks_findings" "test_managed_database_sql_tuning_advisor_tasks_findings" {
  #Required
  managed_database_id        = oci_database_management_managed_database.test_managed_database.id
  sql_tuning_advisor_task_id = oci_database_management_sql_tuning_advisor_task.test_sql_tuning_advisor_task.id

  #Optional
  begin_exec_id     = oci_database_management_begin_exec.test_begin_exec.id
  end_exec_id       = oci_database_management_end_exec.test_end_exec.id
  finding_filter    = var.managed_database_sql_tuning_advisor_tasks_finding_finding_filter
  index_hash_filter = var.managed_database_sql_tuning_advisor_tasks_finding_index_hash_filter
  search_period     = var.managed_database_sql_tuning_advisor_tasks_finding_search_period
  stats_hash_filter = var.managed_database_sql_tuning_advisor_tasks_finding_stats_hash_filter
}

data "oci_database_management_managed_database_sql_tuning_advisor_tasks_recommendations" "test_managed_database_sql_tuning_advisor_tasks_recommendations" {
  #Required
  execution_id               = oci_database_management_execution.test_execution.id
  managed_database_id        = oci_database_management_managed_database.test_managed_database.id
  sql_object_id              = oci_objectstorage_object.test_object.id
  sql_tuning_advisor_task_id = oci_database_management_sql_tuning_advisor_task.test_sql_tuning_advisor_task.id
}

data "oci_database_management_managed_database_sql_tuning_advisor_tasks_sql_execution_plan" "test_managed_database_sql_tuning_advisor_tasks_sql_execution_plan" {
  #Required
  attribute                  = var.managed_database_sql_tuning_advisor_tasks_sql_execution_plan_attribute
  managed_database_id        = oci_database_management_managed_database.test_managed_database.id
  sql_object_id              = oci_objectstorage_object.test_object.id
  sql_tuning_advisor_task_id = oci_database_management_sql_tuning_advisor_task.test_sql_tuning_advisor_task.id
}

data "oci_database_management_managed_database_sql_tuning_advisor_tasks_summary_report" "test_managed_database_sql_tuning_advisor_tasks_summary_report" {
  #Required
  managed_database_id        = oci_database_management_managed_database.test_managed_database.id
  sql_tuning_advisor_task_id = oci_database_management_sql_tuning_advisor_task.test_sql_tuning_advisor_task.id

  #Optional
  begin_exec_id_greater_than_or_equal_to = var.managed_database_sql_tuning_advisor_tasks_summary_report_begin_exec_id_greater_than_or_equal_to
  end_exec_id_less_than_or_equal_to      = var.managed_database_sql_tuning_advisor_tasks_summary_report_end_exec_id_less_than_or_equal_to
  search_period                          = var.managed_database_sql_tuning_advisor_tasks_summary_report_search_period
  time_greater_than_or_equal_to          = var.managed_database_sql_tuning_advisor_tasks_summary_report_time_greater_than_or_equal_to
  time_less_than_or_equal_to             = var.managed_database_sql_tuning_advisor_tasks_summary_report_time_less_than_or_equal_to
}
