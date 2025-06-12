---
subcategory: "Capacity Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_capacity_management_internal_occm_demand_signal_items"
sidebar_current: "docs-oci-datasource-capacity_management-internal_occm_demand_signal_items"
description: |-
  Provides the list of Internal Occm Demand Signal Items in Oracle Cloud Infrastructure Capacity Management service
---

# Data Source: oci_capacity_management_internal_occm_demand_signal_items
This data source provides the list of Internal Occm Demand Signal Items in Oracle Cloud Infrastructure Capacity Management service.

This internal API will list the detailed information about the resources demanded as part of the demand signal.


## Example Usage

```hcl
data "oci_capacity_management_internal_occm_demand_signal_items" "test_internal_occm_demand_signal_items" {
	#Required
	compartment_id = var.compartment_id
	occ_customer_group_id = oci_capacity_management_occ_customer_group.test_occ_customer_group.id

	#Optional
	demand_signal_namespace = var.internal_occm_demand_signal_item_demand_signal_namespace
	occm_demand_signal_id = oci_capacity_management_occm_demand_signal.test_occm_demand_signal.id
	resource_name = oci_cloud_guard_resource.test_resource.name
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ocid of the compartment or tenancy in which resources are to be listed. This will also be used for authorization purposes.
* `demand_signal_namespace` - (Optional) A query parameter to filter the list of demand signal details based on the namespace. 
* `occ_customer_group_id` - (Required) The customer group ocid by which we would filter the list.
* `occm_demand_signal_id` - (Optional) A query parameter to filter the list of demand signal items based on a demand signal id. 
* `resource_name` - (Optional) A query parameter to filter the list of demand signal details based on the resource name. 


## Attributes Reference

The following attributes are exported:

* `internal_occm_demand_signal_item_collection` - The list of internal_occm_demand_signal_item_collection.

### InternalOccmDemandSignalItem Reference

The following attributes are exported:

* `items` - An array of items containing detailed information about different resource demanded as part of a demand signal. 
	* `availability_domain` - The name of the availability domain for which you want to request the Oracle Cloud Infrastructure resource. 
	* `compartment_id` - The OCID of the tenancy from which the demand signal item was created. 
	* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
	* `demand_signal_catalog_resource_id` - The OCID of the corresponding demand signal catalog resource. 
	* `demand_signal_id` - The OCID of the demand signal under which this item will be grouped. 
	* `demand_signal_namespace` - The name of the Oracle Cloud Infrastructure service in consideration for demand signal submission. For example: COMPUTE, NETWORK, GPU etc. 
	* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
	* `id` - The OCID of the demand signal resource request. 
	* `notes` - This field will serve as notes section for you. You can use this section to convey a message to Oracle Cloud Infrastructure regarding your resource request.

		NOTE: The previous value gets overwritten with the new one for this once updated. 
	* `occ_customer_group_id` - The OCID of the customer group in which the demand signal is created. 
	* `quantity` - The quantity of the resource that you want to demand from Oracle Cloud Infrastructure or return to OCI. 
	* `region` - The name of region for which you want to request the Oracle Cloud Infrastructure resource. 
	* `request_type` - The type of request (DEMAND or RETURN) made against a particular demand signal item. 
	* `resource_name` - The name of the Oracle Cloud Infrastructure resource that you want to request. 
	* `resource_properties` - A map of various properties associated with the Oracle Cloud Infrastructure resource. 
	* `state` - The current lifecycle state of the demand signal item. 
	* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
	* `target_compartment_id` - The ocid of the tenancy for which you want to request the Oracle Cloud Infrastructure resource for. This is an optional parameter. 
	* `time_needed_before` - the date before which you would ideally like the Oracle Cloud Infrastructure resource to be delivered to you. 

