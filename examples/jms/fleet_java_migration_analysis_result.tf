// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "java_migration_analysis_result_id" {
  default = "example-java-migration-analysis-result-id"
}

data "oci_jms_fleet_java_migration_analysis_results" "test_fleet_java_migration_analysis_results" {
	#Required
	fleet_id = var.fleet_ocid

	#Optional
	managed_instance_id = var.managed_instance_ocid
	host_name = var.host_name
	application_name = var.application_name
	time_start = var.time_start
	time_end = var.time_end
}

data "oci_jms_fleet_java_migration_analysis_result" "test_fleet_java_migration_analysis_result" {
	#Required
	java_migration_analysis_result_id = var.java_migration_analysis_result_id
	fleet_id = var.fleet_ocid
}
