// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "performance_tuning_analysis_result_id" {
  default = "example-performance-tuning-analysis-result-id"
}

data "oci_jms_fleet_performance_tuning_analysis_results" "test_jms_fleet_performance_tuning_analysis_results" {
	#Required
	fleet_id = var.fleet_ocid

	#Optional
  application_id = var.application_id
	managed_instance_id = var.managed_instance_ocid
	host_name = var.host_name
	time_start = var.time_start
	time_end = var.time_end
}

data "oci_jms_fleet_performance_tuning_analysis_result" "test_jms_fleet_performance_tuning_analysis_result" {
	#Required
	performance_tuning_analysis_result_id = var.performance_tuning_analysis_result_id
	fleet_id = var.fleet_ocid
}
