---
subcategory: "Data Catalog"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datacatalog_data_assets"
sidebar_current: "docs-oci-datasource-datacatalog-data_assets"
description: |-
  Provides the list of Data Assets in Oracle Cloud Infrastructure Data Catalog service
---

# Data Source: oci_datacatalog_data_assets
This data source provides the list of Data Assets in Oracle Cloud Infrastructure Data Catalog service.

Returns a list of data assets within a data catalog.

## Example Usage

```hcl
data "oci_datacatalog_data_assets" "test_data_assets" {
	#Required
	catalog_id = oci_datacatalog_catalog.test_catalog.id

	#Optional
	created_by_id = oci_datacatalog_created_by.test_created_by.id
	display_name = var.data_asset_display_name
	display_name_contains = var.data_asset_display_name_contains
	external_key = var.data_asset_external_key
	fields = var.data_asset_fields
	state = var.data_asset_state
	time_created = var.data_asset_time_created
	time_updated = var.data_asset_time_updated
	type_key = var.data_asset_type_key
	updated_by_id = oci_datacatalog_updated_by.test_updated_by.id
}
```

## Argument Reference

The following arguments are supported:

* `catalog_id` - (Required) Unique catalog identifier.
* `created_by_id` - (Optional) OCID of the user who created the resource.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given. The match is not case sensitive.
* `display_name_contains` - (Optional) A filter to return only resources that match display name pattern given. The match is not case sensitive. For Example : /folders?displayNameContains=Cu.* The above would match all folders with display name that starts with "Cu" or has the pattern "Cu" anywhere in between. 
* `external_key` - (Optional) Unique external identifier of this resource in the external source system.
* `fields` - (Optional) Specifies the fields to return in a data asset summary response. 
* `state` - (Optional) A filter to return only resources that match the specified lifecycle state. The value is case insensitive.
* `time_created` - (Optional) Time that the resource was created. An [RFC3339](https://tools.ietf.org/html/rfc3339) formatted datetime string.
* `time_updated` - (Optional) Time that the resource was updated. An [RFC3339](https://tools.ietf.org/html/rfc3339) formatted datetime string.
* `type_key` - (Optional) The key of the object type.
* `updated_by_id` - (Optional) OCID of the user who updated the resource.


## Attributes Reference

The following attributes are exported:

* `data_asset_collection` - The list of data_asset_collection.

### DataAsset Reference

The following attributes are exported:

* `catalog_id` - The data catalog's OCID.
* `created_by_id` - OCID of the user who created the data asset.
* `description` - Detailed description of the data asset.
* `display_name` - A user-friendly display name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `external_key` - External URI that can be used to reference the object. Format will differ based on the type of object. 
* `key` - Unique data asset key that is immutable.
* `properties` - A map of maps that contains the properties which are specific to the asset type. Each data asset type definition defines it's set of required and optional properties. The map keys are category names and the values are maps of property name to property value. Every property is contained inside of a category. Most data assets have required properties within the "default" category. Example: `{"properties": { "default": { "host": "host1", "port": "1521", "database": "orcl"}}}` 
* `state` - The current state of the data asset.
* `time_created` - The date and time the data asset was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: `2019-03-25T21:10:29.600Z` 
* `time_harvested` - The last time that a harvest was performed on the data asset. An [RFC3339](https://tools.ietf.org/html/rfc3339) formatted datetime string. 
* `time_updated` - The last time that any change was made to the data asset. An [RFC3339](https://tools.ietf.org/html/rfc3339) formatted datetime string. 
* `type_key` - The key of the object type. Type key's can be found via the '/types' endpoint.
* `updated_by_id` - OCID of the user who last modified the data asset.
* `uri` - URI to the data asset instance in the API.

