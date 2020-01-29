---
subcategory: "Data Catalog"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datacatalog_catalog_types"
sidebar_current: "docs-oci-datasource-datacatalog-catalog_types"
description: |-
  Provides the list of Catalog Types in Oracle Cloud Infrastructure Data Catalog service
---

# Data Source: oci_datacatalog_catalog_types
This data source provides the list of Catalog Types in Oracle Cloud Infrastructure Data Catalog service.

Returns a list of all types within a data catalog.

## Example Usage

```hcl
data "oci_datacatalog_catalog_types" "test_catalog_types" {
	#Required
	catalog_id = "${oci_datacatalog_catalog.test_catalog.id}"

	#Optional
	external_type_name = "${var.catalog_type_external_type_name}"
	fields = "${var.catalog_type_fields}"
	is_approved = "${var.catalog_type_is_approved}"
	is_internal = "${var.catalog_type_is_internal}"
	is_tag = "${var.catalog_type_is_tag}"
	name = "${var.catalog_type_name}"
	state = "${var.catalog_type_state}"
	type_category = "${var.catalog_type_type_category}"
}
```

## Argument Reference

The following arguments are supported:

* `catalog_id` - (Required) Unique catalog identifier.
* `external_type_name` - (Optional) Data type as defined in an external system.
* `fields` - (Optional) Specifies the fields to return in a type summary response. 
* `is_approved` - (Optional) Indicates whether the type is approved for use as a classifying object.
* `is_internal` - (Optional) Indicates whether the type is internal, making it unavailable for use by metadata elements.
* `is_tag` - (Optional) Indicates whether the type can be used for tagging metadata elements.
* `name` - (Optional) Immutable resource name.
* `state` - (Optional) A filter to return only resources that match the specified lifecycle state. The value is case insensitive.
* `type_category` - (Optional) Indicates the category of this type . For example, data assets or connections.


## Attributes Reference

The following attributes are exported:

* `type_collection` - The list of type_collection.

### CatalogType Reference

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

