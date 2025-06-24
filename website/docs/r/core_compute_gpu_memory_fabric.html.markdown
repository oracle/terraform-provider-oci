---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_compute_gpu_memory_fabric"
sidebar_current: "docs-oci-resource-core-compute_gpu_memory_fabric"
description: |-
  Provides the Compute Gpu Memory Fabric resource in Oracle Cloud Infrastructure Core service
---

# oci_core_compute_gpu_memory_fabric
This resource provides the Compute Gpu Memory Fabric resource in Oracle Cloud Infrastructure Core service.

Customer can update displayName and tags for compute GPU memory fabric record


## Example Usage

```hcl
resource "oci_core_compute_gpu_memory_fabric" "test_compute_gpu_memory_fabric" {
	#Required
	compute_gpu_memory_fabric_id = oci_core_compute_gpu_memory_fabric.test_compute_gpu_memory_fabric.id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.compute_gpu_memory_fabric_display_name
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the compartment. This should always be the root compartment. 
* `compute_gpu_memory_fabric_id` - (Required) The OCID of the compute GPU memory fabric.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Compute Gpu Memory Fabric
	* `update` - (Defaults to 20 minutes), when updating the Compute Gpu Memory Fabric
	* `delete` - (Defaults to 20 minutes), when destroying the Compute Gpu Memory Fabric


## Import

ComputeGpuMemoryFabrics can be imported using the `id`, e.g.

```
$ terraform import oci_core_compute_gpu_memory_fabric.test_compute_gpu_memory_fabric "id"
```

