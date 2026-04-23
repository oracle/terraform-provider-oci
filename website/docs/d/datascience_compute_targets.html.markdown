---
subcategory: "Data Science"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datascience_compute_targets"
sidebar_current: "docs-oci-datasource-datascience-compute_targets"
description: |-
  Provides the list of Compute Targets in Oracle Cloud Infrastructure Data Science service
---

# Data Source: oci_datascience_compute_targets
This data source provides the list of Compute Targets in Oracle Cloud Infrastructure Data Science service.

List all compute targets in the specified compartment. Supports queries on various other parameters in the query alongside compartmentId (must be included).


## Example Usage

```hcl
data "oci_datascience_compute_targets" "test_compute_targets" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.compute_target_display_name
	id = var.compute_target_id
	state = var.compute_target_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) <b>Filter</b> results by the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) <b>Filter</b> results by its user-friendly name.
* `id` - (Optional) <b>Filter</b> results by [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). Must be an OCID of the correct type for the resource type. 
* `state` - (Optional) <b>Filter</b> results by the specified lifecycle state. Must be a valid state for the resource type. 


## Attributes Reference

The following attributes are exported:

* `compute_targets` - The list of compute_targets.

### ComputeTarget Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment associated with the compute target.
* `compute_configuration_details` - Configuration details of the targeted compute.
	* `compute_type` - The type of compute.
	* `instance_configuration` - The compute target instance configuration details for managed compute cluster type compute target.
		* `boot_volume_size_in_gbs` - The size of the boot volume to attach to the instance.
		* `instance_shape` - The shape used to launch the instances in compute target. Supported shapes can be retrieved using compute target shapes api.
		* `instance_shape_details` - Instance shape configuration for managed compute cluster type compute target. Specify only when a flex shape is selected.
			* `memory_in_gbs` - The total amount of memory allocated to the instance, in gigabytes.
			* `ocpus` - The total number of OCPUs allocated to the instance.
	* `scaling_policy` - The scaling policy to apply to managed compute cluster type compute target.
		* `auto_scaling_policies` - The list of autoscaling policy details.
			* `auto_scaling_policy_type` - The type of autoscaling policy.
			* `initial_instance_count` - For a threshold-based autoscaling policy, this value is the initial number of instances to launch in the managed compute cluster type compute target immediately after autoscaling is enabled. Note that anytime this value is updated, the number of instances will be reset to this value. After autoscaling retrieves performance metrics, the number of instances is automatically adjusted from this initial number to a number that is based on the limits that you set. 
			* `maximum_instance_count` - For a threshold-based autoscaling policy, this value is the maximum number of instances the managed compute cluster type compute target is allowed to increase to (scale out).
			* `minimum_instance_count` - For a threshold-based autoscaling policy, this value is the minimum number of instances the managed compute cluster type compute target is allowed to decrease to (scale in).
			* `rules` - The list of autoscaling policy rules.
				* `metric_expression_rule_type` - The metric expression for creating the alarm used to trigger autoscaling actions on the managed compute cluster type compute target .
				* `metric_type` - Metric type
				* `scale_in_configuration` - The scaling configuration for the predefined metric expression rule.
					* `instance_count_adjustment` - The value is used for adjusting the count of instances by.
					* `pending_duration` - The period of time that the condition defined in the alarm must persist before the alarm state changes from "OK" to "FIRING" or vice versa. For example, a value of 5 minutes means that the alarm must persist in breaching the condition for five minutes before the alarm updates its state to "FIRING"; likewise, the alarm must persist in not breaching the condition for five minutes before the alarm updates its state to "OK." The duration is specified as a string in ISO 8601 format (PT10M for ten minutes or PT1H for one hour). Minimum: PT3M. Maximum: PT1H. Default: PT3M. 
					* `query` - The Monitoring Query Language (MQL) expression to evaluate for the alarm. The Alarms feature of the Monitoring service interprets results for each returned time series as Boolean values, where zero represents false and a non-zero value represents true. A true value means that the trigger rule condition has been met. The query must specify a metric, statistic, interval, and trigger rule (threshold or absence). Supported values for interval: 1m-60m (also 1h). You can optionally specify dimensions and grouping functions. Supported grouping functions: grouping(), groupBy(). Example of threshold alarm: ``` CPUUtilization[1m]{resourceId = "Compute_Target_OCID"}.grouping().mean() < 25 CPUUtilization[1m]{resourceId = "Compute_Target_OCID"}.grouping().mean() > 75 ``` 
					* `scaling_configuration_type` - The type of scaling configuration.
					* `threshold` - A metric value at which the scaling operation will be triggered.
				* `scale_out_configuration` - The scaling configuration for the predefined metric expression rule.
					* `instance_count_adjustment` - The value is used for adjusting the count of instances by.
					* `pending_duration` - The period of time that the condition defined in the alarm must persist before the alarm state changes from "OK" to "FIRING" or vice versa. For example, a value of 5 minutes means that the alarm must persist in breaching the condition for five minutes before the alarm updates its state to "FIRING"; likewise, the alarm must persist in not breaching the condition for five minutes before the alarm updates its state to "OK." The duration is specified as a string in ISO 8601 format (PT10M for ten minutes or PT1H for one hour). Minimum: PT3M. Maximum: PT1H. Default: PT3M. 
					* `query` - The Monitoring Query Language (MQL) expression to evaluate for the alarm. The Alarms feature of the Monitoring service interprets results for each returned time series as Boolean values, where zero represents false and a non-zero value represents true. A true value means that the trigger rule condition has been met. The query must specify a metric, statistic, interval, and trigger rule (threshold or absence). Supported values for interval: 1m-60m (also 1h). You can optionally specify dimensions and grouping functions. Supported grouping functions: grouping(), groupBy(). Example of threshold alarm: ``` CPUUtilization[1m]{resourceId = "Compute_Target_OCID"}.grouping().mean() < 25 CPUUtilization[1m]{resourceId = "Compute_Target_OCID"}.grouping().mean() > 75 ``` 
					* `scaling_configuration_type` - The type of scaling configuration.
					* `threshold` - A metric value at which the scaling operation will be triggered.
		* `cool_down_in_seconds` - For threshold-based autoscaling policies, this value is the minimum period of time to wait between scaling actions. The cooldown period gives the system time to stabilize before rescaling. The minimum value is 600 seconds, which is also the default. The cooldown period starts when the managed compute cluster type compute target  becomes ACTIVE after the scaling operation. 
		* `instance_count` - The number of instances for the managed compute cluster type compute target.
		* `is_enabled` - Whether the autoscaling policy is enabled.
		* `policy_type` - The type of scaling policy.
* `compute_target_system_data` - System data of the compute target.
	* `compute_type` - Type of compute target.
	* `current_instance_count` - Current count of the instances in managed compute cluster type compute target.
* `created_by` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user who created the compute target.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - A short description of the compute target.
* `display_name` - A user-friendly display name for the resource.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compute target.
* `lifecycle_details` - Details about the state of the compute target.
* `metadata` - Metadata for the compute target. The size of metadata must be less than 2048 bytes. Key should be under 32 characters. Key should contain only letters, digits and underscore (_) Key should start with a letter. Key should have at least 2 characters. Key should not end with underscore eg. `TEST_` Key if added cannot be empty. Value can be empty. No specific size limits on individual Values. But overall metadata is limited to 2048 bytes. Key can't be reserved Compute Target metadata. 
* `state` - The state of the compute target.
* `time_created` - The date and time the resource was created in the timestamp format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: 2020-08-06T21:10:29.41Z 

