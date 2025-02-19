---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_compute_gpu_memory_cluster_instances"
sidebar_current: "docs-oci-datasource-core-compute_gpu_memory_cluster_instances"
description: |-
  Provides the list of Compute Gpu Memory Cluster Instances in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_compute_gpu_memory_cluster_instances
This data source provides the list of Compute Gpu Memory Cluster Instances in Oracle Cloud Infrastructure Core service.

List all of the GPU memory cluster instances.


## Example Usage

```hcl
data "oci_core_compute_gpu_memory_cluster_instances" "test_compute_gpu_memory_cluster_instances" {
	#Required
	compute_gpu_memory_cluster_id = oci_core_compute_gpu_memory_cluster.test_compute_gpu_memory_cluster.id
}
```

## Argument Reference

The following arguments are supported:

* `compute_gpu_memory_cluster_id` - (Required) The OCID of the compute GPU memory cluster.


## Attributes Reference

The following attributes are exported:

* `compute_gpu_memory_cluster_instance_collection` - The list of compute_gpu_memory_cluster_instance_collection.

### ComputeGpuMemoryClusterInstance Reference

The following attributes are exported:

* `items` - The list of compute GPU memory cluster instances.
	* `availability_domain` - The availability domain of the GPU memory cluster instance. 
	* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the compartment compartment. 
	* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
	* `fault_domain` - The fault domain the GPU memory cluster instance is running in.
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the Customer-unique GPU memory cluster instance 
	* `instance_configuration_id` - Configuration to be used for this GPU Memory Cluster instance. 
	* `instance_shape` - The shape of an instance. The shape determines the number of CPUs, amount of memory,  and other resources allocated to the instance. The shape determines the number of CPUs,  the amount of memory, and other resources allocated to the instance. You can list all available shapes by calling [ListShapes](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Shape/ListShapes). 
	* `region` - The region that contains the availability domain the instance is running in.
	* `state` - The lifecycle state of the GPU memory cluster instance 
	* `time_created` - The date and time the GPU memory cluster instance was created.  Example: `2016-09-15T21:10:29.600Z` 

