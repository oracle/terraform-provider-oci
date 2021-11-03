---
subcategory: "Log Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_log_analytics_log_analytics_resource_categories_management"
sidebar_current: "docs-oci-resource-log_analytics-log_analytics_resource_categories_management"
description: |-
  Provides the Log Analytics Resource Categories Management resource in Oracle Cloud Infrastructure Log Analytics service
---

# oci_log_analytics_log_analytics_resource_categories_management
This resource provides the Log Analytics Resource Categories Management resource in Oracle Cloud Infrastructure Log Analytics service.

Updates the category assignments of DASHBOARD and SAVEDSEARCH resources.


## Example Usage

```hcl
resource "oci_log_analytics_log_analytics_resource_categories_management" "test_log_analytics_resource_categories_management" {
	#Required
	namespace = var.log_analytics_resource_categories_management_namespace
        resource_id = oci_log_analytics_resource_categories_management_resource_id
        resource_type = var.log_analytics_resource_categories_management_resource_type
        resource_categories = var.log_analytics_resource_categories_management_resource_categories
}
```

## Argument Reference

The following arguments are supported:

* `namespace` - (Required) The Logging Analytics namespace used for the request.
* `resource_categories` - (Required) The list of categories to be assigned to the resource.
* `resource_id` - (Required) The resource unique identifier for which catagories are managed.
* `resource_type` - (Required) The resource type of the resource for which categories are managed. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Log Analytics Resource Categories Management
	* `update` - (Defaults to 20 minutes), when updating the Log Analytics Resource Categories Management
	* `delete` - (Defaults to 20 minutes), when destroying the Log Analytics Resource Categories Management


## Import

LogAnalyticsResourceCategoriesManagement cannot be imported.

