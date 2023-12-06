---
subcategory: "Logging"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_logging_unified_agent_configuration"
sidebar_current: "docs-oci-resource-logging-unified_agent_configuration"
description: |-
  Provides the Unified Agent Configuration resource in Oracle Cloud Infrastructure Logging service
---

# oci_logging_unified_agent_configuration
This resource provides the Unified Agent Configuration resource in Oracle Cloud Infrastructure Logging service.

Create unified agent configuration registration.

## Example Usage

```hcl
resource "oci_logging_unified_agent_configuration" "test_unified_agent_configuration" {
	#Required
	compartment_id = var.compartment_id
	is_enabled = var.unified_agent_configuration_is_enabled
    description = var.unified_agent_configuration_description
	display_name = var.unified_agent_configuration_display_name
	service_configuration {
		#Required
		configuration_type = var.unified_agent_configuration_service_configuration_configuration_type

		destination {
			#Required
			log_object_id = oci_objectstorage_object.test_object.id

			#Optional
			operational_metrics_configuration {
				#Required
				destination {
					#Required
					compartment_id = var.compartment_id
				}
				source {
					#Required
					type = var.unified_agent_configuration_service_configuration_destination_operational_metrics_configuration_source_type

					#Optional
					metrics = var.unified_agent_configuration_service_configuration_destination_operational_metrics_configuration_source_metrics
					
					#Required
					record_input {
						#Required
						namespace = var.unified_agent_configuration_service_configuration_destination_operational_metrics_configuration_source_record_input_namespace

						#Optional
						resource_group = var.unified_agent_configuration_service_configuration_destination_operational_metrics_configuration_source_record_input_resource_group
					}
				}
			}
		}
		sources {
			#Required
			source_type = var.unified_agent_configuration_service_configuration_sources_source_type

			#Optional
			channels = var.unified_agent_configuration_service_configuration_sources_channels
			name = var.unified_agent_configuration_service_configuration_sources_name
			parser {
				#Required
				parser_type = var.unified_agent_configuration_service_configuration_sources_parser_parser_type

				#Optional
				delimiter = var.unified_agent_configuration_service_configuration_sources_parser_delimiter
				expression = var.unified_agent_configuration_service_configuration_sources_parser_expression
				field_time_key = var.unified_agent_configuration_service_configuration_sources_parser_field_time_key
				format = var.unified_agent_configuration_service_configuration_sources_parser_format
				format_firstline = var.unified_agent_configuration_service_configuration_sources_parser_format_firstline
				grok_failure_key = var.unified_agent_configuration_service_configuration_sources_parser_grok_failure_key
				grok_name_key = var.unified_agent_configuration_service_configuration_sources_parser_grok_name_key
				is_estimate_current_event = var.unified_agent_configuration_service_configuration_sources_parser_is_estimate_current_event
				is_keep_time_key = var.unified_agent_configuration_service_configuration_sources_parser_is_keep_time_key
				is_merge_cri_fields = var.unified_agent_configuration_service_configuration_sources_parser_is_merge_cri_fields
				is_null_empty_string = var.unified_agent_configuration_service_configuration_sources_parser_is_null_empty_string
				is_support_colonless_ident = var.unified_agent_configuration_service_configuration_sources_parser_is_support_colonless_ident
				is_with_priority = var.unified_agent_configuration_service_configuration_sources_parser_is_with_priority
				keys = var.unified_agent_configuration_service_configuration_sources_parser_keys
				message_format = var.unified_agent_configuration_service_configuration_sources_parser_message_format
				message_key = var.unified_agent_configuration_service_configuration_sources_parser_message_key
				multi_line_start_regexp = var.unified_agent_configuration_service_configuration_sources_parser_multi_line_start_regexp
				nested_parser {

					#Optional
					time_format = var.unified_agent_configuration_service_configuration_sources_parser_nested_parser_time_format
					field_time_key = var.unified_agent_configuration_service_configuration_sources_parser_nested_parser_field_time_key
					is_keep_time_key = var.unified_agent_configuration_service_configuration_sources_parser_nested_parser_is_keep_time_key
				}
				null_value_pattern = var.unified_agent_configuration_service_configuration_sources_parser_null_value_pattern
				patterns {

					#Optional
					field_time_format = var.unified_agent_configuration_service_configuration_sources_parser_patterns_field_time_format
					field_time_key = var.unified_agent_configuration_service_configuration_sources_parser_patterns_field_time_key
					field_time_zone = var.unified_agent_configuration_service_configuration_sources_parser_patterns_field_time_zone
					name = var.unified_agent_configuration_service_configuration_sources_parser_patterns_name
					pattern = var.unified_agent_configuration_service_configuration_sources_parser_patterns_pattern
				}
				rfc5424time_format = var.unified_agent_configuration_service_configuration_sources_parser_rfc5424time_format
				syslog_parser_type = var.unified_agent_configuration_service_configuration_sources_parser_syslog_parser_type
				time_format = var.unified_agent_configuration_service_configuration_sources_parser_time_format
				time_type = var.unified_agent_configuration_service_configuration_sources_parser_time_type
				timeout_in_milliseconds = var.unified_agent_configuration_service_configuration_sources_parser_timeout_in_milliseconds
				types = var.unified_agent_configuration_service_configuration_sources_parser_types
			}
			paths = var.unified_agent_configuration_service_configuration_sources_paths
		}
	}

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
	group_association {

		#Optional
		group_list = var.unified_agent_configuration_group_association_group_list
	}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The OCID of the compartment that the resource belongs to.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Required) (Updatable) Description for this resource.
* `display_name` - (Required) (Updatable) The user-friendly display name. This must be unique within the enclosing resource, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `group_association` - (Optional) (Updatable) Groups using the configuration.
	* `group_list` - (Optional) (Updatable) list of group/dynamic group ids associated with this configuration.
* `is_enabled` - (Required) (Updatable) Whether or not this resource is currently enabled.
* `service_configuration` - (Required) (Updatable) Top level Unified Agent service configuration object.
	* `configuration_type` - (Required) (Updatable) Type of Unified Agent service configuration.
	* `destination` - (Required) (Updatable) Logging destination object.
		* `log_object_id` - (Required) (Updatable) The OCID of the resource.
		* `operational_metrics_configuration` - (Optional) (Updatable) Unified monitoring agent operational metrics configuration object.
			* `destination` - (Required) (Updatable) Unified monitoring agent operational metrics destination object.
				* `compartment_id` - (Required) (Updatable) The OCID of the compartment that the resource belongs to.
			* `source` - (Required) (Updatable) Unified monitoring agent operational metrics source object.
				* `metrics` - (Optional) (Updatable) List of unified monitoring agent operational metrics.
				* `record_input` - (Required) (Updatable) Record section of OperationalMetricsSource object.
					* `namespace` - (Required) (Updatable) Namespace to emit the operational metrics.
					* `resource_group` - (Optional) (Updatable) Resource group to emit the operational metrics.
				* `type` - (Required) (Updatable) Type of the unified monitoring agent operational metrics source object.
	* `sources` - (Required) (Updatable) Logging source object.
		* `channels` - (Required when source_type=WINDOWS_EVENT_LOG) (Updatable) Windows event log channels.
		* `name` - (Required when configuration_type=LOGGING) (Updatable) Unique name for the source.
		* `parser` - (Applicable when source_type=LOG_TAIL) (Updatable) Source parser object.
			* `delimiter` - (Applicable when parser_type=CSV | TSV) (Updatable) CSV delimiter.
			* `expression` - (Required when parser_type=REGEXP) (Updatable) Regex pattern.
			* `field_time_key` - (Applicable when source_type=LOG_TAIL) (Updatable) Specifies the time field for the event time. If the event doesn't have this field, the current time is used.
			* `format` - (Required when parser_type=MULTILINE) (Updatable) Mutiline pattern format.
			* `format_firstline` - (Applicable when parser_type=MULTILINE) (Updatable) First line pattern format.
			* `grok_failure_key` - (Applicable when parser_type=GROK | MULTILINE_GROK) (Updatable) Grok failure key.
			* `grok_name_key` - (Applicable when parser_type=GROK | MULTILINE_GROK) (Updatable) Grok name key.
			* `is_estimate_current_event` - (Applicable when source_type=LOG_TAIL) (Updatable) If true, use Fluent::EventTime.now(current time) as a timestamp when the time_key is specified.
			* `is_keep_time_key` - (Applicable when source_type=LOG_TAIL) (Updatable) If true, keep the time field in the record.
			* `is_merge_cri_fields` - (Applicable when parser_type=CRI) (Updatable) If you don't need stream or logtag fields, set this to false.
			* `is_null_empty_string` - (Applicable when source_type=LOG_TAIL) (Updatable) If true, an empty string field is replaced with a null value.
			* `is_support_colonless_ident` - (Applicable when parser_type=SYSLOG) (Updatable) Specifies whether or not to support colonless ident. Corresponds to the Fluentd support_colonless_ident parameter.
			* `is_with_priority` - (Applicable when parser_type=SYSLOG) (Updatable) Specifies with priority or not. Corresponds to the Fluentd with_priority parameter.
			* `keys` - (Required when parser_type=CSV | TSV) (Updatable) CSV keys.
			* `message_format` - (Applicable when parser_type=SYSLOG) (Updatable) Syslog message format.
			* `message_key` - (Applicable when parser_type=NONE) (Updatable) Specifies the field name to contain logs.
			* `multi_line_start_regexp` - (Applicable when parser_type=MULTILINE_GROK) (Updatable) Multiline start regexp pattern.
			* `nested_parser` - (Applicable when parser_type=CRI) (Updatable) Optional nested JSON Parser for CRI. Supported fields are fieldTimeKey, timeFormat, and isKeepTimeKey.
				* `time_format` - (Applicable when parser_type=CRI) (Updatable) Process time value using the specified format.
				* `time_type` - (Applicable when parser_type=CRI) (Updatable) JSON parser time type.
			* `null_value_pattern` - (Applicable when source_type=LOG_TAIL) (Updatable) Specify the null value pattern.
			* `parser_type` - (Required) (Updatable) Type of fluent parser.
			* `patterns` - (Required when parser_type=GROK | MULTILINE_GROK) (Updatable) Grok pattern object.
				* `field_time_format` - (Applicable when parser_type=GROK | MULTILINE_GROK) (Updatable) Process value using the specified format. This is available only when time_type is a string.
				* `field_time_key` - (Applicable when parser_type=GROK | MULTILINE_GROK) (Updatable) Specify the time field for the event time. If the event doesn't have this field, the current time is used.
				* `field_time_zone` - (Applicable when parser_type=GROK | MULTILINE_GROK) (Updatable) Use the specified time zone. The time value can be parsed or formatted in the specified time zone.
				* `name` - (Applicable when parser_type=GROK | MULTILINE_GROK) (Updatable) The name key to tag this Grok pattern.
				* `pattern` - (Required when parser_type=GROK | MULTILINE_GROK) (Updatable) The Grok pattern.
			* `rfc5424time_format` - (Applicable when parser_type=SYSLOG) (Updatable) RFC 5424 time format.
			* `syslog_parser_type` - (Applicable when parser_type=SYSLOG) (Updatable) Syslog parser type.
			* `time_format` - (Applicable when parser_type=JSON | REGEXP | SYSLOG) (Updatable) Process time value using the specified format.
			* `time_type` - (Applicable when parser_type=JSON) (Updatable) JSON parser time type.
			* `timeout_in_milliseconds` - (Applicable when source_type=LOG_TAIL) (Updatable) Specify the timeout for parse processing. This is mainly for detecting an incorrect regexp pattern.
			* `types` - (Applicable when source_type=LOG_TAIL) (Updatable) Specify types for converting a field into another type. For example, With this configuration: <parse> @type csv keys time,host,req_id,user time_key time </parse>

				This incoming event: "2013/02/28 12:00:00,192.168.0.1,111,-"

				is parsed as: 1362020400 (2013/02/28/ 12:00:00)

				record: { "host"   : "192.168.0.1", "req_id" : "111", "user"   : "-" } 
		* `paths` - (Required when source_type=LOG_TAIL) (Updatable) Absolute paths for log source files. Wildcards can be used.
		* `source_type` - (Required) (Updatable) Unified schema logging source type.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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
		* `channels` - Windows event log channels.
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
				* `time_format` - Process time value using the specified format.
				* `time_type` - JSON parser time type.
			* `null_value_pattern` - Specify the null value pattern.
			* `parser_type` - Type of fluent parser.
			* `patterns` - Grok pattern object.
				* `field_time_format` - Process value using the specified format. This is available only when time_type is a string.
				* `field_time_key` - Specify the time field for the event time. If the event doesn't have this field, the current time is used.
				* `field_time_zone` - Use the specified time zone. The time value can be parsed or formatted in the specified time zone.
				* `name` - The name key to tag this Grok pattern.
				* `pattern` - The Grok pattern.
			* `rfc5424time_format` - RFC 5424 time format.
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
* `state` - The pipeline state.
* `time_created` - Time the resource was created.
* `time_last_modified` - Time the resource was last modified.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Unified Agent Configuration
	* `update` - (Defaults to 20 minutes), when updating the Unified Agent Configuration
	* `delete` - (Defaults to 20 minutes), when destroying the Unified Agent Configuration


## Import

UnifiedAgentConfigurations can be imported using the `id`, e.g.

```
$ terraform import oci_logging_unified_agent_configuration.test_unified_agent_configuration "id"
```

