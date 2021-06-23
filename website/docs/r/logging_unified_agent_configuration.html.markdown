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
				is_null_empty_string = var.unified_agent_configuration_service_configuration_sources_parser_is_null_empty_string
				is_support_colonless_ident = var.unified_agent_configuration_service_configuration_sources_parser_is_support_colonless_ident
				is_with_priority = var.unified_agent_configuration_service_configuration_sources_parser_is_with_priority
				keys = var.unified_agent_configuration_service_configuration_sources_parser_keys
				message_format = var.unified_agent_configuration_service_configuration_sources_parser_message_format
				message_key = var.unified_agent_configuration_service_configuration_sources_parser_message_key
				multi_line_start_regexp = var.unified_agent_configuration_service_configuration_sources_parser_multi_line_start_regexp
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
	* `sources` - (Required) (Updatable) 
		* `channels` - (Applicable when source_type=WINDOWS_EVENT_LOG) (Updatable) 
		* `name` - (Required when configuration_type=LOGGING) (Updatable) unique name for the source
		* `parser` - (Applicable when source_type=LOG_TAIL) (Updatable) source parser object.
			* `delimiter` - (Applicable when parser_type=CSV | TSV) (Updatable) 
			* `expression` - (Applicable when parser_type=REGEXP) (Updatable) 
			* `field_time_key` - (Applicable when source_type=LOG_TAIL) (Updatable) Specify time field for the event time. If the event doesn't have this field, the current time is used.
			* `format` - (Applicable when parser_type=MULTILINE) (Updatable) 
			* `format_firstline` - (Applicable when parser_type=MULTILINE) (Updatable) 
			* `grok_failure_key` - (Applicable when parser_type=GROK | MULTILINE_GROK) (Updatable) 
			* `grok_name_key` - (Applicable when parser_type=GROK | MULTILINE_GROK) (Updatable) 
			* `is_estimate_current_event` - (Applicable when source_type=LOG_TAIL) (Updatable) If true, use Fluent::EventTime.now(current time) as a timestamp when time_key is specified.
			* `is_keep_time_key` - (Applicable when source_type=LOG_TAIL) (Updatable) If true, keep time field in the record.
			* `is_null_empty_string` - (Applicable when source_type=LOG_TAIL) (Updatable) If true, an empty string field is replaced with nil.
			* `is_support_colonless_ident` - (Applicable when parser_type=SYSLOG) (Updatable) 
			* `is_with_priority` - (Applicable when parser_type=SYSLOG) (Updatable) 
			* `keys` - (Applicable when parser_type=CSV | TSV) (Updatable) 
			* `message_format` - (Applicable when parser_type=SYSLOG) (Updatable) 
			* `message_key` - (Applicable when parser_type=NONE) (Updatable) 
			* `multi_line_start_regexp` - (Applicable when parser_type=MULTILINE_GROK) (Updatable) 
			* `null_value_pattern` - (Applicable when source_type=LOG_TAIL) (Updatable) Specify the null value pattern.
			* `parser_type` - (Required) (Updatable) Type of fluent parser.
			* `patterns` - (Applicable when parser_type=GROK | MULTILINE_GROK) (Updatable) 
				* `field_time_format` - (Applicable when parser_type=GROK | MULTILINE_GROK) (Updatable) Process value using the specified format. This is available only when time_type is a string.
				* `field_time_key` - (Applicable when parser_type=GROK | MULTILINE_GROK) (Updatable) Specify the time field for the event time. If the event doesn't have this field, the current time is used.
				* `field_time_zone` - (Applicable when parser_type=GROK | MULTILINE_GROK) (Updatable) Use the specified time zone. The time value can be parsed or formatted in the specified time zone.
				* `name` - (Applicable when parser_type=GROK | MULTILINE_GROK) (Updatable) The name key to tag this grok pattern.
				* `pattern` - (Required when parser_type=GROK | MULTILINE_GROK) (Updatable) The grok pattern.
			* `rfc5424time_format` - (Applicable when parser_type=SYSLOG) (Updatable) 
			* `syslog_parser_type` - (Applicable when parser_type=SYSLOG) (Updatable) 
			* `time_format` - (Applicable when parser_type=JSON | REGEXP | SYSLOG) (Updatable) 
			* `time_type` - (Applicable when parser_type=JSON) (Updatable) 
			* `timeout_in_milliseconds` - (Applicable when source_type=LOG_TAIL) (Updatable) Specify the timeout for parse processing. This is mainly for detecting an incorrect regexp pattern.
			* `types` - (Applicable when source_type=LOG_TAIL) (Updatable) Specify types for converting a field into another type.
		* `paths` - (Applicable when source_type=LOG_TAIL) (Updatable) 
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
	* `sources` - 
		* `channels` - 
		* `name` - unique name for the source
		* `parser` - source parser object.
			* `delimiter` - 
			* `expression` - 
			* `field_time_key` - Specify time field for the event time. If the event doesn't have this field, the current time is used.
			* `format` - 
			* `format_firstline` - 
			* `grok_failure_key` - 
			* `grok_name_key` - 
			* `is_estimate_current_event` - If true, use Fluent::EventTime.now(current time) as a timestamp when time_key is specified.
			* `is_keep_time_key` - If true, keep time field in the record.
			* `is_null_empty_string` - If true, an empty string field is replaced with nil.
			* `is_support_colonless_ident` - 
			* `is_with_priority` - 
			* `keys` - 
			* `message_format` - 
			* `message_key` - 
			* `multi_line_start_regexp` - 
			* `null_value_pattern` - Specify the null value pattern.
			* `parser_type` - Type of fluent parser.
			* `patterns` - 
				* `field_time_format` - Process value using the specified format. This is available only when time_type is a string.
				* `field_time_key` - Specify the time field for the event time. If the event doesn't have this field, the current time is used.
				* `field_time_zone` - Use the specified time zone. The time value can be parsed or formatted in the specified time zone.
				* `name` - The name key to tag this grok pattern.
				* `pattern` - The grok pattern.
			* `rfc5424time_format` - 
			* `syslog_parser_type` - 
			* `time_format` - 
			* `time_type` - 
			* `timeout_in_milliseconds` - Specify the timeout for parse processing. This is mainly for detecting an incorrect regexp pattern.
			* `types` - Specify types for converting a field into another type.
		* `paths` - 
		* `source_type` - Unified schema logging source type.
* `state` - The pipeline state.
* `time_created` - Time the resource was created.
* `time_last_modified` - Time the resource was last modified.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Unified Agent Configuration
	* `update` - (Defaults to 20 minutes), when updating the Unified Agent Configuration
	* `delete` - (Defaults to 20 minutes), when destroying the Unified Agent Configuration


## Import

UnifiedAgentConfigurations can be imported using the `id`, e.g.

```
$ terraform import oci_logging_unified_agent_configuration.test_unified_agent_configuration "id"
```

