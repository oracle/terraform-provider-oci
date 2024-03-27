---
subcategory: "Logging"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_logging_unified_agent_configuration"
sidebar_current: "docs-oci-datasource-logging-unified_agent_configuration"
description: |-
  Provides details about a specific Unified Agent Configuration in Oracle Cloud Infrastructure Logging service
---

# Data Source: oci_logging_unified_agent_configuration
This data source provides details about a specific Unified Agent Configuration resource in Oracle Cloud Infrastructure Logging service.

Get the unified agent configuration for an ID.

## Example Usage

```hcl
data "oci_logging_unified_agent_configuration" "test_unified_agent_configuration" {
	#Required
	unified_agent_configuration_id = oci_logging_unified_agent_configuration.test_unified_agent_configuration.id
}
```

## Argument Reference

The following arguments are supported:

* `unified_agent_configuration_id` - (Required) The OCID of the Unified Agent configuration.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that the resource belongs to.
* `configuration_state` - State of unified agent service configuration.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - Description for this resource.
* `display_name` - The user-friendly display name. This must be unique within the enclosing resource, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `group_association` - Groups using the configuration.
	* `group_list` - list of group/dynamic group ids associated with this configuration.
* `id` - The OCID of the resource.
* `is_enabled` - Whether or not this resource is currently enabled.
* `service_configuration` - Top level Unified Agent service configuration object.
	* `application_configurations` - Unified Agent monitoring application configuration details.
		* `destination` - Kubernetes destination object.
			* `compartment_id` - The OCID of the compartment that the resource belongs to.
			* `metrics_namespace` - Namespace to which metrics will be emitted.
		* `source` - Kubernetes source object.
			* `name` - Unique name for the source.
			* `scrape_targets` - List of UnifiedAgentKubernetesScrapeTarget.
				* `k8s_namespace` - K8s namespace of the resource.
				* `name` - Custom name.
				* `resource_group` - Resource group in Oracle Cloud Infrastructure monitoring.
				* `resource_type` - Type of resource to scrape metrics.
				* `service_name` - Name of the service prepended to the endpoints.
				* `url` - URL from which the metrics are fetched.
		* `source_type` - Type of source of metrics
		* `sources` - Tail log source objects.
			* `advanced_options` - Advanced options for logging configuration
				* `is_read_from_head` - Starts to read the logs from the head of the file or the last read position recorded in pos_file, not tail.
			* `name` - Unique name for the source.
			* `parser` - Source parser object.
				* `delimiter` - CSV delimiter.
				* `expression` - Regex pattern.
				* `field_time_key` - Specifies the time field for the event time. If the event doesn't have this field, the current time is used.
				* `format` - Mutiline pattern format.
				* `format_firstline` - First line pattern format.
				* `grok_failure_key` - Grok failure key.
				* `grok_name_key` - Grok name key.
				* `is_estimate_current_event` - If true, use Fluent::EventTime.now(current time) as a timestamp when the time_key is specified.
				* `is_keep_time_key` - If true, keep the time field in the record.
				* `is_merge_cri_fields` - If you don't need stream or logtag fields, set this to false.
				* `is_null_empty_string` - If true, an empty string field is replaced with a null value.
				* `is_support_colonless_ident` - Specifies whether or not to support colonless ident. Corresponds to the Fluentd support_colonless_ident parameter.
				* `is_with_priority` - Specifies with priority or not. Corresponds to the Fluentd with_priority parameter.
				* `keys` - CSV keys.
				* `message_format` - Syslog message format.
				* `message_key` - Specifies the field name to contain logs.
				* `multi_line_start_regexp` - Multiline start regexp pattern.
				* `nested_parser` - Optional nested JSON Parser for CRI. Supported fields are fieldTimeKey, timeFormat, and isKeepTimeKey.
					* `parse_nested` - If true, a separator parameter can be further defined.
					* `separator` - Keys of adjacent levels are joined by the separator.
					* `time_format` - Process time value using the specified format.
					* `time_type` - JSON parser time type.
				* `null_value_pattern` - Specify the null value pattern.
				* `parse_nested` - If true, a separator parameter can be further defined.
				* `parser_type` - Type of fluent parser.
				* `patterns` - Grok pattern object.
					* `field_time_format` - Process value using the specified format. This is available only when time_type is a string.
					* `field_time_key` - Specify the time field for the event time. If the event doesn't have this field, the current time is used.
					* `field_time_zone` - Use the specified time zone. The time value can be parsed or formatted in the specified time zone.
					* `name` - The name key to tag this Grok pattern.
					* `pattern` - The Grok pattern.
				* `record_input` - record section of openmetrics parser. 
					* `dimensions` - Dimensions to be added for metrics.
					* `namespace` - Namespace to emit metrics.
					* `resource_group` - Resource group to emit metrics.
				* `rfc5424time_format` - RFC 5424 time format.
				* `separator` - Keys of adjacent levels are joined by the separator.
				* `syslog_parser_type` - Syslog parser type.
				* `time_format` - Process time value using the specified format.
				* `time_type` - JSON parser time type.
				* `timeout_in_milliseconds` - Specify the timeout for parse processing. This is mainly for detecting an incorrect regexp pattern.
				* `types` - Specify types for converting a field into another type. For example, With this configuration: <parse> @type csv keys time,host,req_id,user time_key time </parse>

					This incoming event: "2013/02/28 12:00:00,192.168.0.1,111,-"

					is parsed as: 1362020400 (2013/02/28/ 12:00:00)

					record: { "host"   : "192.168.0.1", "req_id" : "111", "user"   : "-" } 
			* `paths` - Absolute paths for log source files. Wildcards can be used.
			* `source_type` - Unified schema logging source type.
		* `unified_agent_configuration_filter` - Kubernetes filter object
			* `allow_list` - List of metrics regex to be allowed.
			* `deny_list` - List of metrics regex to be denied.
			* `filter_type` - Unified schema logging filter type.
			* `name` - Unique name for the filter.
	* `configuration_type` - Type of Unified Agent service configuration.
	* `destination` - Logging destination object.
		* `log_object_id` - The OCID of the resource.
		* `operational_metrics_configuration` - Unified monitoring agent operational metrics configuration object.
			* `destination` - Unified monitoring agent operational metrics destination object.
				* `compartment_id` - The OCID of the compartment that the resource belongs to.
			* `source` - Unified monitoring agent operational metrics source object.
				* `metrics` - List of unified monitoring agent operational metrics.
				* `record_input` - Record section of OperationalMetricsSource object.
					* `namespace` - Namespace to emit the operational metrics.
					* `resource_group` - Resource group to emit the operational metrics.
				* `type` - Type of the unified monitoring agent operational metrics source object.
	* `sources` - Logging source object.
		* `advanced_options` - Advanced options for logging configuration
			* `is_read_from_head` - Starts to read the logs from the head of the file or the last read position recorded in pos_file, not tail.
		* `channels` - Windows event log channels.
		* `custom_plugin` - User customized source plugin.
		* `name` - Unique name for the source.
		* `parser` - Source parser object.
			* `delimiter` - CSV delimiter.
			* `expression` - Regex pattern.
			* `field_time_key` - Specifies the time field for the event time. If the event doesn't have this field, the current time is used.
			* `format` - Mutiline pattern format.
			* `format_firstline` - First line pattern format.
			* `grok_failure_key` - Grok failure key.
			* `grok_name_key` - Grok name key.
			* `is_estimate_current_event` - If true, use Fluent::EventTime.now(current time) as a timestamp when the time_key is specified.
			* `is_keep_time_key` - If true, keep the time field in the record.
			* `is_merge_cri_fields` - If you don't need stream or logtag fields, set this to false.
			* `is_null_empty_string` - If true, an empty string field is replaced with a null value.
			* `is_support_colonless_ident` - Specifies whether or not to support colonless ident. Corresponds to the Fluentd support_colonless_ident parameter.
			* `is_with_priority` - Specifies with priority or not. Corresponds to the Fluentd with_priority parameter.
			* `keys` - CSV keys.
			* `message_format` - Syslog message format.
			* `message_key` - Specifies the field name to contain logs.
			* `multi_line_start_regexp` - Multiline start regexp pattern.
			* `nested_parser` - Optional nested JSON Parser for CRI. Supported fields are fieldTimeKey, timeFormat, and isKeepTimeKey.
				* `parse_nested` - If true, a separator parameter can be further defined.
				* `separator` - Keys of adjacent levels are joined by the separator.
				* `time_format` - Process time value using the specified format.
				* `time_type` - JSON parser time type.
			* `null_value_pattern` - Specify the null value pattern.
			* `parse_nested` - If true, a separator parameter can be further defined.
			* `parser_type` - Type of fluent parser.
			* `patterns` - Grok pattern object.
				* `field_time_format` - Process value using the specified format. This is available only when time_type is a string.
				* `field_time_key` - Specify the time field for the event time. If the event doesn't have this field, the current time is used.
				* `field_time_zone` - Use the specified time zone. The time value can be parsed or formatted in the specified time zone.
				* `name` - The name key to tag this Grok pattern.
				* `pattern` - The Grok pattern.
			* `record_input` - record section of openmetrics parser. 
				* `dimensions` - Dimensions to be added for metrics.
				* `namespace` - Namespace to emit metrics.
				* `resource_group` - Resource group to emit metrics.
			* `rfc5424time_format` - RFC 5424 time format.
			* `separator` - Keys of adjacent levels are joined by the separator.
			* `syslog_parser_type` - Syslog parser type.
			* `time_format` - Process time value using the specified format.
			* `time_type` - JSON parser time type.
			* `timeout_in_milliseconds` - Specify the timeout for parse processing. This is mainly for detecting an incorrect regexp pattern.
			* `types` - Specify types for converting a field into another type. For example, With this configuration: <parse> @type csv keys time,host,req_id,user time_key time </parse>

				This incoming event: "2013/02/28 12:00:00,192.168.0.1,111,-"

				is parsed as: 1362020400 (2013/02/28/ 12:00:00)

				record: { "host"   : "192.168.0.1", "req_id" : "111", "user"   : "-" } 
		* `paths` - Absolute paths for log source files. Wildcards can be used.
		* `source_type` - Unified schema logging source type.
	* `unified_agent_configuration_filter` - Logging filter object.
		* `allow_list` - A list of filtering rules to include logs
			* `key` - The field name to which the regular expression is applied
			* `pattern` - The regular expression
		* `custom_filter_type` - Type of the custom filter
		* `custom_sections` - List of custom sections in custom filter
			* `name` - The name of the custom section
			* `params` - Parameters in the custom section
		* `deny_list` - A list of filtering rules to reject logs
			* `key` - The field name to which the regular expression is applied
			* `pattern` - The regular expression
		* `emit_invalid_record_to_error` - If true, emit invalid record to @ERROR label. Invalid cases are: 1) key does not exist; 2) the format does not match; or 3) an unexpected error. You can rescue unexpected format logs in the @ERROR lable. If you want to ignore these errors, set this to false. 
		* `filter_type` - Unified schema logging filter type.
		* `hash_value_field` - Store the parsed values as a hash value in a field.
		* `inject_key_prefix` - Store the parsed values with the specified key name prefix.
		* `is_auto_typecast_enabled` - If true, automatically casts the field types.
		* `is_renew_record_enabled` - If true, it modifies a new empty hash
		* `is_ruby_enabled` - When set to true, the full Ruby syntax is enabled in the ${} expression.
		* `keep_keys` - A list of keys to keep. Only relevant if isRenewRecordEnabled is set to true
		* `key_name` - The field name in the record to parse.
		* `name` - Unique name for the filter.
		* `params` - Parameters of the custom filter
		* `parser` - Source parser object.
			* `delimiter` - CSV delimiter.
			* `expression` - Regex pattern.
			* `field_time_key` - Specifies the time field for the event time. If the event doesn't have this field, the current time is used.
			* `format` - Mutiline pattern format.
			* `format_firstline` - First line pattern format.
			* `grok_failure_key` - Grok failure key.
			* `grok_name_key` - Grok name key.
			* `is_estimate_current_event` - If true, use Fluent::EventTime.now(current time) as a timestamp when the time_key is specified.
			* `is_keep_time_key` - If true, keep the time field in the record.
			* `is_merge_cri_fields` - If you don't need stream or logtag fields, set this to false.
			* `is_null_empty_string` - If true, an empty string field is replaced with a null value.
			* `is_support_colonless_ident` - Specifies whether or not to support colonless ident. Corresponds to the Fluentd support_colonless_ident parameter.
			* `is_with_priority` - Specifies with priority or not. Corresponds to the Fluentd with_priority parameter.
			* `keys` - CSV keys.
			* `message_format` - Syslog message format.
			* `message_key` - Specifies the field name to contain logs.
			* `multi_line_start_regexp` - Multiline start regexp pattern.
			* `nested_parser` - Optional nested JSON Parser for CRI. Supported fields are fieldTimeKey, timeFormat, and isKeepTimeKey.
				* `parse_nested` - If true, a separator parameter can be further defined.
				* `separator` - Keys of adjacent levels are joined by the separator.
				* `time_format` - Process time value using the specified format.
				* `time_type` - JSON parser time type.
			* `null_value_pattern` - Specify the null value pattern.
			* `parse_nested` - If true, a separator parameter can be further defined.
			* `parser_type` - Type of fluent parser.
			* `patterns` - Grok pattern object.
				* `field_time_format` - Process value using the specified format. This is available only when time_type is a string.
				* `field_time_key` - Specify the time field for the event time. If the event doesn't have this field, the current time is used.
				* `field_time_zone` - Use the specified time zone. The time value can be parsed or formatted in the specified time zone.
				* `name` - The name key to tag this Grok pattern.
				* `pattern` - The Grok pattern.
			* `record_input` - record section of openmetrics parser. 
				* `dimensions` - Dimensions to be added for metrics.
				* `namespace` - Namespace to emit metrics.
				* `resource_group` - Resource group to emit metrics.
			* `rfc5424time_format` - RFC 5424 time format.
			* `separator` - Keys of adjacent levels are joined by the separator.
			* `syslog_parser_type` - Syslog parser type.
			* `time_format` - Process time value using the specified format.
			* `time_type` - JSON parser time type.
			* `timeout_in_milliseconds` - Specify the timeout for parse processing. This is mainly for detecting an incorrect regexp pattern.
			* `types` - Specify types for converting a field into another type. For example, With this configuration: <parse> @type csv keys time,host,req_id,user time_key time </parse>

				This incoming event: "2013/02/28 12:00:00,192.168.0.1,111,-"

				is parsed as: 1362020400 (2013/02/28/ 12:00:00)

				record: { "host"   : "192.168.0.1", "req_id" : "111", "user"   : "-" } 
		* `record_list` - Add new key-value pairs in logs
			* `key` - A new key
			* `value` - A new value
		* `remove_key_name_field` - If true, remove the keyName field when parsing is succeeded.
		* `remove_keys` - A list of keys to delete
		* `renew_time_key` - Overwrites the time of logs with this value, this value must be a Unix timestamp.
		* `replace_invalid_sequence` - If true, the invalid string is replaced with safe characters and is re-parsed.
		* `reserve_data` - If true, keep the original key-value pair in the parsed result.
		* `reserve_time` - If true, keep the original event time in the parsed result.
* `state` - The pipeline state.
* `time_created` - Time the resource was created.
* `time_last_modified` - Time the resource was last modified.

