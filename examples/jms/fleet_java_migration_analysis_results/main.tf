// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

variable "fleet_id" {
  default = "example-fleet-id"
}

variable "managed_instance_id" {}

variable "time_start" {}

variable "time_end" {}

variable "java_migration_analysis_result_id" {
  default = "example-java-migration-analysis-result-id"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_jms_fleet_java_migration_analysis_results" "test_fleet_java_migration_analysis_results" {
	#Required
	fleet_id = var.fleet_id

	#Optional
	managed_instance_id = var.managed_instance_id
	time_start = var.time_start
	time_end = var.time_end
}

data "oci_jms_fleet_java_migration_analysis_result" "test_fleet_java_migration_analysis_result" {
	#Required
	java_migration_analysis_result_id = var.java_migration_analysis_result_id
	fleet_id = var.fleet_id
}
