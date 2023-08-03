---
subcategory: "Log Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_log_analytics_namespace_storage_recalled_data_size"
sidebar_current: "docs-oci-datasource-log_analytics-namespace_storage_recalled_data_size"
description: |-
  Provides details about a specific Namespace Storage Recalled Data Size in Oracle Cloud Infrastructure Log Analytics service
---

# Data Source: oci_log_analytics_namespace_storage_recalled_data_size
This data source provides details about a specific Namespace Storage Recalled Data Size resource in Oracle Cloud Infrastructure Log Analytics service.

This API gets the datasize of recalls for a given timeframe


## Example Usage

```hcl
data "oci_log_analytics_namespace_storage_recalled_data_size" "test_namespace_storage_recalled_data_size" {
	#Required
	namespace = var.namespace_storage_recalled_data_size_namespace

	#Optional
	time_data_ended = var.namespace_storage_recalled_data_size_time_data_ended
	time_data_started = var.namespace_storage_recalled_data_size_time_data_started
}
```

## Argument Reference

The following arguments are supported:

* `namespace` - (Required) The Logging Analytics namespace used for the request. 
* `time_data_ended` - (Optional) This is the end of the time range for recalled data
* `time_data_started` - (Optional) This is the start of the time range for recalled data


## Attributes Reference

The following attributes are exported:

* `not_recalled_data_in_bytes` - This is the size of the archival data not recalled yet
* `recalled_data_in_bytes` - This is the size of the recalled data
* `time_data_ended` - This is the end of the time range of the archival data
* `time_data_started` - This is the start of the time range of the archival data

