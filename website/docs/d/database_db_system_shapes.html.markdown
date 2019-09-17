---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_db_system_shapes"
sidebar_current: "docs-oci-datasource-database-db_system_shapes"
description: |-
  Provides the list of Db System Shapes in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_db_system_shapes
This data source provides the list of Db System Shapes in Oracle Cloud Infrastructure Database service.

Gets a list of the shapes that can be used to launch a new DB system. The shape determines resources to allocate to the DB system - CPU cores and memory for VM shapes; CPU cores, memory and storage for non-VM (or bare metal) shapes.

## Example Usage

```hcl
data "oci_database_db_system_shapes" "test_db_system_shapes" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	availability_domain = "${var.db_system_shape_availability_domain}"
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Optional) The name of the Availability Domain.
* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `db_system_shapes` - The list of db_system_shapes.

### DbSystemShape Reference

The following attributes are exported:

* `available_core_count` - The maximum number of CPU cores that can be enabled on the DB system for this shape.
* `core_count_increment` - The discrete number by which the CPU core count for this shape can be increased or decreased.
* `maximum_node_count` - The maximum number of database nodes available for this shape.
* `minimum_core_count` - The minimum number of CPU cores that can be enabled on the DB system for this shape.
* `minimum_node_count` - The minimum number of database nodes available for this shape.
* `name` - The name of the shape used for the DB system.
* `shape` - Deprecated. Use `name` instead of `shape`.
* `shape_family` - The family of the shape used for the DB system.

