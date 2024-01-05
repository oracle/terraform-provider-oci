// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_stack_monitoring_config" "test_config_autopromote" {
	compartment_id = var.compartment_ocid
	config_type = "AUTO_PROMOTE"
	resource_type = "HOST"
	is_enabled = true
}

resource "oci_stack_monitoring_config" "test_auto_assign_config" {
  compartment_id = var.compartment_ocid
  config_type = "LICENSE_AUTO_ASSIGN"
  license = "STANDARD_EDITION"
}

resource "oci_stack_monitoring_config" "test_enterprise_extensibility_config" {
  compartment_id = var.compartment_ocid
  config_type = "LICENSE_ENTERPRISE_EXTENSIBILITY"
  is_enabled = true
}
