---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_compute_gpu_memory_cluster"
sidebar_current: "docs-oci-resource-core-compute_gpu_memory_cluster"
description: |-
  Provides the Compute Gpu Memory Cluster resource in Oracle Cloud Infrastructure Core service
---

# oci_core_compute_gpu_memory_cluster
This resource provides the Compute Gpu Memory Cluster resource in Oracle Cloud Infrastructure Core service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/iaas/latest/ComputeGpuMemoryCluster

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/

Create a compute GPU memory cluster instance on a specific compute GPU memory fabric


## Example Usage

```hcl
resource "oci_core_compute_gpu_memory_cluster" "test_compute_gpu_memory_cluster" {
	#Required
	availability_domain = var.compute_gpu_memory_cluster_availability_domain
	compartment_id = var.compartment_id
	compute_cluster_id = oci_core_compute_cluster.test_compute_cluster.id
	instance_configuration_id = oci_core_instance_configuration.test_instance_configuration.id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.compute_gpu_memory_cluster_display_name
	freeform_tags = {"Department"= "Finance"}
	gpu_memory_fabric_id = oci_core_gpu_memory_fabric.test_gpu_memory_fabric.id
	size = var.compute_gpu_memory_cluster_size
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Required) The availability domain of the GPU memory cluster. 
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the compute GPU memory cluster. compartment. 
* `compute_cluster_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compute cluster. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `gpu_memory_fabric_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the GPU memory fabric. 
* `instance_configuration_id` - (Required) (Updatable) Instance Configuration to be used for this GPU Memory Cluster 
* `size` - (Optional) (Updatable) The number of instances currently running in the GpuMemoryCluster 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `availability_domain` - The availability domain of the GPU memory cluster. 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the compute GPU memory cluster. 
* `compute_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compute cluster. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `gpu_memory_fabric_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the GPU memory fabric. 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the Customer-unique GPU memory cluster 
* `instance_configuration_id` - The OCID of the Instance Configuration used to source launch details for this instance.
* `size` - The number of instances currently running in the GpuMemoryCluster 
* `state` - The lifecycle state of the GPU memory cluster 
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the GPU memory cluster was created.  Example: `2016-09-15T21:10:29.600Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Compute Gpu Memory Cluster
	* `update` - (Defaults to 20 minutes), when updating the Compute Gpu Memory Cluster
	* `delete` - (Defaults to 20 minutes), when destroying the Compute Gpu Memory Cluster


## Import

ComputeGpuMemoryClusters can be imported using the `id`, e.g.

```
$ terraform import oci_core_compute_gpu_memory_cluster.test_compute_gpu_memory_cluster "id"
```

