---
subcategory: "Demand Signal"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_demand_signal_occ_demand_signal"
sidebar_current: "docs-oci-resource-demand_signal-occ_demand_signal"
description: |-
  Provides the Occ Demand Signal resource in Oracle Cloud Infrastructure Demand Signal service
---

# oci_demand_signal_occ_demand_signal
This resource provides the Occ Demand Signal resource in Oracle Cloud Infrastructure Demand Signal service.

Creates a OccDemandSignal.

  Updates the data of an OccDemandSignal.

## Example Usage

```hcl
resource "oci_demand_signal_occ_demand_signal" "test_occ_demand_signal" {
	#Required
	compartment_id = var.compartment_id
	is_active = var.occ_demand_signal_is_active
	occ_demand_signal_id = var.occ_demand_signal_occ_demand_signal_id
	occ_demand_signals {
		#Required
		resource_type = var.occ_demand_signal_occ_demand_signals_resource_type
		units = var.occ_demand_signal_occ_demand_signals_units
		values {
			#Required
			time_expected = var.occ_demand_signal_occ_demand_signals_values_time_expected
			value = var.occ_demand_signal_occ_demand_signals_values_value

			#Optional
			comments = var.occ_demand_signal_occ_demand_signals_values_comments
		}
	}

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.occ_demand_signal_display_name
	freeform_tags = {"Department"= "Finance"}
	patch_operations {
		#Required
		operation = var.occ_demand_signal_patch_operations_operation
		selection = var.occ_demand_signal_patch_operations_selection

		#Optional
		from = var.occ_demand_signal_patch_operations_from
		position = var.occ_demand_signal_patch_operations_position
		selected_item = var.occ_demand_signal_patch_operations_selected_item
		value = var.occ_demand_signal_patch_operations_value
		values = var.occ_demand_signal_patch_operations_values
	}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to create the OccDemandSignal in. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `is_active` - (Required) (Updatable) Indicator of whether to share the data with Oracle.
* `occ_demand_signal_id` - (Required) 
* `occ_demand_signals` - (Required) The OccDemandSignal data.
	* `resource_type` - (Required) The name of the resource for the data.
	* `units` - (Required) The units of the data.
	* `values` - (Required) The values of forecast.
		* `comments` - (Optional) Space provided for users to make comments regarding the value.
		* `time_expected` - (Required) The date of the Demand Signal Value.
		* `value` - (Required) The Demand Signal Value.
* `patch_operations` - (Optional) (Updatable) 
	* `from` - (Required when operation=MOVE) (Updatable) 
	* `operation` - (Required) (Updatable) The operation can be one of these values: `INSERT`, `INSERT_MULTIPLE`, `MERGE`, `MOVE`, `PROHIBIT`, `REMOVE`, `REPLACE`, `REQUIRE`
	* `position` - (Applicable when operation=INSERT | INSERT_MULTIPLE | MOVE) (Updatable) 
	* `selected_item` - (Applicable when operation=INSERT | INSERT_MULTIPLE) (Updatable) 
	* `selection` - (Required) (Updatable) 
	* `value` - (Required when operation=INSERT | MERGE | PROHIBIT | REPLACE | REQUIRE) (Updatable) 
	* `values` - (Required when operation=INSERT_MULTIPLE) (Updatable) 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Occ Demand Signal
	* `update` - (Defaults to 20 minutes), when updating the Occ Demand Signal
	* `delete` - (Defaults to 20 minutes), when destroying the Occ Demand Signal


## Import

OccDemandSignals can be imported using the `id`, e.g.

```
$ terraform import oci_demand_signal_occ_demand_signal.test_occ_demand_signal "id"
```

