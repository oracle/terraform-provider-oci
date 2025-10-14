---
subcategory: "Apm Config"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_apm_config_config"
sidebar_current: "docs-oci-resource-apm_config-config"
description: |-
  Provides the Config resource in Oracle Cloud Infrastructure Apm Config service
---

# oci_apm_config_config
This resource provides the Config resource in Oracle Cloud Infrastructure Apm Config service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/apm-config/latest/Config

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/apm/apm_config

Creates a new configuration item.

## Example Usage

```hcl
resource "oci_apm_config_config" "test_config" {
	#Required
	apm_domain_id = oci_apm_apm_domain.test_apm_domain.id
	config_type = var.config_config_type

	#Optional
	agent_version = var.config_agent_version
	attach_install_dir = var.config_attach_install_dir
	config {

		#Optional
		config_map {

			#Optional
			file_name = var.config_config_config_map_file_name
			body = var.config_config_config_map_body
			content_type = var.config_config_config_map_content_type
		}
	}
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.config_description
	dimensions {

		#Optional
		name = var.config_dimensions_name
		value_source = var.config_dimensions_value_source
	}
	display_name = var.config_display_name
	filter_id = oci_apm_config_filter.test_filter.id
	filter_text = var.config_filter_text
	freeform_tags = {"bar-key"= "value"}
	group = var.config_group
	management_agent_id = oci_management_agent_management_agent.test_management_agent.id
	match_agents_with_attribute_value = var.config_match_agents_with_attribute_value
	metrics {

		#Optional
		description = var.config_metrics_description
		name = var.config_metrics_name
		unit = var.config_metrics_unit
		value_source = var.config_metrics_value_source
	}
	namespace = var.config_namespace
	opc_dry_run = var.config_opc_dry_run
	options = var.config_options
	overrides {

		#Optional
		override_list {

			#Optional
			agent_filter = var.config_overrides_override_list_agent_filter
			override_map = var.config_overrides_override_list_override_map
		}
	}
	process_filter = var.config_process_filter
	rules {

		#Optional
		display_name = var.config_rules_display_name
		filter_text = var.config_rules_filter_text
		is_apply_to_error_spans = var.config_rules_is_apply_to_error_spans
		is_enabled = var.config_rules_is_enabled
		priority = var.config_rules_priority
		satisfied_response_time = var.config_rules_satisfied_response_time
		tolerating_response_time = var.config_rules_tolerating_response_time
	}
	run_as_user = var.config_run_as_user
	service_name = oci_announcements_service_service.test_service.name
}
```

## Argument Reference

The following arguments are supported:

* `agent_version` - (Required when config_type=MACS_APM_EXTENSION) (Updatable) The version of the referenced agent bundle.
* `apm_domain_id` - (Required) (Updatable) The APM Domain ID the request is intended for. 
* `attach_install_dir` - (Required when config_type=MACS_APM_EXTENSION) (Updatable) The directory owned by runAsUser.
* `config` - (Applicable when config_type=AGENT) (Updatable) Collection of agent configuration files. For agents that use a single configuration file, this SHOULD contain a single entry and the file name MAY be an empty string. For multiple entries, you should use multiple blocks of `config_map`. To apply a different configuration in a subset of the agents, put this block anywhere in the body of the configuration and edit <some variable> and <some content> {{ <some variable> | default <some content> }} Example: com.oracle.apm.agent.tracer.enable.jfr = {{ isJfrEnabled | default false }} Then, in the configuration's overrides, specify a different value for <some variable> along with the desired agent filter. Example: "agentFilter": "ApplicationType='Tomcat'" "overrideMap": { "isJfrEnabled": true } 
	* `config_map` - (Applicable when config_type=AGENT) (Updatable) Map of an agent configuration file.
		* `file_name` - (Applicable when config_type=AGENT) (Updatable) An agent configuration file name.
        * `body` - (Applicable when config_type=AGENT) (Updatable) The Base64 encoded agent configuration file.
		* `content_type` - (Applicable when config_type=AGENT) (Updatable) The MIME Content-Type that describes the content of the body field, for example, text/yaml or text/yaml; charset=utf-8
* `config_type` - (Required) (Updatable) The type of configuration item.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Applicable when config_type=OPTIONS | SPAN_FILTER) (Updatable) An optional string that describes what the options are intended or used for.
* `dimensions` - (Applicable when config_type=METRIC_GROUP) (Updatable) A list of dimensions for the metric. This variable should not be used.
	* `name` - (Required when config_type=METRIC_GROUP) (Updatable) The name of the dimension.
	* `value_source` - (Applicable when config_type=METRIC_GROUP) (Updatable) The source to populate the dimension. This must not be specified. 
