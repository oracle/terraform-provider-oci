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

* `items` - List of dbSystems.
	* `id` - Unique identifier for the shape
	* `memory_size_in_gbs` - The amount of memory in GB
	* `ocpu_count` - The number of OCPUs
	* `shape` - The Compute Shape Name like VM.Standard.E4.Flex

