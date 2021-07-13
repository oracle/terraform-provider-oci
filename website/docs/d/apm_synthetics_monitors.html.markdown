---
subcategory: "Apm Synthetics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_apm_synthetics_monitors"
sidebar_current: "docs-oci-datasource-apm_synthetics-monitors"
description: |-
  Provides the list of Monitors in Oracle Cloud Infrastructure Apm Synthetics service
---

# Data Source: oci_apm_synthetics_monitors
This data source provides the list of Monitors in Oracle Cloud Infrastructure Apm Synthetics service.

Returns a list of monitors.


## Example Usage

```hcl
data "oci_apm_synthetics_monitors" "test_monitors" {
	#Required
	apm_domain_id = oci_apm_synthetics_apm_domain.test_apm_domain.id

	#Optional
	display_name = var.monitor_display_name
	monitor_type = var.monitor_monitor_type
	script_id = oci_apm_synthetics_script.test_script.id
	status = var.monitor_status
}
```

## Argument Reference

The following arguments are supported:

* `apm_domain_id` - (Required) The APM domain ID the request is intended for. 
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `monitor_type` - (Optional) A filter to return only monitors that match the given monitor type. Supported values are SCRIPTED_BROWSER, BROWSER, SCRIPTED_REST and REST. 
* `script_id` - (Optional) A filter to return only monitors using scriptId.
* `status` - (Optional) A filter to return only monitors that match the status given.


## Attributes Reference

The following attributes are exported:

* `monitor_collection` - The list of monitor_collection.

### Monitor Reference

The following attributes are exported:

* `configuration` - Details of monitor configuration.
	* `config_type` - Type of configuration.
	* `is_certificate_validation_enabled` - If certificate validation is enabled, then the call will fail in case of certification errors.
	* `is_failure_retried` - If isFailureRetried is enabled, then a failed call will be retried.
	* `is_redirection_enabled` - If redirection enabled, then redirects will be allowed while accessing target URL.
	* `network_configuration` - Details of the network configuration.
		* `number_of_hops` - Number of hops.
		* `probe_mode` - Type of probe mode when TCP protocol is selected.
		* `probe_per_hop` - Number of probes per hop.
		* `protocol` - Type of protocol.
		* `transmission_rate` - Number of probe packets sent out simultaneously.
	* `req_authentication_details` - Details for request HTTP authentication.
		* `auth_headers` - List of authentication headers. Example: `[{"headerName": "content-type", "headerValue":"json"}]` 
			* `header_name` - Name of the header.
			* `header_value` - Value of the header.
		* `auth_request_method` - Request method.
		* `auth_request_post_body` - Request post body.
		* `auth_token` - Authentication token.
		* `auth_url` - URL to get authetication token.
		* `auth_user_name` - Username for authentication.
		* `auth_user_password` - User password for authentication.
		* `oauth_scheme` - Request http oauth scheme.
	* `req_authentication_scheme` - Request http authentication scheme.
	* `request_headers` - List of request headers. Example: `[{"headerName": "content-type", "headerValue":"json"}]` 
		* `header_name` - Name of the header.
		* `header_value` - Value of the header.
	* `request_method` - Request HTTP method.
	* `request_post_body` - Request post body content.
	* `request_query_params` - List of request query params. Example: `[{"paramName": "sortOrder", "paramValue": "asc"}]` 
		* `param_name` - Name of request query parameter.
		* `param_value` - Value of request query parameter.
	* `verify_response_codes` - Expected HTTP response codes. For status code range, set values such as 2xx, 3xx. 
	* `verify_response_content` - Verify response content against regular expression based string. If response content does not match the verifyResponseContent value, then it will be considered a failure. 
	* `verify_texts` - Verify all the search strings present in response. If any search string is not present in the response, then it will be considered as a failure. 
		* `text` - Verification text in the response.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - Unique name that can be edited. The name should not contain any confidential information.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the monitor.
* `is_run_once` - If runOnce is enabled, then the monitor will run once.
* `monitor_type` - Type of the monitor.
* `repeat_interval_in_seconds` - Interval in seconds after the start time when the job should be repeated. Minimum repeatIntervalInSeconds should be 300 seconds. 
* `script_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the script. scriptId is mandatory for creation of SCRIPTED_BROWSER and SCRIPTED_REST monitor types. For other monitor types, it should be set to null. 
* `script_name` - Name of the script.
* `script_parameters` - List of script parameters. Example: `[{"monitorScriptParameter": {"paramName": "userid", "paramValue":"testuser"}, "isSecret": false, "isOverwritten": false}]` 
	* `is_overwritten` - If parameter value is default or overwritten. 
	* `is_secret` - Describes if  the parameter value is secret and should be kept confidential. isSecret is specified in either CreateScript or UpdateScript API. 
	* `monitor_script_parameter` - Details of the script parameter that can be used to overwrite the parameter present in the script. 
		* `param_name` - Name of the parameter.
		* `param_value` - Value of the parameter.
* `status` - Enables or disables the monitor.
* `target` - Specify the endpoint on which to run the monitor. For BROWSER and REST monitor types, target is mandatory. If target is specified in the SCRIPTED_BROWSER monitor type, then the monitor will run the selected script (specified by scriptId in monitor) against the specified target endpoint. If target is not specified in the SCRIPTED_BROWSER monitor type, then the monitor will run the selected script as it is. 
* `time_created` - The time the resource was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2020-02-12T22:47:12.613Z` 
* `time_updated` - The time the resource was updated, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2020-02-13T22:47:12.613Z` 
* `timeout_in_seconds` - Timeout in seconds. Timeout cannot be more than 30% of repeatIntervalInSeconds time for monitors. Also, timeoutInSeconds should be a multiple of 60. Monitor will be allowed to run only for timeoutInSeconds time. It would be terminated after that. 
* `vantage_point_count` - Number of vantage points where monitor is running.
* `vantage_points` - List of vantage points from where monitor is running.
	* `display_name` - Unique name that can be edited. The name should not contain any confidential information.
	* `name` - Name of the vantage point.

