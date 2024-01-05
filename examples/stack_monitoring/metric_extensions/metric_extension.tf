// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "stack_mon_management_agent_id_resource1" {}
variable "stack_mon_hostname_resource1" {}
variable "publish_trigger" {default = false} // setting to false to create DRAFT Metric extension


provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_stack_monitoring_metric_extension" "test_metric_extension_example" {
	#Required
	compartment_id = var.compartment_ocid
	name = "ME_MetricExtensionTerraformExample"
	resource_type = "ebs_instance"
	display_name = "Count of Running Instances"
	collection_recurrences = "FREQ=MINUTELY;INTERVAL=10"
	metric_list {
		name = "NumberOfUpInstances"
		display_name = "NumberOfUpInstances"
		is_dimension = false
		data_type = "NUMBER"
		is_hidden = false
		metric_category = "AVAILABILITY"
	}
	query_properties {
		collection_method = "SQL"
		sql_type = "STATEMENT"
		sql_details {
			content = "U0VMRUNUIGNvdW50KGluc3RhbmNlX2lkKSBGUk9NIG1vbml0b3JpbmdfdGFibGUgV0hFUkUgc3RhdHVzID0gJ1VQJyBBTkQgY29tcGFydG1lbnRfdHlwZSA9IDox"
		}
		in_param_details {
			in_param_position = 1
			in_param_value = "staging"
		}
	}
	#Optional
	description = "Collects count of instances in 'UP' status in staging compartments from monitoring table"
	publish_trigger = var.publish_trigger
}


data "oci_stack_monitoring_metric_extension" "test_metric_extension_example" {
  #Required
  metric_extension_id = oci_stack_monitoring_metric_extension.test_metric_extension_example.id
}
