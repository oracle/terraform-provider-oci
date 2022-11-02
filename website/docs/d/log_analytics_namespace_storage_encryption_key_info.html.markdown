---
subcategory: "Log Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_log_analytics_namespace_storage_encryption_key_info"
sidebar_current: "docs-oci-datasource-log_analytics-namespace_storage_encryption_key_info"
description: |-
  Provides details about a specific Namespace Storage Encryption Key Info in Oracle Cloud Infrastructure Log Analytics service
---

# Data Source: oci_log_analytics_namespace_storage_encryption_key_info
This data source provides details about a specific Namespace Storage Encryption Key Info resource in Oracle Cloud Infrastructure Log Analytics service.

This API returns the list of customer owned encryption key info.

## Example Usage

```hcl
data "oci_log_analytics_namespace_storage_encryption_key_info" "test_namespace_storage_encryption_key_info" {
	#Required
	namespace = var.namespace_storage_encryption_key_info_namespace
}
```

## Argument Reference

The following arguments are supported:

* `namespace` - (Required) The Logging Analytics namespace used for the request. 


## Attributes Reference

The following attributes are exported:

* `items` - This is an array of encryption key info. There are at most 2 items in the list. 
	* `key_id` - This is the key OCID of the encryption key (null if Oracle-managed).
	* `key_source` - This is the source of the encryption key.
	* `key_type` - This is the type of data to be encrypted. It can be either active or archival.

