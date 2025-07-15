
// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "plugin_error_compartment_id_in_subtree" {
  default = false
}

variable "plugin_error_time_first_seen_greater_than_or_equal_to" {}

variable "plugin_error_time_first_seen_less_than_or_equal_to" {}

variable "plugin_error_time_last_seen_greater_than_or_equal_to" {}

variable "plugin_error_time_last_seen_less_than_or_equal_to" {}

data "oci_jms_plugin_errors" "test_plugin_errors" {

	#Optional
	compartment_id = var.compartment_ocid
	compartment_id_in_subtree = var.plugin_error_compartment_id_in_subtree
	managed_instance_id = var.managed_instance_ocid
	time_first_seen_greater_than_or_equal_to = var.plugin_error_time_first_seen_greater_than_or_equal_to
	time_first_seen_less_than_or_equal_to = var.plugin_error_time_first_seen_less_than_or_equal_to
	time_last_seen_greater_than_or_equal_to = var.plugin_error_time_last_seen_greater_than_or_equal_to
	time_last_seen_less_than_or_equal_to = var.plugin_error_time_last_seen_less_than_or_equal_to
}