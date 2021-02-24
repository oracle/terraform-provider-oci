---
subcategory: "MySQL Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_mysql_shapes"
sidebar_current: "docs-oci-datasource-mysql-shapes"
description: |-
  Provides the list of Shapes in Oracle Cloud Infrastructure MySQL Database service
---

# Data Source: oci_mysql_shapes
This data source provides the list of Shapes in Oracle Cloud Infrastructure MySQL Database service.

Gets a list of the shapes you can use to create a new MySQL DB System.
The shape determines the resources allocated to the DB System:
CPU cores and memory for VM shapes; CPU cores, memory and
storage for non-VM (or bare metal) shapes.


## Example Usage

```hcl
data "oci_mysql_shapes" "test_shapes" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	availability_domain = var.shape_availability_domain
	is_supported_for = var.shape_is_supported_for
	name = var.shape_name
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Optional) The name of the Availability Domain.
* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `is_supported_for` - (Optional) Return shapes that are supported by the service feature. 
* `name` - (Optional) Name


## Attributes Reference

The following attributes are exported:

* `shapes` - The list of shapes.

### Shape Reference

The following attributes are exported:

* `cpu_core_count` - The number of CPU Cores the Instance provides. These are "OCPU"s.
* `is_supported_for` - What service features the shape is supported for. 
* `memory_size_in_gbs` - The amount of RAM the Instance provides. This is an IEC base-2 number.
* `name` - The name of the shape used for the DB System.

