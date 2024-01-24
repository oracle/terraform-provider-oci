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

variable "drs_file_key" {
  default = "example-drs-file-key"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_jms_fleet_drs_files" "test_fleet_drs_files" {
  #Required
  fleet_id = var.fleet_id
}

data "oci_jms_fleet_drs_file" "test_fleet_drs_file" {
	#Required
	fleet_id = var.fleet_id
  drs_file_key = var.drs_file_key
}
