---
subcategory: "Demand Signal"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_demand_signal_occ_metric_alarm"
sidebar_current: "docs-oci-datasource-demand_signal-occ_metric_alarm"
description: |-
  Provides details about a specific Occ Metric Alarm in Oracle Cloud Infrastructure Demand Signal service
---

# Data Source: oci_demand_signal_occ_metric_alarm
This data source provides details about a specific Occ Metric Alarm resource in Oracle Cloud Infrastructure Demand Signal service.

Retrieves the specified OccMetricAlarm resource based on its identifier.


## Example Usage

```hcl
data "oci_demand_signal_occ_metric_alarm" "test_occ_metric_alarm" {
	#Required
	occ_metric_alarm_id = oci_demand_signal_occ_metric_alarm.test_occ_metric_alarm.id
}
```

## Argument Reference

The following arguments are supported:

* `occ_metric_alarm_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the OccMetricAlarm.


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

