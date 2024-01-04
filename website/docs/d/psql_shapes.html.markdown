---
subcategory: "Psql"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_psql_shapes"
sidebar_current: "docs-oci-datasource-psql-shapes"
description: |-
  Provides the list of Shapes in Oracle Cloud Infrastructure Psql service
---

# Data Source: oci_psql_shapes
This data source provides the list of Shapes in Oracle Cloud Infrastructure Psql service.

Returns the list of shapes allowed in the region.

## Example Usage

```hcl
data "oci_psql_shapes" "test_shapes" {

	#Optional
	compartment_id = var.compartment_id
	id = var.shape_id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The ID of the compartment in which to list resources.
* `id` - (Optional) A filter to return the feature by the shape name.


## Attributes Reference

The following attributes are exported:

* `shape_collection` - The list of shape_collection.

### Shape Reference

The following attributes are exported:

* `items` - List of supported shapes.
	* `id` - A unique identifier for the shape.
	* `memory_size_in_gbs` - The amount of memory in gigabytes.
	* `ocpu_count` - The number of OCPUs.
	* `shape` - The name of the Compute VM shape. Example: `VM.Standard.E4.Flex` 

