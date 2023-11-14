---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_compute_capacity_topology_compute_bare_metal_hosts"
sidebar_current: "docs-oci-datasource-core-compute_capacity_topology_compute_bare_metal_hosts"
description: |-
  Provides the list of Compute Capacity Topology Compute Bare Metal Hosts in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_compute_capacity_topology_compute_bare_metal_hosts
This data source provides the list of Compute Capacity Topology Compute Bare Metal Hosts in Oracle Cloud Infrastructure Core service.

Lists compute bare metal hosts in the specified compute capacity topology.

## Example Usage

```hcl
data "oci_core_compute_capacity_topology_compute_bare_metal_hosts" "test_compute_capacity_topology_compute_bare_metal_hosts" {
	#Required
	compute_capacity_topology_id = oci_core_compute_capacity_topology.test_compute_capacity_topology.id

	#Optional
	availability_domain = var.compute_capacity_topology_compute_bare_metal_host_availability_domain
	compartment_id = var.compartment_id
	compute_hpc_island_id = oci_core_compute_hpc_island.test_compute_hpc_island.id
	compute_local_block_id = oci_core_compute_local_block.test_compute_local_block.id
	compute_network_block_id = oci_core_compute_network_block.test_compute_network_block.id
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Optional) The name of the availability domain.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `compute_capacity_topology_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compute capacity topology.
* `compute_hpc_island_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compute HPC island.
* `compute_local_block_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compute local block.
* `compute_network_block_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compute network block.


## Attributes Reference

The following attributes are exported:

* `compute_bare_metal_host_collection` - The list of compute_bare_metal_host_collection.

### ComputeCapacityTopologyComputeBareMetalHost Reference

The following attributes are exported:

* `items` - The list of compute bare metal hosts.
	* `compute_capacity_topology_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compute capacity topology.
	* `compute_hpc_island_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compute HPC island.
	* `compute_local_block_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compute network block.
	* `compute_network_block_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compute local block.
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compute bare metal host.
	* `instance_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compute instance that runs on the compute bare metal host.
	* `instance_shape` - The shape of the compute instance that runs on the compute bare metal host.
	* `lifecycle_details` - The lifecycle state details of the compute bare metal host.
	* `state` - The current state of the compute bare metal host.
	* `time_created` - The date and time that the compute bare metal host was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
	* `time_updated` - The date and time that the compute bare metal host was updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 

