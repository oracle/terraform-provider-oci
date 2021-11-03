---
subcategory: "Log Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_log_analytics_log_analytics_categories_list"
sidebar_current: "docs-oci-datasource-log_analytics-log_analytics_categories_list"
description: |-
  Provides details about Categories in Oracle Cloud Infrastructure Log Analytics service
---

# Data Source: oci_log_analytics_log_analytics_categories_list
This data source provides details about Categories in Oracle Cloud Infrastructure Log Analytics service.

Returns a list of categories, containing detailed information about them. You may limit the number of results, provide sorting order, and filter by information such as category name or description.


## Example Usage

```hcl
data "oci_log_analytics_log_analytics_categories_list" "test_log_analytics_categories_list" {
	#Required
	namespace = var.log_analytics_categories_list_namespace

	#Optional
	category_display_text = var.log_analytics_categories_list_category_display_text
	category_type = var.log_analytics_categories_list_category_type
	name = var.log_analytics_categories_list_name
}
```

## Argument Reference

The following arguments are supported:

* `category_display_text` - (Optional) The category display text used for filtering. Only categories matching the specified display name or description will be returned. 
* `category_type` - (Optional) A comma-separated list of category types used for filtering. Only categories of the specified types will be returned. 
* `name` - (Optional) A filter to return only log analytics category whose name matches the entire name given. The match is case-insensitive. 
* `namespace` - (Required) The Logging Analytics namespace used for the request. 


## Attributes Reference

The following attributes are exported:

* `items` - An array of categories.
	* `description` - The category description.
	* `display_name` - The category display name.
	* `is_system` - The system flag. A value of false denotes a user-created category. A value of true denotes an Oracle-defined category. 
	* `name` - The unique name that identifies the category.
	* `type` - The category type. Values include "PRODUCT", "TIER", "VENDOR" and "GENERIC".

