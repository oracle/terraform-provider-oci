---
subcategory: "Demand Signal"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_demand_signal_occ_metric_alarm"
sidebar_current: "docs-oci-resource-demand_signal-occ_metric_alarm"
description: |-
  Provides the Occ Metric Alarm resource in Oracle Cloud Infrastructure Demand Signal service
---

# oci_demand_signal_occ_metric_alarm
This resource provides the Occ Metric Alarm resource in Oracle Cloud Infrastructure Demand Signal service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/occds/latest/OccMetricAlarm

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/demand_signal

Creates a new OccMetricAlarm resource in the specified compartment with the provided configuration details.


## Example Usage

```hcl
resource "oci_demand_signal_occ_metric_alarm" "test_occ_metric_alarm" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.occ_metric_alarm_display_name
	frequency = var.occ_metric_alarm_frequency
	is_active = var.occ_metric_alarm_is_active
	resource_configuration {
		#Required
		resource = var.occ_metric_alarm_resource_configuration_resource
		usage_type = var.occ_metric_alarm_resource_configuration_usage_type

		#Optional
		compute_hw_generation = var.occ_metric_alarm_resource_configuration_compute_hw_generation
		hw_generation = var.occ_metric_alarm_resource_configuration_hw_generation
		link_role = var.occ_metric_alarm_resource_configuration_link_role
		node_type = var.occ_metric_alarm_resource_configuration_node_type
		occ_metric_alarm_provider = var.occ_metric_alarm_resource_configuration_occ_metric_alarm_provider
		shape = var.occ_metric_alarm_resource_configuration_shape
		storage_type = var.occ_metric_alarm_resource_configuration_storage_type
	}
	threshold = var.occ_metric_alarm_threshold

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.occ_metric_alarm_description
	freeform_tags = {"Department"= "Finance"}
	state = var.occ_metric_alarm_state
	subscribers = var.occ_metric_alarm_subscribers
	threshold_type = var.occ_metric_alarm_threshold_type
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) Compartment OCID in which the alarm is created.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) Optional description for the alarm.
* `display_name` - (Required) (Updatable) Human-readable name for the alarm.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `frequency` - (Required) (Updatable) Frequency at which notifications should be sent.
* `is_active` - (Required) (Updatable) Alarm active status.
* `resource_configuration` - (Required) Configuration for a given 'resource'
	* `compute_hw_generation` - (Applicable when resource=COMPUTE) The hardware generation of the compute resource.
	* `hw_generation` - (Applicable when resource=EXADATA) The hardware generation of the Exadata system.
	* `link_role` - (Applicable when resource=NETWORK) The role of the link in the network.
	* `node_type` - (Required when resource=EXADATA) The type of node in the Exadata system.
	* `occ_metric_alarm_provider` - (Required when resource=NETWORK) The provider of the network service.
	* `resource` - (Required) Resources like COMPUTE, STORAGE, EXADATA etc.
	* `shape` - (Required when resource=COMPUTE) The shape of the compute instance.
	* `storage_type` - (Required when resource=STORAGE) The type of storage resource.
	* `usage_type` - (Required) The type of usage for the resource.
* `state` - (Optional) (Updatable) The current lifecycle state of the resource.
* `subscribers` - (Optional) (Updatable) List of topic OCIDs for notifications.
* `threshold` - (Required) (Updatable) Threshold at which alarm must be triggered.
* `threshold_type` - (Optional) (Updatable) Units in which threshold is being stored.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - Compartment OCID in which the alarm is created.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - Optional description for the alarm.
* `display_name` - Human-readable name for the alarm.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `frequency` - Frequency at which notifications should be sent.
* `id` - Unique OCID for this alarm configuration.
* `is_active` - Alarm active status.
* `resource_configuration` - Configuration for a given 'resource'
	* `compute_hw_generation` - The hardware generation of the compute resource.
	* `hw_generation` - The hardware generation of the Exadata system.
	* `link_role` - The role of the link in the network.
	* `node_type` - The type of node in the Exadata system.
	* `occ_metric_alarm_provider` - The provider of the network service.
	* `resource` - Resources like COMPUTE, STORAGE, EXADATA etc.
	* `shape` - The shape of the compute instance.
	* `storage_type` - The type of storage resource.
	* `usage_type` - The type of usage for the resource.
* `state` - The current lifecycle state of the resource.
* `subscribers` - List of topic OCIDs for notifications.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `threshold` - Threshold at which alarm must be triggered.
* `threshold_type` - Units in which threshold is being stored.
* `time_created` - Creation timestamp (RFC 3339).
* `time_updated` - Last update timestamp (RFC 3339).

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Occ Metric Alarm
	* `update` - (Defaults to 20 minutes), when updating the Occ Metric Alarm
	* `delete` - (Defaults to 20 minutes), when destroying the Occ Metric Alarm


## Import

OccMetricAlarms can be imported using the `id`, e.g.

```
$ terraform import oci_demand_signal_occ_metric_alarm.test_occ_metric_alarm "id"
```

