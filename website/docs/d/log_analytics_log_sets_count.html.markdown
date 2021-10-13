---
subcategory: "Log Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_log_analytics_log_sets_count"
sidebar_current: "docs-oci-datasource-log_analytics-log_sets_count"
description: |-
  Provides details about a specific Log Sets Count in Oracle Cloud Infrastructure Log Analytics service
---

# Data Source: oci_log_analytics_log_sets_count
This data source provides details about a specific Log Sets Count resource in Oracle Cloud Infrastructure Log Analytics service.

This API returns the count of distinct log sets.


## Example Usage

```hcl
data "oci_log_analytics_log_sets_count" "test_log_sets_count" {
	#Required
	namespace = var.log_sets_count_namespace
}
```

## Argument Reference

The following arguments are supported:

* `namespace` - (Required) The Logging Analytics namespace used for the request. 


## Attributes Reference

The following attributes are exported:

* `log_sets_count` - This is the total number of log sets the tenancy has configured.

