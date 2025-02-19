---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_compute_gpu_memory_cluster"
sidebar_current: "docs-oci-datasource-core-compute_gpu_memory_cluster"
description: |-
  Provides details about a specific Compute Gpu Memory Cluster in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_compute_gpu_memory_cluster
This data source provides details about a specific Compute Gpu Memory Cluster resource in Oracle Cloud Infrastructure Core service.

Gets information about the specified compute GPU memory cluster


## Example Usage

```hcl
data "oci_core_compute_gpu_memory_cluster" "test_compute_gpu_memory_cluster" {
	#Required
	compute_gpu_memory_cluster_id = oci_core_compute_gpu_memory_cluster.test_compute_gpu_memory_cluster.id
}
```

## Argument Reference

The following arguments are supported:

* `compute_gpu_memory_cluster_id` - (Required) The OCID of the compute GPU memory cluster.


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

