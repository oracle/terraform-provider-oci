---
subcategory: "Apm Config"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_apm_config_config"
sidebar_current: "docs-oci-datasource-apm_config-config"
description: |-
  Provides details about a specific Config in Oracle Cloud Infrastructure Apm Config service
---

# Data Source: oci_apm_config_config
This data source provides details about a specific Config resource in Oracle Cloud Infrastructure Apm Config service.

Gets the configuration item identified by the OCID.

## Example Usage

```hcl
data "oci_apm_config_config" "test_config" {
	#Required
	apm_domain_id = oci_apm_apm_domain.test_apm_domain.id
	config_id = oci_apm_config_config.test_config.id
}
```

## Argument Reference

The following arguments are supported:

* `apm_domain_id` - (Required) The APM Domain ID the request is intended for. 
* `config_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the configuration item. 


## Attributes Reference

The following attributes are exported:

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
* `metrics` - The list of metrics in this group. 
	* `description` - A description of the metric.
	* `name` - The name of the metric. This must be a known metric name.
	* `unit` - The unit of the metric.
	* `value_source` - This must not be set.
* `namespace` - The namespace to which the metrics are published. It must be one of several predefined namespaces. 
* `options` - The options are stored here as JSON.
* `rules` - 
	* `display_name` - The name by which a configuration entity is displayed to the end user.
	* `filter_text` - The string that defines the Span Filter expression. 
	* `is_apply_to_error_spans` - Specifies whether an Apdex score should be computed for error spans. Setting it to "true" means that the Apdex score is computed in the usual way. Setting it to "false" skips the Apdex computation and sets the Apdex score to "frustrating" regardless of the configured thresholds. The default is "false". 
	* `is_enabled` - Specifies whether the Apdex score should be computed for spans matching the rule. This can be used to disable Apdex score for spans that do not need or require it. The default is "true". 
	* `priority` - The priority controls the order in which multiple rules in a rule set are applied. Lower values indicate higher priorities. Rules with higher priority are applied first, and once a match is found, the rest of the rules are ignored. Rules within the same rule set cannot have the same priority. 
	* `satisfied_response_time` - The maximum response time in milliseconds that is considered "satisfactory" for the end user. 
	* `tolerating_response_time` - The maximum response time in milliseconds that is considered "tolerable" for the end user. A response time beyond this threshold is considered "frustrating". This value cannot be lower than "satisfiedResponseTime". 
* `time_created` - The time the resource was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2020-02-12T22:47:12.613Z` 
* `time_updated` - The time the resource was updated, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2020-02-13T22:47:12.613Z` 
* `updated_by` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a user. 

