---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_auto_scaling_auto_scaling_configuration"
sidebar_current: "docs-oci-resource-auto_scaling-auto_scaling_configuration"
description: |-
  Provides the Auto Scaling Configuration resource in Oracle Cloud Infrastructure Auto Scaling service
---

# oci_auto_scaling_auto_scaling_configuration
This resource provides the Auto Scaling Configuration resource in Oracle Cloud Infrastructure Auto Scaling service.

Create an AutoScalingConfiguration

## Example Usage

```hcl
resource "oci_auto_scaling_auto_scaling_configuration" "test_auto_scaling_configuration" {
	#Required
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
	auto_scaling_resources {
		#Required
		id = "${var.auto_scaling_configuration_resource_id}"
		type = "${var.auto_scaling_configuration_resource_type}"
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

* `compartment_id` - (Required) The OCID of the compartment containing the AutoScalingConfiguration. 
* `cool_down_in_seconds` - (Optional) (Updatable) The minimum period of time between scaling actions. The default is 300 seconds. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name for the AutoScalingConfiguration. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `is_enabled` - (Optional) (Updatable) If the AutoScalingConfiguration is enabled
* `policies` - (Required) AutoScalingConfiguration policy definitions 
	* `capacity` - (Required) The capacity requirements of the Policy
		* `initial` - (Required) The initial size of the pool
		* `max` - (Required) The maximum size the pool is allowed to increase to
		* `min` - (Required) The minimum size the pool is allowed to decrease to
	* `display_name` - (Optional) A user-friendly name for the Policy. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
	* `policy_type` - (Required) Indicates type of Policy
	* `rules` - (Required) 
		* `action` - (Required) 
			* `type` - (Required) Action type to take
			* `value` - (Required) 
		* `display_name` - (Required) A user-friendly name for the AutoScalingConfiguration condition details. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
		* `metric` - (Required) 
			* `metric_type` - (Required) 
			* `threshold` - (Required) 
				* `operator` - (Required) Support for the following operators GT  - Greater than GTE - Greater than equal to LT  - Less than LTE - Less than equal to 
				* `value` - (Required) 
* `auto_scaling_resources` - (Required) 
	* `id` - (Required) The OCID of resource that the AutoScalingConfiguration will manage. 
	* `type` - (Required) Indicates type of derived class


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the AutoScalingConfiguration. 
* `cool_down_in_seconds` - The minimum period of time between scaling actions. The default is 300 seconds. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name for the AutoScalingConfiguration. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the AutoScalingConfiguration
* `is_enabled` - If the AutoScalingConfiguration is enabled
* `policies` - AutoScalingConfiguration policy definitions 
	* `capacity` - The capacity requirements of the Policy
		* `initial` - The initial size of the pool
		* `max` - The maximum size the pool is allowed to increase to
		* `min` - The minimum size the pool is allowed to decrease to
	* `display_name` - A user-friendly name for the Policy. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
	* `id` - The ID of the policy that is assigned after creation
	* `policy_type` - Indicates type of Policy
	* `rules` - 
		* `action` - 
			* `type` - Action type to take
			* `value` - 
		* `display_name` - A user-friendly name for the AutoScalingConfiguration condition details. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
		* `id` - Id of the condition that is assigned after creation
		* `metric` - 
			* `metric_type` - 
			* `threshold` - 
				* `operator` - Support for the following operators GT  - Greater than GTE - Greater than equal to LT  - Less than LTE - Less than equal to 
				* `value` - 
	* `time_created` - The date and time the AutoScalingConfiguration was created, in the format defined by RFC3339. Example: `2016-08-25T21:10:29.600Z` 
* `auto_scaling_resources` - 
	* `id` - The OCID of resource that the AutoScalingConfiguration will manage. 
	* `type` - Indicates type of derived class
* `time_created` - The date and time the AutoScalingConfiguration was created, in the format defined by RFC3339. Example: `2016-08-25T21:10:29.600Z` 

## Import

AutoScalingConfigurations can be imported using the `id`, e.g.

```
$ terraform import oci_auto_scaling_auto_scaling_configuration.test_auto_scaling_configuration "id"
```

