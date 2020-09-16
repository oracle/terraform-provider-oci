---
subcategory: "Auto Scaling"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_autoscaling_auto_scaling_configurations"
sidebar_current: "docs-oci-datasource-autoscaling-auto_scaling_configurations"
description: |-
  Provides the list of Auto Scaling Configurations in Oracle Cloud Infrastructure Auto Scaling service
---

# Data Source: oci_autoscaling_auto_scaling_configurations
This data source provides the list of Auto Scaling Configurations in Oracle Cloud Infrastructure Auto Scaling service.

Lists autoscaling configurations in the specifed compartment.


## Example Usage

```hcl
data "oci_autoscaling_auto_scaling_configurations" "test_auto_scaling_configurations" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.auto_scaling_configuration_display_name
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the resources monitored by the metric that you are searching for. Use tenancyId to search in the root compartment. 
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. 


## Attributes Reference

The following attributes are exported:

* `auto_scaling_configurations` - The list of auto_scaling_configurations.

### AutoScalingConfiguration Reference

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
* `policies` - AutoScaling policy definitions for the autoscaling configuration. An autoscaling policy defines the criteria that trigger autoscaling actions and the actions to take.
* `time_created` - The date and time the AutoScalingConfiguration was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 

