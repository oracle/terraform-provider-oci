---
subcategory: "Data Catalog"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datacatalog_connections"
sidebar_current: "docs-oci-datasource-datacatalog-connections"
description: |-
  Provides the list of Connections in Oracle Cloud Infrastructure Data Catalog service
---

# Data Source: oci_datacatalog_connections
This data source provides the list of Connections in Oracle Cloud Infrastructure Data Catalog service.

Returns a list of all Connections for a data asset.

## Example Usage

```hcl
data "oci_datacatalog_connections" "test_connections" {
	#Required
	catalog_id = oci_datacatalog_catalog.test_catalog.id
	data_asset_key = var.connection_data_asset_key

	#Optional
	created_by_id = oci_datacatalog_created_by.test_created_by.id
	display_name = var.connection_display_name
	display_name_contains = var.connection_display_name_contains
	external_key = var.connection_external_key
	fields = var.connection_fields
	is_default = var.connection_is_default
	state = var.connection_state
	time_created = var.connection_time_created
	time_status_updated = var.connection_time_status_updated
	time_updated = var.connection_time_updated
	updated_by_id = oci_datacatalog_updated_by.test_updated_by.id
}
```

## Argument Reference

The following arguments are supported:

* `catalog_id` - (Required) Unique catalog identifier.
* `created_by_id` - (Optional) OCID of the user who created the resource.
* `data_asset_key` - (Required) Unique data asset key.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given. The match is not case sensitive.
* `display_name_contains` - (Optional) A filter to return only resources that match display name pattern given. The match is not case sensitive. For Example : /folders?displayNameContains=Cu.* The above would match all folders with display name that starts with "Cu" or has the pattern "Cu" anywhere in between. 
* `external_key` - (Optional) Unique external identifier of this resource in the external source system.
* `fields` - (Optional) Specifies the fields to return in a connection summary response. 
* `is_default` - (Optional) Indicates whether this connection is the default connection.
* `state` - (Optional) A filter to return only resources that match the specified lifecycle state. The value is case insensitive.
* `time_created` - (Optional) Time that the resource was created. An [RFC3339](https://tools.ietf.org/html/rfc3339) formatted datetime string.
* `time_status_updated` - (Optional) Time that the resource's status was last updated. An [RFC3339](https://tools.ietf.org/html/rfc3339) formatted datetime string.
* `time_updated` - (Optional) Time that the resource was updated. An [RFC3339](https://tools.ietf.org/html/rfc3339) formatted datetime string.
* `updated_by_id` - (Optional) OCID of the user who updated the resource.


## Attributes Reference

The following attributes are exported:

* `connection_collection` - The list of connection_collection.

### Connection Reference

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

