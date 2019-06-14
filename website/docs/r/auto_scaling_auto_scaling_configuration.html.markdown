---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_autoscaling_auto_scaling_configuration"
sidebar_current: "docs-oci-resource-autoscaling-auto_scaling_configuration"
description: |-
  Provides the Auto Scaling Configuration resource in Oracle Cloud Infrastructure Auto Scaling service
---

# oci_autoscaling_auto_scaling_configuration
This resource provides the Auto Scaling Configuration resource in Oracle Cloud Infrastructure Auto Scaling service.

Creates an autoscaling configuration.

## Example Usage

```hcl
resource "oci_autoscaling_auto_scaling_configuration" "test_auto_scaling_configuration" {
	#Required
	auto_scaling_resources {
		#Required
		id = "${var.auto_scaling_configuration_auto_scaling_resources_id}"
		type = "${var.auto_scaling_configuration_auto_scaling_resources_type}"
	}
	compartment_id = "${var.compartment_id}"
	policies {
		#Required
		capacity {
			#Required
			initial = "${var.auto_scaling_configuration_policies_capacity_initial}"
			max = "${var.auto_scaling_configuration_policies_capacity_max}"
			min = "${var.auto_scaling_configuration_policies_capacity_min}"
		}
		policy_type = "${var.auto_scaling_configuration_policies_policy_type}"
		rules {
			#Required
			action {
				#Required
				type = "${var.auto_scaling_configuration_policies_rules_action_type}"
				value = "${var.auto_scaling_configuration_policies_rules_action_value}"
			}
			metric {
				#Required
				metric_type = "${var.auto_scaling_configuration_policies_rules_metric_metric_type}"
				threshold {
					#Required
					operator = "${var.auto_scaling_configuration_policies_rules_metric_threshold_operator}"
					value = "${var.auto_scaling_configuration_policies_rules_metric_threshold_value}"
				}
			}

			display_name = "${var.auto_scaling_configuration_policies_rules_display_name}"
		}

		#Optional
		display_name = "${var.auto_scaling_configuration_policies_display_name}"
	}

	#Optional
	cool_down_in_seconds = "${var.auto_scaling_configuration_cool_down_in_seconds}"
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = "${var.auto_scaling_configuration_display_name}"
	freeform_tags = {"Department"= "Finance"}
	is_enabled = "${var.auto_scaling_configuration_is_enabled}"
}
```

## Argument Reference

The following arguments are supported:

* `auto_scaling_resources` - (Required) 
	* `id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource that is managed by the autoscaling configuration. 
	* `type` - (Required) The type of resource.
* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the autoscaling configuration. The autoscaling configuration and the instance pool that it manages must be in the same compartment. 
* `cool_down_in_seconds` - (Optional) (Updatable) The minimum period of time to wait between scaling actions. The cooldown period gives the system time to stabilize before rescaling. The minimum value is 300 seconds, which is also the default. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `is_enabled` - (Optional) (Updatable) Whether the autoscaling configuration is enabled.
* `policies` - (Required) Autoscaling policy definitions for the autoscaling configuration. An autoscaling policy defines the criteria that trigger autoscaling actions and the actions to take.

	Each autoscaling configuration can have one autoscaling policy. 
	* `capacity` - (Required) The capacity requirements of the autoscaling policy.
		* `initial` - (Required) The initial number of instances to launch in the instance pool immediately after autoscaling is enabled. After autoscaling retrieves performance metrics, the number of instances is automatically adjusted from this initial number to a number that is based on the limits that you set. 
		* `max` - (Required) The maximum number of instances the instance pool is allowed to increase to (scale out).
		* `min` - (Required) The minimum number of instances the instance pool is allowed to decrease to (scale in).
	* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
	* `policy_type` - (Required) The type of autoscaling policy.
	* `rules` - (Required) 
		* `action` - (Required) 
			* `type` - (Required) The type of action to take.
			* `value` - (Required) To scale out (increase the number of instances), provide a positive value. To scale in (decrease the number of instances), provide a negative value. 
		* `display_name` - (Required) A user-friendly name. Does not have to be unique. Avoid entering confidential information. This value is not changeable through Terraform.
		* `metric` - (Required) 
			* `metric_type` - (Required) 
			* `threshold` - (Required) 
				* `operator` - (Required) The comparison operator to use. Options are greater than (`GT`), greater than or equal to (`GTE`), less than (`LT`), and less than or equal to (`LTE`). 
				* `value` - (Required) 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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
* `policies` - Autoscaling policy definitions for the autoscaling configuration. An autoscaling policy defines the criteria that trigger autoscaling actions and the actions to take.

	Each autoscaling configuration can have one autoscaling policy. 
	* `capacity` - The capacity requirements of the autoscaling policy.
		* `initial` - The initial number of instances to launch in the instance pool immediately after autoscaling is enabled. After autoscaling retrieves performance metrics, the number of instances is automatically adjusted from this initial number to a number that is based on the limits that you set. 
		* `max` - The maximum number of instances the instance pool is allowed to increase to (scale out).
		* `min` - The minimum number of instances the instance pool is allowed to decrease to (scale in).
	* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
	* `id` - The ID of the autoscaling policy that is assigned after creation.
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

## Import

AutoScalingConfigurations can be imported using the `id`, e.g.

```
$ terraform import oci_autoscaling_auto_scaling_configuration.test_auto_scaling_configuration "id"
```

