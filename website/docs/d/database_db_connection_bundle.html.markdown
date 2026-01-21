---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_db_connection_bundle"
sidebar_current: "docs-oci-datasource-database-db_connection_bundle"
description: |-
  Provides details about a specific Db Connection Bundle in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_db_connection_bundle
This data source provides details about a specific Db Connection Bundle resource in Oracle Cloud Infrastructure Database service.

Gets information about the specified database connection bundle.

## Example Usage

```hcl
data "oci_database_db_connection_bundle" "test_db_connection_bundle" {
	#Required
	db_connection_bundle_id = oci_database_db_connection_bundle.test_db_connection_bundle.id
}
```

## Argument Reference

The following arguments are supported:

* `db_connection_bundle_id` - (Required) The OCID of the database connection bundle.


## Attributes Reference

The following attributes are exported:

* `associated_resource_details` - Details about the resources associated with the connection bundle.
	* `resource_ids` - The OCIDs of the associated resources.
* `compartment_id` - The OCID of the compartment containing the database connection bundle.
* `db_connection_bundle_type` - The type of the database connection bundle.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
* `display_name` - Display name for the connection bundle.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
* `id` - The OCID of the database connection bundle.
* `is_protected` - True for the default, service-created Database Connection Bundle.
* `state` - The current lifecycle state of the database connection bundle.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
* `time_created` - The time the database connection bundle was created. An RFC3339 formatted datetime string.
* `time_last_refreshed` - The time the database connection bundle was last refreshed. An RFC3339 formatted datetime string.
* `time_updated` - The time the database connection bundle was updated. An RFC3339 formatted datetime string.

