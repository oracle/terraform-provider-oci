---
subcategory: "Capacity Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_capacity_management_occm_demand_signal_item"
sidebar_current: "docs-oci-resource-capacity_management-occm_demand_signal_item"
description: |-
  Provides the Occm Demand Signal Item resource in Oracle Cloud Infrastructure Capacity Management service
---

# oci_capacity_management_occm_demand_signal_item
This resource provides the Occm Demand Signal Item resource in Oracle Cloud Infrastructure Capacity Management service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/occcm/latest/OccmDemandSignalItem

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/capacity_management

This API will create a demand signal item representing a resource request. This needs to be grouped under a demand signal.


## Example Usage

```hcl
resource "oci_capacity_management_occm_demand_signal_item" "test_occm_demand_signal_item" {
	#Required
	compartment_id = var.compartment_id
	demand_quantity = var.occm_demand_signal_item_demand_quantity
	demand_signal_catalog_resource_id = oci_cloud_guard_resource.test_resource.id
	demand_signal_id = oci_capacity_management_demand_signal.test_demand_signal.id
	region = var.occm_demand_signal_item_region
	request_type = var.occm_demand_signal_item_request_type
	resource_properties = var.occm_demand_signal_item_resource_properties
	time_needed_before = var.occm_demand_signal_item_time_needed_before

	#Optional
	availability_domain = var.occm_demand_signal_item_availability_domain
	defined_tags = {"foo-namespace.bar-key"= "value"}
	freeform_tags = {"bar-key"= "value"}
	notes = var.occm_demand_signal_item_notes
	target_compartment_id = oci_identity_compartment.test_compartment.id
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Optional) (Updatable) The name of the availability domain for which you want to request the Oracle Cloud Infrastructure resource. This is an optional parameter. 
* `compartment_id` - (Required) The OCID of the tenancy from which the demand signal item was created. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `demand_quantity` - (Required) (Updatable) The quantity of the resource that you want to demand from OCI. 
* `demand_signal_catalog_resource_id` - (Required) The OCID of the correponding demand signal catalog resource. 
* `demand_signal_id` - (Required) The OCID of the demand signal under which we need to create this item. 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `notes` - (Optional) (Updatable) This field will serve as notes section for you. You can use this section to convey a message to Oracle Cloud Infrastructure regarding your resource request.

	NOTE: The previous value gets overwritten with the new one for this once updated. 
* `region` - (Required) (Updatable) The name of region for which you want to request the Oracle Cloud Infrastructure resource. 
* `request_type` - (Required) The type of request (DEMAND or RETURN) that you want to make for this demand signal item. 
* `resource_properties` - (Required) (Updatable) A map of various properties associated with the Oracle Cloud Infrastructure resource. 
* `target_compartment_id` - (Optional) (Updatable) The OCID of the tenancy for which you want to request the Oracle Cloud Infrastructure resource for. This is an optional parameter. 
* `time_needed_before` - (Required) (Updatable) the date before which you would ideally like the Oracle Cloud Infrastructure resource to be delivered to you. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `availability_domain` - The name of the availability domain for which you want to request the Oracle Cloud Infrastructure resource. 
* `compartment_id` - The OCID of the tenancy from which the demand signal item was created. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `demand_quantity` - The quantity of the resource that you want to demand from OCI. 
* `demand_signal_catalog_resource_id` - The OCID of the corresponding demand signal catalog resource. 
* `demand_signal_id` - The OCID of the demand signal under which this item will be grouped. 
* `demand_signal_namespace` - The name of the Oracle Cloud Infrastructure service in consideration for demand signal submission. For example: COMPUTE, NETWORK, GPU etc. 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The OCID of the demand signal item. 
* `notes` - This field will serve as notes section for you. You can use this section to convey a message to Oracle Cloud Infrastructure regarding your resource request.

	NOTE: The previous value gets overwritten with the new one for this once updated. 
* `region` - The name of region for which you want to request the Oracle Cloud Infrastructure resource. 
* `request_type` - The type of request (DEMAND or RETURN) made against a particular demand signal item. 
* `resource_name` - The name of the Oracle Cloud Infrastructure resource that you want to request. 
* `resource_properties` - A map of various properties associated with the Oracle Cloud Infrastructure resource. 
* `state` - The current lifecycle state of the resource.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `target_compartment_id` - The OCID of the tenancy for which you want to request the Oracle Cloud Infrastructure resource for. This is an optional parameter. 
* `time_needed_before` - the date before which you would ideally like the Oracle Cloud Infrastructure resource to be delivered to you. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Occm Demand Signal Item
	* `update` - (Defaults to 20 minutes), when updating the Occm Demand Signal Item
	* `delete` - (Defaults to 20 minutes), when destroying the Occm Demand Signal Item


## Import

OccmDemandSignalItems can be imported using the `id`, e.g.

```
$ terraform import oci_capacity_management_occm_demand_signal_item.test_occm_demand_signal_item "id"
```

