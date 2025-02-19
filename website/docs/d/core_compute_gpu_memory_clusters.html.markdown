---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_compute_gpu_memory_clusters"
sidebar_current: "docs-oci-datasource-core-compute_gpu_memory_clusters"
description: |-
  Provides the list of Compute Gpu Memory Clusters in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_compute_gpu_memory_clusters
This data source provides the list of Compute Gpu Memory Clusters in Oracle Cloud Infrastructure Core service.

List all of the compute GPU memory clusters.


## Example Usage

```hcl
data "oci_core_compute_gpu_memory_clusters" "test_compute_gpu_memory_clusters" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	availability_domain = var.compute_gpu_memory_cluster_availability_domain
	compute_cluster_id = oci_core_compute_cluster.test_compute_cluster.id
	compute_gpu_memory_cluster_id = oci_core_compute_gpu_memory_cluster.test_compute_gpu_memory_cluster.id
	display_name = var.compute_gpu_memory_cluster_display_name
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Optional) The name of the availability domain.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `compute_cluster_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compute cluster. A [compute cluster](https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/compute-clusters.htm) is a remote direct memory access (RDMA) network group. 
* `compute_gpu_memory_cluster_id` - (Optional) A filter to return only the listings that matches the given GPU memory cluster id. 
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. 


## Attributes Reference

The following attributes are exported:

* `compute_gpu_memory_cluster_collection` - The list of compute_gpu_memory_cluster_collection.

### ComputeGpuMemoryCluster Reference

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

