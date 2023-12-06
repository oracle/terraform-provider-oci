---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_compute_capacity_topologies"
sidebar_current: "docs-oci-datasource-core-compute_capacity_topologies"
description: |-
  Provides the list of Compute Capacity Topologies in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_compute_capacity_topologies
This data source provides the list of Compute Capacity Topologies in Oracle Cloud Infrastructure Core service.

Lists the compute capacity topologies in the specified compartment. You can filter the list by a compute
capacity topology display name.


## Example Usage

```hcl
data "oci_core_compute_capacity_topologies" "test_compute_capacity_topologies" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	availability_domain = var.compute_capacity_topology_availability_domain
	display_name = var.compute_capacity_topology_display_name
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Optional) The name of the availability domain.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. 


## Attributes Reference

The following attributes are exported:

* `compute_capacity_topology_collection` - The list of compute_capacity_topology_collection.

### ComputeCapacityTopology Reference

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

