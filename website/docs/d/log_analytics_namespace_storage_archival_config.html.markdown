---
subcategory: "Log Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_log_analytics_namespace_storage_archival_config"
sidebar_current: "docs-oci-datasource-log_analytics-namespace_storage_archival_config"
description: |-
  Provides details about a specific Namespace Storage Archival Config in Oracle Cloud Infrastructure Log Analytics service
---

# Data Source: oci_log_analytics_namespace_storage_archival_config
This data source provides details about a specific Namespace Storage Archival Config resource in Oracle Cloud Infrastructure Log Analytics service.

This API gets the storage configuration of a tenancy


## Example Usage

```hcl
data "oci_log_analytics_namespace_storage_archival_config" "test_namespace_storage_archival_config" {
	#Required
	namespace = var.namespace_storage_archival_config_namespace
}
```

## Argument Reference

The following arguments are supported:

* `namespace` - (Required) The Logging Analytics namespace used for the request. 


## Attributes Reference

The following attributes are exported:

* `archiving_configuration` - This is the configuration for data archiving in object storage
	* `active_storage_duration` - This is the duration data in active storage before data is archived, as described in https://en.wikipedia.org/wiki/ISO_8601#Durations. The largest supported unit is D, e.g. P365D (not P1Y) or P14D (not P2W). 
	* `archival_storage_duration` - This is the duration before archived data is deleted from object storage, as described in https://en.wikipedia.org/wiki/ISO_8601#Durations The largest supported unit is D, e.g. P365D (not P1Y) or P14D (not P2W). 
* `is_archiving_enabled` - This indicates if old data can be archived for a tenancy

