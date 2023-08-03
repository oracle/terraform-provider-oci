---
subcategory: "Log Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_log_analytics_namespace_properties_metadata"
sidebar_current: "docs-oci-datasource-log_analytics-namespace_properties_metadata"
description: |-
  Provides the list of Namespace Properties Metadata in Oracle Cloud Infrastructure Log Analytics service
---

# Data Source: oci_log_analytics_namespace_properties_metadata
This data source provides the list of Namespace Properties Metadata in Oracle Cloud Infrastructure Log Analytics service.

Returns a list of properties along with their metadata.


## Example Usage

```hcl
data "oci_log_analytics_namespace_properties_metadata" "test_namespace_properties_metadata" {
	#Required
	namespace = var.namespace_properties_metadata_namespace

	#Optional
	constraints = var.namespace_properties_metadata_constraints
	display_text = var.namespace_properties_metadata_display_text
	level = var.namespace_properties_metadata_level
	name = var.namespace_properties_metadata_name
}
```

## Argument Reference

The following arguments are supported:

* `constraints` - (Optional) The constraints that apply to the properties at a certain level. 
* `display_text` - (Optional) The property display text used for filtering. Only properties matching the specified display name or description will be returned. 
* `level` - (Optional) The level for which applicable properties are to be listed. 
* `name` - (Optional) The property name used for filtering. 
* `namespace` - (Required) The Logging Analytics namespace used for the request. 


## Attributes Reference

The following attributes are exported:

* `property_metadata_summary_collection` - The list of property_metadata_summary_collection.

### NamespacePropertiesMetadata Reference

The following attributes are exported:

* `items` - An array of properties along with their metadata summary.
	* `default_value` - The default property value.
	* `description` - The property description.
	* `display_name` - The property display name.
	* `levels` - A list of levels at which the property could be defined.
		* `constraints` - A string representation of constraints that apply at this level. For example, a property defined at SOURCE level could further be applicable only for SOURCE_TYPE:database_sql. 
		* `name` - The level name.
	* `name` - The property name.

