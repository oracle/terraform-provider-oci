---
subcategory: "Apm Traces"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_apm_traces_trace_aggregated_snapshot_data"
sidebar_current: "docs-oci-datasource-apm_traces-trace_aggregated_snapshot_data"
description: |-
  Provides details about a specific Trace Aggregated Snapshot Data in Oracle Cloud Infrastructure Apm Traces service
---

# Data Source: oci_apm_traces_trace_aggregated_snapshot_data
This data source provides details about a specific Trace Aggregated Snapshot Data resource in Oracle Cloud Infrastructure Apm Traces service.

Gets the aggregated snapshot identified by trace ID.


## Example Usage

```hcl
data "oci_apm_traces_trace_aggregated_snapshot_data" "test_trace_aggregated_snapshot_data" {
	#Required
	apm_domain_id = oci_apm_apm_domain.test_apm_domain.id
	trace_key = var.trace_aggregated_snapshot_data_trace_key

	#Optional
	server_name = var.trace_aggregated_snapshot_data_server_name
	service_name = oci_core_service.test_service.name
	span_key = var.trace_aggregated_snapshot_data_span_key
	span_name = var.trace_aggregated_snapshot_data_span_name
}
```

## Argument Reference

The following arguments are supported:

* `apm_domain_id` - (Required) The APM Domain ID for the intended request. 
* `server_name` - (Optional) Name of the server. 
* `service_name` - (Optional) Name associated with the service. 
* `span_key` - (Optional) Unique Application Performance Monitoring span identifier (spanId). 
* `span_name` - (Optional) Name of the span associated with the trace. 
* `trace_key` - (Required) Unique Application Performance Monitoring trace identifier (traceId). 


## Attributes Reference

The following attributes are exported:

* `details` - Aggregated snapshot details. 
	* `key` - Name of the property. 
	* `value` - Value of the property. 

