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
	is_maintenance_window_active = var.monitor_is_maintenance_window_active
	is_maintenance_window_set = var.monitor_is_maintenance_window_set
	monitor_type = var.monitor_monitor_type
	script_id = oci_apm_synthetics_script.test_script.id
	status = var.monitor_status
	vantage_point = var.monitor_vantage_point
}
```

## Argument Reference

The following arguments are supported:

* `apm_domain_id` - (Required) The APM domain ID the request is intended for. 
* `display_name` - (Optional) A filter to return only the resources that match the entire display name.
* `is_maintenance_window_active` - (Optional) A filter to return the monitors whose maintenance window is currently active.
* `is_maintenance_window_set` - (Optional) A filter to return the monitors whose maintenance window is set.
* `monitor_type` - (Optional) A filter to return only monitors that match the given monitor type. Supported values are SCRIPTED_BROWSER, BROWSER, SCRIPTED_REST, REST, NETWORK, DNS, FTP and SQL. 
* `script_id` - (Optional) A filter to return only monitors using scriptId.
* `status` - (Optional) A filter to return only monitors that match the status given.
* `vantage_point` - (Optional) The name of the public or dedicated vantage point. 


## Attributes Reference

The following attributes are exported:

* `monitor_collection` - The list of monitor_collection.

### Monitor Reference

The following attributes are exported:

* `availability_configuration` - Monitor availability configuration details.
	* `max_allowed_failures_per_interval` - Maximum number of failed runs allowed in an interval. If an interval has more failed runs than the specified value, then the interval will be classified as UNAVAILABLE.
	* `min_allowed_runs_per_interval` - Minimum number of runs allowed in an interval. If an interval has fewer runs than the specified value, then the interval will be classified as UNKNOWN and will be excluded from the availability calculations.
* `batch_interval_in_seconds` - Time interval between two runs in round robin batch mode (SchedulingPolicy - BATCHED_ROUND_ROBIN).
* `configuration` - Details of monitor configuration.
	* `client_certificate_details` - Details for client certificate.
		* `client_certificate` - Client certificate in PEM format.
			* `content` - Content of the client certificate file.
			* `file_name` - Name of the certificate file. The name should not contain any confidential information.
		* `private_key` - The private key associated with the client certificate in PEM format.
			* `content` - Content of the private key file.
			* `file_name` - Name of the private key file.
	* `config_type` - Type of configuration.
	* `connection_string` - Database connection string.
	* `database_authentication_details` - Details for basic authentication.
		* `password` - Password.
			* `password` - Password.
			* `password_type` - Type of method to pass password.
			* `vault_secret_id` - Vault secret OCID.
		* `username` - Username for authentication.
	* `database_connection_type` - Database connection type. Only CUSTOM_JDBC is supported for MYSQL database type.
	* `database_role` - Database role.
	* `database_type` - Database type.
	* `database_wallet_details` - Details for database wallet.
		* `database_wallet` - The database wallet configuration zip file.
		* `service_name` - Service name of the database.
	* `dns_configuration` - Information about the DNS settings.
		* `is_override_dns` - If isOverrideDns is true, then DNS settings will be overridden.
		* `override_dns_ip` - Attribute to override the DNS IP value. This value will be honored only if isOverrideDns is set to true.
	* `download_size_limit_in_bytes` - Download size limit in Bytes, at which to stop the transfer. Maximum download size limit is 5 MiB.
	* `ftp_basic_authentication_details` - Details for basic authentication.
		* `password` - Password.
			* `password` - Password.
			* `password_type` - Type of method to pass password.
			* `vault_secret_id` - Vault secret OCID.
		* `username` - Username for authentication.
	* `ftp_protocol` - FTP protocol type.
	* `ftp_request_type` - FTP monitor request type.
	* `is_active_mode` - If enabled, Active mode will be used for the FTP connection.
	* `is_certificate_validation_enabled` - If certificate validation is enabled, then the call will fail in case of certification errors.
	* `is_default_snapshot_enabled` - If disabled, auto snapshots are not collected.
	* `is_failure_retried` - If isFailureRetried is enabled, then a failed call will be retried.
	* `is_query_recursive` - If isQueryRecursive is enabled, then queries will be sent recursively to the target server.
	* `is_redirection_enabled` - If redirection is enabled, then redirects will be allowed while accessing target URL.
	* `name_server` - Name of the server that will be used to perform DNS lookup.
	* `network_configuration` - Details of the network configuration. For NETWORK monitor type, NetworkConfiguration is mandatory.
		* `number_of_hops` - Number of hops.
		* `probe_mode` - Type of probe mode when TCP protocol is selected.
		* `probe_per_hop` - Number of probes per hop.
		* `protocol` - Type of protocol.
		* `transmission_rate` - Number of probe packets sent out simultaneously.
	* `protocol` - Type of protocol.
	* `query` - SQL query to be executed.
	* `record_type` - DNS record type.
	* `req_authentication_details` - Details for request HTTP authentication.
		* `auth_headers` - List of authentication headers. Example: `[{"headerName": "content-type", "headerValue":"json"}]` 
			* `header_name` - Name of the header.
			* `header_value` - Value of the header.
		* `auth_request_method` - Request method.
		* `auth_request_post_body` - Request post body.
		* `auth_token` - Authentication token.
		* `auth_url` - URL to get authentication token.
		* `auth_user_name` - User name for authentication.
		* `auth_user_password` - User password for authentication.
		* `oauth_scheme` - Request HTTP OAuth scheme.
	* `req_authentication_scheme` - Request HTTP authentication scheme.
	* `request_headers` - List of request headers. Example: `[{"headerName": "content-type", "headerValue":"json"}]` 
		* `header_name` - Name of the header.
		* `header_value` - Value of the header.
	* `request_method` - Request HTTP method.
	* `request_post_body` - Request post body content.
	* `request_query_params` - List of request query params. Example: `[{"paramName": "sortOrder", "paramValue": "asc"}]` 
		* `param_name` - Name of request query parameter.
		* `param_value` - Value of request query parameter.
	* `upload_file_size_in_bytes` - File upload size in Bytes, at which to stop the transfer. Maximum upload size is 5 MiB.
	* `verify_response_codes` - Expected HTTP response codes. For status code range, set values such as 2xx, 3xx. 
	* `verify_response_content` - Verify response content against regular expression based string. If response content does not match the verifyResponseContent value, then it will be considered a failure. 
	* `verify_texts` - Verifies all the search strings present in the response. If any search string is not present in the response, then it will be considered as a failure. 
		* `text` - Verification text in the response.
* `created_by` - Name of the user that created the monitor.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - Unique name that can be edited. The name should not contain any confidential information.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the monitor.
* `is_ipv6` - If enabled, domain name will resolve to an IPv6 address.
* `is_run_now` - If isRunNow is enabled, then the monitor will run immediately.
* `is_run_once` - If runOnce is enabled, then the monitor will run once.
* `last_updated_by` - Name of the user that recently updated the monitor.
* `maintenance_window_schedule` - Details required to schedule maintenance window.
	* `time_ended` - End time of the maintenance window, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2020-02-12T22:47:12.613Z` 
	* `time_started` - Start time of the maintenance window, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2020-02-12T22:47:12.613Z` 
* `monitor_type` - Type of monitor.
* `repeat_interval_in_seconds` - Interval in seconds after the start time when the job should be repeated. Minimum repeatIntervalInSeconds should be 300 seconds for Scripted REST, Scripted Browser and Browser monitors, and 60 seconds for REST monitor. 
* `scheduling_policy` - Scheduling policy to decide the distribution of monitor executions on vantage points.
* `script_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the script. scriptId is mandatory for creation of SCRIPTED_BROWSER and SCRIPTED_REST monitor types. For other monitor types, it should be set to null. 
* `script_name` - Name of the script.
* `script_parameters` - List of script parameters. Example: `[{"monitorScriptParameter": {"paramName": "userid", "paramValue":"testuser"}, "isSecret": false, "isOverwritten": false}]` 
	* `is_overwritten` - If parameter value is default or overwritten. 
	* `is_secret` - Describes if  the parameter value is secret and should be kept confidential. isSecret is specified in either CreateScript or UpdateScript API. 
	* `monitor_script_parameter` - Details of the script parameter that can be used to overwrite the parameter present in the script. 
		* `param_name` - Name of the parameter.
		* `param_value` - Value of the parameter.
