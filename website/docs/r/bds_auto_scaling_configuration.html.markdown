---
subcategory: "Big Data Service"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_bds_auto_scaling_configuration"
sidebar_current: "docs-oci-resource-bds-auto_scaling_configuration"
description: |-
  Provides the Auto Scaling Configuration resource in Oracle Cloud Infrastructure Big Data Service service
---

# oci_bds_auto_scaling_configuration
This resource provides the Auto Scaling Configuration resource in Oracle Cloud Infrastructure Big Data Service service.

Add an autoscale configuration to the cluster.


## Example Usage

```hcl
resource "oci_bds_auto_scaling_configuration" "test_auto_scaling_configuration" {
	#Required
	bds_instance_id = oci_bds_bds_instance.test_bds_instance.id
	cluster_admin_password = var.auto_scaling_configuration_cluster_admin_password
	is_enabled = var.auto_scaling_configuration_is_enabled
	node_type = var.auto_scaling_configuration_node_type

	#Optional
	display_name = var.auto_scaling_configuration_display_name
	policy_details {
		#Required
		policy_type = var.auto_scaling_configuration_policy_details_policy_type

		#Optional
		scale_down_config {

			#Optional
			memory_step_size = var.auto_scaling_configuration_policy_details_scale_down_config_memory_step_size
			metric {

				#Optional
				metric_type = var.auto_scaling_configuration_policy_details_scale_down_config_metric_metric_type
				threshold {

					#Optional
					duration_in_minutes = var.auto_scaling_configuration_policy_details_scale_down_config_metric_threshold_duration_in_minutes
					operator = var.auto_scaling_configuration_policy_details_scale_down_config_metric_threshold_operator
					value = var.auto_scaling_configuration_policy_details_scale_down_config_metric_threshold_value
				}
			}
			min_memory_per_node = var.auto_scaling_configuration_policy_details_scale_down_config_min_memory_per_node
			min_ocpus_per_node = var.auto_scaling_configuration_policy_details_scale_down_config_min_ocpus_per_node
			ocpu_step_size = var.auto_scaling_configuration_policy_details_scale_down_config_ocpu_step_size
		}
		scale_up_config {

			#Optional
			max_memory_per_node = var.auto_scaling_configuration_policy_details_scale_up_config_max_memory_per_node
			max_ocpus_per_node = var.auto_scaling_configuration_policy_details_scale_up_config_max_ocpus_per_node
			memory_step_size = var.auto_scaling_configuration_policy_details_scale_up_config_memory_step_size
			metric {

				#Optional
				metric_type = var.auto_scaling_configuration_policy_details_scale_up_config_metric_metric_type
				threshold {

					#Optional
					duration_in_minutes = var.auto_scaling_configuration_policy_details_scale_up_config_metric_threshold_duration_in_minutes
					operator = var.auto_scaling_configuration_policy_details_scale_up_config_metric_threshold_operator
					value = var.auto_scaling_configuration_policy_details_scale_up_config_metric_threshold_value
				}
			}
			ocpu_step_size = var.auto_scaling_configuration_policy_details_scale_up_config_ocpu_step_size
		}
	}
}
```

## Argument Reference

The following arguments are supported:

* `bds_instance_id` - (Required) The OCID of the cluster.
* `cluster_admin_password` - (Required) (Updatable) Base-64 encoded password for the cluster (and Cloudera Manager) admin user.
* `display_name` - (Optional) (Updatable) A user-friendly name. The name does not have to be unique, and it may be changed. Avoid entering confidential information.
* `is_enabled` - (Required) (Updatable) Whether the autoscale configuration is enabled.
* `node_type` - (Required) A node type that is managed by an autoscale configuration. The only supported types are WORKER, COMPUTE_ONLY_WORKER and KAFKA_BROKER.
* `policy` - (Optional) (Updatable) This model for autoscaling policy is deprecated and not supported for ODH clusters. Use the `AutoScalePolicyDetails` model to manage autoscale policy details for ODH clusters. 
	* `policy_type` - (Required) (Updatable) Types of autoscale policies. Options are SCHEDULE-BASED or THRESHOLD-BASED. (Only THRESHOLD-BASED is supported in this release.)
	* `rules` - (Required) (Updatable) The list of rules for autoscaling. If an action has multiple rules, the last rule in the array will be applied.
		* `action` - (Required) (Updatable) The valid value are CHANGE_SHAPE_SCALE_UP or CHANGE_SHAPE_SCALE_DOWN.
		* `metric` - (Required) (Updatable) Metric and threshold details for triggering an autoscale action.
			* `metric_type` - (Required) (Updatable) Allowed values are CPU_UTILIZATION and MEMORY_UTILIZATION.
			* `threshold` - (Required) (Updatable) An autoscale action is triggered when a performance metric exceeds a threshold.
				* `duration_in_minutes` - (Required) (Updatable) This value is the minimum period of time the metric value exceeds the threshold value before the action is triggered. The value is in minutes.
				* `operator` - (Required) (Updatable) The comparison operator to use. Options are greater than (GT) or less than (LT).
				* `value` - (Required) (Updatable) Integer non-negative value. 0 < value < 100
