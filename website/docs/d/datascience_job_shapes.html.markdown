---
subcategory: "Data Science"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datascience_job_shapes"
sidebar_current: "docs-oci-datasource-datascience-job_shapes"
description: |-
  Provides the list of Job Shapes in Oracle Cloud Infrastructure Data Science service
---

# Data Source: oci_datascience_job_shapes
This data source provides the list of Job Shapes in Oracle Cloud Infrastructure Data Science service.

List job shapes available in the specified compartment.

## Example Usage

```hcl
data "oci_datascience_job_shapes" "test_job_shapes" {
	#Required
	compartment_id = var.compartment_id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) <b>Filter</b> results by the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.


## Attributes Reference

The following attributes are exported:

* `job_shapes` - The list of job_shapes.

### JobShape Reference

The following attributes are exported:

* `core_count` - The number of cores associated with this job run shape. 
* `memory_in_gbs` - The number of cores associated with this job shape. 
* `name` - The name of the job shape. 
* `shape_series` - The family that the compute shape belongs to. 

