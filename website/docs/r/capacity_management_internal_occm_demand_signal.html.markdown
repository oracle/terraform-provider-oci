---
subcategory: "Capacity Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_capacity_management_internal_occm_demand_signal"
sidebar_current: "docs-oci-resource-capacity_management-internal_occm_demand_signal"
description: |-
  Provides the Internal Occm Demand Signal resource in Oracle Cloud Infrastructure Capacity Management service
---

# oci_capacity_management_internal_occm_demand_signal
This resource provides the Internal Occm Demand Signal resource in Oracle Cloud Infrastructure Capacity Management service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/occcm/latest/InternalOccmDemandSignal

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/capacity_management

This is a internal PUT API which shall be used to update the metadata of the demand signal.


## Example Usage

```hcl
resource "oci_capacity_management_internal_occm_demand_signal" "test_internal_occm_demand_signal" {
	#Required
	occm_demand_signal_id = oci_capacity_management_occm_demand_signal.test_occm_demand_signal.id

	#Optional
	lifecycle_details = var.internal_occm_demand_signal_lifecycle_details
}
```

## Argument Reference

The following arguments are supported:

* `lifecycle_details` - (Optional) (Updatable) The subset of demand signal states available for operators for updating the demand signal.

	IN_PROGRESS -> Transitions the demand signal to IN_PROGRESS state. REJECTED -> Transitions the demand signal to REJECTED state. COMPLETED -> This will transition the demand signal to COMPLETED state. 
* `occm_demand_signal_id` - (Required) The OCID of the demand signal. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the tenancy from which the request to create the demand signal was made. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - A short description about the demand signal. 
* `display_name` - The display name of the demand signal. 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The OCID of the demand signal. 
* `lifecycle_details` - The different states associated with a demand signal. 

	CREATED -> A demand signal is by default created in this state.  SUBMITTED -> Once you have reviewed the details of the demand signal, you can transition it to SUBMITTED state so that Oracle Cloud Infrastructure can start working on it. DELETED -> You can delete a demand signal as long as it is in either CREATED or SUBMITTED state. IN_PROGRESS -> Once Oracle Cloud Infrastructure starts working on a given demand signal. They transition it to IN_PROGRESS. REJECTED -> Oracle Cloud Infrastructure can transition the demand signal to this state if all the demand signal items of that demand signal are declined. COMPLETED -> Oracle Cloud Infrastructure will transition the demand signal to COMPLETED state once the quantities which Oracle Cloud Infrastructure committed to deliver to you has been delivered. 
* `occ_customer_group_id` - The OCID of the customer group in which the demand signal is created. 
* `state` - The current lifecycle state of the demand signal. 
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time when the demand signal was created. 
* `time_updated` - The time when the demand signal was last updated. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Internal Occm Demand Signal
	* `update` - (Defaults to 20 minutes), when updating the Internal Occm Demand Signal
	* `delete` - (Defaults to 20 minutes), when destroying the Internal Occm Demand Signal


## Import

InternalOccmDemandSignals can be imported using the `id`, e.g.

```
$ terraform import oci_capacity_management_internal_occm_demand_signal.test_internal_occm_demand_signal "internal/occmDemandSignals/{occmDemandSignalId}" 
```

