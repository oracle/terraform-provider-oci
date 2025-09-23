// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "fleet_error_compartment_id_in_subtree" {
  default = false
}

variable "fleet_error_time_first_seen_greater_than_or_equal_to" {}

variable "fleet_error_time_first_seen_less_than_or_equal_to" {}

variable "fleet_error_time_last_seen_greater_than_or_equal_to" {}

variable "fleet_error_time_last_seen_less_than_or_equal_to" {}


data "oci_jms_fleet_errors" "test_fleet_errors" {

	#Optional
	compartment_id = var.compartment_ocid
	compartment_id_in_subtree = var.fleet_error_compartment_id_in_subtree
	fleet_id = var.fleet_ocid
	time_first_seen_greater_than_or_equal_to = var.fleet_error_time_first_seen_greater_than_or_equal_to
	time_first_seen_less_than_or_equal_to = var.fleet_error_time_first_seen_less_than_or_equal_to
	time_last_seen_greater_than_or_equal_to = var.fleet_error_time_last_seen_greater_than_or_equal_to
	time_last_seen_less_than_or_equal_to = var.fleet_error_time_last_seen_less_than_or_equal_to
}