* `policy_details` - (Optional) (Updatable) Policy definition for the autoscale configuration.

	An autoscaling policy is part of an autoscaling configuration. For more information, see [Autoscaling](https://docs.cloud.oracle.com/iaas/Content/bigdata/create-cluster.htm#cluster-autoscale)

	You can create following type of autoscaling policies:
	* **MetricBasedVerticalScalingPolicy:** Vertical autoscaling action is triggered when a performance metric exceeds a threshold
	* **MetricBasedHorizontalScalingPolicy:** Horizontal autoscaling action is triggered when a performance metric exceeds a threshold
	* **ScheduleBasedVerticalScalingPolicy:** Vertical autoscaling action is triggered at the specific times that you schedule.
	* **ScheduleBasedHorizontalScalingPolicy:** Horizontal autoscaling action is triggered at the specific times that you schedule.

	An autoscaling configuration can have one of above supported policies. 
	* `policy_type` - (Required) Type of autoscaling policy.
	* `scale_down_config` - (Applicable when policy_type=METRIC_BASED_VERTICAL_SCALING_POLICY) (Updatable) Configration for a metric based vertical scale-down policy.
		* `memory_step_size` - (Applicable when policy_type=METRIC_BASED_VERTICAL_SCALING_POLICY) (Updatable) For nodes with [flexible compute shapes](https://docs.cloud.oracle.com/iaas/Content/bigdata/create-cluster.htm#cluster-plan-shape), this value is the size of memory in GBs to remove from each node during a scale-down event. This value is not used for nodes with fixed compute shapes.
		* `metric` - (Applicable when policy_type=METRIC_BASED_VERTICAL_SCALING_POLICY) (Updatable) Metric and threshold details for triggering an autoscale action.
			* `metric_type` - (Required when policy_type=METRIC_BASED_VERTICAL_SCALING_POLICY) (Updatable) Allowed values are CPU_UTILIZATION and MEMORY_UTILIZATION.
			* `threshold` - (Required when policy_type=METRIC_BASED_VERTICAL_SCALING_POLICY) (Updatable) An autoscale action is triggered when a performance metric exceeds a threshold.
				* `duration_in_minutes` - (Required when policy_type=METRIC_BASED_VERTICAL_SCALING_POLICY) (Updatable) This value is the minimum period of time the metric value exceeds the threshold value before the action is triggered. The value is in minutes.
				* `operator` - (Required when policy_type=METRIC_BASED_VERTICAL_SCALING_POLICY) (Updatable) The comparison operator to use. Options are greater than (GT) or less than (LT).
				* `value` - (Required when policy_type=METRIC_BASED_VERTICAL_SCALING_POLICY) (Updatable) Integer non-negative value. 0 < value < 100
		* `min_memory_per_node` - (Applicable when policy_type=METRIC_BASED_VERTICAL_SCALING_POLICY) (Updatable) For nodes with [flexible compute shapes](https://docs.cloud.oracle.com/iaas/Content/bigdata/create-cluster.htm#cluster-plan-shape), this value is the minimum memory in GBs each node can be scaled-down to. This value is not used for nodes with fixed compute shapes.
		* `min_ocpus_per_node` - (Applicable when policy_type=METRIC_BASED_VERTICAL_SCALING_POLICY) (Updatable) For nodes with [flexible compute shapes](https://docs.cloud.oracle.com/iaas/Content/bigdata/create-cluster.htm#cluster-plan-shape), this value is the minimum number of OCPUs each node can be scaled-down to. This value is not used for nodes with fixed compute shapes.
		* `ocpu_step_size` - (Applicable when policy_type=METRIC_BASED_VERTICAL_SCALING_POLICY) (Updatable) For nodes with [flexible compute shapes](https://docs.cloud.oracle.com/iaas/Content/bigdata/create-cluster.htm#cluster-plan-shape), this value is the number of OCPUs to remove from each node during a scale-down event. This value is not used for nodes with fixed compute shapes.
	* `scale_in_config` - (Applicable when policy_type=METRIC_BASED_HORIZONTAL_SCALING_POLICY) (Updatable) Configration for a metric based horizontal scale-in policy.
		* `metric` - (Applicable when policy_type=METRIC_BASED_HORIZONTAL_SCALING_POLICY) (Updatable) Metric and threshold details for triggering an autoscale action.
			* `metric_type` - (Required when policy_type=METRIC_BASED_HORIZONTAL_SCALING_POLICY) (Updatable) Allowed values are CPU_UTILIZATION and MEMORY_UTILIZATION.
			* `threshold` - (Required when policy_type=METRIC_BASED_HORIZONTAL_SCALING_POLICY) (Updatable) An autoscale action is triggered when a performance metric exceeds a threshold.
				* `duration_in_minutes` - (Required when policy_type=METRIC_BASED_HORIZONTAL_SCALING_POLICY) (Updatable) This value is the minimum period of time the metric value exceeds the threshold value before the action is triggered. The value is in minutes.
				* `operator` - (Required when policy_type=METRIC_BASED_HORIZONTAL_SCALING_POLICY) (Updatable) The comparison operator to use. Options are greater than (GT) or less than (LT).
				* `value` - (Required when policy_type=METRIC_BASED_HORIZONTAL_SCALING_POLICY) (Updatable) Integer non-negative value. 0 < value < 100
		* `min_node_count` - (Applicable when policy_type=METRIC_BASED_HORIZONTAL_SCALING_POLICY) (Updatable) This value is the minimum number of nodes the cluster can be scaled-in to.
		* `step_size` - (Applicable when policy_type=METRIC_BASED_HORIZONTAL_SCALING_POLICY) (Updatable) This value is the number of nodes to remove during a scale-in event.
	* `scale_out_config` - (Applicable when policy_type=METRIC_BASED_HORIZONTAL_SCALING_POLICY) (Updatable) Configration for a metric based horizontal scale-out policy.
		* `max_node_count` - (Applicable when policy_type=METRIC_BASED_HORIZONTAL_SCALING_POLICY) (Updatable) This value is the maximum number of nodes the cluster can be scaled-out to.
		* `metric` - (Applicable when policy_type=METRIC_BASED_HORIZONTAL_SCALING_POLICY) (Updatable) Metric and threshold details for triggering an autoscale action.
			* `metric_type` - (Required when policy_type=METRIC_BASED_HORIZONTAL_SCALING_POLICY) (Updatable) Allowed values are CPU_UTILIZATION and MEMORY_UTILIZATION.
			* `threshold` - (Required when policy_type=METRIC_BASED_HORIZONTAL_SCALING_POLICY) (Updatable) An autoscale action is triggered when a performance metric exceeds a threshold.
				* `duration_in_minutes` - (Required when policy_type=METRIC_BASED_HORIZONTAL_SCALING_POLICY) (Updatable) This value is the minimum period of time the metric value exceeds the threshold value before the action is triggered. The value is in minutes.
				* `operator` - (Required when policy_type=METRIC_BASED_HORIZONTAL_SCALING_POLICY) (Updatable) The comparison operator to use. Options are greater than (GT) or less than (LT).
				* `value` - (Required when policy_type=METRIC_BASED_HORIZONTAL_SCALING_POLICY) (Updatable) Integer non-negative value. 0 < value < 100
		* `step_size` - (Applicable when policy_type=METRIC_BASED_HORIZONTAL_SCALING_POLICY) (Updatable) This value is the number of nodes to add during a scale-out event.
	* `scale_up_config` - (Applicable when policy_type=METRIC_BASED_VERTICAL_SCALING_POLICY) (Updatable) Configration for a metric based vertical scale-up policy.
		* `max_memory_per_node` - (Applicable when policy_type=METRIC_BASED_VERTICAL_SCALING_POLICY) (Updatable) For nodes with [flexible compute shapes](https://docs.cloud.oracle.com/iaas/Content/bigdata/create-cluster.htm#cluster-plan-shape), this value is the maximum memory in GBs each node can be scaled-up to. This value is not used for nodes with fixed compute shapes.
		* `max_ocpus_per_node` - (Applicable when policy_type=METRIC_BASED_VERTICAL_SCALING_POLICY) (Updatable) For nodes with [flexible compute shapes](https://docs.cloud.oracle.com/iaas/Content/bigdata/create-cluster.htm#cluster-plan-shape), this value is the maximum number of OCPUs each node can be scaled-up to. This value is not used for nodes with fixed compute shapes. 
		* `memory_step_size` - (Applicable when policy_type=METRIC_BASED_VERTICAL_SCALING_POLICY) (Updatable) For nodes with [flexible compute shapes](https://docs.cloud.oracle.com/iaas/Content/bigdata/create-cluster.htm#cluster-plan-shape), this value is the size of memory in GBs to add to each node during a scale-up event. This value is not used for nodes with fixed compute shapes.
		* `metric` - (Applicable when policy_type=METRIC_BASED_VERTICAL_SCALING_POLICY) (Updatable) Metric and threshold details for triggering an autoscale action.
			* `metric_type` - (Required when policy_type=METRIC_BASED_VERTICAL_SCALING_POLICY) (Updatable) Allowed values are CPU_UTILIZATION and MEMORY_UTILIZATION.
			* `threshold` - (Required when policy_type=METRIC_BASED_VERTICAL_SCALING_POLICY) (Updatable) An autoscale action is triggered when a performance metric exceeds a threshold.
				* `duration_in_minutes` - (Required when policy_type=METRIC_BASED_VERTICAL_SCALING_POLICY) (Updatable) This value is the minimum period of time the metric value exceeds the threshold value before the action is triggered. The value is in minutes.
				* `operator` - (Required when policy_type=METRIC_BASED_VERTICAL_SCALING_POLICY) (Updatable) The comparison operator to use. Options are greater than (GT) or less than (LT).
				* `value` - (Required when policy_type=METRIC_BASED_VERTICAL_SCALING_POLICY) (Updatable) Integer non-negative value. 0 < value < 100
		* `ocpu_step_size` - (Applicable when policy_type=METRIC_BASED_VERTICAL_SCALING_POLICY) (Updatable) For nodes with [flexible compute shapes](https://docs.cloud.oracle.com/iaas/Content/bigdata/create-cluster.htm#cluster-plan-shape), this value is the number of OCPUs to add to each node during a scale-up event. This value is not used for nodes with fixed compute shapes.
	* `schedule_details` - (Applicable when policy_type=SCHEDULE_BASED_HORIZONTAL_SCALING_POLICY | SCHEDULE_BASED_VERTICAL_SCALING_POLICY) (Updatable) Details of a horizontal scaling schedule.
		* `schedule_type` - (Optional) (Updatable) The type of schedule.
		* `time_and_horizontal_scaling_config` - (Optional) (Updatable) Time of day and horizontal scaling configuration.
			* `target_node_count` - (Optional) (Updatable) This value is the desired number of nodes in the cluster.
			* `time_recurrence` - (Optional) (Updatable) Day/time recurrence (specified following RFC 5545) at which to trigger autoscaling action. Currently only WEEKLY frequency is supported. Days of the week are specified using BYDAY field. Time of the day is specified using BYHOUR and BYMINUTE fields. Other fields are not supported. 
		* `time_and_vertical_scaling_config` - (Optional) (Updatable) Time of day and vertical scaling configuration
			* `target_memory_per_node` - (Optional) (Updatable) For nodes with [flexible compute shapes](https://docs.cloud.oracle.com/iaas/Content/bigdata/create-cluster.htm#cluster-plan-shape), this value is the desired memory in GBs on each node. This value is not used for nodes with fixed compute shapes. 
			* `target_ocpus_per_node` - (Optional) (Updatable) For nodes with [flexible compute shapes](https://docs.cloud.oracle.com/iaas/Content/bigdata/create-cluster.htm#cluster-plan-shape), this value is the desired OCPUs count on each node. This value is not used for nodes with fixed compute shapes. 
			* `target_shape` - (Optional) (Updatable) For nodes with [fixed compute shapes](https://docs.cloud.oracle.com/iaas/Content/bigdata/create-cluster.htm#cluster-plan-shape), this value is the desired shape of each node. This value is not used for nodes with flexible compute shapes. 
			* `time_recurrence` - (Optional) (Updatable) Day/time recurrence (specified following RFC 5545) at which to trigger autoscaling action. Currently only WEEKLY frequency is supported. Days of the week are specified using BYDAY field. Time of the day is specified using BYHOUR and BYMINUTE fields. Other fields are not supported. 
	* `timezone` - (Applicable when policy_type=SCHEDULE_BASED_HORIZONTAL_SCALING_POLICY | SCHEDULE_BASED_VERTICAL_SCALING_POLICY) (Updatable) The time zone of the execution schedule, in IANA time zone database name format


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `display_name` - A user-friendly name. The name does not have to be unique, and it may be changed. Avoid entering confidential information.
* `id` - The unique identifier for the autoscale configuration.
* `node_type` - A node type that is managed by an autoscale configuration. The only supported types are WORKER and COMPUTE_ONLY_WORKER.
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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Auto Scaling Configuration
	* `update` - (Defaults to 20 minutes), when updating the Auto Scaling Configuration
	* `delete` - (Defaults to 20 minutes), when destroying the Auto Scaling Configuration


## Import

AutoScalingConfiguration can be imported using the `id`, e.g.

```
$ terraform import oci_bds_auto_scaling_configuration.test_auto_scaling_configuration "bdsInstances/{bdsInstanceId}/autoScalingConfiguration/{autoScalingConfigurationId}" 
```

