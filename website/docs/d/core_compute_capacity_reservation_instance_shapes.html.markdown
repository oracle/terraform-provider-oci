---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_compute_capacity_reservation_instance_shapes"
sidebar_current: "docs-oci-datasource-core-compute_capacity_reservation_instance_shapes"
description: |-
  Provides the list of Compute Capacity Reservation Instance Shapes in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_compute_capacity_reservation_instance_shapes
This data source provides the list of Compute Capacity Reservation Instance Shapes in Oracle Cloud Infrastructure Core service.

Lists the shapes that can be reserved within the specified compartment.


## Example Usage

```hcl
data "oci_core_compute_capacity_reservation_instance_shapes" "test_compute_capacity_reservation_instance_shapes" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	availability_domain = var.compute_capacity_reservation_instance_shape_availability_domain
	display_name = var.compute_capacity_reservation_instance_shape_display_name
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Optional) The name of the availability domain.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. 


## Attributes Reference

The following attributes are exported:

* `compute_capacity_reservation_instance_shapes` - The list of compute_capacity_reservation_instance_shapes.

### ComputeCapacityReservationInstanceShape Reference

The following attributes are exported:

* `availability_domain` - The shape's availability domain. 
* `instance_shape` - The name of the available shape used to launch instances in a compute capacity reservation. 

