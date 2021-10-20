---
subcategory: "Log Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_log_analytics_log_analytics_unprocessed_data_bucket_management"
sidebar_current: "docs-oci-resource-log_analytics-log_analytics_unprocessed_data_bucket_management"
description: |-
  Provides the Log Analytics Unprocessed Data Bucket Management resource in Oracle Cloud Infrastructure Log Analytics service
---

# oci_log_analytics_log_analytics_unprocessed_data_bucket_management
This resource provides the Log Analytics Unprocessed Data Bucket Management resource in Oracle Cloud Infrastructure Log Analytics service.

This API configures a bucket to store unprocessed payloads.
While processing there could be reasons a payload cannot be processed (mismatched structure, corrupted archive format, etc),
if configured the payload would be uploaded to the bucket for verification.


## Example Usage

```hcl
resource "oci_log_analytics_log_analytics_unprocessed_data_bucket_management" "test_log_analytics_unprocessed_data_bucket_management" {
	#Required
	bucket = var.log_analytics_unprocessed_data_bucket_management_bucket
	namespace = var.log_analytics_unprocessed_data_bucket_management_namespace
}
```

## Argument Reference

The following arguments are supported:

* `bucket` - (Required) Name of the Object Storage bucket.
* `namespace` - (Required) The Logging Analytics namespace used for the request. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `bucket` - Name of the Object Storage bucket.
* `is_enabled` - Flag that specifies if this configuration is enabled or not. 
* `namespace` - Object Storage namespace.
* `time_created` - The time when this record is created. An RFC3339 formatted datetime string.
* `time_updated` - The latest time when this record is updated. An RFC3339 formatted datetime string.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Log Analytics Unprocessed Data Bucket Management
	* `update` - (Defaults to 20 minutes), when updating the Log Analytics Unprocessed Data Bucket Management
	* `delete` - (Defaults to 20 minutes), when destroying the Log Analytics Unprocessed Data Bucket Management


## Import

Import is not supported for LogAnalyticsUnprocessedDataBucketManagement
