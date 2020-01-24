---
subcategory: "Data Catalog"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datacatalog_connection"
sidebar_current: "docs-oci-datasource-datacatalog-connection"
description: |-
  Provides details about a specific Connection in Oracle Cloud Infrastructure Data Catalog service
---

# Data Source: oci_datacatalog_connection
This data source provides details about a specific Connection resource in Oracle Cloud Infrastructure Data Catalog service.

Gets a specific data asset connection by key.

## Example Usage

```hcl
data "oci_datacatalog_connection" "test_connection" {
	#Required
	catalog_id = "${oci_datacatalog_catalog.test_catalog.id}"
	connection_key = "${var.connection_connection_key}"
	data_asset_key = "${var.connection_data_asset_key}"

	#Optional
	fields = "${var.connection_fields}"
}
```

## Argument Reference

The following arguments are supported:

* `catalog_id` - (Required) Unique catalog identifier.
* `connection_key` - (Required) Unique connection key.
* `data_asset_key` - (Required) Unique data asset key.
* `fields` - (Optional) Specifies the fields to return in a connection response. 


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

