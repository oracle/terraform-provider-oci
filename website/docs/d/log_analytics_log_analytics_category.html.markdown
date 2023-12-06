---
subcategory: "Log Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_log_analytics_log_analytics_category"
sidebar_current: "docs-oci-datasource-log_analytics-log_analytics_category"
description: |-
  Provides details about a specific Log Analytics Category in Oracle Cloud Infrastructure Log Analytics service
---

# Data Source: oci_log_analytics_log_analytics_category
This data source provides details about a specific Log Analytics Category resource in Oracle Cloud Infrastructure Log Analytics service.

Gets detailed information about the category with the specified name.


## Example Usage

```hcl
data "oci_log_analytics_log_analytics_category" "test_log_analytics_category" {
	#Required
	name = var.log_analytics_category_name
	namespace = var.log_analytics_category_namespace
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The category name.
* `namespace` - (Required) The Logging Analytics namespace used for the request. 


## Attributes Reference

The following attributes are exported:

* `description` - The category description.
* `display_name` - The category display name.
* `is_system` - The system flag. A value of false denotes a user-created category. A value of true denotes an Oracle-defined category. 
* `name` - The unique name that identifies the category.
* `type` - The category type. Values include "PRODUCT", "TIER", "VENDOR" and "GENERIC".

