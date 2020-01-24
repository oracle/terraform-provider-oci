---
subcategory: "Data Catalog"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datacatalog_data_asset"
sidebar_current: "docs-oci-datasource-datacatalog-data_asset"
description: |-
  Provides details about a specific Data Asset in Oracle Cloud Infrastructure Data Catalog service
---

# Data Source: oci_datacatalog_data_asset
This data source provides details about a specific Data Asset resource in Oracle Cloud Infrastructure Data Catalog service.

Get a specific DataAsset for the given key within a data catalog.

## Example Usage

```hcl
data "oci_datacatalog_data_asset" "test_data_asset" {
	#Required
	catalog_id = "${oci_datacatalog_catalog.test_catalog.id}"
	data_asset_key = "${var.data_asset_data_asset_key}"

	#Optional
	fields = "${var.data_asset_fields}"
}
```

## Argument Reference

The following arguments are supported:

* `catalog_id` - (Required) unique Catalog identifier
* `data_asset_key` - (Required) Unique Data Asset key.
* `fields` - (Optional) Used to control which fields are returned in a Data Asset response. 


## Attributes Reference

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

