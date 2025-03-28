---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_compute_host"
sidebar_current: "docs-oci-datasource-core-compute_host"
description: |-
  Provides details about a specific Compute Host in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_compute_host
This data source provides details about a specific Compute Host resource in Oracle Cloud Infrastructure Core service.

Gets information about the specified compute host


## Example Usage

```hcl
data "oci_core_compute_host" "test_compute_host" {
	#Required
	compute_host_id = oci_core_compute_host.test_compute_host.id
}
```

## Argument Reference

The following arguments are supported:

* `compute_host_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compute host. 


## Attributes Reference

The following attributes are exported:

* `additional_data` - Additional data that can be exposed to the customer.  Will include raw fault codes for strategic customers 
* `availability_domain` - The availability domain of the compute host.  Example: `Uocm:US-CHICAGO-1-AD-2` 
* `capacity_reservation_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the Capacity Reserver that is currently on host 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the compartment. This should always be the root compartment. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `fault_domain` - A fault domain is a grouping of hardware and infrastructure within an availability domain. Each availability domain contains three fault domains. Fault domains let you distribute your instances so that they are not on the same physical hardware within a single availability domain. A hardware failure or Compute hardware maintenance that affects one fault domain does not affect instances in other fault domains.

	This field is the Fault domain of the host 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `health` - The heathy state of the host 
* `hpc_island_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for Customer-unique HPC Island 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the Customer-unique host 
* `impacted_component_details` - A list that contains impacted components related to an unhealthy host. An impacted component will be a  free-form structure of key values pairs that will provide more or less details based on data tiering 
* `instance_id` - The public [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the Virtual Machine or Bare Metal instance 
* `lifecycle_details` - A free-form description detailing why the host is in its current state. 
* `local_block_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for Customer-unique Local Block 
* `gpu_memory_fabric_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for Customer-unique GPU Memory Fabric
* `network_block_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for Customer-unique Network Block 
* `shape` - The shape of host 
* `state` - The lifecycle state of the host 
* `time_created` - The date and time that the compute host record was created, in the format defined by [RFC3339](https://tools .ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time that the compute host record was updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 

