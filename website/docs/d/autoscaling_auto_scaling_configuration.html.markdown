---
subcategory: "Auto Scaling"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_autoscaling_auto_scaling_configuration"
sidebar_current: "docs-oci-datasource-autoscaling-auto_scaling_configuration"
description: |-
  Provides details about a specific Auto Scaling Configuration in Oracle Cloud Infrastructure Auto Scaling service
---

# Data Source: oci_autoscaling_auto_scaling_configuration
This data source provides details about a specific Auto Scaling Configuration resource in Oracle Cloud Infrastructure Auto Scaling service.

Gets information about the specified autoscaling configuration.

## Example Usage

```hcl
data "oci_autoscaling_auto_scaling_configuration" "test_auto_scaling_configuration" {
	#Required
	auto_scaling_configuration_id = oci_autoscaling_auto_scaling_configuration.test_auto_scaling_configuration.id
}
```

## Argument Reference

The following arguments are supported:

* `auto_scaling_configuration_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the autoscaling configuration.


## Attributes Reference

The following attributes are exported:

* `auto_scaling_resources` - A resource that is managed by an autoscaling configuration. The only supported type is `instancePool`.

	Each instance pool can have one autoscaling configuration. 
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource that is managed by the autoscaling configuration. 
	* `type` - The type of resource.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the autoscaling configuration. 
* `cool_down_in_seconds` - For threshold-based autoscaling policies, this value is the minimum period of time to wait between scaling actions. The cooldown period gives the system time to stabilize before rescaling. The minimum value is 300 seconds, which is also the default. The cooldown period starts when the instance pool reaches the running state.

	For schedule-based autoscaling policies, this value is not used. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the autoscaling configuration.
* `is_enabled` - Whether the autoscaling configuration is enabled.
* `max_resource_count` - The maximum number of resources to scale out to.
* `min_resource_count` - The minimum number of resources to scale in to.
* `policies` - Autoscaling policy definitions for the autoscaling configuration. An autoscaling policy defines the criteria that trigger autoscaling actions and the actions to take. 
	* `capacity` - The capacity requirements of the autoscaling policy.
		* `initial` - For a threshold-based autoscaling policy, this value is the initial number of instances to launch in the instance pool immediately after autoscaling is enabled. After autoscaling retrieves performance metrics, the number of instances is automatically adjusted from this initial number to a number that is based on the limits that you set.

			For a schedule-based autoscaling policy, this value is the target pool size to scale to when executing the schedule that's defined in the autoscaling policy. 
		* `max` - For a threshold-based autoscaling policy, this value is the maximum number of instances the instance pool is allowed to increase to (scale out).

			For a schedule-based autoscaling policy, this value is not used. 
		* `min` - For a threshold-based autoscaling policy, this value is the minimum number of instances the instance pool is allowed to decrease to (scale in).

			For a schedule-based autoscaling policy, this value is not used. 
	* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
	* `execution_schedule` - The schedule for executing the autoscaling policy.
		* `expression` - A cron expression that represents the time at which to execute the autoscaling policy.

			Cron expressions have this format: `<second> <minute> <hour> <day of month> <month> <day of week> <year>`

			You can use special characters that are supported with the Quartz cron implementation.

			You must specify `0` as the value for seconds.

			Example: `0 15 10 ? * *` 
		* `timezone` - The time zone for the execution schedule.
		* `type` - The type of execution schedule.
	* `id` - The ID of the autoscaling policy that is assigned after creation.
	* `is_enabled` - Whether the autoscaling policy is enabled.
	* `policy_type` - The type of autoscaling policy.
	* `resource_action` - An action to run on a resource, such as stopping or starting an instance pool.
		* `action` - 
		* `action_type` - The category of action to run on the resource.
	* `rules` - 
		* `action` - The action to take when autoscaling is triggered. 
			* `type` - The type of action to take.
			* `value` - To scale out (increase the number of instances), provide a positive value. To scale in (decrease the number of instances), provide a negative value. 
		* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
		* `id` - ID of the condition that is assigned after creation.
		* `metric` - 
			* `metric_compartment_id` - The OCID of the compartment containing the metrics.
			* `metric_source` - Source of the metric data for creating the alarm used to trigger autoscaling actions.

				The following values are supported:
				* `COMPUTE_AGENT`: CPU or memory metrics emitted by the Compute Instance Monitoring plugin.
				* `CUSTOM_QUERY`: A custom Monitoring Query Language (MQL) expression. 
			* `metric_type` - Metric type example: CPU_UTILIZATION, MEMORY_UTILIZATION
			* `namespace` - The namespace for the query.
			* `pending_duration` - The period of time that the condition defined in the alarm must persist before the alarm state changes from "OK" to "FIRING" or vice versa. For example, a value of 5 minutes means that the alarm must persist in breaching the condition for five minutes before the alarm updates its state to "FIRING"; likewise, the alarm must persist in not breaching the condition for five minutes before the alarm updates its state to "OK."

				The duration is specified as a string in ISO 8601 format (`PT10M` for ten minutes or `PT1H` for one hour). Minimum: PT3M. Maximum: PT1H. Default: PT3M. 
			* `query` - The Monitoring Query Language (MQL) expression to evaluate for the alarm. The Alarms feature of the Monitoring service interprets results for each returned time series as Boolean values, where zero represents false and a non-zero value represents true. A true value means that the trigger rule condition has been met. The query must specify a metric, statistic, interval, and trigger rule (threshold or absence). Supported values for interval: `1m`-`60m` (also `1h`). You can optionally specify dimensions and grouping functions. Supported grouping functions: `grouping()`, `groupBy()`.

				Example of threshold alarm:

				-----

				CpuUtilization[1m]{availabilityDomain="cumS:PHX-AD-1"}.groupBy(availabilityDomain).percentile(0.9) > 85

				----- 
			* `resource_group` - The resource group for the query.
			* `threshold` - 
				* `operator` - The comparison operator to use. Options are greater than (`GT`), greater than or equal to (`GTE`), less than (`LT`), and less than or equal to (`LTE`). 
				* `value` - 
	* `time_created` - The date and time the autoscaling configuration was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `time_created` - The date and time the autoscaling configuration was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 

