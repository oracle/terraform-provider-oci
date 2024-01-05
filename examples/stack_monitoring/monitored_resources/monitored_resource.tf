// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "stack_mon_management_agent_id_resource1" {}
variable "stack_mon_management_agent_id_resource2" {}
variable "stack_mon_hostname_resource1" {}
variable "stack_mon_hostname_resource2" {}


provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_stack_monitoring_monitored_resource" "test_monitored_resource" {
	#Required
	compartment_id = var.compartment_ocid
	name = "terraformExample"
	type = "host"

	#Optional
	display_name = "exampleDisplayName"
	host_name = var.stack_mon_hostname_resource1
	management_agent_id = var.stack_mon_management_agent_id_resource1
	properties {
		name = "osName"
		value = "Linux"
	}
	properties {
		name = "osVersion"
		value = "7.0"
	}
	resource_time_zone = "en"
	license = "ENTERPRISE_EDITION"
	freeform_tags = { "bar-key" = "test_monitored_resource.value" }
	lifecycle {
		ignore_changes = [
			properties, external_id]
	}
}

data "oci_stack_monitoring_monitored_resource" "test_monitored_resource" {
  #Required
  monitored_resource_id = oci_stack_monitoring_monitored_resource.test_monitored_resource.id
}

data "oci_stack_monitoring_monitored_resources" "test_monitored_resources" {
	#Required
	compartment_id = oci_stack_monitoring_monitored_resource.test_monitored_resource.compartment_id
	#Optional
	name = 	oci_stack_monitoring_monitored_resource.test_monitored_resource.name
}