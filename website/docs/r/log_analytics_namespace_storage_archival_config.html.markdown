---
subcategory: "Log Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_log_analytics_namespace_storage_archival_config"
sidebar_current: "docs-oci-resource-log_analytics-namespace_storage_archival_config"
description: |-
  Provides the Namespace Storage Archival Config resource in Oracle Cloud Infrastructure Log Analytics service
---

# oci_log_analytics_namespace_storage_archival_config
This resource provides the Namespace Storage Archival Config resource in Oracle Cloud Infrastructure Log Analytics service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/logan-api-spec/latest/NamespaceStorageArchivalConfig

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/log_analytics

This API updates the archiving configuration


## Example Usage

```hcl
resource "oci_log_analytics_namespace_storage_archival_config" "test_namespace_storage_archival_config" {
	#Required
	archiving_configuration {

		#Optional
		active_storage_duration = var.namespace_storage_archival_config_archiving_configuration_active_storage_duration
		archival_storage_duration = var.namespace_storage_archival_config_archiving_configuration_archival_storage_duration
	}
	namespace = var.namespace_storage_archival_config_namespace
}
```

## Argument Reference

The following arguments are supported:

* `archiving_configuration` - (Required) (Updatable) This is the configuration for data archiving in object storage
	* `active_storage_duration` - (Optional) (Updatable) This is the duration data in active storage before data is archived, as described in https://en.wikipedia.org/wiki/ISO_8601#Durations. The largest supported unit is D, e.g. P365D (not P1Y) or P14D (not P2W). 
	* `archival_storage_duration` - (Optional) (Updatable) This is the duration before archived data is deleted from object storage, as described in https://en.wikipedia.org/wiki/ISO_8601#Durations The largest supported unit is D, e.g. P365D (not P1Y) or P14D (not P2W). 
* `namespace` - (Required) The Logging Analytics namespace used for the request.

** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `archiving_configuration` - This is the configuration for data archiving in object storage
	* `active_storage_duration` - This is the duration data in active storage before data is archived, as described in https://en.wikipedia.org/wiki/ISO_8601#Durations. The largest supported unit is D, e.g. P365D (not P1Y) or P14D (not P2W). 
	* `archival_storage_duration` - This is the duration before archived data is deleted from object storage, as described in https://en.wikipedia.org/wiki/ISO_8601#Durations The largest supported unit is D, e.g. P365D (not P1Y) or P14D (not P2W). 
* `is_archiving_enabled` - This indicates if old data can be archived for a tenancy

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Namespace Storage Archival Config
	* `update` - (Defaults to 20 minutes), when updating the Namespace Storage Archival Config
	* `delete` - (Defaults to 20 minutes), when destroying the Namespace Storage Archival Config
