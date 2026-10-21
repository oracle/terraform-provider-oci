---
subcategory: "Generative AI"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_generative_ai_dedicated_ai_cluster_allocatable_gpus"
sidebar_current: "docs-oci-datasource-generative_ai-dedicated_ai_cluster_allocatable_gpus"
description: |-
  Provides the list of Dedicated Ai Cluster Allocatable Gpus in Oracle Cloud Infrastructure Generative AI service
---

# Data Source: oci_generative_ai_dedicated_ai_cluster_allocatable_gpus
This data source provides the list of Dedicated Ai Cluster Allocatable Gpus in Oracle Cloud Infrastructure Generative AI service.

Lists the allocatable GPUs for launching dedicated AI clusters tied to a specific capacity reservation id.

## Example Usage

```hcl
data "oci_generative_ai_dedicated_ai_cluster_allocatable_gpus" "test_dedicated_ai_cluster_allocatable_gpus" {
	#Required
	capacity_reservation_id = oci_generative_ai_capacity_reservation.test_capacity_reservation.id
	compartment_id = var.compartment_id

	#Optional
	compute_shape = var.dedicated_ai_cluster_allocatable_gpu_compute_shape
}
```

## Argument Reference

The following arguments are supported:

* `capacity_reservation_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the capacity reservation.
* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
* `compute_shape` - (Optional) The compute shape to use to filter allocatable GPUs. For example, `BM.GPU.A10.4`.


## Attributes Reference

The following attributes are exported:

* `allocatable_gpu_collection` - The list of allocatable_gpu_collection.

### DedicatedAiClusterAllocatableGpu Reference

The following attributes are exported:

* `items` - List of allocatable GPU summaries.
	* `compute_shape` - The compute shape for the allocatable GPUs.
	* `gpu_allocations` - The GPU allocation details for the compute shape. Each item represents one underlying capacity allocation.
		* `available_gpus` - The number of available GPUs in this allocation.
		* `used_gpus` - The number of used GPUs in this allocation.
	* `max_allocatable_gpu_units` - The maximum number of GPU units that can be allocated to a single dedicated AI cluster for the compute shape without requiring additional provisioning.
	* `total_available_gpu_units` - The total number of available GPU units for the compute shape.
	* `total_used_gpu_units` - The total number of used GPU units for the compute shape.
