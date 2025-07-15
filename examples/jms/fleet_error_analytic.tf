// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "fleet_error_analytic_compartment_id_in_subtree" {
  default = false
}

data "oci_jms_fleet_error_analytics" "test_fleet_error_analytics" {

	#Optional
	compartment_id = var.compartment_ocid
	compartment_id_in_subtree = var.fleet_error_analytic_compartment_id_in_subtree
}