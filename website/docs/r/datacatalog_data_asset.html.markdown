---
subcategory: "Data Catalog"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datacatalog_data_asset"
sidebar_current: "docs-oci-resource-datacatalog-data_asset"
description: |-
  Provides the Data Asset resource in Oracle Cloud Infrastructure Data Catalog service
---

# oci_datacatalog_data_asset
This resource provides the Data Asset resource in Oracle Cloud Infrastructure Data Catalog service.

Create a new data asset.

## Example Usage

```hcl
resource "oci_datacatalog_data_asset" "test_data_asset" {
	#Required
	catalog_id = oci_datacatalog_catalog.test_catalog.id
	display_name = var.data_asset_display_name
	type_key = var.data_asset_type_key

	#Optional
	description = var.data_asset_description
	properties = var.data_asset_properties
}
```

## Argument Reference

The following arguments are supported:

* `catalog_id` - (Required) Unique catalog identifier.
* `description` - (Optional) (Updatable) Detailed description of the data asset.
* `display_name` - (Required) (Updatable) A user-friendly display name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `properties` - (Optional) (Updatable) A map of maps that contains the properties which are specific to the data asset type. Each data asset type definition defines it's set of required and optional properties. The map keys are category names and the values are maps of property name to property value. Every property is contained inside of a category. Most data assets have required properties within the "default" category. To determine the set of optional and required properties for a data asset type, a query can be done on '/types?type=dataAsset' that returns a collection of all data asset types. The appropriate data asset type, which includes definitions of all of it's properties, can be identified from this collection. Example: `{"properties": { "default": { "host": "host1", "port": "1521", "database": "orcl"}}}` . Terraform treats all map of maps as a flattened map with `.` denoting each level. For more information check out this [example](https://github.com/terraform-providers/terraform-provider-oci/blob/master/examples/datacatalog/main.tf)
* `type_key` - (Required) The key of the data asset type. This can be obtained via the '/types' endpoint.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Data Asset
	* `update` - (Defaults to 20 minutes), when updating the Data Asset
	* `delete` - (Defaults to 20 minutes), when destroying the Data Asset


## Import

DataAssets can be imported using the `id`, e.g.

```
$ terraform import oci_datacatalog_data_asset.test_data_asset "catalogs/{catalogId}/dataAssets/{dataAssetKey}" 
```

