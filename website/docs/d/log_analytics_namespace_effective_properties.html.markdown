---
subcategory: "Log Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_log_analytics_namespace_effective_properties"
sidebar_current: "docs-oci-datasource-log_analytics-namespace_effective_properties"
description: |-
  Provides the list of Namespace Effective Properties in Oracle Cloud Infrastructure Log Analytics service
---

# Data Source: oci_log_analytics_namespace_effective_properties
This data source provides the list of Namespace Effective Properties in Oracle Cloud Infrastructure Log Analytics service.

Returns a list of effective properties for the specified resource.


## Example Usage

```hcl
data "oci_log_analytics_namespace_effective_properties" "test_namespace_effective_properties" {
	#Required
	namespace = var.namespace_effective_property_namespace

	#Optional
	agent_id = oci_cloud_bridge_agent.test_agent.id
	entity_id = oci_log_analytics_log_analytics_entity.test_log_analytics_entity.id
	is_include_patterns = var.namespace_effective_property_is_include_patterns
	name = var.namespace_effective_property_name
	pattern_id = oci_log_analytics_pattern.test_pattern.id
	source_name = var.namespace_effective_property_source_name
}
```

## Argument Reference

The following arguments are supported:

* `agent_id` - (Optional) The agent ocid. 
* `entity_id` - (Optional) The entity ocid. 
* `is_include_patterns` - (Optional) The include pattern flag. 
* `name` - (Optional) The property name used for filtering. 
* `namespace` - (Required) The Logging Analytics namespace used for the request. 
* `pattern_id` - (Optional) The pattern id. 
* `source_name` - (Optional) The source name.


## Attributes Reference

The following attributes are exported:

* `effective_property_collection` - The list of effective_property_collection.

### NamespaceEffectiveProperty Reference

The following attributes are exported:

* `items` - A list of properties and their effective values.
	* `effective_level` - The level from which the effective value was determined.
	* `name` - The property name.
	* `patterns` - A list of pattern level override values for the property.
		* `effective_level` - The effective level of the property value.
		* `id` - The pattern id.
		* `value` - The value of the property.
	* `value` - The effective value of the property. This is determined by considering the value set at the most effective level. 

