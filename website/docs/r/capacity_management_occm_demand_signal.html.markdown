---
subcategory: "Capacity Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_capacity_management_occm_demand_signal"
sidebar_current: "docs-oci-resource-capacity_management-occm_demand_signal"
description: |-
  Provides the Occm Demand Signal resource in Oracle Cloud Infrastructure Capacity Management service
---

# oci_capacity_management_occm_demand_signal
This resource provides the Occm Demand Signal resource in Oracle Cloud Infrastructure Capacity Management service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/occcm/latest/OccmDemandSignal

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/capacity_management

This is a post API to create occm demand signal.


## Example Usage

```hcl
resource "oci_capacity_management_occm_demand_signal" "test_occm_demand_signal" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.occm_demand_signal_display_name

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.occm_demand_signal_description
	freeform_tags = {"bar-key"= "value"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the tenancy where we would like to create a demand signal. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) A short description of the demand signal. 
* `display_name` - (Required) (Updatable) The user-friendly name of the demand signal. Does not have to be unique. Avoid entering anything confidential. 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the tenancy from which the request to create the demand signal was made. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - Meaningful text about the demand signal. 
* `display_name` - The display name of the demand signal. 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The OCID of the demand signal. 
* `lifecycle_details` - The different states associated with a demand signal. 

	CREATED -> A demand signal is by default created in this state.  SUBMITTED -> Once you have reviewed the details of the demand signal, you can transition it to SUBMITTED state so that Oracle Cloud Infrastructure can start working on it. DELETED -> You can delete a demand signal as long as it is in either CREATED or SUBMITTED state. IN_PROGRESS -> Once Oracle Cloud Infrastructure starts working on a given demand signal. They transition it to IN_PROGRESS. CANCELLED -> Oracle Cloud Infrastructure can transition the demand signal to this state. COMPLETED -> Oracle Cloud Infrastructure will transition the demand signal to COMPLETED state once the quantities which Oracle Cloud Infrastructure committed to deliver to you has been delivered. 
* `state` - The current lifecycle state of the resource. 
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time when the demand signal was created. 
* `time_updated` - The time when the demand signal was last updated. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Occm Demand Signal
	* `update` - (Defaults to 20 minutes), when updating the Occm Demand Signal
	* `delete` - (Defaults to 20 minutes), when destroying the Occm Demand Signal


## Import

OccmDemandSignals can be imported using the `id`, e.g.

```
$ terraform import oci_capacity_management_occm_demand_signal.test_occm_demand_signal "id"
```

