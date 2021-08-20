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

Get the configuration of the item identified by the OCID.

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

* `apm_domain_id` - (Required) The APM Domain Id the request is intended for. 
* `config_id` - (Required) The OCID of the ConfiguredItem.


## Attributes Reference

The following attributes are exported:

* `config_type` - The type of configuration item
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - An optional string that describes what the filter is intended or used for.
* `dimensions` - A list of dimensions for this metric
	* `name` - The dimension name
	* `value_source` - The source to populate the dimension. Must be NULL at the moment. 
* `display_name` - The name by which this rule set can be displayed to the user.
* `filter_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a Span Filter. The filterId is mandatory for the creation of MetricGroups. A filterId will be generated when a Span Filter is created. 
* `filter_text` - The string that defines the Span Filter expression. 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the configuration item. An OCID will be generated when the item is created. 
* `metrics` - 
	* `description` - A description of the metric
	* `name` - The name of the metric
	* `unit` - The unit of the metric
	* `value_source` - Must be NULL at the moment, and "name" must be a known metric.
* `namespace` - The namespace to write the metrics to
* `rules` - 
	* `display_name` - A user-friendly name that provides a short description this rule.
	* `filter_text` - The string that defines the Span Filter expression. 
	* `is_apply_to_error_spans` - If true, the rule will compute the actual Apdex score for spans that have been marked as errors. If false, the rule will always set the Apdex for error spans to frustrating, regardless of the configured thresholds. Default is false. 
	* `is_enabled` - Specifies if the Apdex rule will be computed for spans matching the rule. Can be used to make sure certain spans don't get an Apdex score. The default is "true". 
	* `priority` - The priority controls the order in which multiple rules in a rule set are applied. Lower values indicate higher priorities. Rules with higher priority are applied first, and once a match is found, the rest of the rules are ignored. Rules within the same rule set cannot have the same priority. 
	* `satisfied_response_time` - The maximum response time in milliseconds that will be considered satisfactory for the end user. 
	* `tolerating_response_time` - The maximum response time in milliseconds that will be considered tolerable for the end user. Response times beyond this threshold will be considered frustrating. This value cannot be lower than "satisfiedResponseTime". 
* `time_created` - The time the resource was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2020-02-12T22:47:12.613Z` 
* `time_updated` - The time the resource was updated, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2020-02-13T22:47:12.613Z` 

