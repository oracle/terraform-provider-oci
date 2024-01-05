// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "stack_mon_baselineable_metric_evaluate_resource_id" {}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_stack_monitoring_baselineable_metric" "test_baselineable_metric_evaluate_baselineable_metric" {
	#Required
	compartment_id = var.compartment_ocid
	resource_group = "my_resource"
	namespace = "my_namespace"
	name = "metric_name"
	column = "metric_column"
}

data "oci_stack_monitoring_baselineable_metrics_evaluate" "test_baselineable_metrics_evaluate" {
	#Required
	baselineable_metric_id = oci_stack_monitoring_baselineable_metric.test_baselineable_metric_evaluate_baselineable_metric.id
	resource_id = var.stack_mon_baselineable_metric_evaluate_resource_id
	items {
        evaluation_data_points {
            timestamp = "2023-05-15T05:00:00.001Z"
            value = "1.0"
        }
        training_data_points {
            timestamp = "2023-05-14T05:00:00.001Z"
            value = "1.0"
        }
        training_data_points {
            timestamp = "2023-05-14T05:01:00.001Z"
            value = "1.0"
        }
        training_data_points {
            timestamp = "2023-05-14T05:02:00.001Z"
            value = "1.0"
        }
        training_data_points {
            timestamp = "2023-05-14T05:03:00.001Z"
            value = "1.0"
        }
        training_data_points {
            timestamp = "2023-05-14T05:04:00.001Z"
            value = "1.0"
        }
        training_data_points {
            timestamp = "2023-05-14T05:05:00.001Z"
            value = "1.0"
        }
        training_data_points {
            timestamp = "2023-05-14T05:06:00.001Z"
            value = "1.0"
        }
        training_data_points {
            timestamp = "2023-05-14T05:07:00.001Z"
            value = "1.0"
        }
        training_data_points {
            timestamp = "2023-05-14T05:08:00.001Z"
            value = "1.0"
        }
        training_data_points {
            timestamp = "2023-05-14T05:09:00.001Z"
            value = "1.0"
        }
        training_data_points {
            timestamp = "2023-05-14T05:10:00.001Z"
            value = "1.0"
        }
        training_data_points {
            timestamp = "2023-05-14T05:11:00.001Z"
            value = "1.0"
        }
        training_data_points {
            timestamp = "2023-05-14T05:12:00.001Z"
            value = "1.0"
        }
        training_data_points {
            timestamp = "2023-05-14T05:13:00.001Z"
            value = "1.0"
        }
        training_data_points {
            timestamp = "2023-05-14T05:14:00.001Z"
            value = "1.0"
        }
        training_data_points {
            timestamp = "2023-05-14T05:15:00.001Z"
            value = "1.0"
        }
        training_data_points {
            timestamp = "2023-05-14T05:16:00.001Z"
            value = "1.0"
        }
        training_data_points {
            timestamp = "2023-05-14T05:17:00.001Z"
            value = "1.0"
        }
        training_data_points {
            timestamp = "2023-05-14T05:18:00.001Z"
            value = "1.0"
        }
        training_data_points {
            timestamp = "2023-05-14T05:19:00.001Z"
            value = "1.0"
        }
        training_data_points {
            timestamp = "2023-05-14T05:20:00.001Z"
            value = "1.0"
        }
        training_data_points {
            timestamp = "2023-05-14T05:21:00.001Z"
            value = "1.0"
        }
        training_data_points {
            timestamp = "2023-05-14T05:22:00.001Z"
            value = "1.0"
        }
        training_data_points {
            timestamp = "2023-05-14T05:23:00.001Z"
            value = "1.0"
        }
        training_data_points {
            timestamp = "2023-05-14T05:24:00.001Z"
            value = "1.0"
        }
        training_data_points {
            timestamp = "2023-05-14T05:25:00.001Z"
            value = "1.0"
        }
        training_data_points {
            timestamp = "2023-05-14T05:26:00.001Z"
            value = "1.0"
        }
        training_data_points {
            timestamp = "2023-05-14T05:27:00.001Z"
            value = "1.0"
        }
        training_data_points {
            timestamp = "2023-05-14T05:28:00.001Z"
            value = "1.0"
        }
        training_data_points {
            timestamp = "2023-05-14T05:29:00.001Z"
            value = "1.0"
        }
        #Optional
        dimensions = {
            "dimension1" = "value1"
        }
	}
}
