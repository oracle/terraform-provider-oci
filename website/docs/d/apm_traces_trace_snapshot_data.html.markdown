---
subcategory: "Apm Traces"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_apm_traces_trace_snapshot_data"
sidebar_current: "docs-oci-datasource-apm_traces-trace_snapshot_data"
description: |-
  Provides details about a specific Trace Snapshot Data in Oracle Cloud Infrastructure Apm Traces service
---

# Data Source: oci_apm_traces_trace_snapshot_data
This data source provides details about a specific Trace Snapshot Data resource in Oracle Cloud Infrastructure Apm Traces service.

Gets the trace snapshots data identified by trace ID.


## Example Usage

```hcl
data "oci_apm_traces_trace_snapshot_data" "test_trace_snapshot_data" {
	#Required
	apm_domain_id = oci_apm_apm_domain.test_apm_domain.id
	trace_key = var.trace_snapshot_data_trace_key

	#Optional
	is_summarized = var.trace_snapshot_data_is_summarized
	snapshot_time = var.trace_snapshot_data_snapshot_time
	thread_id = oci_apm_traces_thread.test_thread.id
}
```

## Argument Reference

The following arguments are supported:

* `apm_domain_id` - (Required) The APM Domain ID for the intended request. 
* `is_summarized` - (Optional) If enabled, only span level details are sent. 
* `snapshot_time` - (Optional) Epoch time of snapshot. 
* `thread_id` - (Optional) Thread ID for which snapshots need to be retrieved. This identifier of a thread is a long positive number generated when a thread is created. 
* `trace_key` - (Required) Unique Application Performance Monitoring trace identifier (traceId). 


## Attributes Reference

The following attributes are exported:

* `key` - Unique identifier (traceId) for the trace that represents the span set.  Note that this field is defined as traceKey in the API and it maps to the traceId in the trace data in Application Performance Monitoring. 
* `time_ended` - End time of the trace. 
* `time_started` - Start time of the trace. 
* `trace_snapshot_details` - Trace snapshots properties. 
	* `key` - Name of the property. 
	* `value` - Value of the property. 

