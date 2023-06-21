---
subcategory: "Data Flow"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dataflow_run_statement"
sidebar_current: "docs-oci-resource-dataflow-run_statement"
description: |-
  Provides the Run Statement resource in Oracle Cloud Infrastructure Data Flow service
---

# oci_dataflow_run_statement
This resource provides the Run Statement resource in Oracle Cloud Infrastructure Data Flow service.

Executes a statement for a Session run.


## Example Usage

```hcl
resource "oci_dataflow_run_statement" "test_run_statement" {
	#Required
	code = var.run_statement_code
	run_id = oci_dataflow_run.test_run.id
}
```

## Argument Reference

The following arguments are supported:

* `code` - (Required) The statement code to execute. Example: `println(sc.version)` 
* `run_id` - (Required) The unique ID for the run 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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
* `time_created` - The date and time the resource was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2018-04-03T21:10:29.600Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Run Statement
	* `update` - (Defaults to 20 minutes), when updating the Run Statement
	* `delete` - (Defaults to 20 minutes), when destroying the Run Statement


## Import

RunStatements can be imported using the `id`, e.g.

```
$ terraform import oci_dataflow_run_statement.test_run_statement "runs/{runId}/statements/{statementId}" 
```

