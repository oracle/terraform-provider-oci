---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_compute_gpu_memory_fabrics"
sidebar_current: "docs-oci-datasource-core-compute_gpu_memory_fabrics"
description: |-
  Provides the list of Compute Gpu Memory Fabrics in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_compute_gpu_memory_fabrics
This data source provides the list of Compute Gpu Memory Fabrics in Oracle Cloud Infrastructure Core service.

Lists the compute GPU memory fabrics that match the specified criteria and compartmentId.

## Example Usage

```hcl
data "oci_core_compute_gpu_memory_fabrics" "test_compute_gpu_memory_fabrics" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	availability_domain = var.compute_gpu_memory_fabric_availability_domain
	compute_gpu_memory_fabric_health = var.compute_gpu_memory_fabric_compute_gpu_memory_fabric_health
	compute_gpu_memory_fabric_id = oci_core_compute_gpu_memory_fabric.test_compute_gpu_memory_fabric.id
	compute_gpu_memory_fabric_lifecycle_state = var.compute_gpu_memory_fabric_compute_gpu_memory_fabric_lifecycle_state
	compute_hpc_island_id = oci_core_compute_hpc_island.test_compute_hpc_island.id
	compute_network_block_id = oci_core_compute_network_block.test_compute_network_block.id
	display_name = var.compute_gpu_memory_fabric_display_name
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Optional) The name of the availability domain.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `compute_gpu_memory_fabric_health` - (Optional) A filter to return ComputeGpuMemoryFabricSummary resources that match the given fabric health. 
* `compute_gpu_memory_fabric_id` - (Optional) A filter to return only the listings that matches the given GPU memory fabric id. 
* `compute_gpu_memory_fabric_lifecycle_state` - (Optional) A filter to return ComputeGpuMemoryFabricSummary resources that match the given lifecycle state. 
* `compute_hpc_island_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compute HPC island.
* `compute_network_block_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compute network block.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. 


## Attributes Reference

The following attributes are exported:

* `compute_gpu_memory_fabric_collection` - The list of compute_gpu_memory_fabric_collection.

### ComputeGpuMemoryFabric Reference

The following attributes are exported:

* `additional_data` - Additional data that can be exposed to the customer. Right now it will include the switch tray ids. 
* `available_host_count` - The total number of available bare metal hosts located in this compute GPU memory fabric.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the compartment. This should always be the root compartment. 
* `compute_hpc_island_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for Customer-unique HPC Island 
* `compute_local_block_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for Customer-unique Local Block 
* `compute_network_block_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for Customer-unique Network Block 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `fabric_health` - The health state of the GPU memory fabric 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `healthy_host_count` - The total number of healthy bare metal hosts located in this compute GPU memory fabric.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the Customer-unique GPU memory fabric 
* `state` - The lifecycle state of the GPU memory fabric 
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time that the compute GPU memory fabric record was created, in the format defined by [RFC3339] (https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `total_host_count` - The total number of bare metal hosts located in this compute GPU memory fabric.

