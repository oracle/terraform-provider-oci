---
subcategory: "Data Catalog"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datacatalog_catalog"
sidebar_current: "docs-oci-resource-datacatalog-catalog"
description: |-
  Provides the Catalog resource in Oracle Cloud Infrastructure Data Catalog service
---

# oci_datacatalog_catalog
This resource provides the Catalog resource in Oracle Cloud Infrastructure Data Catalog service.

Creates a new Data Catalog instance which includes a console and api url for managing metadata operations.  
For more information, please see the documentation.


## Example Usage

```hcl
resource "oci_datacatalog_catalog" "test_catalog" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	display_name = "${var.catalog_display_name}"
	freeform_tags = {"bar-key"= "value"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) Compartment Identifier
* `defined_tags` - (Optional) (Updatable) Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Optional) (Updatable) Catalog Identifier
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Import

Catalogs can be imported using the `id`, e.g.

```
$ terraform import oci_datacatalog_catalog.test_catalog "id"
```

