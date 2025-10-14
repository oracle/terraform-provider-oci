---
subcategory: "Capacity Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_capacity_management_internal_occm_demand_signal_delivery"
sidebar_current: "docs-oci-resource-capacity_management-internal_occm_demand_signal_delivery"
description: |-
  Provides the Internal Occm Demand Signal Delivery resource in Oracle Cloud Infrastructure Capacity Management service
---

# oci_capacity_management_internal_occm_demand_signal_delivery
This resource provides the Internal Occm Demand Signal Delivery resource in Oracle Cloud Infrastructure Capacity Management service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/occcm/latest/InternalOccmDemandSignalDelivery

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/capacity_management

This is a post API which is used to create a demand signal delivery resource.
operationId: CreateInternalOccmDemandSignalDelivery
summary: A post call to create a demand signal delivery.


## Example Usage

```hcl
resource "oci_capacity_management_internal_occm_demand_signal_delivery" "test_internal_occm_demand_signal_delivery" {
	#Required
	accepted_quantity = var.internal_occm_demand_signal_delivery_accepted_quantity
	compartment_id = var.compartment_id
	demand_signal_id = oci_capacity_management_demand_signal.test_demand_signal.id
	demand_signal_item_id = oci_capacity_management_demand_signal_item.test_demand_signal_item.id
	occ_customer_group_id = oci_capacity_management_occ_customer_group.test_occ_customer_group.id

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	freeform_tags = {"bar-key"= "value"}
	justification = var.internal_occm_demand_signal_delivery_justification
	notes = var.internal_occm_demand_signal_delivery_notes
}
```

## Argument Reference

The following arguments are supported:

* `accepted_quantity` - (Required) (Updatable) The quantity of the resource that Oracle Cloud Infrastructure will supply to the customer. 
* `compartment_id` - (Required) The OCID of the tenancy from which the demand signal delivery resource is created. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `demand_signal_id` - (Required) The OCID of the demand signal under which this delivery will be grouped. 
* `demand_signal_item_id` - (Required) The OCID of the demand signal item corresponding to which this delivery is made. 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `justification` - (Optional) (Updatable) This field could be used by Oracle Cloud Infrastructure to communicate the reason for declining the request. 
* `notes` - (Optional) (Updatable) This field acts as a notes section for operators. 
* `occ_customer_group_id` - (Required) The OCID of the corresponding customer group to which this demand signal delivery resource belongs to. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `accepted_quantity` - The quantity of the resource that Oracle Cloud Infrastructure will supply to the customer. 
* `compartment_id` - The OCID of the tenancy from which the demand signal delivery resource is created. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `demand_signal_id` - The OCID of the demand signal under which this delivery will be grouped. 
* `demand_signal_item_id` - The OCID of the demand signal item corresponding to which this delivery is made. 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The OCID of this demand signal delivery resource. 
* `justification` - This field could be used by Oracle Cloud Infrastructure to communicate the reason for accepting or declining the request. 
* `lifecycle_details` - The enum values corresponding to the various states associated with the delivery resource.

	SUBMITTED -> The state where operators have started working and thinking on the quantity that Oracle Cloud Infrastructure can delivery for the corresponding demand signal item. IN_REVIEW -> The operators are waiting on approvals from different teams/folks in this state. ACCEPTED -> Oracle Cloud Infrastructure has accepted your resource request and will deliver the quantity as specified by acceptance quantity of this resource. DECLINED -> Oracle Cloud Infrastructure has declined you resource request. DELIVERED -> Oracle Cloud Infrastructure has delivered the accepted quantity to the customers.

	NOTE: The resource becomes visible to customers in ACCEPTED, DECLINED or DELIVERED state. 
* `notes` - This field acts as a notes section for operators. 
* `occ_customer_group_id` - The OCID of the corresponding customer group to which this demand signal delivery resource belongs to. 
* `state` - The current lifecycle state of the resource. 
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_delivered` - The date on which the Oracle Cloud Infrastructure delivered the resource to the customers. The default value for this will be the corresponding demand signal item resource's need by date. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Internal Occm Demand Signal Delivery
	* `update` - (Defaults to 20 minutes), when updating the Internal Occm Demand Signal Delivery
	* `delete` - (Defaults to 20 minutes), when destroying the Internal Occm Demand Signal Delivery


## Import

InternalOccmDemandSignalDeliveries can be imported using the `id`, e.g.

```
$ terraform import oci_capacity_management_internal_occm_demand_signal_delivery.test_internal_occm_demand_signal_delivery "internal/occmDemandSignalDeliveries/{occmDemandSignalDeliveryId}" 
```

