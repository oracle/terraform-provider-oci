---
subcategory: "Golden Gate"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_golden_gate_pipeline_schemas"
sidebar_current: "docs-oci-datasource-golden_gate-pipeline_schemas"
description: |-
  Provides the list of Pipeline Schemas in Oracle Cloud Infrastructure Golden Gate service
---

# Data Source: oci_golden_gate_pipeline_schemas
This data source provides the list of Pipeline Schemas in Oracle Cloud Infrastructure Golden Gate service.

Returns an array of schemas based on mapping rules for a pipeline.


## Example Usage

```hcl
data "oci_golden_gate_pipeline_schemas" "test_pipeline_schemas" {
	#Required
	pipeline_id = oci_golden_gate_pipeline.test_pipeline.id

	#Optional
	display_name = var.pipeline_schema_display_name
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Optional) A filter to return only the resources that match the entire 'displayName' given. 
* `pipeline_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the pipeline created. 


## Attributes Reference

The following attributes are exported:

* `pipeline_schema_collection` - The list of pipeline_schema_collection.

### PipelineSchema Reference

The following attributes are exported:

* `items` - Array of pipeline schemas 
	* `source_schema_name` - The schema name from the database connection. 
	* `target_schema_name` - The schema name from the database connection. 

