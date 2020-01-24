---
subcategory: "Data Catalog"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datacatalog_catalogs"
sidebar_current: "docs-oci-datasource-datacatalog-catalogs"
description: |-
  Provides the list of Catalogs in Oracle Cloud Infrastructure Data Catalog service
---

# Data Source: oci_datacatalog_catalogs
This data source provides the list of Catalogs in Oracle Cloud Infrastructure Data Catalog service.

Returns a list of Catalogs.


## Example Usage

```hcl
data "oci_datacatalog_catalogs" "test_catalogs" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	display_name = "${var.catalog_display_name}"
	state = "${var.catalog_state}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given. The match is not case sensitive.
* `state` - (Optional) A filter to return only resources that match the specified lifecycle state. The value is case insensitive.


## Attributes Reference

The following attributes are exported:

* `catalogs` - The list of catalogs.

### Catalog Reference

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

