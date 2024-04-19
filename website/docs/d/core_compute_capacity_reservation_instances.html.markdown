---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_compute_capacity_reservation_instances"
sidebar_current: "docs-oci-datasource-core-compute_capacity_reservation_instances"
description: |-
  Provides the list of Compute Capacity Reservation Instances in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_compute_capacity_reservation_instances
This data source provides the list of Compute Capacity Reservation Instances in Oracle Cloud Infrastructure Core service.

Lists the instances launched under a capacity reservation. You can filter results by specifying criteria.


## Example Usage

```hcl
data "oci_core_compute_capacity_reservation_instances" "test_compute_capacity_reservation_instances" {
	#Required
	capacity_reservation_id = oci_core_capacity_reservation.test_capacity_reservation.id

	#Optional
	availability_domain = var.compute_capacity_reservation_instance_availability_domain
	compartment_id = var.compartment_id
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Optional) The name of the availability domain.  Example: `Uocm:PHX-AD-1` 
* `capacity_reservation_id` - (Required) The OCID of the compute capacity reservation.
* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.


## Attributes Reference

The following attributes are exported:

* `capacity_reservation_instances` - The list of capacity_reservation_instances.

### ComputeCapacityReservationInstance Reference

The following attributes are exported:

* `availability_domain` - The availability domain the instance is running in.
* `cluster_placement_group_id` - The OCID of the cluster placement group of the instance.
* `compartment_id` - The OCID of the compartment that contains the instance.
* `fault_domain` - The fault domain the instance is running in.
* `id` - The OCID of the instance.
* `shape` - The shape of the instance. The shape determines the number of CPUs, amount of memory, and other resources allocated to the instance.

	You can enumerate all available shapes by calling [ListComputeCapacityReservationInstanceShapes](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/computeCapacityReservationInstanceShapes/ListComputeCapacityReservationInstanceShapes). 
* `shape_config` - The shape configuration requested when launching instances in a compute capacity reservation.

	If the parameter is provided, the reservation is created with the resources that you specify. If some properties are missing or the parameter is not provided, the reservation is created with the default configuration values for the `shape` that you specify.

	Each shape only supports certain configurable values. If the values that you provide are not valid for the specified `shape`, an error is returned.

	For more information about customizing the resources that are allocated to flexible shapes, see [Flexible Shapes](https://docs.cloud.oracle.com/iaas/Content/Compute/References/computeshapes.htm#flexible). 
	* `memory_in_gbs` - The total amount of memory available to the instance, in gigabytes. 
	* `ocpus` - The total number of OCPUs available to the instance. 

