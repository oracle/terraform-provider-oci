---
subcategory: "Capacity Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_capacity_management_occ_handover_resource_blocks"
sidebar_current: "docs-oci-datasource-capacity_management-occ_handover_resource_blocks"
description: |-
  Provides the list of Occ Handover Resource Blocks in Oracle Cloud Infrastructure Capacity Management service
---

# Data Source: oci_capacity_management_occ_handover_resource_blocks
This data source provides the list of Occ Handover Resource Blocks in Oracle Cloud Infrastructure Capacity Management service.

List Occ Handover Resource blocks.


## Example Usage

```hcl
data "oci_capacity_management_occ_handover_resource_blocks" "test_occ_handover_resource_blocks" {

	#Optional
	compartment_id = var.compartment_id
	handover_date_greater_than_or_equal_to = var.occ_handover_resource_block_handover_date_greater_than_or_equal_to
	handover_date_less_than_or_equal_to = var.occ_handover_resource_block_handover_date_less_than_or_equal_to
	handover_resource_name = oci_cloud_guard_resource.test_resource.name
	namespace = var.occ_handover_resource_block_namespace
	occ_handover_resource_block_id = oci_capacity_management_occ_handover_resource_block.test_occ_handover_resource_block.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The OCID of the compartment or tenancy in which resources are to be listed.
* `handover_date_greater_than_or_equal_to` - (Optional) This filter helps in fetching all handed over resources for which the recordDate is greater than or equal to the startDate. 
* `handover_date_less_than_or_equal_to` - (Optional) This filter helps in fetching all handed over resources for which the recordDate is less than or equal to the endDate. 
* `handover_resource_name` - (Optional) A filter to return only the list of resources that match the name provided in this filter. 
* `namespace` - (Optional) The namespace by which we would filter the list.
* `occ_handover_resource_block_id` - (Optional) This filter helps in fetching the handed over resource for which the occHandoverResourceId is equal to the one provided here. 


## Attributes Reference

The following attributes are exported:

* `occ_handover_resource_block_collection` - The list of occ_handover_resource_block_collection.

### OccHandoverResourceBlock Reference

The following attributes are exported:

* `items` - An array of occ handover resource blocks. 
	* `associated_capacity_requests` - A list containing details about the capacity requests against which the resources were provisioned by oracle. 
		* `handover_quantity` - The total quantity of the bare metal hardware that was made available corresponding to the capacity request ocid. 
		* `occ_capacity_request_id` - The OCID of the capacity request against which the resources were provisioned. 
	* `compartment_id` - The OCID of the compartment where the resource block's are placed. 
	* `handover_date` - The date on which the resource was handed over to the customer. 
	* `handover_resource_name` - The name of the resource handed over by oracle.  For instance for compute namespace this will be the name of the bare metal hardware resource. 
	* `id` - The OCID of the resource block. 
	* `namespace` - The name of the Oracle Cloud Infrastructure service in consideration.  For example Compute, Exadata and so on. 
	* `occ_customer_group_id` - The OCID of the customer group for which the resources were provisioned. 
	* `placement_details` - Details like building, room and block where the resource was placed after provisioning in the datacenter. 
		* `availability_domain` - The availability domain (AD) for which the resources were provisioned. 
		* `block` - The block in the datacenter room where the resource was placed. 
		* `building` - The datacenter building where the resource was placed. 
		* `region` - The name of the region for which the resources were provisioned. 
		* `room` - The name of the room in the dataacenter building where the resource was placed. 
		* `workload_type` - The type of workload to which these resources were provisioned. 
	* `total_handover_quantity` - The total quantity of the resource that was made available to the customer by Oracle. 

