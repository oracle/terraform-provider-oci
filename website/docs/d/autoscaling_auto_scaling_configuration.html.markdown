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

* `auto_scaling_resources` - 
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource that is managed by the autoscaling configuration. 
	* `type` - The type of resource.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the autoscaling configuration. 
* `cool_down_in_seconds` - The minimum period of time to wait between scaling actions. The cooldown period gives the system time to stabilize before rescaling. The minimum value is 300 seconds, which is also the default. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the autoscaling configuration.
* `is_enabled` - Whether the autoscaling configuration is enabled.
* `max_resource_count` - The maximum number of resources to scale out to.
* `min_resource_count` - The minimum number of resources to scale in to.
* `policies` - Autoscaling policy definitions for the autoscaling configuration. An autoscaling policy defines the criteria that trigger autoscaling actions and the actions to take.

	Each autoscaling configuration can have one autoscaling policy. 
	* `capacity` - The capacity requirements of the autoscaling policy.
		* `initial` - The initial number of instances to launch in the instance pool immediately after autoscaling is enabled. After autoscaling retrieves performance metrics, the number of instances is automatically adjusted from this initial number to a number that is based on the limits that you set. 
		* `max` - The maximum number of instances the instance pool is allowed to increase to (scale out).
		* `min` - The minimum number of instances the instance pool is allowed to decrease to (scale in).
	* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
	* `execution_schedule` - 
		* `expression` - The value representing the execution schedule, as defined by cron format. 
		* `timezone` - Specifies the time zone the schedule is in.
		* `type` - The type of ExecutionSchedule.
	* `id` - The ID of the autoscaling policy that is assigned after creation.
	* `is_enabled` - Boolean field indicating whether this policy is enabled or not.
	* `policy_type` - The type of autoscaling policy.
	* `rules` - 
		* `action` - 
			* `type` - The type of action to take.
			* `value` - To scale out (increase the number of instances), provide a positive value. To scale in (decrease the number of instances), provide a negative value. 
		* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
		* `id` - ID of the condition that is assigned after creation.
		* `metric` - 
			* `metric_type` - 
			* `threshold` - 
				* `operator` - The comparison operator to use. Options are greater than (`GT`), greater than or equal to (`GTE`), less than (`LT`), and less than or equal to (`LTE`). 
				* `value` - 
	* `time_created` - The date and time the autoscaling configuration was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `time_created` - The date and time the AutoScalingConfiguration was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 

