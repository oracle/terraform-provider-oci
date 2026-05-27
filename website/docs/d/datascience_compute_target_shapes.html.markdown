---
subcategory: "Data Science"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datascience_compute_target_shapes"
sidebar_current: "docs-oci-datasource-datascience-compute_target_shapes"
description: |-
  Provides the list of Compute Target Shapes in Oracle Cloud Infrastructure Data Science service
---

# Data Source: oci_datascience_compute_target_shapes
This data source provides the list of Compute Target Shapes in Oracle Cloud Infrastructure Data Science service.

Lists the valid compute target shapes.

## Example Usage

```hcl
data "oci_datascience_compute_target_shapes" "test_compute_target_shapes" {
	#Required
	compartment_id = var.compartment_id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) <b>Filter</b> results by the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.


## Attributes Reference

The following attributes are exported:

* `compute_target_shapes` - The list of compute_target_shapes.

### ComputeTargetShape Reference

The following attributes are exported:

* `core_count` - The number of cores associated with this compute target shape. 
* `memory_in_gbs` - The amount of memory in GBs associated with this compute target shape. 
* `name` - The name of the compute target shape. 
* `shape_series` - The family that the compute shape belongs to. 

