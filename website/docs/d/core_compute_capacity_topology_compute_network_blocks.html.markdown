---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_compute_capacity_topology_compute_network_blocks"
sidebar_current: "docs-oci-datasource-core-compute_capacity_topology_compute_network_blocks"
description: |-
  Provides the list of Compute Capacity Topology Compute Network Blocks in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_compute_capacity_topology_compute_network_blocks
This data source provides the list of Compute Capacity Topology Compute Network Blocks in Oracle Cloud Infrastructure Core service.

Lists compute network blocks in the specified compute capacity topology.

## Example Usage

```hcl
data "oci_core_compute_capacity_topology_compute_network_blocks" "test_compute_capacity_topology_compute_network_blocks" {
	#Required
	compute_capacity_topology_id = oci_core_compute_capacity_topology.test_compute_capacity_topology.id

	#Optional
	availability_domain = var.compute_capacity_topology_compute_network_block_availability_domain
	compartment_id = var.compartment_id
	compute_hpc_island_id = oci_core_compute_hpc_island.test_compute_hpc_island.id
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Optional) The name of the availability domain.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `compute_capacity_topology_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compute capacity topology.
* `compute_hpc_island_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compute HPC island.


## Attributes Reference

The following attributes are exported:

* `compute_network_block_collection` - The list of compute_network_block_collection.

### ComputeCapacityTopologyComputeNetworkBlock Reference

The following attributes are exported:

* `items` - The list of compute network blocks.
	* `compute_capacity_topology_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compute capacity topology.
	* `compute_hpc_island_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compute HPC island.
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compute network block.
	* `state` - The current state of the compute network block.
	* `time_created` - The date and time that the compute network block was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
	* `time_updated` - The date and time that the compute network block was updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
	* `total_compute_bare_metal_host_count` - The total number of compute bare metal hosts located in the compute network block.

