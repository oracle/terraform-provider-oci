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
* `node_type` - A node type that is managed by an autoscale configuration. The only supported types are WORKER, COMPUTE_ONLY_WORKER, KAFKA_BROKER.
* `policy` - This model for autoscaling policy is deprecated and not supported for ODH clusters. Use the `AutoScalePolicyDetails` model to manage autoscale policy details for ODH clusters. 
	* `policy_type` - Types of autoscale policies. Options are SCHEDULE-BASED or THRESHOLD-BASED. (Only THRESHOLD-BASED is supported in this release.)
	* `rules` - The list of rules for autoscaling. If an action has multiple rules, the last rule in the array will be applied.
		* `action` - The valid value are CHANGE_SHAPE_SCALE_UP or CHANGE_SHAPE_SCALE_DOWN.
		* `metric` - Metric and threshold details for triggering an autoscale action.
			* `metric_type` - Allowed values are CPU_UTILIZATION and MEMORY_UTILIZATION.
			* `threshold` - An autoscale action is triggered when a performance metric exceeds a threshold.
				* `duration_in_minutes` - This value is the minimum period of time the metric value exceeds the threshold value before the action is triggered. The value is in minutes.
				* `operator` - The comparison operator to use. Options are greater than (GT) or less than (LT).
				* `value` - Integer non-negative value. 0 < value < 100
