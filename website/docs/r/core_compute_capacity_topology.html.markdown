---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_compute_capacity_topology"
sidebar_current: "docs-oci-resource-core-compute_capacity_topology"
description: |-
  Provides the Compute Capacity Topology resource in Oracle Cloud Infrastructure Core service
---

# oci_core_compute_capacity_topology
This resource provides the Compute Capacity Topology resource in Oracle Cloud Infrastructure Core service.

Creates a new compute capacity topology in the specified compartment and availability domain.

Compute capacity topologies provide the RDMA network topology of your bare metal hosts so that you can launch
instances on your bare metal hosts with targeted network locations.

Compute capacity topologies report the health status of your bare metal hosts.


## Example Usage

```hcl
resource "oci_core_compute_capacity_topology" "test_compute_capacity_topology" {
	#Required
	availability_domain = var.compute_capacity_topology_availability_domain
	capacity_source {
		#Required
		capacity_type = var.compute_capacity_topology_capacity_source_capacity_type

		#Optional
		compartment_id = var.compartment_id
	}
	compartment_id = var.compartment_id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.compute_capacity_topology_display_name
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Required) The availability domain of this compute capacity topology.  Example: `Uocm:US-CHICAGO-1-AD-2` 
* `capacity_source` - (Required) (Updatable) A capacity source of bare metal hosts. 
	* `capacity_type` - (Required) (Updatable) The capacity type of bare metal hosts.
	* `compartment_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment of this capacity source. 
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains this compute capacity topology. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `availability_domain` - The availability domain of the compute capacity topology.  Example: `Uocm:US-CHICAGO-1-AD-2` 
* `capacity_source` - A capacity source of bare metal hosts. 
	* `capacity_type` - The capacity type of bare metal hosts.
	* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment of this capacity source. 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the compute capacity topology. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compute capacity topology.
* `state` - The current state of the compute capacity topology.
* `time_created` - The date and time that the compute capacity topology was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time that the compute capacity topology was updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Compute Capacity Topology
	* `update` - (Defaults to 20 minutes), when updating the Compute Capacity Topology
	* `delete` - (Defaults to 20 minutes), when destroying the Compute Capacity Topology


## Import

ComputeCapacityTopologies can be imported using the `id`, e.g.

```
$ terraform import oci_core_compute_capacity_topology.test_compute_capacity_topology "id"
```

