---
subcategory: "Apm Config"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_apm_config_configs"
sidebar_current: "docs-oci-datasource-apm_config-configs"
description: |-
  Provides the list of Configs in Oracle Cloud Infrastructure Apm Config service
---

# Data Source: oci_apm_config_configs
This data source provides the list of Configs in Oracle Cloud Infrastructure Apm Config service.

Returns all configuration items, which can optionally be filtered by configuration type.

## Example Usage

```hcl
data "oci_apm_config_configs" "test_configs" {
	#Required
	apm_domain_id = oci_apm_apm_domain.test_apm_domain.id

	#Optional
	config_type = var.config_config_type
	defined_tag_equals = var.config_defined_tag_equals
	defined_tag_exists = var.config_defined_tag_exists
	display_name = var.config_display_name
	freeform_tag_equals = var.config_freeform_tag_equals
	freeform_tag_exists = var.config_freeform_tag_exists
	options_group = var.config_options_group
}
```

## Argument Reference

The following arguments are supported:

* `apm_domain_id` - (Required) The APM Domain ID the request is intended for. 
* `config_type` - (Optional) A filter to match configuration items of a given type. Supported values are SPAN_FILTER, METRIC_GROUP, APDEX, AGENT and MACS_APM_EXTENSION. 
* `defined_tag_equals` - (Optional) A list of tag filters to apply.  Only resources with a defined tag matching the value will be returned. Each item in the list has the format "{namespace}.{tagName}.{value}".  All inputs are case-insensitive. Multiple values for the same key (i.e. same namespace and tag name) are interpreted as "OR". Values for different keys (i.e. different namespaces, different tag names, or both) are interpreted as "AND". 
* `defined_tag_exists` - (Optional) A list of tag existence filters to apply.  Only resources for which the specified defined tags exist will be returned. Each item in the list has the format "{namespace}.{tagName}.true" (for checking existence of a defined tag) or "{namespace}.true".  All inputs are case-insensitive. Currently, only existence ("true" at the end) is supported. Absence ("false" at the end) is not supported. Multiple values for the same key (i.e. same namespace and tag name) are interpreted as "OR". Values for different keys (i.e. different namespaces, different tag names, or both) are interpreted as "AND". 
* `display_name` - (Applicable when config_type=APDEX | MACS_APM_EXTENSION | METRIC_GROUP | OPTIONS | SPAN_FILTER) A filter to return resources that match the given display name.
* `freeform_tag_equals` - (Optional) A list of tag filters to apply.  Only resources with a freeform tag matching the value will be returned. The key for each tag is "{tagName}.{value}".  All inputs are case-insensitive. Multiple values for the same tag name are interpreted as "OR".  Values for different tag names are interpreted as "AND". 
* `freeform_tag_exists` - (Optional) A list of tag existence filters to apply.  Only resources for which the specified freeform tags exist the value will be returned. The key for each tag is "{tagName}.true".  All inputs are case-insensitive. Currently, only existence ("true" at the end) is supported. Absence ("false" at the end) is not supported. Multiple values for different tag names are interpreted as "AND". 
* `options_group` - (Optional) A filter to return OPTIONS resources that match the given group.


## Attributes Reference

The following attributes are exported:

* `config_collection` - The list of config_collection.

### Config Reference

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

