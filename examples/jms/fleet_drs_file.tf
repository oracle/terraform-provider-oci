// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "drs_file_key" {
  default = "example-drs-file-key"
}

data "oci_jms_fleet_drs_files" "test_fleet_drs_files" {
  #Required
  fleet_id = "example-fleet-id"
}

data "oci_jms_fleet_drs_file" "test_fleet_drs_file" {
	#Required
	fleet_id = "example-fleet-id"
  drs_file_key = var.drs_file_key
}
