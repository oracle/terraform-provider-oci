---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_health_checks_ping_probe_results"
sidebar_current: "docs-oci-datasource-health_checks-ping_probe_results"
description: |-
  Provides the list of Ping Probe Results in Oracle Cloud Infrastructure Health Checks service
---

# Data Source: oci_health_checks_ping_probe_results
This data source provides the list of Ping Probe Results in Oracle Cloud Infrastructure Health Checks service.

Returns the results for the specified probe, where the `probeConfigurationId`
is the OCID of either a monitor or an on-demand probe.

Results are paginated based on `page` and `limit`.  The `opc-next-page` header provides
a URL for fetching the next page.  Use `sortOrder` to set the order of the
results.  If `sortOrder` is unspecified, results are sorted in ascending order by
`startTime`.


## Example Usage

```hcl
data "oci_health_checks_ping_probe_results" "test_ping_probe_results" {
	#Required
	probe_configuration_id = "${oci_health_checks_probe_configuration.test_probe_configuration.id}"

	#Optional
	start_time_greater_than_or_equal_to = "${var.ping_probe_result_start_time_greater_than_or_equal_to}"
	start_time_less_than_or_equal_to = "${var.ping_probe_result_start_time_less_than_or_equal_to}"
	target = "${var.ping_probe_result_target}"
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

* `ping_probe_results` - The list of ping_probe_results.

### PingProbeResult Reference

The following attributes are exported:

* `connection` - 
	* `address` - The connection IP address.
	* `port` - The port.
* `dns` - 
	* `addresses` - The addresses returned by DNS resolution.
	* `domain_lookup_duration` - Total DNS resolution duration, in milliseconds. Calculated using `domainLookupEnd` minus `domainLookupStart`. 
* `domain_lookup_end` - The time immediately before the vantage point finishes the domain name lookup for the resource. 
* `domain_lookup_start` - The time immediately before the vantage point starts the domain name lookup for the resource. 
* `error_category` - The category of error if an error occurs executing the probe. The `errorMessage` field provides a message with the error details.
	* NONE - No error
	* DNS - DNS errors
	* TRANSPORT - Transport-related errors, for example a "TLS certificate expired" error.
	* NETWORK - Network-related errors, for example a "network unreachable" error.
	* SYSTEM - Internal system errors. 
* `error_message` - The error information indicating why a probe execution failed.
* `icmp_code` - The ICMP code of the response message.  This field is not used when the protocol is set to TCP.  For more information on ICMP codes, see [Internet Control Message Protocol (ICMP) Parameters](https://www.iana.org/assignments/icmp-parameters/icmp-parameters.xhtml). 
* `is_healthy` - True if the probe result is determined to be healthy based on probe type-specific criteria.  For HTTP probes, a probe result is considered healthy if the HTTP response code is greater than or equal to 200 and less than 300. 
* `is_timed_out` - True if the probe did not complete before the configured `timeoutInSeconds` value. 
* `key` - A value identifying this specific probe result. The key is only unique within the results of its probe configuration. The key may be reused after 90 days. 
* `latency_in_ms` - The latency of the probe execution, in milliseconds. 
* `probe_configuration_id` - The OCID of the monitor or on-demand probe responsible for creating this result. 
* `protocol` - The protocols for ping probes.
* `start_time` - The date and time the probe was executed, expressed in milliseconds since the POSIX epoch. This field is defined by the PerformanceResourceTiming interface of the W3C Resource Timing specification. For more information, see [Resource Timing](https://w3c.github.io/resource-timing/#sec-resource-timing). 
* `target` - The target hostname or IP address of the probe.
* `vantage_point_name` - The name of the vantage point that executed the probe.

