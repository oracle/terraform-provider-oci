---
subcategory: "Log Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_log_analytics_log_analytics_resource_categories_list"
sidebar_current: "docs-oci-datasource-log_analytics-log_analytics_resource_categories_list"
description: |-
  Provides details about Resource Categories in Oracle Cloud Infrastructure Log Analytics service
---

# Data Source: oci_log_analytics_log_analytics_resource_categories_list
This data source provides details about Resource Categories in Oracle Cloud Infrastructure Log Analytics service.

Returns a list of resources and their category assignments.


## Example Usage

```hcl
data "oci_log_analytics_log_analytics_resource_categories_list" "test_log_analytics_resource_categories_list" {
	#Required
	namespace = var.log_analytics_resource_categories_list_namespace

	#Optional
        resource_ids = var.log_analytics_resource_categories_list_resource_ids
        resource_types = var.log_analytics_resource_categories_list_resource_types
        resource_categories = var.log_analytics_resource_categories_list_resource_categories
}
```

## Argument Reference

The following arguments are supported:

* `namespace` - (Required) The Logging Analytics namespace used for the request. 
* `resource_categories` - (Optional) A comma-separated list of category names used for filtering
* `resource_ids` - (Optional) A comma-separated list of resource unique identifiers used for filtering. Only resources with matching unique identifiers will be returned. 
* `resource_types` - (Optional) A comma-separated list of resource types used for filtering. Only resources of the types specified will be returned. Examples include SOURCE, PARSER, LOOKUP, etc. 


## Attributes Reference

The following attributes are exported:

* `categories` - An array of categories. The array contents include detailed information about the distinct set of categories assigned to all the listed resources under items. 
	* `description` - The category description.
	* `display_name` - The category display name.
	* `is_system` - The system flag. A value of false denotes a user-created category. A value of true denotes an Oracle-defined category. 
	* `name` - The unique name that identifies the category.
	* `type` - The category type. Values include "PRODUCT", "TIER", "VENDOR" and "GENERIC".
* `items` - A list of resources and their category assignments
	* `category_name` - The category name to which this resource belongs.
	* `is_system` - The system flag. A value of false denotes a user-created category assignment. A value of true denotes an Oracle-defined category assignment. 
	* `resource_id` - The unique identifier of the resource, usually a name or ocid.
	* `resource_type` - The resource type.

