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

Creates a new Configuration item.

## Example Usage

```hcl
resource "oci_apm_config_config" "test_config" {
	#Required
	apm_domain_id = oci_apm_apm_domain.test_apm_domain.id
	config_type = var.config_config_type
	display_name = var.config_display_name

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.config_description
	dimensions {

		#Optional
		name = var.config_dimensions_name
		value_source = var.config_dimensions_value_source
	}
	filter_id = oci_apm_config_filter.test_filter.id
	filter_text = var.config_filter_text
	freeform_tags = {"bar-key"= "value"}
	metrics {

		#Optional
		description = var.config_metrics_description
		name = var.config_metrics_name
		unit = var.config_metrics_unit
		value_source = var.config_metrics_value_source
	}
	namespace = var.config_namespace
	opc_dry_run = var.config_opc_dry_run
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
}
```

## Argument Reference

The following arguments are supported:

* `apm_domain_id` - (Required) (Updatable) The APM Domain Id the request is intended for. 
* `config_type` - (Required) (Updatable) The type of configuration item
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Applicable when config_type=SPAN_FILTER) (Updatable) An optional string that describes what the filter is intended or used for.
* `dimensions` - (Applicable when config_type=METRIC_GROUP) (Updatable) A list of dimensions for this metric. Must be NULL at the moment.
	* `name` - (Required when config_type=METRIC_GROUP) (Updatable) The dimension name
	* `value_source` - (Applicable when config_type=METRIC_GROUP) (Updatable) The source to populate the dimension. Must be NULL at the moment. 
* `display_name` - (Required) (Updatable) The name by which this rule set can be displayed to the user.
* `filter_id` - (Required when config_type=METRIC_GROUP) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a Span Filter. The filterId is mandatory for the creation of MetricGroups. A filterId will be generated when a Span Filter is created. 
* `filter_text` - (Required when config_type=SPAN_FILTER) (Updatable) The string that defines the Span Filter expression. 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `metrics` - (Required when config_type=METRIC_GROUP) (Updatable) 
	* `description` - (Applicable when config_type=METRIC_GROUP) (Updatable) A description of the metric
	* `name` - (Required when config_type=METRIC_GROUP) (Updatable) The name of the metric
	* `unit` - (Applicable when config_type=METRIC_GROUP) (Updatable) The unit of the metric
	* `value_source` - (Applicable when config_type=METRIC_GROUP) (Updatable) Must be NULL at the moment, and "name" must be a known metric.
* `namespace` - (Applicable when config_type=METRIC_GROUP) (Updatable) The namespace to write the metrics to
* `opc_dry_run` - (Optional) (Updatable) Indicates that this request is a dry-run. If set to "true", nothing will be modified, only the validation will be performed. 
* `rules` - (Required when config_type=APDEX) (Updatable) 
	* `display_name` - (Applicable when config_type=APDEX) (Updatable) A user-friendly name that provides a short description this rule.
	* `filter_text` - (Required when config_type=APDEX) (Updatable) The string that defines the Span Filter expression. 
	* `is_apply_to_error_spans` - (Applicable when config_type=APDEX) (Updatable) If true, the rule will compute the actual Apdex score for spans that have been marked as errors. If false, the rule will always set the Apdex for error spans to frustrating, regardless of the configured thresholds. Default is false. 
	* `is_enabled` - (Applicable when config_type=APDEX) (Updatable) Specifies if the Apdex rule will be computed for spans matching the rule. Can be used to make sure certain spans don't get an Apdex score. The default is "true". 
	* `priority` - (Required when config_type=APDEX) (Updatable) The priority controls the order in which multiple rules in a rule set are applied. Lower values indicate higher priorities. Rules with higher priority are applied first, and once a match is found, the rest of the rules are ignored. Rules within the same rule set cannot have the same priority. 
	* `satisfied_response_time` - (Applicable when config_type=APDEX) (Updatable) The maximum response time in milliseconds that will be considered satisfactory for the end user. 
	* `tolerating_response_time` - (Applicable when config_type=APDEX) (Updatable) The maximum response time in milliseconds that will be considered tolerable for the end user. Response times beyond this threshold will be considered frustrating. This value cannot be lower than "satisfiedResponseTime". 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Config
	* `update` - (Defaults to 20 minutes), when updating the Config
	* `delete` - (Defaults to 20 minutes), when destroying the Config


## Import

Configs can be imported using the `id`, e.g.

```
$ terraform import oci_apm_config_config.test_config "configs/{configId}/apmDomainId/{apmDomainId}" 
```

