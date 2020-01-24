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

Gets a Catalog by identifier

## Example Usage

```hcl
data "oci_datacatalog_catalog" "test_catalog" {
	#Required
	catalog_id = "${oci_datacatalog_catalog.test_catalog.id}"
}
```

## Argument Reference

The following arguments are supported:

* `catalog_id` - (Required) unique Catalog identifier


## Attributes Reference

The following attributes are exported:

* `compartment_id` - Compartment Identifier
* `defined_tags` - Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - Catalog Identifier, can be renamed
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - Unique identifier that is immutable on creation
* `lifecycle_details` - An message describing the current state in more detail.  For example, can be used to provide actionable information for a resource in Failed state. 
* `number_of_objects` - The number of data objects added to the Catalog. Please see the Data Catalog documentation for further information on how this is calculated. 
* `service_api_url` - The REST front endpoint url to the catalog instance
* `service_console_url` - The console front endpoint url to the catalog instance
* `state` - The current state of the catalog resource.
* `time_created` - The time the the Catalog was created. An RFC3339 formatted datetime string
* `time_updated` - The time the Catalog was updated. An RFC3339 formatted datetime string

