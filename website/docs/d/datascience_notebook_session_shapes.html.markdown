---
subcategory: "Data Science"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datascience_notebook_session_shapes"
sidebar_current: "docs-oci-datasource-datascience-notebook_session_shapes"
description: |-
  Provides the list of Notebook Session Shapes in Oracle Cloud Infrastructure Data Science service
---

# Data Source: oci_datascience_notebook_session_shapes
This data source provides the list of Notebook Session Shapes in Oracle Cloud Infrastructure Data Science service.

Lists the valid notebook session shapes.

## Example Usage

```hcl
data "oci_datascience_notebook_session_shapes" "test_notebook_session_shapes" {
	#Required
	compartment_id = var.compartment_id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) <b>Filter</b> results by the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.


## Attributes Reference

The following attributes are exported:

* `notebook_session_shapes` - The list of notebook_session_shapes.

### NotebookSessionShape Reference

The following attributes are exported:

* `core_count` - The number of cores associated with this notebook session shape. 
* `memory_in_gbs` - The amount of memory in GBs associated with this notebook session shape. 
* `name` - The name of the notebook session shape. 
* `shape_series` - The family that the compute shape belongs to. 

