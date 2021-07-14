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

Returns details of the autoscale configuration identified by the given ID.


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

* `auto_scaling_configuration_id` - (Required) Unique Oracle-assigned identifier of the autoscale configuration.
* `bds_instance_id` - (Required) The OCID of the cluster.


## Attributes Reference

The following attributes are exported:

* `display_name` - A user-friendly name. The name does not have to be unique, and it may be changed. Avoid entering confidential information.
* `id` - The unique identifier for the autoscale configuration.
* `node_type` - A node type that is managed by an autoscale configuration. The only supported type is WORKER.
* `policy` - Policy definitions for the autoscale configuration.
	* `policy_type` - Types of autoscale policies. Options are SCHEDULE-BASED or THRESHOLD-BASED. (Only THRESHOLD-BASED is supported in this release.)
	* `rules` - The list of rules for autoscaling. If an action has multiple rules, the last rule in the array will be applied.
		* `action` - The valid value are CHANGE_SHAPE_SCALE_UP or CHANGE_SHAPE_SCALE_DOWN.
		* `metric` - Metric and threshold details for triggering an autoscale action.
			* `metric_type` - Allowed value is CPU_UTILIZATION.
			* `threshold` - An autoscale action is triggered when a performance metric meets or exceeds a threshold.
				* `duration_in_minutes` - This value is the minimum period of time the metric value meets or exceeds the threshold value before the action is triggered. The value is in minutes.
				* `operator` - The comparison operator to use. Options are greater than (GT) or less than (LT).
				* `value` - Integer non-negative value. 0 < value < 100
* `state` - The state of the autoscale configuration.
* `time_created` - The time the cluster was created, shown as an RFC 3339 formatted datetime string.
* `time_updated` - The time the autoscale configuration was updated, shown as an RFC 3339 formatted datetime string. 

