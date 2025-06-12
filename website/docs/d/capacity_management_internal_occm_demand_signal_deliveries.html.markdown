---
subcategory: "Capacity Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_capacity_management_internal_occm_demand_signal_deliveries"
sidebar_current: "docs-oci-datasource-capacity_management-internal_occm_demand_signal_deliveries"
description: |-
  Provides the list of Internal Occm Demand Signal Deliveries in Oracle Cloud Infrastructure Capacity Management service
---

# Data Source: oci_capacity_management_internal_occm_demand_signal_deliveries
This data source provides the list of Internal Occm Demand Signal Deliveries in Oracle Cloud Infrastructure Capacity Management service.

This GET call is used to list all demand signal delivery resources within the customer group passed as a query parameter.


## Example Usage

```hcl
data "oci_capacity_management_internal_occm_demand_signal_deliveries" "test_internal_occm_demand_signal_deliveries" {
	#Required
	compartment_id = var.compartment_id
	occ_customer_group_id = oci_capacity_management_occ_customer_group.test_occ_customer_group.id

	#Optional
	id = var.internal_occm_demand_signal_delivery_id
	occm_demand_signal_item_id = oci_capacity_management_occm_demand_signal_item.test_occm_demand_signal_item.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ocid of the compartment or tenancy in which resources are to be listed. This will also be used for authorization purposes.
* `id` - (Optional) A query parameter to filter the list of demand signals based on it's OCID. 
* `occ_customer_group_id` - (Required) The customer group ocid by which we would filter the list.
* `occm_demand_signal_item_id` - (Optional) A query parameter to filter the list of demand signal items based on it's OCID. 


## Attributes Reference

The following attributes are exported:

* `internal_occm_demand_signal_delivery_collection` - The list of internal_occm_demand_signal_delivery_collection.

### InternalOccmDemandSignalDelivery Reference

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

