---
subcategory: "Data Catalog"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datacatalog_catalog_type"
sidebar_current: "docs-oci-datasource-datacatalog-catalog_type"
description: |-
  Provides details about a specific Catalog Type in Oracle Cloud Infrastructure Data Catalog service
---

# Data Source: oci_datacatalog_catalog_type
This data source provides details about a specific Catalog Type resource in Oracle Cloud Infrastructure Data Catalog service.

Gets a specific type by key within a data catalog.

## Example Usage

```hcl
data "oci_datacatalog_catalog_type" "test_catalog_type" {
	#Required
	catalog_id = "${oci_datacatalog_catalog.test_catalog.id}"
	type_key = "${var.catalog_type_type_key}"

	#Optional
	fields = "${var.catalog_type_fields}"
}
```

## Argument Reference

The following arguments are supported:

* `catalog_id` - (Required) Unique catalog identifier.
* `fields` - (Optional) Specifies the fields to return in a type response. 
* `type_key` - (Required) Unique type key.


## Attributes Reference

The following attributes are exported:

* `catalog_id` - The data catalog's OCID.
* `description` - Detailed description of the type.
* `external_type_name` - Mapping type equivalence in the external system.
* `is_approved` - Indicates whether the type is approved for use as a classifying object.
* `is_internal` - Indicates whether the type is internal, making it unavailable for use by metadata elements.
* `is_tag` - Indicates whether the type can be used for tagging metadata elements.
* `key` - Unique type key that is immutable.
* `name` - The immutable name of the type.
* `properties` - A map of arrays which defines the type specific properties, both required and optional. The map keys are category names and the values are arrays contiaing all property details. Every property is contained inside of a category. Most types have required properties within the "default" category. Example: `{ "properties": { "default": { "attributes:": [ { "name": "host", "type": "string", "isRequired": true, "isUpdatable": false }, ... ] } } }` 
* `state` - The current state of the type.
* `type_category` - Indicates the category this type belongs to. For instance, data assets, connections.
* `uri` - URI to the type instance in the API.

