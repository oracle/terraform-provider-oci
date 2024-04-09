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

resource "oci_stack_monitoring_monitored_resource" "test_monitored_resource1" {
	#Required
	compartment_id = var.compartment_ocid
	name = "terraformAssocTestExample"
	type = "host"

	#Optional
	display_name = "exampleDisplayName"
	host_name = var.stack_mon_hostname_resource1
	management_agent_id = var.stack_mon_management_agent_id_resource1
	license = "STANDARD_EDITION"
	properties {
		name = "osName"
		value = "Linux"
	}
	properties {
		name = "osVersion"
		value = "7.0"
	}
	resource_time_zone = "en"
	lifecycle {
		ignore_changes = [
			credentials,
			properties,
			external_id,
			defined_tags]
	}
}
resource "oci_stack_monitoring_monitored_resource" "test_monitored_resource2" {
	#Required
	compartment_id = var.compartment_ocid
	name = "terraformAssocTestExample2"
	type = "host"

	#Optional
	display_name = "exampleDisplayName2"
	host_name = var.stack_mon_hostname_resource2
	management_agent_id = var.stack_mon_management_agent_id_resource2
	properties {
		name = "osName"
		value = "Linux"
	}
	properties {
		name = "osVersion"
		value = "7.0"
	}
	resource_time_zone = "en"
	lifecycle {
		ignore_changes = [
			credentials,
			properties,
			external_id,
			defined_tags]
	}
}

resource "oci_stack_monitoring_monitored_resources_associate_monitored_resource" "test_monitored_resources_associate_monitored_resource" {
	#Required
	association_type = "uses"
	compartment_id = var.compartment_ocid
	destination_resource_id = oci_stack_monitoring_monitored_resource.test_monitored_resource2.id
	source_resource_id = oci_stack_monitoring_monitored_resource.test_monitored_resource1.id
}
