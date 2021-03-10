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

