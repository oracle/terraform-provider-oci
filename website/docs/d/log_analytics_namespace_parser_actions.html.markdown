---
subcategory: "Log Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_log_analytics_namespace_parser_actions"
sidebar_current: "docs-oci-datasource-log_analytics-namespace_parser_actions"
description: |-
  Provides the list of Namespace Parser Actions in Oracle Cloud Infrastructure Log Analytics service
---

# Data Source: oci_log_analytics_namespace_parser_actions
This data source provides the list of Namespace Parser Actions in Oracle Cloud Infrastructure Log Analytics service.

Returns a list of parser actions. You may limit the number of results and provide sorting order.


## Example Usage

```hcl
data "oci_log_analytics_namespace_parser_actions" "test_namespace_parser_actions" {
	#Required
	namespace = var.namespace_parser_action_namespace

	#Optional
	action_display_text = var.namespace_parser_action_action_display_text
	name = var.namespace_parser_action_name
}
```

## Argument Reference

The following arguments are supported:

* `action_display_text` - (Optional) The parser action display text used for filtering. 
* `name` - (Optional) The parser action name used for filtering. 
* `namespace` - (Required) The Logging Analytics namespace used for the request. 


## Attributes Reference

The following attributes are exported:

* `parser_action_summary_collection` - The list of parser_action_summary_collection.

### NamespaceParserAction Reference

The following attributes are exported:

* `items` - An array of parser action summary objects.
	* `description` - The parser action description.
	* `display_name` - The parser action display name.
	* `name` - The parser action name.

