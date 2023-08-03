---
subcategory: "Log Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_log_analytics_namespace_storage_recall_count"
sidebar_current: "docs-oci-datasource-log_analytics-namespace_storage_recall_count"
description: |-
  Provides details about a specific Namespace Storage Recall Count in Oracle Cloud Infrastructure Log Analytics service
---

# Data Source: oci_log_analytics_namespace_storage_recall_count
This data source provides details about a specific Namespace Storage Recall Count resource in Oracle Cloud Infrastructure Log Analytics service.

This API gets the number of recalls made and the maximum recalls that can be made


## Example Usage

```hcl
data "oci_log_analytics_namespace_storage_recall_count" "test_namespace_storage_recall_count" {
	#Required
	namespace = var.namespace_storage_recall_count_namespace
}
```

## Argument Reference

The following arguments are supported:

* `namespace` - (Required) The Logging Analytics namespace used for the request. 


## Attributes Reference

The following attributes are exported:

* `recall_count` - This is the total number of recalls made so far
* `recall_failed` - This is the number of recalls that failed
* `recall_limit` - This is the maximum number of recalls (including successful and pending recalls) allowed
* `recall_pending` - This is the number of recalls in pending state
* `recall_succeeded` - This is the number of recalls that succeeded

