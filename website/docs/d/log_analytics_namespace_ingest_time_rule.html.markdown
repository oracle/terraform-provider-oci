---
subcategory: "Log Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_log_analytics_namespace_ingest_time_rule"
sidebar_current: "docs-oci-datasource-log_analytics-namespace_ingest_time_rule"
description: |-
  Provides details about a specific Namespace Ingest Time Rule in Oracle Cloud Infrastructure Log Analytics service
---

# Data Source: oci_log_analytics_namespace_ingest_time_rule
This data source provides details about a specific Namespace Ingest Time Rule resource in Oracle Cloud Infrastructure Log Analytics service.

Gets detailed information about the specified ingest time rule such as description, defined tags, and free-form tags.


## Example Usage

```hcl
data "oci_log_analytics_namespace_ingest_time_rule" "test_namespace_ingest_time_rule" {
	#Required
	ingest_time_rule_id = oci_events_rule.test_rule.id
	namespace = var.namespace_ingest_time_rule_namespace
}
```

## Argument Reference

The following arguments are supported:

* `ingest_time_rule_id` - (Required) Unique ocid of the ingest time rule. 
* `namespace` - (Required) The Logging Analytics namespace used for the request. 


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

