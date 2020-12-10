---
subcategory: "Audit"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_audit_events"
sidebar_current: "docs-oci-datasource-audit-events"
description: |-
  Provides the list of Audit Events in Oracle Cloud Infrastructure Audit service
---

# Data Source: oci_audit_events
This data source provides the list of Audit Events in Oracle Cloud Infrastructure Audit service.

Returns all the audit events processed for the specified compartment within the specified
time range.


## Example Usage

```hcl
data "oci_audit_events" "test_audit_events" {
	#Required
	compartment_id = var.compartment_id
	end_time = var.audit_event_end_time
	start_time = var.audit_event_start_time
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `end_time` - (Required) Returns events that were processed before this end date and time, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format.

	For example, a start value of `2017-01-01T00:00:00Z` and an end value of `2017-01-02T00:00:00Z` will retrieve a list of all events processed on January 1, 2017. Similarly, a start value of `2017-01-01T00:00:00Z` and an end value of `2017-02-01T00:00:00Z` will result in a list of all events processed between January 1, 2017 and January 31, 2017. You can specify a value with granularity to the minute. Seconds (and milliseconds, if included) must be set to `0`. 
* `start_time` - (Required) Returns events that were processed at or after this start date and time, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format.

	For example, a start value of `2017-01-15T11:30:00Z` will retrieve a list of all events processed since 30 minutes after the 11th hour of January 15, 2017, in Coordinated Universal Time (UTC). You can specify a value with granularity to the minute. Seconds (and milliseconds, if included) must be set to `0`. 


## Attributes Reference

The following attributes are exported:

* `audit_events` - The list of audit_events.

### AuditEvent Reference

The following attributes are exported:

* `cloud_events_version` - The version of the CloudEvents specification. The structure of the envelope follows the  [CloudEvents](https://github.com/cloudevents/spec) industry standard format hosted by the [Cloud Native Computing Foundation ( CNCF)](https://www.cncf.io/).

	Audit uses version 0.1 specification of the CloudEvents event envelope. 

	Example: `0.1` 
* `content_type` - The content type of the data contained in `data`.  Example: `application/json` 
* `data` - The payload of the event. Information within `data` comes from the resource emitting the event. 
	* `additional_details` - A container object for attribues unique to the resource emitting the event.

		Example:

		  -----
		    {
		      "imageId": "ocid1.image.oc1.phx.<unique_ID>",
		      "shape": "VM.Standard1.1",
		      "type": "CustomerVmi"
		    }
		  -----
		
	* `availability_domain` - The availability domain where the resource resides. 
	* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment of the resource  emitting the event. 
	* `compartment_name` - The name of the compartment. This value is the friendly name associated with compartmentId. This value can change, but the service logs the value that appeared at the time of the audit event.  Example: `CompartmentA` 
	* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
	* `event_grouping_id` - This value links multiple audit events that are part of the same API operation. For example,  a long running API operations that emit an event at the start and the end of an operation would use the same value in this field for both events. 
	* `event_name` - Name of the API operation that generated this event.  Example: `GetInstance` 
	* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name,  type, or namespace. Exists for cross-compatibility only. For more information,  see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
	* `identity` - A container object for identity attributes. 
		* `auth_type` - The type of authentication used.  Example: `natv` 
		* `caller_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the caller. The caller that made a  request on behalf of the prinicpal. 
		* `caller_name` - The name of the user or service. This value is the friendly name associated with `callerId`. 
		* `console_session_id` - This value identifies any Console session associated with this request. 
		* `credentials` - The credential ID of the user. This value is extracted from the HTTP 'Authorization' request header. It consists of the tenantId, userId, and user fingerprint, all delimited by a slash (/). 
		* `ip_address` - The IP address of the source of the request.  Example: `172.24.80.88` 
		* `principal_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the principal. 
		* `principal_name` - The name of the user or service. This value is the friendly name associated with `principalId`.  Example: `ExampleName` 
		* `tenant_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the tenant. 
		* `user_agent` - The user agent of the client that made the request.  Example: `Jersey/2.23 (HttpUrlConnection 1.8.0_212)` 
	* `request` - A container object for request attributes. 
		* `action` - The HTTP method of the request.  Example: `GET` 
		* `headers` - The HTTP header fields and values in the request.

			Example: 

			  -----
			    {
			      "opc-principal": [
			        "{\"tenantId\":\"ocid1.tenancy.oc1..<unique_ID>\",\"subjectId\":\"ocid1.user.oc1..<unique_ID>\",\"claims\":[{\"key\":\"pstype\",\"value\":\"natv\",\"issuer\":\"authService.oracle.com\"},{\"key\":\"h_host\",\"value\":\"iaas.r2.oracleiaas.com\",\"issuer\":\"h\"},{\"key\":\"h_opc-request-id\",\"value\":\"<unique_ID>\",\"issuer\":\"h\"},{\"key\":\"ptype\",\"value\":\"user\",\"issuer\":\"authService.oracle.com\"},{\"key\":\"h_date\",\"value\":\"Wed, 18 Sep 2019 00:10:58 UTC\",\"issuer\":\"h\"},{\"key\":\"h_accept\",\"value\":\"application/json\",\"issuer\":\"h\"},{\"key\":\"authorization\",\"value\":\"Signature headers=\\\"date (request-target) host accept opc-request-id\\\",keyId=\\\"ocid1.tenancy.oc1..<unique_ID>/ocid1.user.oc1..<unique_ID>/8c:b4:5f:18:e7:ec:db:08:b8:fa:d2:2a:7d:11:76:ac\\\",algorithm=\\\"rsa-pss-sha256\\\",signature=\\\"<unique_ID>\\\",version=\\\"1\\\"\",\"issuer\":\"h\"},{\"key\":\"h_(request-target)\",\"value\":\"get /20160918/instances/ocid1.instance.oc1.phx.<unique_ID>\",\"issuer\":\"h\"}]}"
			      ],
			      "Accept": [
			        "application/json"
			      ],
			      "X-Oracle-Auth-Client-CN": [
			        "splat-proxy-se-02302.node.ad2.r2"
			      ],
			      "X-Forwarded-Host": [
			        "compute-api.svc.ad1.r2"
			      ],
			      "Connection": [
			        "close"
			      ],
			      "User-Agent": [
			        "Jersey/2.23 (HttpUrlConnection 1.8.0_212)"
			      ],
			      "X-Forwarded-For": [
			        "172.24.80.88"
			      ],
			      "X-Real-IP": [
			        "172.24.80.88"
			      ],
			      "oci-original-url": [
			        "https://iaas.r2.oracleiaas.com/20160918/instances/ocid1.instance.oc1.phx.<unique_ID>"
			      ],
			      "opc-request-id": [
			        "<unique_ID>"
			      ],
			      "Date": [
			        "Wed, 18 Sep 2019 00:10:58 UTC"
			      ]
			    }              
			  -----
			
		* `id` - The opc-request-id of the request. 
		* `parameters` - The parameters supplied by the caller during this operation. 
		* `path` - The full path of the API request.  Example: `/20160918/instances/ocid1.instance.oc1.phx.<unique_ID>` 
	* `resource_id` - An [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) or some other ID for the resource emitting the event. 
	* `resource_name` - The name of the resource emitting the event. 
	* `response` - A container object for response attributes. 
		* `headers` - The headers of the response.

			Example:

			  -----
			    {
			      "ETag": [
			        "<unique_ID>"
			      ],
			      "Connection": [
			        "close"
			      ],
			      "Content-Length": [
			        "1828"
			      ],
			      "opc-request-id": [
			        "<unique_ID>"
			      ],
			      "Date": [
			        "Wed, 18 Sep 2019 00:10:59 GMT"
			      ],
			      "Content-Type": [
			        "application/json"
			      ]
			    }
			  -----
			
		* `message` - A friendly description of what happened during the operation. Use this for troubleshooting. 
		* `payload` - This value is included for backward compatibility with the Audit version 1 schema, where  it contained metadata of interest from the response payload.

			Example: 

			  -----
			    {
			      "resourceName": "my_instance",
			      "id": "ocid1.instance.oc1.phx.<unique_ID>"
			    }
			  -----
			
		* `response_time` - The time of the response to the audited request, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format.  Example: `2019-09-18T00:10:59.278Z` 
		* `status` - The status code of the response.  Example: `200` 
	* `state_change` - A container object for state change attributes. 
		* `current` - Provides the current state of fields that may have changed during an operation. To determine how the current operation changed a resource, compare the information in this attribute to  `previous`. 
		* `previous` - Provides the previous state of fields that may have changed during an operation. To determine how the current operation changed a resource, compare the information in this attribute to  `current`. 
* `event_id` - The GUID of the event. 
* `event_time` - The time the event occurred, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format.  Example: `2019-09-18T00:10:59.252Z` 
* `event_type` - The type of event that happened. 

	The service that produces the event can also add, remove, or change the meaning of a field. A service implementing these type changes would publish a new version of an `eventType` and revise the `eventTypeVersion` field.

	Example: `com.oraclecloud.ComputeApi.GetInstance` 
* `event_type_version` - The version of the event type. This version applies to the payload of the event, not the envelope. Use `cloudEventsVersion` to determine the version of the envelope.  Example: `2.0` 
* `source` - The source of the event.  Example: `ComputeApi` 

