---
subcategory: "Golden Gate"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_golden_gate_pipeline_running_processes"
sidebar_current: "docs-oci-datasource-golden_gate-pipeline_running_processes"
description: |-
  Provides the list of Pipeline Running Processes in Oracle Cloud Infrastructure Golden Gate service
---

# Data Source: oci_golden_gate_pipeline_running_processes
This data source provides the list of Pipeline Running Processes in Oracle Cloud Infrastructure Golden Gate service.

Retrieves a Pipeline's running replication process's status like Capture/Apply.


## Example Usage

```hcl
data "oci_golden_gate_pipeline_running_processes" "test_pipeline_running_processes" {
	#Required
	pipeline_id = oci_golden_gate_pipeline.test_pipeline.id
}
```

## Argument Reference

The following arguments are supported:

* `pipeline_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the pipeline created. 


## Attributes Reference

The following attributes are exported:

* `pipeline_running_process_collection` - The list of pipeline_running_process_collection.

### PipelineRunningProcess Reference

The following attributes are exported:

* `items` - The list of replication processes and their details. 
	* `last_record_lag_in_seconds` - The latency, in seconds, of a process running in a replication. This option applies when retrieving running processes. 
	* `name` - An object's Display Name. 
	* `process_type` - The type of process running in a replication. For example, Extract or Replicat. This option applies when retrieving running processes. 
	* `status` - The status of the Extract or Replicat process. This option applies when retrieving running processes. 
	* `time_last_processed` - The date and time the last record was processed by an Extract or Replicat. This option applies when retrieving running processes. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2024-07-25T21:10:29.600Z`. 

