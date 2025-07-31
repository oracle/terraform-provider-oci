---
subcategory: "Apm Traces"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_apm_traces_log"
sidebar_current: "docs-oci-datasource-apm_traces-log"
description: |-
  Provides details about a specific Log in Oracle Cloud Infrastructure Apm Traces service
---

# Data Source: oci_apm_traces_log
This data source provides details about a specific Log resource in Oracle Cloud Infrastructure Apm Traces service.

Retrieve a log in the APM Domain.


## Example Usage

```hcl
data "oci_apm_traces_log" "test_log" {
	#Required
	apm_domain_id = oci_apm_apm_domain.test_apm_domain.id
	log_key = var.log_log_key
	time_log_ended_less_than = var.log_time_log_ended_less_than
	time_log_started_greater_than_or_equal_to = var.log_time_log_started_greater_than_or_equal_to
}
```

## Argument Reference

The following arguments are supported:

* `apm_domain_id` - (Required) The APM Domain ID for the intended request. 
* `log_key` - (Required) Log key. 
* `time_log_ended_less_than` - (Required) Include logs with log time less than this value. 
* `time_log_started_greater_than_or_equal_to` - (Required) Include logs with log time equal to or greater than this value. 


## Attributes Reference

The following attributes are exported:

* `attribute_metadata` - Metadata about the attributes in the logs. 
	* `attribute_type` - Type associated with the attribute key. 
	* `attribute_unit` - Unit associated with the attribute key.  If unit is not specified, it defaults to NONE. 
* `attributes` - List of attributes associated with the logs. 
	* `attribute_name` - Key that specifies the attribute name. 
	* `attribute_value` - Value associated with the attribute key. 
* `body` - Log body (Body). 
* `event_name` - Name of the event. 
* `log_key` - Unique identifier (logId) for the logKey.  Note that this field is defined as logKey in the API and it maps to the logId in Application Performance Monitoring. 
* `overflow_attributes` - Full values for attributes that are too long to be stored as a log attribute (Overflow). 
* `severity_number` - Log Severity number (SeverityNumber). 
* `severity_text` - Log Severity text (SeverityText).  Also known as Log level. 
* `span_key` - Unique identifier for the span (spanId) associated with this log. 
* `time_created` - Time that the log event occurred (CreatedTime). 
* `time_observed` - Time that the log was received by apm (ObservedTime). 
* `timestamp` - Time used by the time picker (RecordedTime).  Either the timeCreated if present or the timeObserved. 
* `trace_flags` - Trace flags. 
* `trace_key` - Unique identifier for the trace (traceId) associated with this log. 

