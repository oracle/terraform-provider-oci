---
subcategory: "Log Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_log_analytics_log_analytics_log_groups_summary"
sidebar_current: "docs-oci-datasource-log_analytics-log_analytics_log_groups_summary"
description: |-
  Provides details about a specific Log Analytics Log Groups Summary in Oracle Cloud Infrastructure Log Analytics service
---

# Data Source: oci_log_analytics_log_analytics_log_groups_summary
This data source provides details about a specific Log Analytics Log Groups Summary resource in Oracle Cloud Infrastructure Log Analytics service.

Returns the count of log groups in a compartment.


## Example Usage

```hcl
data "oci_log_analytics_log_analytics_log_groups_summary" "test_log_analytics_log_groups_summary" {
	#Required
	compartment_id = var.compartment_id
	namespace = var.log_analytics_log_groups_summary_namespace
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `namespace` - (Required) The Logging Analytics namespace used for the request. 


## Attributes Reference

The following attributes are exported:

* `count` - The log group count.

