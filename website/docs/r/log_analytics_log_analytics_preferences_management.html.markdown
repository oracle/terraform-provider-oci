---
subcategory: "Log Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_log_analytics_log_analytics_preferences_management"
sidebar_current: "docs-oci-resource-log_analytics-log_analytics_preferences_management"
description: |-
  Provides the Log Analytics Preferences Management resource in Oracle Cloud Infrastructure Log Analytics service
---

# oci_log_analytics_log_analytics_preferences_management
This resource provides the Log Analytics Preferences Management resource in Oracle Cloud Infrastructure Log Analytics service.

Updates the tenant preferences. Currently, only "DEFAULT_HOMEPAGE" is supported.


## Example Usage

```hcl
resource "oci_log_analytics_log_analytics_preferences_management" "test_log_analytics_preferences_management" {
	#Required
	namespace = var.log_analytics_preferences_management_namespace

	#Optional
	items {

		#Optional
		name = var.log_analytics_preferences_management_items_name
		value = var.log_analytics_preferences_management_items_value
	}
}
```

## Argument Reference

The following arguments are supported:

* `items` - (Optional) An array of tenant preference details.
	* `name` - (Optional) The preference name. Currently, only "DEFAULT_HOMEPAGE" is supported.
	* `value` - (Optional) The preference value.
* `namespace` - (Required) The Logging Analytics namespace used for the request. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Log Analytics Preferences Management
	* `update` - (Defaults to 20 minutes), when updating the Log Analytics Preferences Management
	* `delete` - (Defaults to 20 minutes), when destroying the Log Analytics Preferences Management


## Import

Import is not supported for LogAnalyticsPreferencesManagement
