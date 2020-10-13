---
subcategory: "Health Checks"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_health_checks_http_probe_results"
sidebar_current: "docs-oci-datasource-health_checks-http_probe_results"
description: |-
  Provides the list of Http Probe Results in Oracle Cloud Infrastructure Health Checks service
---

# Data Source: oci_health_checks_http_probe_results
This data source provides the list of Http Probe Results in Oracle Cloud Infrastructure Health Checks service.

Gets the HTTP probe results for the specified probe or monitor, where
the `probeConfigurationId` is the OCID of either a monitor or an
on-demand probe.


## Example Usage

```hcl
data "oci_health_checks_http_probe_results" "test_http_probe_results" {
	#Required
	probe_configuration_id = oci_health_checks_probe_configuration.test_probe_configuration.id

	#Optional
	start_time_greater_than_or_equal_to = var.http_probe_result_start_time_greater_than_or_equal_to
	start_time_less_than_or_equal_to = var.http_probe_result_start_time_less_than_or_equal_to
	target = var.http_probe_result_target
}
```

## Argument Reference

The following arguments are supported:

* `probe_configuration_id` - (Required) The OCID of a monitor or on-demand probe.
* `start_time_greater_than_or_equal_to` - (Optional) Returns results with a `startTime` equal to or greater than the specified value.
* `start_time_less_than_or_equal_to` - (Optional) Returns results with a `startTime` equal to or less than the specified value.
* `target` - (Optional) Filters results that match the `target`.


## Attributes Reference

The following attributes are exported:

* `http_probe_results` - The list of http_probe_results.

### HttpProbeResult Reference

The following attributes are exported:

* `connect_end` - The time immediately after the vantage point finishes establishing the connection to the server to retrieve the resource. 
* `connect_start` - The time immediately before the vantage point starts establishing the connection to the server to retrieve the resource. 
* `connection` - TCP connection results.  All durations are in milliseconds.
	* `address` - The connection IP address.
	* `connect_duration` - Total connect duration, calculated using `connectEnd` minus `connectStart`.
	* `port` - The port.
	* `secure_connect_duration` - The duration to secure the connection.  This value will be zero for insecure connections.  Calculated using `connectEnd` minus `secureConnectionStart`. 
* `dns` - The DNS resolution results.
	* `addresses` - The addresses returned by DNS resolution.
	* `domain_lookup_duration` - Total DNS resolution duration, in milliseconds. Calculated using `domainLookupEnd` minus `domainLookupStart`. 
* `domain_lookup_end` - The time immediately before the vantage point finishes the domain name lookup for the resource. 
* `domain_lookup_start` - The time immediately before the vantage point starts the domain name lookup for the resource. 
* `duration` - The total duration from start of request until response is fully consumed or the connection is closed. 
* `encoded_body_size` - The size, in octets, of the payload body prior to removing any applied content-codings. 
* `error_category` - The category of error if an error occurs executing the probe. The `errorMessage` field provides a message with the error details.
	* NONE - No error
	* DNS - DNS errors
	* TRANSPORT - Transport-related errors, for example a "TLS certificate expired" error.
	* NETWORK - Network-related errors, for example a "network unreachable" error.
	* SYSTEM - Internal system errors. 
* `error_message` - The error information indicating why a probe execution failed.
* `fetch_start` - The time immediately before the vantage point starts to fetch the resource. 
* `is_healthy` - True if the probe result is determined to be healthy based on probe type-specific criteria.  For HTTP probes, a probe result is considered healthy if the HTTP response code is greater than or equal to 200 and less than 300. 
* `is_timed_out` - True if the probe did not complete before the configured `timeoutInSeconds` value. 
* `key` - A value identifying this specific probe result. The key is only unique within the results of its probe configuration. The key may be reused after 90 days. 
* `probe_configuration_id` - The OCID of the monitor or on-demand probe responsible for creating this result. 
* `protocol` - The supported protocols available for HTTP probes.
* `request_start` - The time immediately before the vantage point starts requesting the resource from the server. 
* `response_end` - The time immediately after the vantage point receives the last byte of the response or immediately before the transport connection is closed, whichever comes first. 
* `response_start` - The time immediately after the vantage point's HTTP parser receives the first byte of the response. 
* `secure_connection_start` - The time immediately before the vantage point starts the handshake process to secure the current connection. 
* `start_time` - The date and time the probe was executed, expressed in milliseconds since the POSIX epoch. This field is defined by the PerformanceResourceTiming interface of the W3C Resource Timing specification. For more information, see [Resource Timing](https://w3c.github.io/resource-timing/#sec-resource-timing). 
* `status_code` - The HTTP response status code.
* `target` - The target hostname or IP address of the probe.
* `vantage_point_name` - The name of the vantage point that executed the probe.

