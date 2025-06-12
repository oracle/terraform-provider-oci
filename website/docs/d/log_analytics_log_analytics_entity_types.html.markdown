---
subcategory: "Log Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_log_analytics_log_analytics_entity_types"
sidebar_current: "docs-oci-datasource-log_analytics-log_analytics_entity_types"
description: |-
  Provides the list of Log Analytics Entity Types in Oracle Cloud Infrastructure Log Analytics service
---

# Data Source: oci_log_analytics_log_analytics_entity_types
This data source provides the list of Log Analytics Entity Types in Oracle Cloud Infrastructure Log Analytics service.

Return a list of log analytics entity types.

## Example Usage

```hcl
data "oci_log_analytics_log_analytics_entity_types" "test_log_analytics_entity_types" {
	#Required
	namespace = var.log_analytics_entity_type_namespace

	#Optional
	cloud_type = var.log_analytics_entity_type_cloud_type
	name = var.log_analytics_entity_type_name
	name_contains = var.log_analytics_entity_type_name_contains
	state = var.log_analytics_entity_type_state
}
```

## Argument Reference

The following arguments are supported:

* `cloud_type` - (Optional) A filter to return CLOUD or NON_CLOUD entity types. 
* `name` - (Optional) A filter to return only log analytics entity types whose name matches the entire name given. The match is case-insensitive. 
* `name_contains` - (Optional) A filter to return only log analytics entity types whose name or internalName contains name given. The match is case-insensitive. 
* `namespace` - (Required) The Logging Analytics namespace used for the request. 
* `state` - (Optional) A filter to return only those log analytics entity types with the specified lifecycle state. The state value is case-insensitive.


## Attributes Reference

The following attributes are exported:

* `log_analytics_entity_type_collection` - The list of log_analytics_entity_type_collection.

### LogAnalyticsEntityType Reference

The following attributes are exported:

* `items` - Array of log analytics entity type summary.
	* `category` - Log analytics entity type category. Category will be used for grouping and filtering. 
	* `cloud_type` - Log analytics entity type group. This can be CLOUD (OCI) or NON_CLOUD otherwise. 
	* `internal_name` - Internal name for the log analytics entity type. 
	* `name` - Log analytics entity type name. 
	* `state` - The current lifecycle state of the log analytics entity type.
	* `time_created` - Time the log analytics entity type was created. An RFC3339 formatted datetime string. 
	* `time_updated` - Time the log analytics entity type was updated. An RFC3339 formatted datetime string. 

