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

Returns a list of Data Assets within a data catalog.

## Example Usage

```hcl
data "oci_datacatalog_data_assets" "test_data_assets" {
	#Required
	catalog_id = "${oci_datacatalog_catalog.test_catalog.id}"

	#Optional
	created_by_id = "${oci_datacatalog_created_by.test_created_by.id}"
	display_name = "${var.data_asset_display_name}"
	external_key = "${var.data_asset_external_key}"
	fields = "${var.data_asset_fields}"
	state = "${var.data_asset_state}"
	time_created = "${var.data_asset_time_created}"
	time_updated = "${var.data_asset_time_updated}"
	type_key = "${var.data_asset_type_key}"
	updated_by_id = "${oci_datacatalog_updated_by.test_updated_by.id}"
}
```

## Argument Reference

The following arguments are supported:

* `catalog_id` - (Required) unique Catalog identifier
* `created_by_id` - (Optional) Id (OCID) of the user who created the resource.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given. The match is not case sensitive.
* `external_key` - (Optional) Unique external identifier of this resource in the external source system.
* `fields` - (Optional) Used to control which fields are returned in a Data Asset summary response. 
* `state` - (Optional) A filter to return only resources that match the specified lifecycle state. The value is case insensitive.
* `time_created` - (Optional) Time that the Resource was created. An RFC3339 formatted datetime string.
* `time_updated` - (Optional) Time that the Resource was updated. An RFC3339 formatted datetime string.
* `type_key` - (Optional) The key of the object type.
* `updated_by_id` - (Optional) Id of the user who updated the resource.


## Attributes Reference

The following attributes are exported:

* `data_asset_collection` - The list of data_asset_collection.

### DataAsset Reference

The following attributes are exported:

* `catalog_id` - The Catalog's Oracle ID (OCID).
* `created_by_id` - Id (OCID) of the user who created the Data Asset.
* `description` - Detailed description of the Data Asset.
* `display_name` - The display name of a user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `external_key` - External uri which can be used to reference the object. Format will differ based on the type of object. 
* `key` - Unique Data Asset key that is immutable.
* `properties` - A map of maps which contains the properties which are specific to the asset type. Each Data Asset type definition defines it's set of required and optional properties. The map keys are category names and the values are maps of property name to property value. Every property is contained inside of a category. Most Data Assets have required properties within the "default" category. Example: `{"properties": { "default": { "host": "host1", "port": "1521", "database": "orcl"}}}` 
* `state` - The current state of the Data Asset.
* `time_created` - The date and time the DataAsset was created, in the format defined by RFC3339. Example: `2019-03-25T21:10:29.600Z` 
* `time_updated` - The last time that any change was made to the Data Asset. An RFC3339 formatted datetime string. 
* `type_key` - The key of the object type. Type key's can be found via the '/types' endpoint.
* `updated_by_id` - Id (OCID) of the user who last modified the Data Asset.
* `uri` - URI to the Data Asset instance in the API.

