---
subcategory: "Demand Signal"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_demand_signal_occ_demand_signal"
sidebar_current: "docs-oci-datasource-demand_signal-occ_demand_signal"
description: |-
  Provides details about a specific Occ Demand Signal in Oracle Cloud Infrastructure Demand Signal service
---

# Data Source: oci_demand_signal_occ_demand_signal
This data source provides details about a specific Occ Demand Signal resource in Oracle Cloud Infrastructure Demand Signal service.

Gets information about a OccDemandSignal.

## Example Usage

```hcl
data "oci_demand_signal_occ_demand_signal" "test_occ_demand_signal" {
	#Required
	occ_demand_signal_id = oci_demand_signal_occ_demand_signal.test_occ_demand_signal.id
}
```

## Argument Reference

The following arguments are supported:

* `occ_demand_signal_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the OccDemandSignal.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the OccDemandSignal.
* `is_active` - Indicator of whether to share the data with Oracle.
* `lifecycle_details` - A message that describes the current state of the OccDemandSignal in more detail. For example, can be used to provide actionable information for a resource in the Failed state. 
* `occ_demand_signals` - The OccDemandSignal data.
	* `resource_type` - The name of the resource for the data.
	* `units` - The units of the data.
	* `values` - The values of forecast.
		* `comments` - Space provided for users to make comments regarding the value.
		* `time_expected` - The date of the Demand Signal Value.
		* `value` - The Demand Signal Value.
* `state` - The current state of the OccDemandSignal.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the OccDemandSignal was created, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time the OccDemandSignal was updated, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 

