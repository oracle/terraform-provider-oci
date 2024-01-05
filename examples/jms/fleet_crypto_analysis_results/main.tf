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

variable "aggregation_mode" {
  default = "JFR"
}

variable "managed_instance_id" {}

variable "time_start" {}

variable "time_end" {}

variable "crypto_analysis_result_id" {
  default = "example-crypto-analysis-result-id"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_jms_fleet_crypto_analysis_results" "test_fleet_crypto_analysis_results" {
	#Required
	fleet_id = var.fleet_id

	#Optional
	aggregation_mode = var.aggregation_mode
	managed_instance_id = var.managed_instance_id
	time_end = var.time_end
	time_start = var.time_start
}

data "oci_jms_fleet_crypto_analysis_result" "test_fleet_crypto_analysis_result" {
	#Required
	crypto_analysis_result_id = var.crypto_analysis_result_id
	fleet_id = var.fleet_id
}