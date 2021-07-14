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
	policy {
		#Required
		policy_type = var.auto_scaling_configuration_policy_policy_type
		rules {
			#Required
			action = var.auto_scaling_configuration_policy_rules_action
			metric {
				#Required
				metric_type = var.auto_scaling_configuration_policy_rules_metric_metric_type
				threshold {
					#Required
					duration_in_minutes = var.auto_scaling_configuration_policy_rules_metric_threshold_duration_in_minutes
					operator = var.auto_scaling_configuration_policy_rules_metric_threshold_operator
					value = var.auto_scaling_configuration_policy_rules_metric_threshold_value
				}
			}
		}
	}

	#Optional
	display_name = var.auto_scaling_configuration_display_name
}
```

## Argument Reference

The following arguments are supported:

* `bds_instance_id` - (Required) The OCID of the cluster.
* `cluster_admin_password` - (Required) (Updatable) Base-64 encoded password for the cluster (and Cloudera Manager) admin user.
* `display_name` - (Optional) (Updatable) A user-friendly name. The name does not have to be unique, and it may be changed. Avoid entering confidential information.
* `is_enabled` - (Required) (Updatable) Whether the autoscale configuration is enabled.
* `node_type` - (Required) A node type that is managed by an autoscale configuration. The only supported type is WORKER.
* `policy` - (Required) (Updatable) Policy definitions for the autoscale configuration.
	* `policy_type` - (Required) (Updatable) Types of autoscale policies. Options are SCHEDULE-BASED or THRESHOLD-BASED. (Only THRESHOLD-BASED is supported in this release.)
	* `rules` - (Required) (Updatable) The list of rules for autoscaling. If an action has multiple rules, the last rule in the array will be applied.
		* `action` - (Required) (Updatable) The valid value are CHANGE_SHAPE_SCALE_UP or CHANGE_SHAPE_SCALE_DOWN.
		* `metric` - (Required) (Updatable) Metric and threshold details for triggering an autoscale action.
			* `metric_type` - (Required) (Updatable) Allowed value is CPU_UTILIZATION.
			* `threshold` - (Required) (Updatable) An autoscale action is triggered when a performance metric meets or exceeds a threshold.
				* `duration_in_minutes` - (Required) (Updatable) This value is the minimum period of time the metric value meets or exceeds the threshold value before the action is triggered. The value is in minutes.
				* `operator` - (Required) (Updatable) The comparison operator to use. Options are greater than (GT) or less than (LT).
				* `value` - (Required) (Updatable) Integer non-negative value. 0 < value < 100


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Auto Scaling Configuration
	* `update` - (Defaults to 20 minutes), when updating the Auto Scaling Configuration
	* `delete` - (Defaults to 20 minutes), when destroying the Auto Scaling Configuration


## Import

AutoScalingConfiguration can be imported using the `id`, e.g.

```
$ terraform import oci_bds_auto_scaling_configuration.test_auto_scaling_configuration "bdsInstances/{bdsInstanceId}/autoScalingConfiguration/{autoScalingConfigurationId}" 
```