* `policy_details` - Details of an autoscale policy.

	You can create following types of autoscaling policies:
	* **MetricBasedVerticalScalingPolicy:** Vertical autoscaling action is triggered when a performance metric exceeds a threshold
	* **MetricBasedHorizontalScalingPolicy:** Horizontal autoscaling action is triggered when a performance metric exceeds a threshold
	* **ScheduleBasedVerticalScalingPolicy:** Vertical autoscaling action is triggered at the specific times that you schedule.
	* **ScheduleBasedHorizontalScalingPolicy:** Horizontal autoscaling action is triggered at the specific times that you schedule. 
	* `action_type` - The type of autoscaling action to take.
	* `policy_type` - Type of autoscaling policy.
	* `scale_down_config` - Configration for a metric based vertical scale-down policy.
		* `memory_step_size` - For nodes with [flexible compute shapes](https://docs.cloud.oracle.com/iaas/Content/bigdata/create-cluster.htm#cluster-plan-shape), this value is the size of memory in GBs to remove from each node during a scale-down event. This value is not used for nodes with fixed compute shapes.
		* `metric` - Metric and threshold details for triggering an autoscale action.
			* `metric_type` - Allowed values are CPU_UTILIZATION and MEMORY_UTILIZATION.
			* `threshold` - An autoscale action is triggered when a performance metric exceeds a threshold.
				* `duration_in_minutes` - This value is the minimum period of time the metric value exceeds the threshold value before the action is triggered. The value is in minutes.
				* `operator` - The comparison operator to use. Options are greater than (GT) or less than (LT).
				* `value` - Integer non-negative value. 0 < value < 100
		* `min_memory_per_node` - For nodes with [flexible compute shapes](https://docs.cloud.oracle.com/iaas/Content/bigdata/create-cluster.htm#cluster-plan-shape), this value is the minimum memory in GBs each node can be scaled-down to. This value is not used for nodes with fixed compute shapes.
		* `min_ocpus_per_node` - For nodes with [flexible compute shapes](https://docs.cloud.oracle.com/iaas/Content/bigdata/create-cluster.htm#cluster-plan-shape), this value is the minimum number of OCPUs each node can be scaled-down to. This value is not used for nodes with fixed compute shapes.
		* `ocpu_step_size` - For nodes with [flexible compute shapes](https://docs.cloud.oracle.com/iaas/Content/bigdata/create-cluster.htm#cluster-plan-shape), this value is the number of OCPUs to remove from each node during a scale-down event. This value is not used for nodes with fixed compute shapes.
	* `scale_in_config` - Configration for a metric based horizontal scale-in policy.
		* `metric` - Metric and threshold details for triggering an autoscale action.
			* `metric_type` - Allowed values are CPU_UTILIZATION and MEMORY_UTILIZATION.
			* `threshold` - An autoscale action is triggered when a performance metric exceeds a threshold.
				* `duration_in_minutes` - This value is the minimum period of time the metric value exceeds the threshold value before the action is triggered. The value is in minutes.
				* `operator` - The comparison operator to use. Options are greater than (GT) or less than (LT).
				* `value` - Integer non-negative value. 0 < value < 100
		* `min_node_count` - This value is the minimum number of nodes the cluster can be scaled-in to.
		* `step_size` - This value is the number of nodes to remove during a scale-in event.
	* `scale_out_config` - Configration for a metric based horizontal scale-out policy.
		* `max_node_count` - This value is the maximum number of nodes the cluster can be scaled-out to.
		* `metric` - Metric and threshold details for triggering an autoscale action.
			* `metric_type` - Allowed values are CPU_UTILIZATION and MEMORY_UTILIZATION.
			* `threshold` - An autoscale action is triggered when a performance metric exceeds a threshold.
				* `duration_in_minutes` - This value is the minimum period of time the metric value exceeds the threshold value before the action is triggered. The value is in minutes.
				* `operator` - The comparison operator to use. Options are greater than (GT) or less than (LT).
				* `value` - Integer non-negative value. 0 < value < 100
		* `step_size` - This value is the number of nodes to add during a scale-out event.
	* `scale_up_config` - Configration for a metric based vertical scale-up policy.
		* `max_memory_per_node` - For nodes with [flexible compute shapes](https://docs.cloud.oracle.com/iaas/Content/bigdata/create-cluster.htm#cluster-plan-shape), this value is the maximum memory in GBs each node can be scaled-up to. This value is not used for nodes with fixed compute shapes.
		* `max_ocpus_per_node` - For nodes with [flexible compute shapes](https://docs.cloud.oracle.com/iaas/Content/bigdata/create-cluster.htm#cluster-plan-shape), this value is the maximum number of OCPUs each node can be scaled-up to. This value is not used for nodes with fixed compute shapes. 
		* `memory_step_size` - For nodes with [flexible compute shapes](https://docs.cloud.oracle.com/iaas/Content/bigdata/create-cluster.htm#cluster-plan-shape), this value is the size of memory in GBs to add to each node during a scale-up event. This value is not used for nodes with fixed compute shapes.
		* `metric` - Metric and threshold details for triggering an autoscale action.
			* `metric_type` - Allowed values are CPU_UTILIZATION and MEMORY_UTILIZATION.
			* `threshold` - An autoscale action is triggered when a performance metric exceeds a threshold.
				* `duration_in_minutes` - This value is the minimum period of time the metric value exceeds the threshold value before the action is triggered. The value is in minutes.
				* `operator` - The comparison operator to use. Options are greater than (GT) or less than (LT).
				* `value` - Integer non-negative value. 0 < value < 100
		* `ocpu_step_size` - For nodes with [flexible compute shapes](https://docs.cloud.oracle.com/iaas/Content/bigdata/create-cluster.htm#cluster-plan-shape), this value is the number of OCPUs to add to each node during a scale-up event. This value is not used for nodes with fixed compute shapes.
	* `schedule_details` - Details of a horizontal scaling schedule.
		* `schedule_type` - The type of schedule.
		* `time_and_horizontal_scaling_config` - Time of day and horizontal scaling configuration.
			* `target_node_count` - This value is the desired number of nodes in the cluster.
			* `time_recurrence` - Day/time recurrence (specified following RFC 5545) at which to trigger autoscaling action. Currently only WEEKLY frequency is supported. Days of the week are specified using BYDAY field. Time of the day is specified using BYHOUR and BYMINUTE fields. Other fields are not supported. 
		* `time_and_vertical_scaling_config` - Time of day and vertical scaling configuration
			* `target_memory_per_node` - For nodes with [flexible compute shapes](https://docs.cloud.oracle.com/iaas/Content/bigdata/create-cluster.htm#cluster-plan-shape), this value is the desired memory in GBs on each node. This value is not used for nodes with fixed compute shapes. 
			* `target_ocpus_per_node` - For nodes with [flexible compute shapes](https://docs.cloud.oracle.com/iaas/Content/bigdata/create-cluster.htm#cluster-plan-shape), this value is the desired OCPUs count on each node. This value is not used for nodes with fixed compute shapes. 
			* `target_shape` - For nodes with [fixed compute shapes](https://docs.cloud.oracle.com/iaas/Content/bigdata/create-cluster.htm#cluster-plan-shape), this value is the desired shape of each node. This value is not used for nodes with flexible compute shapes. 
			* `time_recurrence` - Day/time recurrence (specified following RFC 5545) at which to trigger autoscaling action. Currently only WEEKLY frequency is supported. Days of the week are specified using BYDAY field. Time of the day is specified using BYHOUR and BYMINUTE fields. Other fields are not supported. 
	* `timezone` - The time zone of the execution schedule, in IANA time zone database name format
	* `trigger_type` - The type of autoscaling trigger.
* `state` - The state of the autoscale configuration.
* `time_created` - The time the cluster was created, shown as an RFC 3339 formatted datetime string.
* `time_updated` - The time the autoscale configuration was updated, shown as an RFC 3339 formatted datetime string. 

