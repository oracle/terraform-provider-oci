// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "stack_mon_source_resource_id" {}
variable "stack_mon_destination_resource_id" {}


provider "oci" {
	tenancy_ocid     = var.tenancy_ocid
	user_ocid        = var.user_ocid
	fingerprint      = var.fingerprint
	private_key_path = var.private_key_path
	region           = var.region
}


resource "oci_stack_monitoring_monitored_resources_list_member" "test_monitored_resources_list_member" {
  #Required
  monitored_resource_id = var.stack_mon_source_resource_id

  #Optional
  destination_resource_id = var.stack_mon_destination_resource_id
}
