---
subcategory: "Data Catalog"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datacatalog_connection"
sidebar_current: "docs-oci-resource-datacatalog-connection"
description: |-
  Provides the Connection resource in Oracle Cloud Infrastructure Data Catalog service
---

# oci_datacatalog_connection
This resource provides the Connection resource in Oracle Cloud Infrastructure Data Catalog service.

Creates a new connection.

## Example Usage

```hcl
resource "oci_datacatalog_connection" "test_connection" {
	#Required
	catalog_id = oci_datacatalog_catalog.test_catalog.id
	data_asset_key = var.connection_data_asset_key
	display_name = var.connection_display_name
	properties = var.connection_properties
	type_key = var.connection_type_key

	#Optional
	description = var.connection_description
	enc_properties = var.connection_enc_properties
	is_default = var.connection_is_default
}
```

## Argument Reference

The following arguments are supported:

* `catalog_id` - (Required) Unique catalog identifier.
* `data_asset_key` - (Required) Unique data asset key.
* `description` - (Optional) (Updatable) A description of the connection.
* `display_name` - (Required) (Updatable) A user-friendly display name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `enc_properties` - (Optional) (Updatable) A map of maps that contains the encrypted values for sensitive properties which are specific to the connection type. Each connection type definition defines it's set of required and optional properties. The map keys are category names and the values are maps of property name to property value. Every property is contained inside of a category. Most connections have required properties within the "default" category. To determine the set of optional and required properties for a connection type, a query can be done on '/types?type=connection' that returns a collection of all connection types. The appropriate connection type, which will include definitions of all of it's properties, can be identified from this collection. Example: `{"encProperties": { "default": { "password": "example-password"}}}` 
* `is_default` - (Optional) (Updatable) Indicates whether this connection is the default connection. The first connection of a data asset defaults to being the default, subsequent connections default to not being the default. If a default connection already exists, then trying to create a connection as the default will fail. In this case the default connection would need to be updated not to be the default and then the new connection can then be created as the default. 
* `properties` - (Required) (Updatable) A map of maps that contains the properties which are specific to the connection type. Each connection type definition defines it's set of required and optional properties. The map keys are category names and the values are maps of property name to property value. Every property is contained inside of a category. Most connections have required properties within the "default" category. To determine the set of optional and required properties for a connection type, a query can be done on '/types?type=connection' that returns a collection of all connection types. The appropriate connection type, which will include definitions of all of it's properties, can be identified from this collection. Example: `{"properties": { "default": { "username": "user1"}}}` . Terraform treats all map of maps as a flattened map with `.` denoting each level. For more information check out this [example](https://github.com/terraform-providers/terraform-provider-oci/blob/master/examples/datacatalog/main.tf)
* `type_key` - (Required) The key of the object type. Type key's can be found via the '/types' endpoint.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `created_by_id` - OCID of the user who created the connection.
* `data_asset_key` - Unique key of the parent data asset.
* `description` - A description of the connection.
* `display_name` - A user-friendly display name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `external_key` - Unique external key of this object from the source system.
* `is_default` - Indicates whether this connection is the default connection.
* `key` - Unique connection key that is immutable.
* `properties` - A map of maps that contains the properties which are specific to the connection type. Each connection type definition defines it's set of required and optional properties. The map keys are category names and the values are maps of property name to property value. Every property is contained inside of a category. Most connections have required properties within the "default" category. Example: `{"properties": { "default": { "username": "user1"}}}` 
* `state` - The current state of the connection.
* `time_created` - The date and time the connection was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: `2019-03-25T21:10:29.600Z` 
* `time_status_updated` - Time that the connections status was last updated. An [RFC3339](https://tools.ietf.org/html/rfc3339) formatted datetime string.
* `time_updated` - The last time that any change was made to the connection. An [RFC3339](https://tools.ietf.org/html/rfc3339) formatted datetime string. 
* `type_key` - The key of the object type. Type key's can be found via the '/types' endpoint.
* `updated_by_id` - OCID of the user who modified the connection.
* `uri` - URI to the connection instance in the API.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Connection
	* `update` - (Defaults to 20 minutes), when updating the Connection
	* `delete` - (Defaults to 20 minutes), when destroying the Connection


## Import

Connections can be imported using the `id`, e.g.

```
$ terraform import oci_datacatalog_connection.test_connection "catalogs/{catalogId}/dataAssets/{dataAssetKey}/connections/{connectionKey}" 
```

