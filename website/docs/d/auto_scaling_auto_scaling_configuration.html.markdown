---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_auto_scaling_auto_scaling_configuration"
sidebar_current: "docs-oci-datasource-auto_scaling-auto_scaling_configuration"
description: |-
  Provides details about a specific Auto Scaling Configuration in Oracle Cloud Infrastructure Auto Scaling service
---

# Data Source: oci_auto_scaling_auto_scaling_configuration
This data source provides details about a specific Auto Scaling Configuration resource in Oracle Cloud Infrastructure Auto Scaling service.

Get AutoScalingConfiguration

## Example Usage

```hcl
data "oci_auto_scaling_auto_scaling_configuration" "test_auto_scaling_configuration" {
	#Required
	auto_scaling_configuration_id = "${oci_auto_scaling_auto_scaling_configuration.test_auto_scaling_configuration.id}"
}
```

## Argument Reference

The following arguments are supported:

* `auto_scaling_configuration_id` - (Required) The OCID of the auto scaling configuration.


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

