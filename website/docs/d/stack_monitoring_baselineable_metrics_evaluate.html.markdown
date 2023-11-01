---
subcategory: "Stack Monitoring"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_stack_monitoring_baselineable_metrics_evaluate"
sidebar_current: "docs-oci-datasource-stack_monitoring-baselineable_metrics_evaluate"
description: |-
  Provides details about a specific Baselineable Metrics Evaluate in Oracle Cloud Infrastructure Stack Monitoring service
---

# Data Source: oci_stack_monitoring_baselineable_metrics_evaluate
This data source provides details about a specific Baselineable Metrics Evaluate resource in Oracle Cloud Infrastructure Stack Monitoring service.

Evaluates metric for anomalies for the given data points

## Example Usage

```hcl
data "oci_stack_monitoring_baselineable_metrics_evaluate" "test_baselineable_metrics_evaluate" {
	#Required
	baselineable_metric_id = oci_stack_monitoring_baselineable_metric.test_baselineable_metric.id
	items {
		#Required
		evaluation_data_points {
			#Required
			timestamp = var.baselineable_metrics_evaluate_items_evaluation_data_points_timestamp
			value = var.baselineable_metrics_evaluate_items_evaluation_data_points_value
		}
		training_data_points {
			#Required
			timestamp = var.baselineable_metrics_evaluate_items_training_data_points_timestamp
			value = var.baselineable_metrics_evaluate_items_training_data_points_value
		}

		#Optional
		dimensions = var.baselineable_metrics_evaluate_items_dimensions
	}
	resource_id = oci_usage_proxy_resource.test_resource.id
}
```

## Argument Reference

The following arguments are supported:

* `baselineable_metric_id` - (Required) Identifier for the metric
* `items` - (Required) List of Metric data
	* `dimensions` - (Optional) list of dimensions for the metric
	* `evaluation_data_points` - (Required) list of data points for the metric for evaluation of anomalies
		* `timestamp` - (Required) timestamp of when the metric was collected
		* `value` - (Required) value for the metric data point
	* `training_data_points` - (Required) list of data points for the metric for training of baseline
		* `timestamp` - (Required) timestamp of when the metric was collected
		* `value` - (Required) value for the metric data point
* `resource_id` - (Required) OCID of the resource


## Attributes Reference

The following attributes are exported:

* `items` - List of Metric data
	* `data_points` - list of anomaly data points for the metric
		* `anomaly` - if the value is anomaly or not 0 indicates not an anomaly -1 indicates value is below the threshold +1 indicates value is above the threshold
		* `high` - upper threshold for the metric value
		* `low` - lower threshold for the metric value
		* `timestamp` - timestamp of when the metric was collected
		* `value` - value for the metric data point
	* `dimensions` - list of dimensions for the metric
* `resource_id` - OCID of the resource

