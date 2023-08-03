---
subcategory: "Log Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_log_analytics_log_analytics_preference"
sidebar_current: "docs-oci-datasource-log_analytics-log_analytics_preference"
description: |-
  Provides details about a specific Log Analytics Preference in Oracle Cloud Infrastructure Log Analytics service
---

# Data Source: oci_log_analytics_log_analytics_preference
This data source provides details about a specific Log Analytics Preference resource in Oracle Cloud Infrastructure Log Analytics service.

Lists the tenant preferences such as DEFAULT_HOMEPAGE and collection properties.


## Example Usage

```hcl
data "oci_log_analytics_log_analytics_preference" "test_log_analytics_preference" {
	#Required
	namespace = var.log_analytics_preference_namespace
}
```

## Argument Reference

The following arguments are supported:

* `namespace` - (Required) The Logging Analytics namespace used for the request. 


## Attributes Reference

The following attributes are exported:

* `items` - An array of tenant preferences.
	* `name` - The preference name.
	* `value` - The preference value.

