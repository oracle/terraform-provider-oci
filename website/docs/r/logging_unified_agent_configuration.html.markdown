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
	description = var.unified_agent_configuration_description
	display_name = var.unified_agent_configuration_display_name
	is_enabled = var.unified_agent_configuration_is_enabled
    description = var.unified_agent_configuration_description
	display_name = var.unified_agent_configuration_display_name
	service_configuration {
		#Required
		configuration_type = var.unified_agent_configuration_service_configuration_configuration_type

		#Optional
		application_configurations {
			#Required
			destination {

				#Optional
				compartment_id = var.compartment_id
				metrics_namespace = var.unified_agent_configuration_service_configuration_application_configurations_destination_metrics_namespace
			}
			source_type = var.unified_agent_configuration_service_configuration_application_configurations_source_type

			#Optional
			source {

				#Optional
				name = var.unified_agent_configuration_service_configuration_application_configurations_source_name
				scrape_targets {

					#Optional
					k8s_namespace = var.unified_agent_configuration_service_configuration_application_configurations_source_scrape_targets_k8s_namespace
					name = var.unified_agent_configuration_service_configuration_application_configurations_source_scrape_targets_name
					resource_group = var.unified_agent_configuration_service_configuration_application_configurations_source_scrape_targets_resource_group
					resource_type = var.unified_agent_configuration_service_configuration_application_configurations_source_scrape_targets_resource_type
					service_name = oci_core_service.test_service.name
					url = var.unified_agent_configuration_service_configuration_application_configurations_source_scrape_targets_url
				}
			}
			sources {

				#Optional
				advanced_options {

					#Optional
					is_read_from_head = var.unified_agent_configuration_service_configuration_application_configurations_sources_advanced_options_is_read_from_head
				}
				name = var.unified_agent_configuration_service_configuration_application_configurations_sources_name
				parser {
					#Required
					parser_type = var.unified_agent_configuration_service_configuration_application_configurations_sources_parser_parser_type

					#Optional
					delimiter = var.unified_agent_configuration_service_configuration_application_configurations_sources_parser_delimiter
					expression = var.unified_agent_configuration_service_configuration_application_configurations_sources_parser_expression
					field_time_key = var.unified_agent_configuration_service_configuration_application_configurations_sources_parser_field_time_key
					format = var.unified_agent_configuration_service_configuration_application_configurations_sources_parser_format
					format_firstline = var.unified_agent_configuration_service_configuration_application_configurations_sources_parser_format_firstline
					grok_failure_key = var.unified_agent_configuration_service_configuration_application_configurations_sources_parser_grok_failure_key
					grok_name_key = var.unified_agent_configuration_service_configuration_application_configurations_sources_parser_grok_name_key
					is_estimate_current_event = var.unified_agent_configuration_service_configuration_application_configurations_sources_parser_is_estimate_current_event
					is_keep_time_key = var.unified_agent_configuration_service_configuration_application_configurations_sources_parser_is_keep_time_key
					is_merge_cri_fields = var.unified_agent_configuration_service_configuration_application_configurations_sources_parser_is_merge_cri_fields
					is_null_empty_string = var.unified_agent_configuration_service_configuration_application_configurations_sources_parser_is_null_empty_string
					is_support_colonless_ident = var.unified_agent_configuration_service_configuration_application_configurations_sources_parser_is_support_colonless_ident
					is_with_priority = var.unified_agent_configuration_service_configuration_application_configurations_sources_parser_is_with_priority
					keys = var.unified_agent_configuration_service_configuration_application_configurations_sources_parser_keys
					message_format = var.unified_agent_configuration_service_configuration_application_configurations_sources_parser_message_format
					message_key = var.unified_agent_configuration_service_configuration_application_configurations_sources_parser_message_key
					multi_line_start_regexp = var.unified_agent_configuration_service_configuration_application_configurations_sources_parser_multi_line_start_regexp
					nested_parser {

						#Optional
						parse_nested = var.unified_agent_configuration_service_configuration_application_configurations_sources_parser_nested_parser_parse_nested
						separator = var.unified_agent_configuration_service_configuration_application_configurations_sources_parser_nested_parser_separator
						time_format = var.unified_agent_configuration_service_configuration_application_configurations_sources_parser_nested_parser_time_format
						time_type = var.unified_agent_configuration_service_configuration_application_configurations_sources_parser_nested_parser_time_type
					}
					null_value_pattern = var.unified_agent_configuration_service_configuration_application_configurations_sources_parser_null_value_pattern
					parse_nested = var.unified_agent_configuration_service_configuration_application_configurations_sources_parser_parse_nested
					patterns {

						#Optional
						field_time_format = var.unified_agent_configuration_service_configuration_application_configurations_sources_parser_patterns_field_time_format
						field_time_key = var.unified_agent_configuration_service_configuration_application_configurations_sources_parser_patterns_field_time_key
						field_time_zone = var.unified_agent_configuration_service_configuration_application_configurations_sources_parser_patterns_field_time_zone
						name = var.unified_agent_configuration_service_configuration_application_configurations_sources_parser_patterns_name
						pattern = var.unified_agent_configuration_service_configuration_application_configurations_sources_parser_patterns_pattern
					}
					record_input {

						#Optional
						dimensions = var.unified_agent_configuration_service_configuration_application_configurations_sources_parser_record_input_dimensions
						namespace = var.unified_agent_configuration_service_configuration_application_configurations_sources_parser_record_input_namespace
						resource_group = var.unified_agent_configuration_service_configuration_application_configurations_sources_parser_record_input_resource_group
					}
					rfc5424time_format = var.unified_agent_configuration_service_configuration_application_configurations_sources_parser_rfc5424time_format
					separator = var.unified_agent_configuration_service_configuration_application_configurations_sources_parser_separator
					syslog_parser_type = var.unified_agent_configuration_service_configuration_application_configurations_sources_parser_syslog_parser_type
					time_format = var.unified_agent_configuration_service_configuration_application_configurations_sources_parser_time_format
					time_type = var.unified_agent_configuration_service_configuration_application_configurations_sources_parser_time_type
					timeout_in_milliseconds = var.unified_agent_configuration_service_configuration_application_configurations_sources_parser_timeout_in_milliseconds
					types = var.unified_agent_configuration_service_configuration_application_configurations_sources_parser_types
				}
				paths = var.unified_agent_configuration_service_configuration_application_configurations_sources_paths
				source_type = var.unified_agent_configuration_service_configuration_application_configurations_sources_source_type
			}
			unified_agent_configuration_filter {

				#Optional
				allow_list = var.unified_agent_configuration_service_configuration_application_configurations_unified_agent_configuration_filter_allow_list
				deny_list = var.unified_agent_configuration_service_configuration_application_configurations_unified_agent_configuration_filter_deny_list
				filter_type = var.unified_agent_configuration_service_configuration_application_configurations_unified_agent_configuration_filter_filter_type
				name = var.unified_agent_configuration_service_configuration_application_configurations_unified_agent_configuration_filter_name
			}
		}
		
		destination {

			#Optional
			log_object_id = oci_objectstorage_object.test_object.id
			operational_metrics_configuration {

				#Optional
				destination {

					#Optional
					compartment_id = var.compartment_id
				}
				source {

					#Optional
					metrics = var.unified_agent_configuration_service_configuration_destination_operational_metrics_configuration_source_metrics
					
					#Required
					record_input {

						#Optional
						namespace = var.unified_agent_configuration_service_configuration_destination_operational_metrics_configuration_source_record_input_namespace
						resource_group = var.unified_agent_configuration_service_configuration_destination_operational_metrics_configuration_source_record_input_resource_group
					}
					type = var.unified_agent_configuration_service_configuration_destination_operational_metrics_configuration_source_type
				}
			}
		}
		sources {
			#Required
			source_type = var.unified_agent_configuration_service_configuration_sources_source_type

			#Optional
			advanced_options {

				#Optional
				is_read_from_head = var.unified_agent_configuration_service_configuration_sources_advanced_options_is_read_from_head
			}
			channels = var.unified_agent_configuration_service_configuration_sources_channels
			custom_plugin = var.unified_agent_configuration_service_configuration_sources_custom_plugin
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
					parse_nested = var.unified_agent_configuration_service_configuration_sources_parser_nested_parser_parse_nested
					separator = var.unified_agent_configuration_service_configuration_sources_parser_nested_parser_separator
					time_format = var.unified_agent_configuration_service_configuration_sources_parser_nested_parser_time_format
					field_time_key = var.unified_agent_configuration_service_configuration_sources_parser_nested_parser_field_time_key
					is_keep_time_key = var.unified_agent_configuration_service_configuration_sources_parser_nested_parser_is_keep_time_key
				}
				null_value_pattern = var.unified_agent_configuration_service_configuration_sources_parser_null_value_pattern
				parse_nested = var.unified_agent_configuration_service_configuration_sources_parser_parse_nested
				patterns {

					#Optional
					field_time_format = var.unified_agent_configuration_service_configuration_sources_parser_patterns_field_time_format
					field_time_key = var.unified_agent_configuration_service_configuration_sources_parser_patterns_field_time_key
					field_time_zone = var.unified_agent_configuration_service_configuration_sources_parser_patterns_field_time_zone
					name = var.unified_agent_configuration_service_configuration_sources_parser_patterns_name
					pattern = var.unified_agent_configuration_service_configuration_sources_parser_patterns_pattern
				}
				record_input {

					#Optional
					dimensions = var.unified_agent_configuration_service_configuration_sources_parser_record_input_dimensions
					namespace = var.unified_agent_configuration_service_configuration_sources_parser_record_input_namespace
					resource_group = var.unified_agent_configuration_service_configuration_sources_parser_record_input_resource_group
				}
				rfc5424time_format = var.unified_agent_configuration_service_configuration_sources_parser_rfc5424time_format
				separator = var.unified_agent_configuration_service_configuration_sources_parser_separator
				syslog_parser_type = var.unified_agent_configuration_service_configuration_sources_parser_syslog_parser_type
				time_format = var.unified_agent_configuration_service_configuration_sources_parser_time_format
				time_type = var.unified_agent_configuration_service_configuration_sources_parser_time_type
				timeout_in_milliseconds = var.unified_agent_configuration_service_configuration_sources_parser_timeout_in_milliseconds
				types = var.unified_agent_configuration_service_configuration_sources_parser_types
			}
			paths = var.unified_agent_configuration_service_configuration_sources_paths
		}
		unified_agent_configuration_filter {
			#Required
			filter_type = var.unified_agent_configuration_service_configuration_unified_agent_configuration_filter_filter_type

			#Optional
			allow_list {

				#Optional
				key = var.unified_agent_configuration_service_configuration_unified_agent_configuration_filter_allow_list_key
				pattern = var.unified_agent_configuration_service_configuration_unified_agent_configuration_filter_allow_list_pattern
			}
			custom_filter_type = var.unified_agent_configuration_service_configuration_unified_agent_configuration_filter_custom_filter_type
			custom_sections {

				#Optional
				name = var.unified_agent_configuration_service_configuration_unified_agent_configuration_filter_custom_sections_name
				params = var.unified_agent_configuration_service_configuration_unified_agent_configuration_filter_custom_sections_params
			}
			deny_list {

				#Optional
				key = var.unified_agent_configuration_service_configuration_unified_agent_configuration_filter_deny_list_key
				pattern = var.unified_agent_configuration_service_configuration_unified_agent_configuration_filter_deny_list_pattern
			}
			emit_invalid_record_to_error = var.unified_agent_configuration_service_configuration_unified_agent_configuration_filter_emit_invalid_record_to_error
			hash_value_field = var.unified_agent_configuration_service_configuration_unified_agent_configuration_filter_hash_value_field
			inject_key_prefix = var.unified_agent_configuration_service_configuration_unified_agent_configuration_filter_inject_key_prefix
			is_auto_typecast_enabled = var.unified_agent_configuration_service_configuration_unified_agent_configuration_filter_is_auto_typecast_enabled
			is_renew_record_enabled = var.unified_agent_configuration_service_configuration_unified_agent_configuration_filter_is_renew_record_enabled
			is_ruby_enabled = var.unified_agent_configuration_service_configuration_unified_agent_configuration_filter_is_ruby_enabled
			keep_keys = var.unified_agent_configuration_service_configuration_unified_agent_configuration_filter_keep_keys
			key_name = oci_kms_key.test_key.name
			name = var.unified_agent_configuration_service_configuration_unified_agent_configuration_filter_name
			params = var.unified_agent_configuration_service_configuration_unified_agent_configuration_filter_params
			parser {
				#Required
				parser_type = var.unified_agent_configuration_service_configuration_unified_agent_configuration_filter_parser_parser_type

				#Optional
				delimiter = var.unified_agent_configuration_service_configuration_unified_agent_configuration_filter_parser_delimiter
				expression = var.unified_agent_configuration_service_configuration_unified_agent_configuration_filter_parser_expression
				field_time_key = var.unified_agent_configuration_service_configuration_unified_agent_configuration_filter_parser_field_time_key
				format = var.unified_agent_configuration_service_configuration_unified_agent_configuration_filter_parser_format
				format_firstline = var.unified_agent_configuration_service_configuration_unified_agent_configuration_filter_parser_format_firstline
				grok_failure_key = var.unified_agent_configuration_service_configuration_unified_agent_configuration_filter_parser_grok_failure_key
				grok_name_key = var.unified_agent_configuration_service_configuration_unified_agent_configuration_filter_parser_grok_name_key
				is_estimate_current_event = var.unified_agent_configuration_service_configuration_unified_agent_configuration_filter_parser_is_estimate_current_event
				is_keep_time_key = var.unified_agent_configuration_service_configuration_unified_agent_configuration_filter_parser_is_keep_time_key
				is_merge_cri_fields = var.unified_agent_configuration_service_configuration_unified_agent_configuration_filter_parser_is_merge_cri_fields
				is_null_empty_string = var.unified_agent_configuration_service_configuration_unified_agent_configuration_filter_parser_is_null_empty_string
				is_support_colonless_ident = var.unified_agent_configuration_service_configuration_unified_agent_configuration_filter_parser_is_support_colonless_ident
				is_with_priority = var.unified_agent_configuration_service_configuration_unified_agent_configuration_filter_parser_is_with_priority
				keys = var.unified_agent_configuration_service_configuration_unified_agent_configuration_filter_parser_keys
				message_format = var.unified_agent_configuration_service_configuration_unified_agent_configuration_filter_parser_message_format
				message_key = var.unified_agent_configuration_service_configuration_unified_agent_configuration_filter_parser_message_key
				multi_line_start_regexp = var.unified_agent_configuration_service_configuration_unified_agent_configuration_filter_parser_multi_line_start_regexp
				nested_parser {

					#Optional
					parse_nested = var.unified_agent_configuration_service_configuration_unified_agent_configuration_filter_parser_nested_parser_parse_nested
					separator = var.unified_agent_configuration_service_configuration_unified_agent_configuration_filter_parser_nested_parser_separator
					time_format = var.unified_agent_configuration_service_configuration_unified_agent_configuration_filter_parser_nested_parser_time_format
					time_type = var.unified_agent_configuration_service_configuration_unified_agent_configuration_filter_parser_nested_parser_time_type
				}
				null_value_pattern = var.unified_agent_configuration_service_configuration_unified_agent_configuration_filter_parser_null_value_pattern
				parse_nested = var.unified_agent_configuration_service_configuration_unified_agent_configuration_filter_parser_parse_nested
				patterns {

					#Optional
					field_time_format = var.unified_agent_configuration_service_configuration_unified_agent_configuration_filter_parser_patterns_field_time_format
					field_time_key = var.unified_agent_configuration_service_configuration_unified_agent_configuration_filter_parser_patterns_field_time_key
					field_time_zone = var.unified_agent_configuration_service_configuration_unified_agent_configuration_filter_parser_patterns_field_time_zone
					name = var.unified_agent_configuration_service_configuration_unified_agent_configuration_filter_parser_patterns_name
					pattern = var.unified_agent_configuration_service_configuration_unified_agent_configuration_filter_parser_patterns_pattern
				}
				record_input {

					#Optional
					dimensions = var.unified_agent_configuration_service_configuration_unified_agent_configuration_filter_parser_record_input_dimensions
					namespace = var.unified_agent_configuration_service_configuration_unified_agent_configuration_filter_parser_record_input_namespace
					resource_group = var.unified_agent_configuration_service_configuration_unified_agent_configuration_filter_parser_record_input_resource_group
				}
				rfc5424time_format = var.unified_agent_configuration_service_configuration_unified_agent_configuration_filter_parser_rfc5424time_format
				separator = var.unified_agent_configuration_service_configuration_unified_agent_configuration_filter_parser_separator
				syslog_parser_type = var.unified_agent_configuration_service_configuration_unified_agent_configuration_filter_parser_syslog_parser_type
				time_format = var.unified_agent_configuration_service_configuration_unified_agent_configuration_filter_parser_time_format
				time_type = var.unified_agent_configuration_service_configuration_unified_agent_configuration_filter_parser_time_type
				timeout_in_milliseconds = var.unified_agent_configuration_service_configuration_unified_agent_configuration_filter_parser_timeout_in_milliseconds
				types = var.unified_agent_configuration_service_configuration_unified_agent_configuration_filter_parser_types
			}
			record_list {

				#Optional
				key = var.unified_agent_configuration_service_configuration_unified_agent_configuration_filter_record_list_key
				value = var.unified_agent_configuration_service_configuration_unified_agent_configuration_filter_record_list_value
			}
			remove_key_name_field = var.unified_agent_configuration_service_configuration_unified_agent_configuration_filter_remove_key_name_field
			remove_keys = var.unified_agent_configuration_service_configuration_unified_agent_configuration_filter_remove_keys
			renew_time_key = var.unified_agent_configuration_service_configuration_unified_agent_configuration_filter_renew_time_key
			replace_invalid_sequence = var.unified_agent_configuration_service_configuration_unified_agent_configuration_filter_replace_invalid_sequence
			reserve_data = var.unified_agent_configuration_service_configuration_unified_agent_configuration_filter_reserve_data
			reserve_time = var.unified_agent_configuration_service_configuration_unified_agent_configuration_filter_reserve_time
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
	* `application_configurations` - (Required when configuration_type=MONITORING) (Updatable) Unified Agent monitoring application configuration details.
		* `destination` - (Required) (Updatable) Kubernetes destination object.
			* `compartment_id` - (Required when source_type=KUBERNETES | TAIL | URL) (Updatable) The OCID of the compartment that the resource belongs to.
			* `metrics_namespace` - (Required when source_type=KUBERNETES | TAIL | URL) (Updatable) Namespace to which metrics will be emitted.
		* `source` - (Required when source_type=KUBERNETES | URL) (Updatable) Kubernetes source object.
			* `name` - (Required when source_type=KUBERNETES | URL) (Updatable) Unique name for the source.
			* `scrape_targets` - (Required when source_type=KUBERNETES | URL) (Updatable) List of UnifiedAgentKubernetesScrapeTarget.
				* `k8s_namespace` - (Required when source_type=KUBERNETES) (Updatable) K8s namespace of the resource.
				* `name` - (Applicable when source_type=URL) (Updatable) Custom name.
				* `resource_group` - (Applicable when source_type=KUBERNETES) (Updatable) Resource group in Oracle Cloud Infrastructure monitoring.
				* `resource_type` - (Required when source_type=KUBERNETES) (Updatable) Type of resource to scrape metrics.
				* `service_name` - (Applicable when source_type=KUBERNETES) (Updatable) Name of the service prepended to the endpoints.
				* `url` - (Required when source_type=URL) (Updatable) URL from which the metrics are fetched.
		* `source_type` - (Required) (Updatable) Type of source of metrics
		* `sources` - (Required when source_type=TAIL) (Updatable) Tail log source objects.
			* `advanced_options` - (Applicable when source_type=TAIL) (Updatable) Advanced options for logging configuration
				* `is_read_from_head` - (Applicable when source_type=TAIL) (Updatable) Starts to read the logs from the head of the file or the last read position recorded in pos_file, not tail.
			* `name` - (Required when source_type=TAIL) (Updatable) Unique name for the source.
			* `parser` - (Applicable when source_type=TAIL) (Updatable) Source parser object.
				* `delimiter` - (Applicable when parser_type=CSV | TSV) (Updatable) CSV delimiter.
				* `expression` - (Required when parser_type=REGEXP) (Updatable) Regex pattern.
				* `field_time_key` - (Applicable when source_type=TAIL) (Updatable) Specifies the time field for the event time. If the event doesn't have this field, the current time is used.
				* `format` - (Required when parser_type=MULTILINE) (Updatable) Mutiline pattern format.
				* `format_firstline` - (Applicable when parser_type=MULTILINE) (Updatable) First line pattern format.
				* `grok_failure_key` - (Applicable when parser_type=GROK | MULTILINE_GROK) (Updatable) Grok failure key.
				* `grok_name_key` - (Applicable when parser_type=GROK | MULTILINE_GROK) (Updatable) Grok name key.
				* `is_estimate_current_event` - (Applicable when source_type=TAIL) (Updatable) If true, use Fluent::EventTime.now(current time) as a timestamp when the time_key is specified.
				* `is_keep_time_key` - (Applicable when source_type=TAIL) (Updatable) If true, keep the time field in the record.
				* `is_merge_cri_fields` - (Applicable when parser_type=CRI) (Updatable) If you don't need stream or logtag fields, set this to false.
				* `is_null_empty_string` - (Applicable when source_type=TAIL) (Updatable) If true, an empty string field is replaced with a null value.
				* `is_support_colonless_ident` - (Applicable when parser_type=SYSLOG) (Updatable) Specifies whether or not to support colonless ident. Corresponds to the Fluentd support_colonless_ident parameter.
				* `is_with_priority` - (Applicable when parser_type=SYSLOG) (Updatable) Specifies with priority or not. Corresponds to the Fluentd with_priority parameter.
				* `keys` - (Required when parser_type=CSV | TSV) (Updatable) CSV keys.
				* `message_format` - (Applicable when parser_type=SYSLOG) (Updatable) Syslog message format.
				* `message_key` - (Applicable when parser_type=NONE) (Updatable) Specifies the field name to contain logs.
				* `multi_line_start_regexp` - (Applicable when parser_type=MULTILINE_GROK) (Updatable) Multiline start regexp pattern.
				* `nested_parser` - (Applicable when parser_type=CRI) (Updatable) Optional nested JSON Parser for CRI. Supported fields are fieldTimeKey, timeFormat, and isKeepTimeKey.
					* `parse_nested` - (Applicable when parser_type=CRI) (Updatable) If true, a separator parameter can be further defined.
					* `separator` - (Applicable when parser_type=CRI) (Updatable) Keys of adjacent levels are joined by the separator.
					* `time_format` - (Applicable when parser_type=CRI) (Updatable) Process time value using the specified format.
					* `time_type` - (Applicable when parser_type=CRI) (Updatable) JSON parser time type.
				* `null_value_pattern` - (Applicable when source_type=TAIL) (Updatable) Specify the null value pattern.
				* `parse_nested` - (Applicable when parser_type=JSON) (Updatable) If true, a separator parameter can be further defined.
				* `parser_type` - (Required) (Updatable) Type of fluent parser.
				* `patterns` - (Required when parser_type=GROK | MULTILINE_GROK) (Updatable) Grok pattern object.
					* `field_time_format` - (Applicable when parser_type=GROK | MULTILINE_GROK) (Updatable) Process value using the specified format. This is available only when time_type is a string.
					* `field_time_key` - (Applicable when parser_type=GROK | MULTILINE_GROK) (Updatable) Specify the time field for the event time. If the event doesn't have this field, the current time is used.
					* `field_time_zone` - (Applicable when parser_type=GROK | MULTILINE_GROK) (Updatable) Use the specified time zone. The time value can be parsed or formatted in the specified time zone.
					* `name` - (Applicable when parser_type=GROK | MULTILINE_GROK) (Updatable) The name key to tag this Grok pattern.
					* `pattern` - (Required when parser_type=GROK | MULTILINE_GROK) (Updatable) The Grok pattern.
				* `record_input` - (Applicable when parser_type=OPENMETRICS) (Updatable) record section of openmetrics parser. 
					* `dimensions` - (Applicable when parser_type=OPENMETRICS) (Updatable) Dimensions to be added for metrics.
					* `namespace` - (Applicable when parser_type=OPENMETRICS) (Updatable) Namespace to emit metrics.
					* `resource_group` - (Applicable when parser_type=OPENMETRICS) (Updatable) Resource group to emit metrics.
				* `rfc5424time_format` - (Applicable when parser_type=SYSLOG) (Updatable) RFC 5424 time format.
				* `separator` - (Applicable when parser_type=JSON) (Updatable) Keys of adjacent levels are joined by the separator.
				* `syslog_parser_type` - (Applicable when parser_type=SYSLOG) (Updatable) Syslog parser type.
				* `time_format` - (Applicable when parser_type=JSON | REGEXP | SYSLOG) (Updatable) Process time value using the specified format.
				* `time_type` - (Applicable when parser_type=JSON) (Updatable) JSON parser time type.
				* `timeout_in_milliseconds` - (Applicable when source_type=TAIL) (Updatable) Specify the timeout for parse processing. This is mainly for detecting an incorrect regexp pattern.
				* `types` - (Applicable when source_type=TAIL) (Updatable) Specify types for converting a field into another type. For example, With this configuration: <parse> @type csv keys time,host,req_id,user time_key time </parse>

					This incoming event: "2013/02/28 12:00:00,192.168.0.1,111,-"

					is parsed as: 1362020400 (2013/02/28/ 12:00:00)

					record: { "host"   : "192.168.0.1", "req_id" : "111", "user"   : "-" } 
			* `paths` - (Required when source_type=TAIL) (Updatable) Absolute paths for log source files. Wildcards can be used.
			* `source_type` - (Required when source_type=TAIL) (Updatable) Unified schema logging source type.
		* `unified_agent_configuration_filter` - (Applicable when source_type=KUBERNETES | URL) (Updatable) Kubernetes filter object
			* `allow_list` - (Applicable when source_type=KUBERNETES | URL) (Updatable) List of metrics regex to be allowed.
			* `deny_list` - (Applicable when source_type=KUBERNETES | URL) (Updatable) List of metrics regex to be denied.
			* `filter_type` - (Required when source_type=KUBERNETES | URL) (Updatable) Unified schema logging filter type.
			* `name` - (Required when source_type=KUBERNETES | URL) (Updatable) Unique name for the filter.
	* `configuration_type` - (Required) (Updatable) Type of Unified Agent service configuration.
	* `destination` - (Required when configuration_type=LOGGING) (Updatable) Logging destination object.
		* `log_object_id` - (Required when configuration_type=LOGGING) (Updatable) The OCID of the resource.
		* `operational_metrics_configuration` - (Applicable when configuration_type=LOGGING) (Updatable) Unified monitoring agent operational metrics configuration object.
			* `destination` - (Required when configuration_type=LOGGING) (Updatable) Unified monitoring agent operational metrics destination object.
				* `compartment_id` - (Required when configuration_type=LOGGING) (Updatable) The OCID of the compartment that the resource belongs to.
			* `source` - (Required when configuration_type=LOGGING) (Updatable) Unified monitoring agent operational metrics source object.
				* `metrics` - (Applicable when configuration_type=LOGGING) (Updatable) List of unified monitoring agent operational metrics.
				* `record_input` - (Required when configuration_type=LOGGING) (Updatable) Record section of OperationalMetricsSource object.
					* `namespace` - (Required when configuration_type=LOGGING) (Updatable) Namespace to emit the operational metrics.
					* `resource_group` - (Applicable when configuration_type=LOGGING) (Updatable) Resource group to emit the operational metrics.
				* `type` - (Required when configuration_type=LOGGING) (Updatable) Type of the unified monitoring agent operational metrics source object.
	* `sources` - (Required when configuration_type=LOGGING) (Updatable) Logging source object.
		* `advanced_options` - (Applicable when source_type=LOG_TAIL) (Updatable) Advanced options for logging configuration
			* `is_read_from_head` - (Applicable when source_type=LOG_TAIL) (Updatable) Starts to read the logs from the head of the file or the last read position recorded in pos_file, not tail.
		* `channels` - (Required when source_type=WINDOWS_EVENT_LOG) (Updatable) Windows event log channels.
		* `custom_plugin` - (Required when source_type=CUSTOM_PLUGIN) (Updatable) User customized source plugin.
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
				* `parse_nested` - (Applicable when parser_type=CRI) (Updatable) If true, a separator parameter can be further defined.
				* `separator` - (Applicable when parser_type=CRI) (Updatable) Keys of adjacent levels are joined by the separator.
				* `time_format` - (Applicable when parser_type=CRI) (Updatable) Process time value using the specified format.
				* `time_type` - (Applicable when parser_type=CRI) (Updatable) JSON parser time type.
			* `null_value_pattern` - (Applicable when source_type=LOG_TAIL) (Updatable) Specify the null value pattern.
			* `parse_nested` - (Applicable when parser_type=JSON) (Updatable) If true, a separator parameter can be further defined.
			* `parser_type` - (Required) (Updatable) Type of fluent parser.
			* `patterns` - (Required when parser_type=GROK | MULTILINE_GROK) (Updatable) Grok pattern object.
				* `field_time_format` - (Applicable when parser_type=GROK | MULTILINE_GROK) (Updatable) Process value using the specified format. This is available only when time_type is a string.
				* `field_time_key` - (Applicable when parser_type=GROK | MULTILINE_GROK) (Updatable) Specify the time field for the event time. If the event doesn't have this field, the current time is used.
				* `field_time_zone` - (Applicable when parser_type=GROK | MULTILINE_GROK) (Updatable) Use the specified time zone. The time value can be parsed or formatted in the specified time zone.
				* `name` - (Applicable when parser_type=GROK | MULTILINE_GROK) (Updatable) The name key to tag this Grok pattern.
				* `pattern` - (Required when parser_type=GROK | MULTILINE_GROK) (Updatable) The Grok pattern.
			* `record_input` - (Applicable when parser_type=OPENMETRICS) (Updatable) record section of openmetrics parser. 
				* `dimensions` - (Applicable when parser_type=OPENMETRICS) (Updatable) Dimensions to be added for metrics.
				* `namespace` - (Applicable when parser_type=OPENMETRICS) (Updatable) Namespace to emit metrics.
				* `resource_group` - (Applicable when parser_type=OPENMETRICS) (Updatable) Resource group to emit metrics.
			* `rfc5424time_format` - (Applicable when parser_type=SYSLOG) (Updatable) RFC 5424 time format.
			* `separator` - (Applicable when parser_type=JSON) (Updatable) Keys of adjacent levels are joined by the separator.
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
	* `unified_agent_configuration_filter` - (Applicable when configuration_type=LOGGING) (Updatable) Logging filter object.
		* `allow_list` - (Applicable when filter_type=GREP_FILTER) (Updatable) A list of filtering rules to include logs
			* `key` - (Applicable when filter_type=GREP_FILTER) (Updatable) The field name to which the regular expression is applied
			* `pattern` - (Applicable when filter_type=GREP_FILTER) (Updatable) The regular expression
		* `custom_filter_type` - (Required when filter_type=CUSTOM_FILTER) (Updatable) Type of the custom filter
		* `custom_sections` - (Applicable when filter_type=CUSTOM_FILTER) (Updatable) List of custom sections in custom filter
			* `name` - (Applicable when filter_type=CUSTOM_FILTER) (Updatable) The name of the custom section
			* `params` - (Applicable when filter_type=CUSTOM_FILTER) (Updatable) Parameters in the custom section
		* `deny_list` - (Applicable when filter_type=GREP_FILTER) (Updatable) A list of filtering rules to reject logs
			* `key` - (Applicable when filter_type=GREP_FILTER) (Updatable) The field name to which the regular expression is applied
			* `pattern` - (Applicable when filter_type=GREP_FILTER) (Updatable) The regular expression
		* `emit_invalid_record_to_error` - (Applicable when filter_type=PARSER_FILTER) (Updatable) If true, emit invalid record to @ERROR label. Invalid cases are: 1) key does not exist; 2) the format does not match; or 3) an unexpected error. You can rescue unexpected format logs in the @ERROR lable. If you want to ignore these errors, set this to false. 
		* `filter_type` - (Required) (Updatable) Unified schema logging filter type.
		* `hash_value_field` - (Applicable when filter_type=PARSER_FILTER) (Updatable) Store the parsed values as a hash value in a field.
		* `inject_key_prefix` - (Applicable when filter_type=PARSER_FILTER) (Updatable) Store the parsed values with the specified key name prefix.
		* `is_auto_typecast_enabled` - (Applicable when filter_type=RECORD_TRANSFORMER_FILTER) (Updatable) If true, automatically casts the field types.
		* `is_renew_record_enabled` - (Applicable when filter_type=RECORD_TRANSFORMER_FILTER) (Updatable) If true, it modifies a new empty hash
		* `is_ruby_enabled` - (Applicable when filter_type=RECORD_TRANSFORMER_FILTER) (Updatable) When set to true, the full Ruby syntax is enabled in the ${} expression.
		* `keep_keys` - (Applicable when filter_type=RECORD_TRANSFORMER_FILTER) (Updatable) A list of keys to keep. Only relevant if isRenewRecordEnabled is set to true
		* `key_name` - (Required when filter_type=PARSER_FILTER) (Updatable) The field name in the record to parse.
		* `name` - (Required when configuration_type=LOGGING) (Updatable) Unique name for the filter.
		* `params` - (Applicable when filter_type=CUSTOM_FILTER) (Updatable) Parameters of the custom filter
		* `parser` - (Required when filter_type=PARSER_FILTER) (Updatable) Source parser object.
			* `delimiter` - (Applicable when parser_type=CSV | TSV) (Updatable) CSV delimiter.
			* `expression` - (Required when parser_type=REGEXP) (Updatable) Regex pattern.
			* `field_time_key` - (Applicable when filter_type=PARSER_FILTER) (Updatable) Specifies the time field for the event time. If the event doesn't have this field, the current time is used.
			* `format` - (Required when parser_type=MULTILINE) (Updatable) Mutiline pattern format.
			* `format_firstline` - (Applicable when parser_type=MULTILINE) (Updatable) First line pattern format.
			* `grok_failure_key` - (Applicable when parser_type=GROK | MULTILINE_GROK) (Updatable) Grok failure key.
			* `grok_name_key` - (Applicable when parser_type=GROK | MULTILINE_GROK) (Updatable) Grok name key.
			* `is_estimate_current_event` - (Applicable when filter_type=PARSER_FILTER) (Updatable) If true, use Fluent::EventTime.now(current time) as a timestamp when the time_key is specified.
			* `is_keep_time_key` - (Applicable when filter_type=PARSER_FILTER) (Updatable) If true, keep the time field in the record.
			* `is_merge_cri_fields` - (Applicable when parser_type=CRI) (Updatable) If you don't need stream or logtag fields, set this to false.
			* `is_null_empty_string` - (Applicable when filter_type=PARSER_FILTER) (Updatable) If true, an empty string field is replaced with a null value.
			* `is_support_colonless_ident` - (Applicable when parser_type=SYSLOG) (Updatable) Specifies whether or not to support colonless ident. Corresponds to the Fluentd support_colonless_ident parameter.
			* `is_with_priority` - (Applicable when parser_type=SYSLOG) (Updatable) Specifies with priority or not. Corresponds to the Fluentd with_priority parameter.
			* `keys` - (Required when parser_type=CSV | TSV) (Updatable) CSV keys.
			* `message_format` - (Applicable when parser_type=SYSLOG) (Updatable) Syslog message format.
			* `message_key` - (Applicable when parser_type=NONE) (Updatable) Specifies the field name to contain logs.
			* `multi_line_start_regexp` - (Applicable when parser_type=MULTILINE_GROK) (Updatable) Multiline start regexp pattern.
			* `nested_parser` - (Applicable when parser_type=CRI) (Updatable) Optional nested JSON Parser for CRI. Supported fields are fieldTimeKey, timeFormat, and isKeepTimeKey.
				* `parse_nested` - (Applicable when parser_type=CRI) (Updatable) If true, a separator parameter can be further defined.
				* `separator` - (Applicable when parser_type=CRI) (Updatable) Keys of adjacent levels are joined by the separator.
				* `time_format` - (Applicable when parser_type=CRI) (Updatable) Process time value using the specified format.
				* `time_type` - (Applicable when parser_type=CRI) (Updatable) JSON parser time type.
			* `null_value_pattern` - (Applicable when filter_type=PARSER_FILTER) (Updatable) Specify the null value pattern.
			* `parse_nested` - (Applicable when parser_type=JSON) (Updatable) If true, a separator parameter can be further defined.
			* `parser_type` - (Required) (Updatable) Type of fluent parser.
			* `patterns` - (Required when parser_type=GROK | MULTILINE_GROK) (Updatable) Grok pattern object.
				* `field_time_format` - (Applicable when parser_type=GROK | MULTILINE_GROK) (Updatable) Process value using the specified format. This is available only when time_type is a string.
				* `field_time_key` - (Applicable when parser_type=GROK | MULTILINE_GROK) (Updatable) Specify the time field for the event time. If the event doesn't have this field, the current time is used.
				* `field_time_zone` - (Applicable when parser_type=GROK | MULTILINE_GROK) (Updatable) Use the specified time zone. The time value can be parsed or formatted in the specified time zone.
				* `name` - (Applicable when parser_type=GROK | MULTILINE_GROK) (Updatable) The name key to tag this Grok pattern.
				* `pattern` - (Required when parser_type=GROK | MULTILINE_GROK) (Updatable) The Grok pattern.
			* `record_input` - (Applicable when parser_type=OPENMETRICS) (Updatable) record section of openmetrics parser. 
				* `dimensions` - (Applicable when parser_type=OPENMETRICS) (Updatable) Dimensions to be added for metrics.
				* `namespace` - (Applicable when parser_type=OPENMETRICS) (Updatable) Namespace to emit metrics.
				* `resource_group` - (Applicable when parser_type=OPENMETRICS) (Updatable) Resource group to emit metrics.
			* `rfc5424time_format` - (Applicable when parser_type=SYSLOG) (Updatable) RFC 5424 time format.
			* `separator` - (Applicable when parser_type=JSON) (Updatable) Keys of adjacent levels are joined by the separator.
			* `syslog_parser_type` - (Applicable when parser_type=SYSLOG) (Updatable) Syslog parser type.
			* `time_format` - (Applicable when parser_type=JSON | REGEXP | SYSLOG) (Updatable) Process time value using the specified format.
			* `time_type` - (Applicable when parser_type=JSON) (Updatable) JSON parser time type.
			* `timeout_in_milliseconds` - (Applicable when filter_type=PARSER_FILTER) (Updatable) Specify the timeout for parse processing. This is mainly for detecting an incorrect regexp pattern.
			* `types` - (Applicable when filter_type=PARSER_FILTER) (Updatable) Specify types for converting a field into another type. For example, With this configuration: <parse> @type csv keys time,host,req_id,user time_key time </parse>

				This incoming event: "2013/02/28 12:00:00,192.168.0.1,111,-"

				is parsed as: 1362020400 (2013/02/28/ 12:00:00)

				record: { "host"   : "192.168.0.1", "req_id" : "111", "user"   : "-" } 
		* `record_list` - (Required when filter_type=RECORD_TRANSFORMER_FILTER) (Updatable) Add new key-value pairs in logs
			* `key` - (Applicable when filter_type=RECORD_TRANSFORMER_FILTER) (Updatable) A new key
			* `value` - (Applicable when filter_type=RECORD_TRANSFORMER_FILTER) (Updatable) A new value
		* `remove_key_name_field` - (Applicable when filter_type=PARSER_FILTER) (Updatable) If true, remove the keyName field when parsing is succeeded.
		* `remove_keys` - (Applicable when filter_type=RECORD_TRANSFORMER_FILTER) (Updatable) A list of keys to delete
		* `renew_time_key` - (Applicable when filter_type=RECORD_TRANSFORMER_FILTER) (Updatable) Overwrites the time of logs with this value, this value must be a Unix timestamp.
		* `replace_invalid_sequence` - (Applicable when filter_type=PARSER_FILTER) (Updatable) If true, the invalid string is replaced with safe characters and is re-parsed.
		* `reserve_data` - (Applicable when filter_type=PARSER_FILTER) (Updatable) If true, keep the original key-value pair in the parsed result.
		* `reserve_time` - (Applicable when filter_type=PARSER_FILTER) (Updatable) If true, keep the original event time in the parsed result.


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

