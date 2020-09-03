---
subcategory: "Data Catalog"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datacatalog_catalog"
sidebar_current: "docs-oci-datasource-datacatalog-catalog"
description: |-
  Provides details about a specific Catalog in Oracle Cloud Infrastructure Data Catalog service
---

# Data Source: oci_datacatalog_catalog
This data source provides details about a specific Catalog resource in Oracle Cloud Infrastructure Data Catalog service.

Gets a data catalog by identifier.

## Example Usage

```hcl
data "oci_datacatalog_catalog" "test_catalog" {
	#Required
	catalog_id = oci_datacatalog_catalog.test_catalog.id
}
```

## Argument Reference

The following arguments are supported:

* `catalog_id` - (Required) Unique catalog identifier.


## Attributes Reference

The following attributes are exported:

* `attached_catalog_private_endpoints` - The list of private reverse connection endpoints attached to the catalog
* `compartment_id` - Compartment identifier.
* `defined_tags` - Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - Data catalog identifier, which can be renamed.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - Unique identifier that is immutable on creation.
* `lifecycle_details` - An message describing the current state in more detail.  For example, it can be used to provide actionable information for a resource in 'Failed' state. 
* `number_of_objects` - The number of data objects added to the data catalog. Please see the data catalog documentation for further information on how this is calculated. 
* `service_api_url` - The REST front endpoint URL to the data catalog instance.
* `service_console_url` - The console front endpoint URL to the data catalog instance.
* `state` - The current state of the data catalog resource.
* `time_created` - The time the data catalog was created. An [RFC3339](https://tools.ietf.org/html/rfc3339) formatted datetime string.
* `time_updated` - The time the data catalog was updated. An [RFC3339](https://tools.ietf.org/html/rfc3339) formatted datetime string.

