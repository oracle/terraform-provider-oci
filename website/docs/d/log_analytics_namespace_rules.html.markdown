---
subcategory: "Log Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_log_analytics_namespace_rules"
sidebar_current: "docs-oci-datasource-log_analytics-namespace_rules"
description: |-
  Provides the list of Namespace Rules in Oracle Cloud Infrastructure Log Analytics service
---

# Data Source: oci_log_analytics_namespace_rules
This data source provides the list of Namespace Rules in Oracle Cloud Infrastructure Log Analytics service.

Returns a list of ingest time rules and scheduled tasks in a compartment. You may limit the number of items returned, provide sorting options, and filter the results.


## Example Usage

```hcl
data "oci_log_analytics_namespace_rules" "test_namespace_rules" {
	#Required
	compartment_id = var.compartment_id
	namespace = var.namespace_rule_namespace

	#Optional
	display_name = var.namespace_rule_display_name
	kind = var.namespace_rule_kind
	state = var.namespace_rule_state
	target_service = var.namespace_rule_target_service
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return rules whose displayName matches in whole or in part the specified value. The match is case-insensitive. 
* `kind` - (Optional) The rule kind used for filtering. Only rules of the specified kind will be returned. 
* `namespace` - (Required) The Logging Analytics namespace used for the request. 
* `state` - (Optional) The rule lifecycle state used for filtering. Currently supported values are ACTIVE and DELETED. 
* `target_service` - (Optional) The target service to use for filtering. 


## Attributes Reference

The following attributes are exported:

* `rule_summary_collection` - The list of rule_summary_collection.

### NamespaceRule Reference

The following attributes are exported:

* `items` - An array of rule summary objects.
	* `compartment_id` - Compartment Identifier [OCID] (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
	* `description` - Description for this resource. 
	* `display_name` - The ingest time rule or scheduled task display name.
	* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
	* `id` - The log analytics entity OCID. This ID is a reference used by log analytics features and it represents a resource that is provisioned and managed by the customer on their premises or on the cloud. 
	* `is_enabled` - A flag indicating whether or not the ingest time rule or scheduled task is enabled.
	* `kind` - The kind of rule - either an ingest time rule or a scheduled task. 
	* `last_execution_status` - The most recent task execution status.
	* `state` - The current state of the logging analytics rule. 
	* `target_service` - The target service.
	* `time_created` - The date and time the resource was created, in the format defined by RFC3339. 
	* `time_last_executed` - The date and time the scheduled task last executed, in the format defined by RFC3339.
	* `time_updated` - The date and time the resource was last updated, in the format defined by RFC3339. 

