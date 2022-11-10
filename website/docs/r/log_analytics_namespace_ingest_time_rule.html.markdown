---
subcategory: "Log Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_log_analytics_namespace_ingest_time_rule"
sidebar_current: "docs-oci-resource-log_analytics-namespace_ingest_time_rule"
description: |-
  Provides the Namespace Ingest Time Rule resource in Oracle Cloud Infrastructure Log Analytics service
---

# oci_log_analytics_namespace_ingest_time_rule
This resource provides the Namespace Ingest Time Rule resource in Oracle Cloud Infrastructure Log Analytics service.

Creates a new ingest time rule in the specified compartment. You may also specify optional information such as description, defined tags, and free-form tags.


## Example Usage

```hcl
resource "oci_log_analytics_namespace_ingest_time_rule" "test_namespace_ingest_time_rule" {
	#Required
	actions {
		#Required
		compartment_id = var.compartment_id
		metric_name = oci_monitoring_metric.test_metric.name
		namespace = var.namespace_ingest_time_rule_actions_namespace
		type = var.namespace_ingest_time_rule_actions_type

		#Optional
		dimensions = var.namespace_ingest_time_rule_actions_dimensions
		resource_group = var.namespace_ingest_time_rule_actions_resource_group
	}
	compartment_id = var.compartment_id
	conditions {
		#Required
		field_name = var.namespace_ingest_time_rule_conditions_field_name
		field_operator = var.namespace_ingest_time_rule_conditions_field_operator
		field_value = var.namespace_ingest_time_rule_conditions_field_value
		kind = var.namespace_ingest_time_rule_conditions_kind

		#Optional
		additional_conditions {
			#Required
			condition_field = var.namespace_ingest_time_rule_conditions_additional_conditions_condition_field
			condition_operator = var.namespace_ingest_time_rule_conditions_additional_conditions_condition_operator
			condition_value = var.namespace_ingest_time_rule_conditions_additional_conditions_condition_value
		}
	}
	display_name = var.namespace_ingest_time_rule_display_name
	namespace = var.namespace_ingest_time_rule_namespace

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.namespace_ingest_time_rule_description
	freeform_tags = {"bar-key"= "value"}
}
```

## Argument Reference

The following arguments are supported:

* `actions` - (Required) (Updatable) The action(s) to be performed if the ingest time rule condition(s) are satisfied. 
	* `compartment_id` - (Required) (Updatable) The compartment OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the extracted metric. 
	* `dimensions` - (Optional) (Updatable) Additional dimensions to publish for the extracted metric. A valid list contains the source field names whose values are to be published as dimensions. The source name itself is specified using a special macro SOURCE_NAME 
	* `metric_name` - (Required) (Updatable) The metric name of the extracted metric. A valid value starts with an alphabetical character and includes only alphanumeric characters, periods (.), underscores (_), hyphens (-), and dollar signs ($). 
	* `namespace` - (Required) (Updatable) The namespace of the extracted metric. A valid value starts with an alphabetical character and includes only alphanumeric characters and underscores (_). 
	* `resource_group` - (Optional) (Updatable) The resourceGroup of the extracted metric. A valid value starts with an alphabetical character and includes only alphanumeric characters, periods (.), underscores (_), hyphens (-), and dollar signs ($). 
	* `type` - (Required) (Updatable) Discriminator.
* `compartment_id` - (Required) (Updatable) Compartment Identifier [OCID] (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `conditions` - (Required) (Updatable) The condition(s) to evaluate for an ingest time rule.
	* `additional_conditions` - (Optional) (Updatable) Optional additional condition(s) to be evaluated.
		* `condition_field` - (Required) (Updatable) The additional field name to be evaluated.
		* `condition_operator` - (Required) (Updatable) The operator to be used for evaluating the additional field.
		* `condition_value` - (Required) (Updatable) The additional field value to be evaluated.
	* `field_name` - (Required) (Updatable) The field name to be evaluated.
	* `field_operator` - (Required) (Updatable) The operator to be used for evaluating the field.
	* `field_value` - (Required) (Updatable) The field value to be evaluated.
	* `kind` - (Required) (Updatable) Discriminator.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) Description for this resource. 
* `display_name` - (Required) (Updatable) The ingest time rule display name.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `namespace` - (Required) The Logging Analytics namespace used for the request. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `actions` - The action(s) to be performed if the ingest time rule condition(s) are satisfied. 
	* `compartment_id` - The compartment OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the extracted metric. 
	* `dimensions` - Additional dimensions to publish for the extracted metric. A valid list contains the source field names whose values are to be published as dimensions. The source name itself is specified using a special macro SOURCE_NAME 
	* `metric_name` - The metric name of the extracted metric. A valid value starts with an alphabetical character and includes only alphanumeric characters, periods (.), underscores (_), hyphens (-), and dollar signs ($). 
	* `namespace` - The namespace of the extracted metric. A valid value starts with an alphabetical character and includes only alphanumeric characters and underscores (_). 
	* `resource_group` - The resourceGroup of the extracted metric. A valid value starts with an alphabetical character and includes only alphanumeric characters, periods (.), underscores (_), hyphens (-), and dollar signs ($). 
	* `type` - Discriminator.
* `compartment_id` - Compartment Identifier [OCID] (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `conditions` - The condition(s) to evaluate for an ingest time rule.
	* `additional_conditions` - Optional additional condition(s) to be evaluated.
		* `condition_field` - The additional field name to be evaluated.
		* `condition_operator` - The operator to be used for evaluating the additional field.
		* `condition_value` - The additional field value to be evaluated.
	* `field_name` - The field name to be evaluated.
	* `field_operator` - The operator to be used for evaluating the field.
	* `field_value` - The field value to be evaluated.
	* `kind` - Discriminator.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - Description for this resource. 
* `display_name` - The ingest time rule display name.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The log analytics entity OCID. This ID is a reference used by log analytics features and it represents a resource that is provisioned and managed by the customer on their premises or on the cloud. 
* `is_enabled` - A flag indicating whether or not the ingest time rule is enabled.
* `state` - The current state of the ingest time rule. 
* `time_created` - The date and time the resource was created, in the format defined by RFC3339. 
* `time_updated` - The date and time the resource was last updated, in the format defined by RFC3339. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Namespace Ingest Time Rule
	* `update` - (Defaults to 20 minutes), when updating the Namespace Ingest Time Rule
	* `delete` - (Defaults to 20 minutes), when destroying the Namespace Ingest Time Rule


## Import

NamespaceIngestTimeRules can be imported using the `id`, e.g.

```
$ terraform import oci_log_analytics_namespace_ingest_time_rule.test_namespace_ingest_time_rule "namespaces/{namespaceName}/ingestTimeRules/{ingestTimeRuleId}" 
```