* `status` - Enables or disables the monitor.
* `target` - Specify the endpoint on which to run the monitor. For BROWSER, REST, NETWORK, DNS and FTP monitor types, target is mandatory. If target is specified in the SCRIPTED_BROWSER monitor type, then the monitor will run the selected script (specified by scriptId in monitor) against the specified target endpoint. If target is not specified in the SCRIPTED_BROWSER monitor type, then the monitor will run the selected script as it is. For NETWORK monitor with TCP protocol, a port needs to be provided along with target. Example: 192.168.0.1:80. 
* `time_created` - The time the resource was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2020-02-12T22:47:12.613Z` 
* `time_updated` - The time the resource was updated, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2020-02-13T22:47:12.613Z` 
* `timeout_in_seconds` - Timeout in seconds. If isFailureRetried is true, then timeout cannot be more than 30% of repeatIntervalInSeconds time for monitors. If isFailureRetried is false, then timeout cannot be more than 50% of repeatIntervalInSeconds time for monitors. Also, timeoutInSeconds should be a multiple of 60 for Scripted REST, Scripted Browser and Browser monitors. Monitor will be allowed to run only for timeoutInSeconds time. It would be terminated after that. 
* `vantage_point_count` - Number of vantage points where monitor is running.
* `vantage_points` - List of public, dedicated and onPremise vantage points where the monitor is running.
	* `display_name` - Unique name that can be edited. The name should not contain any confidential information.
	* `name` - Name of the vantage point.
	* `worker_list` - List of workers running the assigned monitor.

