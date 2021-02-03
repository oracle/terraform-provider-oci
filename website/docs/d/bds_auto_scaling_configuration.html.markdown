---
subcategory: "Big Data Service"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_bds_auto_scaling_configuration"
sidebar_current: "docs-oci-datasource-bds-auto_scaling_configuration"
description: |-
  Provides details about a specific Auto Scaling Configuration in Oracle Cloud Infrastructure Big Data Service service
---

# Data Source: oci_bds_auto_scaling_configuration
This data source provides details about a specific Auto Scaling Configuration resource in Oracle Cloud Infrastructure Big Data Service service.

Gets information about the specified autoscaling configuration.


## Example Usage

```hcl
data "oci_bds_auto_scaling_configuration" "test_auto_scaling_configuration" {
	#Required
	auto_scaling_configuration_id = oci_autoscaling_auto_scaling_configuration.test_auto_scaling_configuration.id
	bds_instance_id = oci_bds_bds_instance.test_bds_instance.id
}
```

## Argument Reference

The following arguments are supported:

* `auto_scaling_configuration_id` - (Required) Unique Oracle-assigned identifier of the autoscaling configuration.
* `bds_instance_id` - (Required) The OCID of the BDS instance


## Attributes Reference

The following attributes are exported:

* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `id` - The unique identifier for autoscaling configuration.
* `node_type` - A node type that is managed by an autoscaling configuration. The only supported type is WORKER.
* `policy` - Policy definitions for the autoscaling configuration
	* `policy_type` - Types of autoscaling policies. SCHEDULE-BASED or  THRESHOLD-BASED, current only supported THRESHOLD-BASED.
	* `rules` - The list of rules for autoscaling. If an action have multiple rules, last rule in the array will be applied.
		* `action` - The valid value are - CHANGE_SHAPE_SCALE_UP or CHANGE_SHAPE_SCALE_DOWN
		* `metric` - Metric and threshold details for triggering an autoscaling action
			* `metric_type` - Allowed value is CPU_UTILIZATION currently
			* `threshold` - An autoscaling action is triggered when a performance metric meets or exceeds a threshold
				* `duration_in_minutes` - This value is the minimum period of time metric value meets or exceeds threshold value before action is trigger. The value is in minutes.
				* `operator` - The comparison operator to use. Options are greater than (GT), less than (LT).
				* `value` - integer non negative value. 0 < value < 100
* `state` - The state of the autoscaling configuration
* `time_created` - The time the BDS instance was created. An RFC3339 formatted datetime string
* `time_updated` - The time the autoscale configuration was updated. An RFC3339 formatted datetime string 

