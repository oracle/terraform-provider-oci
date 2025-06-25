---
subcategory: "Capacity Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_capacity_management_internal_occm_demand_signal_delivery"
sidebar_current: "docs-oci-datasource-capacity_management-internal_occm_demand_signal_delivery"
description: |-
  Provides details about a specific Internal Occm Demand Signal Delivery in Oracle Cloud Infrastructure Capacity Management service
---

# Data Source: oci_capacity_management_internal_occm_demand_signal_delivery
This data source provides details about a specific Internal Occm Demand Signal Delivery resource in Oracle Cloud Infrastructure Capacity Management service.

This is an internal GET API to get the details of a demand signal delivery resource corresponding to a demand signal item.


## Example Usage

```hcl
data "oci_capacity_management_internal_occm_demand_signal_delivery" "test_internal_occm_demand_signal_delivery" {
	#Required
	occm_demand_signal_delivery_id = oci_capacity_management_occm_demand_signal_delivery.test_occm_demand_signal_delivery.id
}
```

## Argument Reference

The following arguments are supported:

* `occm_demand_signal_delivery_id` - (Required) The OCID of the demand signal delivery. 


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

