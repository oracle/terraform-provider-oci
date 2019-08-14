---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_audit_events"
sidebar_current: "docs-oci-datasource-audit-events"
description: |-
  Provides the list of Audit Events in Oracle Cloud Infrastructure Audit service
---

# Data Source: oci_audit_events
This data source provides the list of Audit Events in Oracle Cloud Infrastructure Audit service.

Returns all audit events for the specified compartment that were processed within the specified time range.

## Example Usage

```hcl
data "oci_audit_events" "test_audit_events" {
	#Required
	compartment_id = "${var.compartment_id}"
	end_time = "${var.audit_event_end_time}"
	start_time = "${var.audit_event_start_time}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `end_time` - (Required) Returns events that were processed before this end date and time, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. For example, a start value of `2017-01-01T00:00:00Z` and an end value of `2017-01-02T00:00:00Z` will retrieve a list of all events processed on January 1, 2017. Similarly, a start value of `2017-01-01T00:00:00Z` and an end value of `2017-02-01T00:00:00Z` will result in a list of all events processed between January 1, 2017 and January 31, 2017. You can specify a value with granularity to the minute. Seconds (and milliseconds, if included) must be set to `0`. 
* `start_time` - (Required) Returns events that were processed at or after this start date and time, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. For example, a start value of `2017-01-15T11:30:00Z` will retrieve a list of all events processed since 30 minutes after the 11th hour of January 15, 2017, in Coordinated Universal Time (UTC). You can specify a value with granularity to the minute. Seconds (and milliseconds, if included) must be set to `0`.
* `limit` - (Optional) The number of pages of events to request from the service. Default to 1. Large `start_time` and `end_time` ranges or very active tenancies may result in very large data sets that could cause performance issues running Terraform commands. This default value mitigates that risk by requiring intentionally setting a higher tolerance for slow running Terarform commands with potentially large statefiles. 


## Attributes Reference

The following attributes are exported:

* `audit_events` - The list of audit_events.

### AuditEvent Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment.
* `compartment_name` - The name of the compartment. This value is the friendly name associated with compartmentId. This value can change, but the service logs the value that appeared at the time of the audit event. 
* `credential_id` - The credential ID of the user. This value is extracted from the HTTP 'Authorization' request header. It consists of the tenantId, userId, and user fingerprint, all delimited by a slash (/).
* `event_id` - The GUID of the event.
* `event_name` - The name of the event. Example: `LaunchInstance` 
* `event_source` - The source of the event.
* `event_time` - The time the event occurred, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format.
* `event_type` - The type of the event.
* `principal_id` - The OCID of the user whose action triggered the event.
* `request_action` - The HTTP method of the request.
* `request_agent` - The user agent of the client that made the request.
* `request_headers` - The HTTP header fields and values in the request.
* `request_id` - The opc-request-id of the request.
* `request_origin` - The IP address of the source of the request.
* `request_parameters` - The query parameter fields and values for the request.
* `request_resource` - The resource targeted by the request.
* `response_headers` - The headers of the response.
* `response_payload` - Metadata of interest from the response payload. For example, the OCID of a resource.
* `response_status` - The status code of the response.
* `response_time` - The time of the response to the audited request, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format.
* `tenant_id` - The OCID of the tenant.
* `user_name` - The name of the user or service. This value is the friendly name associated with principalId.