* `display_name` - (Required when config_type=APDEX | MACS_APM_EXTENSION | METRIC_GROUP | OPTIONS | SPAN_FILTER) (Updatable) The name by which a configuration entity is displayed to the end user.
* `filter_id` - (Required when config_type=METRIC_GROUP) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a Span Filter. The filterId is mandatory for the creation of MetricGroups. A filterId is generated when a Span Filter is created. 
* `filter_text` - (Required when config_type=SPAN_FILTER) (Updatable) The string that defines the Span Filter expression. 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `group` - (Applicable when config_type=OPTIONS) (Updatable) A string that specifies the group that an OPTIONS item belongs to. 
* `management_agent_id` - (Required when config_type=MACS_APM_EXTENSION) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Management Agent that will provision the APM Agent. 
* `match_agents_with_attribute_value` - (Required when config_type=AGENT) The agent attribute VALUE by which an agent configuration is matched to an agent.  Each agent configuration object must specify a different value.  The attribute KEY corresponding to this VALUE is in the matchAgentsWithAttributeKey field. 
* `metrics` - (Required when config_type=METRIC_GROUP) (Updatable) The list of metrics in this group. 
	* `description` - (Applicable when config_type=METRIC_GROUP) (Updatable) A description of the metric.
	* `name` - (Required when config_type=METRIC_GROUP) (Updatable) The name of the metric. This must be a known metric name.
	* `unit` - (Applicable when config_type=METRIC_GROUP) (Updatable) The unit of the metric.
	* `value_source` - (Applicable when config_type=METRIC_GROUP) (Updatable) This must not be set.
* `namespace` - (Applicable when config_type=METRIC_GROUP) (Updatable) The namespace to which the metrics are published. It must be one of several predefined namespaces. 
* `opc_dry_run` - (Optional) (Updatable) Indicates that the request is a dry run, if set to "true". A dry run request does not modify the configuration item details and is used only to perform validation on the submitted data. 
* `options` - (Applicable when config_type=OPTIONS) (Updatable) The options are stored here as JSON.
* `overrides` - (Applicable when config_type=AGENT) (Updatable) Agent configuration overrides that should apply to a subset of the agents associated with an Agent Config object.
	* `override_list` - (Applicable when config_type=AGENT) (Updatable) 
		* `agent_filter` - (Applicable when config_type=AGENT) (Updatable) The string that defines the Agent Filter expression. 
		* `override_map` - (Applicable when config_type=AGENT) (Updatable) A map whose key is a substitution variable specified within the configuration's body. For example, if below was specified in the configuration's body {{ isJfrEnabled | default false }} Then a valid map key would be "isJfrEnabled". The value is typically different than the default specified in the configuration's body. Thus, in this example, the map entry could be "isJfrEnabled": true 
* `process_filter` - (Required when config_type=MACS_APM_EXTENSION) (Updatable) Filter patterns used to discover active Java processes for provisioning the APM Agent.
* `rules` - (Required when config_type=APDEX) (Updatable) 
	* `display_name` - (Applicable when config_type=APDEX) (Updatable) The name by which a configuration entity is displayed to the end user.
	* `filter_text` - (Required when config_type=APDEX) (Updatable) The string that defines the Span Filter expression. 
	* `is_apply_to_error_spans` - (Applicable when config_type=APDEX) (Updatable) Specifies whether an Apdex score should be computed for error spans. Setting it to "true" means that the Apdex score is computed in the usual way. Setting it to "false" skips the Apdex computation and sets the Apdex score to "frustrating" regardless of the configured thresholds. The default is "false". 
	* `is_enabled` - (Applicable when config_type=APDEX) (Updatable) Specifies whether the Apdex score should be computed for spans matching the rule. This can be used to disable Apdex score for spans that do not need or require it. The default is "true". 
	* `priority` - (Required when config_type=APDEX) (Updatable) The priority controls the order in which multiple rules in a rule set are applied. Lower values indicate higher priorities. Rules with higher priority are applied first, and once a match is found, the rest of the rules are ignored. Rules within the same rule set cannot have the same priority. 
	* `satisfied_response_time` - (Applicable when config_type=APDEX) (Updatable) The maximum response time in milliseconds that is considered "satisfactory" for the end user. 
	* `tolerating_response_time` - (Applicable when config_type=APDEX) (Updatable) The maximum response time in milliseconds that is considered "tolerable" for the end user. A response time beyond this threshold is considered "frustrating". This value cannot be lower than "satisfiedResponseTime". 
* `run_as_user` - (Required when config_type=MACS_APM_EXTENSION) (Updatable) The OS user that should be used to discover Java processes.
* `service_name` - (Required when config_type=MACS_APM_EXTENSION) (Updatable) The name of the service being monitored. This argument enables you to filter by service and view traces and other signals in the APM Explorer user interface. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `agent_version` - The version of the referenced agent bundle.
* `attach_install_dir` - The directory owned by runAsUser.
* `config` - Collection of agent configuration files. For agents that use a single configuration file, this SHOULD contain a single entry and the file name MAY be an empty string. For multiple entries, you should use multiple blocks of `config_map`. To apply a different configuration in a subset of the agents, put this block anywhere in the body of the configuration and edit <some variable> and <some content> {{ <some variable> | default <some content> }} Example: com.oracle.apm.agent.tracer.enable.jfr = {{ isJfrEnabled | default false }} Then, in the configuration's overrides, specify a different value for <some variable> along with the desired agent filter. Example: "agentFilter": "ApplicationType='Tomcat'" "overrideMap": { "isJfrEnabled": true }
	* `config_map` - Map of an agent configuration file.
		* `file_name` - An agent configuration file name.
		* `body` - The Base64 encoded agent configuration file.
		* `content_type` - The MIME Content-Type that describes the content of the body field, for example, text/yaml or text/yaml; charset=utf-8
