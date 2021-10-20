---
subcategory: "Log Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_log_analytics_log_analytics_unprocessed_data_bucket"
sidebar_current: "docs-oci-datasource-log_analytics-log_analytics_unprocessed_data_bucket"
description: |-
  Provides details about a specific Log Analytics Unprocessed Data Bucket in Oracle Cloud Infrastructure Log Analytics service
---

# Data Source: oci_log_analytics_log_analytics_unprocessed_data_bucket
This data source provides details about a specific Log Analytics Unprocessed Data Bucket resource in Oracle Cloud Infrastructure Log Analytics service.

This API retrieves details of the configured bucket that stores unprocessed payloads.


## Example Usage

```hcl
data "oci_log_analytics_log_analytics_unprocessed_data_bucket" "test_log_analytics_unprocessed_data_bucket" {
	#Required
	namespace = var.log_analytics_unprocessed_data_bucket_namespace
}
```

## Argument Reference

The following arguments are supported:

* `namespace` - (Required) The Logging Analytics namespace used for the request. 


## Attributes Reference

The following attributes are exported:

* `bucket` - Name of the Object Storage bucket.
* `is_enabled` - Flag that specifies if this configuration is enabled or not. 
* `namespace` - Object Storage namespace.
* `time_created` - The time when this record is created. An RFC3339 formatted datetime string.
* `time_updated` - The latest time when this record is updated. An RFC3339 formatted datetime string.

