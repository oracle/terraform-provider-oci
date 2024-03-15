---
subcategory: "Apm Traces"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_apm_traces_trace"
sidebar_current: "docs-oci-datasource-apm_traces-trace"
description: |-
  Provides details about a specific Trace in Oracle Cloud Infrastructure Apm Traces service
---

# Data Source: oci_apm_traces_trace
This data source provides details about a specific Trace resource in Oracle Cloud Infrastructure Apm Traces service.

Gets the trace details identified by traceId.


## Example Usage

```hcl
data "oci_apm_traces_trace" "test_trace" {
	#Required
	apm_domain_id = oci_apm_apm_domain.test_apm_domain.id
	trace_key = var.trace_trace_key

	#Optional
	time_trace_started_greater_than_or_equal_to = var.trace_time_trace_started_greater_than_or_equal_to
	time_trace_started_less_than = var.trace_time_trace_started_less_than
	trace_namespace = var.trace_trace_namespace
}
```

## Argument Reference

The following arguments are supported:

* `apm_domain_id` - (Required) The APM Domain ID for the intended request. 
* `time_trace_started_greater_than_or_equal_to` - (Optional) Include traces that have a `minTraceStartTime` equal to or greater than this value. 
* `time_trace_started_less_than` - (Optional) Include traces that have a `minTraceStartTime` less than this value. 
* `trace_key` - (Required) Unique Application Performance Monitoring trace identifier (traceId). 
* `trace_namespace` - (Optional) Name space from which the trace details need to be retrieved. 


## Attributes Reference

The following attributes are exported:

* `error_span_count` - The number of spans with errors that have been processed by the system for the trace. Note that the number of spans with errors will be less than or equal to the total number of spans in the trace.  
* `is_fault` - Boolean flag that indicates whether the trace has an error. 
* `key` - Unique identifier (traceId) for the trace that represents the span set.  Note that this field is defined as traceKey in the API and it maps to the traceId in the trace data in Application Performance Monitoring. 
* `root_span_duration_in_ms` - Time taken for the root span operation to complete in milliseconds. 
* `root_span_operation_name` - Root span name associated with the trace. This is the flow start operation name. Null is displayed if the root span is not yet completed. 
* `root_span_service_name` - Service associated with the trace. 
* `service_summaries` - A summary of the spans by service. 
	* `error_spans` - Number of spans with errors for serviceName in the trace. 
	* `span_service_name` - Name associated with the service. 
	* `total_spans` - Number of spans for serviceName in the trace. 
* `source_name` - Source of trace (traces, syn_traces). 
* `span_count` - The number of spans that have been processed by the system for the trace.  Note that there could be additional spans that have not been processed or reported yet if the trace is still in progress. 
* `span_summary` - Summary of the information pertaining to the spans in the trace window that is being queried. 
	* `error_span_count` - The number of spans with errors that have been processed by the system for the trace. Note that the number of spans with errors will be less than or equal to the total number of spans in the trace.  
	* `is_fault` - Boolean flag that indicates whether the trace has an error. 
	* `key` - Unique identifier (traceId) for the trace that represents the span set.  Note that this field is defined as traceKey in the API and it maps to the traceId in the trace data in Application Performance Monitoring. 
	* `root_span_duration_in_ms` - Time taken for the root span operation to complete in milliseconds. 
	* `root_span_operation_name` - Root span name associated with the trace. This is the flow start operation name. Null is displayed if the root span is not yet completed. 
	* `root_span_service_name` - Service associated with the trace. 
	* `service_summaries` - A summary of the spans by service. 
		* `error_spans` - Number of spans with errors for serviceName in the trace. 
		* `span_service_name` - Name associated with the service. 
		* `total_spans` - Number of spans for serviceName in the trace. 
	* `span_count` - The number of spans that have been processed by the system for the trace.  Note that there could be additional spans that have not been processed or reported yet if the trace is still in progress. 
	* `time_earliest_span_started` - Start time of the earliest span in the span collection. 
	* `time_latest_span_ended` - End time of the span that most recently ended in the span collection. 
	* `time_root_span_ended` - End time of the root span for the span collection. 
	* `time_root_span_started` - Start time of the root span for the span collection. 
	* `trace_duration_in_ms` - Time between the start of the earliest span and the end of the most recent span in milliseconds. 
	* `trace_error_code` - Error code of the trace. 
	* `trace_error_type` - Error type of the trace. 
	* `trace_status` - The status of the trace. The trace statuses are defined as follows: complete - a root span has been recorded, but there is no information on the errors. success - a complete root span is recorded there is a successful error type and error code - HTTP 200. incomplete - the root span has not yet been received. error - the root span returned with an error. There may or may not be an associated error code or error type. 
* `spans` - An array of spans in the trace. 
	* `duration_in_ms` - Total span duration in milliseconds. 
	* `is_error` - Indicates if the span has an error. 
	* `key` - Unique identifier (spanId) for the span.  Note that this field is defined as spanKey in the API and it maps to the spanId in the trace data in Application Performance Monitoring. 
	* `kind` - Kind associated with the span. 
	* `logs` - List of logs associated with the span. 
		* `event_name` - Name of the event for which the log is created. 
		* `span_logs` - List of logs associated with the span at the given timestamp. 
			* `log_key` - Key that specifies the log name. 
			* `log_value` - Value associated with the log key. 
		* `time_created` - Timestamp at which the log is created. 
	* `operation_name` - Span name associated with the trace.  This is usually the method or URI of the request. 
	* `parent_span_key` - Unique parent identifier for the span if one exists. For root spans this will be null. 
	* `service_name` - Service name associated with the span. 
	* `source_name` - Source of span (spans, syn_spans). 
	* `tags` - List of tags associated with the span. 
		* `tag_name` - Key that specifies the tag name. 
		* `tag_value` - Value associated with the tag key. 
	* `tags_metadata` - Metadata about the tags in the span. 
		* `tag_type` - Type associated with the tag key. 
		* `tag_unit` - Unit associated with the tag key. 
	* `time_ended` - Span end time.  Timestamp when the span was completed. 
	* `time_started` - Span start time.  Timestamp when the span was started. 
	* `trace_key` - Unique identifier for the trace. 
* `time_earliest_span_started` - Start time of the earliest span in the span collection. 
* `time_latest_span_ended` - End time of the span that most recently ended in the span collection. 
* `time_root_span_ended` - End time of the root span for the span collection. 
* `time_root_span_started` - Start time of the root span for the span collection. 
* `trace_duration_in_ms` - Time between the start of the earliest span and the end of the most recent span in milliseconds. 
* `trace_error_code` - Error code of the trace. 
* `trace_error_type` - Error type of the trace. 
* `trace_status` - The status of the trace. The trace statuses are defined as follows: complete - a root span has been recorded, but there is no information on the errors. success - a complete root span is recorded there is a successful error type and error code - HTTP 200. incomplete - the root span has not yet been received. error - the root span returned with an error. There may or may not be an associated error code or error type. 