* `config_type` - The type of configuration item.
* `created_by` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a user. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - An optional string that describes what the span filter is intended or used for.
* `dimensions` - A list of dimensions for the metric. This variable should not be used.
	* `name` - The name of the dimension.
	* `value_source` - The source to populate the dimension. This must not be specified. 
* `display_name` - The name by which a configuration entity is displayed to the end user.
* `etag` - For optimistic concurrency control. See `if-match`. 
* `filter_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a Span Filter. The filterId is mandatory for the creation of MetricGroups. A filterId is generated when a Span Filter is created. 
* `filter_text` - The string that defines the Span Filter expression. 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `group` - A string that specifies the group that an OPTIONS item belongs to. 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the configuration item. An OCID is generated when the item is created. 
* `in_use_by` - The list of configuration items that reference the span filter.
	* `config_type` - The type of configuration item.
	* `display_name` - The name by which a configuration entity is displayed to the end user.
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the configuration item. An OCID is generated when the item is created. 
	* `options_group` - A string that specifies the group that an OPTIONS item belongs to. 
* `management_agent_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Management Agent that will provision the APM Agent. 
* `match_agents_with_attribute_key` - The agent attribute KEY by which an Agent configuration is matched to an agent.  All agent configuration objects share the same key. It is [ServiceName, service.name] by default.  The attribute VALUE corresponding to this KEY is in the matchAgentsWithAttributeValue field. 
* `match_agents_with_attribute_value` - The agent attribute VALUE by which an agent configuration is matched to an agent.  Each agent configuration object must specify a different value.  The attribute KEY corresponding to this VALUE is in the matchAgentsWithAttributeKey field. 
* `metrics` - The list of metrics in this group. 
	* `description` - A description of the metric.
	* `name` - The name of the metric. This must be a known metric name.
	* `unit` - The unit of the metric.
	* `value_source` - This must not be set.
* `namespace` - The namespace to which the metrics are published. It must be one of several predefined namespaces. 
* `options` - The options are stored here as JSON.
* `overrides` - Agent configuration overrides that should apply to a subset of the agents associated with an Agent Config object.
	* `override_list` - 
		* `agent_filter` - The string that defines the Agent Filter expression. 
		* `override_map` - A map whose key is a substitution variable specified within the configuration's body. For example, if below was specified in the configuration's body {{ isJfrEnabled | default false }} Then a valid map key would be "isJfrEnabled". The value is typically different than the default specified in the configuration's body. Thus, in this example, the map entry could be "isJfrEnabled": true 
* `process_filter` - Filter patterns used to discover active Java processes for provisioning the APM Agent.
* `rules` - 
	* `display_name` - The name by which a configuration entity is displayed to the end user.
	* `filter_text` - The string that defines the Span Filter expression. 
	* `is_apply_to_error_spans` - Specifies whether an Apdex score should be computed for error spans. Setting it to "true" means that the Apdex score is computed in the usual way. Setting it to "false" skips the Apdex computation and sets the Apdex score to "frustrating" regardless of the configured thresholds. The default is "false". 
	* `is_enabled` - Specifies whether the Apdex score should be computed for spans matching the rule. This can be used to disable Apdex score for spans that do not need or require it. The default is "true". 
	* `priority` - The priority controls the order in which multiple rules in a rule set are applied. Lower values indicate higher priorities. Rules with higher priority are applied first, and once a match is found, the rest of the rules are ignored. Rules within the same rule set cannot have the same priority. 
	* `satisfied_response_time` - The maximum response time in milliseconds that is considered "satisfactory" for the end user. 
	* `tolerating_response_time` - The maximum response time in milliseconds that is considered "tolerable" for the end user. A response time beyond this threshold is considered "frustrating". This value cannot be lower than "satisfiedResponseTime". 
* `run_as_user` - The OS user that should be used to discover Java processes.
* `service_name` - The name of the service being monitored. This argument enables you to filter by service and view traces and other signals in the APM Explorer user interface. 
* `time_created` - The time the resource was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2020-02-12T22:47:12.613Z` 
* `time_updated` - The time the resource was updated, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2020-02-13T22:47:12.613Z` 
* `updated_by` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a user. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Config
	* `update` - (Defaults to 20 minutes), when updating the Config
	* `delete` - (Defaults to 20 minutes), when destroying the Config


## Import

Configs can be imported using the `id`, e.g.

```
$ terraform import oci_apm_config_config.test_config "configs/{configId}/apmDomainId/{apmDomainId}" 
```

