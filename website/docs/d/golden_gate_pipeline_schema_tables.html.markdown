---
subcategory: "Golden Gate"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_golden_gate_pipeline_schema_tables"
sidebar_current: "docs-oci-datasource-golden_gate-pipeline_schema_tables"
description: |-
  Provides the list of Pipeline Schema Tables in Oracle Cloud Infrastructure Golden Gate service
---

# Data Source: oci_golden_gate_pipeline_schema_tables
This data source provides the list of Pipeline Schema Tables in Oracle Cloud Infrastructure Golden Gate service.

Returns an array of tables under the given schemas of the pipeline for given source and target schemas passed as query params.


## Example Usage

```hcl
data "oci_golden_gate_pipeline_schema_tables" "test_pipeline_schema_tables" {
	#Required
	pipeline_id = oci_golden_gate_pipeline.test_pipeline.id
	source_schema_name = var.pipeline_schema_table_source_schema_name
	target_schema_name = var.pipeline_schema_table_target_schema_name

	#Optional
	display_name = var.pipeline_schema_table_display_name
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Optional) A filter to return only the resources that match the entire 'displayName' given. 
* `pipeline_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the pipeline created. 
* `source_schema_name` - (Required) Name of the source schema obtained from get schema endpoint of the created pipeline. 
* `target_schema_name` - (Required) Name of the target schema obtained from get schema endpoint of the created pipeline. 


## Attributes Reference

The following attributes are exported:

* `pipeline_schema_table_collection` - The list of pipeline_schema_table_collection.

### PipelineSchemaTable Reference

The following attributes are exported:

* `items` - Array of source or target schema tables of a pipeline's assigned connection. 
	* `source_table_name` - The table name from the schema of database connection. 
	* `target_table_name` - The table name from the schema of database connection. 
* `source_schema_name` - The schema name from the database connection. 
* `target_schema_name` - The schema name from the database connection. 

