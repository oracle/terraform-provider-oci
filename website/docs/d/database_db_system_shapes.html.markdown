---
subcategory: "Database"
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
	compartment_id = var.compartment_id

	#Optional
	availability_domain = var.db_system_shape_availability_domain
	shape_attribute = var.db_system_shape_shape_attribute
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Optional) The name of the Availability Domain.
* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `shape_attribute` - (Optional) If provided and applicable, return DB System shape parameters based on the shapeAttribute provided


## Attributes Reference

The following attributes are exported:

* `db_system_shapes` - The list of db_system_shapes.

### DbSystemShape Reference

The following attributes are exported:

* `are_server_types_supported` - If true, the shape supports configurable DB and Storage Server types.
* `available_core_count` - The maximum number of CPU cores that can be enabled on the DB system for this shape.
* `available_core_count_per_node` - The maximum number of CPU cores per database node that can be enabled for this shape. Only applicable to the flex Exadata shape, ExaCC Elastic shapes and VM Flex shapes.
* `available_data_storage_in_tbs` - The maximum DATA storage that can be enabled for this shape.
* `available_data_storage_per_server_in_tbs` - The maximum data storage available per storage server for this shape. Only applicable to ExaCC Elastic shapes.
* `available_db_node_per_node_in_gbs` - The maximum Db Node storage available per database node for this shape. Only applicable to ExaCC Elastic shapes.
* `available_db_node_storage_in_gbs` - The maximum Db Node storage that can be enabled for this shape.
* `available_memory_in_gbs` - The maximum memory that can be enabled for this shape.
* `available_memory_per_node_in_gbs` - The maximum memory available per database node for this shape. Only applicable to ExaCC Elastic shapes.
* `compute_model` - The compute model of the Autonomous Database. This is required if using the `computeCount` parameter. If using `cpuCoreCount` then it is an error to specify `computeModel` to a non-null value. ECPU compute model is the recommended model and OCPU compute model is legacy.
* `core_count_increment` - The discrete number by which the CPU core count for this shape can be increased or decreased.
* `display_name` - The display name of the shape used for the DB system.
* `max_storage_count` - The maximum number of Exadata storage servers available for the Exadata infrastructure.
* `maximum_node_count` - The maximum number of compute servers available for this shape.
* `min_core_count_per_node` - The minimum number of CPU cores that can be enabled per node for this shape.
* `min_data_storage_in_tbs` - The minimum data storage that need be allocated for this shape.
* `min_db_node_storage_per_node_in_gbs` - The minimum Db Node storage that need be allocated per node for this shape.
* `min_memory_per_node_in_gbs` - The minimum memory that need be allocated per node for this shape.
* `min_storage_count` - The minimum number of Exadata storage servers available for the Exadata infrastructure.
* `minimum_core_count` - The minimum number of CPU cores that can be enabled on the DB system for this shape.
* `minimum_node_count` - The minimum number of compute servers available for this shape.
* `name` - The name of the shape used for the DB system.
* `shape` - Deprecated. Use `name` instead of `shape`.
* `shape_attributes` - The shapeAttributes of the DB system shape. 
* `shape_family` - The family of the shape used for the DB system.
* `shape_type` - The shape type for the virtual machine DB system. Shape type is determined by CPU hardware. Valid values are `AMD` , `INTEL` or `INTEL_FLEX_X9`.

