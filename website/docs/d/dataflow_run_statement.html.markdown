---
subcategory: "Data Flow"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dataflow_run_statement"
sidebar_current: "docs-oci-datasource-dataflow-run_statement"
description: |-
  Provides details about a specific Run Statement in Oracle Cloud Infrastructure Data Flow service
---

# Data Source: oci_dataflow_run_statement
This data source provides details about a specific Run Statement resource in Oracle Cloud Infrastructure Data Flow service.

Retrieves the statement corresponding to the `statementId` for a Session run specified by `runId`.


## Example Usage

```hcl
data "oci_dataflow_run_statement" "test_run_statement" {
	#Required
	run_id = oci_dataflow_run.test_run.id
	statement_id = oci_dataflow_statement.test_statement.id
}
```

## Argument Reference

The following arguments are supported:

* `run_id` - (Required) The unique ID for the run 
* `statement_id` - (Required) The unique ID for the statement. 


## Attributes Reference

The following attributes are exported:

* `code` - The statement code to execute. Example: `println(sc.version)` 
* `id` - The statement ID. 
* `output` - The execution output of a statement. 
	* `data` - An object representing execution output of a statement. 
		* `type` - The type of the `StatementOutputData` like `TEXT_PLAIN`, `TEXT_HTML` or `IMAGE_PNG`. 
		* `value` - The statement code execution output in html format. 
	* `error_name` - The name of the error in the statement output. 
	* `error_value` - The value of the error in the statement output. 
	* `status` - Status of the statement output. 
	* `traceback` - The traceback of the statement output. 
* `progress` - The execution progress. 
* `run_id` - The ID of a run. 
* `state` - The current state of this statement. 
* `time_completed` - The date and time a statement execution was completed, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2022-05-31T21:10:29.600Z` 
* `time_created` - The date and time a application was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2018-04-03T21:10:29.600Z` 

