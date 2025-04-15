---
subcategory: "Log Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_log_analytics_namespace_field_usage"
sidebar_current: "docs-oci-datasource-log_analytics-namespace_field_usage"
description: |-
  Provides details about a specific Namespace Field Usage in Oracle Cloud Infrastructure Log Analytics service
---

# Data Source: oci_log_analytics_namespace_field_usage
This data source provides details about a specific Namespace Field Usage resource in Oracle Cloud Infrastructure Log Analytics service.

Gets usage information about the field with the specified name.


## Example Usage

```hcl
data "oci_log_analytics_namespace_field_usage" "test_namespace_field_usage" {
	#Required
	field_name = var.namespace_field_usage_field_name
	namespace = var.namespace_field_usage_namespace
}
```

## Argument Reference

The following arguments are supported:

* `field_name` - (Required) The field name.
* `namespace` - (Required) The Logging Analytics namespace used for the request. 


## Attributes Reference

The following attributes are exported:

* `dependent_parsers` - Parsers that depend on or use the field.
	* `dependencies` - The list of dependencies of the parser.
		* `reference_display_name` - The display name of the dependency object
		* `reference_id` - The unique identifier of the reference, if available.
		* `reference_name` - The name of the dependency object
		* `reference_type` - The type of reference that defines the dependency.
		* `type` - The dependency type.
	* `is_system` - The system flag.  A value of false denotes a custom, or user defined object.  A value of true denotes a built in object. 
	* `parser_display_name` - The parser display name.
	* `parser_id` - The parser unique identifier.
	* `parser_name` - The parser name.
	* `parser_type` - The parser type
* `dependent_sources` - Sources that depend on or use the field.
	* `dependencies` - The list of dependencies defined by the source.
		* `reference_display_name` - The display name of the dependency object
		* `reference_id` - The unique identifier of the reference, if available.
		* `reference_name` - The name of the dependency object
		* `reference_type` - The type of reference that defines the dependency.
		* `type` - The dependency type.
	* `entity_types` - The entity types.
		* `entity_type` - The entity type.
		* `entity_type_category` - The type category.
		* `entity_type_display_name` - The entity type display name.
		* `source_id` - The source unique identifier.
	* `is_auto_association_enabled` - A flag indicating whether or not the source is marked for auto association. 
	* `is_system` - The system flag.  A value of false denotes a custom, or user defined object.  A value of true denotes a built in object. 
	* `source_display_name` - The source display name.
	* `source_id` - The source unique identifier.
	* `source_name` - The source name.
	* `source_type` - The source type.

