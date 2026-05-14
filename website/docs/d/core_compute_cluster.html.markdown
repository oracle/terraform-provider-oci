---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_compute_cluster"
sidebar_current: "docs-oci-datasource-core-compute_cluster"
description: |-
  Provides details about a specific Compute Cluster in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_compute_cluster
This data source provides details about a specific Compute Cluster resource in Oracle Cloud Infrastructure Core service.

Gets information about a compute cluster. A [compute cluster](https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/compute-clusters.htm)
is a remote direct memory access (RDMA) network group.


## Example Usage

```hcl
data "oci_core_compute_cluster" "test_compute_cluster" {
	#Required
	compute_cluster_id = oci_core_compute_cluster.test_compute_cluster.id
}
```

## Argument Reference

The following arguments are supported:

* `compute_cluster_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compute cluster. A [compute cluster](https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/compute-clusters.htm) is a remote direct memory access (RDMA) network group.


## Attributes Reference

The following attributes are exported:

* `availability_domain` - The availability domain the compute cluster is running in.  Example: `Uocm:PHX-AD-1`
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the compute cluster.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compute cluster.
* `placement_constraint_details` - The details for providing placement constraints.
	* `hpc_island_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the HPC island for the compute cluster.

		This field cannot be updated after creation of the compute cluster.
	* `logical_placement_constraint` - The logical placement strategy to apply. Allowed values are `SINGLE_TIER`, `SINGLE_BLOCK`, and `PACKED_DISTRIBUTION_MULTI_BLOCK`.
	* `target_memory_fabric_ids` - The list of target GPU memory fabric OCIDs to constrain placement. Up to 15 items are allowed.

		If GMFs are passed in, the `hpcIslandId` must be set on the compute cluster, and the provided GMFs must belong to that same HPC island.

		The ordering of the array will be preserved. Ensure that all items in the array are unique.
	* `target_network_block_ids` - The list of target network block OCIDs to constrain placement. Up to 15 items are allowed.

		If `targetNetworkBlockIds` is provided, the `hpcIslandId` must be set on the compute cluster, and the provided network blocks must belong to that same HPC island.

		The ordering of the array will be preserved. Ensure that all items in the array are unique.
	* `type` - The type for the placement constraints. Supported value: `COMPUTE_CLUSTER`.
* `state` - The current state of the compute cluster.
* `time_created` - The date and time the compute cluster was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z`
* `time_updated` - The date and time the compute cluster was updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z`
