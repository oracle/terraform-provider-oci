// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

data "oci_jms_fleet_diagnoses" "test_jms_fleet_diagnoses" {
	#Required
	fleet_id = var.fleet_ocid
}
