---
subcategory: "Capacity Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_capacity_management_occm_demand_signal_item"
sidebar_current: "docs-oci-datasource-capacity_management-occm_demand_signal_item"
description: |-
  Provides details about a specific Occm Demand Signal Item in Oracle Cloud Infrastructure Capacity Management service
---

# Data Source: oci_capacity_management_occm_demand_signal_item
This data source provides details about a specific Occm Demand Signal Item resource in Oracle Cloud Infrastructure Capacity Management service.

This is a GET API to get the details of a demand signal item resource representing the details of the resource demanded by you.


## Example Usage

```hcl
data "oci_capacity_management_occm_demand_signal_item" "test_occm_demand_signal_item" {
	#Required
	occm_demand_signal_item_id = oci_capacity_management_occm_demand_signal_item.test_occm_demand_signal_item.id
}
```

## Argument Reference

The following arguments are supported:

* `occm_demand_signal_item_id` - (Required) The OCID of the demand signal item. 


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

