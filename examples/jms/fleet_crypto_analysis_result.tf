// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "aggregation_mode" {
  default = "JFR"
}

variable "finding_count" {
	default = 10
}

variable "finding_count_greater_than" {
	default = 10
}

variable "non_compliant_finding_count" {
	default = 10
}

variable "non_compliant_finding_count_greater_than" {
	default = 10
}

variable "crypto_analysis_result_id" {
  default = "example-crypto-analysis-result-id"
}

data "oci_jms_fleet_crypto_analysis_results" "test_fleet_crypto_analysis_results" {
	#Required
	fleet_id = var.fleet_ocid

	#Optional
	aggregation_mode = var.aggregation_mode
	managed_instance_id = var.managed_instance_ocid
	host_name = var.host_name
	finding_count = var.finding_count
	finding_count_greater_than = var.finding_count_greater_than
	non_compliant_finding_count = var.non_compliant_finding_count
	non_compliant_finding_count_greater_than = var.non_compliant_finding_count_greater_than
	time_end = var.time_end
	time_start = var.time_start
}

data "oci_jms_fleet_crypto_analysis_result" "test_fleet_crypto_analysis_result" {
	#Required
	crypto_analysis_result_id = var.crypto_analysis_result_id
	fleet_id = var.fleet_ocid
}