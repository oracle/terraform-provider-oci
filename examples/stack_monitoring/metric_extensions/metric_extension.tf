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

resource "oci_stack_monitoring_metric_extension" "test_metric_extension_example_http" {
	#Required
	compartment_id = var.compartment_ocid
	name = "ME_MetricExtensionTerraformExampleHttp"
	resource_type = "oracle_goldengate_admin_server"
	display_name = "Golden Gate IO Read Performance"
	collection_recurrences = "FREQ=MINUTELY;INTERVAL=10"
	metric_list {
		name = "IoReadBytes"
		display_name = "IO Read Bytes"
		is_dimension = false
		data_type = "NUMBER"
		is_hidden = true 
		metric_category = "AVAILABILITY"
	} 
  metric_list {
		name = "IOReadRate"
		display_name = "IO Read Rate"
		is_dimension = false
		data_type = "NUMBER"
		is_hidden = false 
		metric_category = "AVAILABILITY"
    compute_expression = "(IoReadBytes > 0 ? (__interval == 0 ? 0 : (IoReadBytes > _IoReadBytes ? (((IoReadBytes - _IoReadBytes) / __interval) / (1024 * 1024)) : 0))  : 0)"
	}
	query_properties {
		collection_method = "HTTP"
		url = "%pm_server_url%/services/v2/mpoints/%api_process_name%/processPerformance"
		response_content_type = "APPLICATION_JSON"
		protocol_type = "HTTPS"
		script_details {
			name = "ioRead_performance.js"
			content = "ZnVuY3Rpb24gcnVuTWV0aG9kKG1ldHJpY09ic2VydmF0aW9uLCBzb3VyY2VQcm9wcykKewogICAgbGV0IHJlc3BvbnNlX3JhdyA9IEpTT04ucGFyc2UobWV0cmljT2JzZXJ2YXRpb24pOwogICAgbGV0IHJlc3BvbnNlID0gcmVzcG9uc2VfcmF3LnJlc3BvbnNlOwogICAgbGV0IGlvUmVhZEJ5dGVzID0gcmVzcG9uc2VbImlvUmVhZEJ5dGVzIl07CiAgICByZXR1cm4gW1tpb1JlYWRCeXRlc11dOwp9"
		}
	}
	#Optional
	description = "Collects count of instances in 'UP' status in staging compartments from monitoring table"
	publish_trigger = var.publish_trigger
}

data "oci_stack_monitoring_metric_extension" "test_metric_extension_example_http" {
  # Required
  metric_extension_id = oci_stack_monitoring_metric_extension.test_metric_extension_example_http.id
}
