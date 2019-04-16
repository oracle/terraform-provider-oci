---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_autoscaling_auto_scaling_configurations"
sidebar_current: "docs-oci-datasource-autoscaling-auto_scaling_configurations"
description: |-
  Provides the list of Auto Scaling Configurations in Oracle Cloud Infrastructure Autoscaling service
---

# Data Source: oci_autoscaling_auto_scaling_configurations
This data source provides the list of Auto Scaling Configurations in Oracle Cloud Infrastructure Autoscaling service.

Lists AutoScalingConfigurations in the specific compartment.


## Example Usage

```hcl
data "oci_autoscaling_auto_scaling_configurations" "test_auto_scaling_configurations" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	display_name = "${var.auto_scaling_configuration_display_name}"
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
	* `id` - The OCID of resource that the AutoScalingConfiguration will manage. 
	* `type` - Indicates type of derived class
* `compartment_id` - The OCID of the compartment containing the AutoScalingConfiguration. 
* `cool_down_in_seconds` - The minimum period of time between scaling actions. The default is 300 seconds. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name for the AutoScalingConfiguration. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the AutoScalingConfiguration
* `is_enabled` - If the AutoScalingConfiguration is enabled
* `policies` - AutoScalingConfiguration policy definitions 
* `time_created` - The date and time the AutoScalingConfiguration was created, in the format defined by RFC3339. Example: `2016-08-25T21:10:29.600Z` 

