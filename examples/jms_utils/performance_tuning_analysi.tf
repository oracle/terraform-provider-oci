// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "performance_tuning_analysi_id" {
  default = "id"
}

variable "performance_tuning_analysi_performance_tuning_analysis_result" {
  default = "ACTION_RECOMMENDED"
}

variable "performance_tuning_analysi_project_name" {
  default = "project_name"
}


data "oci_jms_utils_performance_tuning_analysis" "test_performance_tuning_analysis" {

  #Optional
  analysis_project_name              = var.performance_tuning_analysi_project_name
  compartment_id                     = var.tenancy_ocid
  id                                 = var.performance_tuning_analysi_id
  performance_tuning_analysis_result = var.performance_tuning_analysi_performance_tuning_analysis_result
}