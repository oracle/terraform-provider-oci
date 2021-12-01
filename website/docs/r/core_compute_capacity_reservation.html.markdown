---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_compute_capacity_reservation"
sidebar_current: "docs-oci-resource-core-compute_capacity_reservation"
description: |-
  Provides the Compute Capacity Reservation resource in Oracle Cloud Infrastructure Core service
---

# oci_core_compute_capacity_reservation
This resource provides the Compute Capacity Reservation resource in Oracle Cloud Infrastructure Core service.

Creates a new compute capacity reservation in the specified compartment and availability domain.
Compute capacity reservations let you reserve instances in a compartment.
When you launch an instance using this reservation, you are assured that you have enough space for your instance, 
and you won't get out of capacity errors.
For more information, see [Reserved Capacity](https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/reserve-capacity.htm).


## Example Usage

```hcl
resource "oci_core_compute_capacity_reservation" "test_compute_capacity_reservation" {
	#Required
	availability_domain = var.compute_capacity_reservation_availability_domain
	compartment_id = var.compartment_id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.compute_capacity_reservation_display_name
	freeform_tags = {"Department"= "Finance"}
	instance_reservation_configs {
		#Required
		instance_shape = var.compute_capacity_reservation_instance_reservation_configs_instance_shape
		reserved_count = var.compute_capacity_reservation_instance_reservation_configs_reserved_count

		#Optional
		fault_domain = var.compute_capacity_reservation_instance_reservation_configs_fault_domain
		instance_shape_config {

			#Optional
			memory_in_gbs = var.compute_capacity_reservation_instance_reservation_configs_instance_shape_config_memory_in_gbs
			ocpus = var.compute_capacity_reservation_instance_reservation_configs_instance_shape_config_ocpus
		}
	}
	is_default_reservation = var.compute_capacity_reservation_is_default_reservation
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Required) The availability domain of this compute capacity reservation.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the capacity reservation. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `instance_reservation_configs` - (Optional) (Updatable) The capacity configurations for the capacity reservation.

	To use the reservation for the desired shape, specify the shape, count, and optionally the fault domain where you want this configuration. 
	* `fault_domain` - (Optional) (Updatable) The fault domain to use for instances created using this capacity configuration. For more information, see [Fault Domains](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/regions.htm#fault). If you do not specify the fault domain, the capacity is available for an instance that does not specify a fault domain. To change the fault domain for a reservation, delete the reservation and create a new one in the preferred fault domain.

		To retrieve a list of fault domains, use the `ListFaultDomains` operation in the [Identity and Access Management Service API](/iaas/api/#/en/identity/20160918/).

		Example: `FAULT-DOMAIN-1` 
	* `instance_shape` - (Required) (Updatable) The shape requested when launching instances using reserved capacity. The shape determines the number of CPUs, amount of memory, and other resources allocated to the instance. You can list all available shapes by calling [ListComputeCapacityReservationInstanceShapes](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/computeCapacityReservationInstanceShapes/ListComputeCapacityReservationInstanceShapes). 
	* `instance_shape_config` - (Optional) (Updatable) The shape configuration requested when launching instances in a compute capacity reservation.

		If the parameter is provided, the reservation is created with the resources that you specify. If some properties are missing or the parameter is not provided, the reservation is created with the default configuration values for the `shape` that you specify.

		Each shape only supports certain configurable values. If the values that you provide are not valid for the specified `shape`, an error is returned.

		For more information about customizing the resources that are allocated to flexible shapes, see [Flexible Shapes](https://docs.cloud.oracle.com/iaas/Content/Compute/References/computeshapes.htm#flexible). 
		* `memory_in_gbs` - (Optional) (Updatable) The total amount of memory available to the instance, in gigabytes. 
		* `ocpus` - (Optional) (Updatable) The total number of OCPUs available to the instance. 
	* `reserved_count` - (Required) (Updatable) The total number of instances that can be launched from the capacity configuration.
* `is_default_reservation` - (Optional) (Updatable) Whether this capacity reservation is the default. For more information, see [Capacity Reservations](https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/reserve-capacity.htm#default). 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `availability_domain` - The availability domain of the compute capacity reservation.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the compute capacity reservation. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compute capacity reservation.
* `instance_reservation_configs` - The capacity configurations for the capacity reservation.

	To use the reservation for the desired shape, specify the shape, count, and optionally the fault domain where you want this configuration. 
	* `fault_domain` - The fault domain of this capacity configuration. If a value is not supplied, this capacity configuration is applicable to all fault domains in the specified availability domain. For more information, see [Capacity Reservations](https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/reserve-capacity.htm). 
	* `instance_shape` - The shape to use when launching instances using compute capacity reservations. The shape determines the number of CPUs, the amount of memory, and other resources allocated to the instance. You can list all available shapes by calling [ListComputeCapacityReservationInstanceShapes](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/computeCapacityReservationInstanceShapes/ListComputeCapacityReservationInstanceShapes). 
	* `instance_shape_config` - The shape configuration requested when launching instances in a compute capacity reservation.

		If the parameter is provided, the reservation is created with the resources that you specify. If some properties are missing or the parameter is not provided, the reservation is created with the default configuration values for the `shape` that you specify.

		Each shape only supports certain configurable values. If the values that you provide are not valid for the specified `shape`, an error is returned.

		For more information about customizing the resources that are allocated to flexible shapes, see [Flexible Shapes](https://docs.cloud.oracle.com/iaas/Content/Compute/References/computeshapes.htm#flexible). 
		* `memory_in_gbs` - The total amount of memory available to the instance, in gigabytes. 
		* `ocpus` - The total number of OCPUs available to the instance. 
	* `reserved_count` - The total number of instances that can be launched from the capacity configuration.
	* `used_count` - The amount of capacity in use out of the total capacity reserved in this capacity configuration.
* `is_default_reservation` - Whether this capacity reservation is the default. For more information, see [Capacity Reservations](https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/reserve-capacity.htm#default). 
* `reserved_instance_count` - The number of instances for which capacity will be held with this compute capacity reservation. This number is the sum of the values of the `reservedCount` fields for all of the instance capacity configurations under this reservation. The purpose of this field is to calculate the percentage usage of the reservation. 
* `state` - The current state of the compute capacity reservation.
* `time_created` - The date and time the compute capacity reservation was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time the compute capacity reservation was updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `used_instance_count` - The total number of instances currently consuming space in this compute capacity reservation. This number is the sum of the values of the `usedCount` fields for all of the instance capacity configurations under this reservation. The purpose of this field is to calculate the percentage usage of the reservation. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Compute Capacity Reservation
	* `update` - (Defaults to 20 minutes), when updating the Compute Capacity Reservation
	* `delete` - (Defaults to 20 minutes), when destroying the Compute Capacity Reservation


## Import

ComputeCapacityReservations can be imported using the `id`, e.g.

```
$ terraform import oci_core_compute_capacity_reservation.test_compute_capacity_reservation "id"
```

