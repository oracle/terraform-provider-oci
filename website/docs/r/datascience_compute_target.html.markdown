---
subcategory: "Data Science"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datascience_compute_target"
sidebar_current: "docs-oci-resource-datascience-compute_target"
description: |-
  Provides the Compute Target resource in Oracle Cloud Infrastructure Data Science service
---

# oci_datascience_compute_target
This resource provides the Compute Target resource in Oracle Cloud Infrastructure Data Science service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/data-science/latest/ComputeTarget

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/datascience

Creates a new compute target resource.

## Example Usage

```hcl
resource "oci_datascience_compute_target" "test_compute_target" {
	#Required
	compartment_id = var.compartment_id
	compute_configuration_details {
		#Required
		compute_type = var.compute_target_compute_configuration_details_compute_type
		instance_configuration {
			#Required
			instance_shape = var.compute_target_compute_configuration_details_instance_configuration_instance_shape

			#Optional
			boot_volume_size_in_gbs = var.compute_target_compute_configuration_details_instance_configuration_boot_volume_size_in_gbs
			instance_shape_details {

				#Optional
				memory_in_gbs = var.compute_target_compute_configuration_details_instance_configuration_instance_shape_details_memory_in_gbs
				ocpus = var.compute_target_compute_configuration_details_instance_configuration_instance_shape_details_ocpus
			}
		}

		#Optional
		scaling_policy {
			#Required
			policy_type = var.compute_target_compute_configuration_details_scaling_policy_policy_type

			#Optional
			auto_scaling_policies {
				#Required
				auto_scaling_policy_type = var.compute_target_compute_configuration_details_scaling_policy_auto_scaling_policies_auto_scaling_policy_type
				initial_instance_count = var.compute_target_compute_configuration_details_scaling_policy_auto_scaling_policies_initial_instance_count
				maximum_instance_count = var.compute_target_compute_configuration_details_scaling_policy_auto_scaling_policies_maximum_instance_count
				minimum_instance_count = var.compute_target_compute_configuration_details_scaling_policy_auto_scaling_policies_minimum_instance_count
				rules {
					#Required
					metric_expression_rule_type = var.compute_target_compute_configuration_details_scaling_policy_auto_scaling_policies_rules_metric_expression_rule_type
					scale_in_configuration {

						#Optional
						instance_count_adjustment = var.compute_target_compute_configuration_details_scaling_policy_auto_scaling_policies_rules_scale_in_configuration_instance_count_adjustment
						pending_duration = var.compute_target_compute_configuration_details_scaling_policy_auto_scaling_policies_rules_scale_in_configuration_pending_duration
						query = var.compute_target_compute_configuration_details_scaling_policy_auto_scaling_policies_rules_scale_in_configuration_query
						scaling_configuration_type = var.compute_target_compute_configuration_details_scaling_policy_auto_scaling_policies_rules_scale_in_configuration_scaling_configuration_type
						threshold = var.compute_target_compute_configuration_details_scaling_policy_auto_scaling_policies_rules_scale_in_configuration_threshold
					}
					scale_out_configuration {

						#Optional
						instance_count_adjustment = var.compute_target_compute_configuration_details_scaling_policy_auto_scaling_policies_rules_scale_out_configuration_instance_count_adjustment
						pending_duration = var.compute_target_compute_configuration_details_scaling_policy_auto_scaling_policies_rules_scale_out_configuration_pending_duration
						query = var.compute_target_compute_configuration_details_scaling_policy_auto_scaling_policies_rules_scale_out_configuration_query
						scaling_configuration_type = var.compute_target_compute_configuration_details_scaling_policy_auto_scaling_policies_rules_scale_out_configuration_scaling_configuration_type
						threshold = var.compute_target_compute_configuration_details_scaling_policy_auto_scaling_policies_rules_scale_out_configuration_threshold
					}

					#Optional
					metric_type = var.compute_target_compute_configuration_details_scaling_policy_auto_scaling_policies_rules_metric_type
				}
			}
			cool_down_in_seconds = var.compute_target_compute_configuration_details_scaling_policy_cool_down_in_seconds
			instance_count = var.compute_target_compute_configuration_details_scaling_policy_instance_count
			is_enabled = var.compute_target_compute_configuration_details_scaling_policy_is_enabled
		}
	}

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.compute_target_description
	display_name = var.compute_target_display_name
	freeform_tags = {"Department"= "Finance"}
	metadata = var.compute_target_metadata
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where you want to create the compute target.
* `compute_configuration_details` - (Required) (Updatable) Configuration details of the targeted compute.
	* `compute_type` - (Required) (Updatable) The type of compute.
	* `instance_configuration` - (Required) (Updatable) The compute target instance configuration details for managed compute cluster type compute target.
		* `boot_volume_size_in_gbs` - (Optional) (Updatable) The size of the boot volume to attach to the instance.
		* `instance_shape` - (Required) (Updatable) The shape used to launch the instances in compute target. Supported shapes can be retrieved using compute target shapes api.
		* `instance_shape_details` - (Optional) (Updatable) Instance shape configuration for managed compute cluster type compute target. Specify only when a flex shape is selected.
			* `memory_in_gbs` - (Optional) (Updatable) The total amount of memory allocated to the instance, in gigabytes.
			* `ocpus` - (Optional) (Updatable) The total number of OCPUs allocated to the instance.
	* `scaling_policy` - (Optional) (Updatable) The scaling policy to apply to managed compute cluster type compute target.
		* `auto_scaling_policies` - (Required when policy_type=AUTOSCALING) (Updatable) The list of autoscaling policy details.
			* `auto_scaling_policy_type` - (Required) (Updatable) The type of autoscaling policy.
			* `initial_instance_count` - (Required) (Updatable) For a threshold-based autoscaling policy, this value is the initial number of instances to launch in the managed compute cluster type compute target immediately after autoscaling is enabled. Note that anytime this value is updated, the number of instances will be reset to this value. After autoscaling retrieves performance metrics, the number of instances is automatically adjusted from this initial number to a number that is based on the limits that you set. 
			* `maximum_instance_count` - (Required) (Updatable) For a threshold-based autoscaling policy, this value is the maximum number of instances the managed compute cluster type compute target is allowed to increase to (scale out).
			* `minimum_instance_count` - (Required) (Updatable) For a threshold-based autoscaling policy, this value is the minimum number of instances the managed compute cluster type compute target is allowed to decrease to (scale in).
			* `rules` - (Required) (Updatable) The list of autoscaling policy rules.
				* `metric_expression_rule_type` - (Required) (Updatable) The metric expression for creating the alarm used to trigger autoscaling actions on the managed compute cluster type compute target .
				* `metric_type` - (Required when metric_expression_rule_type=PREDEFINED_EXPRESSION) (Updatable) Metric type
				* `scale_in_configuration` - (Required) (Updatable) The scaling configuration for the predefined metric expression rule.
					* `instance_count_adjustment` - (Applicable when metric_expression_rule_type=CUSTOM_EXPRESSION | PREDEFINED_EXPRESSION) (Updatable) The value is used for adjusting the count of instances by.
					* `pending_duration` - (Applicable when metric_expression_rule_type=CUSTOM_EXPRESSION | PREDEFINED_EXPRESSION) (Updatable) The period of time that the condition defined in the alarm must persist before the alarm state changes from "OK" to "FIRING" or vice versa. For example, a value of 5 minutes means that the alarm must persist in breaching the condition for five minutes before the alarm updates its state to "FIRING"; likewise, the alarm must persist in not breaching the condition for five minutes before the alarm updates its state to "OK." The duration is specified as a string in ISO 8601 format (PT10M for ten minutes or PT1H for one hour). Minimum: PT3M. Maximum: PT1H. Default: PT3M. 
					* `query` - (Required when metric_expression_rule_type=CUSTOM_EXPRESSION) (Updatable) The Monitoring Query Language (MQL) expression to evaluate for the alarm. The Alarms feature of the Monitoring service interprets results for each returned time series as Boolean values, where zero represents false and a non-zero value represents true. A true value means that the trigger rule condition has been met. The query must specify a metric, statistic, interval, and trigger rule (threshold or absence). Supported values for interval: 1m-60m (also 1h). You can optionally specify dimensions and grouping functions. Supported grouping functions: grouping(), groupBy(). Example of threshold alarm: ``` CPUUtilization[1m]{resourceId = "Compute_Target_OCID"}.grouping().mean() < 25 CPUUtilization[1m]{resourceId = "Compute_Target_OCID"}.grouping().mean() > 75 ``` 
					* `scaling_configuration_type` - (Required when metric_expression_rule_type=CUSTOM_EXPRESSION | PREDEFINED_EXPRESSION) (Updatable) The type of scaling configuration.
					* `threshold` - (Required when metric_expression_rule_type=PREDEFINED_EXPRESSION) (Updatable) A metric value at which the scaling operation will be triggered.
				* `scale_out_configuration` - (Required) (Updatable) The scaling configuration for the predefined metric expression rule.
					* `instance_count_adjustment` - (Applicable when metric_expression_rule_type=CUSTOM_EXPRESSION | PREDEFINED_EXPRESSION) (Updatable) The value is used for adjusting the count of instances by.
					* `pending_duration` - (Applicable when metric_expression_rule_type=CUSTOM_EXPRESSION | PREDEFINED_EXPRESSION) (Updatable) The period of time that the condition defined in the alarm must persist before the alarm state changes from "OK" to "FIRING" or vice versa. For example, a value of 5 minutes means that the alarm must persist in breaching the condition for five minutes before the alarm updates its state to "FIRING"; likewise, the alarm must persist in not breaching the condition for five minutes before the alarm updates its state to "OK." The duration is specified as a string in ISO 8601 format (PT10M for ten minutes or PT1H for one hour). Minimum: PT3M. Maximum: PT1H. Default: PT3M. 
					* `query` - (Required when metric_expression_rule_type=CUSTOM_EXPRESSION) (Updatable) The Monitoring Query Language (MQL) expression to evaluate for the alarm. The Alarms feature of the Monitoring service interprets results for each returned time series as Boolean values, where zero represents false and a non-zero value represents true. A true value means that the trigger rule condition has been met. The query must specify a metric, statistic, interval, and trigger rule (threshold or absence). Supported values for interval: 1m-60m (also 1h). You can optionally specify dimensions and grouping functions. Supported grouping functions: grouping(), groupBy(). Example of threshold alarm: ``` CPUUtilization[1m]{resourceId = "Compute_Target_OCID"}.grouping().mean() < 25 CPUUtilization[1m]{resourceId = "Compute_Target_OCID"}.grouping().mean() > 75 ``` 
					* `scaling_configuration_type` - (Required when metric_expression_rule_type=CUSTOM_EXPRESSION | PREDEFINED_EXPRESSION) (Updatable) The type of scaling configuration.
					* `threshold` - (Required when metric_expression_rule_type=PREDEFINED_EXPRESSION) (Updatable) A metric value at which the scaling operation will be triggered.
		* `cool_down_in_seconds` - (Applicable when policy_type=AUTOSCALING) (Updatable) For threshold-based autoscaling policies, this value is the minimum period of time to wait between scaling actions. The cooldown period gives the system time to stabilize before rescaling. The minimum value is 300 seconds, which is also the default. The cooldown period starts when the managed compute cluster type compute target  becomes ACTIVE after the scaling operation. 
		* `instance_count` - (Required when policy_type=FIXED_SIZE) (Updatable) The number of instances for the managed compute cluster type compute target.
		* `is_enabled` - (Applicable when policy_type=AUTOSCALING) (Updatable) Whether the autoscaling policy is enabled.
		* `policy_type` - (Required) (Updatable) The type of scaling policy.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) A short description of the compute target.
* `display_name` - (Optional) (Updatable) A user-friendly display name for the resource.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `metadata` - (Optional) (Updatable) Metadata for the compute target. The size of metadata must be less than 2048 bytes. Key should be under 32 characters. Key should contain only letters, digits and underscore (_) Key should start with a letter. Key should have at least 2 characters. Key should not end with underscore eg. `TEST_` Key if added cannot be empty. Value can be empty. No specific size limits on individual Values. But overall metadata is limited to 2048 bytes. Key can't be reserved Compute Target metadata. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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
		* `cool_down_in_seconds` - For threshold-based autoscaling policies, this value is the minimum period of time to wait between scaling actions. The cooldown period gives the system time to stabilize before rescaling. The minimum value is 300 seconds, which is also the default. The cooldown period starts when the managed compute cluster type compute target  becomes ACTIVE after the scaling operation. 
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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Compute Target
	* `update` - (Defaults to 20 minutes), when updating the Compute Target
	* `delete` - (Defaults to 20 minutes), when destroying the Compute Target


## Import

ComputeTargets can be imported using the `id`, e.g.

```
$ terraform import oci_datascience_compute_target.test_compute_target "id"
```

