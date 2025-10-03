---
subcategory: "Managed Kafka"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_managed_kafka_node_shapes"
sidebar_current: "docs-oci-datasource-managed_kafka-node_shapes"
description: |-
  Provides the list of Node Shapes in Oracle Cloud Infrastructure Managed Kafka service
---

# Data Source: oci_managed_kafka_node_shapes
This data source provides the list of Node Shapes in Oracle Cloud Infrastructure Managed Kafka service.

Returns the list of shapes allowed in the region.

## Example Usage

```hcl
data "oci_managed_kafka_node_shapes" "test_node_shapes" {

	#Optional
	compartment_id = var.compartment_id
	name = var.node_shape_name
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
* `name` - (Optional) The name to filter on.


## Attributes Reference

The following attributes are exported:

* `node_shape_collection` - The list of node_shape_collection.

### NodeShape Reference

The following attributes are exported:

* `items` - List of NodeShapeSummary.
	* `name` - The name of the shape

